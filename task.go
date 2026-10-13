package main

import (
	"time"
)

type Task struct{
	Id int
	Title string
	Description string
	Create_date time.Time
	Deadline time.Time
	Done bool

}