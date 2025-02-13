-- +migrate Up
alter table `{{.DB_TABLE_PREFIX}}lti_routings`
  add is_default boolean;

-- +migrate Down
alter table `{{.DB_TABLE_PREFIX}}lti_routings` drop column is_default;
