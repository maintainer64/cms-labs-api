-- +migrate Up
CREATE TABLE `{{.DB_TABLE_PREFIX}}lti_access_tokens`
(
  `id`         bigint unsigned AUTO_INCREMENT,
  `index`      longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `payload`    longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  PRIMARY KEY (`id`),
  INDEX        `idx_{{.DB_TABLE_PREFIX}}lti_access_tokens` (`index`(768)),
  INDEX        `idx_{{.DB_TABLE_PREFIX}}lti_access_tokens_created_at` (`created_at`)
);

CREATE TABLE `{{.DB_TABLE_PREFIX}}lti_forms`
(
  `id`                 bigint unsigned AUTO_INCREMENT,
  `created_at`         datetime(3),
  `updated_at`         datetime(3),
  `name`               varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `lti_client_id`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `lti_deployment_id`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `base_uri`           varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `lti_auth_token_uri` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `lti_auth_login_uri` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `key_set_uri`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `target_link_uri`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `public_key`         text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `private_key`        text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`),
  INDEX                `idx_{{.DB_TABLE_PREFIX}}lti_forms_lti_client_id` (`lti_client_id`),
  INDEX                `idx_{{.DB_TABLE_PREFIX}}lti_forms_lti_deployment_id` (`lti_deployment_id`),
  INDEX                `idx_{{.DB_TABLE_PREFIX}}lti_forms_lti_base_uri` (`base_uri`),
  INDEX                `idx_{{.DB_TABLE_PREFIX}}lti_forms_created_at` (`created_at`)
);

CREATE TABLE `{{.DB_TABLE_PREFIX}}lti_launch_data`
(
  `id`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `lti_form_id` bigint unsigned,
  `created_at`  datetime(3),
  `updated_at`  datetime(3),
  `launch_data` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`),
  INDEX         `idx_{{.DB_TABLE_PREFIX}}lti_launch_data_lti_form_id` (`lti_form_id`),
  INDEX         `idx_{{.DB_TABLE_PREFIX}}lti_launch_data_created_at` (`created_at`),
  INDEX         `idx_{{.DB_TABLE_PREFIX}}lti_launch_data_updated_at` (`updated_at`)
);

CREATE TABLE `{{.DB_TABLE_PREFIX}}lti_nonce_tokens`
(
  `id`              bigint unsigned AUTO_INCREMENT,
  `created_at`      datetime(3),
  `updated_at`      datetime(3),
  `nonce`           varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `target_link_uri` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`),
  INDEX             `idx_{{.DB_TABLE_PREFIX}}lti_nonce_tokens_nonce` (`nonce`),
  INDEX             `idx_{{.DB_TABLE_PREFIX}}lti_nonce_tokens_created_at` (`created_at`)
);

CREATE TABLE `{{.DB_TABLE_PREFIX}}pnet_servers`
(
  `id`                     bigint unsigned AUTO_INCREMENT,
  `created_at`             datetime(3),
  `updated_at`             datetime(3),
  `name`                   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `url`                    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `is_active`              boolean,
  `minutes_for_disconnect` bigint,
  `max_count_users_limit`  bigint,
  `last_online_status`     datetime(3),
  `last_count_users`       bigint,
  `unit_rate`              bigint,
  `token`                  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_{{.DB_TABLE_PREFIX}}pnet_servers_token` (`token`),
  INDEX                    `idx_{{.DB_TABLE_PREFIX}}pnet_servers_created_at` (`created_at`),
  INDEX                    `idx_{{.DB_TABLE_PREFIX}}pnet_servers_updated_at` (`updated_at`)
);

CREATE TABLE `{{.DB_TABLE_PREFIX}}users`
(
  `id`             bigint unsigned AUTO_INCREMENT,
  `created_at`     datetime(3),
  `updated_at`     datetime(3),
  `deleted_at`     datetime(3),
  `email`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `name`           varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `user_role`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `group_name`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `lti_user_id`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `last_launch_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_{{.DB_TABLE_PREFIX}}users_email` (`email`),
  INDEX            `idx_{{.DB_TABLE_PREFIX}}users_last_launch_id` (`last_launch_id`),
  INDEX            `idx_{{.DB_TABLE_PREFIX}}users_created_at` (`created_at`)
);

CREATE TABLE `{{.DB_TABLE_PREFIX}}user_tokens`
(
  `user_id`       bigint unsigned,
  `refresh_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `hash_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `created_at`    datetime(3),
  `updated_at`    datetime(3),
  PRIMARY KEY (`user_id`),
  FOREIGN KEY (`user_id`)
    REFERENCES `{{.DB_TABLE_PREFIX}}users` (`id`)
    ON DELETE CASCADE,
  INDEX           `idx_{{.DB_TABLE_PREFIX}}user_tokens_refresh_token` (`refresh_token`),
  INDEX           `idx_{{.DB_TABLE_PREFIX}}user_tokens_created_at` (`created_at`)
);

