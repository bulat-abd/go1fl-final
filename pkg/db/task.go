package db

type Task struct {
	ID      int64  `db:"id" json:"id,string"`
	Date    string `db:"date" json:"date"`
	Title   string `db:"title" json:"title"`
	Comment string `db:"comment" json:"comment"`
	Repeat  string `db:"repeat" json:"repeat"`
}

type TaskService struct {
	store TaskStore
}

func NewTaskService(store TaskStore) TaskService {
	return TaskService{store: store}
}
