package model

import "github.com/jinzhu/gorm"

// "golang.org/x/crypto/bcrypt"

// Video 视频模型
type Video struct {
	gorm.Model
	VideoID    int
	VideoTitle string
	VideoInfo  string
}
