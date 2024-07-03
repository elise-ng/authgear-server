package botprotection

import (
	"strings"

	"github.com/authgear/authgear-server/pkg/util/slice"
)

// raw API response from recaptchav2, which is not visible to consumer of this API client
type RecaptchaV2RawResponse struct {
	Success *bool `json:"success,omitempty"`

	// specific to Success == false
	ErrorCodes []RecaptchaV2ErrorCode `json:"error-codes"`

	// specific to Success == true
	ChallengeTs string `json:"challenge_ts"` // timestamp of the challenge load (ISO format yyyy-MM-dd'T'HH:mm:ssZZ)
	Hostname    string `json:"hostname"`     // hostname for which the challenge was served.
}

// parsed success response
type RecaptchaV2SuccessResponse struct {
	ChallengeTs string `json:"challenge_ts"` // timestamp of the challenge load (ISO format yyyy-MM-dd'T'HH:mm:ssZZ)
	Hostname    string `json:"hostname"`     // hostname for which the challenge was served.
}

// parsed error response
type RecaptchaV2ErrorResponse struct {
	ErrorCodes []RecaptchaV2ErrorCode `json:"error-codes"`
}

// returns a comma separated string of error codes
func (e *RecaptchaV2ErrorResponse) Error() string {
	errorCodeStrings := slice.Map(e.ErrorCodes, func(c RecaptchaV2ErrorCode) string {
		return string(c)
	})

	return strings.Join(errorCodeStrings, ",")
}

type RecaptchaV2ErrorCode string

const (
	RecaptchaV2ErrorCodeMissingInputSecret   RecaptchaV2ErrorCode = "missing-input-secret"
	RecaptchaV2ErrorCodeInvalidInputSecret   RecaptchaV2ErrorCode = "invalid-input-secret"
	RecaptchaV2ErrorCodeMissingInputResponse RecaptchaV2ErrorCode = "missing-input-response"
	RecaptchaV2ErrorCodeInvalidInputResponse RecaptchaV2ErrorCode = "invalid-input-response"
	RecaptchaV2ErrorCodeBadRequest           RecaptchaV2ErrorCode = "bad-request"
	RecaptchaV2ErrorCodeTimeoutOrDuplicate   RecaptchaV2ErrorCode = "timeout-or-duplicate"
)
