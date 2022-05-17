create table if not exists `t_log_detail` (
    `id` int unsigned not null auto_increment,
    `pid` int unsigned not null default 0 comment '主表ID',
    `key` varchar(64) not null default '' comment '字段',
    `old_value` longtext comment '旧值',
    `new_value` longtext comment '新值',
    primary key(`id`),
    key `idx_pid` (`pid`)
) engine=InnoDB default charset utf8mb4 comment '日志详情表';
