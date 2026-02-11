package db

type Task struct {
	ID      int64  `db:"id"`
	Date    string `db:"date"`
	Title   string `db:"title"`
	Comment string `db:"comment"`
	Repeat  string `db:"repeat"`
}

type TaskService struct {
	store TaskStore
}

func NewTaskService(store TaskStore) TaskService {
	return TaskService{store: store}
}
