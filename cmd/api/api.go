package main

import (
	"adm/api/pkg/api"
	"adm/api/pkg/utls/config"
)

func main() {
	c := config.NewConf()

	api.Start(c)
}
