package routes

import (
	"project/pasarwarga/controllers"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {

	//category
	e.GET("/categories", controllers.GetAllCategories)
	e.POST("/category", controllers.CreateCategory)
	e.GET("/category/:category_slug", controllers.GetDetailCategory)
	e.PUT("/category", controllers.EditCategory)
	e.DELETE("/category", controllers.DeleteCategory)

	//article
	e.GET("/articles", controllers.GetAllArticles)
	e.POST("/article", controllers.CreateArticle)
	e.GET("/article/:slug", controllers.GetDetailArticle)
	e.PUT("/article", controllers.EditArticle)
	e.DELETE("/article", controllers.DeleteArticle)
}
