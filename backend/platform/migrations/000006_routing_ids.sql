-- +migrate Up
DROP INDEX `idx_{{.DB_TABLE_PREFIX}}lti_routings_lti_task_id` ON `{{.DB_TABLE_PREFIX}}lti_routings`;

-- +migrate Up
alter table `{{.DB_TABLE_PREFIX}}lti_routings`
  add lti_sub_id varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- +migrate Up
alter table `{{.DB_TABLE_PREFIX}}lti_routings`
  add lti_course_id varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;


-- +migrate Down
alter table `{{.DB_TABLE_PREFIX}}lti_routings` drop column lti_course_id;

-- +migrate Down
alter table `{{.DB_TABLE_PREFIX}}lti_routings` drop column lti_sub_id;

-- +migrate Down
ALTER TABLE `{{.DB_TABLE_PREFIX}}lti_routings`
  ADD INDEX `idx_{{.DB_TABLE_PREFIX}}lti_routings_lti_task_id` (`lti_task_id`);
