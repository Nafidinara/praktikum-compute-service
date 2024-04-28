package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"Praktikum/middlewares"
	"Praktikum/models"
	"Praktikum/services"
)

type UserController struct {
	service services.UserService
}

func InitUserController(jwtAuth *middlewares.JWTConfig) UserController {
	return UserController{
		service: services.InitUserService(jwtAuth),
	}
}

func (controller *UserController) Register(c echo.Context) error {
	var userInput models.RegisterInput

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	user, err := controller.service.Register(userInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.Response{
		Status:  "success",
		Message: "success create user",
		Data:    user,
	})
}
func (controller *UserController) Login(c echo.Context) error {
	var userInput models.LoginInput

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	token, err := controller.service.Login(userInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success login",
		Data:    token,
	})
}

func (controller *UserController) GetUser(c echo.Context) error {
	claims, err := middlewares.GetUser(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	user, err := controller.service.GetUser(strconv.Itoa(claims.ID))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success get user",
		Data:    user,
	})
}
