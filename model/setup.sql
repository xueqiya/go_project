CREATE DATABASE IF NOT EXISTS `gin-blog` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE `gin-blog`;

SET FOREIGN_KEY_CHECKS=0;

DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '',
  `created_on` int(10) unsigned DEFAULT '0',
  `modified_on` int(10) unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;