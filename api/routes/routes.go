package routes

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewReceiptRoutes),
	fx.Provide(NewRoutes),
	
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	receiptRoutes ReceiptRoutes,
) Routes {
	return Routes{
		receiptRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}