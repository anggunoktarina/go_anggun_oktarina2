package database

import (
	"presensee_project/config"
	"presensee_project/model"
)

func CreateCategory(category *model.Category) error {
	if err := config.DB.Create(category).Error; err != nil {
		return err
	}
	return nil
}


func GetCategorys() (categorys []model.Category, err error) {
	if err = config.DB.Model(&model.Category{}).Find(&categorys).Error; err != nil {
		return
	}
	return
}

func GetCategoryByID(id uint) (category model.Category, err error) {
	if err = config.DB.First(&category, id).Error; err != nil {
		return
	}
	return
}

func GetCategoryByName(name string) ([]model.Category, error) {
	var categorys []model.Category

	err := config.DB.Where("name LIKE ? ", "%"+name+"%").Find(&categorys).Error
	if err != nil {
		return nil, err
	}
	return categorys, nil
}

func UpdateCategory(category *model.Category) error {
	if err := config.DB.Save(category).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCategory(id uint) error {
	if err := config.DB.Delete(&model.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
