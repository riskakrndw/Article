package controllers

import (
	"fmt"
	"net/http"
	"project/pasarwarga/databases"
	"project/pasarwarga/models"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type OutputCategory struct {
	ID           uint      `json:"id"`
	CategoryName string    `json:"category_name"`
	CategorySlug string    `json:"category_slug"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type OutputCategory2 struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
}

func CreateCategory(c echo.Context) error {
	//get user's input
	input_user := models.Category{}
	c.Bind(&input_user)

	//check is data nil?
	if input_user.CategoryName == "" || input_user.CategorySlug == "" {
		return c.JSON(http.StatusBadRequest, "Please fill all data")
	}

	//check category
	category, is_category_exists, err := databases.CheckCategoryExists(input_user.CategoryName, input_user.CategorySlug)
	category_deleted, is_category_deleted, _ := databases.CheckCategoryDeleted(input_user.CategoryName)
	fmt.Println(is_category_exists, is_category_deleted)
	if !is_category_exists && is_category_deleted {
		category, err = databases.UpdateCategoryDeleted(category_deleted)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot update data")
		}
	} else if !is_category_exists && !is_category_deleted {
		category, err = databases.CreateCategory(input_user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot insert data")
		}
	} else if is_category_exists && !is_category_deleted {
		return c.JSON(http.StatusBadRequest, "Name or Slug already exists")
	}

	//customize output
	output := OutputCategory{
		ID:           category.ID,
		CategoryName: category.CategoryName,
		CategorySlug: category.CategorySlug,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
	}

	return c.JSON(http.StatusOK, output)
}

func GetAllCategories(c echo.Context) error {
	//get param
	name := c.QueryParam("name")

	var all_categories []models.Category

	if name == "" {
		//get all categories
		all_categories, _ = databases.GetAllCategories()
		if len(all_categories) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "cannot find categories")
		}
	} else {
		//get all categories
		all_categories, _ = databases.GetCategoriesByName(name)
		if len(all_categories) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "cannot find categories")
		}
	}

	//customize output
	var output []OutputCategory2
	for i := 0; i < len(all_categories); i++ {
		new_result := OutputCategory2{
			ID:           all_categories[i].ID,
			CategoryName: all_categories[i].CategoryName,
		}
		output = append(output, new_result)
	}

	return c.JSON(http.StatusOK, output)
}

func GetDetailCategory(c echo.Context) error {
	//get param
	category_slug := c.Param("category_slug")

	//get detail category
	category, _ := databases.GetDetailCategory(category_slug)
	if category.ID == 0 {
		return c.JSON(http.StatusInternalServerError, "cannot find category")
	}

	//customize output
	output := OutputCategory{
		ID:           category.ID,
		CategoryName: category.CategoryName,
		CategorySlug: category.CategorySlug,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
	}

	return c.JSON(http.StatusOK, output)
}

func EditCategory(c echo.Context) error {
	//get param
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	//get user's input
	category, err := databases.GetCategoryById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	c.Bind(&category)

	//update data
	category_updated, err := databases.EditCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot update data")
	}

	//customize output
	output := OutputCategory{
		ID:           category_updated.ID,
		CategoryName: category_updated.CategoryName,
		CategorySlug: category_updated.CategorySlug,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
	}

	return c.JSON(http.StatusOK, output)
}

func DeleteCategory(c echo.Context) error {
	//get param
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	//delete data
	category, err := databases.DeleteCategory(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "invalid id",
			"data":    category,
		})
	}

	return c.JSON(http.StatusOK, "delete category success")
}
