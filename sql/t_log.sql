create table if not exists `t_log` (
    `id` int unsigned not null auto_increment,
    `uid` int unsigned not null default 0 comment '操作用户',
    `name` varchar(64) not null default '' comment '表名',
    `ip` varchar(64) not null default '' comment '访问IP',
    `type` tinyint unsigned not null default 0 comment '操作类型:新增=1,修改=2,删除=3',
    `time` timestamp not null default current_timestamp comment '时间',
    primary key(`id`),
    key `idx_uid` (`uid`),
    key `idx_name` (`name`),
    key `idx_type` (`type`),
    key `idx_time` (`time`)
) engine=InnoDB default charset utf8 comment '日志表';