CREATE TABLE `{{.DB_TABLE_PREFIX}}round_queue_pools`
(
  `id`             bigint unsigned AUTO_INCREMENT,
  `created_at`     datetime(3),
  `updated_at`     datetime(3),
  `type`           varchar(255),
  `connected_at`   datetime(3),
  `last_used`      boolean,
  `is_active`      boolean,
  `pnet_server_id` bigint unsigned  null,
  FOREIGN KEY (`pnet_server_id`)
    REFERENCES `{{.DB_TABLE_PREFIX}}pnet_servers` (`id`)
    ON DELETE CASCADE,
  PRIMARY KEY (`id`),
  INDEX            `idx_{{.DB_TABLE_PREFIX}}round_queue_pools_connected_at` (`connected_at`),
  INDEX            `idx_{{.DB_TABLE_PREFIX}}round_queue_pools_type` (`type`),
  INDEX            `idx_{{.DB_TABLE_PREFIX}}round_queue_pools_is_active` (`is_active`),
  INDEX            `idx_{{.DB_TABLE_PREFIX}}round_queue_pools_pnet_server_id` (`pnet_server_id`),
  INDEX            `idx_{{.DB_TABLE_PREFIX}}round_queue_pools_created_at` (`created_at`)
);

CREATE TABLE `{{.DB_TABLE_PREFIX}}lti_routings`
(
  `id`                     bigint unsigned AUTO_INCREMENT,
  `created_at`             datetime(3),
  `updated_at`             datetime(3),
  `name`                   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `lti_title`              varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `lti_description`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `lti_task_id`            varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `lti_params_task`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `collaboration`          bigint,
  `pinned_session_minutes` bigint,
  `pnet_labs_path`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `pnet_test_path`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`),
  INDEX                    `idx_{{.DB_TABLE_PREFIX}}lti_routings_created_at` (`created_at`),
  INDEX                    `idx_{{.DB_TABLE_PREFIX}}lti_routings_lti_title` (`lti_title`),
  INDEX                    `idx_{{.DB_TABLE_PREFIX}}lti_routings_lti_task_id` (`lti_task_id`),
  INDEX                    `idx_{{.DB_TABLE_PREFIX}}lti_routings_lti_params_task` (`lti_params_task`)
);

CREATE TABLE `{{.DB_TABLE_PREFIX}}lti_attempts`
(
  `id`             bigint unsigned AUTO_INCREMENT,
  `created_at`     datetime(3),
  `updated_at`     datetime(3),
  `user_id`        bigint unsigned,
  `pnet_server_id` bigint unsigned,
  `lti_routing_id` bigint unsigned,
  `expired_at`     datetime(3),
  `room_number`    bigint,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`)
    REFERENCES `{{.DB_TABLE_PREFIX}}users` (`id`)
    ON DELETE CASCADE,
  FOREIGN KEY (`pnet_server_id`)
    REFERENCES `{{.DB_TABLE_PREFIX}}pnet_servers` (`id`)
    ON DELETE CASCADE,
  FOREIGN KEY (`lti_routing_id`)
    REFERENCES `{{.DB_TABLE_PREFIX}}lti_routings` (`id`)
    ON DELETE CASCADE,
  INDEX            `idx_{{.DB_TABLE_PREFIX}}lti_attempts_created_at` (`created_at`),
  INDEX            `idx_{{.DB_TABLE_PREFIX}}lti_attempts_expired_at` (`expired_at`)
);


-- +migrate Down

drop table if exists `{{.DB_TABLE_PREFIX}}lti_attempts` cascade;

drop table if exists `{{.DB_TABLE_PREFIX}}lti_access_tokens` cascade;

drop table if exists `{{.DB_TABLE_PREFIX}}lti_forms` cascade;

drop table if exists `{{.DB_TABLE_PREFIX}}lti_launch_data` cascade;

drop table if exists `{{.DB_TABLE_PREFIX}}lti_nonce_tokens` cascade;

drop table if exists `{{.DB_TABLE_PREFIX}}pnet_servers` cascade;

drop table if exists `{{.DB_TABLE_PREFIX}}lti_routings` cascade;

drop table if exists `{{.DB_TABLE_PREFIX}}round_queue_pools` cascade;

drop table if exists `{{.DB_TABLE_PREFIX}}user_tokens` cascade;

drop table if exists `{{.DB_TABLE_PREFIX}}users` cascade;
