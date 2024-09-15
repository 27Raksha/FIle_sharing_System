package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type File struct {
	gorm.Model
	UserID      uint
	Name        string
	Size        int64
	Location    string 
	UploadedAt  time.Time
	IsShared    bool
	SharedLink  string
}
