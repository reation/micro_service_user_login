CREATE TABLE `user` (
                        `id` int(11) NOT NULL AUTO_INCREMENT,
                        `user_name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
                        `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
                        `nickname` varchar(255) NOT NULL DEFAULT '' COMMENT '昵称',
                        `sex` tinyint(4) NOT NULL DEFAULT '0' COMMENT '性别 0:男 1:女',
                        `birthday` date NOT NULL DEFAULT '0000-00-00' COMMENT '生日',
                        `state` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态 1:正常 2:冻结 3:注销 4:禁封 5:待审核 ',
                        `is_member` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否是会员 0:否 1:是',
                        `update_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
                        `create_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='用户表';
