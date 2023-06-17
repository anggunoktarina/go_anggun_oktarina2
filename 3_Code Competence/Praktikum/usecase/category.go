package usecase

import (
	"presensee_project/model"
	"presensee_project/repository/database"
)

func CreateCategory(category *model.Category) error {
	err := database.CreateCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func GetCategory(id uint) (model.Category, error) {
	category, err := database.GetCategoryByID(id)
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func GetCategoryByName(name string) ([]model.Category, error) {
	category, err := database.GetCategoryByName(name)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func GetListCategorys() ([]model.Category, error) {
	categorys, err := database.GetCategorys()
	if err != nil {
		return nil, err
	}
	return categorys, nil
}

func UpdateCategory(category *model.Category) error {
	err := database.UpdateCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCategory(id uint) error {
	err := database.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}