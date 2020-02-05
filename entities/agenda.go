package entity

type Agenda struct {
	Base
	Meeting string
	Items   []AgendaItem
	Date    string
}

type AgendaItem struct {
	Title   string
	Details string
	Index   int
}

func CreateAgenda()    {}
func UpdateAgenda()    {}
func CloneAgenda()     {}
func DeleteItem()      {}
func PopItem()         {}
func PushItem()        {}
func ChangeItemIndex() {}
