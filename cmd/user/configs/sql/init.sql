CREATE TABLE `user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uid`        bigint NOT NULL COMMENT 'uid',
    `user_name`  varchar(128) NOT NULL DEFAULT '' COMMENT 'UserName',
    `email`      varchar(128) NOT NULL DEFAULT '' COMMENT 'email',
    `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
    `bio`        varchar(256) NOT NULL DEFAULT '' COMMENT 'bio',
    `image`      varchar(256) NOT NULL DEFAULT '' COMMENT 'image',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`uid`),
    UNIQUE KEY `email` (`email`),
    KEY          `idx_user_name` (`user_name`) COMMENT 'UserName index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';