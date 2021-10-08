package main

import (
	"github.com/mirzaakhena/zapp2/application"
	"github.com/mirzaakhena/zapp2/application/registry"
)

func main() {

	application.Run(registry.NewZapp()())

}
