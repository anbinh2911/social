package server

func (s *Server) RegisterRoutes() {

	userGroup := s.Group("/users")
	userGroup.Post("/v1/regsiter", s.registerUserHandler)
	userGroup.Post("/v1/login", s.loginHandler)
	userGroup.Post("/v1/forgot", s.forgotPasswordHandler)
	userGroup.Get("/v1/detail", s.userDetailHandler)

}
