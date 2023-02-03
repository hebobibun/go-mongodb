package main

import (
	"context"
	"go-mongodb/config"
)

func main() {
	clt := config.MgConnect()
	defer clt.Disconnect(context.TODO())
}

