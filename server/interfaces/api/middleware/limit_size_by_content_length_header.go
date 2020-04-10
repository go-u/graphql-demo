package middleware

import (
	"net/http"
	"server/etc/config"
)

func LimitSizeByContentLengthHeader(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength > config.MaxRequestSize {
			_ = BuildErrorResponse(w, "content length over")
			return // http.HandlerFunc(fn) を返さずに処理を終了する
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
