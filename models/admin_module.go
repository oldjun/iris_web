package models

import (
	"time"
)

type AdminModule struct {
	Id   int       `xorm:"not null pk autoincr INT(10)"`
	Name string    `xorm:"not null default '''' comment('名称') index VARCHAR(32)"`
	Sort int       `xorm:"not null default 0 comment('排序') index TINYINT(3)"`
	Time time.Time `xorm:"not null default 'current_timestamp()' comment('时间') index TIMESTAMP"`
}

func (this *AdminModule) TableName() string {
	return "t_admin_module"
}
