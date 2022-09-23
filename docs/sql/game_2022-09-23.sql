# ************************************************************
# Sequel Ace SQL dump
# Version 20035
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# Host: 127.0.0.1 (MySQL 5.5.5-10.8.3-MariaDB)
# Database: game
# Generation Time: 2022-09-23 08:05:23 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table account
# ------------------------------------------------------------

DROP TABLE IF EXISTS `account`;

CREATE TABLE `account` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(50) DEFAULT NULL,
  `username` varchar(50) DEFAULT NULL,
  `login_type` varchar(11) DEFAULT NULL,
  `country` varchar(50) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `user_picture` varchar(200) DEFAULT NULL,
  `team_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `account` WRITE;
/*!40000 ALTER TABLE `account` DISABLE KEYS */;

INSERT INTO `account` (`id`, `user_id`, `username`, `login_type`, `country`, `email`, `user_picture`, `team_id`)
VALUES
	(1,'107415537797726047891','chun jay','google','Korea','chc3484mum@gmail.com','https://lh3.googleusercontent.com/a/AATXAJyYCyN5DRf35JWY7yfz-BybcSSJ5arHmgAr1b63=s96-c',1);

/*!40000 ALTER TABLE `account` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table player
# ------------------------------------------------------------

DROP TABLE IF EXISTS `player`;

CREATE TABLE `player` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `player_name` varchar(50) DEFAULT NULL,
  `squad_number` varchar(50) DEFAULT NULL,
  `team_name` varchar(50) DEFAULT NULL,
  `team_id` int(11) unsigned NOT NULL,
  `nation` varchar(11) DEFAULT NULL,
  `height` int(11) unsigned DEFAULT NULL,
  `weight` int(11) unsigned DEFAULT NULL,
  `age` int(11) unsigned DEFAULT NULL,
  `foot` varchar(11) DEFAULT NULL,
  `position` varchar(11) DEFAULT NULL,
  `ball_control` int(11) unsigned DEFAULT NULL,
  `dribbling` int(11) unsigned DEFAULT NULL,
  `low_pass` int(11) unsigned DEFAULT NULL,
  `lofted_pass` int(11) unsigned DEFAULT NULL,
  `speed` int(11) unsigned DEFAULT NULL,
  `jump` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `player` WRITE;
/*!40000 ALTER TABLE `player` DISABLE KEYS */;

INSERT INTO `player` (`id`, `player_name`, `squad_number`, `team_name`, `team_id`, `nation`, `height`, `weight`, `age`, `foot`, `position`, `ball_control`, `dribbling`, `low_pass`, `lofted_pass`, `speed`, `jump`)
VALUES
	(8574,'H. LLORIS','1','Tottenham',1,'FRANCE',188,78,34,'Left foot','GK',61,44,62,60,63,84),
	(43063,'SON HEUNG-MIN','7','Tottenham',1,'KOREA',183,79,28,'Right foot','LMF',86,86,78,80,93,72),
	(43166,'M. DOHERTY','2','Tottenham',1,'IRELAND',186,89,28,'Right foot','RB',74,73,75,80,82,81),
	(47287,'H. KANE','10','Tottenham',1,'ENGLAND',188,76,27,'Right foot','CF',83,78,82,84,78,80),
	(112523,'REGUILÓN','3','Tottenham',1,'SPAIN',178,68,24,'Left foot','LB',77,80,78,80,87,68);

/*!40000 ALTER TABLE `player` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table team
# ------------------------------------------------------------

DROP TABLE IF EXISTS `team`;

CREATE TABLE `team` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `team_name` varchar(50) DEFAULT NULL,
  `manager_name` varchar(50) DEFAULT NULL,
  `formation` varchar(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `team` WRITE;
/*!40000 ALTER TABLE `team` DISABLE KEYS */;

INSERT INTO `team` (`id`, `team_name`, `manager_name`, `formation`)
VALUES
	(1,'Tottenham','Antonio Conte','3:4:3'),
	(2,'Manchester City','Josep Guardiola','4:3:3'),
	(3,'Manchester United','Erik ten Hag','4:2:3:1');

/*!40000 ALTER TABLE `team` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
