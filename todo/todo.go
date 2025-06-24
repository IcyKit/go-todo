package todo

type ToDo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_сompleted"`
}

func (ToDo) TableName() string {
	return "todo"
}
