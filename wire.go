//go:build wireinject
// +build wireinject

package wireExp

import (
	"github.com/google/wire"
	"wireExp/internal/api/handler"
	"wireExp/internal/api/repository"
	"wireExp/internal/api/router"
	"wireExp/internal/api/service"
	"wireExp/internal/clients"
)

type Server struct {
	Router *router.Router
}

func NewServer() *Server {
	wire.Build(
		wire.Struct(new(Server), "*"),
		router.NewRouter,
		handler.HandlerWireSet,
		service.ServiceWireSet,
		repository.RepoWireSet,
		clients.ClientsWireSet,
	)
	return &Server{}
}