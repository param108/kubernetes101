package main

type Line struct {
	ID   int    `json:"id",gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Data string `json:"data"`
}
