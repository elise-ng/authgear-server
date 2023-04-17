package service

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	texttemplate "text/template"

	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/config/configsource"
	portalconfig "github.com/authgear/authgear-server/pkg/portal/config"
)

type AuthzAdder interface {
	AddAuthz(
		auth config.AdminAPIAuth,
		appID config.AppID,
		authKey *config.AdminAPIAuthKey,
		auditContext interface{},
		hdr http.Header) (err error)
}

type AdminAPIService struct {
	AuthgearConfig *portalconfig.AuthgearConfig
	AdminAPIConfig *portalconfig.AdminAPIConfig
	ConfigSource   *configsource.ConfigSource
	AuthzAdder     AuthzAdder
}

type AuthzAuditContext struct {
	Triggerer   string `json:"triggerer"`
	ActorUserID string `json:"actor_user_id"`
	Referrer    string `json:"refererer"`
}

func (s *AdminAPIService) ResolveConfig(appID string) (*config.Config, error) {
	appCtx, err := s.ConfigSource.ContextResolver.ResolveContext(appID)
	if err != nil {
		return nil, err
	}
	return appCtx.Config, nil
}

func (s *AdminAPIService) ResolveHost(appID string) (host string, err error) {
	t := texttemplate.New("host-template")
	_, err = t.Parse(s.AdminAPIConfig.HostTemplate)
	if err != nil {
		return
	}
	var buf strings.Builder

	data := map[string]interface{}{
		"AppID": appID,
	}
	err = t.Execute(&buf, data)
	if err != nil {
		return
	}

	host = buf.String()
	return
}

func (s *AdminAPIService) ResolveEndpoint(appID string) (*url.URL, error) {
	switch s.AdminAPIConfig.Type {
	case portalconfig.AdminAPITypeStatic:
		endpoint, err := url.Parse(s.AdminAPIConfig.Endpoint)
		if err != nil {
			return nil, err
		}
		return endpoint, nil
	default:
		panic(fmt.Errorf("portal: unexpected admin API type: %v", s.AdminAPIConfig.Type))
	}
}

func (s *AdminAPIService) Director(appID string, p string, actorUserID string) (director func(*http.Request), err error) {
	cfg, err := s.ResolveConfig(appID)
	if err != nil {
		return
	}

	authKey, ok := cfg.SecretConfig.LookupData(config.AdminAPIAuthKeyKey).(*config.AdminAPIAuthKey)
	if !ok {
		err = fmt.Errorf("failed to look up admin API auth key: %v", appID)
		return
	}

	endpoint, err := s.ResolveEndpoint(appID)
	if err != nil {
		return
	}
	endpoint.Path = p

	host, err := s.ResolveHost(appID)
	if err != nil {
		return
	}

	director = func(r *http.Request) {
		// It is important to preserve raw query so that GraphiQL ?query=... is not broken.
		rawQuery := r.URL.RawQuery
		r.URL = endpoint
		r.URL.RawQuery = rawQuery
		r.Host = host
		r.Header.Set("X-Forwarded-Host", r.Host)

		err = s.AuthzAdder.AddAuthz(
			s.AdminAPIConfig.Auth,
			config.AppID(appID),
			authKey,
			AuthzAuditContext{
				Triggerer:   "portal",
				ActorUserID: actorUserID,
				Referrer:    r.Header.Get("Referer"),
			},
			r.Header,
		)
		if err != nil {
			panic(err)
		}
	}
	return
}

func (s *AdminAPIService) SelfDirector(actorUserID string) (director func(*http.Request), err error) {
	return s.Director(s.AuthgearConfig.AppID, "/graphql", actorUserID)
}
