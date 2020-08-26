package controller

import "github.com/ashishb26/rzpbool/middleware"

// InitRoutes initializes the routes for the Router in the Server struct
func (s *Server) InitRoutes() {
	api := s.Router.Group("/api")
	api.Use(middleware.AuthUser())
	{
		api.POST("/", s.AddBool)
		api.GET("/:id", s.GetBool)
		api.PATCH("/:id", s.UpdateBool)
		api.DELETE("/:id", s.DeleteBool)
	}

	auth := s.Router.Group("/user")
	{
		auth.POST("/login", s.UserLogin)
	}
}
