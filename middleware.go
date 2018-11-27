package precompiler

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// GinMiddleware is used to set the cache-control header for all files under the specified asset path (including leading slash)
func GinMiddleware(assetPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, assetPath) {
			c.Header("cache-control", "max-age=315360000; public")
		}
	}
}
