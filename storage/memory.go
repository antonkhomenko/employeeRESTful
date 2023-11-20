package storage

import (
	"encoding/json"
	"sync"
)

type MemoryStorage struct {
	counter int
	data    map[int]Employee
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		counter: 1,
		data:    make(map[int]Employee),
	}
}

func (ms *MemoryStorage) Insert(e *Employee) error {
	ms.Lock()
	e.ID = ms.counter
	ms.data[e.ID] = *e
	ms.counter++
	ms.Unlock()
	return nil
}

func (ms *MemoryStorage) Get(id int) (Employee, error) {
	ms.Lock()
	e, ok := ms.data[id]
	ms.Unlock()
	if !ok {
		return e, invalidID
	}
	return e, nil
}

func (ms *MemoryStorage) GetAll() ([]Employee, error) {
	ms.Lock()
	result := make([]Employee, 0, len(ms.data))
	for _, value := range ms.data {
		result = append(result, value)
	}
	ms.Unlock()
	return result, nil
}

func (ms *MemoryStorage) Update(id int, inputData map[string]any) error {
	ms.Lock()
	employee, ok := ms.data[id]
	if !ok {
		return invalidID
	}
	employeeB, err := json.Marshal(employee)
	if err != nil {
		return err
	}
	var employeeMap map[string]any
	err = json.Unmarshal(employeeB, &employeeMap)
	if err != nil {
		return err
	}
	for key, value := range inputData {
		if value != employeeMap[key] {
			employeeMap[key] = value
		}
	}
	dataB, err := json.Marshal(employeeMap)
	if err != nil {
		return err
	}
	err = json.Unmarshal(dataB, &employee)
	if err != nil {
		return err
	}
	ms.data[id] = employee
	ms.Unlock()
	return nil
}

func (ms *MemoryStorage) Delete(id int) error {
	ms.Lock()
	_, ok := ms.data[id]
	if !ok {
		return invalidID
	}
	delete(ms.data, id)
	ms.Unlock()
	return nil
}
