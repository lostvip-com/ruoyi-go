/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.88.114-3307
 Source Server Type    : MySQL
 Source Server Version : 80019 (8.0.19)
 Source Host           : 192.168.88.114:3307
 Source Schema         : dpc

 Target Server Type    : MySQL
 Target Server Version : 80019 (8.0.19)
 File Encoding         : 65001

 Date: 15/05/2024 16:45:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app_gen_params
-- ----------------------------
DROP TABLE IF EXISTS `app_gen_params`;
CREATE TABLE `app_gen_params`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `use_flag` int NULL DEFAULT NULL,
  `param_no` int NULL DEFAULT NULL,
  `param_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '参量名称',
  `param_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '参量类型',
  `unit` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '单位',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  `monitor_type_id` int NULL DEFAULT NULL,
  `del_flag` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `create_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `update_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unq_param_monitor`(`param_no` ASC, `monitor_type_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28952 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '生成的参量表，相同的参量编号一定相同，监控类型可能不同 如 A相电压，不管在哪个监控类型下，都是参量X' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of app_gen_params
-- ----------------------------

-- ----------------------------
-- Table structure for dpc_task
-- ----------------------------
DROP TABLE IF EXISTS `dpc_task`;
CREATE TABLE `dpc_task`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '工号',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '密码',
  `prj_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '项  目  号',
  `task_content` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '任务内容',
  `start_date` date NULL DEFAULT NULL COMMENT '开始日期',
  `end_date` date NULL DEFAULT NULL COMMENT '结束日期',
  `work_days` bigint NULL DEFAULT NULL COMMENT '本月工时',
  `auto_submit` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '自动提交',
  `status` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '任务状态',
  `sort` bigint NULL DEFAULT NULL COMMENT '排序，大的优先',
  `update_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '更新者',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `create_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '创建人',
  `del_flag` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dpc_task
-- ----------------------------

-- ----------------------------
-- Table structure for gen_table
-- ----------------------------
DROP TABLE IF EXISTS `gen_table`;
CREATE TABLE `gen_table`  (
  `table_id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '表名称',
  `table_comment` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '表描述',
  `class_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '实体类名称',
  `tpl_category` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'crud' COMMENT '使用的模板（crud单表操作 tree树表操作）',
  `package_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '生成包路径',
  `module_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '生成模块名',
  `business_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '生成业务名',
  `function_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '生成功能名',
  `function_author` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '生成功能作者',
  `options` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '其它生成选项',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`table_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '代码生成业务表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of gen_table
-- ----------------------------
INSERT INTO `gen_table` VALUES (7, 'app_gen_params', '参量字典', 'AppGenParams', 'crud', 'mywork', 'main', 'params', '基础参量管理', 'lv', '', 'admin', '2024-02-28 14:21:50', 'admin', '2024-02-28 16:10:05', '所有参量的汇总');

