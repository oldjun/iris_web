package models

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

type Admin struct {
	Id       int       `xorm:"not null pk autoincr INT(10)"`
	Username string    `xorm:"not null default '''' comment('用户名') unique VARCHAR(16)"`
	Phone    string    `xorm:"not null default '''' comment('手机号') index VARCHAR(16)"`
	Password string    `xorm:"not null default '''' comment('密码') VARCHAR(64)"`
	Role     int       `xorm:"not null default 0 comment('角色ID') index INT(10)"`
	Type     int       `xorm:"not null default 0 comment('类型:0=管理员,1=超级管理员') index TINYINT(3)"`
	Lock     int       `xorm:"not null default 0 comment('锁定:0=未锁定,1=已锁定') TINYINT(3)"`
	Time     time.Time `xorm:"not null default 'current_timestamp()' comment('注册时间') index TIMESTAMP"`
}

func (this *Admin) TableName() string {
	return "t_admin"
}

func (this *Admin) CheckPassword(password string) bool {
	hash := md5.New()
	hash.Write([]byte(password))
	passwordHash := hex.EncodeToString(hash.Sum(nil))
	return this.Password == passwordHash
}

func (this *Admin) SetPassword(password string) {
	hash := md5.New()
	hash.Write([]byte(password))
	passwordHash := hex.EncodeToString(hash.Sum(nil))
	this.Password = passwordHash
}

func (this *Admin) IsLock() bool {
	if this.Lock == 1 {
		return true
	} else {
		return false
	}
}

func (this *Admin) IsSuper() bool {
	if this.Type == 1 {
		return true
	} else {
		return false
	}
}
