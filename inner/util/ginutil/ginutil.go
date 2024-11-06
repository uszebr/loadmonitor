package ginutil

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

// util for easier handling Templ components with Gin
func Render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}
