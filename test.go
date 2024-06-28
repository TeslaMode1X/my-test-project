package tester

import "strconv"

var age int = 15
var name string = "Not Given"

func AddInts(a, b int) int {
	return a + b
}

func SayHello(userName string) string {
	name = userName
	return "Hi, " + name
}

func GetInfo() string {
	return name + " " + strconv.Itoa(age)
}
