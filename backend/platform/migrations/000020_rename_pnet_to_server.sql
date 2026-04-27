-- Rename pnet_servers table to servers and rename pnet_ prefix columns to server_/labs_/test_
-- 1. Rename table pnet_servers to servers
RENAME TABLE `{{.DB_TABLE_PREFIX}}pnet_servers` TO `{{.DB_TABLE_PREFIX}}servers`;

-- 2. Rename columns in lti_routings (pnet_ -> server_/labs_/test_)
ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_routings`
    CHANGE COLUMN `pnet_server_id` `server_id` bigint unsigned,
    CHANGE COLUMN `pnet_labs_type` `labs_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    CHANGE COLUMN `pnet_labs_path` `labs_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    CHANGE COLUMN `pnet_test_path` `test_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 3. Rename column in lti_attempts (pnet_server_id -> server_id)
ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_attempts`
    CHANGE COLUMN `pnet_server_id` `server_id` bigint unsigned;

-- 4. Rename column in round_queue_pools (pnet_server_id -> server_id)
ALTER TABLE `{{.DB_TABLE_PREFIX}}round_queue_pools`
    CHANGE COLUMN `pnet_server_id` `server_id` bigint unsigned;

-- 5. Update indexes (drop old, create new)
ALTER TABLE `{{.DB_TABLE_PREFIX}}round_queue_pools`
    DROP INDEX IF EXISTS `idx_{{.DB_TABLE_PREFIX}}round_queue_pools_pnet_server_id`,
    ADD INDEX `idx_{{.DB_TABLE_PREFIX}}round_queue_pools_server_id` (`server_id`);

-- 6. Update foreign keys
ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_routings`
    DROP FOREIGN KEY IF EXISTS `lti_routings_ibfk_1`,
    ADD CONSTRAINT `lti_routings_ibfk_server` FOREIGN KEY (`server_id`) REFERENCES `{{.DB_TABLE_PREFIX}}servers` (`id`) ON DELETE CASCADE;

ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_attempts`
    DROP FOREIGN KEY IF EXISTS `lti_attempts_ibfk_2`,
    ADD CONSTRAINT `lti_attempts_ibfk_server` FOREIGN KEY (`server_id`) REFERENCES `{{.DB_TABLE_PREFIX}}servers` (`id`) ON DELETE SET NULL;

ALTER TABLE `{{.DB_TABLE_PREFIX}}round_queue_pools`
    DROP FOREIGN KEY IF EXISTS `round_queue_pools_ibfk_1`,
    ADD CONSTRAINT `round_queue_pools_ibfk_server` FOREIGN KEY (`server_id`) REFERENCES `{{.DB_TABLE_PREFIX}}servers` (`id`) ON DELETE CASCADE;

-- 7. Update role_relations foreign key
ALTER TABLE `{{.DB_TABLE_PREFIX}}role_relations`
    DROP FOREIGN KEY IF EXISTS `{{.DB_TABLE_PREFIX}}_role_relations_ibfk_server`,
    ADD CONSTRAINT `{{.DB_TABLE_PREFIX}}_role_relations_ibfk_server` FOREIGN KEY (`server_id`) REFERENCES `{{.DB_TABLE_PREFIX}}servers` (`id`) ON DELETE CASCADE;