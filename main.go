package main

import (
	"fmt"
	"study/task/domain"
	"time"
)

var Tasks []*domain.Task

func main() {
	task1, error1 := domain.NewTask(1, "Goのお勉強", time.Time{})
	if error1 != nil {
		println(error1.Error())
		return
	}
	task2, error2 := domain.NewTask(2, "Rustのお勉強", time.Now().AddDate(0, 0, 1))
	if error2 != nil {
		println(error2.Error())
		return
	}
	task3, error3 := domain.NewTask(3, "kotlinのお勉強", time.Date(2022, 10, 1, 0, 0, 0, 0, &time.Location{}))
	if error3 != nil {
		println(error3.Error())
		return
	}

	Tasks = append(Tasks, task1)
	Tasks = append(Tasks, task2)
	Tasks = append(Tasks, task3)

	fmt.Println("全部表示")
	for i, task := range Tasks {
		fmt.Println("---------------------------------------------------")
		fmt.Printf("index: %d\n", i)
		fmt.Printf("taskId: %d\n", task.GetId())
		fmt.Printf("taskTitle: %s\n", task.GetTitle())
		year, month, date := task.GetExpirationDate().Date()
		fmt.Printf("taskExpirationDate: %d/%d/%d\n", year, month, date)
	}
}
