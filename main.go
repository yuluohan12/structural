package main

import (
	"fmt"
	"time"
)

type Quenstiom struct {
	id        int64
	content   string
	CreatedAt time.Time
	DeletedAt time.Time
	Update    time.Time
}

func main() {
	var q Quenstiom
	q.id = 15204
	q.content = "我是大帅逼"
	q.CreatedAt = time.Now()
	q.Update = time.Now()
	q.DeletedAt = time.Now()
	fmt.Print(q.id, q.content, q.Update, q.CreatedAt, q.DeletedAt)
}
