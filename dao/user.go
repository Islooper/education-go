package dao

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Id        int            `gorm:"column:id" db:"id" json:"id" form:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at" form:"deleted_at"`
	UserType  int8           `gorm:"column:user_type" db:"user_type" json:"user_type" form:"user_type"` //用户类型:1.学生,2.教师
	Name      string         `gorm:"column:name" db:"name" json:"name" form:"name"`                     //姓名
	EMail     string         `gorm:"column:e_mail" db:"e_mail" json:"e_mail" form:"e_mail"`             //邮件地址
	Ip        string         `gorm:"column:ip" db:"ip" json:"ip" form:"ip"`                             //局域网内ip地址
	UserName  string         `gorm:"column:user_name" db:"user_name" json:"user_name" form:"user_name"` //账号名
	Password  string         `gorm:"column:password" db:"password" json:"password" form:"password"`     //密码
}

type UserDao struct {
}

func (*Users) TableName() string {
	return "users"
}

func (u *UserDao) ReadByUserNameAndPass(userName, password string) (*Users, error) {
	if userName == "" {
		return nil, errors.New("参数为空")
	}

	var userDo Users
	err := Db.Model(Users{}).Where("user_name = ? and password = ?  ", userName, password).Find(&userDo).Error
	if err != nil {
		return nil, err
	}

	return &userDo, nil
}

func (u *UserDao) UpdateIp(userDo *Users) error {
	return Db.Model(Users{}).Where("id = ? ", userDo.Id).Updates(userDo).Error
}
