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

 Date: 25/09/2021 14:10:59
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for brand
-- ----------------------------
DROP TABLE IF EXISTS `brand`;
CREATE TABLE `brand` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '品牌id',
  `name` varchar(20) DEFAULT NULL COMMENT '品牌名称',
  `sort` int(20) DEFAULT NULL COMMENT '品牌排序',
  `created` char(20) DEFAULT NULL COMMENT '创建时间',
  `updated` char(20) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=158 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of brand
-- ----------------------------
BEGIN;
INSERT INTO `brand` VALUES (149, '苹果', 23, '2021-09-21 14:41:16', NULL);
INSERT INTO `brand` VALUES (150, '华为', 34, '2021-09-21 14:41:25', '2021-09-25 12:37:09');
INSERT INTO `brand` VALUES (153, '小米', 22, '2021-09-24 09:04:23', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
