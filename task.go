package main

import (
	"time"
)

type Task struct{
	id int
	title string
	description string
	add_date time.Time
	deadline time.Time
	done bool

}

func main()