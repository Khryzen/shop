package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	context := map[string]interface{}{}

	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")
	page := strings.TrimSuffix(r.URL.Path, "/")

	switch page {
	case "index":
		context = DashboardHandler(w, r)
	default:
		page = "index"
		context = DashboardHandler(w, r)

	}
	context["Page"] = strings.Title(page)
	InterfaceRender(w, r, page, context)

}

func InterfaceRender(w http.ResponseWriter, r *http.Request, tpl string, context map[string]interface{}) {
	templateList := []string{}
	templateList = append(templateList, "./templates/client/base.html")

	path := "./templates/client/" + tpl + ".html"
	templateList = append(templateList, path)

	uadmin.RenderMultiHTML(w, r, templateList, context)
}
