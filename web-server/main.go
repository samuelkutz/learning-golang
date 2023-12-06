package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
go run main.go

OR

GOOS=windows go build
and then you can run the binary with ./filename
*/

// OBS ABOUT STRUCTS AND THEIR METHODS: names starting with upper case are public to other packages and lower case are accessible only here
type Course struct {
	Name        string  `json:"course"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// we can create methods for structs (actually cooler than classes)
func (c Course) getAllAtributes() string {
	return fmt.Sprintf("Name: %s, Description: %s, Price: %f", c.Name, c.Description, c.Price)
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

// we can also define functions like this
func home(w http.ResponseWriter, r *http.Request) {
	course := Course{
		Name:        "golang",
		Description: "golang",
		Price:       1.0,
	}

	fmt.Println((course.getAllAtributes()))

	json.NewEncoder(w).Encode((course))
}
