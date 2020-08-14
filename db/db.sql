show global variables like '%lower_case%';
SHOW COLLATION like 'utf8%';

show variables like '%time_zone%';
# 查看当前会话时区
SELECT @@session.time_zone;
# # 设置当前会话时区
# SET time_zone = 'Europe/Helsinki';
SET time_zone = "+8:00";
# # 数据库全局时区设置
SELECT @@global.time_zone;
# # 设置全局时区
SET GLOBAL time_zone = '+8:00';
# SET GLOBAL time_zone = 'Europe/Helsinki';
SELECT NOW();
SELECT UTC_TIMESTAMP;
SELECT LOCALTIMESTAMP;
SELECT CURRENT_TIMESTAMP;

# DROP DATABASE `ok-short`;
CREATE DATABASE `ok-short` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
USE `ok-short`;

# -------------------短链数据--------------------
DROP TABLE IF EXISTS ok_link;
CREATE TABLE ok_link
(
    `id`                    INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `origin_url`            VARCHAR(255)     Not NULL,
    `short_code`            VARCHAR(10)
                                CHARACTER SET utf8mb4
                                    COLLATE utf8mb4_bin # 大小写敏感
                                             NOT NULL,
    `created_by`            INT(10) UNSIGNED NOT NULL,
    `created_at`            TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `expiration_in_minutes` INT(10) UNSIGNED NOT NULL DEFAULT 30,
    PRIMARY KEY (`id`)
) ENGINE = INNODB
  AUTO_INCREMENT = 1
  COLLATE = utf8mb4_unicode_ci
  DEFAULT CHARSET = utf8mb4;
ALTER TABLE ok_link
    ADD UNIQUE (`short_code`);

# -------------------短链访问记录--------------
DROP TABLE IF EXISTS `ok_link_visited_log`;
CREATE TABLE `ok_link_visited_log`
(
    `id`          INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `remote_addr` VARCHAR(32)      Not NULL,
    `short_code`  VARCHAR(10)
                      CHARACTER SET utf8mb4
                          COLLATE utf8mb4_bin # 大小写敏感
                                   NOT NULL,
    `ua`          VARCHAR(500)     Not NULL,
    `cookie`      VARCHAR(500)     Not NULL,
    `visitor_id`  INT(10) UNSIGNED NOT NULL,
    `visited_at`  TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    PRIMARY KEY (`id`)
) ENGINE = INNODB
  AUTO_INCREMENT = 1
  COLLATE = utf8mb4_unicode_ci
  DEFAULT CHARSET = utf8mb4;

# -------------------短链作者(服务用户)--------------
DROP TABLE IF EXISTS `ok_link_author`;
CREATE TABLE `ok_link_author`
(
    `id`         INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `sid`        VARCHAR(10)
                     CHARACTER SET utf8mb4
                         COLLATE utf8mb4_bin # 大小写敏感
                                  NOT NULL,
    `name`       VARCHAR(20)      Not NULL,
    `password`   VARCHAR(64)      Not NULL,
    `email`      VARCHAR(30)      Not NULL,
    `avatar_url` VARCHAR(100)     Not NULL,
    `created_at` TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `updated_at` TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    PRIMARY KEY (`id`)
) ENGINE = INNODB
  AUTO_INCREMENT = 1
  COLLATE = utf8mb4_unicode_ci
  DEFAULT CHARSET = utf8mb4;

# -------------------后台管理人员--------------
DROP TABLE IF EXISTS `ok_link_admin_user`;
CREATE TABLE `ok_link_admin_user`
(
    `id`         INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`       VARCHAR(20)      Not NULL,
    `password`   VARCHAR(64)      Not NULL,
    `email`      VARCHAR(30),
    `avatar_url` VARCHAR(100),
    `created_at` TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `updated_at` TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `deleted_at` TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = INNODB
  AUTO_INCREMENT = 1
  COLLATE = utf8mb4_unicode_ci
  DEFAULT CHARSET = utf8mb4;

# -------------------------------------
DROP TABLE IF EXISTS `visitors`;
CREATE TABLE `visitors`
(
    `id`       INT(10) UNSIGNED           NOT NULL AUTO_INCREMENT,
    `ip`       INT(16) UNSIGNED           NOT NULL,
    `browser`  VARCHAR(500)               NOT NULL,
    `version`  VARCHAR(500)               NOT NULL,
    `platform` ENUM ('w','l','m','a','i') NOT NULL,
    `date`     TIMESTAMP                  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
#     ,UNIQUE KEY `person` (`ip`, `date`)
) ENGINE = INNODB
  AUTO_INCREMENT = 1
  COLLATE = utf8mb4_unicode_ci
  DEFAULT CHARSET = utf8mb4;

# -------------------------
SELECT *
FROM ok_link
ORDER BY id DESC;
# WHERE short_code = 'Yt';

SELECT *
FROM ok_link_visited_log
ORDER BY id desc;


SELECT * FROM ok_link_admin_user ORDER BY id DESC ;

# TODO 云服务器MySQL时区问题