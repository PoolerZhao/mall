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

 Date: 25/09/2021 14:11:28
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品id',
  `category_id` bigint(20) DEFAULT NULL COMMENT '类目id',
  `kind` tinyint(1) DEFAULT NULL COMMENT '商品类型，1为全新，0为二手',
  `title` varchar(240) DEFAULT NULL COMMENT '商品标题',
  `brand_id` bigint(20) DEFAULT NULL COMMENT '品牌id',
  `name` varchar(80) DEFAULT NULL COMMENT '商品名称',
  `price` decimal(10,2) DEFAULT NULL COMMENT '价格',
  `amount` int(10) DEFAULT NULL COMMENT '库存',
  `sales` int(10) DEFAULT NULL COMMENT '销量',
  `image_url` varchar(300) DEFAULT NULL COMMENT '图片路径',
  `send_address` varchar(20) DEFAULT NULL COMMENT '发货地址',
  `parcel_type` varchar(10) DEFAULT NULL COMMENT '快递类型',
  `sales_service` varchar(50) DEFAULT NULL COMMENT '售后服务',
  `creator_id` bigint(20) DEFAULT NULL COMMENT '创建者id',
  `status` tinyint(1) DEFAULT NULL COMMENT '1为上架，0为下架',
  `created` char(20) DEFAULT NULL COMMENT '创建时间',
  `updated` char(20) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=285 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of product
-- ----------------------------
BEGIN;
INSERT INTO `product` VALUES (103, 148, 1, '荣耀Play5T手机 华为手机新款正品', 111, '荣耀手机', 8888.00, 950, 900, 'http://localhost:8000/image/IMG_1190.JPG', '广东深圳', '申通快递', '提供发票,退换货承诺,服务承诺：该类商品，必须支持【七天退货】服务,,,', 100030, 1, '2021-07-28 13:53:59', '2021-09-24 16:14:01');
INSERT INTO `product` VALUES (104, 151, 1, '红米Redmi Note 9 Pro 手机 骁龙750G', 112, '小米手机', 1199.00, 1330, 54, 'http://localhost:8000/image/IMG_1191.JPG', '广东深圳', '圆通速递', '提供发票,服务承诺：该类商品，必须支持【七天退货】服务,退换货承诺,', 100030, 0, '2021-07-28 14:03:50', '2021-08-11 11:41:34');
INSERT INTO `product` VALUES (105, 152, 1, 'Apple/苹果 iPhone11 全新国行正品 全网通双卡', 113, '苹果11', 2388.00, 550, 4829, 'http://localhost:8000/image/IMG_1193.JPG', '浙江杭州', '顺丰快递', '提供发票,服务承诺：该类商品，必须支持【七天退货】服务,', 100030, 1, '2021-07-28 14:13:31', NULL);
INSERT INTO `product` VALUES (106, 156, 1, 'Apple苹果 MacBook Pro超薄便捷商务办公笔记本', 113, '苹果笔记本电脑', 6299.00, 1020, 445, 'http://localhost:8000/image/IMG_1194.JPG', '广东广州', '顺丰快递', '提供发票,保修服务,服务承诺：该类商品，必须支持【七天退货】服务,', 100030, 0, '2021-07-28 14:26:01', '2021-08-10 18:48:59');
INSERT INTO `product` VALUES (107, 158, 1, '联想小新air14/pro13 商务学生办公笔记本电脑', 114, '联想笔记本电脑', 3199.00, 505, 122, 'http://localhost:8000/image/IMG_1195.JPG', '湖北武汉', '申通快递', '提供发票,退换货承诺,', 100030, 0, '2021-07-28 14:36:30', '2021-08-10 18:49:00');
INSERT INTO `product` VALUES (108, 161, 1, '小米全面屏电视Pro 65英寸E65S 4K超清智能网络电视', 112, '小米全面屏电', 2219.00, 1920, 89, 'http://localhost:8000/image/IMG_1196.JPG', '浙江杭州', '顺丰快递', '提供发票,退换货承诺,服务承诺：该类商品，必须支持【七天退货】服务,', 100030, 1, '2021-07-28 14:45:15', NULL);
INSERT INTO `product` VALUES (109, 164, 1, 'Apple/苹果iPad8 正品平板电脑 10.2英寸', 113, 'Apple平板电脑', 2320.00, 1006, 103, 'http://localhost:8000/image/IMG_1197.JPG', '广东深圳', '申通快递', '提供发票,服务承诺：该类商品，必须支持【七天退货】服务,', 100030, 1, '2021-07-28 14:54:50', NULL);
INSERT INTO `product` VALUES (110, 167, 1, '小米手环6智能血氧心率监测 蓝牙男女款运动计步器', 112, '小米手环6', 229.00, 1288, 5, 'http://localhost:8000/image/IMG_1198.JPG', '北京', '顺丰快递', '提供发票,退换货承诺,,,', 100030, 1, '2021-07-28 15:03:03', '2021-07-29 10:01:39');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
