/*
 Navicat Premium Dump SQL

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80406 (8.4.6)
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 80406 (8.4.6)
 File Encoding         : 65001

 Date: 11/09/2025 18:19:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `comment` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `post_id` bigint UNSIGNED NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_comments_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_comments_user`(`user_id` ASC) USING BTREE,
  INDEX `fk_posts_comments`(`post_id` ASC) USING BTREE,
  CONSTRAINT `fk_comments_post` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_comments_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_posts_comments` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (1, '2025-09-10 17:33:45.000', '2025-09-10 17:33:48.000', NULL, '好看', 1, 1);
INSERT INTO `comments` VALUES (2, '2025-09-10 17:33:45.000', '2025-09-10 17:33:48.000', NULL, '好看3', 1, 1);

-- ----------------------------
-- Table structure for posts
-- ----------------------------
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `body` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `comment_num` bigint UNSIGNED NULL DEFAULT NULL,
  `status` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_posts_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_posts_user`(`user_id` ASC) USING BTREE,
  CONSTRAINT `fk_posts_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of posts
-- ----------------------------
INSERT INTO `posts` VALUES (1, '2025-09-10 15:34:04.000', '2025-09-10 15:34:07.000', NULL, '你好', '<html>1111</html>', 1, 2, NULL);
INSERT INTO `posts` VALUES (2, '2025-09-10 15:34:04.000', '2025-09-10 15:34:07.000', NULL, '你好0', '<html>1111</html>', 1, 0, NULL);
INSERT INTO `posts` VALUES (3, '2025-09-11 14:16:44.760', '2025-09-11 14:16:44.760', NULL, '新的一个', '新的一个', 1, 0, NULL);
INSERT INTO `posts` VALUES (4, '2025-09-11 14:29:48.464', '2025-09-11 14:29:48.464', NULL, '新的一个', '新的一个', 1, 0, NULL);
INSERT INTO `posts` VALUES (5, '2025-09-11 15:12:16.583', '2025-09-11 15:12:16.583', NULL, '新的一个', '新的一个', 1, 0, '');
INSERT INTO `posts` VALUES (6, '2025-09-11 15:13:34.759', '2025-09-11 15:13:34.759', NULL, '新的一个', '新的一个', 1, 0, '');
INSERT INTO `posts` VALUES (7, '2025-09-11 15:20:59.102', '2025-09-11 15:20:59.102', NULL, '新的一个', '新的一个', 1, 0, '');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `article_num` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '2025-09-10 15:33:52.000', '2025-09-10 15:33:55.000', NULL, '张三', 7);

SET FOREIGN_KEY_CHECKS = 1;
