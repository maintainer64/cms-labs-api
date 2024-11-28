CREATE TABLE `lti_access_tokens`
(
  `id`         bigint unsigned AUTO_INCREMENT,
  `index`      longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `payload`    longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  PRIMARY KEY (`id`),
  INDEX `idx_lti_access_tokens_index` (`index`(768)),
  INDEX `lti_access_tokens_created_at` (`created_at`)
);

CREATE TABLE `lti_forms`
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
  INDEX `idx_lti_forms_lti_client_id` (`lti_client_id`),
  INDEX `idx_lti_forms_lti_deployment_id` (`lti_deployment_id`),
  INDEX `idx_lti_forms_lti_base_uri` (`base_uri`),
  INDEX `idx_lti_forms_created_at` (`created_at`)
);

CREATE TABLE `lti_launch_data`
(
  `id`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `lti_form_id` bigint unsigned,
  `created_at`  datetime(3),
  `updated_at`  datetime(3),
  `launch_data` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`),
  INDEX `idx_lti_launch_data_lti_form_id` (`lti_form_id`),
  INDEX `idx_lti_launch_data_created_at` (`created_at`),
  INDEX `idx_lti_launch_data_updated_at` (`updated_at`)
);

CREATE TABLE `lti_nonce_tokens`
(
  `id`              bigint unsigned AUTO_INCREMENT,
  `created_at`      datetime(3),
  `updated_at`      datetime(3),
  `nonce`           varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `target_link_uri` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`),
  INDEX `idx_lti_nonce_tokens_nonce` (`nonce`),
  INDEX `idx_lti_nonce_tokens_created_at` (`created_at`)
);

CREATE TABLE `pnet_servers`
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
  UNIQUE INDEX `idx_pnet_servers_token` (`token`),
  INDEX `idx_pnet_servers_created_at` (`created_at`),
  INDEX `idx_pnet_servers_updated_at` (`updated_at`)
);

CREATE TABLE `users`
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
  UNIQUE INDEX `idx_users_email` (`email`),
  INDEX `idx_users_last_launch_id` (`last_launch_id`),
  INDEX `users_created_at` (`created_at`)
);

CREATE TABLE `user_tokens`
(
  `user_id`       bigint unsigned,
  `refresh_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `hash_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `created_at`    datetime(3),
  `updated_at`    datetime(3),
  PRIMARY KEY (`user_id`),
  FOREIGN KEY (`user_id`)
    REFERENCES users (`id`)
    ON DELETE CASCADE,
  INDEX `idx_user_tokens_refresh_token` (`refresh_token`),
  INDEX `user_tokens_created_at` (`created_at`)
);

CREATE TABLE `round_queue_pools`
(
  `id`             bigint unsigned AUTO_INCREMENT,
  `created_at`     datetime(3),
  `updated_at`     datetime(3),
  `type`           varchar(255),
  `connected_at`   datetime(3),
  `last_used`      boolean,
  `is_active`      boolean,
  `pnet_server_id` bigint null,
  PRIMARY KEY (`id`),
  INDEX `idx_round_queue_pools_connected_at` (`connected_at`),
  INDEX `idx_round_queue_pools_type` (`type`),
  INDEX `idx_round_queue_pools_is_active` (`is_active`),
  INDEX `idx_round_queue_pools_pnet_server_id` (`pnet_server_id`),
  INDEX `uround_queue_pools_created_at` (`created_at`)
);

CREATE TABLE `lti_routings`
(
  `id`              bigint unsigned AUTO_INCREMENT,
  `created_at`      datetime(3),
  `updated_at`      datetime(3),
  `name`            varchar(255),
  `lti_title`       varchar(255),
  `lti_description` varchar(255),
  `lti_task_id`     varchar(255),
  `lti_params_task` varchar(255),
  `collaboration`   bigint,
  `pnet_labs_path`  varchar(255),
  `pnet_test_path`  varchar(255),
  PRIMARY KEY (`id`),
  INDEX `lti_routings_created_at` (`created_at`),
  INDEX `lti_routings_lti_title` (`lti_title`),
  INDEX `lti_routings_lti_task_id` (`lti_task_id`),
  INDEX `lti_routings_lti_params_task` (`lti_params_task`)
);