-- ----------------------------
-- Table structure for gen_table_column
-- ----------------------------
DROP TABLE IF EXISTS `gen_table_column`;
CREATE TABLE `gen_table_column`  (
  `column_id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_id` bigint NULL DEFAULT NULL COMMENT '归属表编号',
  `column_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '列名称',
  `column_comment` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '列描述',
  `column_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '列类型',
  `go_type` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Go类型',
  `go_field` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Go字段名',
  `html_field` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'html字段名',
  `is_pk` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否主键（1是）',
  `is_increment` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否自增（1是）',
  `is_required` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否必填（1是）',
  `is_insert` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否为插入字段（1是）',
  `is_edit` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否编辑字段（1是）',
  `is_list` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否列表字段（1是）',
  `is_query` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否查询字段（1是）',
  `query_type` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'EQ' COMMENT '查询方式（等于、不等于、大于、小于、范围）',
  `html_type` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）',
  `dict_type` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `sort` int NULL DEFAULT NULL COMMENT '排序',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`column_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 548 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '代码生成业务表字段' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of gen_table_column
-- ----------------------------
INSERT INTO `gen_table_column` VALUES (401, 37, 'id', '主键', 'bigint(20)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (402, 37, 'pid', '问题ID', 'bigint(20)', 'int64', 'Pid', 'pid', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (403, 37, 'atype', '回复人类别', 'tinyint(1)', 'int', 'Atype', 'atype', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (404, 37, 'user_id', '回复人ID', 'bigint(20)', 'int64', 'UserId', 'userId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (405, 37, 'nick_name', '回复人名称', 'varchar(50)', 'string', 'NickName', 'nickName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (406, 37, 'avatar', '回复人头像', 'varchar(50)', 'string', 'Avatar', 'avatar', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (407, 37, 'remark', '回复内容', 'varchar(250)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (408, 37, 'img1', '回复图片1', 'varchar(100)', 'string', 'Img1', 'img1', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (409, 37, 'img2', '回复图片2', 'varchar(100)', 'string', 'Img2', 'img2', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (410, 37, 'img3', '回复图片3', 'varchar(100)', 'string', 'Img3', 'img3', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (411, 37, 'status', '状态', 'tinyint(1)', 'int', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (412, 37, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (413, 37, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (414, 38, 'id', '主键', 'bigint(20)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (415, 38, 'zx_id', '用户ID', 'bigint(20)', 'int64', 'ZxId', 'zxId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (416, 38, 'zx_name', '用户昵称', 'varchar(30)', 'string', 'ZxName', 'zxName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (417, 38, 'zx_avatar', '咨询头像', 'varchar(50)', 'string', 'ZxAvatar', 'zxAvatar', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (418, 38, 'zj_id', '专家ID', 'bigint(20)', 'int64', 'ZjId', 'zjId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (419, 38, 'remark', '问题描述', 'varchar(250)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (420, 38, 'img1', '问题图片1', 'varchar(100)', 'string', 'Img1', 'img1', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (421, 38, 'img2', '问题图片2', 'varchar(100)', 'string', 'Img2', 'img2', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (422, 38, 'img3', '问题图片3', 'varchar(100)', 'string', 'Img3', 'img3', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (423, 38, 'ispay', '是否付费', 'tinyint(1)', 'int', 'Ispay', 'ispay', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (424, 38, 'pay_time', '付费时间', 'datetime', 'Time', 'PayTime', 'payTime', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'datatime', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (425, 38, 'pay_no', '订单号', 'varchar(50)', 'string', 'PayNo', 'payNo', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (426, 38, 'status', '状态', 'tinyint(4)', 'int', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (427, 38, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (428, 38, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (429, 39, 'wid', '申请ID', 'int(50)', 'int64', 'Wid', 'wid', '1', '1', '0', '1', '0', '1', '1', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (430, 39, 'uid', '用户ID', 'int(50)', 'int64', 'Uid', 'uid', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (431, 39, 'realname', '姓名', 'varchar(50)', 'string', 'Realname', 'realname', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (432, 39, 'idno', '身份证号', 'varchar(50)', 'string', 'Idno', 'idno', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (433, 39, 'wmoney', '提现金额', 'decimal(10,0)', 'int64', 'Wmoney', 'wmoney', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (434, 39, 'wbankname', '银行名称', 'varchar(50)', 'string', 'Wbankname', 'wbankname', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (435, 39, 'wbankaccount', '银行帐户', 'varchar(50)', 'string', 'Wbankaccount', 'wbankaccount', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (436, 39, 'wrealname', '银行户名', 'varchar(50)', 'string', 'Wrealname', 'wrealname', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (437, 39, 'wopenbank', '开户行', 'varchar(50)', 'string', 'Wopenbank', 'wopenbank', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (438, 39, 'wcreatetime', '申请时间', 'varchar(50)', 'string', 'Wcreatetime', 'wcreatetime', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (439, 39, 'wstatus', '审核结果', 'int(11)', 'int64', 'Wstatus', 'wstatus', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (440, 39, 'wremark', '审核意见', 'varchar(255)', 'string', 'Wremark', 'wremark', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (441, 39, 'waudittime', '审核时间', 'varchar(50)', 'string', 'Waudittime', 'waudittime', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (442, 40, 'id', '主键', 'bigint(20)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (443, 40, 'uname', '手机号', 'varchar(11)', 'string', 'Uname', 'uname', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (444, 40, 'puid', '推荐人ID', 'bigint(20)', 'int64', 'Puid', 'puid', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (445, 40, 'puname', '推荐人手机号', 'varchar(11)', 'string', 'Puname', 'puname', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (446, 40, 'real_name', '姓名', 'varchar(30)', 'string', 'RealName', 'realName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (447, 40, 'idno', '身份证号', 'varchar(18)', 'string', 'Idno', 'idno', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (448, 40, 'avatar', '头像', 'varchar(50)', 'string', 'Avatar', 'avatar', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (449, 40, 'job', '职业', 'varchar(50)', 'string', 'Job', 'job', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (450, 40, 'utype', '用户类别', 'tinyint(1)', 'int', 'Utype', 'utype', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (451, 40, 'upwd', '密码', 'varchar(50)', 'string', 'Upwd', 'upwd', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (452, 40, 'salt', '密码盐', 'varchar(45)', 'string', 'Salt', 'salt', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (453, 40, 'idpic1', '身份证正面', 'varchar(100)', 'string', 'Idpic1', 'idpic1', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (454, 40, 'idpic2', '身份证反面', 'varchar(100)', 'string', 'Idpic2', 'idpic2', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (455, 40, 'pimg1', '职业资格认证1', 'varchar(100)', 'string', 'Pimg1', 'pimg1', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (456, 40, 'pimg2', '职业资格认证2', 'varchar(100)', 'string', 'Pimg2', 'pimg2', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (457, 40, 'pimg3', '职业资格认证3', 'varchar(100)', 'string', 'Pimg3', 'pimg3', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (458, 40, 'status', '状态', 'tinyint(1)', 'int', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (459, 40, 'im_money', '在线咨询费用', 'decimal(8,2)', 'float64', 'ImMoney', 'imMoney', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (460, 40, 'tel_money', '电话咨询费用', 'decimal(8,2)', 'float64', 'TelMoney', 'telMoney', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (461, 40, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (462, 40, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (463, 41, 'id', '主键', 'bigint(20)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (464, 41, 'source_id', '注册ID', 'varchar(50)', 'string', 'SourceId', 'sourceId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (465, 41, 'source_type', '注册类别', 'tinyint(1)', 'int', 'SourceType', 'sourceType', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (466, 41, 'recommender', '推荐人', 'varchar(20)', 'string', 'Recommender', 'recommender', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (467, 41, 'nick_name', '昵称', 'varchar(50)', 'string', 'NickName', 'nickName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (468, 41, 'avatar', '头像', 'varchar(50)', 'string', 'Avatar', 'avatar', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (469, 41, 'zj_id', '专家ID', 'bigint(20)', 'int64', 'ZjId', 'zjId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (470, 41, 'status', '状态', 'tinyint(11)', 'int64', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (471, 41, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (472, 41, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (473, 1, 'config_id', '参数主键', 'int(5)', 'int', 'ConfigId', 'configId', '1', '1', '0', '1', '0', '1', '1', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (474, 1, 'config_name', '参数名称', 'varchar(100)', 'string', 'ConfigName', 'configName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (475, 1, 'config_key', '参数键名', 'varchar(100)', 'string', 'ConfigKey', 'configKey', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (476, 1, 'config_value', '参数键值', 'varchar(500)', 'string', 'ConfigValue', 'configValue', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (477, 1, 'config_type', '系统内置（Y是 N否）', 'char(1)', 'string', 'ConfigType', 'configType', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (478, 1, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (479, 1, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (480, 1, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (481, 1, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (482, 1, 'remark', '备注', 'varchar(500)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (483, 2, 'config_id', '参数主键', 'int(5)', 'int', 'ConfigId', 'configId', '1', '1', '0', '1', '0', '1', '1', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (484, 2, 'config_name', '参数名称', 'varchar(100)', 'string', 'ConfigName', 'configName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (485, 2, 'config_key', '参数键名', 'varchar(100)', 'string', 'ConfigKey', 'configKey', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (486, 2, 'config_value', '参数键值', 'varchar(500)', 'string', 'ConfigValue', 'configValue', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (487, 2, 'config_type', '系统内置（Y是 N否）', 'char(1)', 'string', 'ConfigType', 'configType', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (488, 2, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (489, 2, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (490, 2, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (491, 2, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (492, 2, 'remark', '备注', 'varchar(500)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (493, 3, 'config_id', '参数主键', 'int(5)', 'int', 'ConfigId', 'configId', '1', '1', '0', '1', '0', '1', '1', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (494, 3, 'config_name', '参数名称', 'varchar(100)', 'string', 'ConfigName', 'configName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (495, 3, 'config_key', '参数键名', 'varchar(100)', 'string', 'ConfigKey', 'configKey', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (496, 3, 'config_value', '参数键值', 'varchar(500)', 'string', 'ConfigValue', 'configValue', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (497, 3, 'config_type', '系统内置（Y是 N否）', 'char(1)', 'string', 'ConfigType', 'configType', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (498, 3, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (499, 3, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (500, 3, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (501, 3, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (502, 3, 'remark', '备注', 'varchar(500)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (503, 4, 'config_id', '参数主键', 'int(5)', 'int', 'ConfigId', 'configId', '1', '1', '0', '1', '0', '1', '1', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (504, 4, 'config_name', '参数名称', 'varchar(100)', 'string', 'ConfigName', 'configName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (505, 4, 'config_key', '参数键名', 'varchar(100)', 'string', 'ConfigKey', 'configKey', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (506, 4, 'config_value', '参数键值', 'varchar(500)', 'string', 'ConfigValue', 'configValue', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (507, 4, 'config_type', '系统内置（Y是 N否）', 'char(1)', 'string', 'ConfigType', 'configType', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (508, 4, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (509, 4, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (510, 4, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (511, 4, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (512, 4, 'remark', '备注', 'varchar(500)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (513, 5, 'config_id', '参数主键11', 'int(5)', 'int', 'ConfigId', 'configId', '1', '1', '0', '1', '0', '1', '1', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (514, 5, 'config_name', '参数名称111', 'varchar(100)', 'string', 'ConfigName', 'configName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (515, 5, 'config_key', '参数键名111', 'varchar(100)', 'string', 'ConfigKey', 'configKey', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (516, 5, 'config_value', '参数键值', 'varchar(500)', 'string', 'ConfigValue', 'configValue', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (517, 5, 'config_type', '系统内置（Y是 N否）', 'char(1)', 'string', 'ConfigType', 'configType', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (518, 5, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (519, 5, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (520, 5, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (521, 5, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (522, 5, 'remark', '备注', 'varchar(500)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (523, 6, 'oper_id', '日志主键', 'bigint(20)', 'int64', 'OperId', 'operId', '1', '1', '0', '1', '0', '1', '1', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (524, 6, 'title', '模块标题', 'varchar(50)', 'string', 'Title', 'title', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (525, 6, 'business_type', '业务类型（0其它 1新增 2修改 3删除）', 'int(11)', 'int64', 'BusinessType', 'businessType', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (526, 6, 'method', '方法名称', 'varchar(100)', 'string', 'Method', 'method', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (527, 6, 'request_method', '请求方式', 'varchar(10)', 'string', 'RequestMethod', 'requestMethod', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (528, 6, 'operator_type', '操作类别（0其它 1后台用户 2手机端用户）', 'int(11)', 'int64', 'OperatorType', 'operatorType', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (529, 6, 'oper_name', '操作人员', 'varchar(50)', 'string', 'OperName', 'operName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (530, 6, 'dept_name', '部门名称', 'varchar(50)', 'string', 'DeptName', 'deptName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (531, 6, 'oper_url', '请求URL', 'varchar(255)', 'string', 'OperUrl', 'operUrl', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (532, 6, 'oper_ip', '主机地址', 'varchar(50)', 'string', 'OperIp', 'operIp', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (533, 6, 'oper_location', '操作地点', 'varchar(255)', 'string', 'OperLocation', 'operLocation', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (534, 6, 'oper_param', '请求参数', 'varchar(2000)', 'string', 'OperParam', 'operParam', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (535, 6, 'json_result', '返回参数', 'varchar(2000)', 'string', 'JsonResult', 'jsonResult', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (536, 6, 'status', '操作状态（0正常 1异常）', 'int(11)', 'int64', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (537, 6, 'error_msg', '错误消息', 'varchar(2000)', 'string', 'ErrorMsg', 'errorMsg', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (538, 6, 'oper_time', '操作时间', 'datetime', 'Time', 'OperTime', 'operTime', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (539, 7, 'id', 'ID', 'bigint', 'int64', 'Id', 'id', '1', '1', '0', '0', '1', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (540, 7, 'use_flag', '是否使用', 'int', 'int', 'UseFlag', 'useFlag', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (541, 7, 'param_no', '参量号', 'int', 'int', 'ParamNo', 'paramNo', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (542, 7, 'param_name', '参量名称', 'varchar(255)', 'string', 'ParamName', 'paramName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (543, 7, 'param_type', '参量类型', 'varchar(255)', 'string', 'ParamType', 'paramType', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (544, 7, 'unit', '单位', 'varchar(255)', 'string', 'Unit', 'unit', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (545, 7, 'remark', '备注信息', 'varchar(255)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (546, 7, 'monitor_type_id', '监控类型', 'int', 'int', 'MonitorTypeId', 'monitorTypeId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (547, 7, 'create_time', '', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 9, 'admin', NULL, '', NULL);

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `config_id` int NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '参数键值',
  `config_type` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'N' COMMENT '系统内置（Y是 N否）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`config_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '参数配置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (1, '主框架页-默认皮肤样式名称', 'system.index.skinName', 'skin-blue', 'Y', 'admin', '2018-03-16 11:33:00', '', '2021-06-20 13:51:55', '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow');
INSERT INTO `sys_config` VALUES (2, '用户管理-账号初始密码', 'system.user.initPassword', '123456', 'Y', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '初始化密码 123456');
INSERT INTO `sys_config` VALUES (3, '主框架页-侧边栏主题', 'system.index.sideTheme', 'theme-dark', 'Y', 'admin', '2018-03-16 11:33:00', '', '2020-02-05 10:46:28', '深黑主题theme-dark，浅色主题theme-light，深蓝主题theme-blue');
INSERT INTO `sys_config` VALUES (4, '静态资源网盘存储', 'system.resource.url', '/static', 'Y', 'admin', '2020-02-18 20:10:33', '', '2020-03-23 20:51:39', 'public目录下的静态资源存储到OSS/COS等网盘，将键值设为/public表示本地');

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `dept_id` bigint NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint NULL DEFAULT 0 COMMENT '父部门id',
  `ancestors` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `order_num` int NULL DEFAULT 0 COMMENT '显示顺序',
  `leader` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `tenant_id` bigint NULL DEFAULT 0 COMMENT '租户id',
  PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 115 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '部门表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (100, 0, '0', 'LOSTVIP', 0, 'admin', '15888888888', '110@qq.com', '0', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2023-09-15 11:34:01', 0);
