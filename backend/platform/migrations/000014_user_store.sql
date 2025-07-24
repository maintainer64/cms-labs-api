-- +migrate Up
ALTER TABLE `{{.DB_TABLE_PREFIX}}users`
  ADD COLUMN store JSON NULL DEFAULT NULL;

-- +migrate Down
alter table `{{.DB_TABLE_PREFIX}}users` drop column store;

