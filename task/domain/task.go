package domain

import (
	"errors"
	"time"
)

type Task struct {
	id              int
	title           string
	expiredDateTime time.Time
}

func NewTask(id int, title string, expiredDateTime time.Time) (*Task, error) {
	if title == "" {
		return nil, errors.New("タイトルを指定してください")
	}

	task := &Task{
		id:              id,
		title:           title,
		expiredDateTime: expiredDateTime,
	}
	return task, nil
}

func (task Task) GetTitle() string {
	return task.title
}

func (task Task) GetId() int {
	return task.id
}

func (task Task) GetExpirationDate() time.Time {
	return task.expiredDateTime
}

func (task *Task) SetExpirationDate(expirationDate time.Time) {
	task.expiredDateTime = expirationDate
}
