CREATE TABLE `comment`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uid`        bigint NOT NULL COMMENT 'uid',
    `target_uid`        bigint NOT NULL COMMENT 'target_uid',
    `comment_id` bigint NOT NULL COMMENT 'comment_id',
    `target_article_id` bigint NOT NULL COMMENT 'target_article_id',
    `target_comment_id` bigint DEFAULT 0 COMMENT 'target_comment_id 二级评论时候不为空',
    `comment_type` tinyint(3) NOT NULL DEFAULT 0 COMMENT 'comment_type 0:一级评论(文章评论) 1:二级评论(评论子评论)',
    `status`     tinyint(3) NOT NULL DEFAULT 0 COMMENT "status 0:待审核 1:审核通过 2:审核未通过",
    `body`       varchar(256) NOT NULL DEFAULT '' COMMENT 'comment body',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'comment create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'comment update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'comment delete time',
    PRIMARY KEY  (`id`),
    UNIQUE KEY   `target_article_id_comment_id` (`target_article_id`, `comment_id`),
    KEY          `uid` (`uid`),
    KEY          `comment_id` (`comment_id`),
    KEY          `target_comment_id` (`target_comment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='comment info table';


