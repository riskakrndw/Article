package controllers

import (
	"net/http"
	"project/pasarwarga/databases"
	"project/pasarwarga/models"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type OutputArticle struct {
	ID         uint      `json:"id"`
	CategoryID uint      `json:"category_id"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ListArticle struct {
	ID       uint   `json:"id"`
	Category string `json:"category_name"`
	Title    string `json:"title"`
	Slug     string `json:"slug"`
}

func CreateArticle(c echo.Context) error {
	//get user's input
	input_user := models.Article{}
	c.Bind(&input_user)

	//check is data nil?
	if input_user.CategoryID == 0 || input_user.Title == "" || input_user.Slug == "" || input_user.Content == "" {
		return c.JSON(http.StatusBadRequest, "Please fill all data")
	}

	//check category
	_, err := databases.GetCategoryById(int(input_user.CategoryID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot find category")
	}

	//create article
	article, err := databases.CreateArticle(input_user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot insert data")
	}

	//customize output
	output := OutputArticle{
		ID:         article.ID,
		CategoryID: article.CategoryID,
		Title:      article.Title,
		Slug:       article.Slug,
		Content:    article.Content,
		CreatedAt:  article.CreatedAt,
		UpdatedAt:  article.UpdatedAt,
	}

	return c.JSON(http.StatusOK, output)
}

func GetAllArticles(c echo.Context) error {
	//get param
	category := c.QueryParam("category")
	title := c.QueryParam("title")

	var articles []databases.ListArticle

	if title == "" && category == "" {
		//get all articles
		articles, _ = databases.GetAllArticles()
		if len(articles) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "cannot find article")
		}
	} else if category == "" {
		//get all articles by title
		articles, _ = databases.GetArticlesByTitle(title)
		if len(articles) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "cannot find article")
		}
	} else if title == "" {
		//get all articles by category
		articles, _ = databases.GetArticlesByCategory(category)
		if len(articles) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "cannot find article")
		}
	} else if title != "" && category != "" {
		//get all articles by title and category
		articles, _ = databases.GetArticlesByTitleAndCategory(title, category)
		if len(articles) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "cannot find article")
		}
	}

	//customize output
	var output []ListArticle
	for i := 0; i < len(articles); i++ {
		new_result := ListArticle{
			ID:       articles[i].ID,
			Category: articles[i].CategoryName,
			Title:    articles[i].Title,
			Slug:     articles[i].Slug,
		}
		output = append(output, new_result)
	}

	return c.JSON(http.StatusOK, output)
}

func GetDetailArticle(c echo.Context) error {
	//get param
	slug := c.Param("slug")

	//get detail article
	article, _ := databases.GetDetailArticle(slug)
	if article.ID == 0 {
		return c.JSON(http.StatusInternalServerError, "cannot find article")
	}

	return c.JSON(http.StatusOK, article)
}

func EditArticle(c echo.Context) error {
	//get param
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	//get user's input
	article, err := databases.GetArticleById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	c.Bind(&article)

	//update data
	article_updated, err := databases.EditArticle(article)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot update data")
	}

	//customize output
	output := OutputArticle{
		ID:         article_updated.ID,
		CategoryID: article_updated.CategoryID,
		Title:      article_updated.Title,
		Slug:       article_updated.Slug,
		Content:    article_updated.Content,
		CreatedAt:  article_updated.CreatedAt,
		UpdatedAt:  article_updated.UpdatedAt,
	}

	return c.JSON(http.StatusOK, output)
}

func DeleteArticle(c echo.Context) error {
	//get param
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	//delete data
	article, err := databases.DeleteArticle(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "invalid id",
			"data":    article,
		})
	}

	return c.JSON(http.StatusOK, "delete article success")
}
