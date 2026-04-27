create
definer = root@`%` procedure generate_random_users(IN num_users int)
BEGIN
    DECLARE
i INT DEFAULT 0;
    WHILE
i < num_users
        DO
            INSERT INTO core_users (created_at, updated_at, deleted_at, email, name, user_role, group_name, lti_user_id,
                                    last_launch_id)
            VALUES (NOW(), NOW(), NULL, CONCAT('user', i, '@example.com'), CONCAT('User ', i),
                    IF(i % 2 = 0, 'admin', 'student'), NULL, NULL, '');
            SET
i = i + 1;
END WHILE;
END;

CREATE
DEFINER = root@`%` PROCEDURE generate_random_servers(IN num_servers INT)
BEGIN
    DECLARE
i INT DEFAULT 0;
    WHILE
i < num_servers
        DO
            INSERT INTO core_servers (created_at, updated_at, name, url, is_active, minutes_for_disconnect,
                                           max_count_users_limit, last_online_status, last_count_users, unit_rate,
                                           token)
            VALUES (NOW(), NOW(), CONCAT('Server ', i), CONCAT('http://server', i, '.example.com'), 1,
                    FLOOR(RAND() * 100),
                    FLOOR(RAND() * 1000), NOW(), FLOOR(RAND() * 100), FLOOR(RAND() * 10), MD5(RAND()));
            SET
i = i + 1;
END WHILE;
END;
create
definer = root@`%` procedure add_all_rooms(IN num_rooms INT)
BEGIN
    DECLARE
i INT DEFAULT 0;
    WHILE
i < num_rooms
        DO
INSERT INTO core_lti_rooms (created_at, updated_at, room_number)
VALUES (NOW(), NOW(), i);
SET
i = i + 1;
END WHILE;
END;
