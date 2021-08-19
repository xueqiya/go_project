CREATE
DATABASE IF NOT EXISTS `go_project` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE
`go_project`;

SET
FOREIGN_KEY_CHECKS=0;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`          BIGINT      NOT NULL AUTO_INCREMENT,
    `phone`       VARCHAR(11) NOT NULL DEFAULT '',
    `password`    VARCHAR(20) NOT NULL DEFAULT '',
    `nike_name`   VARCHAR(20)          DEFAULT '',
    `age`         int                  DEFAULT '0',
    `created_on`  BIGINT               DEFAULT '0',
    `modified_on` BIGINT               DEFAULT '0',
    `status`      TINYINT              DEFAULT '1',
    PRIMARY KEY (`id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods`
(
    `id`            BIGINT NOT NULL AUTO_INCREMENT,
    `fk_user_goods` BIGINT       DEFAULT '0',
    `price`         int          DEFAULT '0',
    `keyword`       VARCHAR(20)  DEFAULT '',
    `content`       VARCHAR(200) DEFAULT '',
    `image`         VARCHAR(100) DEFAULT '',
    `location`      VARCHAR(100) DEFAULT '',
    `address`       VARCHAR(100) DEFAULT '',
    `created_on`    BIGINT       DEFAULT '0',
    `modified_on`   BIGINT       DEFAULT '0',
    `status`        TINYINT      DEFAULT '1',
    PRIMARY KEY (`id`),
    FOREIGN KEY (`fk_user_goods`) REFERENCES user (`id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;