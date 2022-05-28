/*
 Navicat Premium Data Transfer

 Source Server         : GoCRUD
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3388
 Source Schema         : gocrud

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 28/05/2022 21:47:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `gender` varchar(10) DEFAULT NULL,
  `dateofbirth` date DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `is_actived` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` (`id`, `name`, `gender`, `dateofbirth`, `email`, `password`, `is_actived`, `created_at`, `updated_at`) VALUES (1, 'Demo User', 'Male', '1992-08-04', 'demo@demo.demo', '$2a$10$PODMD/Cu/ipmK11xhMBLPuAdUACWO3ghi8YsO/yqre8N97x1IvIoW', '0', '2022-05-28 21:35:29', '2022-05-28 21:35:29');
INSERT INTO `users` (`id`, `name`, `gender`, `dateofbirth`, `email`, `password`, `is_actived`, `created_at`, `updated_at`) VALUES (2, 'Admin Admin', 'Male', '1998-07-13', 'admin@admin.admin', '$2a$10$Be4MHuqS3cy2GuE6bUhIf.GsAwnLyUW/hbH/U.Ul7nbikhYVczUNK', '0', '2022-05-28 21:37:30', '2022-05-28 21:37:30');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
