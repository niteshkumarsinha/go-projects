package main

import (
	"errors"
	"os"
	"strconv"
	"time"
	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) ValidateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		return errors.New("invalid index")
	}
	return nil
}

func (todos *Todos) Delete(index int) error {
	t := *todos
	if err := todos.ValidateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) Toggle(index int) error {
	if err := todos.ValidateIndex(index); err != nil {
		return err
	}
	todo := &(*todos)[index]
	todo.Completed = !todo.Completed
	if todo.Completed {
		completionTime := time.Now()
		todo.CompletedAt = &completionTime
	} else {
		todo.CompletedAt = nil
	}
	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	t := *todos
	if err := t.ValidateIndex(index); err != nil {
		return err
	}
	t[index].Title = title
	return nil
}

func (todos *Todos) List() {
	t := table.New(os.Stdout)
	t.SetRowLines(false)
	t.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for i, todo := range *todos {
		completed := "❌"
		completedAt := ""
		if todo.Completed {
			completed = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}

		t.AddRow(
			strconv.Itoa(i),
			todo.Title,
			completed,
			todo.CreatedAt.Format(time.RFC1123),
			completedAt,
		)
	}

	t.Render()
}
