package main

import (
	"fmt"

	"app/src/inits"
)

func main() {
	fmt.Println("Project started")
	// uncomment to turn off stdout
	//gin.SetMode(gin.ReleaseMode)
	// gin.DefaultWriter = ioutil.Discard
	inits.Init()
	// services.InitServices()
	// store.InitDb()
	// defer store.CloseDb()
	// routes.InitRoutes()
}
