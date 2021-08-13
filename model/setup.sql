CREATE
DATABASE IF NOT EXISTS `go_api` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE
`go_api`;

SET
FOREIGN_KEY_CHECKS=0;

DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag`
(
    `id`          BIGINT ( 10 ) NOT NULL AUTO_INCREMENT,
    `name`        VARCHAR(100) DEFAULT '',
    `created_on`  INT ( 10 ) DEFAULT '0',
    `modified_on` INT ( 10 ) DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`          BIGINT      NOT NULL AUTO_INCREMENT,
    `phone`       VARCHAR(11) NOT NULL DEFAULT '',
    `password`    VARCHAR(20) NOT NULL DEFAULT '',
    `nike_name`   VARCHAR(20)          DEFAULT '',
    `age`         VARCHAR(3)           DEFAULT '0',
    `created_on`  int                  DEFAULT '0',
    `modified_on` int                  DEFAULT '0',
    `status`      TINYINT              DEFAULT '1',
    PRIMARY KEY (`id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;