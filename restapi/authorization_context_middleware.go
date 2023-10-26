package restapi

import (
	"Parallels/pd-api-service/basecontext"
	"Parallels/pd-api-service/constants"
	"context"
	"net/http"
)

func AddAuthorizationContextMiddlewareAdapter() Adapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := r.Context().Value(REQUEST_ID_KEY)
			authorizationContext := basecontext.CloneAuthorizationContext()

			// Adding the request id if it exist
			if id != nil {
				authorizationContext.RequestId = id.(string)
			}

			// Adding a new Authorization Request to the Request
			ctx := context.WithValue(r.Context(), constants.AUTHORIZATION_CONTEXT_KEY, authorizationContext)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
