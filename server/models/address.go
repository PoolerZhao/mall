package models

type Address struct {
	Id              uint   `gorm:"primaryKey"`
	UserId          string `gorm:"userId"`
	Name            string `gorm:"name"`
	Mobile          string `gorm:"mobile"`
	Province        string `gorm:"province"`
	City            string `gorm:"city"`
	District        string `gorm:"district"`
	DetailedAddress string `gorm:"detailedAddress"`
	IsDefault       int    `gorm:"IsDefault"`
	Created         string `gorm:"created"`
	Updated         string `gorm:"updated"`
}
