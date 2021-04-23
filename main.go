package main

import (
	"gin-demo/Databases"
	"gin-demo/Router"
)

func main() {
	defer Databases.Db.Close()
	Router.InitRouter()
}
