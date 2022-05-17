create table if not exists `t_article_list` (
    `id` int unsigned not null auto_increment,
    `pid` int unsigned not null default 0 comment '父类ID',
    `cid` int unsigned not null default 0 comment '子类ID',
    `aid` int unsigned not null default 0 comment '文章ID',
    `time` timestamp not null default current_timestamp comment '时间',
    primary key(`id`),
    key `idx_pid` (`pid`),
    key `idx_cid` (`cid`),
    key `idx_aid` (`aid`),
    key `idx_time` (`time`)
) engine=InnoDB default charset utf8 comment '文章列表';
