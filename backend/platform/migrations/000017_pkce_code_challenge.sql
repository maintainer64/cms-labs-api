-- +migrate Up
ALTER TABLE `{{.DB_TABLE_PREFIX}}token_attempts`
ADD COLUMN `code_challenge` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
ADD COLUMN `code_challenge_method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- +migrate Down
ALTER TABLE `{{.DB_TABLE_PREFIX}}token_attempts`
DROP COLUMN IF EXISTS `code_challenge`,
DROP COLUMN IF EXISTS `code_challenge_method`;
