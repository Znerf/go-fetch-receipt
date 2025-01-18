package main

import (
	"receipt/bootstrap"
	"go.uber.org/fx"
)


func main() {
	fx.New(bootstrap.Module).Run()
}