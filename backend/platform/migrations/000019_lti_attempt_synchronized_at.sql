-- +migrate Up
ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_attempts`
ADD COLUMN `synchronized_at` datetime(3),
ADD INDEX `idx_{{.DB_TABLE_PREFIX}}lti_attempts_synchronized_at` (`synchronized_at`);

-- +migrate Up
ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_launch_data`
  ADD COLUMN `attempt_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
ADD INDEX `idx_{{.DB_TABLE_PREFIX}}lti_launch_data_attempt_id` (`attempt_id`);

-- +migrate Down
ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_launch_data`
DROP INDEX `idx_{{.DB_TABLE_PREFIX}}lti_launch_data_attempt_id`,
DROP COLUMN `attempt_id`;

-- +migrate Down
ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_attempts`
DROP INDEX `idx_{{.DB_TABLE_PREFIX}}lti_attempts_synchronized_at`,
DROP COLUMN `synchronized_at`;
