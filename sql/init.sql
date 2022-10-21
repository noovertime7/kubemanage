
create database `kubemanage`;

use  kubemanage;


-- auto-generated definition
create table t_admin
(
    id        int auto_increment
        primary key,
    user_name varchar(255) not null comment '用户名',
    salt      varchar(255) not null comment '盐',
    password  varchar(255) not null comment '密码',
    update_at datetime     not null comment '更新时间',
    create_at datetime     not null comment '创建时间',
    is_delete int          not null comment '是否删除',
    status    int(10)      not null comment '是否在线(1为在线)'
)
    charset = utf8;

INSERT INTO kubemanage.t_admin (id, user_name, salt, password, update_at, create_at, is_delete, status) VALUES (1, 'admin', 'admin', '29c09a3c055e47f704fb7c6df5b530e25f80ee3ab2a3ce44858284f929157389', '2022-07-05 13:01:35', '2022-06-15 20:55:55', 0, 1);

-- auto-generated definition
create table t_workflow
(
    id           int(32) auto_increment
        primary key,
    name         varchar(32) null,
    replicas     int(32)     null,
    namespace    varchar(32) null,
    deployment   varchar(32) null,
    service      varchar(32) null,
    ingress      varchar(32) null,
    service_type varchar(32) null,
    created_at   datetime(6) null,
    updated_at   datetime(6) null,
    deleted_at   datetime(6) null,
    is_deleted   int(32)     null
);

