create table if not exists `t_article_cate` (
    `id` int unsigned not null auto_increment,
    `pid` int unsigned not null default 0 comment '上级',
    `name` varchar(32) not null default '' comment '名称',
    `sort` int unsigned not null default 0 comment '排序',
    `time` timestamp not null default current_timestamp comment '时间',
    primary key(`id`),
    key `idx_pid` (`pid`),
    key `idx_name` (`name`),
    key `idx_sort` (`sort`),
    key `idx_time` (`time`)
) engine=InnoDB default charset utf8 comment '文章分类表';
