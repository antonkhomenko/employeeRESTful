package storage

import (
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

var naruto = Employee{1, "Naruto", "Uzumaki", 18, 500}

func TestMongoDB_Update(t *testing.T) {
	data := struct {
		input  Employee
		result bson.D
	}{
		naruto,
		bson.D{{"_id", 1}, {"name", "Naruto"}, {"surname", "Uzumaki"}, {"age", 18}, {"salary", 500}},
	}
	narutoByte, err := bson.Marshal(naruto)
	if err != nil {
		t.Fatal()
	}
	var narutoBsonD bson.D
	err = bson.Unmarshal(narutoByte, &narutoBsonD)
	if err != nil {
		t.Fatal()
	}

}
