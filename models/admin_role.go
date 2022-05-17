package models

import (
	"time"
)

type AdminRole struct {
	Id   int       `xorm:"not null pk autoincr INT(10)"`
	Name string    `xorm:"not null default '''' comment('角色名称') index VARCHAR(16)"`
	Time time.Time `xorm:"not null default 'current_timestamp()' comment('时间') index TIMESTAMP"`
}

func (this *AdminRole) TableName() string {
	return "t_admin_role"
}
