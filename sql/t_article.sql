create table if not exists `t_article` (
    `id` int unsigned not null auto_increment,
    `uid` int unsigned not null default 0 comment '发布者',
    `title` varchar(32) not null default '' comment '标题',
    `image` varchar(64) not null default '' comment '封面',
    `brief` varchar(256) not null default '' comment '简介',
    `content` longtext comment '内容',
    `clicks` int unsigned not null default 0 comment '阅读',
    `status` tinyint unsigned not null default 0 comment '状态:0=隐藏,1=显示',
    `time` timestamp not null default current_timestamp comment '时间',
    primary key(`id`),
    key `idx_uid` (`uid`),
    key `idx_title` (`title`),
    key `idx_status` (`status`),
    key `idx_time` (`time`)
) engine=InnoDB default charset utf8mb4 comment '文章表';
