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

 Date: 25/09/2021 14:10:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '地址id',
  `user_id` varchar(200) DEFAULT NULL COMMENT '用户id',
  `name` varchar(255) DEFAULT NULL COMMENT '收货人姓名',
  `mobile` char(16) DEFAULT NULL COMMENT '手机号',
  `province` char(30) DEFAULT NULL COMMENT '省',
  `city` char(30) DEFAULT NULL COMMENT '城市',
  `district` char(30) DEFAULT NULL COMMENT '区/县',
  `detailed_address` varchar(200) DEFAULT NULL COMMENT '详细地址',
  `is_default` tinyint(1) DEFAULT NULL COMMENT '1为默认，0为非默认',
  `created` char(20) DEFAULT NULL COMMENT '创建时间',
  `updated` char(20) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1007 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of address
-- ----------------------------
BEGIN;
INSERT INTO `address` VALUES (1001, 'oUT385ZLmRr6R_a9xKSfSW9SekY1', '张三', '13900882903', '北京市AA', '北京市', 'xx区', 'xxx街道', 1, '2021-08-07 12:23:52', NULL);
INSERT INTO `address` VALUES (1002, 'oUT385ZLmRr6R_a9xKSfSW9SekY1', '张三', '13900882903', '北京市AA', '北京市', 'xx区', 'xxx街道', 1, '2021-08-07 12:24:15', NULL);
INSERT INTO `address` VALUES (1006, '100030', '王小明', '13500892350', '安徽省', '芜湖市', '芜湖县', 'xx小区xx栋', 0, '2021-08-08 09:41:09', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
