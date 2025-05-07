package repository

import (
	"role-based-to-do-api/internal/model"
	"time"
)

var (
	todoLists []model.TodoList
	todoSteps []model.TodoStep
	listID    = 1
	stepID    = 1
)

func AddTodoList(name string, owner string) model.TodoList {
	now := time.Now()
	list := model.TodoList{
		ID:         listID,
		Name:       name,
		CreatedAt:  now,
		UpdatedAt:  now,
		Completion: 0,
		Owner:      owner,
	}
	listID++
	todoLists = append(todoLists, list)
	return list
}

func GetTodoListsByOwner(owner string, isAdmin bool) []model.TodoList {
	var result []model.TodoList
	for _, l := range todoLists {
		if l.DeletedAt == nil && (isAdmin || l.Owner == owner) {
			result = append(result, l)
		}
	}
	return result
}
