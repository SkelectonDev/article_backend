package models

import (
	"time"
)

// Model definition same as gorm.Model, but including column and json tags
type Model struct {
	ID        uint       `gorm:"primary_key;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

// article Post
type Post struct {
	Model
	Title          string    `gorm:"column:title" json:"title"`
	Text           string    `gorm:"column:text" json:"text"`
	Author         string    `gorm:"column:author" json:"author"`
	Next           *Post     `gorm:"column:next" json:"next"`
	NextPostID     uint      `gorm:"type:int REFERENCES posts(id) ON DELETE CASCADE;column:next_post_id" json:"next_post_id"`
	Previous       *Post     `gorm:"column:previous" json:"previous"`
	PreviousPostID uint      `gorm:"type:int REFERENCES posts(id) ON DELETE CASCADE;column:previous_post_id" json:"previous_post_id"`
	PostedOn       time.Time `gorm:"column:posted_on" json:"posted_on"`
	Sections       []Section `gorm:"column:sections" json:"sections"`
	Tags           []Tag     `gorm:"column:tags" json:"tags"`
}

// Section of article Post (headings)
type Section struct {
	Model
	PostID uint   `gorm:"type:int REFERENCES posts(id) ON DELETE CASCADE;column:post_id" json:"post_id"`
	Name   string `gorm:"column:name" json:"name"`
}

// Project that I developed
type Project struct {
	Model
	Name                string `gorm:"column:name" json:"name"`
	ThumbnailPictureURL string `gorm:"column:thumbnail_url" json:"src"`
	URL                 string `gorm:"column:url" json:"url"`
	Description         string `gorm:"column:description" json:"description"`
	Tags                []Tag  `gorm:"column:tags" json:"tags"`
}

// Tag of article Post (hashtag)
type Tag struct {
	Model
	PostID    uint   `gorm:"type:int REFERENCES posts(id) ON DELETE CASCADE;column:post_id" json:"post_id"`
	ProjectID uint   `gorm:"type:int REFERENCES projects(id) ON DELETE CASCADE;column:project_id" json:"project_id"`
	Name      string `gorm:"column:name" json:"name"`
}

// Book that I Read
type Book struct {
	Model
	Title           string `gorm:"column:title" json:"title"`
	CoverPictureURL string `gorm:"column:cover_url" json:"src"`
	URL             string `gorm:"column:url" json:"url"`
	AlternativeText string `gorm:"column:alt" json:"alt"`
}
