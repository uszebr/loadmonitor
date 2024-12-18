// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package rtimemonitor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/uszebr/loadmonitor/inner/view/baseview"
	"strconv"
)

// idea TODO: add heap and stack data
type RunTimeData struct {
	NumCpu       int
	NumGorutines int
	GoOs         string
	GoArch       string
	Version      string
	//Mem
	MemAlloc      uint64
	MemTotalAlloc uint64
	MemSys        uint64
	//Gc
	GcCycles    uint32
	GcSys       uint64
	GcNext      uint64
	GcSinceLast float64
}

func RuntimeMonitorPage(rtd RunTimeData) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"page-inner\"><div class=\"page-header\"><h4 class=\"page-title\">Go RunTime Data</h4></div><div class=\"row\"><!-- Card With Icon States Color --><div class=\"row\"><div class=\"col-sm-6 col-md-3\"><div class=\"card card-stats card-round\"><div class=\"card-body\"><div class=\"row\"><div class=\"col-5\"><div class=\"icon-big text-center\"><i class=\"bi bi-airplane-engines text-danger\"></i></div></div><div class=\"col-7 col-stats\"><div class=\"numbers\"><p class=\"card-category\">Os</p><h4 class=\"card-title\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(rtd.GoOs)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `inner/view/rtimemonitor/rtimemonitorview.templ`, Line: 48, Col: 44}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h4></div></div></div></div></div></div><div class=\"col-sm-6 col-md-3\"><div class=\"card card-stats card-round\"><div class=\"card-body\"><div class=\"row\"><div class=\"col-5\"><div class=\"icon-big text-center\"><i class=\"bi bi-buildings text-success\"></i></div></div><div class=\"col-7 col-stats\"><div class=\"numbers\"><p class=\"card-category\">Arch</p><h4 class=\"card-title\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(rtd.GoArch)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `inner/view/rtimemonitor/rtimemonitorview.templ`, Line: 67, Col: 46}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h4></div></div></div></div></div></div><div class=\"col-sm-6 col-md-3\"><div class=\"card card-stats card-round\"><div class=\"card-body\"><div class=\"row\"><div class=\"col-5\"><div class=\"icon-big text-center\"><i class=\"fa fa-microchip\"></i></div></div><div class=\"col-7 col-stats\"><div class=\"numbers\"><p class=\"card-category\">CPU</p><h4 class=\"card-title\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(rtd.NumCpu))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `inner/view/rtimemonitor/rtimemonitorview.templ`, Line: 86, Col: 60}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h4></div></div></div></div></div></div><div class=\"col-sm-6 col-md-3\"><div class=\"card card-stats card-round\"><div class=\"card-body\"><div class=\"row\"><div class=\"col-5\"><div class=\"icon-big text-center\"><i class=\"bi bi-signpost-2 text-secondary\"></i></div></div><div class=\"col-7 col-stats\"><div class=\"numbers\"><p class=\"card-category\">GoRutines</p><h4 class=\"card-title\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(rtd.NumGorutines))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `inner/view/rtimemonitor/rtimemonitorview.templ`, Line: 105, Col: 66}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h4></div></div></div></div></div></div><div class=\"col-sm-6 col-md-3\"><div class=\"card card-stats card-round\"><div class=\"card-body\"><div class=\"row\"><div class=\"col-5\"><div class=\"icon-big text-center\"><i class=\"bi bi-list-columns-reverse text-primary\"></i></div></div><div class=\"col-7 col-stats\"><div class=\"numbers\"><p class=\"card-category\">Runtime</p><h4 class=\"card-title\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(rtd.Version)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `inner/view/rtimemonitor/rtimemonitorview.templ`, Line: 124, Col: 47}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h4></div></div></div></div></div></div><!-- MemStat --><h3 class=\"fw-bold mb-3\">Memory Stat</h3><div class=\"row justify-content-center align-items-center mb-1\"><div class=\"col-md-5 ps-md-0\"><div class=\"card card-pricing\"><div class=\"card-header\"><h4 class=\"card-title\">Memory</h4><div class=\"card-price\"><span class=\"price\"><i class=\"bi bi-memory text-warning\"></i></span></div></div><div class=\"card-body\"><ul class=\"specification-list\"><li><span class=\"name-specification\">MemAlloc</span> <span class=\"status-specification\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(int(rtd.MemAlloc)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `inner/view/rtimemonitor/rtimemonitorview.templ`, Line: 146, Col: 79}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></li><li><span class=\"name-specification\">MemTotalAlloc</span> <span class=\"status-specification\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(int(rtd.MemTotalAlloc)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `inner/view/rtimemonitor/rtimemonitorview.templ`, Line: 150, Col: 84}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></li><li><span class=\"name-specification\">MemSys</span> <span class=\"status-specification\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(int(rtd.MemSys)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `inner/view/rtimemonitor/rtimemonitorview.templ`, Line: 154, Col: 77}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></li></ul></div></div></div><div class=\"col-md-5 pe-md-0\"><div class=\"card card-pricing\"><div class=\"card-header\"><h4 class=\"card-title\">Garbage Collector</h4><div class=\"card-price\"><span class=\"price\"><i class=\"bi bi-recycle text-warning\"></i></span></div></div><div class=\"card-body\"><ul class=\"specification-list\"><li><span class=\"name-specification\">Cycles</span> <span class=\"status-specification\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(int(rtd.GcCycles)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `inner/view/rtimemonitor/rtimemonitorview.templ`, Line: 172, Col: 79}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></li><li><span class=\"name-specification\">Sys</span> <span class=\"status-specification\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(int(rtd.GcSys)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `inner/view/rtimemonitor/rtimemonitorview.templ`, Line: 176, Col: 76}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></li><li><span class=\"name-specification\">Next</span> <span class=\"status-specification\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(int(rtd.GcNext)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `inner/view/rtimemonitor/rtimemonitorview.templ`, Line: 180, Col: 77}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></li><li><span class=\"name-specification\">Since Last</span> <span class=\"status-specification\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(int(rtd.GcSinceLast)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `inner/view/rtimemonitor/rtimemonitorview.templ`, Line: 184, Col: 82}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></li></ul></div></div></div></div></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = baseview.BasePage(baseview.BaseParam{Title: "Go Runtime data", HTMX: false}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
