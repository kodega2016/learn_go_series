package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

func main() {
	r := mux.NewRouter()
	// seeding data
	c1 := Course{
		CourseID:    "1",
		CourseName:  "Flutter and firebase",
		CoursePrice: 12,
		Author: &Author{
			Fullname: "khadga bahadur shrestha",
			Website:  "kodega.com",
		},
	}

	c2 := Course{
		CourseID:    "2",
		CourseName:  "Microservices",
		CoursePrice: 15,
		Author: &Author{
			Fullname: "khadga bahadur shrestha",
			Website:  "kodega.com",
		},
	}

	courses = append(courses, c1)
	courses = append(courses, c2)

	// routing
	r.HandleFunc("/", serveHome)
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	fmt.Println("web server is running on 8080")
	http.ListenAndServe(":8080", r)

}

// controllers- file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllCourses(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Println("get all courses...")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	// grab id from the request
	params := mux.Vars(r)
	id := params["id"]

	// loop through the courses, find the mathcing id and return the response
	for _, course := range courses {
		if course.CourseID == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")
	return

}

func createOneCourse(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No valid data inside the json")
		return
	}

	// generate unique id,string
	// append course into courses
	rand.Seed(time.Now().Unix())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	// first grab the id from the req
	params := mux.Vars(r)
	id := params["id"]

	for index, course := range courses {
		if course.CourseID == id {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = id
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}
func deleteOneCourse(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	// first grab the id from the req
	params := mux.Vars(r)
	id := params["id"]

	for index, course := range courses {
		if course.CourseID == id {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

}
