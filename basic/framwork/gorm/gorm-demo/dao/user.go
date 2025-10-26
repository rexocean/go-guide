package dao

import "log"

// User 定义User模型，绑定users表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
type User struct {
	ID int64
	//通过在字段后面的标签说明，定义golang字段和表字段的关系
	//例如 `gorm:"column:username"` 标签说明含义是: Mysql表的列名（字段名)为username
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:createtime"`
}

func (u User) TableName() string {
	return "users"
}

func SaveUser(user *User) {
	// 数据库的操作， 链接数据库
	err := DB.Create(user).Error
	if err != nil {
		log.Println("insert user error", err)
	}
}