INSERT INTO `sys_dept` VALUES (110, 100, '0,100', '铁塔公司', 100, '', '', '', '0', '0', 'admin', '2019-12-02 17:07:02', 'admin', '2023-09-15 11:54:50', 0);
INSERT INTO `sys_dept` VALUES (111, 100, '0,100', '电影学院', 51, '曾尚兵1', '18788996255', 'ddd@163.com', '0', '2', 'admin', '2020-03-01 09:40:48', 'admin', '2020-03-01 09:40:55', 0);
INSERT INTO `sys_dept` VALUES (112, 100, '0,100', '公司运维', 10, '曾尚兵', '18788996255', 'ddd@163.com', '0', '0', 'admin', '2020-03-21 16:30:26', 'admin', '2023-09-15 11:55:24', 0);
INSERT INTO `sys_dept` VALUES (113, 100, '0,100', '测试一', 2, '测试一', '', '', '0', '2', 'admin', '2023-09-15 11:38:29', 'admin', '2023-09-15 11:39:15', 0);
INSERT INTO `sys_dept` VALUES (114, 100, '0,100', '测试二', 100, '', '', '', '0', '0', 'admin', '2023-09-15 11:39:05', 'admin', '2023-09-15 11:55:13', 0);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `dict_code` bigint NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `dict_sort` int NULL DEFAULT 0 COMMENT '字典排序',
  `dict_label` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `css_class` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '表格回显样式',
  `is_default` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'N' COMMENT '是否默认（Y是 N否）',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '字典数据表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (1, 1, '男', '0', 'sys_user_sex', '', 'default', 'Y', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2020-03-23 20:45:09', '[[*{remark}]]');
