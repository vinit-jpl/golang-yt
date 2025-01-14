package main

// Model for course - file
type Course struct {
	CourseId    string  `json:"Courseid"`
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
func IsEmpty(c *Course) bool {
	return c.CourseId == "" && c.CourseName == ""
}
func main() {

}
