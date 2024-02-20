CREATE TABLE `article`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uid`        bigint NOT NULL COMMENT 'uid',
    `article_id` bigint NOT NULL COMMENT 'article_id',
    `title`      varchar(64) NOT NULL DEFAULT '' COMMENT 'title',
    `description`varchar(128) NOT NULL DEFAULT '' COMMENT 'description',
    `body`       varchar(1024) NOT NULL DEFAULT '' COMMENT 'body',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'article create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'article update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'article delete time',
    PRIMARY KEY  (`id`),
    UNIQUE KEY   `article_id` (`article_id`),
    KEY          `uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='article info table';

CREATE TABLE `article_tag_map`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `article_id` bigint NOT NULL COMMENT 'article_id',
    `tag_name`   varchar(64) NOT NULL DEFAULT '' COMMENT 'tag_name',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'tag create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'tag update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'tag delete time',
    PRIMARY KEY  (`id`),
    KEY          `article_id` (`article_id`),
    KEY          `tag_name` (`tag_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='article_tag_map info table';

CREATE TABLE `article_favourite_info`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uid`        bigint NOT NULL COMMENT 'uid',
    `article_id` bigint NOT NULL COMMENT 'article_id',
    `status`     tinyint(3) NOT NULL DEFAULT 0 COMMENT "status 0:未操作 1:喜欢 2:不喜欢";
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'delete time',
    PRIMARY KEY  (`id`),
    UNIQUE KEY   `uid_article_id` (`uid`,`article_id`),
    KEY          `article_id` (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='article_favourite info table';


