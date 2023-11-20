package main

import (
	"database-lesson/storage"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string
}

type Handler struct {
	storage storage.Storage
}

func NewHandler(s storage.Storage) *Handler {
	return &Handler{
		storage: s,
	}
}

func (h *Handler) CreateEmployee(c *gin.Context) {
	var employee storage.Employee
	if err := c.Bind(&employee); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}
	err := h.storage.Insert(&employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": employee.ID,
	})
}

func (h *Handler) ReadEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("failed to convert param id into int")
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}
	employee, err := h.storage.Get(id)
	if err != nil {
		fmt.Printf("failed to get employee with id: %v\n", id)
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"employee": employee,
	})
}

func (h *Handler) ReadAllEmployees(c *gin.Context) {
	employees, err := h.storage.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"employees": employees,
	})
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}
	var inputData map[string]any
	err = c.Bind(&inputData)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}
	fmt.Printf("inputData = %v\n", inputData)
	err = h.storage.Update(id, inputData)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}
	err = h.storage.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
