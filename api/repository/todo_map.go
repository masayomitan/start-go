package repository

import(
    "errors"
    "sync"
    "todo/domain"
)

type todoRepository struct
{
    m sync.Map
}

//インターフェイスを通す
func NewSyncMapTodoRepository() domain.TodoRepository {
    return &todoRepository{}
}

//全てのtodoを取得
func (t *todoRepository) AllGet() ([]domain.Todo, error) {
		var todos []domain.Todo
		t.m.Range(func(key interface{}, value interface{}) bool {
			todos = append(
					todos,
					value.(domain.Todo),
			)
			return true
		})
		return todos, nil
}

//ステータス更新
func (t * todoRepository) StatusUpdate( id int ) error  {
    r, ok := t.m.LoadAndDelete( id )
		if !ok {
        return errors.New("データを取得できませんでした")
		}

		newTodo := r.( domain.Todo )
		if newTodo.Completed {
				newTodo.Completed = false
		} else {
        newTodo.Completed = true
		}
		t.Store(newTodo)
		return nil
}

// 保存
func ( t * todoRepository ) Store( todo domain.Todo ) error {
		t.m.Store( todo.ID, todo )
		return nil
}

// 削除
func (t *todoRepository) Delete(id int) error {
	t.m.Delete(id)
	return nil
}