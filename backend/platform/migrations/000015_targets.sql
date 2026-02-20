-- +migrate Up
CREATE TABLE `{{.DB_TABLE_PREFIX}}targets`
(
  `id`              varchar(255) NOT NULL,
  `type`            varchar(50)  NOT NULL,
  `name`            varchar(255) NOT NULL,
  `description`     text,
  `links`           json,
  `internal_links`  json,
  `tags`            json,
  `internal_tags`   json,
  `synchronized_at` datetime(3),
  `created_at`      datetime(3),
  `updated_at`      datetime(3),
  PRIMARY KEY (`id`),
  INDEX             `idx_targets_type` (`type`),
  UNIQUE INDEX `idx_targets_name` (`name`)
);

-- +migrate Up
CREATE TABLE `{{.DB_TABLE_PREFIX}}target_addons`
(
  `id`         bigint unsigned AUTO_INCREMENT,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `target_id`  varchar(255) NOT NULL,
  `addon_type` varchar(50)  NOT NULL,
  `addon_id`   varchar(36)  NOT NULL,
  `config`     json         NOT NULL,
  PRIMARY KEY (`id`),
  INDEX        `idx_target_addons_target_id` (`target_id`),
  CONSTRAINT `fk_target_addons_target` FOREIGN KEY (`target_id`) REFERENCES `{{.DB_TABLE_PREFIX}}targets` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +migrate Up
CREATE TABLE `{{.DB_TABLE_PREFIX}}target_relations`
(
  `id`             bigint unsigned AUTO_INCREMENT,
  `from_target_id` varchar(255) NOT NULL,
  `to_target_id`   varchar(255) NOT NULL,
  `relation_type`  varchar(255) NOT NULL,
  `created_at`     datetime(3),
  `updated_at`     datetime(3),
  PRIMARY KEY (`id`),
  INDEX            `idx_target_relations_from_target_id` (`from_target_id`),
  INDEX            `idx_target_relations_to_target_id` (`to_target_id`),
  CONSTRAINT `fk_target_relations_from` FOREIGN KEY (`from_target_id`) REFERENCES `{{.DB_TABLE_PREFIX}}targets` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_target_relations_to` FOREIGN KEY (`to_target_id`) REFERENCES `{{.DB_TABLE_PREFIX}}targets` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `uniq_target_relations` UNIQUE (`from_target_id`, `to_target_id`, `relation_type`)
);

-- +migrate Up
CREATE TABLE `{{.DB_TABLE_PREFIX}}target_users`
(
  `id`         bigint unsigned AUTO_INCREMENT,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `target_id`  varchar(255) NOT NULL,
  `user_id`    bigint unsigned,
  `roles`      json         NOT NULL,
  PRIMARY KEY (`id`),
  INDEX        `idx_target_users_target_id` (`target_id`),
  INDEX        `idx_target_users_user_id` (`user_id`),
  CONSTRAINT `fk_target_users_target` FOREIGN KEY (`target_id`) REFERENCES `{{.DB_TABLE_PREFIX}}targets` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_target_users_user` FOREIGN KEY (`user_id`) REFERENCES `{{.DB_TABLE_PREFIX}}users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +migrate Up
ALTER TABLE `{{.DB_TABLE_PREFIX}}token_attempts`
  ADD COLUMN `target_id` varchar(255) NULL AFTER `server_id`;

-- +migrate Up
ALTER TABLE `{{.DB_TABLE_PREFIX}}token_attempts`
  ADD CONSTRAINT `fk_token_attempts_target` FOREIGN KEY (`target_id`) REFERENCES `{{.DB_TABLE_PREFIX}}targets` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE `{{.DB_TABLE_PREFIX}}token_attempts` DROP FOREIGN KEY `fk_token_attempts_target`;
-- +migrate Down
ALTER TABLE `{{.DB_TABLE_PREFIX}}token_attempts` DROP COLUMN `target_id`;
-- +migrate Down
ALTER TABLE `{{.DB_TABLE_PREFIX}}target_users` DROP FOREIGN KEY `fk_target_users_target`;
-- +migrate Down
ALTER TABLE `{{.DB_TABLE_PREFIX}}target_users` DROP FOREIGN KEY `fk_target_users_user`;
-- +migrate Down
ALTER TABLE `{{.DB_TABLE_PREFIX}}target_relations` DROP FOREIGN KEY `fk_target_relations_from`;
-- +migrate Down
ALTER TABLE `{{.DB_TABLE_PREFIX}}target_relations` DROP FOREIGN KEY `fk_target_relations_to`;
-- +migrate Down
ALTER TABLE `{{.DB_TABLE_PREFIX}}target_addons` DROP FOREIGN KEY `fk_target_addons_target`;
-- +migrate Down
DROP TABLE IF EXISTS `{{.DB_TABLE_PREFIX}}target_users`;
-- +migrate Down
DROP TABLE IF EXISTS `{{.DB_TABLE_PREFIX}}target_relations`;
-- +migrate Down
DROP TABLE IF EXISTS `{{.DB_TABLE_PREFIX}}target_addons`;
-- +migrate Down
DROP TABLE IF EXISTS `{{.DB_TABLE_PREFIX}}targets`;
