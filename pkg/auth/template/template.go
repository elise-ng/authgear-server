package template

import (
	"github.com/authgear/authgear-server/pkg/auth/config"
	"github.com/authgear/authgear-server/pkg/auth/dependency/forgotpassword"
	"github.com/authgear/authgear-server/pkg/auth/dependency/welcomemessage"
	"github.com/authgear/authgear-server/pkg/auth/handler/webapp"
	"github.com/authgear/authgear-server/pkg/otp"
	"github.com/authgear/authgear-server/pkg/template"
)

func NewEngineWithConfig(
	serverConfig *config.ServerConfig,
	c *config.Config,
) *template.Engine {
	e := template.NewEngine(template.NewEngineOptions{
		DefaultTemplatesDirectory: serverConfig.DefaultTemplateDirectory,
		TemplateItems:             c.AppConfig.Template.Items,
		FallbackLanguage:          c.AppConfig.Localization.FallbackLanguage,
	})

	e.Register(welcomemessage.TemplateWelcomeEmailTXT)
	e.Register(welcomemessage.TemplateWelcomeEmailHTML)

	e.Register(otp.TemplateOTPMessageSMSTXT)
	e.Register(otp.TemplateOTPMessageEmailTXT)
	e.Register(otp.TemplateOTPMessageEmailHTML)

	// Auth UI
	e.Register(webapp.TemplateAuthUITranslationJSON)

	e.Register(webapp.TemplateAuthUIHTMLHeadHTML)
	e.Register(webapp.TemplateAuthUIHeaderHTML)
	e.Register(webapp.TemplateAuthUIFooterHTML)

	e.Register(webapp.TemplateAuthUILoginHTML)
	e.Register(webapp.TemplateAuthUISignupHTML)
	e.Register(webapp.TemplateAuthUIPromoteHTML)

	e.Register(webapp.TemplateAuthUIEnterPasswordHTML)
	e.Register(webapp.TemplateAuthUICreatePasswordHTML)
	e.Register(webapp.TemplateAuthUIOOBOTPHTML)
	e.Register(webapp.TemplateAuthUIEnterLoginIDHTML)

	e.Register(webapp.TemplateAuthUIForgotPasswordHTML)
	e.Register(webapp.TemplateAuthUIForgotPasswordSuccessHTML)
	e.Register(webapp.TemplateAuthUIResetPasswordHTML)
	e.Register(webapp.TemplateAuthUIResetPasswordSuccessHTML)
	e.Register(webapp.TemplateAuthUILogoutHTML)

	e.Register(webapp.TemplateAuthUISettingsHTML)
	e.Register(webapp.TemplateAuthUISettingsIdentityHTML)

	e.Register(forgotpassword.TemplateForgotPasswordEmailTXT)
	e.Register(forgotpassword.TemplateForgotPasswordEmailHTML)
	e.Register(forgotpassword.TemplateForgotPasswordSMSTXT)

	return e
}
