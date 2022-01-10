package repository

import "github.com/google/wire"

var RepoWireSet = wire.NewSet(
	NewUserRepo,
	)