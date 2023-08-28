package routes

import (
	"github.com/FogonaBabilonia/femboy_uploads/controllers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	address string
	router  *gin.Engine
}

func NewServer(addr string) *Server {
	return &Server{
		address: addr,
		router:  gin.Default(),
	}
}

func (s *Server) setup() {
	s.router.HTMLRender = create_render()
	s.router.Static("/assets/static", "assets/static")

	s.router.GET("/", controllers.HandleRoot)
	s.router.GET("/login", controllers.HandleLogin)
	s.router.GET("/register", controllers.HandleRegister)
	s.router.GET("/public", controllers.HandlePublic)
	s.router.POST("/api/public/upload", controllers.UploadFile)
	s.router.POST("/api/register", controllers.Create_User)
	s.router.POST("/api/login", nil)
}

func (s *Server) Run() {
	s.setup()
	s.router.Run(s.address)
}
