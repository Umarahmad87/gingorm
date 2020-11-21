package inits

import (
	"app/src/controllers"
	"app/src/models"
	"app/src/store"
	"app/src/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/namsral/flag"
	"github.com/gin-contrib/gzip"
)

// database variables
var (
	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	Port             string
)

// InitRoutes initializes all the routes
func InitRoutes() {
	router := gin.Default()

	router.Use(utils.CORSMiddleware())
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	apiRouter := router.Group("/api")
	{
		/*** START USER ***/
		user := new(controllers.UserController)

		apiRouter.POST("/user/login", user.Login)
		// apiRouter.GET("/user/logout", user.Logout)

	}

	// router.LoadHTMLGlob("./public/html/*")

	// router.Static("/public", "./public")

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,  gin.H{
			"ginGormBoilerplateVersion": "v1.0",
			"goVersion":             runtime.Version(),
		})
	})

	/*router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})*/

	router.Run(":" + Port)

}

// InitDb initializes database connection
func InitDb() {

	conString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		DatabaseHost, DatabasePort, DatabaseUser, DatabaseName, DatabasePassword)

	store.OpenDb(conString)
	store.Db.AutoMigrate(&models.User{})
	seed()
}

func seed() {
	log.Println("Database seed")
	seedUser()
}

func seedUser() {
	var user = &models.User{Email: "admin@gmail.com", Name: "admin", Password: "admin" }
	models.UserManager.CreateUser(user)
}


// Init initializes env variables and starts the server
func Init() {

	//database variables
	flag.StringVar(&DatabaseHost, "DATABASE_HOST", "localhost", "Database Host")
	flag.StringVar(&DatabasePort, "DATABASE_PORT", "5432", "Database Port")
	flag.StringVar(&DatabaseUser, "DATABASE_USER", "postgres", "Database User")
	flag.StringVar(&DatabasePassword, "DATABASE_PASSWORD", "umar", "Database Password")
	flag.StringVar(&DatabaseName, "DATABASE_NAME", "postgres", "Database Name")

	// server port
	flag.StringVar(&Port, "PORT", "3000", "Port")

	flag.Parse()
	fmt.Printf("############### Database Host:%s \n", DatabaseHost)

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	InitDb()
	InitRoutes()

}
