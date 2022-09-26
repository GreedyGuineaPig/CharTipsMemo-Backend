package main

import (
	"gomysql-api/router"
)

func main() {
	rt := router.Init()
	rt.Run()
}
