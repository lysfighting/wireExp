package clients

import "github.com/google/wire"

var ClientsWireSet = wire.NewSet(
	NewExpDB,
)