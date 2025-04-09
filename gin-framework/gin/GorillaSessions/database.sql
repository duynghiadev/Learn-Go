CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `email` varchar(254) NOT NULL,
  `pswd_hash` varchar(70) NOT NULL,
  `created_at` datetime NOT NULL,
  `is_active` tinyint(1) NOT NULL,
  `ver_hash` varchar(500) NOT NULL,
  `timeout` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `username_UNIQUE` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- account in video
INSERT INTO users(username, email, pswd_hash, created_at, is_active, ver_hash, timeout)
VALUES ('Carlton', 'myemailaddr@email.com', '$2a$14$4Zo1x2n5VHMsXvAaaySI/.alvPRYvvOSGtdw/8HJFt/pRJB7g5EvG', '2022-05-25 02:25:37', '1', '$2a$14$4Zo1x2n5VHMsXvAaaySI/.alvPRYvvOSGtdw/8HJFt/pRJB7g5EvG', '2022-11-25 02:25:37');

-- this is my account
-- NOTE: $2a$10$cSPFnk6QjI41W6RJg3JUNuDjaFwzBChuqqQsPL204lNF5JWWT7Y82 -> duynghia123
INSERT INTO users(username, email, pswd_hash, created_at, is_active, ver_hash, timeout)
VALUES ('duynghia', 'test@gmail.com', '$2a$10$cSPFnk6QjI41W6RJg3JUNuDjaFwzBChuqqQsPL204lNF5JWWT7Y82',
        '2025-02-22 11:55:30', '1', '$2a$10$cSPFnk6QjI41W6RJg3JUNuDjaFwzBChuqqQsPL204lNF5JWWT7Y82',
        '2030-02-22 11:55:30');