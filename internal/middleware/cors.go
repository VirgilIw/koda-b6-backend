package middleware

import (
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetHeader() = Dipakai untuk mengambil header yang dikirim client / frontend.
// Header() / Set Header Dipakai server untuk mengirim header ke client.
func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		origin := ctx.GetHeader("Origin")

		allowOrigins := []string{
			os.Getenv("FRONTEND_URL"),
		}

		allowHeaders := []string{
			"Origin",
			"Content-Type",
			"Authorization",
		}

		allowMethods := []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		}

		// contains di situ dipakai untuk mengecek apakah origin dari request ada di daftar origin yang diizinkan.

		if slices.Contains(allowOrigins, origin) {
			ctx.Header("Access-Control-Allow-Origin", origin)
		}

		ctx.Header("Access-Control-Allow-Headers", strings.Join(allowHeaders, ", "))
		ctx.Header("Access-Control-Allow-Methods", strings.Join(allowMethods, ", "))

		// preflight
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.Next()
	}
}
