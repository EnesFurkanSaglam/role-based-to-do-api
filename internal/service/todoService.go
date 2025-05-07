package service

import (
	"role-based-to-do-api/internal/model"
	"role-based-to-do-api/internal/repository"
)

func CreateList(name string, owner string) model.TodoList {
	return repository.AddTodoList(name, owner)
}

func GetUserLists(owner string, role string) []model.TodoList {
	isAdmin := role == "admin"
	return repository.GetTodoListsByOwner(owner, isAdmin)
}
