package main

import (
	"mekar/master"
	"mekar/master/config"
)

func main() {
	db := config.EnvConn()
	useActivityLog := config.AuthSwitch()
	router := config.CreateRouter()
	master.InitAll(router, db, useActivityLog)
	config.RunServer(router)
}
