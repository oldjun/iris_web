package models

import (
	"time"
)

type AdminAuth struct {
	Id     int       `xorm:"not null pk autoincr INT(10)"`
	Role   int       `xorm:"not null default 0 comment('角色ID') index INT(10)"`
	Action int       `xorm:"not null default 0 comment('权限ID') index INT(10)"`
	Time   time.Time `xorm:"not null default 'current_timestamp()' comment('时间') index TIMESTAMP"`
}

func (this *AdminAuth) TableName() string {
	return "t_admin_auth"
}
