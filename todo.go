package todo

type TodoList struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

type UserList struct {
	Id int
	UserId int 
	ListId int 
}

type TodoItem struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"descritpion"`
	Done bool `json:"done"`
}

type ListsItem struct {
	Id int
	ListId int
	ItemId int  
}