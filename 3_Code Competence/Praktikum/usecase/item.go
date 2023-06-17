package usecase

import (
	"presensee_project/model"
	"presensee_project/repository/database"
)

func CreateItem(item *model.Item) error {
	err := database.CreateItem(item)
	if err != nil {
		return err
	}
	return nil
}

func GetItem(id uint) (model.Item, error) {
	item, err := database.GetItemByID(id)
	if err != nil {
		return model.Item{}, err
	}
	return item, nil
}

func GetItemByName(name string) ([]model.Item, error) {
	item, err := database.GetItemByName(name)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func GetListItems() ([]model.Item, error) {
	items, err := database.GetItems()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func UpdateItem(item *model.Item) error {
	err := database.UpdateItem(item)
	if err != nil {
		return err
	}
	return nil
}

func DeleteItem(id uint) error {
	err := database.DeleteItem(id)
	if err != nil {
		return err
	}
	return nil
}