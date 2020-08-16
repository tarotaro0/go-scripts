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
}

func (d *DueDate) UnmarshalJSON(raw []byte) error {
	epoch, err := strconv.Atoi(string(raw))
	if err != nil {
		return err
	}

	d.Time = time.Unix(int64(epoch), 0)
	return nil
}

var jsonString = []byte(`[
	{"name": "lunch", "time": 1486600200}
]`)

func main() {
	var tasks []Task
	if err := json.Unmarshal(jsonString, &tasks); err != nil {
		panic(nil)
	}

	for _, task := range tasks {
		fmt.Printf("%s: %v\n", task.Name, task.Time)
	}
}
