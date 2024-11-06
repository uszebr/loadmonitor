package ginutil

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

// video about gin/ https://github.com/matheusgomes28/urchin/blob/main/migrations/20240115222144_add_posts_table.sql
// Render util for easier handling Templ components with Gin
func Render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}
