package domain

type Todo struct {

    ID        int    `json:"id"`
    Text      string `json:"text"`
    Completed bool   `json:"completed"`
}

type TodoUsecase interface {
		AllGet() ([]Todo, error)
		StatusUpdate(id int) error
		Store(todo Todo) error
		Delete(id int) error
		Search(key string) ([]Todo, error) // 追加
}

type TodoRepository interface {
		AllGet() ([]Todo, error)
		StatusUpdate(id int) error
		Store(todo Todo) error
		Delete(id int) error
		Search(key string) ([]Todo, error) // 追加
}
