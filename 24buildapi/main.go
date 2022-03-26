package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"name"`
	Website  string `json:"website"`
}

// dummy DB
var courses []Course

func (course *Course) IsEmpty() bool {
	// return course.CourseId == "" && course.CourseName == ""
	return course.CourseName == ""
}

func main() {
	fmt.Println("Building APIs in Golang!")

	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "Apache Kafka Mastery", CoursePrice: 499, Author: &Author{FullName: "Aditya Garde", Website: "www.github.com/adityagarde"}})

	courses = append(courses, Course{CourseId: "4", CourseName: "Reactive Microservices with Webflux", CoursePrice: 499, Author: &Author{FullName: "Aditya Garde", Website: "www.github.com/adityagarde"}})

	courses = append(courses, Course{CourseId: "5", CourseName: "Introduction to Go", CoursePrice: 399, Author: &Author{FullName: "Aditya Garde", Website: "www.linkedin.com/in/adityagarde/"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal((http.ListenAndServe(":4000", r)))
}

// controllers
// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Building APIs in Golang!</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses.")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// Loop through courses [] to find matching id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found for id == " + params["id"])
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// if body is {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data in the JSON.")
		return
	}

	// generate unique ID & append it to course

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// Loop through courses - find the required id, remove it, add back with same Id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found for id == " + params["id"])

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// Loop through courses - find the required id, remove it
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

}
