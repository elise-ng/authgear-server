package sso

import (
	"net/http"
	"testing"

	"github.com/authgear/authgear-server/pkg/lib/config"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/h2non/gock.v1"
)

func TestAzureadb2cImpl(t *testing.T) {
	Convey("Azureadb2cImpl", t, func() {
		gock.InterceptClient(http.DefaultClient)
		defer gock.Off()

		g := &Azureadb2cImpl{
			RedirectURL: mockRedirectURLProvider{},
			ProviderConfig: config.OAuthSSOProviderConfig{
				ClientID: "client_id",
				Type:     config.OAuthSSOProviderTypeAzureADB2C,
				Tenant:   "tenant",
				Policy:   "policy",
			},
		}

		gock.New("https://tenant.b2clogin.com/tenant.onmicrosoft.com/policy/v2.0/.well-known/openid-configuration").
			Reply(200).
			BodyString(`
{
  "authorization_endpoint": "https://localhost/authorize"
}
			`)
		defer func() { gock.Flush() }()

		u, err := g.GetAuthURL(GetAuthURLParam{
			ResponseMode: ResponseModeFormPost,
			Nonce:        "nonce",
			State:        "state",
			Prompt:       []string{"login"},
		})
		So(err, ShouldBeNil)
		So(u, ShouldEqual, "https://localhost/authorize?client_id=client_id&nonce=nonce&prompt=login&redirect_uri=https%3A%2F%2Flocalhost%2F&response_mode=form_post&response_type=code&scope=openid&state=state")
	})
}
