package databases

import (
	"project/pasarwarga/config"
	"project/pasarwarga/models"
)

func CheckCategoryExists(name, slug string) (models.Category, bool, error) {
	var category models.Category

	if err := config.DB.Model(&category).Where("category_name = ? OR category_slug = ?", name, slug).First(&category).Error; err != nil {
		return category, false, err
	}

	if category.CategoryName == name || category.CategorySlug == slug {
		return category, true, nil
	} else {
		return category, false, nil
	}
}

func CheckCategoryDeleted(name string) (models.Category, bool, error) {
	var category models.Category

	if err := config.DB.Raw("SELECT * FROM categories WHERE category_name = ? AND deleted_at IS NOT NULL", name).Scan(&category).Error; err != nil {
		return category, false, err
	}

	if category.ID == 0 {
		return category, false, nil
	} else {
		return category, true, nil
	}
}

func CreateCategory(category models.Category) (models.Category, error) {
	if err := config.DB.Save(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func GetAllCategories() ([]models.Category, error) {
	var all_categories []models.Category
	if err := config.DB.Find(&all_categories).Error; err != nil {
		return all_categories, err
	}
	return all_categories, nil
}

func GetCategoriesByName(name string) ([]models.Category, error) {
	var all_categories []models.Category
	search := "%" + name + "%"
	if err := config.DB.Find(&all_categories, "category_name LIKE ?", search).Error; err != nil {
		return all_categories, err
	}
	return all_categories, nil
}

func GetDetailCategory(category_slug string) (models.Category, error) {
	var category models.Category
	if err := config.DB.Find(&category, "category_slug = ?", category_slug).Error; err != nil {
		return category, err
	}
	return category, nil
}

func GetCategoryById(id int) (models.Category, error) {
	var category models.Category

	if err := config.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func EditCategory(category models.Category) (models.Category, error) {
	if err := config.DB.Save(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func UpdateCategoryDeleted(category models.Category) (models.Category, error) {
	if err := config.DB.Raw("UPDATE categories SET deleted_at = NULL WHERE id = ?", category.ID).Scan(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func DeleteCategory(id int) (models.Category, error) {
	var category models.Category
	if err := config.DB.Delete(&category, "id = ?", id).Error; err != nil {
		return category, err
	}
	return category, nil
}
