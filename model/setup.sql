CREATE DATABASE IF NOT EXISTS `go_api` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE `go_api`;

SET FOREIGN_KEY_CHECKS=0;

DROP TABLE IF EXISTS `tag`;
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '',
  `created_on` int(10) unsigned DEFAULT '0',
  `modified_on` int(10) unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;