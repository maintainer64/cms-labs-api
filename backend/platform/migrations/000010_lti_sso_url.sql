-- +migrate Up
alter table `{{.DB_TABLE_PREFIX}}lti_forms`
  ADD COLUMN `sso_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT (NULL);


-- +migrate Down
alter table `{{.DB_TABLE_PREFIX}}lti_forms` drop column sso_url;

