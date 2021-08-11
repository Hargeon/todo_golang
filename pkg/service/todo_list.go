package service

import (
	"github.com/Hargeon/todo"
	"github.com/Hargeon/todo/pkg/repository"
)

type TodoService struct {
	repo repository.TodoList
}

func NewTodoService(repo repository.TodoList) *TodoService {
	return &TodoService{repo: repo}
}

func (t *TodoService) Create(userId int, list todo.TodoList) (int, error) {
	return t.repo.Create(userId, list)
}

func (t *TodoService) GetLists(userId int) ([]*todo.TodoList, error) {
	return t.repo.GetLists(userId)
}
func (t *TodoService) GetList(listId, userId int) (todo.TodoList, error) {
	return t.repo.GetList(listId, userId)
}