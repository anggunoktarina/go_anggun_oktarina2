package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"user_id"`
}

func main() {

	e := echo.New()
	e.GET("/", HelloController)
	e.GET("/users", GetUser)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.GET("/books", GetBookController)
	e.GET("/books/:id", GetBookByIdController)
	e.POST("/books", CreateBookController)
	e.DELETE("/books/:id", DeleteBookByIdController)
	e.Start(":8000")

}

func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.QueryParam("name")
	user := User{
		Id:          id,
		Name:        name,
		Email:       "alta@example.com",
		PhoneNumber: "08123123123",
	}
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to get user",
		Status:  "Success",
		Data: map[string]interface{}{
			"user": user,
		},
	})
}

func GetUser(c echo.Context) error {
	user := User{
		Name:        "Alta",
		Email:       "alta@example.com",
		PhoneNumber: "08123123123",
	}
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to get user",
		Status:  "Success",
		Data:    user,
	})
}

func HelloController(c echo.Context) error {
	println("hit this endpoint")
	return c.String(http.StatusOK, "Hello World")
}

func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to create user",
		Status:  "Success",
		Data:    user,
	})
}

func GetBookController(c echo.Context) error {
	response, _ := http.Get("https://jsonplaceholder.typicode.com/posts")

	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var books []Book
	json.Unmarshal(responseData, &books)
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to get Book",
		Status:  "Success",
		Data:    books,
	})
}

func GetBookByIdController(c echo.Context) error {
	id := c.Param("id")
	response, _ := http.Get("https://jsonplaceholder.typicode.com/posts" + "/" + id)

	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var books Book
	json.Unmarshal(responseData, &books)
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to get Book",
		Status:  "Success",
		Data:    books,
	})
}

func CreateBookController(c echo.Context) error {
	books := Book{}
	c.Bind(&books)
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success to create book",
		Status:  "Success",
		Data:    books,
	})
}

func DeleteBookByIdController(c echo.Context) error {
	id := c.Param("id")

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/posts/"+id, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create request",
			Status:  "Error",
			Data:    nil,
		})
	}

	response, err := client.Do(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete book",
			Status:  "Error",
			Data:    nil,
		})
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		return c.JSON(http.StatusOK, Response{
			Code:    http.StatusOK,
			Message: "Book deleted successfully",
			Status:  "Success",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusNotFound, Response{
		Code:    http.StatusNotFound,
		Message: "Book not found",
		Status:  "Error",
		Data:    nil,
	})
}
