create table if not exists `t_admin_action` (
    `id` int unsigned not null auto_increment,
    `mid` int unsigned not null default 0 comment '模块',
    `name` varchar(32) not null default '' comment '名称',
    `action` varchar(32) not null default '' comment '方法',
    `menu` tinyint unsigned not null default 0 comment '菜单:0=否,1=是',
    `sort` tinyint unsigned not null default 0 comment '排序',
    `time` timestamp not null default current_timestamp comment '时间',
    primary key(`id`),
    key `idx_mid` (`mid`),
    key `idx_name` (`name`),
    key `idx_action` (`action`),
    key `idx_sort` (`sort`),
    key `idx_time` (`time`)
) engine=InnoDB default charset utf8 comment '权限表';
