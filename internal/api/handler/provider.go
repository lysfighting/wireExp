package handler

import "github.com/google/wire"

var HandlerWireSet = wire.NewSet(
	NewUserHandler,
)