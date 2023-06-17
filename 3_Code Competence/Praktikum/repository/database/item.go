package database

import (
	"presensee_project/config"
	"presensee_project/model"
)

func CreateItem(item *model.Item) error {
	if err := config.DB.Create(item).Error; err != nil {
		return err
	}
	return nil
}


func GetItems() (items []model.Item, err error) {
	if err = config.DB.Model(&model.Item{}).Find(&items).Error; err != nil {
		return
	}
	return
}

func GetItemByID(id uint) (item model.Item, err error) {
	if err = config.DB.First(&item, id).Error; err != nil {
		return
	}
	return
}

func GetItemByName(name string) ([]model.Item, error) {
	var items []model.Item

	err := config.DB.Where("name LIKE ? ", "%"+name+"%").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func UpdateItem(item *model.Item) error {
	if err := config.DB.Save(item).Error; err != nil {
		return err
	}
	return nil
}

func DeleteItem(id uint) error {
	if err := config.DB.Delete(&model.Item{}, id).Error; err != nil {
		return err
	}
	return nil
}
