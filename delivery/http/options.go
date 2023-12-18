package http

type ServerOption func(server *Server)

func Address(addr string) ServerOption {
	return func(server *Server) {
		server.address = addr
	}
}
