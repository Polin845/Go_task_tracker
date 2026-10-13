package main

import (
	"fmt"
)

type Storage interface {
	 CreateTask(task Task) (Task, error)
    GetTask(id int) (Task, error)
    UpdateTask(task Task) error
    DeleteTask(id int) error
}

type MemoryStorage struct {
	tasks map[int]Task
	nextID int // Чтобы id автоматически создавались для тасков,а не вводились руками
}

func (s *MemoryStorage) CreateTask(task Task) {
	task.Id = s.nextID
	s.tasks[s.nextID] = task
	s.nextID++
}

func (s *MemoryStorage) GetTask(id int) (Task, error) {
    task, exists := s.tasks[id]

    if !exists {
        return Task{}, fmt.Errorf("task with ID %d not found", id)
    }

    return task, nil
}

func (s *MemoryStorage) UpdateTask(task Task) error {
    _, exists := s.tasks[task.Id]

    if !exists {
        return fmt.Errorf("task with Id %d not found", task.Id)
    }

    s.tasks[task.Id] = task

    return nil
}

func (s *MemoryStorage) DeleteTask(id int) error {
    _, exists := s.tasks[id]

    if !exists {
        return fmt.Errorf("task with ID %d not found", id)
    }

    delete(s.tasks, id)

    return nil
}