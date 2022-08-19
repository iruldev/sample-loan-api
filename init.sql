DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `ktp` varchar(255) NOT NULL,
  `birth_date` varchar(255) NOT NULL,
  `sex` tinyint NOT NULL COMMENT '1 : Pria, 2 : Wanita',
  PRIMARY KEY (`id`),
  UNIQUE KEY `customers_ktp_index` (`ktp`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `loans`;
CREATE TABLE `loans` (
  `id` int NOT NULL AUTO_INCREMENT,
  `customer_id` int NOT NULL,
  `amount` int NOT NULL,
  `period` int NOT NULL,
  `purpose` enum('vacation','renovation','electronics','wedding','rent','car','investment') NOT NULL,
  PRIMARY KEY (`id`),
  KEY `customer_id` (`customer_id`),
  CONSTRAINT `loans_ibfk_1` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;