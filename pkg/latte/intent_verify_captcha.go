package latte

import (
	"context"

	"github.com/authgear/authgear-server/pkg/lib/workflow"
	"github.com/authgear/authgear-server/pkg/util/validation"
)

func init() {
	workflow.RegisterPrivateIntent(&IntentVerifyCaptcha{})
}

var IntentVerifyCaptchaSchema = validation.NewSimpleSchema(`{}`)

type IntentVerifyCaptcha struct {
}

func (*IntentVerifyCaptcha) Kind() string {
	return "latte.IntentVerifyCaptcha"
}

func (*IntentVerifyCaptcha) JSONSchema() *validation.SimpleSchema {
	return IntentVerifyIdentitySchema
}

func (*IntentVerifyCaptcha) CanReactTo(ctx context.Context, deps *workflow.Dependencies, w *workflow.Workflow) ([]workflow.Input, error) {
	if len(w.Nodes) == 0 {
		return []workflow.Input{
			&InputTakeCaptchaToken{},
		}, nil
	}
	return nil, workflow.ErrEOF
}

func (i *IntentVerifyCaptcha) ReactTo(ctx context.Context, deps *workflow.Dependencies, w *workflow.Workflow, input workflow.Input) (*workflow.Node, error) {
	var inputTakeCaptchaToken inputTakeCaptchaToken

	switch {
	case workflow.AsInput(input, &inputTakeCaptchaToken):
		// FIXME(tung): Verify the token
		// token := inputTakeCaptchaToken.GetToken()
		node := NodeVerifiedCaptcha{}
		return workflow.NewNodeSimple(&node), nil
	default:
		return nil, workflow.ErrIncompatibleInput
	}
}

func (*IntentVerifyCaptcha) GetEffects(ctx context.Context, deps *workflow.Dependencies, w *workflow.Workflow) (effs []workflow.Effect, err error) {
	return nil, nil
}

func (i *IntentVerifyCaptcha) OutputData(ctx context.Context, deps *workflow.Dependencies, w *workflow.Workflow) (interface{}, error) {
	return nil, nil
}
