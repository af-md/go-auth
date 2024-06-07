-- MySQL dump 10.13  Distrib 8.0.34, for macos13 (arm64)
--
-- Host: 127.0.0.1    Database: account
-- ------------------------------------------------------
-- Server version	8.0.37

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `matches`
--

DROP TABLE IF EXISTS `matches`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `matches` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `UserId1` int NOT NULL,
  `UserId2` int NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `matches`
--

LOCK TABLES `matches` WRITE;
/*!40000 ALTER TABLE `matches` DISABLE KEYS */;
INSERT INTO `matches` VALUES (1,21,2),(2,23,4),(3,0,0);
/*!40000 ALTER TABLE `matches` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `swipe`
--

DROP TABLE IF EXISTS `swipe`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `swipe` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `UserId` int NOT NULL,
  `SwipedUserId` int NOT NULL,
  `Preference` varchar(45) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `swipe`
--

LOCK TABLES `swipe` WRITE;
/*!40000 ALTER TABLE `swipe` DISABLE KEYS */;
INSERT INTO `swipe` VALUES (1,17,2313,'YES'),(2,21,2,'YES'),(3,2,21,'YES'),(4,21,4,'YES'),(5,23,4,'YES'),(6,4,23,'YES');
/*!40000 ALTER TABLE `swipe` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(45) NOT NULL,
  `Email` varchar(45) NOT NULL,
  `Password` varchar(200) NOT NULL,
  `Gender` varchar(45) NOT NULL,
  `Age` int NOT NULL,
  `Location` int NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `Email_UNIQUE` (`Email`),
  UNIQUE KEY `Password_UNIQUE` (`Password`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Bob Davis','bob.davis@example.com','$2a$10$OXibkcD7R4pAeGTVZw9Hd.MIoC1wWROlNI/vXv3eCMQHYOkRhAnuu','Male',72,0),(2,'Omar Farooq','omar.farooq@hotmail.com','$2a$10$GA5MSnywJFA.4rY53dfyMe/YzkDA5l/Sweby43C9zxx6sVgPdaiDS','Male',66,0),(3,'Mahmoud Al-Mansour','mahmoud.al-mansour@hotmail.com','$2a$10$gD/Vo7UkMloWxJBeC4AYNe9tAdhdCW0dLCcmQCc4kafcWIB6A2KWO','Male',25,0),(4,'Aisha Syed','aisha.syed@yahoo.com','$2a$10$29etxlXnPWH8H0nf.d7Gj.oW3N6XQScXtceA35AAYqDCAfn9TYQ.S','Male',22,0),(5,'Hassan Siddiqi','hassan.siddiqi@example.com','$2a$10$q7SlCDiN1DCsi.TdsmWlpe8SZiRGPrtRvT3GHxpHoiis9D7y7Dc.S','Femal',18,0),(6,'Samira Syed','samira.syed@yahoo.com','$2a$10$4k27Taznapa/3c1/dOn77.sEJB/c82SkioPWMeeDiOdITXWnZ.2FC','Male',26,0),(7,'Muhammad Patel','muhammad.patel@outlook.com','$2a$10$PWIGY9jytvLgrB51zGNgfOuT0eYfnr.a4HrDVRuob0MWzTCqfUI6m','Femal',23,0),(8,'Hassan Iqbal','hassan.iqbal@example.com','$2a$10$cIYyi7MagHjQ.3flEOCN6.Isl/mtXUqmHen8V0t9E2.0C7urRuuKK','Male',27,0),(9,'Ahmad Iqbal','ahmad.iqbal@outlook.com','$2a$10$RRpPlqS8tRhHRkBYVMRepe1ntyQJp3KZw9tHplAn7McTupFzmLzyq','Femal',39,0),(10,'Zainab Hassan','zainab.hassan@outlook.com','$2a$10$Rba1HWC7zw8RusRDRz1oXOn91QpDiHP/m/5vZ90uTzzw79uFuGnpO','Female',38,0),(11,'Ibrahim Khan','ibrahim.khan@hotmail.com','$2a$10$akeIpz99ApeYD9lJtrund.jH8uqEM7AJGj.QVsONwkqn2dLLoLGae','Male',41,0),(12,'Omar Bakri','omar.bakri@yahoo.com','$2a$10$cm2BamY2Xg/XM7odjK2VSex78M/sA45q08PGJVGlnweAenEZ.1tSS','Female',23,0),(13,'Fatima Al-Hakim','fatima.al-hakim@outlook.com','$2a$10$o5TXpTmwaswF/Ows/YS4rO095.fLwI96kcg2sJ5gD7OAoLTlGOVXC','Female',46,0),(14,'Ali Hassan','ali.hassan@outlook.com','$2a$10$gHKVOsJ91xSDBsmdoAKY9eI2oRNuOJNeOJMtiMyB2BhYnmozVwrOG','Female',41,0),(15,'Hassan Al-Mansour','hassan.al-mansour@example.com','$2a$10$ZsaSHF8C9dQdSeJ8dDYa0eIcS7rldpPaztwwIIrqkVk3jIkpKCYdi','Male',38,0),(16,'Muhammad Bakri','muhammad.bakri@hotmail.com','$2a$10$q/TnT7AZ2vofxoiofLDjNO.c1JQksa4q5o9O4DVfnXHpTo0pR5.CC','Female',21,0),(17,'Ibrahim Bakri','ibrahim.bakri@gmail.com','$2a$10$dbbSL5P.8xbs8C/xC2kOO.kSAyigFm3zbLIsGZ0NTO3Z.ZWeqGuyi','Female',18,0),(18,'Hassan Al-Mansour','hassan.al-mansour@yahoo.com','$2a$10$AJcQ5cO4YAsKcfr4fZvH5Op7XWsEl6i0WJPtNo20YwpKRxCKK7SQO','Male',40,0),(19,'Maryam Al-Hakim','maryam.al-hakim@outlook.com','$2a$10$jfB/0Gb9v44bM.HAJWW7NujOeEyR43ZenbhTsiRNFZTj3hXrssJrS','Male',44,0),(20,'Fatima Hassan','fatima.hassan@yahoo.com','$2a$10$LSMovwkylyOU4cRZeQzjfu99sFEkPeuv60uHgepQ8I0PgWK3jg5ty','Female',24,0),(21,'Hassan Al-Hakim','hassan.al-hakim@hotmail.com','$2a$10$ZRImsMQYLqhUlGZ0gHBKcOKDEqW6df.zY1YFcxfJP2VNOTdnKVeYa','Female',43,0),(22,'Ahmad Malik','ahmad.malik@outlook.com','$2a$10$0ZrSoqu2pLz25etttFnL4eLXdvRugL5YjspOtMKxAbuAUohG50FIe','Male',20,0),(23,'Rania Saleh','rania.saleh@example.com','$2a$10$dXslnypJJHk8Y7yGP.48iO0dM6SSu8uRj4tKjRDd9OhycGrHgEynO','Female',47,57),(24,'Ali Al-Mansour','ali.al-mansour@gmail.com','$2a$10$7kfL1N0AuCq.4X5tKS.a/uiPDH6NTrpoKbgKHbuxmnYO7OVY68ikq','Female',21,45),(25,'Omar Yilmaz','omar.yilmaz@yahoo.com','$2a$10$BpVlzQ7iABTPwGvIP4vDSOaSx7Dtj5hA.jEd.Bn/uDcvojJs.3Cku','Male',21,48);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'account'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-06-07 23:44:14
