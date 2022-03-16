package main

import "fmt"

// here global scope

var gpa int
var studentName string
var isGraduated bool


func main(){
//custom  scoped
gpa = 3
studentName = "Ram"
isGraduated = true
fmt.Println(gpa)
fmt.Println(studentName)
fmt.Println(isGraduated)
fmt.Println("GPA: ",gpa,"Student: ", studentName, "IsGraduted: ", isGraduated)
}