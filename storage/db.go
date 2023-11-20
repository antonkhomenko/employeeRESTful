package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	c *mongo.Collection
}

func NewMongoDB(client *mongo.Client) *MongoDB {
	return &MongoDB{
		c: client.Database("employeebox").Collection("employee"),
	}
}

func (m *MongoDB) Insert(e *Employee) error {
	_, err := m.c.InsertOne(context.TODO(), e)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) Get(id int) (Employee, error) {
	var employee Employee
	result := m.c.FindOne(context.TODO(), bson.D{{"_id", id}})
	err := result.Decode(&employee)
	if err != nil {
		return employee, err
	}
	return employee, nil
}

func (m *MongoDB) GetAll() ([]Employee, error) {
	var employeeList []Employee
	cursor, err := m.c.Find(context.TODO(), bson.D{})
	if err != nil {
		return employeeList, err
	}
	err = cursor.All(context.TODO(), &employeeList)
	if err != nil {
		return employeeList, err
	}
	return employeeList, nil
}

func (m *MongoDB) Update(id int, data map[string]any) error {
	employeeBson, err := bson.Marshal(data)
	if err != nil {
		return err
	}
	var filter = bson.D{{"_id", id}}
	var updateItem bson.D
	err = bson.Unmarshal(employeeBson, &updateItem)
	fmt.Println(updateItem)
	if err != nil {
		return err
	}
	_, err = m.c.UpdateOne(context.TODO(), filter, bson.D{{"$set", updateItem}})
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) Delete(id int) error {
	filter := bson.D{{"_id", id}}
	_, err := m.c.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
