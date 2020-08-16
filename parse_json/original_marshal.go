package main

import (
	"encoding/json"
	"strconv"
	"time"

	"golang.org/x/exp/errors/fmt"
)

type DueDate struct {
	time.Time
}

type Task struct {
	Name string  `json:"name"`
	Time DueDate `json:"time"`
	Done bool    `json:"done"`
}

// 日付のシリアライズ
func (d *DueDate) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(d.Unix()))), nil
}

type Todolist []Task

// done: falseのみmarshalされる
func (t Todolist) MarshalJSON() ([]byte, error) {
	activeTasks := make([]Task, 0, len(t))
	for _, task := range t {
		if !task.Done {
			activeTasks = append(activeTasks, task)
		}
	}

	return json.Marshal(activeTasks)
}

func main() {
	tasks := []Task{
		{
			Name: "lunch",
			Time: DueDate{Time: time.Now()},
			Done: true,
		},
		{
			Name: "dinner",
			Time: DueDate{Time: time.Now().Add(time.Hour * 6)},
			Done: false,
		},
	}

	d, _ := json.Marshal(Todolist(tasks))
	fmt.Println(string(d))
}
