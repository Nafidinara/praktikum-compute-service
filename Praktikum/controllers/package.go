package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"Praktikum/database"
	"Praktikum/models"
)

func GetAllPackage(c echo.Context) error {
	var packages []models.Package

	err := database.DB.Find(&packages).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "err",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"message": "success get all package",
		"data":   packages,
	})
}

func CreatePackage(c echo.Context) error {
	pkg := models.Package{}
	c.Bind(&pkg)

	err := database.DB.Save(&pkg).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "err",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"message": "success create package",
		"data":   pkg,
	})
}

func UpdatePackage(c echo.Context) error {
	pkg := models.Package{}
	c.Bind(&pkg)

	existingPkg := models.Package{}
	id := c.Param("id")

	err := database.DB.First(&existingPkg, id).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "err",
			"message": err.Error(),
		})
	}

	existingPkg.Name = pkg.Name
	existingPkg.Sender = pkg.Sender
	existingPkg.Receiver = pkg.Receiver
	existingPkg.SenderLocation = pkg.SenderLocation
	existingPkg.ReceiverLocation = pkg.ReceiverLocation
	existingPkg.Fee = pkg.Fee
	existingPkg.Weight = pkg.Weight

	err = database.DB.Save(&existingPkg).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "err",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"message": "package updated successfully",
		"data":   existingPkg,
	})
}

func GetOnePackage(c echo.Context) error {
	pkg := models.Package{}
	id := c.Param("id")

	err := database.DB.First(&pkg, id).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "err",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"message": "success get one package",
		"data":   pkg,
	})
}

func DeletePackage(c echo.Context) error {
	pkg := models.Package{}
	id := c.Param("id")	

	err := database.DB.Delete(&pkg, id).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "err",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"message": "package deleted successfully",
	})
}