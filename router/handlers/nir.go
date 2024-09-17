package handlers

import (
	"cleanarch/infra/db/conn"
	nirRepository "cleanarch/infra/db/repository/nir"
	"cleanarch/services/nir"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetNIRResults(c echo.Context) error {
	dbConn, err := conn.GetPSQLConn()
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error connecting to database: %v", err))
	}
	defer dbConn.Close(context.Background())

	service := nir.NewNirUseCases(nirRepository.NewINirRepository(dbConn))
	results, err := service.GetNirResults()
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error getting NIR results: %v", err))
	}

	return c.JSON(http.StatusOK, results)
}

func GetNIRResultsByID(c echo.Context) error {
	dbConn, err := conn.GetPSQLConn()
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error connecting to database: %v", err))
	}
	defer dbConn.Close(context.Background())

	service := nir.NewNirUseCases(nirRepository.NewINirRepository(dbConn))
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error converting ID to integer: %v", err))
	}

	result, err := service.GetNIRResultsByID(int64(idInt))
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error getting NIR result: %v", err))
	}

	return c.JSON(http.StatusOK, result)
}
