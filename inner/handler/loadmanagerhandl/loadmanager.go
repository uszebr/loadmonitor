package loadmanagerhandl

import (
	"github.com/gin-gonic/gin"
	"github.com/uszebr/loadmonitor/inner/util/ginutil"
	"github.com/uszebr/loadmonitor/inner/view/loadmanagerview"
)

type LoadManagerHandler struct{}

func (h *LoadManagerHandler) HandlePage(c *gin.Context) {
	_ = ginutil.Render(c, 200, loadmanagerview.LoadManagerPage())
	// TODO: log err here
}
