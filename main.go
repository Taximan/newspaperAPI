package main

import "github.com/gin-gonic/gin"
import "github.com/joho/godotenv"
import "os"

//SECRET used to token encrpytion
var SECRET = ""

func main() {
  if err := godotenv.Load(); err != nil {
    panic(err.Error())
  }
  
  SECRET = os.Getenv("JWT_SECRET")
  
  InitDB(os.Getenv("PG_CONNECTION_STRING"))
  
  app := gin.Default()
  
  auth := app.Group("/auth")
  {
    auth.POST("/register", PostAuthRegister)
    auth.POST("/login", PostAuthLogin)
  }
  
  api := app.Group("/api")
  {
    article := api.Group("/article")
    {
      article.POST("/", Auth, PostAPIArticle)
      article.GET("/:id", GetAPIArticle)
      article.PUT("/:id", PutAPIArticle)
      article.DELETE("/:id", Auth, DeleteAPIArticle)
    }
  }
  
  app.Run()
  
}