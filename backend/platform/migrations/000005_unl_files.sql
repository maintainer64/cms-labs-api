-- +migrate Up
CREATE TABLE `{{.DB_TABLE_PREFIX}}unl_files`
(
  `id`         bigint unsigned AUTO_INCREMENT,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `path`       longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `type`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `deleted_at` datetime(3),
  `synced_id`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `content`    mediumblob,
  PRIMARY KEY (`id`)
);

-- +migrate Down
drop table if exists `{{.DB_TABLE_PREFIX}}unl_files` cascade;
