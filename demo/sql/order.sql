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

 Date: 25/09/2021 14:11:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '订单id',
  `payment_status` int(1) DEFAULT '1' COMMENT '支付状态',
  `product_item` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '商品项',
  `total_price` decimal(20,0) DEFAULT NULL COMMENT '合计',
  `status` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '未发货' COMMENT '订单状态',
  `address_id` bigint(20) DEFAULT NULL COMMENT '地址id',
  `user_id` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '用户id',
  `admin_id` bigint(20) DEFAULT NULL COMMENT '管理员id',
  `created` char(20) DEFAULT NULL COMMENT '创建时间',
  `updated` char(20) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1004 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of order
-- ----------------------------
BEGIN;
INSERT INTO `order` VALUES (86, 1, '103,104,', 3598, '已发货', 1002, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', 100030, '2021-08-08 16:16:48', NULL);
INSERT INTO `order` VALUES (87, 1, '103,106,', 8698, '已发货', 1002, 'oUT385ZLmRr6111R_a9xKSfSW9SekYI', 100030, '2021-08-08 18:32:53', NULL);
INSERT INTO `order` VALUES (88, 1, '107,109,', 5519, '已发货', 1002, 'oUT385ZLm1Rr6R_a9xKSfSW9SekYI', 100030, '2021-08-08 18:33:06', NULL);
INSERT INTO `order` VALUES (89, 1, '107,109,110,', 5748, '已发货', 1002, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', 100030, '2021-08-08 18:33:42', NULL);
INSERT INTO `order` VALUES (1003, 1, '103,106,', 15187, '未发货', NULL, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', NULL, '2021-08-12 16:25:34', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
