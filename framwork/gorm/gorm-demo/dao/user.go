package dao

import (
	"gorm.io/gorm"
	"log"
)

// User 定义User模型，绑定users表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
type User struct {
	// gorm默认的字段：id、更新时间、创建时间、删除时间 组合的方式
	//gorm.Model
	ID int64
	//通过在字段后面的标签说明，定义golang字段和表字段的关系
	//例如 `gorm:"column:username"` 标签说明含义是: Mysql表的列名（字段名)为username
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:createtime"`
	Admin      bool   `gorm:"-"`
	// 默认填充无需业务填充
	//CreatedAt  time.Time
}

func (u User) TableName() string {
	return "users"
}

func UserTable(user User) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if user.Admin {
			return tx.Table("admin_users")
		}

		return tx.Table("users")
	}
}

func SaveUser(user *User) {
	// 数据库的操作， 链接数据库
	// 动态表名
	//err := DB.Scopes(UserTable(*user)).Create(user).Error
	//err := DB.Table("admin_users").Create(user).Error
	//err := DB.Create(user).Error
	// 操作之后如果有err就返回err
	tx := DB.Create(user)
	// 批量插入和批量分批插入
	//users := []User{}
	//tx := DB.Create(users)
	//DB.CreateInBatches(user,2)
	err := tx.Error
	affected := tx.RowsAffected
	log.Println("affected rows:", affected)

	if err != nil {
		log.Println("insert user error", err)
	}
}

func GetById(id int64) User {
	var user User
	err := DB.Where("id=?", id).First(&user)
	if err != nil {
		log.Println("get user error", err)
	}
	return user
}

func GetAll() []User {
	var users []User
	err := DB.Find(&users).Error
	if err != nil {
		log.Println("get users error", err)
	}
	return users
}

func Update(id int64) {
	err := DB.Model(&User{}).Where("id = ?", id).Update("username", "lisi")
	if err != nil {
		log.Println("update user error", err)
	}
}

func Delete(id int64) {
	err := DB.Where("id = ?", id).Delete(&User{})
	if err != nil {
		log.Println("delete user error", err)
	}
}
