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

func AddTodoStep(listID int, content string) model.TodoStep {
	now := time.Now()
	step := model.TodoStep{
		ID:        stepID,
		ListID:    listID,
		Content:   content,
		Done:      false,
		CreatedAt: now,
		UpdatedAt: now,
	}
	stepID++
	todoSteps = append(todoSteps, step)
	updateListCompletion(listID)
	return step
}

func GetStepsByListID(listID int) []model.TodoStep {
	var result []model.TodoStep
	for _, s := range todoSteps {
		if s.ListID == listID && s.DeletedAt == nil {
			result = append(result, s)
		}
	}
	return result
}

func CompleteStep(stepID int) bool {
	for i, s := range todoSteps {
		if s.ID == stepID && s.DeletedAt == nil {
			todoSteps[i].Done = true
			todoSteps[i].UpdatedAt = time.Now()
			updateListCompletion(s.ListID)
			return true
		}
	}
	return false
}

func DeleteStep(stepID int) bool {
	now := time.Now()
	for i, s := range todoSteps {
		if s.ID == stepID && s.DeletedAt == nil {
			todoSteps[i].DeletedAt = &now
			todoSteps[i].UpdatedAt = now
			updateListCompletion(s.ListID)
			return true
		}
	}
	return false
}

func updateListCompletion(listID int) {
	var total, done int
	for _, s := range todoSteps {
		if s.ListID == listID && s.DeletedAt == nil {
			total++
			if s.Done {
				done++
			}
		}
	}

	for i, l := range todoLists {
		if l.ID == listID {
			if total == 0 {
				todoLists[i].Completion = 0
			} else {
				todoLists[i].Completion = float64(done) / float64(total) * 100
			}
			todoLists[i].UpdatedAt = time.Now()
		}
	}
}

func SoftDeleteList(listID int) bool {
	now := time.Now()
	for i, l := range todoLists {
		if l.ID == listID && l.DeletedAt == nil {
			todoLists[i].DeletedAt = &now
			todoLists[i].UpdatedAt = now
			return true
		}
	}
	return false
}
