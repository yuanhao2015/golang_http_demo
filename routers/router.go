package routers

import (
	"../controllers"
	"../utils"
	"io"
	"net/http"
	"regexp"
	//"strings"
	"fmt"
	"time"
)

type WebController struct {
	Function func(http.ResponseWriter, *http.Request)
	Method   string
	Pattern  string
}

var mux []WebController

func init() {
	mux = append(mux, WebController{controllers.Index, "GET", "^/$"})
	mux = append(mux, WebController{controllers.MemberList, "GET", "^/memberlist"})
	mux = append(mux, WebController{controllers.MemberAdd, "POST", "^/memberadd"})
	mux = append(mux, WebController{controllers.MemberDelete, "GET", "^/memberdel"})
	mux = append(mux, WebController{controllers.MemberGet, "GET", "^/memberget"})

}

type HttpHandler struct{}

func (*HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if n, _ := regexp.MatchString("^/static", r.URL.Path); n {
		Static(w, r)
		return
	}
	t := time.Now()
	for _, webController := range mux {
		if m, _ := regexp.MatchString(webController.Pattern, r.URL.Path); m {
			if r.Method == webController.Method {
				webController.Function(w, r)
				go utils.WriteLog(r, t, "match", webController.Pattern)
				return
			}
		}
	}
	go utils.WriteLog(r, t, "unmatch", "")
	io.WriteString(w, "找不到路由")
	return
}

//添加静态资源
func Static(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deal Static: ", r.URL.Path)
	http.ServeFile(w, r, "."+r.URL.Path)
}
