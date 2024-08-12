package main

import "time"

type doc struct {
	Id   string    `json:"id"`
	Num  string    `json:"num"`
	Date time.Time `json:"date"`
}
