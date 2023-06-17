CREATE DATABASE IF NOT EXISTS `go_crud_api` DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
USE `go_crud_api`;
CREATE TABLE IF NOT EXISTS `category` (
      `id` int unsigned NOT NULL AUTO_INCREMENT,
      `name` varchar(100),
      PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;