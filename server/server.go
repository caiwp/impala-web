package server

import (
	"github.com/gin-contrib/multitemplate"
)

const (
	tplLogin string = "login"
	tplIndex string = "index"
	tplQuery string = "query"

	tplBase   string = "template/base.html"
	tplLayout string = "template/layout.html"
)

func CreateMyRender() multitemplate.Render {
	r := multitemplate.New()

	r.AddFromFiles(tplLogin, "template/login.html")
	r.AddFromFiles(tplIndex, tplBase, tplLayout, "template/index.html")
	r.AddFromFiles(tplQuery, tplBase, tplLayout, "template/query.html")

	return r
}
