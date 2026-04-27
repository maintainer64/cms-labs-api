-- +migrate Up
-- Rename pnet_servers table to servers and rename pnet_ prefix columns to server_/labs_/test_

SET foreign_key_checks = 0;

-- Rename table pnet_servers to servers
RENAME TABLE `{{.DB_TABLE_PREFIX}}pnet_servers` TO `{{.DB_TABLE_PREFIX}}servers`;

-- Rename columns in lti_routings (pnet_ -> server_/labs_/test_)
ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_routings`
    CHANGE COLUMN `pnet_server_id` `server_id` bigint unsigned,
    CHANGE COLUMN `pnet_labs_type` `labs_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    CHANGE COLUMN `pnet_labs_path` `labs_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    CHANGE COLUMN `pnet_test_path` `test_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Rename column in lti_attempts (pnet_server_id -> server_id)
ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_attempts`
    CHANGE COLUMN `pnet_server_id` `server_id` bigint unsigned;

-- Rename column in round_queue_pools (pnet_server_id -> server_id)
ALTER TABLE `{{.DB_TABLE_PREFIX}}round_queue_pools`
    CHANGE COLUMN `pnet_server_id` `server_id` bigint unsigned;

-- Rename column in role_relations (server_id already has correct name)
-- No need to rename, just update FK later

-- Update index name
ALTER TABLE `{{.DB_TABLE_PREFIX}}round_queue_pools`
    DROP INDEX `idx_{{.DB_TABLE_PREFIX}}round_queue_pools_pnet_server_id`,
    ADD INDEX `idx_{{.DB_TABLE_PREFIX}}round_queue_pools_server_id` (`server_id`);

SET foreign_key_checks = 1;

-- +migrate Down
SET foreign_key_checks = 0;

-- Rename columns back in round_queue_pools
ALTER TABLE `{{.DB_TABLE_PREFIX}}round_queue_pools`
    DROP INDEX `idx_{{.DB_TABLE_PREFIX}}round_queue_pools_server_id`,
    ADD INDEX `idx_{{.DB_TABLE_PREFIX}}round_queue_pools_pnet_server_id` (`pnet_server_id`);

ALTER TABLE `{{.DB_TABLE_PREFIX}}round_queue_pools`
    CHANGE COLUMN `server_id` `pnet_server_id` bigint unsigned;

-- Rename columns back in lti_attempts
ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_attempts`
    CHANGE COLUMN `server_id` `pnet_server_id` bigint unsigned;

-- Rename columns back in lti_routings
ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_routings`
    CHANGE COLUMN `server_id` `pnet_server_id` bigint unsigned,
    CHANGE COLUMN `labs_type` `pnet_labs_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    CHANGE COLUMN `labs_path` `pnet_labs_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    CHANGE COLUMN `test_path` `pnet_test_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Rename table back
RENAME TABLE `{{.DB_TABLE_PREFIX}}servers` TO `{{.DB_TABLE_PREFIX}}pnet_servers`;

SET foreign_key_checks = 1;