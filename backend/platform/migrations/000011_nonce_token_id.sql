-- +migrate Up

alter table `{{.DB_TABLE_PREFIX}}token_attempts`
  add nonce varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;


-- +migrate Down
alter table `{{.DB_TABLE_PREFIX}}token_attempts` drop column nonce;

