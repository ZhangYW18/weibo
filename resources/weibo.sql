-- MySQL dump 10.13  Distrib 8.0.15, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: weibo
-- ------------------------------------------------------
-- Server version	5.7.25-0ubuntu0.18.10.2

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `comment` (
  `commentid` int(11) NOT NULL AUTO_INCREMENT,
  `weiboid` int(11) NOT NULL,
  `userid` int(11) NOT NULL,
  `username` varchar(45) NOT NULL,
  `created_at` datetime NOT NULL,
  `comment_text` varchar(280) NOT NULL,
  PRIMARY KEY (`commentid`),
  UNIQUE KEY `commentID_UNIQUE` (`commentid`),
  KEY `fk_comment_2_idx` (`userid`,`username`),
  KEY `fk_comment_1_idx` (`weiboid`),
  CONSTRAINT `fk_comment_1` FOREIGN KEY (`weiboid`) REFERENCES `weibo` (`weiboid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_comment_2` FOREIGN KEY (`userid`, `username`) REFERENCES `user` (`userid`, `username`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` VALUES (1,1,0,'admin','2018-04-05 00:00:00','评论1'),(2,1,1,'WhatIf','2018-09-08 00:00:00','评论2'),(3,5,0,'admin','2019-05-06 20:52:42','Melodrama牛逼'),(4,6,0,'admin','2019-05-06 21:46:26','你在这里，活着');
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `follow`
--

DROP TABLE IF EXISTS `follow`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `follow` (
  `userid` int(11) NOT NULL,
  `followid` int(11) NOT NULL,
  PRIMARY KEY (`userid`,`followid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `follow`
--

LOCK TABLES `follow` WRITE;
/*!40000 ALTER TABLE `follow` DISABLE KEYS */;
INSERT INTO `follow` VALUES (0,1),(0,2),(1,0);
/*!40000 ALTER TABLE `follow` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `user` (
  `userid` int(11) NOT NULL,
  `username` varchar(16) NOT NULL,
  `password` varchar(16) NOT NULL,
  `sex` tinyint(1) DEFAULT NULL,
  `age` int(3) DEFAULT NULL,
  PRIMARY KEY (`userid`,`username`),
  UNIQUE KEY `userID_UNIQUE` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (0,'admin','123456',1,25),(1,'WhatIf','1',1,35),(2,'IfThenElse','1',0,12),(3,'if','1',0,17),(4,'if','1',1,34);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `weibo`
--

DROP TABLE IF EXISTS `weibo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `weibo` (
  `weiboid` int(11) NOT NULL AUTO_INCREMENT,
  `userid` int(11) NOT NULL,
  `username` varchar(16) NOT NULL,
  `created_at` datetime NOT NULL,
  `text` varchar(280) NOT NULL,
  `like` int(11) DEFAULT NULL,
  `comment_count` int(11) DEFAULT NULL,
  PRIMARY KEY (`weiboid`),
  UNIQUE KEY `weiboID_UNIQUE` (`weiboid`),
  KEY `fk_Weibo_1_idx` (`userid`,`username`),
  CONSTRAINT `fk_Weibo_1` FOREIGN KEY (`userid`, `username`) REFERENCES `user` (`userid`, `username`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `weibo`
--

LOCK TABLES `weibo` WRITE;
/*!40000 ALTER TABLE `weibo` DISABLE KEYS */;
INSERT INTO `weibo` VALUES (1,0,'admin','2018-10-05 12:00:00','Father of the Bride',0,2),(4,0,'admin','2019-05-06 19:33:06','oooo',0,0),(5,0,'admin','2019-05-06 19:34:44','Lorde牛逼',0,1),(6,1,'WhatIf','2019-05-05 19:00:00','我身在何处',0,1);
/*!40000 ALTER TABLE `weibo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'weibo'
--

--
-- Dumping routines for database 'weibo'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-05-08 21:38:57
