package models

import (
	"time"
)

type AdminAction struct {
	Id     int       `xorm:"not null pk autoincr INT(10)"`
	Mid    int       `xorm:"not null default 0 comment('模块') index INT(10)"`
	Name   string    `xorm:"not null default '''' comment('名称') index VARCHAR(32)"`
	Action string    `xorm:"not null default '''' comment('方法') index VARCHAR(32)"`
	Sort   int       `xorm:"not null default 0 comment('排序') index TINYINT(3)"`
	Menu   int       `xorm:"not null default 0 comment('菜单:0=否,1=是') index TINYINT(3)"`
	Time   time.Time `xorm:"not null default 'current_timestamp()' comment('时间') index TIMESTAMP"`
}

func (this *AdminAction) TableName() string {
	return "t_admin_action"
}
