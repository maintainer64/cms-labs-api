-- +migrate Up
alter table `{{.DB_TABLE_PREFIX}}lti_attempts`
  add attempt_id varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- +migrate Up
alter table `{{.DB_TABLE_PREFIX}}lti_attempts`
  add index `idx_{{.DB_TABLE_PREFIX}}lti_attempts_attempt_id` (`attempt_id`);

-- +migrate Down
alter table `{{.DB_TABLE_PREFIX}}lti_attempts` drop column attempt_id;

