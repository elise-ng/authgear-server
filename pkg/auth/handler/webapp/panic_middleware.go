package webapp

import (
	"net/http"

	"github.com/felixge/httpsnoop"

	"github.com/authgear/authgear-server/pkg/api/apierrors"
	"github.com/authgear/authgear-server/pkg/auth/handler/webapp/viewmodels"
	"github.com/authgear/authgear-server/pkg/auth/webapp"
	"github.com/authgear/authgear-server/pkg/util/log"
	"github.com/authgear/authgear-server/pkg/util/panicutil"
)

type PanicMiddlewareLogger struct{ *log.Logger }

func NewPanicMiddlewareLogger(lf *log.Factory) PanicMiddlewareLogger {
	return PanicMiddlewareLogger{lf.New("webapp-panic-middleware")}
}

type PanicMiddleware struct {
	ErrorCookie   *webapp.ErrorCookie
	Logger        PanicMiddlewareLogger
	BaseViewModel *viewmodels.BaseViewModeler
	Renderer      Renderer
}

func (m *PanicMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		written := false

		w = httpsnoop.Wrap(w, httpsnoop.Hooks{
			WriteHeader: func(f httpsnoop.WriteHeaderFunc) httpsnoop.WriteHeaderFunc {
				return func(code int) {
					written = true
					f(code)
				}
			},
			Write: func(f httpsnoop.WriteFunc) httpsnoop.WriteFunc {
				return func(b []byte) (int, error) {
					written = true
					return f(b)
				}
			},
		})

		defer func() {
			if e := recover(); e != nil {
				err := panicutil.MakeError(e)
				m.Logger.WithError(err).Error("panic occurred")

				apiError := apierrors.AsAPIError(err)

				if !written {
					// Render the HTML directly and DO NOT redirect.
					// If we redirect to the original URL, then GET request will result in infinite redirect.
					// See https://github.com/authgear/authgear-server/issues/3509

					cookie, cookieErr := m.ErrorCookie.SetRecoverableError(r, apiError)
					if cookieErr != nil {
						panic(cookieErr)
					}

					r.AddCookie(cookie)

					data := make(map[string]interface{})
					baseViewModel := m.BaseViewModel.ViewModel(r, w)
					viewmodels.Embed(data, baseViewModel)

					m.Renderer.RenderHTML(w, r, TemplateWebFatalErrorHTML, data)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}
