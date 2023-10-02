package routes

import "github.com/gin-contrib/multitemplate"

func create_render() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r.AddFromFiles("index", "assets/templates/index.html", "assets/templates/root.html")
	r.AddFromFiles("public", "assets/templates/public.html", "assets/templates/root.html")

	return r
}
