package server

import (
	"path/filepath"

	"fmt"

	"github.com/caiwp/impala-web/setting"
	"github.com/gin-contrib/multitemplate"
)

var (
	tplLogin = "login"
	tplIndex = "index"
	tplQuery = "query"

	tplBase   = "base"
	tplLayout = "layout"
)

func CreateMyRender() multitemplate.Render {
	r := multitemplate.New()

	tplBase = getHTML(tplBase)
	tplLayout = getHTML(tplLayout)

	r.AddFromFiles(tplLogin, getHTML(tplLogin))
	r.AddFromFiles(tplIndex, tplBase, tplLayout, getHTML(tplIndex))
	r.AddFromFiles(tplQuery, tplBase, tplLayout, getHTML(tplQuery))

	return r
}

func getHTML(s string) string {
	t := fmt.Sprintf("template/%s.html", s)
	return filepath.Join(setting.RootPath, t)
}
