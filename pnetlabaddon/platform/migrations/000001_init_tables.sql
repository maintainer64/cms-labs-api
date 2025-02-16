-- +migrate Up
CREATE TABLE IF NOT EXISTS `wiresharks`
(
  `ws_id`
  bigint
  AUTO_INCREMENT,
  `ws_tenant`
  bigint,
  `ws_lab`
  varchar
(
  200
),`ws_node` bigint,`ws_if` bigint,`ws_net` bigint,`ws_node_name` varchar
(
  150
),`ws_if_name` varchar
(
  150
),`ws_dc_name` varchar
(
  150
),`ws_port` bigint,`ws_ip` varchar
(
  150
), PRIMARY KEY
(
  `ws_id`
));
CREATE TABLE IF NOT EXISTS `users`
(
  `pod`
  bigint
  AUTO_INCREMENT,
  `username`
  text,
  `cookie`
  text,
  `email`
  varchar
(
  150
),`expiration` bigint DEFAULT -1,`name` text,`password` text,`session` bigint,`ip` text,`role` text,`folder` text,`lab_session` bigint,`html5` boolean,`license` text,`online_time` bigint,`note` text,`offline` bigint,`active_time` bigint,`expired_time` bigint,`user_status` bigint DEFAULT 1,`user_workspace` text,`max_node` bigint,`max_node_lab` bigint, PRIMARY KEY
(
  `pod`
), CONSTRAINT `uni_users_email` UNIQUE
(
  `email`
));
CREATE TABLE IF NOT EXISTS `user_roles`
(
  `user_role_id`
  bigint
  AUTO_INCREMENT,
  `user_role_name`
  varchar
(
  150
),`user_role_workspace` text,`user_role_note` text,`user_role_ram` double,`user_role_cpu` double,`user_role_hdd` double, PRIMARY KEY
(
  `user_role_id`
));
CREATE TABLE IF NOT EXISTS `user_permission`
(
  `user_per_id`
  bigint
  AUTO_INCREMENT,
  `user_per_role`
  bigint,
  `user_per_name`
  varchar
(
  150
), PRIMARY KEY
(
  `user_per_id`
));
CREATE TABLE IF NOT EXISTS `process_device`
(
  `process_device_id`
  varchar
(
  150
),`process_device_dtotal` bigint,`process_device_dnow` bigint,`process_device_utotal` bigint,`process_device_unow` bigint,`process_device_log` text, PRIMARY KEY
(
  `process_device_id`
));
CREATE TABLE IF NOT EXISTS `process`
(
  `process_id`
  varchar
(
  200
),`process_dtotal` bigint,`process_dnow` bigint,`process_utotal` bigint,`process_unow` bigint,`process_finish` bigint, PRIMARY KEY
(
  `process_id`
));
CREATE TABLE IF NOT EXISTS `node_sessions`
(
  `node_session_id`
  bigint
  AUTO_INCREMENT,
  `node_session_nid`
  bigint,
  `node_session_lab`
  bigint,
  `node_session_port`
  bigint,
  `node_session_type`
  varchar
(
  150
),`node_session_workspace` text,`node_session_ram` double,`node_session_cpu` double,`node_session_hdd` double,`node_session_running` bigint,`node_session_pod` bigint,`node_session_iol` bigint, PRIMARY KEY
(
  `node_session_id`
));
CREATE TABLE IF NOT EXISTS `lab_sessions`
(
  `lab_session_id`
  bigint
  AUTO_INCREMENT,
  `lab_session_lid`
  varchar
(
  150
),`lab_session_pod` bigint,`lab_session_joined` text,`lab_session_path` text,`lab_session_running` bigint, PRIMARY KEY
(
  `lab_session_id`
));
CREATE TABLE IF NOT EXISTS `if_sessions`
(
  `if_session_id`
  bigint
  AUTO_INCREMENT,
  `if_session_lab`
  bigint,
  `if_session_node`
  bigint,
  `if_session_ifid`
  bigint,
  `if_session_type`
  varchar
(
  150
),`if_session_quality` text,`if_session_suspend` bigint, PRIMARY KEY
(
  `if_session_id`
));
CREATE TABLE IF NOT EXISTS `html5`
(
  `username`
  text,
  `pod`
  bigint,
  `token`
  text
);


-- +migrate Down
