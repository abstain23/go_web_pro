CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(64) COLLATE utf8mb4_general_ci,
  `gender` tinyint(4) NOT NULL DEFAULT '0',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`) USING BTREE,
  UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



CREATE TABLE `community` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `community_id` int(10) unsigned NOT NULL,
  `community_name` VARCHAR(128) COLLATE utf8mb4_general_ci NOT NULL,
  `introduction`  VARCHAR(256) COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(`id`),
  UNIQUE KEY `idx_community_id` (`community_id`),
  UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `community` VALUES('1', '1', 'go', 'golang', '2023-03-12 21:51:00', '2023-03-12 21:51:00');

INSERT INTO `community` VALUES('2', '2', 'leetcode', '', '2023-03-12 21:52:00', '2023-03-12 21:52:00');



CREATE TABLE `post` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `post_id` bigint(20) NOT NULL COMMENT '帖子id',
  `title` VARCHAR(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `content` VARCHAR(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `author_id` bigint(20)  NOT NULL COMMENT '作者的用户id',
  `community_id` bigint(20) NOT NULL COMMENT '所属社区id',
  `status` TINYINT(4) NOT NULL DEFAULT '1' COMMENT '帖子状态',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY(`id`),
  UNIQUE KEY `idx_post_id` (`post_id`),
  KEY `idx_author_id` (`author_id`),
  KEY `idx_community_id` (`community_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
