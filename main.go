package main

import (
	"fmt"
	"net/http"
)

/*
go run main.go

OR

GOOS=windows go build
and then you can run the binary with ./filename
*/

func main() {
	course := Course{
		Name:        "golang",
		Description: "golang",
		Price:       1.0,
	}

	fmt.Println(course.getAllAtributes())

	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

// we can also define functions like this
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GOing well"))
}

// OBS ABOUT STRUCTS AND THEIR METHODS: names starting with upper case are public to other packages and lower case are accessible only here
type Course struct {
	Name        string
	Description string
	Price       float64
}

// we can create methods for structs (actually cooler than classes)
func (c Course) getAllAtributes() string {
	return fmt.Sprintf("Name: %s, Description: %s, Price: %f", c.Name, c.Description, c.Price)
}
