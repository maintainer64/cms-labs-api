-- +migrate Up
create table `{{.DB_TABLE_PREFIX}}lti_rooms`
(
  `id`          bigint unsigned AUTO_INCREMENT,
  `created_at`  datetime(3),
  `updated_at`  datetime(3),
  `room_number` bigint,
  PRIMARY KEY (`id`),
  INDEX         `idx_{{.DB_TABLE_PREFIX}}created_at` (`created_at`)
);


alter table `{{.DB_TABLE_PREFIX}}lti_attempts` drop column room_number;

alter table `{{.DB_TABLE_PREFIX}}lti_attempts`
  add column room_id bigint unsigned DEFAULT(NULL),
  add foreign key (`room_id`) references `{{.DB_TABLE_PREFIX}}lti_rooms` (`id`) on
delete
cascade;


-- +migrate Down
alter table `{{.DB_TABLE_PREFIX}}lti_attempts` drop foreign key `{{.DB_TABLE_PREFIX}}lti_attempts_ibfk_4`;

alter table `{{.DB_TABLE_PREFIX}}lti_attempts` drop column room_id;

alter table `{{.DB_TABLE_PREFIX}}lti_attempts`
  add column room_number bigint DEFAULT (NULL);

drop table if exists `{{.DB_TABLE_PREFIX}}lti_rooms` cascade;

