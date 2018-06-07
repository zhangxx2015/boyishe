/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50631
Source Host           : localhost:3306
Source Database       : beedb

Target Server Type    : MYSQL
Target Server Version : 50631
File Encoding         : 65001

Date: 2017-07-31 02:43:15
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `carousels`
-- ----------------------------
DROP TABLE IF EXISTS `carousels`;
CREATE TABLE `carousels` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键ID',
  `uniqueId` varchar(20) NOT NULL DEFAULT '' COMMENT '编号',
  `addDate` varchar(15) NOT NULL COMMENT '记录添加日期',
  `addTime` varchar(15) NOT NULL COMMENT '记录添加时间',
  `addUtc` int(11) NOT NULL COMMENT '记录添加时间戳',
  `appType` varchar(20) NOT NULL COMMENT '数据类型',
  `isDeleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  `opatator` varchar(255) DEFAULT '' COMMENT '操作人',
  `lastUpdateUtc` int(11) NOT NULL COMMENT '操作时间戳',
  `title` varchar(255) DEFAULT '' COMMENT '标题',
  `imageUrl` varchar(255) DEFAULT '' COMMENT '图片地址',
  `showcaseUniqueId` varchar(20) NOT NULL COMMENT '展示项ID',
  `sortIndex` int(11) NOT NULL COMMENT '顺序',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique` (`uniqueId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of carousels
-- ----------------------------

-- ----------------------------
-- Table structure for `contents`
-- ----------------------------
DROP TABLE IF EXISTS `contents`;
CREATE TABLE `contents` (
  `views` int(11) DEFAULT NULL COMMENT '观看次数',
  `specialUniqueId` varchar(20) NOT NULL DEFAULT '' COMMENT '专题编号',
  `providerUniqueId` varchar(20) DEFAULT '' COMMENT '提供者编号',
  `title` varchar(255) DEFAULT '' COMMENT '名称',
  `description` varchar(255) DEFAULT '' COMMENT '描述',
  `price` int(11) NOT NULL COMMENT '价格(单位：分)',
  `chargingRule` int(11) NOT NULL COMMENT '计费规则(1.免费资源,2.会员专享,3.付费资源)',
  `days` int(11) NOT NULL COMMENT '有效期(单位：天数)',
  `invalidUtc` int(11) DEFAULT NULL COMMENT '过期时间戳',
  `resourceType` int(11) NOT NULL COMMENT '资源类型(1.音频,2.视频,3.活动)',
  `resourceUrl` varchar(255) NOT NULL DEFAULT '' COMMENT '资源地址',
  `uniqueId` varchar(20) NOT NULL DEFAULT '' COMMENT '编号',
  `addDate` varchar(15) NOT NULL COMMENT '记录添加日期',
  `addTime` varchar(15) NOT NULL COMMENT '记录添加时间',
  `addUtc` int(11) NOT NULL COMMENT '记录添加时间戳',
  `isDeleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  `opatator` varchar(255) DEFAULT '' COMMENT '操作人',
  `lastUpdateUtc` int(11) NOT NULL COMMENT '操作时间戳',
  PRIMARY KEY (`uniqueId`),
  UNIQUE KEY `unique` (`uniqueId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for `journals`
-- ----------------------------
DROP TABLE IF EXISTS `journals`;
CREATE TABLE `journals` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键ID',
  `uniqueId` varchar(20) NOT NULL DEFAULT '' COMMENT '编号',
  `addDate` varchar(15) NOT NULL COMMENT '记录添加日期',
  `addTime` varchar(15) NOT NULL COMMENT '记录添加时间',
  `addUtc` int(11) NOT NULL COMMENT '记录添加时间戳',
  `appType` varchar(20) NOT NULL COMMENT '数据类型',
  `isDeleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  `opatator` varchar(255) DEFAULT '' COMMENT '操作人',
  `lastUpdateUtc` int(11) NOT NULL COMMENT '操作时间戳',
  `userUniqueId` varchar(20) DEFAULT '' COMMENT '用户编号',
  `fee` int(11) NOT NULL COMMENT '金额(单位：分)',
  `itemUniqueId` varchar(20) DEFAULT '' COMMENT '项目编号',
  `chargingUniqueId` varchar(20) DEFAULT '' COMMENT '计费规则',
  `userLevel` int(11) NOT NULL COMMENT '用户级别',
  `discount` int(11) NOT NULL COMMENT '折扣(范围：0-100)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique` (`uniqueId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of journals
-- ----------------------------

-- ----------------------------
-- Table structure for `playhistory`
-- ----------------------------
DROP TABLE IF EXISTS `playhistory`;
CREATE TABLE `playhistory` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键ID',
  `uniqueId` varchar(20) NOT NULL DEFAULT '' COMMENT '编号',
  `addDate` varchar(15) NOT NULL COMMENT '记录添加日期',
  `addTime` varchar(15) NOT NULL COMMENT '记录添加时间',
  `addUtc` int(11) NOT NULL COMMENT '记录添加时间戳',
  `appType` varchar(20) NOT NULL COMMENT '数据类型',
  `isDeleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  `opatator` varchar(255) DEFAULT '' COMMENT '操作人',
  `lastUpdateUtc` int(11) NOT NULL COMMENT '操作时间戳',
  `userUniqueId` varchar(20) DEFAULT '' COMMENT '用户',
  `lastAccessUtc` int(11) NOT NULL COMMENT '最后更新',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique` (`uniqueId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of playhistory
-- ----------------------------

-- ----------------------------
-- Table structure for `purchased`
-- ----------------------------
DROP TABLE IF EXISTS `purchased`;
CREATE TABLE `purchased` (
  `showCaseUniqueId` varchar(20) NOT NULL COMMENT '展示项',
  `userUniqueId` varchar(20) DEFAULT '' COMMENT '用户',
  `itemUniqueId` varchar(20) DEFAULT '' COMMENT '项目',
  `uniqueId` varchar(20) NOT NULL DEFAULT '' COMMENT '编号',
  `addDate` varchar(15) NOT NULL COMMENT '记录添加日期',
  `addTime` varchar(15) NOT NULL COMMENT '记录添加时间',
  `addUtc` int(11) NOT NULL COMMENT '记录添加时间戳',
  `isDeleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  `opatator` varchar(255) DEFAULT '' COMMENT '操作人',
  `lastUpdateUtc` int(11) NOT NULL COMMENT '操作时间戳',
  UNIQUE KEY `unique` (`uniqueId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of purchased
-- ----------------------------

-- ----------------------------
-- Table structure for `specials`
-- ----------------------------
DROP TABLE IF EXISTS `specials`;
CREATE TABLE `specials` (
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `topicName` varchar(255) NOT NULL DEFAULT '' COMMENT '栏目信息',
  `sortIndex` int(11) DEFAULT NULL COMMENT '排序',
  `provider` varchar(255) NOT NULL DEFAULT '' COMMENT '提供者名称',
  `cover` varchar(255) NOT NULL DEFAULT '' COMMENT '展示图片的URL',
  `introduce` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `uniqueId` varchar(20) NOT NULL DEFAULT '' COMMENT '编号',
  `addDate` varchar(15) NOT NULL COMMENT '记录添加日期',
  `addTime` varchar(15) NOT NULL COMMENT '记录添加时间',
  `addUtc` int(11) NOT NULL COMMENT '记录添加时间戳',
  `isDeleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  `opatator` varchar(255) DEFAULT '' COMMENT '操作人',
  `lastUpdateUtc` int(11) NOT NULL COMMENT '操作时间戳',
  PRIMARY KEY (`title`),
  UNIQUE KEY `unique` (`uniqueId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for `topics`
-- ----------------------------
DROP TABLE IF EXISTS `topics`;
CREATE TABLE `topics` (
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '栏目名称',
  `uniqueId` varchar(20) NOT NULL DEFAULT '' COMMENT '编号',
  `addDate` varchar(15) NOT NULL COMMENT '记录添加日期',
  `addTime` varchar(15) NOT NULL COMMENT '记录添加时间',
  `addUtc` int(11) NOT NULL COMMENT '记录添加时间戳',
  `isDeleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  `opatator` varchar(255) DEFAULT '' COMMENT '操作人',
  `lastUpdateUtc` int(11) NOT NULL COMMENT '操作时间戳',
  PRIMARY KEY (`title`),
  UNIQUE KEY `unique` (`uniqueId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for `users`
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `loginName` varchar(255) NOT NULL DEFAULT '' COMMENT '账号(手机账号)',
  `password` varchar(255) DEFAULT '' COMMENT '登陆密码',
  `accessKey` varchar(16) DEFAULT NULL COMMENT '访问令牌',
  `expires` int(11) DEFAULT NULL COMMENT '令牌过期时间',
  `userRole` int(11) NOT NULL COMMENT '角色',
  `clientIMEI` varchar(255) DEFAULT '' COMMENT '客户端识别码',
  `clientIP` varchar(255) DEFAULT '' COMMENT '客户端IP地址',
  `focusIndustry` varchar(255) DEFAULT '' COMMENT '关注行业',
  `userName` varchar(255) DEFAULT '' COMMENT '真实姓名',
  `company` varchar(255) DEFAULT '' COMMENT '公司',
  `job` varchar(255) DEFAULT '' COMMENT '职务',
  `city` varchar(255) DEFAULT '' COMMENT '所在城市',
  `faceIcon` varchar(255) DEFAULT '' COMMENT '头像',
  `userLevel` int(11) DEFAULT '0' COMMENT '用户级别',
  `idCard` varchar(255) DEFAULT '' COMMENT '身份证',
  `bankCard` varchar(255) DEFAULT '' COMMENT '银行卡',
  `bankName` varchar(255) DEFAULT '' COMMENT '银行',
  `bankAddress` varchar(255) DEFAULT '' COMMENT '开户行地址',
  `ownedSales` varchar(255) DEFAULT '' COMMENT '隶属销售',
  `uniqueId` varchar(20) NOT NULL DEFAULT '' COMMENT '编号',
  `addDate` varchar(15) NOT NULL COMMENT '记录添加日期',
  `addTime` varchar(15) NOT NULL COMMENT '记录添加时间',
  `addUtc` int(11) NOT NULL COMMENT '记录添加时间戳',
  `isDeleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  `opatator` varchar(255) DEFAULT '' COMMENT '操作人',
  `lastUpdateUtc` int(11) NOT NULL COMMENT '操作时间戳',
  PRIMARY KEY (`loginName`,`uniqueId`),
  UNIQUE KEY `unique` (`uniqueId`) USING BTREE,
  UNIQUE KEY `loginName` (`loginName`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for `vcodes`
-- ----------------------------
DROP TABLE IF EXISTS `vcodes`;
CREATE TABLE `vcodes` (
  `loginName` varchar(255) NOT NULL DEFAULT '' COMMENT '账号(手机账号)',
  `vcode` varchar(6) NOT NULL DEFAULT '' COMMENT '验证码',
  `freeze` int(11) DEFAULT NULL COMMENT '冻结时间',
  `uniqueId` varchar(20) NOT NULL DEFAULT '' COMMENT '编号',
  `addDate` varchar(15) NOT NULL COMMENT '记录添加日期',
  `addTime` varchar(15) NOT NULL COMMENT '记录添加时间',
  `addUtc` int(11) NOT NULL COMMENT '记录添加时间戳',
  `isDeleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  `opatator` varchar(255) DEFAULT '' COMMENT '操作人',
  `lastUpdateUtc` int(11) NOT NULL COMMENT '操作时间戳',
  PRIMARY KEY (`loginName`,`uniqueId`),
  UNIQUE KEY `unique` (`uniqueId`) USING BTREE,
  UNIQUE KEY `loginName` (`loginName`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

