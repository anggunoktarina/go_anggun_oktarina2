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
func GetCategorysController(c echo.Context) error {
	categorys, err := usecase.GetListCategorys()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success",
		"categorys": categorys,
	})
}

// GetMahasiswaController returns mahasiswa data based on ID
func GetCategoryController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	category, err := usecase.GetCategory(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"category": category,
	})
}

func GetCategoryByCategoryNameController(c echo.Context) error {
	name := c.QueryParam("name")

	category, err := usecase.GetCategoryByName(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"category": category,
	})
}

// CreateMahasiswaController creates a new mahasiswa
func CreateCategoryController(c echo.Context) error {
	requestPayload := new(payload.CreateCategoryRequest)

	// Bind and validate the payload
	if err := c.Bind(requestPayload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(requestPayload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Check if the user_id exists in the users table
	_, err := usecase.GetCategory(requestPayload.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user_id")
	}

	category := &model.Category{
		Name: requestPayload.Name,
	}
	err = usecase.CreateCategory(category)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	responsePayload := &payload.CreateCategoryResponse{
		CategoryID: category.ID,
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":   "success",
		"category": responsePayload,
	})
}

// UpdateMahasiswaController updates mahasiswa data based on ID
func UpdateCategoryController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	categoryToUpdate, err := usecase.GetCategory(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	updatedCategory := new(payload.UpdateCategoryRequest)

	// Bind and validate the payload
	if err := c.Bind(updatedCategory); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(updatedCategory); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Update mahasiswa data
	categoryToUpdate.Name = updatedCategory.Name

	err = usecase.UpdateCategory(&categoryToUpdate) // Pass the pointer to mahasiswaToUpdate
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &payload.UpdateCategoryResponse{
		CategoryID: categoryToUpdate.ID,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success",
		"mahasiswa": response,
	})
}

// DeleteMahasiswaController deletes mahasiswa data based on ID
func DeleteCategoryController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = usecase.DeleteCategory(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "category deleted successfully",
	})
}
