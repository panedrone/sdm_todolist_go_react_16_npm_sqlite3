package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"runtime"
	"sdm_demo_todolist/shared"
	"sdm_demo_todolist/sqlx/dbal"
	"sdm_demo_todolist/sqlx/handlers"
	"sdm_demo_todolist/sqlx/swagger"
)

// @schemes	http
// @produce	json
// @version	0.0.1
//
// @title		SDM TodoApp API
// @in			header
// @BasePath	/api
// @accept		json
// @host		127.0.0.1:8080
func main() {

	err := dbal.OpenDB()
	if err != nil {
		println(err.Error())
		return
	}
	defer func() {
		_ = dbal.CloseDB()
	}()

	gin.SetMode(gin.ReleaseMode)

	myRouter := gin.New()

	swagger.Init(myRouter)

	myOS, myArch := runtime.GOOS, runtime.GOARCH
	inContainer := " docker,"
	if _, err := os.Lstat("/.dockerenv"); err != nil && os.IsNotExist(err) {
		inContainer = ""
	}

	// whoIam := fmt.Sprintf(`%v, %v,%v sqlx, sqlite3, <a target="_blank" href="swagger/index.html">swagger</a>`, myOS, myArch, inContainer)
	whoIam := fmt.Sprintf(`%v, %v,%v sqlx, sqlite3`, myOS, myArch, inContainer)

	shared.AssignHandlers(myRouter, whoIam, handlers.NewProjectHandlers(), handlers.NewTaskHandlers())

	log.Fatal(myRouter.Run(":8080"))
}
