-- +migrate Up
CREATE TABLE `{{.DB_TABLE_PREFIX}}roles`
(
  `id`         bigint unsigned AUTO_INCREMENT,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `code`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `name`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`)
);

-- +migrate Up
CREATE TABLE `{{.DB_TABLE_PREFIX}}role_relations`
(
  `id`         bigint unsigned AUTO_INCREMENT,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `role_id`    bigint unsigned,
  `user_id`    bigint unsigned NULL,
  `server_id`  bigint unsigned NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `{{.DB_TABLE_PREFIX}}_role_relations_ibfk_role` FOREIGN KEY (`role_id`) REFERENCES `{{.DB_TABLE_PREFIX}}roles` (`id`) ON DELETE CASCADE,
  CONSTRAINT `{{.DB_TABLE_PREFIX}}_role_relations_ibfk_user` FOREIGN KEY (`user_id`) REFERENCES `{{.DB_TABLE_PREFIX}}users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `{{.DB_TABLE_PREFIX}}_role_relations_ibfk_server` FOREIGN KEY (`server_id`) REFERENCES `{{.DB_TABLE_PREFIX}}pnet_servers` (`id`) ON DELETE CASCADE,
  CONSTRAINT `{{.DB_TABLE_PREFIX}}_role_relations_xor` CHECK (
    (`user_id` IS NOT NULL AND `server_id` IS NULL) OR
    (`user_id` IS NULL AND `server_id` IS NOT NULL)
    )
);


-- +migrate Down
drop table if exists `{{.DB_TABLE_PREFIX}}role_relations` cascade;

-- +migrate Down
drop table if exists `{{.DB_TABLE_PREFIX}}roles` cascade;


