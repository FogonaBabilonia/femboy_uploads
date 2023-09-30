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

	api_routes := s.router.Group("/api")
	user_routes := s.router.Group("/")

	user_routes.Use() // add middleware later on

	user_routes.GET("/", controllers.HandleRootPage)
	user_routes.GET("/login", controllers.HandleLoginPage)
	user_routes.GET("/register", controllers.HandleRegisterPage)
	user_routes.GET("/public", controllers.HandlePublicPage)
	user_routes.GET("/user", controllers.HandleUserPage)

	api_routes.POST("/public/upload", controllers.HandleUploadFile)
	api_routes.POST("/register", controllers.HandleCreate_User)
	api_routes.POST("/login", controllers.HandleLogin)
}

func (s *Server) Run() {
	s.setup()
	s.router.Run(s.address)
}
