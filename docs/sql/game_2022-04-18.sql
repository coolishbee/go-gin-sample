# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.5.5-10.7.3-MariaDB)
# Database: game
# Generation Time: 2022-04-17 15:17:16 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table player
# ------------------------------------------------------------

DROP TABLE IF EXISTS `player`;

CREATE TABLE `player` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `player_name` varchar(50) DEFAULT NULL,
  `squad_number` varchar(50) DEFAULT NULL,
  `team_name` varchar(11) DEFAULT NULL,
  `team_id` int(11) unsigned DEFAULT NULL,
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



# Dump of table team
# ------------------------------------------------------------

DROP TABLE IF EXISTS `team`;

CREATE TABLE `team` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `team_name` varchar(11) DEFAULT NULL,
  `manager_name` varchar(11) DEFAULT NULL,
  `formation` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
