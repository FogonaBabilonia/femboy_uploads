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

	s.router.GET("/", controllers.HandleRootPage)
	s.router.GET("/login", controllers.HandleLoginPage)
	s.router.GET("/register", controllers.HandleRegisterPage)
	s.router.GET("/public", controllers.HandlePublicPage)
	s.router.GET("/user", controllers.HandleUserPage)
	s.router.POST("/api/public/upload", controllers.HandleUploadFile)
	s.router.POST("/api/register", controllers.HandleCreate_User)
	s.router.POST("/api/login", controllers.HandleLogin)
}

func (s *Server) Run() {
	s.setup()
	s.router.Run(s.address)
}
