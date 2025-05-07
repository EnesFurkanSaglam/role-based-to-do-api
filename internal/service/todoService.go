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

func AddStep(listID int, content string) model.TodoStep {
	return repository.AddTodoStep(listID, content)
}

func GetSteps(listID int) []model.TodoStep {
	return repository.GetStepsByListID(listID)
}

func CompleteStep(stepID int) bool {
	return repository.CompleteStep(stepID)
}

func DeleteStep(stepID int) bool {
	return repository.DeleteStep(stepID)
}
