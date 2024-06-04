-- Create "healths" table
CREATE TABLE `healths` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `message` varchar(128) NOT NULL,
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
