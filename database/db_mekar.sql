CREATE DATABASE  IF NOT EXISTS `db_mekar` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `db_mekar`;
-- MySQL dump 10.13  Distrib 8.0.20, for Win64 (x86_64)
--
-- Host: localhost    Database: db_mekar
-- ------------------------------------------------------
-- Server version	8.0.20

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
-- Table structure for table `admin`
--

DROP TABLE IF EXISTS `admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin` (
  `username` varchar(50) NOT NULL,
  `password` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin`
--

LOCK TABLES `admin` WRITE;
/*!40000 ALTER TABLE `admin` DISABLE KEYS */;
INSERT INTO `admin` VALUES ('admin','$2y$14$I82sSL0y4ZvGq6K1KKbPXe.Kyjh67he17Xu714GrlLd2nWK.JtCZG');
/*!40000 ALTER TABLE `admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pekerjaan`
--

DROP TABLE IF EXISTS `pekerjaan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pekerjaan` (
  `id_pekerjaan` int NOT NULL,
  `nama_pekerjaan` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id_pekerjaan`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pekerjaan`
--

LOCK TABLES `pekerjaan` WRITE;
/*!40000 ALTER TABLE `pekerjaan` DISABLE KEYS */;
INSERT INTO `pekerjaan` VALUES (1,'PNS'),(2,'Wirausaha'),(3,'Wiraswasta'),(4,'Belum Bekerja');
/*!40000 ALTER TABLE `pekerjaan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pendidikan`
--

DROP TABLE IF EXISTS `pendidikan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pendidikan` (
  `id_pendidikan` int NOT NULL,
  `jenjang_pendidikan` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id_pendidikan`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pendidikan`
--

LOCK TABLES `pendidikan` WRITE;
/*!40000 ALTER TABLE `pendidikan` DISABLE KEYS */;
INSERT INTO `pendidikan` VALUES (1,'SD'),(2,'SMP'),(3,'SMA'),(4,'DIPLOMA'),(5,'SARJANA'),(6,'MAGISTER'),(7,'DOKTOR');
/*!40000 ALTER TABLE `pendidikan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id_user` varchar(64) NOT NULL,
  `nama_user` varchar(60) DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `no_ktp` varchar(16) DEFAULT NULL,
  `pekerjaan` int NOT NULL,
  `pendidikan` int NOT NULL,
  `status` varchar(1) DEFAULT 'A',
  PRIMARY KEY (`id_user`),
  KEY `pekerjaan_idx` (`pekerjaan`),
  KEY `pendidikan_idx` (`pendidikan`),
  CONSTRAINT `pekerjaan` FOREIGN KEY (`pekerjaan`) REFERENCES `pekerjaan` (`id_pekerjaan`),
  CONSTRAINT `pendidikan` FOREIGN KEY (`pendidikan`) REFERENCES `pendidikan` (`id_pendidikan`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES ('223b1535-4e24-4580-a7c1-3f3788ca8222','Dinda Maani','1998-11-01','16710941198003',4,5,'P'),('4bfb21b6-4bee-4265-bd24-2e00906b5cf5','Dinda Maharani','1998-11-01','16710941198003',4,5,'A'),('cdd91171-ee08-4107-be01-8827c7bd5147','Tiara Arung Rahayu','1996-07-29','1671094290795002',2,5,'P');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'db_mekar'
--

--
-- Dumping routines for database 'db_mekar'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-03-01  9:49:58
