package webapp

import (
	"fmt"
	"net/http"

	"github.com/authgear/authgear-server/pkg/auth/handler/webapp/viewmodels"
	"github.com/authgear/authgear-server/pkg/auth/webapp"
	"github.com/authgear/authgear-server/pkg/lib/authn"
	"github.com/authgear/authgear-server/pkg/lib/infra/mail"
	"github.com/authgear/authgear-server/pkg/lib/interaction"
	"github.com/authgear/authgear-server/pkg/lib/interaction/nodes"
	"github.com/authgear/authgear-server/pkg/util/httproute"
	"github.com/authgear/authgear-server/pkg/util/phone"
	"github.com/authgear/authgear-server/pkg/util/template"
	"github.com/authgear/authgear-server/pkg/util/validation"
)

var TemplateWebEnterOOBOTPHTML = template.RegisterHTML(
	"web/enter_oob_otp.html",
	components...,
)

var EnterOOBOTPSchema = validation.NewSimpleSchema(`
	{
		"type": "object",
		"properties": {
			"x_code": { "type": "string" }
		},
		"required": ["x_code"]
	}
`)

func ConfigureEnterOOBOTPRoute(route httproute.Route) httproute.Route {
	return route.
		WithMethods("OPTIONS", "POST", "GET").
		WithPathPattern("/enter_oob_otp")
}

type EnterOOBOTPViewModel struct {
	OOBOTPTarget           string
	OOBOTPCodeSendCooldown int
	OOBOTPCodeLength       int
	OOBOTPChannel          string
}

type EnterOOBOTPHandler struct {
	ControllerFactory ControllerFactory
	BaseViewModel     *viewmodels.BaseViewModeler
	Renderer          Renderer
}

type EnterOOBOTPNode interface {
	GetOOBOTPTarget() string
	GetOOBOTPChannel() string
	GetOOBOTPCodeSendCooldown() int
	GetOOBOTPCodeLength() int
}

func (h *EnterOOBOTPHandler) GetData(r *http.Request, rw http.ResponseWriter, session *webapp.Session, graph *interaction.Graph) (map[string]interface{}, error) {
	data := map[string]interface{}{}

	baseViewModel := h.BaseViewModel.ViewModel(r, rw)
	viewModel := EnterOOBOTPViewModel{}
	var n EnterOOBOTPNode
	if graph.FindLastNode(&n) {
		viewModel.OOBOTPCodeSendCooldown = n.GetOOBOTPCodeSendCooldown()
		viewModel.OOBOTPCodeLength = n.GetOOBOTPCodeLength()
		viewModel.OOBOTPChannel = n.GetOOBOTPChannel()

		switch authn.AuthenticatorOOBChannel(viewModel.OOBOTPChannel) {
		case authn.AuthenticatorOOBChannelEmail:
			viewModel.OOBOTPTarget = mail.MaskAddress(n.GetOOBOTPTarget())
		case authn.AuthenticatorOOBChannelSMS:
			viewModel.OOBOTPTarget = phone.Mask(n.GetOOBOTPTarget())
		}
	}

	currentNode := graph.CurrentNode()
	alternatives := viewmodels.AlternativeStepsViewModel{}
	switch currentNode.(type) {
	case *nodes.NodeAuthenticationOOBTrigger:
		err := alternatives.AddAuthenticationAlternatives(graph, webapp.SessionStepEnterOOBOTPAuthn)
		if err != nil {
			return nil, err
		}
	case *nodes.NodeCreateAuthenticatorOOBSetup:
		err := alternatives.AddCreateAuthenticatorAlternatives(graph, webapp.SessionStepEnterOOBOTPSetup)
		if err != nil {
			return nil, err
		}
	default:
		panic(fmt.Errorf("enter_oob_otp: unexpected node: %T", currentNode))
	}

	viewmodels.Embed(data, baseViewModel)
	viewmodels.Embed(data, viewModel)
	viewmodels.Embed(data, alternatives)

	return data, nil
}

func (h *EnterOOBOTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctrl, err := h.ControllerFactory.New(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer ctrl.Serve()

	ctrl.Get(func() error {
		session, err := ctrl.InteractionSession()
		if err != nil {
			return err
		}

		graph, err := ctrl.InteractionGet()
		if err != nil {
			return err
		}

		data, err := h.GetData(r, w, session, graph)
		if err != nil {
			return err
		}

		h.Renderer.RenderHTML(w, r, TemplateWebEnterOOBOTPHTML, data)
		return nil
	})

	ctrl.PostAction("resend", func() error {
		result, err := ctrl.InteractionPost(func() (input interface{}, err error) {
			input = &InputResendCode{}
			return
		})
		if err != nil {
			return err
		}

		result.WriteResponse(w, r)
		return nil
	})

	ctrl.PostAction("submit", func() error {
		result, err := ctrl.InteractionPost(func() (input interface{}, err error) {
			err = EnterOOBOTPSchema.Validator().ValidateValue(FormToJSON(r.Form))
			if err != nil {
				return
			}

			code := r.Form.Get("x_code")
			deviceToken := r.Form.Get("x_device_token") == "true"

			input = &InputAuthOOB{
				Code:        code,
				DeviceToken: deviceToken,
			}
			return
		})
		if err != nil {
			return err
		}

		result.WriteResponse(w, r)
		return nil
	})

	handleAlternativeSteps(ctrl)
}
