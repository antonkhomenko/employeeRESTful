package storage

import (
	"fmt"
)

var (
	invalidID error = fmt.Errorf("employee with such id dosent exist")
)

type Employee struct {
	ID      int     `json:"id" bson:"_id"`
	Name    string  `json:"name" bson:"name"`
	Surname string  `json:"surname" bson:"surname"`
	Age     int     `json:"age" bson:"age"`
	Salary  float64 `json:"salary" bson:"salary"`
}

type Storage interface {
	Insert(e *Employee) error
	Get(id int) (Employee, error)
	GetAll() ([]Employee, error)
	Update(id int, s map[string]any) error
	Delete(id int) error
}
