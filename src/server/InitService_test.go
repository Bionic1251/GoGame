package server

import (
	"log"
	"testing"
)

/**
* Run To test Coordinate functionality
* Test cases
	* Coordinate value cannot be x <= 0
	* Max coordinate

* Practical unified approach for error handling
* Let operating system end the program as failed status
	* os.Exit(1) - failed
	* os.Exit(0) - success

func CheckError (err error) {
	if err != nil {
		log.Println ("Fatal Error" , err.Error())
		os.Exit(1)
	}
}
*/
func TestSetLocation(t *testing.T) {
	var x, y, radius float64
	x, y, radius = 0, 5.7, 2

	if _, err := SetLocation(radius, Coordinate{x, y}); err != nil {
		t.Errorf("problems with coordinates %f %f %f", x, y, radius)
	} else {
		log.Println("worked out ok")
	}
}
