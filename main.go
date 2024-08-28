package main

import (
	"go-base/component/appctx"
	"go-base/middleware"
	authmodel "go-base/module/auth/model"
	authtransport "go-base/module/auth/transport"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func loadEnv() {
	env := os.Getenv("APP_ENV")
	if env == "test" {
		err := godotenv.Load(".env.test")
		if err != nil {
			log.Fatalf("Error loading .env.test file")
		}
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}
}
func main() { 
	loadEnv()
	dsn:=os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&authmodel.User{})
	if err!=nil {
		log.Fatalf("cannot migrate database")
	}
	db = db.Debug()
	//router http
	appContext:=appctx.NewAppContext(db, os.Getenv("SECRET_KEY"))
	router:=gin.Default()
	v1:=router.Group("/v1")
	setUpRouter(v1,appContext)
	router.Run()
	
}
func setUpRouter(v1 *gin.RouterGroup,appContext  appctx.AppContext) {
	v1=v1.Group("/",middleware.Recover(appContext))
	v1.POST("/register",authtransport.Register(appContext))
	v1.POST("/login",authtransport.Login(appContext))
	v1.POST("/generate-token",authtransport.GenerateToken(appContext))
	v1.GET("/ping",middleware.RequireAuth(appContext),func(c *gin.Context){
		c.JSON(200,gin.H{"message":"pong"})
	})
}

