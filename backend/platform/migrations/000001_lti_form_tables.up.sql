CREATE TABLE `lti_forms`
(
  `id`         bigint unsigned AUTO_INCREMENT,
  `created_at` datetime(3) default now(3) not null,
  `updated_at` datetime(3) default now(3) not null on update now(3),
  `name`       varchar(255)               not null,
  `version`    varchar(255)               not null,
  `lti_key`    varchar(255)               not null,
  `lti_secret` varchar(255)               not null,
  PRIMARY KEY (`id`),
  INDEX `idx_lti_forms_lti_key` (`lti_key`)
);
