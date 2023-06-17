package controller

import (
	"net/http"
	"presensee_project/model"
	"presensee_project/model/payload"
	"presensee_project/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetMahasiswasController returns all mahasiswa data
func GetItemsController(c echo.Context) error {
	items, err := usecase.GetListItems()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"items":  items,
	})
}

// GetMahasiswaController returns mahasiswa data based on ID
func GetItemController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	item, err := usecase.GetItem(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success",
		"item": item,
	})
}

func GetItemByItemNameController(c echo.Context) error {
	name := c.QueryParam("name")

	item, err := usecase.GetItemByName(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"item":   item,
	})
}


// CreateMahasiswaController creates a new mahasiswa
func CreateItemController(c echo.Context) error {
	requestPayload := new(payload.CreateItemRequest)

	// Bind and validate the payload
	if err := c.Bind(requestPayload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(requestPayload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Check if the user_id exists in the users table
	_, err := usecase.GetItem(requestPayload.CategoryID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user_id")
	}

	item := &model.Item{
		Name:        requestPayload.Name,
		Description: requestPayload.Description,
		Stock:       requestPayload.Stock,
		Price:       requestPayload.Price,
		CategoryID:  requestPayload.CategoryID,
	}
	err = usecase.CreateItem(item)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	responsePayload := &payload.CreateItemResponse{
		ItemID: item.ID,
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"item":   responsePayload,
	})
}

// UpdateMahasiswaController updates mahasiswa data based on ID
func UpdateItemController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	itemToUpdate, err := usecase.GetItem(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	updatedItem := new(payload.UpdateItemRequest)

	// Bind and validate the payload
	if err := c.Bind(updatedItem); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(updatedItem); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Update mahasiswa data
	itemToUpdate.Name = updatedItem.Name
	itemToUpdate.Description = updatedItem.Description
	itemToUpdate.Stock = updatedItem.Stock
	itemToUpdate.CategoryID = updatedItem.CategoryID
	itemToUpdate.Price = updatedItem.Price

	err = usecase.UpdateItem(&itemToUpdate) // Pass the pointer to mahasiswaToUpdate
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &payload.UpdateItemResponse{
		ItemID: itemToUpdate.ID,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success",
		"mahasiswa": response,
	})
}

// DeleteMahasiswaController deletes mahasiswa data based on ID
func DeleteItemController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = usecase.DeleteItem(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Item deleted successfully",
	})
}
