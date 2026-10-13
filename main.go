package main

import "fmt"

func main() {
	storage := MemoryStorage{
		tasks:  make(map[int]Task),
		nextID: 1,
	}

	task := Task{
		Title:       "Изучить Go",
		Description: "Разобраться с указателями!!!",
		Done:        false,
	}

	storage.CreateTask(task)
	fmt.Println("Создана!")

	foundTask, err := storage.GetTask(1)
    if err != nil {
        fmt.Println("Ошибочка:", err)
    } else {
        fmt.Println("Есть:", foundTask)
    }

	foundTask.Done = true

    err = storage.UpdateTask(foundTask)
    if err != nil {
        fmt.Println("Ошибочка:", err)
    } else {
        fmt.Println("Обновлено!!!")
    }

	updatedTask, err := storage.GetTask(1)
	fmt.Println("Вот обновлённая) :", updatedTask)

	err = storage.DeleteTask(1)
	_, err = storage.GetTask(1)
    if err != nil {
        fmt.Println("Ожидаемо, так и надо", err)
    } else {
        fmt.Println("Не удалилась")
    }

}