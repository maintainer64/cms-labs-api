-- +migrate Up
alter table `{{.DB_TABLE_PREFIX}}lti_routings`
  add pnet_server_id bigint unsigned;

-- +migrate Down
alter table `{{.DB_TABLE_PREFIX}}lti_routings` drop column pnet_server_id;

