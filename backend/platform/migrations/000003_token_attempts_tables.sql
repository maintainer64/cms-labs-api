-- +migrate Up
CREATE TABLE `{{.DB_TABLE_PREFIX}}token_attempts`
(
  `id`                 bigint unsigned AUTO_INCREMENT,
  `created_at`         datetime(3),
  `updated_at`         datetime(3),
  `user_id`            bigint unsigned,
  `server_id`          bigint,
  `token`              varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  INDEX                `idx_{{.DB_TABLE_PREFIX}}token_attempts_token` (`token`),
  `state`              varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `authorization_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  INDEX                `idx_{{.DB_TABLE_PREFIX}}token_attempts_authorization_code` (`authorization_code`),
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`)
    REFERENCES `{{.DB_TABLE_PREFIX}}users` (`id`)
    ON DELETE CASCADE
);

-- +migrate Down
drop table if exists `{{.DB_TABLE_PREFIX}}token_attempts` cascade;
