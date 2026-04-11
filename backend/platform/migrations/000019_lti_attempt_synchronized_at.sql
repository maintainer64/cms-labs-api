-- +migrate Up
ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_attempts`
ADD COLUMN `synchronized_at` datetime(3),
ADD INDEX `idx_{{.DB_TABLE_PREFIX}}lti_attempts_synchronized_at` (`synchronized_at`);

-- +migrate Down
ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_attempts`
DROP COLUMN IF EXISTS `synchronized_at`,
DROP INDEX IF EXISTS `idx_{{.DB_TABLE_PREFIX}}lti_attempts_synchronized_at`;