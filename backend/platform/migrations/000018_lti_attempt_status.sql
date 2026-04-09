-- +migrate Up
ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_attempts`
ADD COLUMN `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'pending',
ADD COLUMN `result` json,
ADD INDEX `idx_{{.DB_TABLE_PREFIX}}lti_attempts_status` (`status`);

UPDATE `{{.DB_TABLE_PREFIX}}lti_attempts` SET status = 'completed';

ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_attempts`
DROP INDEX `idx_{{.DB_TABLE_PREFIX}}lti_attempts_expired_at`,
DROP COLUMN `expired_at`;

-- +migrate Down
ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_attempts`
ADD COLUMN `expired_at` datetime(3),
ADD INDEX `idx_{{.DB_TABLE_PREFIX}}lti_attempts_expired_at` (`expired_at`);

ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_attempts`
DROP INDEX `idx_{{.DB_TABLE_PREFIX}}lti_attempts_status`,
DROP COLUMN `status`,
DROP COLUMN `result`;
