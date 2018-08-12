package model

import (
	"sync"
	"time"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/10 下午7:22'
*/

type BaseModel struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	CreatedAt time.Time  `gorm:"column:createdAt"`
	UpdatedAt time.Time  `gorm:"column:updatedAt"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index"`
}

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"say_hello"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}
