/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80016
 Source Host           : localhost:3306
 Source Schema         : mall

 Target Server Type    : MySQL
 Target Server Version : 80016
 File Encoding         : 65001

 Date: 25/09/2021 14:11:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id（主键）',
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '用户名称',
  `password` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户密码',
  `avatar` varchar(300) DEFAULT NULL COMMENT '用户头像',
  `role` char(50) DEFAULT NULL COMMENT '用户角色',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '用户状态，1表示正常，0表示暂停',
  `created` char(20) DEFAULT NULL COMMENT '创建时间',
  `updated` char(20) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100037 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (100030, 'admin', '123', 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png', 'ROLE_admin', 1, '2021-07-16 10:35:02', NULL);
INSERT INTO `user` VALUES (100034, 'wang', '123', 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png', '会员', 0, '2021-09-09 23:10:19', '2021-07-16 12:11:57');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