INSERT INTO `sys_dict_data` VALUES (2, 2, '女', '1', 'sys_user_sex', '', '', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '性别女');
INSERT INTO `sys_dict_data` VALUES (4, 1, '显示', '0', 'sys_show_hide', '', 'primary', 'Y', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '显示菜单');
INSERT INTO `sys_dict_data` VALUES (5, 2, '隐藏', '1', 'sys_show_hide', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '隐藏菜单');
INSERT INTO `sys_dict_data` VALUES (6, 1, '正常', '0', 'sys_normal_disable', '', 'primary', 'Y', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '正常状态');
INSERT INTO `sys_dict_data` VALUES (7, 2, '停用', '1', 'sys_normal_disable', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '停用状态');
INSERT INTO `sys_dict_data` VALUES (8, 1, '正常', '0', 'sys_job_status', '', 'primary', 'Y', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '正常状态');
INSERT INTO `sys_dict_data` VALUES (9, 2, '暂停', '1', 'sys_job_status', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '停用状态');
INSERT INTO `sys_dict_data` VALUES (10, 1, '默认', 'DEFAULT', 'sys_job_group', '', '', 'Y', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '默认分组');
INSERT INTO `sys_dict_data` VALUES (11, 2, '系统', 'SYSTEM', 'sys_job_group', '', '', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统分组');
INSERT INTO `sys_dict_data` VALUES (12, 1, '是', 'Y', 'sys_yes_no', '', 'primary', 'Y', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统默认是');
INSERT INTO `sys_dict_data` VALUES (13, 2, '否', 'N', 'sys_yes_no', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统默认否');
INSERT INTO `sys_dict_data` VALUES (14, 1, '通知', '1', 'sys_notice_type', '', 'warning', 'Y', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '通知');
INSERT INTO `sys_dict_data` VALUES (15, 2, '公告', '2', 'sys_notice_type', '', 'success', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '公告');
INSERT INTO `sys_dict_data` VALUES (16, 1, '正常', '0', 'sys_notice_status', '', 'primary', 'Y', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '正常状态');
INSERT INTO `sys_dict_data` VALUES (17, 2, '关闭', '1', 'sys_notice_status', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '关闭状态');
INSERT INTO `sys_dict_data` VALUES (18, 1, '新增', '1', 'sys_oper_type', '', 'info', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '新增操作');
INSERT INTO `sys_dict_data` VALUES (19, 2, '修改', '2', 'sys_oper_type', '', 'info', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '修改操作');
INSERT INTO `sys_dict_data` VALUES (20, 3, '删除', '3', 'sys_oper_type', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '删除操作');
INSERT INTO `sys_dict_data` VALUES (21, 4, '授权', '4', 'sys_oper_type', '', 'primary', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '授权操作');
INSERT INTO `sys_dict_data` VALUES (22, 5, '导出', '5', 'sys_oper_type', '', 'warning', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '导出操作');
INSERT INTO `sys_dict_data` VALUES (23, 6, '导入', '6', 'sys_oper_type', '', 'warning', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '导入操作');
INSERT INTO `sys_dict_data` VALUES (24, 7, '强退', '7', 'sys_oper_type', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '强退操作');
INSERT INTO `sys_dict_data` VALUES (25, 8, '生成代码', '8', 'sys_oper_type', '', 'warning', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '生成操作');
INSERT INTO `sys_dict_data` VALUES (26, 9, '清空数据', '9', 'sys_oper_type', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '清空操作');
INSERT INTO `sys_dict_data` VALUES (27, 1, '成功', '0', 'sys_common_status', '', 'primary', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '正常状态');
INSERT INTO `sys_dict_data` VALUES (28, 2, '失败', '1', 'sys_common_status', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '停用状态');
INSERT INTO `sys_dict_data` VALUES (29, 0, '免费用户', '0', 'zjuser_type', NULL, 'default', 'Y', '0', 'admin', '2019-12-02 16:56:16', 'admin', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (30, 1, '付费用户', '1', 'zjuser_type', NULL, 'primary', 'Y', '0', 'admin', '2019-12-02 16:56:40', 'admin', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (31, 0, '微信用户', '0', 'zxuser_type', NULL, 'default', 'Y', '0', 'admin', '2019-12-02 17:14:32', 'admin', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (32, 1, 'QQ用户', '1', 'zxuser_type', NULL, 'primary', 'N', '0', 'admin', '2019-12-02 17:14:55', 'admin', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (33, 2, '抖音用户', '2', 'zxuser_type', NULL, 'primary', 'N', '0', 'admin', '2019-12-02 17:15:21', 'admin', NULL, NULL);

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
  `dict_id` bigint NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `dict_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_id`) USING BTREE,
  UNIQUE INDEX `dict_type`(`dict_type` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '字典类型表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES (1, '用户性别', 'sys_user_sex', '0', 'admin', '2018-03-16 11:33:00', '', '2020-03-01 09:41:19', '用户性别列表');
INSERT INTO `sys_dict_type` VALUES (2, '菜单状态', 'sys_show_hide', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '菜单状态列表');
INSERT INTO `sys_dict_type` VALUES (3, '系统开关', 'sys_normal_disable', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统开关列表');
INSERT INTO `sys_dict_type` VALUES (4, '任务状态', 'sys_job_status', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '任务状态列表');
INSERT INTO `sys_dict_type` VALUES (5, '任务分组', 'sys_job_group', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '任务分组列表');
INSERT INTO `sys_dict_type` VALUES (6, '系统是否', 'sys_yes_no', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统是否列表');
INSERT INTO `sys_dict_type` VALUES (7, '通知类型', 'sys_notice_type', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '通知类型列表');
INSERT INTO `sys_dict_type` VALUES (8, '通知状态', 'sys_notice_status', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '通知状态列表');
INSERT INTO `sys_dict_type` VALUES (9, '操作类型', 'sys_oper_type', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '操作类型列表');
INSERT INTO `sys_dict_type` VALUES (10, '系统状态', 'sys_common_status', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '登录状态列表');
INSERT INTO `sys_dict_type` VALUES (11, '专家用户类别', 'zjuser_type', '0', 'admin', '2019-12-02 16:55:42', 'admin', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (12, '咨询用户类别', 'zxuser_type', '0', 'admin', '2019-12-02 17:14:07', 'admin', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (13, '测试11', 'test3dddd', '0', 'admin', '2020-02-05 16:23:06', '', '2020-03-23 20:23:25', '');

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job`  (
  `job_id` bigint NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `job_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_params` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '参数',
  `job_group` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `invoke_target` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '调用目标字符串',
  `cron_expression` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT 'cron执行表达式',
  `misfire_policy` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1' COMMENT '计划执行策略（1多次执行 2执行一次）',
  `concurrent` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1' COMMENT '是否并发执行（0允许 1禁止）',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '状态（0正常 1暂停）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`job_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '定时任务调度表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_job
-- ----------------------------
INSERT INTO `sys_job` VALUES (10, 'test1', '', 'DEFAULT', 'test1', '* * * * *', '1', '1', '1', 'admin', '2020-02-26 15:30:27', '', '2020-03-24 16:12:46', '');
INSERT INTO `sys_job` VALUES (12, 'test2', 'helloworld|yjgo', 'DEFAULT', 'test2', '@every 3s', '1', '1', '1', 'admin', '2020-02-27 10:20:26', '', NULL, '');

-- ----------------------------
-- Table structure for sys_job_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_job_log`;
CREATE TABLE `sys_job_log`  (
  `job_log_id` bigint NOT NULL AUTO_INCREMENT COMMENT '任务日志ID',
  `job_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '任务名称',
  `job_group` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '任务组名',
  `invoke_target` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '调用目标字符串',
  `job_message` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '日志信息',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '执行状态（0正常 1失败）',
  `exception_info` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '异常信息',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`job_log_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '定时任务调度日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_job_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_logininfor
-- ----------------------------
DROP TABLE IF EXISTS `sys_logininfor`;
CREATE TABLE `sys_logininfor`  (
  `info_id` bigint NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `login_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '操作系统',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '提示消息',
  `login_time` datetime NULL DEFAULT NULL COMMENT '访问时间',
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 660 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '系统访问记录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_logininfor
-- ----------------------------

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `menu_id` bigint NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `menu_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '菜单名称',
  `parent_id` bigint NULL DEFAULT 0 COMMENT '父菜单ID',
  `order_num` int NULL DEFAULT 0 COMMENT '显示顺序',
  `url` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '#' COMMENT '请求地址',
  `target` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '打开方式（menuItem页签 menuBlank新窗口）',
  `menu_type` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '菜单类型（M目录 C菜单 F按钮）',
  `visible` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '菜单状态（0显示 1隐藏）',
  `perms` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '权限标识',
  `icon` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '#' COMMENT '菜单图标',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1085 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '菜单权限表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, '系统管理', 0, 1, '#', '', 'M', '0', '', 'fa fa-gear', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统管理目录');
INSERT INTO `sys_menu` VALUES (2, '系统监控', 0, 2, '#', '', 'M', '0', '', 'fa fa-video-camera', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统监控目录');
INSERT INTO `sys_menu` VALUES (3, '系统工具', 0, 3, '#', '', 'M', '0', '', 'fa fa-bars', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统工具目录');
INSERT INTO `sys_menu` VALUES (4, '业务管理', 0, 1, '#', 'menuItem', 'M', '0', NULL, 'fa fa-newspaper-o', 'admin', '2019-12-02 16:39:09', 'admin', NULL, '');
INSERT INTO `sys_menu` VALUES (5, '实例演示', 0, 5, '#', '', 'M', '0', '', 'fa fa-desktop', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统工具目录');
INSERT INTO `sys_menu` VALUES (6, '表单演示', 5, 1, '#', 'menuItem', 'M', '0', '', 'fa fa-asterisk', 'admin', '2018-03-16 11:33:00', 'admin', '2021-06-19 23:28:01', '表单演示');
INSERT INTO `sys_menu` VALUES (7, '表格演示', 5, 2, '#', '', 'M', '0', '', '', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '表格演示');
INSERT INTO `sys_menu` VALUES (8, '弹框演示', 5, 3, '#', '', 'M', '0', '', '', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '弹框演示');
INSERT INTO `sys_menu` VALUES (9, '操作演示', 5, 4, '#', '', 'M', '0', '', '', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '操作演示');
INSERT INTO `sys_menu` VALUES (10, '报表演示', 5, 5, '#', '', 'M', '0', '', '', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '报表演示');
INSERT INTO `sys_menu` VALUES (11, '图标演示', 5, 6, '#', 'menuItem', 'M', '0', '', 'fa fa-bookmark-o', 'admin', '2018-03-16 11:33:00', 'admin', '2024-02-29 02:23:37', '图标演示');
INSERT INTO `sys_menu` VALUES (12, '栅格演示', 6, 2, '/demo/form/grid', '', 'C', '0', 'demo:grid:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (13, '下拉框', 6, 3, '/demo/form/select', '', 'C', '0', 'demo:select:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (14, '时间轴', 6, 4, '/demo/form/timeline', '', 'C', '0', 'demo:timeline:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (15, '基本表单', 6, 5, '/demo/form/basic', '', 'C', '0', 'demo:basic:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (16, '卡片列表', 6, 6, '/demo/form/cards', '', 'C', '0', 'demo:cards:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (17, '功能扩展', 6, 7, '/demo/form/jasny', '', 'C', '0', 'demo:jasny:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (18, '拖动排序', 6, 8, '/demo/form/sortable', '', 'C', '0', 'demo:sortable:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (19, '选项卡&面板', 6, 9, '/demo/form/tabs_panels', '', 'C', '0', 'demo:tabs_panels:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (20, '表单校验', 6, 10, '/demo/form/validate', '', 'C', '0', 'demo:validate:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (21, '表单向导', 6, 11, '/demo/form/wizard', '', 'C', '0', 'demo:wizard:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (22, '文件上传', 6, 12, '/demo/form/upload', '', 'C', '0', 'demo:upload:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (23, '日期和时间', 6, 13, '/demo/form/datetime', '', 'C', '0', 'demo:datetime:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (24, '富文本编辑器', 6, 14, '/demo/form/summernote', '', 'C', '0', 'demo:summernote:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (25, '左右互选', 6, 15, '/demo/form/duallistbox', '', 'C', '0', 'demo:duallistbox:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (27, '按钮演示', 6, 1, '/demo/form/button', 'menuItem', 'C', '0', 'demo:button:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', '', '2020-02-04 08:46:48', '');
INSERT INTO `sys_menu` VALUES (28, '数据汇总', 7, 2, '/demo/table/footer', '', 'C', '0', 'demo:footer:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (29, '组合表头', 7, 3, '/demo/table/groupHeader', '', 'C', '0', 'demo:groupHeader:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (31, '记住翻页', 7, 5, '/demo/table/remember', '', 'C', '0', 'demo:remember:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (32, '跳转指定页', 7, 6, '/demo/table/pageGo', '', 'C', '0', 'demo:pageGo:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (33, '查询参数', 7, 7, '/demo/table/params', '', 'C', '0', 'demo:params:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (35, '点击加载表格', 7, 9, '/demo/table/button', '', 'C', '0', 'demo:button:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (36, '表格冻结列', 7, 10, '/demo/table/fixedColumns', '', 'C', '0', 'demo:fixedColumns:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (37, '触发事件', 7, 11, '/demo/table/event', '', 'C', '0', 'demo:event:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (38, '细节视图', 7, 12, '/demo/table/detail', '', 'C', '0', 'demo:detail:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (39, '父子视图', 7, 13, '/demo/table/child', '', 'C', '0', 'demo:child:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (40, '图片预览', 7, 14, '/demo/table/image', '', 'C', '0', 'demo:image:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (41, '动态增删改查', 7, 15, '/demo/table/curd', '', 'C', '0', 'demo:curd:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (42, '表格拖曳', 7, 16, '/demo/table/recorder', '', 'C', '0', 'demo:recorder:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (43, '行内编辑', 7, 17, '/demo/table/editable', '', 'C', '0', 'demo:editable:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (44, '其它操作', 7, 18, '/demo/table/other', '', 'C', '0', 'demo:other:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (45, '查询条件', 7, 1, '/demo/table/lv_sql', '', 'C', '0', 'demo:lv_sql:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (46, '弹层组件', 8, 2, '/demo/modal/layer', '', 'C', '0', 'demo:layer:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (47, '弹层表格', 8, 3, '/demo/modal/table', '', 'C', '0', 'demo:table:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (48, '模态窗口', 8, 1, '/demo/modal/dialog', '', 'C', '0', 'demo:dialog:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (49, '其他操作', 9, 2, '/demo/operate/other', '', 'C', '0', 'demo:other:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (50, '表格操作', 9, 1, '/demo/operate/table', '', 'C', '0', 'demo:table:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (51, 'Peity', 10, 2, '/demo/report/metrics', '', 'C', '0', 'demo:metrics:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (52, 'SparkLine', 10, 3, '/demo/report/peity', '', 'C', '0', 'demo:peity:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (53, '图表组合', 10, 4, '/demo/report/sparkline', '', 'C', '0', 'demo:sparkline:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (54, '百度Echarts', 10, 1, '/demo/report/echarts', '', 'C', '0', 'demo:echarts:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (55, 'Glyphicons', 11, 2, '/demo/icon/glyphicons', '', 'C', '0', 'demo:glyphicons:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (56, 'Font Awesome', 11, 1, '/demo/icon/fontawesome', '', 'C', '0', 'demo:fontawesome:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (100, '用户管理', 1, 1, '/system/user', '', 'C', '0', 'system:user:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '用户管理菜单');
INSERT INTO `sys_menu` VALUES (101, '角色管理', 1, 2, '/system/role', '', 'C', '0', 'system:role:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '角色管理菜单');
INSERT INTO `sys_menu` VALUES (102, '菜单管理', 1, 3, '/system/menu', '', 'C', '0', 'system:menu:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '菜单管理菜单');
INSERT INTO `sys_menu` VALUES (103, '部门管理', 1, 4, '/system/dept', 'menuItem', 'C', '0', 'system:dept:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2023-09-14 16:55:37', '部门管理菜单');
INSERT INTO `sys_menu` VALUES (104, '岗位管理', 1, 5, '/system/post', 'menuItem', 'C', '0', 'system:post:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2023-09-14 16:55:51', '岗位管理菜单');
INSERT INTO `sys_menu` VALUES (105, '字典管理', 1, 6, '/system/dict', '', 'C', '0', 'system:dict:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '字典管理菜单');
INSERT INTO `sys_menu` VALUES (106, '参数设置', 1, 7, '/system/config', '', 'C', '0', 'system:config:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '参数设置菜单');
INSERT INTO `sys_menu` VALUES (109, '在线用户', 2, 1, '/monitor/online', '', 'C', '0', 'monitor:online:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '在线用户菜单');
INSERT INTO `sys_menu` VALUES (110, '定时任务', 2, 2, '/monitor/job', '', 'C', '0', 'monitor:job:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '定时任务菜单');
INSERT INTO `sys_menu` VALUES (112, '服务监控', 2, 3, '/monitor/server', '', 'C', '0', 'monitor:server:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '服务监控菜单');
INSERT INTO `sys_menu` VALUES (113, '表单构建', 3, 1, '/tool/build', '', 'C', '0', 'tool:build:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '表单构建菜单');
INSERT INTO `sys_menu` VALUES (114, '代码生成', 3, 2, '/tool/gen', '', 'C', '0', 'tool:gen:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '代码生成菜单');
INSERT INTO `sys_menu` VALUES (115, '系统接口', 3, 3, '/tool/swagger', '', 'C', '0', 'tool:swagger:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统接口菜单');
INSERT INTO `sys_menu` VALUES (500, '操作日志', 2, 4, '/monitor/operlog', '', 'C', '0', 'monitor:operlog:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '操作日志菜单');
INSERT INTO `sys_menu` VALUES (501, '登录日志', 2, 5, '/monitor/logininfor', '', 'C', '0', 'monitor:logininfor:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '登录日志菜单');
INSERT INTO `sys_menu` VALUES (1000, '用户查询', 100, 1, '#', '', 'F', '0', 'system:user:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1001, '用户新增', 100, 2, '#', '', 'F', '0', 'system:user:add', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1002, '用户修改', 100, 3, '#', '', 'F', '0', 'system:user:edit', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1003, '用户删除', 100, 4, '#', '', 'F', '0', 'system:user:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1004, '用户导出', 100, 5, '#', '', 'F', '0', 'system:user:export', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1005, '用户导入', 100, 6, '#', '', 'F', '0', 'system:user:import', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1006, '重置密码', 100, 7, '#', '', 'F', '0', 'system:user:resetPwd', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1007, '角色查询', 101, 1, '#', '', 'F', '0', 'system:role:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1008, '角色新增', 101, 2, '#', '', 'F', '0', 'system:role:add', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1009, '角色修改', 101, 3, '#', '', 'F', '0', 'system:role:edit', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1010, '角色删除', 101, 4, '#', '', 'F', '0', 'system:role:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1011, '角色导出', 101, 5, '#', '', 'F', '0', 'system:role:export', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1012, '菜单查询', 102, 1, '#', '', 'F', '0', 'system:menu:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1013, '菜单新增', 102, 2, '#', '', 'F', '0', 'system:menu:add', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1014, '菜单修改', 102, 3, '#', '', 'F', '0', 'system:menu:edit', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1015, '菜单删除', 102, 4, '#', '', 'F', '0', 'system:menu:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1016, '部门查询', 103, 1, '#', '', 'F', '0', 'system:dept:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1017, '部门新增', 103, 2, '#', '', 'F', '0', 'system:dept:add', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1018, '部门修改', 103, 3, '#', '', 'F', '0', 'system:dept:edit', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1019, '部门删除', 103, 4, '#', '', 'F', '0', 'system:dept:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1020, '岗位查询', 104, 1, '#', '', 'F', '0', 'system:post:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1021, '岗位新增', 104, 2, '#', '', 'F', '0', 'system:post:add', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1022, '岗位修改', 104, 3, '#', '', 'F', '0', 'system:post:edit', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1023, '岗位删除', 104, 4, '#', '', 'F', '0', 'system:post:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1024, '岗位导出', 104, 5, '#', '', 'F', '0', 'system:post:export', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1025, '字典查询', 105, 1, '#', '', 'F', '0', 'system:dict:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1026, '字典新增', 105, 2, '#', '', 'F', '0', 'system:dict:add', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1027, '字典修改', 105, 3, '#', '', 'F', '0', 'system:dict:edit', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1028, '字典删除', 105, 4, '#', '', 'F', '0', 'system:dict:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1029, '字典导出', 105, 5, '#', '', 'F', '0', 'system:dict:export', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1030, '参数查询', 106, 1, '#', '', 'F', '0', 'system:config:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1031, '参数新增', 106, 2, '#', '', 'F', '0', 'system:config:add', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1032, '参数修改', 106, 3, '#', '', 'F', '0', 'system:config:edit', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1033, '参数删除', 106, 4, '#', '', 'F', '0', 'system:config:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1034, '参数导出', 106, 5, '#', '', 'F', '0', 'system:config:export', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1039, '操作查询', 500, 1, '#', '', 'F', '0', 'monitor:operlog:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1040, '操作删除', 500, 2, '#', '', 'F', '0', 'monitor:operlog:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1041, '详细信息', 500, 3, '#', '', 'F', '0', 'monitor:operlog:detail', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1042, '日志导出', 500, 4, '#', '', 'F', '0', 'monitor:operlog:export', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1043, '登录查询', 501, 1, '#', '', 'F', '0', 'monitor:logininfor:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1044, '登录删除', 501, 2, '#', '', 'F', '0', 'monitor:logininfor:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1045, '日志导出', 501, 3, '#', '', 'F', '0', 'monitor:logininfor:export', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1046, '账户解锁', 501, 4, '#', '', 'F', '0', 'monitor:logininfor:unlock', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1047, '在线查询', 109, 1, '#', '', 'F', '0', 'monitor:online:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1048, '批量强退', 109, 2, '#', '', 'F', '0', 'monitor:online:batchForceLogout', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1049, '单条强退', 109, 3, '#', '', 'F', '0', 'monitor:online:forceLogout', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1050, '任务查询', 110, 1, '#', '', 'F', '0', 'monitor:job:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1051, '任务新增', 110, 2, '#', '', 'F', '0', 'monitor:job:add', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1052, '任务修改', 110, 3, '#', '', 'F', '0', 'monitor:job:edit', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1053, '任务删除', 110, 4, '#', '', 'F', '0', 'monitor:job:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1054, '状态修改', 110, 5, '#', '', 'F', '0', 'monitor:job:changeStatus', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1055, '任务详细', 110, 6, '#', '', 'F', '0', 'monitor:job:detail', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1056, '任务导出', 110, 7, '#', '', 'F', '0', 'monitor:job:export', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1057, '生成查询', 114, 1, '#', '', 'F', '0', 'tool:gen:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1058, '生成修改', 114, 2, '#', '', 'F', '0', 'tool:gen:edit', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1059, '生成删除', 114, 3, '#', '', 'F', '0', 'tool:gen:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1060, '预览代码', 114, 4, '#', '', 'F', '0', 'tool:gen:preview', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1061, '生成代码', 114, 5, '#', '', 'F', '0', 'tool:gen:code', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1062, '字典详情', 105, 5, '#', '', 'F', '0', 'system:dict:detail', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (1064, '监控平台', 4, 1, 'http://iciom.dpc.com.cn/', 'menuItem', 'C', '0', '', 'fa fa-globe', '', '2021-06-19 23:12:45', '', '2023-09-19 21:44:36', '');
INSERT INTO `sys_menu` VALUES (1065, '搜索自动补全', 6, 100, '/demo/form/autocomplete', 'menuItem', 'C', '', '', 'fa fa-sticky-note-o', '', '2023-09-18 10:15:49', '', '2023-09-18 10:17:03', '');
INSERT INTO `sys_menu` VALUES (1066, '发电量重算', 4, 100, 'iot/hisData/toWizard', 'menuItem', 'C', '', '', 'fa fa-sticky-note-o', '', '2023-09-18 21:42:44', '', '2023-09-19 21:36:52', '');
INSERT INTO `sys_menu` VALUES (1079, '基础参量管理', 4, 1, '/mywork/params', '', 'C', '0', 'mywork:params:view', 'fa fa-sticky-note-o', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '基础参量管理菜单');
INSERT INTO `sys_menu` VALUES (1080, '基础参量管理新增', 1079, 2, '#', '', 'F', '0', 'mywork:params:add', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1081, '基础参量管理查询', 1079, 1, '#', '', 'F', '0', 'mywork:params:list', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1082, '基础参量管理修改', 1079, 3, '#', '', 'F', '0', 'mywork:params:edit', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1083, '基础参量管理删除', 1079, 4, '#', '', 'F', '0', 'mywork:params:remove', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1084, '基础参量管理导出', 1079, 5, '#', '', 'F', '0', 'mywork:params:export', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');

-- ----------------------------
-- Table structure for sys_notice
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice`;
CREATE TABLE `sys_notice`  (
  `notice_id` int NOT NULL AUTO_INCREMENT COMMENT '公告ID',
  `notice_title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '公告标题',
  `notice_type` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '公告类型（1通知 2公告）',
  `notice_content` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '公告内容',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '公告状态（0正常 1关闭）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`notice_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '通知公告表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_notice
-- ----------------------------

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log`  (
  `oper_id` bigint NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '模块标题',
  `business_type` int NULL DEFAULT 0 COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '请求方式',
  `operator_type` int NULL DEFAULT 0 COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '操作人员',
  `dept_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '请求URL',
  `oper_ip` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '操作地点',
  `oper_param` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '请求参数',
  `json_result` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '返回参数',
  `status` int NULL DEFAULT 0 COMMENT '操作状态（0正常 1异常）',
  `error_msg` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '错误消息',
  `oper_time` datetime NULL DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 688 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '操作日志记录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `post_id` bigint NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '岗位名称',
  `post_sort` int NOT NULL COMMENT '显示顺序',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '状态（0正常 1停用）',
  `create_by` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint NULL DEFAULT NULL COMMENT '租户id',
  `del_flag` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '岗位信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES (1, 'manager', '经理', 1, '0', 'admin', '2018-03-16 11:33:00', '', '2023-09-15 11:35:40', 'manager', 0, 0);
INSERT INTO `sys_post` VALUES (2, 'operator', '运维', 2, '0', 'admin', '2018-03-16 11:33:00', 'admin', '2023-09-15 11:36:02', '', 0, 0);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `role_id` bigint NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色名称',
  `role_key` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色权限字符串',
  `role_sort` int NOT NULL COMMENT '显示顺序',
  `data_scope` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1' COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色状态（0正常 1停用）',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '管理员', 'admin', 1, '1', '0', '0', 'admin', '2018-03-16 11:33:00', '', '2021-06-02 00:01:45', '管理员');
INSERT INTO `sys_role` VALUES (3, '厂家运维人员', 'oper', 0, '1', '0', '0', 'admin', '2020-03-01 09:13:21', 'admin', '2023-09-15 11:03:22', '');

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept`  (
  `role_id` bigint NOT NULL COMMENT '角色ID',
  `dept_id` bigint NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`, `dept_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色和部门关联表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------
INSERT INTO `sys_role_dept` VALUES (2, 100);
INSERT INTO `sys_role_dept` VALUES (2, 110);
INSERT INTO `sys_role_dept` VALUES (3, 100);
INSERT INTO `sys_role_dept` VALUES (3, 110);
INSERT INTO `sys_role_dept` VALUES (3, 112);

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `role_id` bigint NOT NULL COMMENT '角色ID',
  `menu_id` bigint NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色和菜单关联表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES (1, 1);
INSERT INTO `sys_role_menu` VALUES (1, 2);
INSERT INTO `sys_role_menu` VALUES (1, 3);
INSERT INTO `sys_role_menu` VALUES (1, 4);
INSERT INTO `sys_role_menu` VALUES (1, 100);
INSERT INTO `sys_role_menu` VALUES (1, 101);
INSERT INTO `sys_role_menu` VALUES (1, 102);
INSERT INTO `sys_role_menu` VALUES (1, 103);
INSERT INTO `sys_role_menu` VALUES (1, 104);
INSERT INTO `sys_role_menu` VALUES (1, 105);
INSERT INTO `sys_role_menu` VALUES (1, 106);
INSERT INTO `sys_role_menu` VALUES (1, 109);
INSERT INTO `sys_role_menu` VALUES (1, 110);
INSERT INTO `sys_role_menu` VALUES (1, 112);
INSERT INTO `sys_role_menu` VALUES (1, 113);
INSERT INTO `sys_role_menu` VALUES (1, 114);
INSERT INTO `sys_role_menu` VALUES (1, 115);
INSERT INTO `sys_role_menu` VALUES (1, 500);
INSERT INTO `sys_role_menu` VALUES (1, 501);
INSERT INTO `sys_role_menu` VALUES (1, 1001);
INSERT INTO `sys_role_menu` VALUES (1, 1004);
INSERT INTO `sys_role_menu` VALUES (1, 1005);
INSERT INTO `sys_role_menu` VALUES (1, 1007);
INSERT INTO `sys_role_menu` VALUES (1, 1008);
INSERT INTO `sys_role_menu` VALUES (1, 1009);
INSERT INTO `sys_role_menu` VALUES (1, 1010);
INSERT INTO `sys_role_menu` VALUES (1, 1011);
INSERT INTO `sys_role_menu` VALUES (1, 1012);
INSERT INTO `sys_role_menu` VALUES (1, 1013);
INSERT INTO `sys_role_menu` VALUES (1, 1016);
INSERT INTO `sys_role_menu` VALUES (1, 1017);
INSERT INTO `sys_role_menu` VALUES (1, 1018);
INSERT INTO `sys_role_menu` VALUES (1, 1019);
INSERT INTO `sys_role_menu` VALUES (1, 1020);
INSERT INTO `sys_role_menu` VALUES (1, 1021);
INSERT INTO `sys_role_menu` VALUES (1, 1022);
INSERT INTO `sys_role_menu` VALUES (1, 1023);
INSERT INTO `sys_role_menu` VALUES (1, 1024);
INSERT INTO `sys_role_menu` VALUES (1, 1025);
INSERT INTO `sys_role_menu` VALUES (1, 1026);
INSERT INTO `sys_role_menu` VALUES (1, 1029);
INSERT INTO `sys_role_menu` VALUES (1, 1030);
INSERT INTO `sys_role_menu` VALUES (1, 1031);
INSERT INTO `sys_role_menu` VALUES (1, 1034);
INSERT INTO `sys_role_menu` VALUES (1, 1039);
INSERT INTO `sys_role_menu` VALUES (1, 1040);
INSERT INTO `sys_role_menu` VALUES (1, 1041);
INSERT INTO `sys_role_menu` VALUES (1, 1042);
INSERT INTO `sys_role_menu` VALUES (1, 1043);
INSERT INTO `sys_role_menu` VALUES (1, 1044);
INSERT INTO `sys_role_menu` VALUES (1, 1045);
INSERT INTO `sys_role_menu` VALUES (1, 1046);
INSERT INTO `sys_role_menu` VALUES (1, 1047);
INSERT INTO `sys_role_menu` VALUES (1, 1048);
INSERT INTO `sys_role_menu` VALUES (1, 1049);
INSERT INTO `sys_role_menu` VALUES (1, 1050);
INSERT INTO `sys_role_menu` VALUES (1, 1051);
INSERT INTO `sys_role_menu` VALUES (1, 1052);
INSERT INTO `sys_role_menu` VALUES (1, 1053);
INSERT INTO `sys_role_menu` VALUES (1, 1054);
INSERT INTO `sys_role_menu` VALUES (1, 1055);
INSERT INTO `sys_role_menu` VALUES (1, 1056);
INSERT INTO `sys_role_menu` VALUES (1, 1057);
INSERT INTO `sys_role_menu` VALUES (1, 1058);
INSERT INTO `sys_role_menu` VALUES (1, 1059);
INSERT INTO `sys_role_menu` VALUES (1, 1060);
INSERT INTO `sys_role_menu` VALUES (1, 1061);
INSERT INTO `sys_role_menu` VALUES (2, 1);
INSERT INTO `sys_role_menu` VALUES (2, 2);
INSERT INTO `sys_role_menu` VALUES (2, 3);
INSERT INTO `sys_role_menu` VALUES (2, 4);
INSERT INTO `sys_role_menu` VALUES (2, 5);
INSERT INTO `sys_role_menu` VALUES (2, 6);
INSERT INTO `sys_role_menu` VALUES (2, 7);
INSERT INTO `sys_role_menu` VALUES (2, 8);
INSERT INTO `sys_role_menu` VALUES (2, 9);
INSERT INTO `sys_role_menu` VALUES (2, 10);
INSERT INTO `sys_role_menu` VALUES (2, 11);
INSERT INTO `sys_role_menu` VALUES (2, 12);
INSERT INTO `sys_role_menu` VALUES (2, 13);
INSERT INTO `sys_role_menu` VALUES (2, 14);
INSERT INTO `sys_role_menu` VALUES (2, 15);
INSERT INTO `sys_role_menu` VALUES (2, 16);
INSERT INTO `sys_role_menu` VALUES (2, 17);
INSERT INTO `sys_role_menu` VALUES (2, 18);
INSERT INTO `sys_role_menu` VALUES (2, 19);
INSERT INTO `sys_role_menu` VALUES (2, 20);
INSERT INTO `sys_role_menu` VALUES (2, 21);
INSERT INTO `sys_role_menu` VALUES (2, 22);
INSERT INTO `sys_role_menu` VALUES (2, 23);
INSERT INTO `sys_role_menu` VALUES (2, 24);
INSERT INTO `sys_role_menu` VALUES (2, 25);
INSERT INTO `sys_role_menu` VALUES (2, 26);
INSERT INTO `sys_role_menu` VALUES (2, 27);
INSERT INTO `sys_role_menu` VALUES (2, 28);
INSERT INTO `sys_role_menu` VALUES (2, 29);
INSERT INTO `sys_role_menu` VALUES (2, 30);
INSERT INTO `sys_role_menu` VALUES (2, 31);
INSERT INTO `sys_role_menu` VALUES (2, 32);
INSERT INTO `sys_role_menu` VALUES (2, 33);
INSERT INTO `sys_role_menu` VALUES (2, 34);
INSERT INTO `sys_role_menu` VALUES (2, 35);
INSERT INTO `sys_role_menu` VALUES (2, 36);
INSERT INTO `sys_role_menu` VALUES (2, 37);
INSERT INTO `sys_role_menu` VALUES (2, 38);
INSERT INTO `sys_role_menu` VALUES (2, 39);
INSERT INTO `sys_role_menu` VALUES (2, 40);
INSERT INTO `sys_role_menu` VALUES (2, 41);
INSERT INTO `sys_role_menu` VALUES (2, 42);
INSERT INTO `sys_role_menu` VALUES (2, 43);
INSERT INTO `sys_role_menu` VALUES (2, 44);
INSERT INTO `sys_role_menu` VALUES (2, 45);
INSERT INTO `sys_role_menu` VALUES (2, 46);
INSERT INTO `sys_role_menu` VALUES (2, 47);
INSERT INTO `sys_role_menu` VALUES (2, 48);
INSERT INTO `sys_role_menu` VALUES (2, 49);
INSERT INTO `sys_role_menu` VALUES (2, 50);
INSERT INTO `sys_role_menu` VALUES (2, 51);
INSERT INTO `sys_role_menu` VALUES (2, 52);
INSERT INTO `sys_role_menu` VALUES (2, 53);
INSERT INTO `sys_role_menu` VALUES (2, 54);
INSERT INTO `sys_role_menu` VALUES (2, 55);
INSERT INTO `sys_role_menu` VALUES (2, 100);
INSERT INTO `sys_role_menu` VALUES (2, 101);
INSERT INTO `sys_role_menu` VALUES (2, 102);
INSERT INTO `sys_role_menu` VALUES (2, 103);
INSERT INTO `sys_role_menu` VALUES (2, 104);
INSERT INTO `sys_role_menu` VALUES (2, 105);
INSERT INTO `sys_role_menu` VALUES (2, 106);
INSERT INTO `sys_role_menu` VALUES (2, 107);
INSERT INTO `sys_role_menu` VALUES (2, 109);
INSERT INTO `sys_role_menu` VALUES (2, 110);
INSERT INTO `sys_role_menu` VALUES (2, 112);
INSERT INTO `sys_role_menu` VALUES (2, 113);
INSERT INTO `sys_role_menu` VALUES (2, 114);
INSERT INTO `sys_role_menu` VALUES (2, 115);
INSERT INTO `sys_role_menu` VALUES (2, 500);
INSERT INTO `sys_role_menu` VALUES (2, 501);
INSERT INTO `sys_role_menu` VALUES (2, 1001);
INSERT INTO `sys_role_menu` VALUES (2, 1002);
INSERT INTO `sys_role_menu` VALUES (2, 1003);
INSERT INTO `sys_role_menu` VALUES (2, 1004);
INSERT INTO `sys_role_menu` VALUES (2, 1005);
INSERT INTO `sys_role_menu` VALUES (2, 1006);
INSERT INTO `sys_role_menu` VALUES (2, 1007);
INSERT INTO `sys_role_menu` VALUES (2, 1008);
INSERT INTO `sys_role_menu` VALUES (2, 1009);
INSERT INTO `sys_role_menu` VALUES (2, 1010);
INSERT INTO `sys_role_menu` VALUES (2, 1011);
INSERT INTO `sys_role_menu` VALUES (2, 1012);
INSERT INTO `sys_role_menu` VALUES (2, 1013);
INSERT INTO `sys_role_menu` VALUES (2, 1014);
INSERT INTO `sys_role_menu` VALUES (2, 1015);
INSERT INTO `sys_role_menu` VALUES (2, 1016);
INSERT INTO `sys_role_menu` VALUES (2, 1017);
INSERT INTO `sys_role_menu` VALUES (2, 1018);
INSERT INTO `sys_role_menu` VALUES (2, 1019);
INSERT INTO `sys_role_menu` VALUES (2, 1020);
INSERT INTO `sys_role_menu` VALUES (2, 1021);
INSERT INTO `sys_role_menu` VALUES (2, 1022);
INSERT INTO `sys_role_menu` VALUES (2, 1023);
INSERT INTO `sys_role_menu` VALUES (2, 1024);
INSERT INTO `sys_role_menu` VALUES (2, 1025);
INSERT INTO `sys_role_menu` VALUES (2, 1026);
INSERT INTO `sys_role_menu` VALUES (2, 1027);
INSERT INTO `sys_role_menu` VALUES (2, 1028);
INSERT INTO `sys_role_menu` VALUES (2, 1029);
INSERT INTO `sys_role_menu` VALUES (2, 1030);
INSERT INTO `sys_role_menu` VALUES (2, 1031);
INSERT INTO `sys_role_menu` VALUES (2, 1032);
INSERT INTO `sys_role_menu` VALUES (2, 1033);
INSERT INTO `sys_role_menu` VALUES (2, 1034);
INSERT INTO `sys_role_menu` VALUES (2, 1035);
INSERT INTO `sys_role_menu` VALUES (2, 1036);
INSERT INTO `sys_role_menu` VALUES (2, 1037);
INSERT INTO `sys_role_menu` VALUES (2, 1038);
INSERT INTO `sys_role_menu` VALUES (2, 1039);
INSERT INTO `sys_role_menu` VALUES (2, 1040);
INSERT INTO `sys_role_menu` VALUES (2, 1041);
INSERT INTO `sys_role_menu` VALUES (2, 1042);
INSERT INTO `sys_role_menu` VALUES (2, 1043);
INSERT INTO `sys_role_menu` VALUES (2, 1044);
INSERT INTO `sys_role_menu` VALUES (2, 1045);
INSERT INTO `sys_role_menu` VALUES (2, 1046);
INSERT INTO `sys_role_menu` VALUES (2, 1047);
INSERT INTO `sys_role_menu` VALUES (2, 1048);
INSERT INTO `sys_role_menu` VALUES (2, 1049);
INSERT INTO `sys_role_menu` VALUES (2, 1050);
INSERT INTO `sys_role_menu` VALUES (2, 1051);
INSERT INTO `sys_role_menu` VALUES (2, 1052);
INSERT INTO `sys_role_menu` VALUES (2, 1053);
INSERT INTO `sys_role_menu` VALUES (2, 1054);
INSERT INTO `sys_role_menu` VALUES (2, 1055);
INSERT INTO `sys_role_menu` VALUES (2, 1056);
INSERT INTO `sys_role_menu` VALUES (2, 1057);
INSERT INTO `sys_role_menu` VALUES (2, 1058);
INSERT INTO `sys_role_menu` VALUES (2, 1059);
INSERT INTO `sys_role_menu` VALUES (2, 1060);
INSERT INTO `sys_role_menu` VALUES (2, 1061);
INSERT INTO `sys_role_menu` VALUES (3, 4);
INSERT INTO `sys_role_menu` VALUES (3, 1064);

-- ----------------------------
-- Table structure for sys_tenant
-- ----------------------------
DROP TABLE IF EXISTS `sys_tenant`;
CREATE TABLE `sys_tenant`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建者',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '商户名称',
  `address` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系地址',
  `manager` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '负责人',
  `phone` varchar(18) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  `start_time` datetime NULL DEFAULT NULL COMMENT '起租时间',
  `end_time` datetime NULL DEFAULT NULL COMMENT '结束时间',
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '安全邮箱',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_tenant
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `user_id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `dept_id` bigint NULL DEFAULT NULL COMMENT '部门ID',
  `login_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '登录账号',
  `user_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户昵称',
  `user_type` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '00' COMMENT '用户类型（00系统用户）',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '用户邮箱',
  `phonenumber` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '手机号码',
  `sex` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '用户性别（0男 1女 2未知）',
  `avatar` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '头像路径',
  `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '盐加密',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '帐号状态（0正常 1停用）',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `login_ip` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '最后登录IP',
  `login_date` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint NULL DEFAULT 0 COMMENT '租户id',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 110, 'admin', '超级管理员', '00', 'dd122221111d@163.com', '15888888881', '0', '/upload/admin1694703638771106900.png', '609ac514e6ef0b9a5f4eee66693fd7f8', 'NcSB3H', '0', '0', '127.0.0.1', '2020-01-13 13:20:40', 'admin', '2018-03-16 11:33:00', 'admin', '2021-06-20 13:52:46', '管理员111111', 0);
INSERT INTO `sys_user` VALUES (2, 110, 'qy1156010001', 'qy1156010001', '', 'uhvs@163.com', '18518276786', '0', '', '49efab53d90058dd6086aed4158a5bbe', 'YYYYYd', '0', '0', '', NULL, 'admin', '2023-09-15 11:01:46', 'admin', '2023-09-19 22:01:29', '', 0);

-- ----------------------------
-- Table structure for sys_user_online
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_online`;
CREATE TABLE `sys_user_online`  (
  `session_id` varchar(250) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户会话id',
  `login_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登录账号',
  `dept_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `ipaddr` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '操作系统',
  `status` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '在线状态on_line在线off_line离线',
  `start_timestamp` datetime NULL DEFAULT NULL COMMENT 'session创建时间',
  `last_access_time` datetime NULL DEFAULT NULL COMMENT 'session最后访问时间',
  `expire_time` int NULL DEFAULT 0 COMMENT '超时时间，单位为分钟',
  PRIMARY KEY (`session_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '在线用户记录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user_online
-- ----------------------------
INSERT INTO `sys_user_online` VALUES ('6de1f46577c2c02bb92a3cbeecf8718a', '超级管理员', '', '192.168.240.162:49630', '', 'Chrome', 'Windows 10', 'on_line', '2024-05-15 08:35:24', '2024-05-15 08:35:24', 1440);
INSERT INTO `sys_user_online` VALUES ('7IHPSW3T77IZ5EGASLEYR2TGROQCIWQQK2CPEDOLP4CQEPTL3QGQ', 'admin', '', '::1', '内网IP', 'Chrome', 'Windows 10', 'on_line', '2021-06-20 15:59:42', '2021-06-20 15:59:42', 1440);

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post`  (
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `post_id` bigint NOT NULL COMMENT '岗位ID',
  `del_flag` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
  PRIMARY KEY (`user_id`, `post_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户与岗位关联表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `role_id` bigint NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户和角色关联表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES (2, 3);

SET FOREIGN_KEY_CHECKS = 1;
