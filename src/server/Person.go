package server

import (
	"fmt"
)

type Person struct {
	Coordinate
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("First name: %s, last name: %s \nCoordinates: %v ", p.FirstName, p.LastName, p.Coordinate)
}

/**
func (a Area) Locate (p Person) (Container, error)  {
	for i := 0 ; i < len (a); i++ {
		if a[i].Contains(p) {
			return a[i], nil
		}
	}
	return nil, errors.New("not found")
}


*/
