package entity

type Person struct {
	Base
	Name    string
	Email   string
	Id      int
	History []Meeting
}
