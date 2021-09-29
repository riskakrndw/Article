package databases

import (
	"project/pasarwarga/config"
	"project/pasarwarga/models"
	"time"
)

type ListArticle struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	Title        string `json:"title"`
	Slug         string `json:"slug"`
}

type DetailArticle struct {
	ID           uint      `json:"id"`
	CategoryName string    `json:"category_name"`
	CategorySlug string    `json:"category_slug"`
	Title        string    `json:"title"`
	Slug         string    `json:"slug"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func CreateArticle(article models.Article) (models.Article, error) {
	if err := config.DB.Save(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

func GetAllArticles() ([]ListArticle, error) {
	var list_articles []ListArticle
	var articles []models.Article
	if err := config.DB.Model(&articles).Select("articles.id, categories.category_name, articles.title, articles.slug ").Joins("JOIN categories ON categories.id = articles.category_id").Scan(&list_articles).Error; err != nil {
		return list_articles, err
	}
	return list_articles, nil
}

func GetArticlesByTitle(title string) ([]ListArticle, error) {
	var list_articles []ListArticle
	var articles []models.Article
	search := "%" + title + "%"
	if err := config.DB.Model(&articles).Select("articles.id, categories.category_name, articles.title, articles.slug ").Joins("JOIN categories ON categories.id = articles.category_id AND articles.title LIKE ?", search).Scan(&list_articles).Error; err != nil {
		return list_articles, err
	}
	return list_articles, nil
}

func GetArticlesByCategory(category string) ([]ListArticle, error) {
	var list_articles []ListArticle
	var articles []models.Article
	search := "%" + category + "%"
	if err := config.DB.Model(&articles).Select("articles.id, categories.category_name, articles.title, articles.slug ").Joins("JOIN categories ON categories.id = articles.category_id AND categories.category_name LIKE ? ", search).Scan(&list_articles).Error; err != nil {
		return list_articles, err
	}
	return list_articles, nil
}

func GetArticlesByTitleAndCategory(title, category string) ([]ListArticle, error) {
	var list_articles []ListArticle
	var articles []models.Article
	category = "%" + category + "%"
	title = "%" + title + "%"
	if err := config.DB.Model(&articles).Select("articles.id, categories.category_name, articles.title, articles.slug ").Joins("JOIN categories ON categories.id = articles.category_id AND categories.category_name LIKE ? AND articles.title LIKE ?", category, title).Scan(&list_articles).Error; err != nil {
		return list_articles, err
	}
	return list_articles, nil
}

func GetDetailArticle(slug string) (DetailArticle, error) {
	var detail_article DetailArticle
	var article models.Article
	if err := config.DB.Model(&article).Select("articles.*, categories.category_name, categories.category_slug").Joins("JOIN categories ON categories.id = articles.category_id AND articles.slug = ? ", slug).Scan(&detail_article).Error; err != nil {
		return detail_article, err
	}
	return detail_article, nil
}

func GetArticleById(id int) (models.Article, error) {
	var article models.Article

	if err := config.DB.Where("id = ?", id).First(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

func EditArticle(article models.Article) (models.Article, error) {
	if err := config.DB.Save(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

func DeleteArticle(id int) (models.Article, error) {
	var article models.Article
	if err := config.DB.Delete(&article, "id = ?", id).Error; err != nil {
		return article, err
	}
	return article, nil
}
