CREATE TABLE `follow_info`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uid`        bigint NOT NULL COMMENT 'uid',
    `follow_uid` bigint NOT NULL COMMENT 'follow_uid',
    `status`     tinyint(3) NOT NULL COMMENT 'status 0:未关注, 1:已关注, 2:拉黑态',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'follow info create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'follow info update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'follow info delete time',
    PRIMARY KEY (`id`),
    KEY (`follow_uid`),
    UNIQUE KEY          `uid_follow_uid` (`uid`, `follow_uid`) COMMENT 'uid->follow_uid index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='follow info table';