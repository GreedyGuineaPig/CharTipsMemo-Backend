package model

import (
	"strconv"

	"gorm.io/gorm"
)

type Post struct {
	ID         uint   `json:"id" gorm:"AUTO_INCREMENT;primaryKey"`
	Version    string `json:"version" gorm:"type:varchar(255);not null"`
	Char       string `json:"char" gorm:"not null"`
	IsAntiChar bool   `json:"anti"`
	Body       string `json:"body" gorm:"type:varchar(255);not null"`
}

// may need to implement isAntiChar filter

func (p *Post) FirstById(id string) (db *gorm.DB) {
	uintID, _ := strconv.ParseUint(id, 10, 32)
	return DB.First(&p, uintID)
}

func (p *Post) Find() (db *gorm.DB) {
	return DB.Find(&p)
}

func FindAll() (ps []Post) {
	var posts []Post
	DB.Find(&posts)
	return posts
}

func FindAllByChar(char string) (ps []Post) {
	var posts []Post
	DB.Find(&posts, Post{Char: char})
	return posts
}

func (p *Post) Create() (db *gorm.DB) {
	return DB.Create(&p)
}

func (p *Post) Save() (db *gorm.DB) {
	return DB.Save(&p)
}

func (p *Post) Updates() (db *gorm.DB) {
	return DB.Model(&p).Updates(p)
}

func (p *Post) Delete() (db *gorm.DB) {
	return DB.Delete(&p)
}

// func (p *Post) DeleteById(id uint) (db *gorm.DB) {
// 	return DB.Where("id = ?", id).Delete(&p)
// }
