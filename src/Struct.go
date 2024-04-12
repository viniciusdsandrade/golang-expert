package main

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"time"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	YearBorn  int    `json:"year_born"`
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) Age() int {
	currentYear := time.Now().Year()
	return currentYear - p.YearBorn
}

func (p Person) ToString() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	return string(bytes)
}

func (p Person) HashCode() uint32 {
	h := fnv.New32a()
	_, err := h.Write([]byte(p.ToString()))
	if err != nil {
		return 0
	}
	return h.Sum32()
}

func (p Person) Equals(other Person) bool {
	return p.FirstName == other.FirstName &&
		p.LastName == other.LastName &&
		p.YearBorn == other.YearBorn
}

func (p Person) Clone() (*Person, error) {
	bytes, err := json.Marshal(p)

	if err != nil {
		return nil, err
	}

	var myCopy Person
	err = json.Unmarshal(bytes, &myCopy)
	if err != nil {
		return nil, err
	}

	return &myCopy, nil
}
