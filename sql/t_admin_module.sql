create table if not exists `t_admin_module` (
    `id` int unsigned not null auto_increment,
    `name` varchar(32) not null default '' comment '名称',
    `sort` tinyint unsigned not null default 0 comment '排序',
    `time` timestamp not null default current_timestamp comment '时间',
    primary key(`id`),
    key `idx_name` (`name`),
    key `idx_sort` (`sort`),
    key `idx_time` (`time`)
) engine=InnoDB default charset utf8 comment '模块表';
