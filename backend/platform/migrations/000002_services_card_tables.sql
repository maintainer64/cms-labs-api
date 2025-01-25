-- +migrate Up
CREATE TABLE `{{.DB_TABLE_PREFIX}}service_cards`
(
  `id`          bigint unsigned AUTO_INCREMENT,
  `created_at`  datetime(3),
  `updated_at`  datetime(3),
  `image_url`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `url`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `name`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `order`       bigint,
  `is_active`   boolean,
  PRIMARY KEY (`id`)
);

-- +migrate Down
drop table if exists `{{.DB_TABLE_PREFIX}}service_cards` cascade;
