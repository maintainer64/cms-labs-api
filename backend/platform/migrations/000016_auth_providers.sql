-- +migrate Up
RENAME
TABLE `{{.DB_TABLE_PREFIX}}lti_forms` TO `{{.DB_TABLE_PREFIX}}auth_providers`;

-- Rename indexes to reflect new table name
-- +migrate Up
ALTER TABLE `{{.DB_TABLE_PREFIX}}auth_providers`
  RENAME INDEX `idx_{{.DB_TABLE_PREFIX}}lti_forms_created_at` TO `idx_{{.DB_TABLE_PREFIX}}auth_providers_created_at`,
  RENAME INDEX `idx_{{.DB_TABLE_PREFIX}}lti_forms_lti_base_uri` TO `idx_{{.DB_TABLE_PREFIX}}auth_providers_lti_base_uri`,
  RENAME INDEX `idx_{{.DB_TABLE_PREFIX}}lti_forms_lti_client_id` TO `idx_{{.DB_TABLE_PREFIX}}auth_providers_lti_client_id`,
  RENAME INDEX `idx_{{.DB_TABLE_PREFIX}}lti_forms_lti_deployment_id` TO `idx_{{.DB_TABLE_PREFIX}}auth_providers_lti_deployment_id`;

-- +migrate Up
ALTER TABLE `{{.DB_TABLE_PREFIX}}auth_providers`
  ADD COLUMN `type` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'lti';

-- +migrate Up
ALTER TABLE `{{.DB_TABLE_PREFIX}}auth_providers`
  ADD INDEX `idx_{{.DB_TABLE_PREFIX}}auth_providers_type` (`type`);

-- +migrate Down
ALTER TABLE `{{.DB_TABLE_PREFIX}}auth_providers` DROP COLUMN `type`;

-- +migrate Down
DROP INDEX `idx_{{.DB_TABLE_PREFIX}}auth_providers_type` ON `{{.DB_TABLE_PREFIX}}auth_providers`;

-- Revert index names to original
-- +migrate Down
ALTER TABLE `{{.DB_TABLE_PREFIX}}auth_providers`
  RENAME INDEX `idx_{{.DB_TABLE_PREFIX}}auth_providers_created_at` TO `idx_{{.DB_TABLE_PREFIX}}lti_forms_created_at`,
  RENAME INDEX `idx_{{.DB_TABLE_PREFIX}}auth_providers_lti_base_uri` TO `idx_{{.DB_TABLE_PREFIX}}lti_forms_lti_base_uri`,
  RENAME INDEX `idx_{{.DB_TABLE_PREFIX}}auth_providers_lti_client_id` TO `idx_{{.DB_TABLE_PREFIX}}lti_forms_lti_client_id`,
  RENAME INDEX `idx_{{.DB_TABLE_PREFIX}}auth_providers_lti_deployment_id` TO `idx_{{.DB_TABLE_PREFIX}}lti_forms_lti_deployment_id`;

-- Rename table back to lti_forms
-- +migrate Down
RENAME
TABLE `{{.DB_TABLE_PREFIX}}auth_providers` TO `{{.DB_TABLE_PREFIX}}lti_forms`;
