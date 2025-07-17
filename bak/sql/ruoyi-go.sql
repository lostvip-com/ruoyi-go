/*
 Navicat Premium Data Transfer

 Source Server         : lv
 Source Server Type    : MySQL
 Source Server Version : 50744
 Source Host           : 47.242.156.150:3307
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 50744
 File Encoding         : 65001

 Date: 17/07/2025 13:42:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for advance_config
-- ----------------------------
DROP TABLE IF EXISTS `advance_config`;
CREATE TABLE `advance_config`  (
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `log_level` int(11) NULL DEFAULT 1 COMMENT '日志等级',
  `persist_storage` tinyint(1) NULL DEFAULT 0 COMMENT '存储开关',
  `storage_hour` int(11) NULL DEFAULT 24 COMMENT '存储时长',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of advance_config
-- ----------------------------

-- ----------------------------
-- Table structure for alert_list
-- ----------------------------
DROP TABLE IF EXISTS `alert_list`;
CREATE TABLE `alert_list`  (
  `created` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `modified` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `alert_rule_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '告警记录ID',
  `trigger_time` bigint(20) NULL DEFAULT NULL COMMENT '触发时间',
  `alert_result` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '告警内容',
  `status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态',
  `treated_time` bigint(20) NULL DEFAULT NULL COMMENT '处理时间',
  `message` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '处理意见',
  `is_send` tinyint(1) NULL DEFAULT NULL COMMENT '是否发送通知',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_alert_list_alert_rule`(`alert_rule_id`) USING BTREE,
  CONSTRAINT `fk_alert_list_alert_rule` FOREIGN KEY (`alert_rule_id`) REFERENCES `alert_rule` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of alert_list
-- ----------------------------

-- ----------------------------
-- Table structure for alert_rule
-- ----------------------------
DROP TABLE IF EXISTS `alert_rule`;
CREATE TABLE `alert_rule`  (
  `created` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `modified` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `device_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `alert_type` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `alert_level` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `status` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `condition` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `sub_rule` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `notify` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `silence_time` bigint(20) NULL DEFAULT NULL,
  `description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_alert_rule_device_id`(`device_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of alert_rule
-- ----------------------------
INSERT INTO `alert_rule` VALUES (1720828131332, 1720828131332, '15030751', 'victoriametrics', '', '设备告警', '次要', '', '', '', '', 0, '');

-- ----------------------------
-- Table structure for category_template
-- ----------------------------
DROP TABLE IF EXISTS `category_template`;
CREATE TABLE `category_template`  (
  `created` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `modified` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `category_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '品类标识',
  `category_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '品类名字',
  `scene` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '场景',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of category_template
-- ----------------------------

-- ----------------------------
-- Table structure for cms_online
-- ----------------------------
DROP TABLE IF EXISTS `cms_online`;
CREATE TABLE `cms_online`  (
  `session_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `login_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登录名',
  `dept_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登录名',
  `ipaddr` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `login_location` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `browser` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `os` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `start_timestamp` datetime(0) NULL DEFAULT NULL COMMENT '开始时间',
  `last_access_time` datetime(0) NULL DEFAULT NULL COMMENT '上次访问',
  `expire_time` bigint(20) NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`session_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cms_online
-- ----------------------------

-- ----------------------------
-- Table structure for device_library
-- ----------------------------
DROP TABLE IF EXISTS `device_library`;
CREATE TABLE `device_library`  (
  `created` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `modified` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名字',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '描述',
  `protocol` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '驱动协议',
  `version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '当前安装版本',
  `container_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '容器名字',
  `docker_config_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '镜像仓库配置表id',
  `docker_repo_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '镜像名称',
  `docker_image_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '镜像ID',
  `support_versions` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '可用版本',
  `is_internal` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否内置，云端内置驱动',
  `language` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '代码语言',
  `manual` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '驱动市场使用说明手册',
  `icon` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '图标',
  `classify_id` bigint(20) NULL DEFAULT NULL COMMENT '分类',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of device_library
-- ----------------------------

-- ----------------------------
-- Table structure for device_service
-- ----------------------------
DROP TABLE IF EXISTS `device_service`;
CREATE TABLE `device_service`  (
  `created` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `modified` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名字',
  `base_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址',
  `device_library_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '驱动ID',
  `config` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '配置',
  `docker_container_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'docker容器ID',
  `expert_mode` tinyint(1) NULL DEFAULT NULL COMMENT '扩展模式',
  `expert_mode_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '扩展内容',
  `docker_params_switch` tinyint(1) NULL DEFAULT NULL COMMENT 'docker启动参数开关',
  `docker_params` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'docker启动参数',
  `container_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '容器名字',
  `log_level` int(11) NULL DEFAULT 1 COMMENT '日志等级',
  `driver_type` bigint(20) NOT NULL DEFAULT 1 COMMENT '驱动类别，1：驱动，2：三方应用',
  `platform` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_device_service_device_library_id`(`device_library_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of device_service
-- ----------------------------

-- ----------------------------
-- Table structure for doc
-- ----------------------------
DROP TABLE IF EXISTS `doc`;
CREATE TABLE `doc`  (
  `created` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `modified` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名字',
  `sort` tinyint(4) NULL DEFAULT NULL COMMENT '排序',
  `jump_link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '跳转地址',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of doc
-- ----------------------------

-- ----------------------------
-- Table structure for docker_config
-- ----------------------------
DROP TABLE IF EXISTS `docker_config`;
CREATE TABLE `docker_config`  (
  `created` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `modified` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址',
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `salt_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '盐值',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of docker_config
-- ----------------------------

-- ----------------------------
-- Table structure for dpc_task
-- ----------------------------
DROP TABLE IF EXISTS `dpc_task`;
CREATE TABLE `dpc_task`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '工号',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `prj_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '项  目  号',
  `task_content` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '任务内容',
  `start_date` date NULL DEFAULT NULL COMMENT '开始日期',
  `end_date` date NULL DEFAULT NULL COMMENT '结束日期',
  `work_days` bigint(20) NULL DEFAULT NULL COMMENT '本月工时',
  `auto_submit` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '自动提交',
  `status` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '任务状态',
  `sort` bigint(20) NULL DEFAULT NULL COMMENT '排序，大的优先',
  `update_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `create_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `del_flag` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dpc_task
-- ----------------------------

-- ----------------------------
-- Table structure for driver_classify
-- ----------------------------
DROP TABLE IF EXISTS `driver_classify`;
CREATE TABLE `driver_classify`  (
  `created` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `modified` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名字',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of driver_classify
-- ----------------------------

-- ----------------------------
-- Table structure for gen_table
-- ----------------------------
DROP TABLE IF EXISTS `gen_table`;
CREATE TABLE `gen_table`  (
  `table_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
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
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`table_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '代码生成业务表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_table
-- ----------------------------
INSERT INTO `gen_table` VALUES (24, 'his_patient', '患者基本信息', 'HisPatient', 'crud', 'robvi', 'biz', 'patient', '患者基本信息', 'lv', '', 'admin', '2023-11-26 16:09:17', 'admin', '2023-11-26 16:12:38', '');
INSERT INTO `gen_table` VALUES (25, 'dpc_task', '计划任务', 'DpcTask', 'crud', 'robvi', 'biz', 'task', '计划任务', 'lv', '', 'admin', '2023-11-26 16:27:18', 'admin', '2023-11-27 23:01:36', '');

-- ----------------------------
-- Table structure for gen_table_column
-- ----------------------------
DROP TABLE IF EXISTS `gen_table_column`;
CREATE TABLE `gen_table_column`  (
  `column_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_id` bigint(20) NULL DEFAULT NULL COMMENT '归属表编号',
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
  `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`column_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 882 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '代码生成业务表字段' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_table_column
-- ----------------------------
INSERT INTO `gen_table_column` VALUES (539, 12, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (540, 12, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (541, 12, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (542, 12, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (543, 12, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (544, 12, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (545, 12, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (546, 12, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (547, 12, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (548, 12, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (549, 12, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (550, 12, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (551, 12, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (552, 12, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (553, 12, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (554, 12, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (555, 12, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (556, 12, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (557, 12, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (558, 12, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (559, 12, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (560, 12, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (561, 12, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (562, 12, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (563, 12, 'create_time', '创建时间', 'date', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (564, 12, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (565, 12, 'update_time', '更新时间', 'date', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (566, 12, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (567, 12, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (568, 13, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (569, 13, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (570, 13, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (571, 13, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (572, 13, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (573, 13, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (574, 13, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (575, 13, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (576, 13, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (577, 13, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (578, 13, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (579, 13, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (580, 13, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'select', 'sys_user_sex', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (581, 13, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (582, 13, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (583, 13, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (584, 13, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (585, 13, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (586, 13, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (587, 13, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (588, 13, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (589, 13, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '0', '0', 'LIKE', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (590, 13, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (591, 13, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (592, 13, 'create_time', '创建时间', 'date', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (593, 13, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (594, 13, 'update_time', '更新时间', 'date', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (595, 13, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (596, 13, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '0', '1', 'EQ', 'select', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (597, 14, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (598, 14, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (599, 14, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (600, 14, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (601, 14, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (602, 14, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (603, 14, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (604, 14, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (605, 14, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (606, 14, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (607, 14, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (608, 14, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (609, 14, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (610, 14, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (611, 14, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (612, 14, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (613, 14, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (614, 14, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (615, 14, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (616, 14, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (617, 14, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (618, 14, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (619, 14, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (620, 14, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (621, 14, 'create_time', '创建时间', 'date', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (622, 14, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (623, 14, 'update_time', '更新时间', 'date', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (624, 14, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (625, 14, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (626, 15, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (627, 15, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (628, 15, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (629, 15, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (630, 15, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (631, 15, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (632, 15, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (633, 15, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (634, 15, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (635, 15, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (636, 15, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (637, 15, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (638, 15, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (639, 15, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (640, 15, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (641, 15, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (642, 15, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (643, 15, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (644, 15, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (645, 15, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (646, 15, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (647, 15, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (648, 15, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (649, 15, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (650, 15, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (651, 15, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (652, 15, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (653, 15, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (654, 15, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (655, 16, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (656, 16, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (657, 16, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (658, 16, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (659, 16, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (660, 16, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (661, 16, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (662, 16, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (663, 16, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (664, 16, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (665, 16, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (666, 16, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (667, 16, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (668, 16, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (669, 16, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (670, 16, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (671, 16, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (672, 16, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (673, 16, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (674, 16, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (675, 16, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (676, 16, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (677, 16, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (678, 16, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (679, 16, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (680, 16, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (681, 16, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (682, 16, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (683, 16, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (684, 17, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (685, 17, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (686, 17, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (687, 17, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (688, 17, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (689, 17, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (690, 17, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (691, 17, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (692, 17, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (693, 17, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (694, 17, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (695, 17, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (696, 17, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (697, 17, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (698, 17, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (699, 17, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (700, 17, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (701, 17, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (702, 17, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (703, 17, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (704, 17, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (705, 17, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (706, 17, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (707, 17, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (708, 17, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (709, 17, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (710, 17, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (711, 17, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (712, 17, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (713, 18, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (714, 18, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (715, 18, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (716, 18, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (717, 18, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (718, 18, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (719, 18, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (720, 18, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (721, 18, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (722, 18, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (723, 18, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (724, 18, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (725, 18, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (726, 18, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (727, 18, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (728, 18, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (729, 18, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (730, 18, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (731, 18, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (732, 18, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (733, 18, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (734, 18, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '0', 'EQ', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (735, 18, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (736, 18, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (737, 18, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (738, 18, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (739, 18, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (740, 18, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (741, 18, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (742, 19, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (743, 19, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (744, 19, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (745, 19, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (746, 19, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (747, 19, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (748, 19, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (749, 19, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (750, 19, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (751, 19, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (752, 19, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (753, 19, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (754, 20, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (755, 19, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (756, 20, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (757, 19, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (758, 20, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (759, 19, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (760, 20, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (761, 19, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (762, 19, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (763, 20, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (764, 19, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (765, 20, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (766, 19, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (767, 20, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (768, 19, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (769, 20, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (770, 19, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (771, 20, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (772, 19, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '0', 'EQ', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (773, 20, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (774, 19, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (775, 19, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (776, 20, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (777, 19, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (778, 20, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (779, 19, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (780, 20, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (781, 19, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (782, 20, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (783, 19, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (784, 20, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (785, 19, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (786, 20, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (787, 20, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (788, 20, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (789, 20, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (790, 20, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (791, 20, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (792, 20, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '0', 'EQ', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (793, 20, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (794, 20, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (795, 20, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (796, 20, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (797, 20, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (798, 20, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (799, 20, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (801, 22, 'id', '', 'int(11)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (802, 22, 'username', '', 'varchar(255)', 'string', 'Username', 'username', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (803, 22, 'password', '', 'varchar(255)', 'string', 'Password', 'password', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (804, 22, 'taskContent', '', 'varchar(255)', 'string', 'TaskContent', 'taskContent', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (805, 22, 'startDate', '', 'date', 'Time', 'StartDate', 'startDate', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (806, 22, 'endDate', '', 'date', 'Time', 'EndDate', 'endDate', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'datetime', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (807, 22, 'workDays', '', 'int(11)', 'int64', 'WorkDays', 'workDays', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (808, 22, 'autoSubmit', '', 'tinyint(1)', 'int', 'AutoSubmit', 'autoSubmit', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (809, 22, 'status', '', 'varchar(255)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'radio', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (810, 22, 'sort', '', 'int(11)', 'int64', 'Sort', 'sort', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (811, 22, 'updateTime', '', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'datetime', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (812, 23, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (813, 23, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (814, 23, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (815, 23, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (816, 23, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (817, 23, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (818, 23, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (819, 23, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (820, 23, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (821, 23, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (822, 23, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (823, 23, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (824, 23, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (825, 23, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (826, 23, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (827, 23, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (828, 23, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (829, 23, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (830, 23, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (831, 23, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (832, 23, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (833, 23, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (834, 23, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'LIKE', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (835, 23, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (836, 23, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (837, 23, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (838, 23, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (839, 23, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (840, 23, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (841, 24, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (842, 24, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (843, 24, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (844, 24, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (845, 24, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (846, 24, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (847, 24, 'idcard_path', '身份证照片', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (848, 24, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (849, 24, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (850, 24, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (851, 24, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (852, 24, 'family_id', '家庭ID', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (853, 24, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'select', 'sys_user_sex', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (854, 24, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (855, 24, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (856, 24, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (857, 24, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (858, 24, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (859, 24, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (860, 24, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (861, 24, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (862, 24, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (863, 24, 'del_flag', '删除标识1删除0未删除', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'LIKE', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (864, 24, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '1', '0', '0', '0', 'LIKE', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (865, 24, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (866, 24, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '1', '1', '0', '0', 'LIKE', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (867, 24, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (868, 24, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (869, 25, 'id', 'ID', 'int(11)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (870, 25, 'username', '工号', 'varchar(16)', 'string', 'Username', 'username', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (871, 25, 'password', '密码', 'varchar(32)', 'string', 'Password', 'password', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (872, 25, 'prj_code', '项  目  号', 'varchar(16)', 'string', 'PrjCode', 'prjCode', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (873, 25, 'task_content', '任务内容', 'varchar(64)', 'string', 'TaskContent', 'taskContent', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (874, 25, 'start_date', '开始日期', 'date', 'Time', 'StartDate', 'startDate', '0', '0', '0', '1', '1', '1', '1', 'GTE', 'datetime', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (875, 25, 'end_date', '结束日期', 'date', 'Time', 'EndDate', 'endDate', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'datetime', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (876, 25, 'work_days', '本月工时', 'int(11)', 'int64', 'WorkDays', 'workDays', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (877, 25, 'auto_submit', '自动提交', 'char(1)', 'string', 'AutoSubmit', 'autoSubmit', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (878, 25, 'status', '任务状态，ready running end', 'varchar(16)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'radio', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (879, 25, 'sort', '排序，大的优先', 'int(11)', 'int64', 'Sort', 'sort', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (880, 25, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (881, 25, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 13, 'admin', NULL, '', NULL);

-- ----------------------------
-- Table structure for his_patient
-- ----------------------------
DROP TABLE IF EXISTS `his_patient`;
CREATE TABLE `his_patient`  (
  `id` bigint(32) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '姓名',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `doctor_id` bigint(20) NULL DEFAULT NULL COMMENT '责任医生Id',
  `idcard` varchar(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '证件号',
  `head_url` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '照片',
  `idcard_path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '身份证照片',
  `bed_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '床号',
  `org_id` bigint(20) NULL DEFAULT NULL COMMENT '建档单位',
  `org_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '建档单位地址',
  `org_establish` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '建档单位',
  `family_id` bigint(32) NULL DEFAULT NULL COMMENT '家庭ID',
  `sex` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '性别',
  `birth` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '生日',
  `weight` float(10, 0) NULL DEFAULT NULL COMMENT '体重',
  `height` float(10, 0) NULL DEFAULT NULL COMMENT '身高',
  `nation` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '民族',
  `native_place` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '籍贯',
  `address` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '现居地址',
  `occupation` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '职业',
  `contactor_phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '联系人手机号',
  `contactor_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '联系人',
  `del_flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '删除标识1删除0未删除',
  `create_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '患者基本信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of his_patient
-- ----------------------------
INSERT INTO `his_patient` VALUES (1, 'test', '17600196256', 1, '63280120170516003X', NULL, NULL, '1', NULL, '的撒范德萨春树暮云大本营', '的撒范德萨', NULL, '0', '2023-01-01', 12, 2, '汉', '的撒法第三', '浦东', '春树暮云', '17600196256', '121', '0', NULL, '2023-08-20 00:00:00', NULL, '2023-08-20 00:00:00', '夺花');
INSERT INTO `his_patient` VALUES (2, '秀梅', '17600196256', 1, '63280120170516003X', NULL, NULL, '1', NULL, '', '', NULL, '1', '2017-05-16', 12, 1, '汉', '1212121', '浦东', '12121', '', '', '0', NULL, '2023-08-20 00:00:00', NULL, NULL, 'fdsfds');
INSERT INTO `his_patient` VALUES (3, '运营成本', '17600196256', 1, '63280120170516003X', NULL, NULL, '1', 103, '', '研发部门', NULL, NULL, '2017-05-16', 21, 21, '汉', '青海', '浦东', '', '', '', '0', 'admin', '2023-08-20 00:00:00', NULL, NULL, '士大夫噶但是');
INSERT INTO `his_patient` VALUES (4, '花木成畦手自栽大本营', '17600196256', 1, '63280120170526003X', NULL, NULL, '1', 103, '', '研发部门', NULL, NULL, '2017-05-16', 11, 21, '21', '21', '浦东', '121', '', '', '0', 'admin', '2023-08-20 00:00:00', NULL, NULL, '');

-- ----------------------------
-- Table structure for iot_alert_ai_di
-- ----------------------------
DROP TABLE IF EXISTS `iot_alert_ai_di`;
CREATE TABLE `iot_alert_ai_di`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `dev_id` bigint(20) NULL DEFAULT NULL COMMENT '设备ID',
  `code` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '告警唯一识别号',
  `type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '告警类型',
  `level` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '告警级别',
  `content` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '告警内容',
  `message` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '处理意见',
  `trigger_time` datetime(0) NULL DEFAULT NULL COMMENT '触发时间',
  `recover_time` datetime(0) NULL DEFAULT NULL COMMENT '恢复时间',
  `treated_time` datetime(0) NULL DEFAULT NULL COMMENT '处理时间',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态',
  `del_flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '删除标记',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建者',
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iot_alert_ai_di
-- ----------------------------

-- ----------------------------
-- Table structure for iot_cmd_ao_do
-- ----------------------------
DROP TABLE IF EXISTS `iot_cmd_ao_do`;
CREATE TABLE `iot_cmd_ao_do`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `dev_id` bigint(20) NULL DEFAULT NULL COMMENT '设备ID',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '属性名',
  `prop_tags` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'TAG',
  `prop_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '属性名',
  `prop_value` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '采集值',
  `dev_time` datetime(0) NULL DEFAULT NULL COMMENT '采集时间',
  `del_flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '删除标记',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建者',
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iot_cmd_ao_do
-- ----------------------------

-- ----------------------------
-- Table structure for iot_data_his
-- ----------------------------
DROP TABLE IF EXISTS `iot_data_his`;
CREATE TABLE `iot_data_his`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `dev_id` bigint(20) NULL DEFAULT NULL COMMENT '设备ID',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '属性名',
  `prop_tags` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'TAG',
  `prop_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '属性名',
  `prop_value` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '采集值',
  `dev_time` datetime(0) NULL DEFAULT NULL COMMENT '采集时间',
  `del_flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '删除标记',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建者',
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iot_data_his
-- ----------------------------

-- ----------------------------
-- Table structure for iot_data_realtime
-- ----------------------------
DROP TABLE IF EXISTS `iot_data_realtime`;
CREATE TABLE `iot_data_realtime`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `dev_id` bigint(20) NULL DEFAULT NULL COMMENT '设备ID',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '属性名',
  `prop_tags` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'TAG',
  `prop_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '属性名',
  `prop_value` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '采集值',
  `dev_time` datetime(0) NULL DEFAULT NULL COMMENT '采集时间',
  `del_flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '删除标记',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建者',
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iot_data_realtime
-- ----------------------------
INSERT INTO `iot_data_realtime` VALUES (1, 1, '1', '1', 'pv', '12', '2024-08-30 10:01:05', '0', '2024-08-30 10:01:09', '2024-08-30 10:01:12', 'test', 'test', 0);

-- ----------------------------
-- Table structure for iot_device
-- ----------------------------
DROP TABLE IF EXISTS `iot_device`;
CREATE TABLE `iot_device`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `dept_id` bigint(20) NULL DEFAULT NULL COMMENT '上级单位ID',
  `sn` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '序列号',
  `dev_no` smallint(6) NULL DEFAULT NULL COMMENT '从机地址',
  `product_id` bigint(20) NULL DEFAULT NULL COMMENT '产品ID',
  `gateway_id` bigint(20) NULL DEFAULT NULL COMMENT '所属网关ID，只对网关子设备有效',
  `drive_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '驱动实例ID',
  `node_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '节点类型3设备2网关',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名字',
  `status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '设备状态',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '描述',
  `secret` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密钥',
  `platform` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '平台名称',
  `install_location` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '安装地址',
  `last_sync_time` datetime(0) NULL DEFAULT NULL COMMENT '最后一次同步时间',
  `last_online_time` datetime(0) NULL DEFAULT NULL COMMENT '最后一次在线时间',
  `del_flag` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建者',
  `drive_instance_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '驱动实例ID',
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  `dept_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `gateway_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iot_device
-- ----------------------------
INSERT INTO `iot_device` VALUES (1, 110, 'SN20240720', 1, 1, 0, '', '12', '1', '0', '1', '1', NULL, '1', NULL, NULL, 1, '2024-07-20 14:45:56', '2024-07-20 14:45:56', '', 'admin', NULL, NULL, NULL, NULL);
INSERT INTO `iot_device` VALUES (2, 110, 'SN20240710', 1, 1, 0, '', '12', '1', '0', '1', '1', NULL, '1', NULL, NULL, 1, '2024-07-20 14:47:59', '2024-07-20 14:47:59', '', 'admin', NULL, NULL, NULL, NULL);
INSERT INTO `iot_device` VALUES (3, 110, 'SN20240721', 1, 1, 0, '', '10', '设备网关1', '0', '网关1', '123', NULL, '动力源老楼', NULL, NULL, 0, '2024-07-29 12:46:16', '2024-07-29 12:46:16', 'admin', 'admin', NULL, NULL, NULL, NULL);
INSERT INTO `iot_device` VALUES (4, 110, 'SN20240711', 1, 2, 3, '', '12', '逆变器-sn123456', '0', '网关子设备1', '123456', NULL, '动力源老楼1', NULL, NULL, 0, '2024-07-25 13:42:24', '2024-07-23 01:17:02', 'admin', 'admin', NULL, NULL, NULL, NULL);
INSERT INTO `iot_device` VALUES (5, 114, 'SN20240712', 1, 2, 3, '', '12', '2', '0', 'test12123', '1', NULL, '1', NULL, NULL, 0, '2024-07-31 08:45:20', '2024-07-31 08:45:10', 'admin', 'admin', NULL, NULL, NULL, NULL);
INSERT INTO `iot_device` VALUES (6, 114, 'SN20240713', 1, 3, 0, '', '11', '2', '0', '直连设备', '1', NULL, '1', NULL, NULL, 0, '2024-07-31 09:26:38', '2024-07-31 09:26:38', 'admin', 'admin', NULL, NULL, NULL, NULL);
INSERT INTO `iot_device` VALUES (7, 1114, 'SN20240714', 1, 2, 0, '', '12', '2', '0', '逆变器2', '2', NULL, '动力源老楼1', NULL, NULL, 0, '2024-07-23 02:00:52', '2024-07-23 02:00:52', 'admin', 'admin', NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for iot_prd_action
-- ----------------------------
DROP TABLE IF EXISTS `iot_prd_action`;
CREATE TABLE `iot_prd_action`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `product_id` int(11) NULL DEFAULT NULL COMMENT '产品ID',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标识符',
  `tag` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名字',
  `call_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '调用方式',
  `input_params` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '输入参数',
  `output_params` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '输入参数',
  `remark` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `del_flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '删除标记',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建者',
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iot_prd_action
-- ----------------------------
INSERT INTO `iot_prd_action` VALUES (1, 0, 'param_xml_imp', '3200', '下发参量表', '1', '{}', '{}', '', '0', '2024-08-02 05:03:24', '2024-08-02 05:03:23', 'admin', 'admin', 0);

-- ----------------------------
-- Table structure for iot_prd_event
-- ----------------------------
DROP TABLE IF EXISTS `iot_prd_event`;
CREATE TABLE `iot_prd_event`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `product_id` int(11) NULL DEFAULT NULL COMMENT '产品ID',
  `event_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '事件类型',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标识符',
  `tag` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名字',
  `remark` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '描述',
  `output_params` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '输入参数',
  `del_flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '删除标记',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建者',
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iot_prd_event
-- ----------------------------
INSERT INTO `iot_prd_event` VALUES (1, 0, '1', 'alert_temp', '1001', '温度超上限告警', 'test', '{}', '0', '2024-08-02 05:03:08', '2024-08-02 05:03:08', 'admin', 'admin', 0);
INSERT INTO `iot_prd_event` VALUES (2, 0, '0', 'report_aidi', '3600', '全量上报属性值', '定时全量上属性信息', '', '0', '2024-08-02 05:07:24', '2024-08-02 05:07:24', 'admin', 'admin', 0);

-- ----------------------------
-- Table structure for iot_prd_property
-- ----------------------------
DROP TABLE IF EXISTS `iot_prd_property`;
CREATE TABLE `iot_prd_property`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `product_id` bigint(20) NULL DEFAULT NULL COMMENT '产品ID',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名字',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标识符',
  `tag` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `access_mode` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '读写模型',
  `type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数据类型',
  `unit` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '单位',
  `remark` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '描述',
  `del_flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '删除标记',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建者',
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  `precision` smallint(6) NULL DEFAULT 1 COMMENT '倍率',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iot_prd_property
-- ----------------------------
INSERT INTO `iot_prd_property` VALUES (1, 0, 'A相电压力、', 'ua', '320', 'r', 'uint16', 'V', '320对应参量号', '0', '2024-07-31 13:04:36', '2024-07-31 13:04:36', 'admin', 'admin', 0, 1);

-- ----------------------------
-- Table structure for iot_product
-- ----------------------------
DROP TABLE IF EXISTS `iot_product`;
CREATE TABLE `iot_product`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名字',
  `key` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '产品标识',
  `cloud_product_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '云产品ID',
  `cloud_instance_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '云实例ID',
  `platform` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '平台',
  `protocol` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '协议',
  `node_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '节点类型',
  `net_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '网络类型',
  `data_format` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数据类型',
  `last_sync_time` bigint(20) NULL DEFAULT NULL COMMENT '最后一次同步时间',
  `factory` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '工厂名称',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '描述',
  `status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '产品状态',
  `extra` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '扩展字段',
  `del_flag` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建者',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '产品编码,对应可监控类型ID',
  `manufacturer` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '生产厂商',
  `active` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '1' COMMENT '是否启用',
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iot_product
-- ----------------------------
INSERT INTO `iot_product` VALUES (1, 'DPC设备网关', '', '', '', '', 'tcp', '10', '', '', 0, 'DPC', '设备网关，通过TCP方式接入平台', '', '', 0, '2024-07-29 12:53:54', '2024-07-29 12:53:54', 'admin', 'admin', '1', NULL, '1', NULL);
INSERT INTO `iot_product` VALUES (2, 'DPC光伏逆变器', NULL, NULL, NULL, NULL, 'modbus', '12', NULL, NULL, NULL, 'DPC', '无', NULL, NULL, 0, '2024-07-29 12:53:04', '2024-07-29 12:53:04', 'admin', 'admin', '33', NULL, '1', NULL);

-- ----------------------------
-- Table structure for iot_rule_engine
-- ----------------------------
DROP TABLE IF EXISTS `iot_rule_engine`;
CREATE TABLE `iot_rule_engine`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名字',
  `description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '描述',
  `status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态',
  `filter` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `data_resource_id` int(11) NULL DEFAULT NULL COMMENT '资源ID',
  `del_flag` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建者',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iot_rule_engine
-- ----------------------------

-- ----------------------------
-- Table structure for iot_scene
-- ----------------------------
DROP TABLE IF EXISTS `iot_scene`;
CREATE TABLE `iot_scene`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `status` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态 running stopped ',
  `conditions` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '条件',
  `actions` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '动作',
  `remark` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '描述',
  `del_flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '删除标记',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建者',
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iot_scene
-- ----------------------------

-- ----------------------------
-- Table structure for language_sdk
-- ----------------------------
DROP TABLE IF EXISTS `language_sdk`;
CREATE TABLE `language_sdk`  (
  `created` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `modified` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名字',
  `icon` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '图标',
  `addr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '描述',
  `sort` tinyint(4) NULL DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of language_sdk
-- ----------------------------

-- ----------------------------
-- Table structure for metrics
-- ----------------------------
DROP TABLE IF EXISTS `metrics`;
CREATE TABLE `metrics`  (
  `key` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `timestamp` bigint(20) NULL DEFAULT NULL,
  `cpu_used_percent` double NULL DEFAULT NULL,
  `memory_used` bigint(20) NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of metrics
-- ----------------------------

-- ----------------------------
-- Table structure for mqtt_auth
-- ----------------------------
DROP TABLE IF EXISTS `mqtt_auth`;
CREATE TABLE `mqtt_auth`  (
  `created` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `modified` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `resource_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '资源ID',
  `resource_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '资源类型',
  `client_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户端ID',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_mqtt_auth_client_id`(`client_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mqtt_auth
-- ----------------------------

-- ----------------------------
-- Table structure for msg_gather
-- ----------------------------
DROP TABLE IF EXISTS `msg_gather`;
CREATE TABLE `msg_gather`  (
  `created` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `modified` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `date` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '日期',
  `count` bigint(20) NULL DEFAULT NULL COMMENT '数量',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of msg_gather
-- ----------------------------

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product`  (
  `created` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `modified` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名字',
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '产品标识',
  `cloud_product_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '云产品ID',
  `cloud_instance_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '云实例ID',
  `platform` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '平台',
  `protocol` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '协议',
  `node_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '节点类型',
  `net_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '网络类型',
  `data_format` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数据类型',
  `last_sync_time` bigint(20) NULL DEFAULT NULL COMMENT '最后一次同步时间',
  `factory` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '工厂名称',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '描述',
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '产品状态',
  `extra` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '扩展字段',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_product_cloud_instance_id`(`cloud_instance_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of product
-- ----------------------------

-- ----------------------------
-- Table structure for quick_navigation
-- ----------------------------
DROP TABLE IF EXISTS `quick_navigation`;
CREATE TABLE `quick_navigation`  (
  `created` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `modified` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `icon` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `sort` bigint(20) NULL DEFAULT NULL,
  `jump_link` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of quick_navigation
-- ----------------------------

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `config_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '参数键值',
  `config_type` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'N' COMMENT '系统内置（Y是 N否）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`config_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '参数配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (1, '主框架页-默认皮肤样式名称', 'sys.index.skinName', 'skin-red', 'Y', 'admin', '2018-03-16 11:33:00', '', '2023-11-23 11:09:43', '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow');
INSERT INTO `sys_config` VALUES (2, '用户管理-账号初始密码', 'sys.user.initPassword', '123456', 'Y', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '初始化密码 123456');
INSERT INTO `sys_config` VALUES (3, '主框架页-侧边栏主题', 'sys.index.sideTheme', 'theme-blue', 'Y', 'admin', '2018-03-16 11:33:00', '', '2023-11-30 20:05:01', '深黑主题theme-dark，浅色主题theme-light，深蓝主题theme-blue');
INSERT INTO `sys_config` VALUES (4, '静态资源网盘存储', 'sys.resource.url', '/static', 'Y', 'admin', '2020-02-18 20:10:33', '', '2024-12-17 01:59:21', 'public目录下的静态资源存储到OSS/COS等网盘，将键值设为/public表示本地');

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint(20) NULL DEFAULT NULL COMMENT '父部门id',
  `ancestors` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '祖级列表',
  `dept_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '部门名称',
  `order_num` int(10) NULL DEFAULT NULL COMMENT '显示顺序',
  `leader` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '部门状态（0正常 1停用）',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '删除标记',
  `create_by` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  `dept_type` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '组织类型',
  PRIMARY KEY (`dept_id`) USING BTREE,
  INDEX `idx_ancestors`(`ancestors`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 115 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '部门表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (100, 0, '0', '某公司', 0, 'admin', '', '331641539@qq.com', '0', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2023-09-15 11:34:01', 0, '0');
INSERT INTO `sys_dept` VALUES (110, 100, '0,100,110', '项目部', 100, 'lostvip', '', '331641539@qq.com', '0', '0', 'admin', '2024-07-11 00:31:13', 'admin', '2024-07-11 00:31:13', 0, '0');
INSERT INTO `sys_dept` VALUES (111, 100, '0,100,111', '测试部', 51, 'lostvip', '', '331641539@qq.com', '0', '0', 'admin', '2020-03-01 09:40:48', 'admin', '2020-03-01 09:40:55', 0, '0');
INSERT INTO `sys_dept` VALUES (112, 100, '0,100,112', '运维部', 10, 'lostvip', '', '331641539@qq.com', '0', '0', 'admin', '2020-03-21 16:30:26', 'admin', '2023-10-29 22:30:15', 0, '0');
INSERT INTO `sys_dept` VALUES (113, 100, '0,100,113', '生产部', 2, 'lostvip', '', '331641539@qq.com', '0', '0', 'admin', '2023-09-15 11:38:29', 'admin', '2023-09-15 11:39:15', 0, '0');
INSERT INTO `sys_dept` VALUES (114, 100, '0,100,114', '研发部', 100, 'lostvip', '', '331641539@qq.com', '0', '0', 'admin', '2024-07-11 00:32:01', 'admin', '2024-07-11 00:32:01', 0, '0');

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `dict_sort` int(11) NULL DEFAULT NULL COMMENT '字典排序',
  `dict_label` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '字典标签',
  `dict_value` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '字典键值',
  `dict_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '字典类型',
  `css_class` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '表格回显样式',
  `is_default` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否默认（Y是 N否）',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '删除标记',
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '字典数据表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (1, 1, '男', '0', 'sys_user_sex', '', 'default', 'Y', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2023-09-21 21:45:51', '[[*{remark}]]', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (2, 2, '女', '1', 'sys_user_sex', '', '', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '性别女', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (4, 1, '显示', '0', 'sys_show_hide', '', 'primary', 'Y', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '显示菜单', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (5, 2, '隐藏', '1', 'sys_show_hide', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '隐藏菜单', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (6, 1, '正常', '0', 'sys_normal_disable', '', 'primary', 'Y', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '正常状态', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (7, 2, '停用', '1', 'sys_normal_disable', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '停用状态', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (8, 1, '正常', '0', 'sys_job_status', '', 'primary', 'Y', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '正常状态', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (9, 2, '暂停', '1', 'sys_job_status', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '停用状态', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (10, 1, '默认', 'DEFAULT', 'sys_job_group', '', '', 'Y', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '默认分组', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (11, 2, '系统', 'SYSTEM', 'sys_job_group', '', '', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统分组', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (12, 1, '是', 'Y', 'sys_yes_no', '', 'primary', 'Y', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统默认是', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (13, 2, '否', 'N', 'sys_yes_no', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统默认否', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (14, 1, '通知', '1', 'sys_notice_type', '', 'warning', 'Y', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '通知', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (15, 2, '公告', '2', 'sys_notice_type', '', 'success', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '公告', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (16, 1, '正常', '0', 'sys_notice_status', '', 'primary', 'Y', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '正常状态', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (17, 2, '关闭', '1', 'sys_notice_status', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '关闭状态', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (18, 1, '新增', '1', 'sys_oper_type', '', 'info', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '新增操作', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (19, 2, '修改', '2', 'sys_oper_type', '', 'info', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '修改操作', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (20, 3, '删除', '3', 'sys_oper_type', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '删除操作', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (21, 4, '授权', '4', 'sys_oper_type', '', 'primary', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '授权操作', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (22, 5, '导出', '5', 'sys_oper_type', '', 'warning', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '导出操作', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (23, 6, '导入', '6', 'sys_oper_type', '', 'warning', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '导入操作', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (24, 7, '强退', '7', 'sys_oper_type', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '强退操作', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (25, 8, '生成代码', '8', 'sys_oper_type', '', 'warning', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '生成操作', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (26, 9, '清空数据', '9', 'sys_oper_type', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '清空操作', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (27, 1, '成功', '0', 'sys_common_status', '', 'primary', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '正常状态', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (28, 2, '失败', '1', 'sys_common_status', '', 'danger', 'N', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '停用状态', '0', NULL);
INSERT INTO `sys_dict_data` VALUES (29, 0, '免费用户', '0', 'zjuser_type', NULL, 'default', 'Y', '0', 'admin', '2019-12-02 16:56:16', 'admin', NULL, NULL, '0', NULL);
INSERT INTO `sys_dict_data` VALUES (30, 1, '付费用户', '1', 'zjuser_type', NULL, 'primary', 'Y', '0', 'admin', '2019-12-02 16:56:40', 'admin', NULL, NULL, '0', NULL);
INSERT INTO `sys_dict_data` VALUES (31, 0, '微信用户', '0', 'zxuser_type', NULL, 'default', 'Y', '0', 'admin', '2019-12-02 17:14:32', 'admin', NULL, NULL, '0', NULL);
INSERT INTO `sys_dict_data` VALUES (32, 1, 'QQ用户', '1', 'zxuser_type', NULL, 'primary', 'N', '0', 'admin', '2019-12-02 17:14:55', 'admin', NULL, NULL, '0', NULL);
INSERT INTO `sys_dict_data` VALUES (33, 2, '抖音用户', '2', 'zxuser_type', NULL, 'primary', 'N', '0', 'admin', '2019-12-02 17:15:21', 'admin', NULL, NULL, '0', NULL);
INSERT INTO `sys_dict_data` VALUES (34, 100, '顶顶顶', '滴滴答答', 'sssdgsfhdfhk', '顶顶顶', 'default', 'Y', '0', 'admin', '2024-12-19 02:03:48', '', '2024-12-19 02:03:48', '', '0', 0);

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
  `dict_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `dict_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '字典名称',
  `dict_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '字典类型',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '删除标记',
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  PRIMARY KEY (`dict_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '字典类型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES (1, '用户性别', 'sys_user_sex', '0', 'admin', '2018-03-16 11:33:00', '', '2020-03-01 09:41:19', '用户性别列表', '0', NULL);
INSERT INTO `sys_dict_type` VALUES (2, '菜单状态', 'sys_show_hide', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '菜单状态列表', '0', NULL);
INSERT INTO `sys_dict_type` VALUES (3, '系统开关', 'sys_normal_disable', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统开关列表', '0', NULL);
INSERT INTO `sys_dict_type` VALUES (4, '任务状态', 'sys_job_status', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '任务状态列表', '0', NULL);
INSERT INTO `sys_dict_type` VALUES (5, '任务分组', 'sys_job_group', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '任务分组列表', '0', NULL);
INSERT INTO `sys_dict_type` VALUES (6, '系统是否', 'sys_yes_no', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统是否列表', '0', NULL);
INSERT INTO `sys_dict_type` VALUES (7, '通知类型', 'sys_notice_type', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '通知类型列表', '0', NULL);
INSERT INTO `sys_dict_type` VALUES (8, '通知状态', 'sys_notice_status', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '通知状态列表', '0', NULL);
INSERT INTO `sys_dict_type` VALUES (9, '操作类型', 'sys_oper_type', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '操作类型列表', '0', NULL);
INSERT INTO `sys_dict_type` VALUES (10, '系统状态', 'sys_common_status', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '登录状态列表', '0', NULL);
INSERT INTO `sys_dict_type` VALUES (11, '专家用户类别', 'zjuser_type', '0', 'admin', '2019-12-02 16:55:42', 'admin', NULL, NULL, '0', NULL);
INSERT INTO `sys_dict_type` VALUES (12, '咨询用户类别', 'zxuser_type', '0', 'admin', '2019-12-02 17:14:07', 'admin', NULL, NULL, '0', NULL);
INSERT INTO `sys_dict_type` VALUES (14, '焚烧', 'sssdgsfhdfhk', '0', 'admin', '2024-12-19 02:02:54', '', '2024-12-19 02:02:54', '', '0', 0);

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job`  (
  `job_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `job_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_params` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '参数',
  `job_group` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `invoke_target` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '调用目标字符串',
  `cron_expression` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT 'cron执行表达式',
  `misfire_policy` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1' COMMENT '计划执行策略（1多次执行 2执行一次）',
  `concurrent` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1' COMMENT '是否并发执行（0允许 1禁止）',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '状态（0正常 1暂停）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`job_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '定时任务调度表' ROW_FORMAT = Dynamic;

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
  `job_log_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务日志ID',
  `job_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '任务名称',
  `job_group` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '任务组名',
  `invoke_target` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '调用目标字符串',
  `job_message` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '日志信息',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '执行状态（0正常 1失败）',
  `exception_info` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '异常信息',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`job_log_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '定时任务调度日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_job_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_logininfor
-- ----------------------------
DROP TABLE IF EXISTS `sys_logininfor`;
CREATE TABLE `sys_logininfor`  (
  `info_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `login_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '操作系统',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '提示消息',
  `login_time` datetime(0) NULL DEFAULT NULL COMMENT '访问时间',
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1051 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '系统访问记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_logininfor
-- ----------------------------
INSERT INTO `sys_logininfor` VALUES (381, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 13:29:30');
INSERT INTO `sys_logininfor` VALUES (382, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 13:59:34');
INSERT INTO `sys_logininfor` VALUES (383, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 14:33:41');
INSERT INTO `sys_logininfor` VALUES (384, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 15:07:56');
INSERT INTO `sys_logininfor` VALUES (385, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 15:20:08');
INSERT INTO `sys_logininfor` VALUES (386, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 15:22:20');
INSERT INTO `sys_logininfor` VALUES (387, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 15:38:52');
INSERT INTO `sys_logininfor` VALUES (388, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 15:39:51');
INSERT INTO `sys_logininfor` VALUES (389, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 15:49:13');
INSERT INTO `sys_logininfor` VALUES (390, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 15:50:04');
INSERT INTO `sys_logininfor` VALUES (391, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 15:52:06');
INSERT INTO `sys_logininfor` VALUES (392, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 15:53:42');
INSERT INTO `sys_logininfor` VALUES (393, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 15:56:50');
INSERT INTO `sys_logininfor` VALUES (394, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 16:01:00');
INSERT INTO `sys_logininfor` VALUES (395, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 16:02:31');
INSERT INTO `sys_logininfor` VALUES (396, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 16:03:26');
INSERT INTO `sys_logininfor` VALUES (397, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 16:09:14');
INSERT INTO `sys_logininfor` VALUES (398, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 16:20:03');
INSERT INTO `sys_logininfor` VALUES (399, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 16:27:41');
INSERT INTO `sys_logininfor` VALUES (400, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 16:34:51');
INSERT INTO `sys_logininfor` VALUES (401, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 18:04:05');
INSERT INTO `sys_logininfor` VALUES (402, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 18:12:21');
INSERT INTO `sys_logininfor` VALUES (403, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 18:17:53');
INSERT INTO `sys_logininfor` VALUES (404, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 18:20:11');
INSERT INTO `sys_logininfor` VALUES (405, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 18:44:10');
INSERT INTO `sys_logininfor` VALUES (406, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 18:45:20');
INSERT INTO `sys_logininfor` VALUES (407, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 18:51:47');
INSERT INTO `sys_logininfor` VALUES (408, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 18:55:35');
INSERT INTO `sys_logininfor` VALUES (409, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 18:57:44');
INSERT INTO `sys_logininfor` VALUES (410, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 18:59:49');
INSERT INTO `sys_logininfor` VALUES (411, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 19:01:42');
INSERT INTO `sys_logininfor` VALUES (412, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 19:04:18');
INSERT INTO `sys_logininfor` VALUES (413, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 19:28:51');
INSERT INTO `sys_logininfor` VALUES (414, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 19:30:01');
INSERT INTO `sys_logininfor` VALUES (415, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 20:42:57');
INSERT INTO `sys_logininfor` VALUES (416, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 20:46:46');
INSERT INTO `sys_logininfor` VALUES (417, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 20:48:01');
INSERT INTO `sys_logininfor` VALUES (418, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 21:08:41');
INSERT INTO `sys_logininfor` VALUES (419, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 21:43:49');
INSERT INTO `sys_logininfor` VALUES (420, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 21:55:36');
INSERT INTO `sys_logininfor` VALUES (421, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 21:56:54');
INSERT INTO `sys_logininfor` VALUES (422, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 22:00:25');
INSERT INTO `sys_logininfor` VALUES (423, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-24 22:07:46');
INSERT INTO `sys_logininfor` VALUES (424, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-25 14:18:52');
INSERT INTO `sys_logininfor` VALUES (425, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-25 14:59:39');
INSERT INTO `sys_logininfor` VALUES (426, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 13:46:03');
INSERT INTO `sys_logininfor` VALUES (427, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 16:19:47');
INSERT INTO `sys_logininfor` VALUES (428, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 16:23:40');
INSERT INTO `sys_logininfor` VALUES (429, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 16:24:22');
INSERT INTO `sys_logininfor` VALUES (430, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 16:26:52');
INSERT INTO `sys_logininfor` VALUES (431, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 17:38:39');
INSERT INTO `sys_logininfor` VALUES (432, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 17:41:48');
INSERT INTO `sys_logininfor` VALUES (433, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 17:43:37');
INSERT INTO `sys_logininfor` VALUES (434, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 17:45:25');
INSERT INTO `sys_logininfor` VALUES (435, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 19:38:09');
INSERT INTO `sys_logininfor` VALUES (436, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 19:39:43');
INSERT INTO `sys_logininfor` VALUES (437, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 19:41:09');
INSERT INTO `sys_logininfor` VALUES (438, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 20:35:01');
INSERT INTO `sys_logininfor` VALUES (439, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 21:07:38');
INSERT INTO `sys_logininfor` VALUES (440, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 21:32:22');
INSERT INTO `sys_logininfor` VALUES (441, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 21:35:00');
INSERT INTO `sys_logininfor` VALUES (442, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 21:36:07');
INSERT INTO `sys_logininfor` VALUES (443, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 21:51:02');
INSERT INTO `sys_logininfor` VALUES (444, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 21:53:04');
INSERT INTO `sys_logininfor` VALUES (445, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 21:54:52');
INSERT INTO `sys_logininfor` VALUES (446, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 21:56:22');
INSERT INTO `sys_logininfor` VALUES (447, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 22:00:20');
INSERT INTO `sys_logininfor` VALUES (448, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 22:08:10');
INSERT INTO `sys_logininfor` VALUES (449, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 22:10:38');
INSERT INTO `sys_logininfor` VALUES (450, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-26 22:18:28');
INSERT INTO `sys_logininfor` VALUES (451, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 09:39:14');
INSERT INTO `sys_logininfor` VALUES (452, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 09:40:27');
INSERT INTO `sys_logininfor` VALUES (453, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 09:46:12');
INSERT INTO `sys_logininfor` VALUES (454, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 09:47:49');
INSERT INTO `sys_logininfor` VALUES (455, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 09:49:05');
INSERT INTO `sys_logininfor` VALUES (456, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 09:52:40');
INSERT INTO `sys_logininfor` VALUES (457, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 10:00:55');
INSERT INTO `sys_logininfor` VALUES (458, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 10:03:04');
INSERT INTO `sys_logininfor` VALUES (459, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 10:03:52');
INSERT INTO `sys_logininfor` VALUES (460, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 10:11:50');
INSERT INTO `sys_logininfor` VALUES (461, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 10:12:47');
INSERT INTO `sys_logininfor` VALUES (462, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 10:17:20');
INSERT INTO `sys_logininfor` VALUES (463, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 10:18:32');
INSERT INTO `sys_logininfor` VALUES (464, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 10:39:41');
INSERT INTO `sys_logininfor` VALUES (465, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 10:44:16');
INSERT INTO `sys_logininfor` VALUES (466, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 11:04:58');
INSERT INTO `sys_logininfor` VALUES (467, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 11:43:43');
INSERT INTO `sys_logininfor` VALUES (468, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 12:02:57');
INSERT INTO `sys_logininfor` VALUES (469, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 21:00:55');
INSERT INTO `sys_logininfor` VALUES (470, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 21:36:47');
INSERT INTO `sys_logininfor` VALUES (471, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 21:39:50');
INSERT INTO `sys_logininfor` VALUES (472, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 21:40:51');
INSERT INTO `sys_logininfor` VALUES (473, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 21:42:14');
INSERT INTO `sys_logininfor` VALUES (474, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 21:43:20');
INSERT INTO `sys_logininfor` VALUES (475, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-27 21:45:45');
INSERT INTO `sys_logininfor` VALUES (476, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-28 20:56:27');
INSERT INTO `sys_logininfor` VALUES (477, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-28 21:06:20');
INSERT INTO `sys_logininfor` VALUES (478, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-03-29 14:54:13');
INSERT INTO `sys_logininfor` VALUES (479, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-04-16 15:50:09');
INSERT INTO `sys_logininfor` VALUES (480, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-04-16 15:56:04');
INSERT INTO `sys_logininfor` VALUES (481, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-04-16 16:06:16');
INSERT INTO `sys_logininfor` VALUES (482, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-05-05 13:50:54');
INSERT INTO `sys_logininfor` VALUES (483, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-05-05 13:53:13');
INSERT INTO `sys_logininfor` VALUES (484, 'admin', '::1', '内网IP', 'Chrome', 'Intel Mac OS X 10_14_6', '0', '登陆成功', '2020-05-05 13:59:27');
INSERT INTO `sys_logininfor` VALUES (485, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2021-06-01 23:55:03');
INSERT INTO `sys_logininfor` VALUES (486, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2021-06-02 00:02:10');
INSERT INTO `sys_logininfor` VALUES (487, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2021-06-19 23:08:06');
INSERT INTO `sys_logininfor` VALUES (488, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2021-06-19 23:09:50');
INSERT INTO `sys_logininfor` VALUES (489, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2021-06-19 23:14:10');
INSERT INTO `sys_logininfor` VALUES (490, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2021-06-19 23:23:11');
INSERT INTO `sys_logininfor` VALUES (491, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2021-06-19 23:42:36');
INSERT INTO `sys_logininfor` VALUES (492, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2021-06-20 00:05:22');
INSERT INTO `sys_logininfor` VALUES (493, 'admin', '::1', '内网IP', 'Chrome', 'Android 6.0', '0', '登陆成功', '2021-06-20 00:12:47');
INSERT INTO `sys_logininfor` VALUES (494, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2021-06-20 00:25:33');
INSERT INTO `sys_logininfor` VALUES (495, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2021-06-20 00:25:42');
INSERT INTO `sys_logininfor` VALUES (496, 'admin', '::1', '内网IP', 'Chrome', 'Android 6.0', '0', '账号或密码不正确', '2021-06-20 00:25:55');
INSERT INTO `sys_logininfor` VALUES (497, 'admin', '::1', '内网IP', 'Chrome', 'Android 6.0', '0', '账号或密码不正确', '2021-06-20 00:26:45');
INSERT INTO `sys_logininfor` VALUES (498, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2021-06-20 00:27:02');
INSERT INTO `sys_logininfor` VALUES (499, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2021-06-20 00:27:14');
INSERT INTO `sys_logininfor` VALUES (500, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2021-06-20 00:28:01');
INSERT INTO `sys_logininfor` VALUES (501, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2021-06-20 13:43:34');
INSERT INTO `sys_logininfor` VALUES (502, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2021-06-20 13:44:46');
INSERT INTO `sys_logininfor` VALUES (503, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2021-06-20 13:48:29');
INSERT INTO `sys_logininfor` VALUES (504, 'admin', '127.0.0.1', '内网IP', 'Firefox', 'Windows 10', '0', '登陆成功', '2021-06-20 13:51:09');
INSERT INTO `sys_logininfor` VALUES (505, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2021-06-20 15:59:42');
INSERT INTO `sys_logininfor` VALUES (506, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2022-11-18 14:52:04');
INSERT INTO `sys_logininfor` VALUES (507, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2022-11-18 15:10:22');
INSERT INTO `sys_logininfor` VALUES (508, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2022-11-18 15:10:33');
INSERT INTO `sys_logininfor` VALUES (509, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-03-15 13:25:42');
INSERT INTO `sys_logininfor` VALUES (510, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-03-15 13:26:23');
INSERT INTO `sys_logininfor` VALUES (511, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 12:39:39');
INSERT INTO `sys_logininfor` VALUES (512, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 14:01:47');
INSERT INTO `sys_logininfor` VALUES (513, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 14:07:49');
INSERT INTO `sys_logininfor` VALUES (514, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 14:10:43');
INSERT INTO `sys_logininfor` VALUES (515, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 14:11:27');
INSERT INTO `sys_logininfor` VALUES (516, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 14:18:41');
INSERT INTO `sys_logininfor` VALUES (517, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 14:19:45');
INSERT INTO `sys_logininfor` VALUES (518, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 14:26:06');
INSERT INTO `sys_logininfor` VALUES (519, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 14:27:37');
INSERT INTO `sys_logininfor` VALUES (520, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 14:27:50');
INSERT INTO `sys_logininfor` VALUES (521, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 14:31:14');
INSERT INTO `sys_logininfor` VALUES (522, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 14:52:09');
INSERT INTO `sys_logininfor` VALUES (523, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 14:56:23');
INSERT INTO `sys_logininfor` VALUES (524, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 15:24:11');
INSERT INTO `sys_logininfor` VALUES (525, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 15:24:26');
INSERT INTO `sys_logininfor` VALUES (526, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 15:24:48');
INSERT INTO `sys_logininfor` VALUES (527, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 15:25:35');
INSERT INTO `sys_logininfor` VALUES (528, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 15:26:26');
INSERT INTO `sys_logininfor` VALUES (529, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 15:31:13');
INSERT INTO `sys_logininfor` VALUES (530, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 15:36:56');
INSERT INTO `sys_logininfor` VALUES (531, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 15:38:33');
INSERT INTO `sys_logininfor` VALUES (532, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 15:48:19');
INSERT INTO `sys_logininfor` VALUES (533, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 15:50:10');
INSERT INTO `sys_logininfor` VALUES (534, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 16:02:17');
INSERT INTO `sys_logininfor` VALUES (535, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 16:03:10');
INSERT INTO `sys_logininfor` VALUES (536, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 16:07:31');
INSERT INTO `sys_logininfor` VALUES (537, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 16:20:05');
INSERT INTO `sys_logininfor` VALUES (538, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 16:28:53');
INSERT INTO `sys_logininfor` VALUES (539, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 16:32:50');
INSERT INTO `sys_logininfor` VALUES (540, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 16:35:27');
INSERT INTO `sys_logininfor` VALUES (541, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 16:53:17');
INSERT INTO `sys_logininfor` VALUES (542, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 21:52:15');
INSERT INTO `sys_logininfor` VALUES (543, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 21:53:02');
INSERT INTO `sys_logininfor` VALUES (544, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 22:30:00');
INSERT INTO `sys_logininfor` VALUES (545, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 22:42:55');
INSERT INTO `sys_logininfor` VALUES (546, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 22:52:28');
INSERT INTO `sys_logininfor` VALUES (547, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 22:53:23');
INSERT INTO `sys_logininfor` VALUES (548, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 22:53:35');
INSERT INTO `sys_logininfor` VALUES (549, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 22:59:20');
INSERT INTO `sys_logininfor` VALUES (550, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-14 23:02:14');
INSERT INTO `sys_logininfor` VALUES (551, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 10:57:59');
INSERT INTO `sys_logininfor` VALUES (552, 'qy1156010001', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 11:03:41');
INSERT INTO `sys_logininfor` VALUES (553, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2023-09-15 11:04:13');
INSERT INTO `sys_logininfor` VALUES (554, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2023-09-15 11:04:20');
INSERT INTO `sys_logininfor` VALUES (555, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2023-09-15 11:04:27');
INSERT INTO `sys_logininfor` VALUES (556, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2023-09-15 11:04:55');
INSERT INTO `sys_logininfor` VALUES (557, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 11:05:08');
INSERT INTO `sys_logininfor` VALUES (558, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 11:06:05');
INSERT INTO `sys_logininfor` VALUES (559, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 11:06:49');
INSERT INTO `sys_logininfor` VALUES (560, 'qy1156010001', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2023-09-15 11:09:08');
INSERT INTO `sys_logininfor` VALUES (561, 'qy1156010001', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 11:09:28');
INSERT INTO `sys_logininfor` VALUES (562, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 11:12:52');
INSERT INTO `sys_logininfor` VALUES (563, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 11:22:25');
INSERT INTO `sys_logininfor` VALUES (564, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 11:31:39');
INSERT INTO `sys_logininfor` VALUES (565, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 11:51:02');
INSERT INTO `sys_logininfor` VALUES (566, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 11:59:07');
INSERT INTO `sys_logininfor` VALUES (567, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 12:00:08');
INSERT INTO `sys_logininfor` VALUES (568, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 13:40:50');
INSERT INTO `sys_logininfor` VALUES (569, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 13:41:20');
INSERT INTO `sys_logininfor` VALUES (570, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 13:42:10');
INSERT INTO `sys_logininfor` VALUES (571, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 13:42:59');
INSERT INTO `sys_logininfor` VALUES (572, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 17:38:39');
INSERT INTO `sys_logininfor` VALUES (573, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 17:55:57');
INSERT INTO `sys_logininfor` VALUES (574, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 17:56:53');
INSERT INTO `sys_logininfor` VALUES (575, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 17:57:49');
INSERT INTO `sys_logininfor` VALUES (576, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 17:58:21');
INSERT INTO `sys_logininfor` VALUES (577, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 17:59:03');
INSERT INTO `sys_logininfor` VALUES (578, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 18:06:15');
INSERT INTO `sys_logininfor` VALUES (579, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 18:07:15');
INSERT INTO `sys_logininfor` VALUES (580, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 18:09:38');
INSERT INTO `sys_logininfor` VALUES (581, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 18:17:37');
INSERT INTO `sys_logininfor` VALUES (582, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 18:31:33');
INSERT INTO `sys_logininfor` VALUES (583, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 18:40:20');
INSERT INTO `sys_logininfor` VALUES (584, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 18:44:15');
INSERT INTO `sys_logininfor` VALUES (585, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 18:55:01');
INSERT INTO `sys_logininfor` VALUES (586, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 19:23:07');
INSERT INTO `sys_logininfor` VALUES (587, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 19:23:50');
INSERT INTO `sys_logininfor` VALUES (588, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 19:25:46');
INSERT INTO `sys_logininfor` VALUES (589, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 19:26:19');
INSERT INTO `sys_logininfor` VALUES (590, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-15 22:09:38');
INSERT INTO `sys_logininfor` VALUES (591, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-16 18:46:02');
INSERT INTO `sys_logininfor` VALUES (592, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-17 18:39:38');
INSERT INTO `sys_logininfor` VALUES (593, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-17 19:02:19');
INSERT INTO `sys_logininfor` VALUES (594, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-17 19:03:39');
INSERT INTO `sys_logininfor` VALUES (595, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-17 19:15:39');
INSERT INTO `sys_logininfor` VALUES (596, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-17 19:19:12');
INSERT INTO `sys_logininfor` VALUES (597, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-17 19:35:40');
INSERT INTO `sys_logininfor` VALUES (598, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-17 19:38:42');
INSERT INTO `sys_logininfor` VALUES (599, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-17 19:41:11');
INSERT INTO `sys_logininfor` VALUES (600, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-17 20:04:07');
INSERT INTO `sys_logininfor` VALUES (601, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-17 22:29:46');
INSERT INTO `sys_logininfor` VALUES (602, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-17 22:42:10');
INSERT INTO `sys_logininfor` VALUES (603, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-17 22:48:04');
INSERT INTO `sys_logininfor` VALUES (604, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-17 23:41:45');
INSERT INTO `sys_logininfor` VALUES (605, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-17 23:43:05');
INSERT INTO `sys_logininfor` VALUES (606, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-18 09:57:25');
INSERT INTO `sys_logininfor` VALUES (607, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-18 09:58:17');
INSERT INTO `sys_logininfor` VALUES (608, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-18 10:02:03');
INSERT INTO `sys_logininfor` VALUES (609, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-18 10:09:47');
INSERT INTO `sys_logininfor` VALUES (610, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-18 10:17:10');
INSERT INTO `sys_logininfor` VALUES (611, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-18 11:36:20');
INSERT INTO `sys_logininfor` VALUES (612, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-18 14:52:17');
INSERT INTO `sys_logininfor` VALUES (613, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-18 15:13:45');
INSERT INTO `sys_logininfor` VALUES (614, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-18 21:36:17');
INSERT INTO `sys_logininfor` VALUES (615, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-18 21:40:43');
INSERT INTO `sys_logininfor` VALUES (616, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-18 21:42:55');
INSERT INTO `sys_logininfor` VALUES (617, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-18 21:52:45');
INSERT INTO `sys_logininfor` VALUES (618, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-18 21:53:51');
INSERT INTO `sys_logininfor` VALUES (619, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-18 21:56:52');
INSERT INTO `sys_logininfor` VALUES (620, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-18 21:58:55');
INSERT INTO `sys_logininfor` VALUES (621, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-18 22:10:26');
INSERT INTO `sys_logininfor` VALUES (622, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-19 13:57:09');
INSERT INTO `sys_logininfor` VALUES (623, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-19 13:59:35');
INSERT INTO `sys_logininfor` VALUES (624, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-19 14:00:12');
INSERT INTO `sys_logininfor` VALUES (625, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-19 14:01:44');
INSERT INTO `sys_logininfor` VALUES (626, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-19 14:02:05');
INSERT INTO `sys_logininfor` VALUES (627, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-19 14:14:16');
INSERT INTO `sys_logininfor` VALUES (628, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-19 20:21:15');
INSERT INTO `sys_logininfor` VALUES (629, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-19 20:22:48');
INSERT INTO `sys_logininfor` VALUES (630, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-19 21:22:54');
INSERT INTO `sys_logininfor` VALUES (631, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-19 21:45:24');
INSERT INTO `sys_logininfor` VALUES (632, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-19 21:53:03');
INSERT INTO `sys_logininfor` VALUES (633, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-19 21:54:07');
INSERT INTO `sys_logininfor` VALUES (634, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-19 21:56:05');
INSERT INTO `sys_logininfor` VALUES (635, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-19 22:00:59');
INSERT INTO `sys_logininfor` VALUES (636, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-20 20:50:09');
INSERT INTO `sys_logininfor` VALUES (637, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-20 21:51:15');
INSERT INTO `sys_logininfor` VALUES (638, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-20 21:56:29');
INSERT INTO `sys_logininfor` VALUES (639, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-20 22:01:04');
INSERT INTO `sys_logininfor` VALUES (640, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-20 22:02:36');
INSERT INTO `sys_logininfor` VALUES (641, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-20 22:05:11');
INSERT INTO `sys_logininfor` VALUES (642, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-20 22:12:08');
INSERT INTO `sys_logininfor` VALUES (643, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-20 22:13:28');
INSERT INTO `sys_logininfor` VALUES (644, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-20 22:21:57');
INSERT INTO `sys_logininfor` VALUES (645, 'admin', '111.193.7.238', '北京市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-20 22:48:39');
INSERT INTO `sys_logininfor` VALUES (646, 'admin', '223.160.129.142', '北京市', 'Chrome', 'Android 12', '0', '登陆成功', '2023-09-21 08:37:06');
INSERT INTO `sys_logininfor` VALUES (647, 'admin', '223.160.129.142', '北京市', 'Chrome', 'Android 12', '0', '登陆成功', '2023-09-21 08:46:43');
INSERT INTO `sys_logininfor` VALUES (648, 'admin', '175.0.61.79', '长沙市', 'Firefox', 'Windows 10', '0', '登陆成功', '2023-09-21 08:49:59');
INSERT INTO `sys_logininfor` VALUES (649, 'admin', '219.143.126.167', '北京市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 09:42:32');
INSERT INTO `sys_logininfor` VALUES (650, 'admin', '117.36.67.20', '西安市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 11:19:32');
INSERT INTO `sys_logininfor` VALUES (652, 'admin', '219.156.133.200', '南阳市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 17:14:52');
INSERT INTO `sys_logininfor` VALUES (653, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 19:46:27');
INSERT INTO `sys_logininfor` VALUES (654, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 20:29:28');
INSERT INTO `sys_logininfor` VALUES (655, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 20:44:28');
INSERT INTO `sys_logininfor` VALUES (656, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 20:51:28');
INSERT INTO `sys_logininfor` VALUES (657, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 20:57:58');
INSERT INTO `sys_logininfor` VALUES (658, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 21:00:44');
INSERT INTO `sys_logininfor` VALUES (659, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 21:02:41');
INSERT INTO `sys_logininfor` VALUES (660, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 21:03:00');
INSERT INTO `sys_logininfor` VALUES (661, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 21:18:52');
INSERT INTO `sys_logininfor` VALUES (662, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 21:21:10');
INSERT INTO `sys_logininfor` VALUES (663, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 21:23:00');
INSERT INTO `sys_logininfor` VALUES (664, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 21:39:51');
INSERT INTO `sys_logininfor` VALUES (665, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 21:41:43');
INSERT INTO `sys_logininfor` VALUES (666, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 21:41:50');
INSERT INTO `sys_logininfor` VALUES (667, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 21:45:18');
INSERT INTO `sys_logininfor` VALUES (668, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-21 21:48:20');
INSERT INTO `sys_logininfor` VALUES (669, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-22 20:44:58');
INSERT INTO `sys_logininfor` VALUES (670, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-22 21:03:03');
INSERT INTO `sys_logininfor` VALUES (671, 'admin', '111.193.7.238', '北京市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-22 13:52:07');
INSERT INTO `sys_logininfor` VALUES (672, 'admin', '111.193.7.238', '北京市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-22 22:08:12');
INSERT INTO `sys_logininfor` VALUES (673, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-22 22:32:37');
INSERT INTO `sys_logininfor` VALUES (674, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-22 22:33:07');
INSERT INTO `sys_logininfor` VALUES (675, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-22 22:33:55');
INSERT INTO `sys_logininfor` VALUES (676, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-22 22:34:32');
INSERT INTO `sys_logininfor` VALUES (677, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Android 6.0', '0', '登陆成功', '2023-09-22 22:37:48');
INSERT INTO `sys_logininfor` VALUES (678, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Android 6.0', '0', '登陆成功', '2023-09-22 22:39:20');
INSERT INTO `sys_logininfor` VALUES (679, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Android 6.0', '0', '登陆成功', '2023-09-22 22:41:13');
INSERT INTO `sys_logininfor` VALUES (680, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-22 22:46:23');
INSERT INTO `sys_logininfor` VALUES (681, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-22 22:48:00');
INSERT INTO `sys_logininfor` VALUES (682, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-22 23:09:45');
INSERT INTO `sys_logininfor` VALUES (683, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-22 23:10:03');
INSERT INTO `sys_logininfor` VALUES (684, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-22 23:16:27');
INSERT INTO `sys_logininfor` VALUES (685, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-22 23:16:46');
INSERT INTO `sys_logininfor` VALUES (686, 'admin', '111.193.7.238', '北京市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-23 11:14:41');
INSERT INTO `sys_logininfor` VALUES (687, 'admin', '111.193.7.238', '北京市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-23 11:15:07');
INSERT INTO `sys_logininfor` VALUES (688, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-23 11:32:46');
INSERT INTO `sys_logininfor` VALUES (689, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-23 11:33:44');
INSERT INTO `sys_logininfor` VALUES (690, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-23 11:40:28');
INSERT INTO `sys_logininfor` VALUES (691, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Android 6.0', '0', '登陆成功', '2023-09-23 11:44:05');
INSERT INTO `sys_logininfor` VALUES (692, 'admin', '111.193.7.238', '北京市', 'Chrome', 'Android 6.0', '0', '登陆成功', '2023-09-23 11:51:22');
INSERT INTO `sys_logininfor` VALUES (693, 'admin', '111.193.7.238', '北京市', 'Chrome', 'Android 6.0', '0', '登陆成功', '2023-09-23 11:54:15');
INSERT INTO `sys_logininfor` VALUES (694, 'admin', '111.193.7.238', '北京市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-23 11:57:32');
INSERT INTO `sys_logininfor` VALUES (695, 'admin', '111.193.7.238', '北京市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-23 12:02:59');
INSERT INTO `sys_logininfor` VALUES (696, 'admin', '111.193.7.238', '北京市', 'Chrome', 'Android 6.0', '0', '登陆成功', '2023-09-23 12:21:22');
INSERT INTO `sys_logininfor` VALUES (697, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-23 12:34:44');
INSERT INTO `sys_logininfor` VALUES (698, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-23 12:34:48');
INSERT INTO `sys_logininfor` VALUES (699, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-23 12:34:50');
INSERT INTO `sys_logininfor` VALUES (700, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-23 12:35:01');
INSERT INTO `sys_logininfor` VALUES (701, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Android 6.0', '0', '登陆成功', '2023-09-23 12:35:35');
INSERT INTO `sys_logininfor` VALUES (702, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Android 6.0', '0', '登陆成功', '2023-09-23 12:35:48');
INSERT INTO `sys_logininfor` VALUES (703, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Android 6.0', '0', '登陆成功', '2023-09-23 12:35:50');
INSERT INTO `sys_logininfor` VALUES (704, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Android 6.0', '0', '登陆成功', '2023-09-23 12:35:53');
INSERT INTO `sys_logininfor` VALUES (705, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-23 12:37:41');
INSERT INTO `sys_logininfor` VALUES (706, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-23 19:23:46');
INSERT INTO `sys_logininfor` VALUES (707, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-23 22:32:32');
INSERT INTO `sys_logininfor` VALUES (708, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-24 12:33:09');
INSERT INTO `sys_logininfor` VALUES (709, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-24 13:26:48');
INSERT INTO `sys_logininfor` VALUES (710, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-24 19:10:41');
INSERT INTO `sys_logininfor` VALUES (711, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-24 20:24:53');
INSERT INTO `sys_logininfor` VALUES (712, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-24 20:27:00');
INSERT INTO `sys_logininfor` VALUES (713, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-24 20:27:46');
INSERT INTO `sys_logininfor` VALUES (714, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-24 20:28:00');
INSERT INTO `sys_logininfor` VALUES (715, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-24 20:29:18');
INSERT INTO `sys_logininfor` VALUES (716, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-24 20:29:38');
INSERT INTO `sys_logininfor` VALUES (717, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-24 20:33:07');
INSERT INTO `sys_logininfor` VALUES (718, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-24 20:53:15');
INSERT INTO `sys_logininfor` VALUES (719, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-24 20:59:07');
INSERT INTO `sys_logininfor` VALUES (720, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-25 20:18:55');
INSERT INTO `sys_logininfor` VALUES (721, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-25 20:35:21');
INSERT INTO `sys_logininfor` VALUES (722, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-25 20:36:01');
INSERT INTO `sys_logininfor` VALUES (723, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-25 21:00:34');
INSERT INTO `sys_logininfor` VALUES (724, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-25 21:56:59');
INSERT INTO `sys_logininfor` VALUES (725, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-25 21:57:48');
INSERT INTO `sys_logininfor` VALUES (726, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-26 18:21:50');
INSERT INTO `sys_logininfor` VALUES (727, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-09-26 20:57:56');
INSERT INTO `sys_logininfor` VALUES (728, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-04 18:55:05');
INSERT INTO `sys_logininfor` VALUES (729, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-05 21:02:15');
INSERT INTO `sys_logininfor` VALUES (730, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-10 17:56:16');
INSERT INTO `sys_logininfor` VALUES (731, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-15 09:48:58');
INSERT INTO `sys_logininfor` VALUES (732, 'admin', '172.20.0.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-15 10:31:08');
INSERT INTO `sys_logininfor` VALUES (733, 'admin', '172.20.0.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-15 10:40:45');
INSERT INTO `sys_logininfor` VALUES (734, 'admin', '172.20.0.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-15 12:37:45');
INSERT INTO `sys_logininfor` VALUES (735, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-15 12:43:57');
INSERT INTO `sys_logininfor` VALUES (736, 'admin', '111.194.208.211', '北京市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-15 21:09:53');
INSERT INTO `sys_logininfor` VALUES (737, 'admin', '111.194.208.211', '北京市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-15 21:15:19');
INSERT INTO `sys_logininfor` VALUES (738, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2023-10-15 22:26:28');
INSERT INTO `sys_logininfor` VALUES (739, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-15 22:26:37');
INSERT INTO `sys_logininfor` VALUES (740, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-16 10:12:18');
INSERT INTO `sys_logininfor` VALUES (741, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-16 19:38:56');
INSERT INTO `sys_logininfor` VALUES (742, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-17 09:03:38');
INSERT INTO `sys_logininfor` VALUES (743, 'admin', '10.0.0.2', '', 'Safari', 'Intel Mac OS X 10_15_7', '0', '登陆成功', '2023-10-17 10:01:21');
INSERT INTO `sys_logininfor` VALUES (744, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2023-10-17 14:08:12');
INSERT INTO `sys_logininfor` VALUES (745, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-17 14:08:20');
INSERT INTO `sys_logininfor` VALUES (746, 'admin', '111.194.208.211', '北京市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-18 22:13:20');
INSERT INTO `sys_logininfor` VALUES (747, 'admin', '60.1.167.118', '石家庄市', 'Firefox', 'Windows 10', '0', '登陆成功', '2023-10-19 09:00:17');
INSERT INTO `sys_logininfor` VALUES (748, 'admin', '117.114.175.206', '北京市', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2023-10-19 13:22:38');
INSERT INTO `sys_logininfor` VALUES (749, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-21 16:26:12');
INSERT INTO `sys_logininfor` VALUES (750, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-21 17:01:21');
INSERT INTO `sys_logininfor` VALUES (751, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-21 17:02:16');
INSERT INTO `sys_logininfor` VALUES (752, 'admin', '10.0.0.2', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', '登陆成功', '2023-10-23 19:15:37');
INSERT INTO `sys_logininfor` VALUES (753, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-24 16:56:41');
INSERT INTO `sys_logininfor` VALUES (754, 'admin', '10.0.0.2', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', '登陆成功', '2023-10-24 17:51:21');
INSERT INTO `sys_logininfor` VALUES (755, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-25 19:23:19');
INSERT INTO `sys_logininfor` VALUES (756, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-26 17:16:13');
INSERT INTO `sys_logininfor` VALUES (757, 'admin', '10.0.0.2', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', '登陆成功', '2023-10-27 01:53:05');
INSERT INTO `sys_logininfor` VALUES (758, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-28 23:30:39');
INSERT INTO `sys_logininfor` VALUES (759, 'admin', '192.168.100.103', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-29 18:41:17');
INSERT INTO `sys_logininfor` VALUES (760, 'admin', '192.168.100.103', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-29 18:42:33');
INSERT INTO `sys_logininfor` VALUES (761, 'admin', '192.168.100.103', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-29 18:42:45');
INSERT INTO `sys_logininfor` VALUES (762, 'admin', '192.168.100.103', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-29 19:44:42');
INSERT INTO `sys_logininfor` VALUES (763, 'admin', '192.168.100.103', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-29 21:19:09');
INSERT INTO `sys_logininfor` VALUES (764, 'admin', '192.168.100.103', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-29 21:24:19');
INSERT INTO `sys_logininfor` VALUES (765, 'admin', '192.168.100.103', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-29 21:32:11');
INSERT INTO `sys_logininfor` VALUES (766, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-29 21:40:41');
INSERT INTO `sys_logininfor` VALUES (767, 'admin', '192.168.100.103', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-29 21:43:43');
INSERT INTO `sys_logininfor` VALUES (768, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2023-10-29 22:24:16');
INSERT INTO `sys_logininfor` VALUES (769, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-29 22:24:29');
INSERT INTO `sys_logininfor` VALUES (770, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-29 22:34:47');
INSERT INTO `sys_logininfor` VALUES (771, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-29 22:43:59');
INSERT INTO `sys_logininfor` VALUES (772, 'admin', '192.168.100.103', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-30 19:10:26');
INSERT INTO `sys_logininfor` VALUES (773, 'admin', '192.168.100.103', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-30 20:14:41');
INSERT INTO `sys_logininfor` VALUES (774, 'admin', '192.168.100.103', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-10-30 21:20:08');
INSERT INTO `sys_logininfor` VALUES (775, 'admin', '10.0.0.2', '', 'Chrome', 'Android 12', '0', '登陆成功', '2023-11-01 15:57:27');
INSERT INTO `sys_logininfor` VALUES (776, 'admin', '10.0.0.2', '', 'Firefox', 'Windows 10', '0', '登陆成功', '2023-11-02 02:22:36');
INSERT INTO `sys_logininfor` VALUES (777, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-03 15:12:11');
INSERT INTO `sys_logininfor` VALUES (778, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-04 16:25:01');
INSERT INTO `sys_logininfor` VALUES (779, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-07 21:59:24');
INSERT INTO `sys_logininfor` VALUES (780, 'admin', '10.0.0.2', '', 'Firefox', 'Windows 10', '0', '登陆成功', '2023-11-08 14:28:42');
INSERT INTO `sys_logininfor` VALUES (781, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2023-11-08 23:01:49');
INSERT INTO `sys_logininfor` VALUES (782, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-08 23:02:13');
INSERT INTO `sys_logininfor` VALUES (783, 'admin', '10.0.0.2', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', '登陆成功', '2023-11-10 22:39:30');
INSERT INTO `sys_logininfor` VALUES (784, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-12 14:44:52');
INSERT INTO `sys_logininfor` VALUES (785, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-13 15:08:57');
INSERT INTO `sys_logininfor` VALUES (786, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-13 16:55:28');
INSERT INTO `sys_logininfor` VALUES (787, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-13 17:25:59');
INSERT INTO `sys_logininfor` VALUES (788, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-14 13:42:01');
INSERT INTO `sys_logininfor` VALUES (789, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-14 13:43:05');
INSERT INTO `sys_logininfor` VALUES (790, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-14 13:45:34');
INSERT INTO `sys_logininfor` VALUES (791, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-17 20:13:57');
INSERT INTO `sys_logininfor` VALUES (792, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-17 22:50:28');
INSERT INTO `sys_logininfor` VALUES (793, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-17 23:10:22');
INSERT INTO `sys_logininfor` VALUES (794, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-17 23:17:54');
INSERT INTO `sys_logininfor` VALUES (795, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-17 23:18:04');
INSERT INTO `sys_logininfor` VALUES (796, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-17 23:18:49');
INSERT INTO `sys_logininfor` VALUES (797, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-17 23:19:25');
INSERT INTO `sys_logininfor` VALUES (798, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-17 23:21:05');
INSERT INTO `sys_logininfor` VALUES (799, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-17 23:24:18');
INSERT INTO `sys_logininfor` VALUES (800, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-17 23:47:33');
INSERT INTO `sys_logininfor` VALUES (801, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-17 23:48:50');
INSERT INTO `sys_logininfor` VALUES (802, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-18 00:02:28');
INSERT INTO `sys_logininfor` VALUES (803, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-18 00:07:36');
INSERT INTO `sys_logininfor` VALUES (804, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-18 10:21:28');
INSERT INTO `sys_logininfor` VALUES (805, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-18 12:09:41');
INSERT INTO `sys_logininfor` VALUES (806, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-18 16:50:40');
INSERT INTO `sys_logininfor` VALUES (807, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-18 17:50:04');
INSERT INTO `sys_logininfor` VALUES (808, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-18 17:50:51');
INSERT INTO `sys_logininfor` VALUES (809, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-18 17:51:06');
INSERT INTO `sys_logininfor` VALUES (810, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-18 17:55:14');
INSERT INTO `sys_logininfor` VALUES (811, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-18 20:41:47');
INSERT INTO `sys_logininfor` VALUES (812, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-18 21:08:30');
INSERT INTO `sys_logininfor` VALUES (813, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-18 21:43:52');
INSERT INTO `sys_logininfor` VALUES (814, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-18 21:46:02');
INSERT INTO `sys_logininfor` VALUES (815, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2023-11-18 22:35:07');
INSERT INTO `sys_logininfor` VALUES (816, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2023-11-18 22:35:39');
INSERT INTO `sys_logininfor` VALUES (817, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-18 22:35:50');
INSERT INTO `sys_logininfor` VALUES (818, 'admin', '111.194.209.161', '北京市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-19 00:47:01');
INSERT INTO `sys_logininfor` VALUES (819, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-19 00:47:55');
INSERT INTO `sys_logininfor` VALUES (820, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-19 00:51:20');
INSERT INTO `sys_logininfor` VALUES (821, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-19 10:52:10');
INSERT INTO `sys_logininfor` VALUES (822, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-19 10:52:22');
INSERT INTO `sys_logininfor` VALUES (823, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-19 10:52:34');
INSERT INTO `sys_logininfor` VALUES (824, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-19 10:53:00');
INSERT INTO `sys_logininfor` VALUES (825, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-19 10:53:34');
INSERT INTO `sys_logininfor` VALUES (826, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-19 10:54:19');
INSERT INTO `sys_logininfor` VALUES (827, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-19 11:23:29');
INSERT INTO `sys_logininfor` VALUES (828, 'admin1', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2023-11-19 11:34:30');
INSERT INTO `sys_logininfor` VALUES (829, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-19 11:34:41');
INSERT INTO `sys_logininfor` VALUES (830, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-19 12:25:34');
INSERT INTO `sys_logininfor` VALUES (831, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-19 13:16:07');
INSERT INTO `sys_logininfor` VALUES (832, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-19 13:27:47');
INSERT INTO `sys_logininfor` VALUES (833, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-19 14:28:01');
INSERT INTO `sys_logininfor` VALUES (834, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-19 19:29:49');
INSERT INTO `sys_logininfor` VALUES (835, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-20 14:58:00');
INSERT INTO `sys_logininfor` VALUES (836, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-20 20:24:08');
INSERT INTO `sys_logininfor` VALUES (837, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-20 23:33:57');
INSERT INTO `sys_logininfor` VALUES (838, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-20 23:38:06');
INSERT INTO `sys_logininfor` VALUES (839, 'admin', '112.66.25.208', '儋州市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-21 20:40:24');
INSERT INTO `sys_logininfor` VALUES (840, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-21 22:43:40');
INSERT INTO `sys_logininfor` VALUES (841, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-22 09:28:27');
INSERT INTO `sys_logininfor` VALUES (842, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-22 10:50:05');
INSERT INTO `sys_logininfor` VALUES (843, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-22 23:21:49');
INSERT INTO `sys_logininfor` VALUES (844, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-23 11:08:13');
INSERT INTO `sys_logininfor` VALUES (845, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-23 15:27:19');
INSERT INTO `sys_logininfor` VALUES (846, 'admin', '10.0.0.2', '', 'Safari', 'Intel Mac OS X 10_15_7', '0', '登陆成功', '2023-11-23 17:26:30');
INSERT INTO `sys_logininfor` VALUES (847, 'admin', '60.169.85.7', '芜湖市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-24 08:27:29');
INSERT INTO `sys_logininfor` VALUES (848, 'awdaawd', '119.4.174.254', '成都市', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2023-11-24 15:01:50');
INSERT INTO `sys_logininfor` VALUES (849, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-24 15:21:08');
INSERT INTO `sys_logininfor` VALUES (850, 'admin', '116.113.44.250', '呼和浩特市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-24 15:24:40');
INSERT INTO `sys_logininfor` VALUES (851, 'admin', '116.113.44.250', '呼和浩特市', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-24 15:25:10');
INSERT INTO `sys_logininfor` VALUES (852, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-24 18:23:13');
INSERT INTO `sys_logininfor` VALUES (853, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-24 19:11:26');
INSERT INTO `sys_logininfor` VALUES (854, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-24 21:43:46');
INSERT INTO `sys_logininfor` VALUES (855, 'admin', '10.0.0.2', '', 'Safari', 'Intel Mac OS X 10_15_7', '0', '登陆成功', '2023-11-25 01:18:26');
INSERT INTO `sys_logininfor` VALUES (856, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-26 16:00:20');
INSERT INTO `sys_logininfor` VALUES (857, 'admin', '192.168.230.1', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-26 17:06:57');
INSERT INTO `sys_logininfor` VALUES (858, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-27 22:57:19');
INSERT INTO `sys_logininfor` VALUES (859, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-28 11:33:09');
INSERT INTO `sys_logininfor` VALUES (860, 'admin', '10.0.0.2', '', 'Safari', 'Intel Mac OS X 10_15_7', '0', '登陆成功', '2023-11-28 14:44:00');
INSERT INTO `sys_logininfor` VALUES (861, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-29 00:22:21');
INSERT INTO `sys_logininfor` VALUES (862, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2023-11-29 02:02:15');
INSERT INTO `sys_logininfor` VALUES (863, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-29 02:02:27');
INSERT INTO `sys_logininfor` VALUES (864, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-29 12:45:37');
INSERT INTO `sys_logininfor` VALUES (865, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-30 11:55:07');
INSERT INTO `sys_logininfor` VALUES (866, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-30 14:14:24');
INSERT INTO `sys_logininfor` VALUES (867, 'admin', '10.0.0.2', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', '登陆成功', '2023-11-30 15:22:26');
INSERT INTO `sys_logininfor` VALUES (868, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-30 19:58:36');
INSERT INTO `sys_logininfor` VALUES (869, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-11-30 22:03:35');
INSERT INTO `sys_logininfor` VALUES (870, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-12-01 15:08:34');
INSERT INTO `sys_logininfor` VALUES (871, 'admin', '10.0.0.2', '', 'Chrome', 'Windows 10', '0', '登陆成功', '2023-12-02 14:35:24');
INSERT INTO `sys_logininfor` VALUES (872, 'admin', '104.245.96.153', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-05 01:17:10');
INSERT INTO `sys_logininfor` VALUES (873, 'admin', '125.229.210.185', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-05 02:28:14');
INSERT INTO `sys_logininfor` VALUES (874, 'admin', '125.229.210.185', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-05 02:33:14');
INSERT INTO `sys_logininfor` VALUES (875, 'admin', '39.188.150.29', '嘉兴市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-05 08:29:24');
INSERT INTO `sys_logininfor` VALUES (876, 'admin', '101.69.227.130', '杭州市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-05 08:36:46');
INSERT INTO `sys_logininfor` VALUES (877, 'admin', '113.84.33.20', '汕头市', 'Chrome', 'Android 10', '0', 'login success', '2024-11-05 10:21:29');
INSERT INTO `sys_logininfor` VALUES (878, 'admin', '103.166.96.133', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-05 10:47:54');
INSERT INTO `sys_logininfor` VALUES (879, 'admin', '112.51.42.34', '厦门市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-05 16:17:11');
INSERT INTO `sys_logininfor` VALUES (880, 'admin', '117.182.43.63', '桂林市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-05 17:16:41');
INSERT INTO `sys_logininfor` VALUES (881, 'admin', '119.130.105.145', '广州市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-06 01:50:50');
INSERT INTO `sys_logininfor` VALUES (882, 'admin', '164.52.53.230', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-06 02:21:39');
INSERT INTO `sys_logininfor` VALUES (883, 'admin', '36.110.95.4', '北京市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-06 02:54:39');
INSERT INTO `sys_logininfor` VALUES (884, 'admin', '45.90.58.21', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-06 05:36:32');
INSERT INTO `sys_logininfor` VALUES (885, 'admin', '139.162.11.71', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-07 02:55:54');
INSERT INTO `sys_logininfor` VALUES (886, 'admin', '183.14.134.191', '深圳市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-07 05:49:55');
INSERT INTO `sys_logininfor` VALUES (887, 'admin', '38.207.136.2', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-07 08:20:53');
INSERT INTO `sys_logininfor` VALUES (888, 'admin', '153.3.178.113', '南京市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-07 09:29:24');
INSERT INTO `sys_logininfor` VALUES (889, 'admin', '58.216.162.130', '常州市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-08 01:24:18');
INSERT INTO `sys_logininfor` VALUES (890, 'admin', '171.43.146.38', '武汉市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-08 02:02:45');
INSERT INTO `sys_logininfor` VALUES (891, 'admin', '153.3.178.113', '南京市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-08 06:42:50');
INSERT INTO `sys_logininfor` VALUES (892, 'admin', '183.209.150.244', '南京市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-08 16:11:10');
INSERT INTO `sys_logininfor` VALUES (893, 'admin', '103.167.135.120', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-09 08:00:41');
INSERT INTO `sys_logininfor` VALUES (894, 'admin', '112.97.201.63', '东莞市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-09 12:22:29');
INSERT INTO `sys_logininfor` VALUES (895, 'admin', '144.24.13.34', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-11 03:06:18');
INSERT INTO `sys_logininfor` VALUES (896, 'admin', '171.223.208.204', '成都市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-11 07:34:08');
INSERT INTO `sys_logininfor` VALUES (897, 'admin', '125.121.102.122', '杭州市', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-11 11:57:29');
INSERT INTO `sys_logininfor` VALUES (898, 'admin', '39.188.155.2', '嘉兴市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-11 14:56:02');
INSERT INTO `sys_logininfor` VALUES (899, 'admin', '61.140.242.88', '广州市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-12 03:25:14');
INSERT INTO `sys_logininfor` VALUES (900, 'admin', '1.206.9.200', '六盘水市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-12 10:47:59');
INSERT INTO `sys_logininfor` VALUES (901, 'admin', '183.250.153.98', '厦门市', 'Android', 'Android 12', '0', 'login success', '2024-11-12 13:46:39');
INSERT INTO `sys_logininfor` VALUES (902, 'admin', '222.172.135.118', '昆明市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-13 08:18:37');
INSERT INTO `sys_logininfor` VALUES (903, 'admin', '119.39.129.93', '株洲市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-13 09:57:24');
INSERT INTO `sys_logininfor` VALUES (904, 'admin', '148.135.125.178', '', 'Safari', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-14 01:47:28');
INSERT INTO `sys_logininfor` VALUES (905, 'admin', '43.132.141.20', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-14 07:55:37');
INSERT INTO `sys_logininfor` VALUES (906, 'admin', '60.249.1.75', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-14 09:35:53');
INSERT INTO `sys_logininfor` VALUES (907, 'admin', '50.116.4.142', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-14 16:34:18');
INSERT INTO `sys_logininfor` VALUES (908, 'admin', '118.163.200.38', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-15 04:01:08');
INSERT INTO `sys_logininfor` VALUES (909, 'admin', '220.246.101.223', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-15 08:36:04');
INSERT INTO `sys_logininfor` VALUES (910, 'admin', '175.2.114.86', '娄底市', 'Chrome', 'CrOS x86_64 14541.0.0', '0', 'login success', '2024-11-15 11:57:16');
INSERT INTO `sys_logininfor` VALUES (911, 'admin', '122.245.148.18', '宁波市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-16 02:34:16');
INSERT INTO `sys_logininfor` VALUES (912, 'admin', '65.75.222.18', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-16 06:15:14');
INSERT INTO `sys_logininfor` VALUES (913, 'admin', '212.107.30.70', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-16 15:31:21');
INSERT INTO `sys_logininfor` VALUES (914, 'admin', '212.107.30.70', '', 'Safari', 'CPU iPhone OS 17_6_1 like Mac OS X', '0', 'login success', '2024-11-16 15:34:04');
INSERT INTO `sys_logininfor` VALUES (915, 'admin', '182.88.56.127', '南宁市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-17 08:14:07');
INSERT INTO `sys_logininfor` VALUES (916, 'admin', '182.255.32.50', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-17 08:21:33');
INSERT INTO `sys_logininfor` VALUES (917, 'admin', '182.255.32.50', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-17 08:22:15');
INSERT INTO `sys_logininfor` VALUES (918, 'admin', '116.23.220.114', '广州市', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-23 10:36:21');
INSERT INTO `sys_logininfor` VALUES (919, 'admin', '165.154.72.189', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-11-23 12:14:12');
INSERT INTO `sys_logininfor` VALUES (920, 'admin', '165.154.72.189', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-11-23 12:14:17');
INSERT INTO `sys_logininfor` VALUES (921, 'admin', '165.154.72.189', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-23 12:14:22');
INSERT INTO `sys_logininfor` VALUES (922, 'admin', '103.229.54.131', '', 'Android', 'Android 12', '0', 'login success', '2024-11-23 16:39:12');
INSERT INTO `sys_logininfor` VALUES (923, 'admin', '38.181.219.240', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-24 07:46:54');
INSERT INTO `sys_logininfor` VALUES (924, 'admin', '165.225.38.98', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-24 13:12:01');
INSERT INTO `sys_logininfor` VALUES (925, 'admin', '182.88.58.112', '南宁市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-24 13:36:33');
INSERT INTO `sys_logininfor` VALUES (926, 'admin', '175.2.113.174', '娄底市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-26 00:44:32');
INSERT INTO `sys_logininfor` VALUES (927, 'admin', '101.230.232.2', '上海市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-26 03:17:55');
INSERT INTO `sys_logininfor` VALUES (928, 'admin', '175.2.113.174', '娄底市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-26 03:33:43');
INSERT INTO `sys_logininfor` VALUES (929, 'admin', '45.90.208.186', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-26 05:18:30');
INSERT INTO `sys_logininfor` VALUES (930, 'admin', '58.241.94.197', '常州市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-26 06:29:00');
INSERT INTO `sys_logininfor` VALUES (931, 'admin', '47.242.148.45', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-26 14:34:03');
INSERT INTO `sys_logininfor` VALUES (932, 'admin', '183.162.254.246', '淮北市', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-27 02:55:37');
INSERT INTO `sys_logininfor` VALUES (933, 'admin', '147.78.177.161', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-27 06:38:24');
INSERT INTO `sys_logininfor` VALUES (934, 'admin', '43.198.151.157', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-27 08:44:41');
INSERT INTO `sys_logininfor` VALUES (935, 'admin', '38.148.227.39', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-27 12:56:09');
INSERT INTO `sys_logininfor` VALUES (936, 'admin', '117.188.6.17', '贵阳市', 'Safari', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-27 16:36:17');
INSERT INTO `sys_logininfor` VALUES (937, 'admin', '113.70.130.179', '佛山市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-28 01:23:30');
INSERT INTO `sys_logininfor` VALUES (938, 'admin', '45.153.5.169', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-28 09:37:15');
INSERT INTO `sys_logininfor` VALUES (939, 'admin', '86.98.35.130', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-28 10:52:08');
INSERT INTO `sys_logininfor` VALUES (940, 'admin', '38.150.11.214', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-29 01:29:57');
INSERT INTO `sys_logininfor` VALUES (941, 'admin', '222.209.94.142', '成都市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-29 03:52:04');
INSERT INTO `sys_logininfor` VALUES (942, 'admin', '116.80.100.82', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-29 07:23:12');
INSERT INTO `sys_logininfor` VALUES (943, 'admin', '123.151.209.114', '天津市', 'Chrome', 'Windows 10', '0', 'login success', '2024-11-29 08:50:29');
INSERT INTO `sys_logininfor` VALUES (944, 'admin', '163.53.18.106', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-11-29 09:02:07');
INSERT INTO `sys_logininfor` VALUES (945, 'admin', '59.33.170.207', '惠州市', 'Android', 'Android 10', '0', 'login success', '2024-11-29 16:37:28');
INSERT INTO `sys_logininfor` VALUES (946, 'admin', '39.65.150.143', '青岛市', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-12-01 08:39:02');
INSERT INTO `sys_logininfor` VALUES (947, 'admin', '18.183.55.122', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-02 00:17:08');
INSERT INTO `sys_logininfor` VALUES (948, 'admin', '178.236.36.200', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-02 04:41:14');
INSERT INTO `sys_logininfor` VALUES (949, 'admin', '58.19.99.221', '武汉市', 'python-requests', '', '0', 'login success', '2024-12-02 10:19:34');
INSERT INTO `sys_logininfor` VALUES (950, 'admin', '119.123.76.83', '深圳市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-02 14:25:23');
INSERT INTO `sys_logininfor` VALUES (951, 'admin', '104.234.233.236', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-02 15:11:42');
INSERT INTO `sys_logininfor` VALUES (952, 'admin', '58.19.99.221', '武汉市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-02 16:32:05');
INSERT INTO `sys_logininfor` VALUES (953, 'admin', '61.227.208.146', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-12-03 03:35:56');
INSERT INTO `sys_logininfor` VALUES (954, 'admin', '3.1.101.19', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-03 12:35:37');
INSERT INTO `sys_logininfor` VALUES (955, 'admin', '171.9.37.111', '驻马店市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-03 13:42:15');
INSERT INTO `sys_logininfor` VALUES (956, 'admin', '113.249.46.44', '重庆市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-03 14:06:04');
INSERT INTO `sys_logininfor` VALUES (957, 'admin', '120.230.84.64', '广州市', 'Safari', 'CPU iPhone OS 18_1_1 like Mac OS X', '0', 'login success', '2024-12-03 14:55:04');
INSERT INTO `sys_logininfor` VALUES (958, 'admin', '115.171.80.9', '北京市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-04 00:55:03');
INSERT INTO `sys_logininfor` VALUES (959, 'admin', '208.87.243.137', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-04 06:34:11');
INSERT INTO `sys_logininfor` VALUES (960, 'admin', '175.9.217.47', '长沙市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-04 07:34:13');
INSERT INTO `sys_logininfor` VALUES (961, 'admin', '13.115.211.126', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-05 06:12:58');
INSERT INTO `sys_logininfor` VALUES (962, 'admin', '14.153.7.198', '深圳市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-05 10:26:46');
INSERT INTO `sys_logininfor` VALUES (963, 'admin', '120.235.0.1', '湛江市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-05 15:08:01');
INSERT INTO `sys_logininfor` VALUES (964, 'admin', '218.85.122.10', '厦门市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-06 00:46:35');
INSERT INTO `sys_logininfor` VALUES (965, 'admin', '13.214.220.94', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-06 02:21:08');
INSERT INTO `sys_logininfor` VALUES (966, 'admin', '43.230.11.174', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-12-06 03:11:43');
INSERT INTO `sys_logininfor` VALUES (967, 'admin', '101.198.192.116', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-06 09:06:39');
INSERT INTO `sys_logininfor` VALUES (968, 'admin', '103.172.81.154', '', 'Firefox', 'Linux x86_64', '0', 'login success', '2024-12-06 12:19:46');
INSERT INTO `sys_logininfor` VALUES (969, 'admin', '118.114.94.226', '成都市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-06 12:22:50');
INSERT INTO `sys_logininfor` VALUES (970, 'admin', '219.145.109.245', '咸阳市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-08 00:35:09');
INSERT INTO `sys_logininfor` VALUES (971, 'admin', '38.207.136.2', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-12-08 07:20:49');
INSERT INTO `sys_logininfor` VALUES (972, 'admin', '5.31.223.89', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-08 19:13:50');
INSERT INTO `sys_logininfor` VALUES (973, 'admin', '211.141.83.125', '南昌市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-09 08:44:52');
INSERT INTO `sys_logininfor` VALUES (974, 'admin', '42.92.166.115', '兰州市', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-12-10 02:43:44');
INSERT INTO `sys_logininfor` VALUES (975, 'admin', '38.150.12.164', '', 'Firefox', 'Linux x86_64', '0', 'login success', '2024-12-10 07:54:51');
INSERT INTO `sys_logininfor` VALUES (976, 'admin', '109.166.36.159', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-10 09:46:30');
INSERT INTO `sys_logininfor` VALUES (977, 'admin', '47.236.14.136', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-10 14:42:26');
INSERT INTO `sys_logininfor` VALUES (978, 'admin', '47.236.14.136', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-10 14:42:29');
INSERT INTO `sys_logininfor` VALUES (979, 'admin', '47.236.14.136', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-10 14:42:27');
INSERT INTO `sys_logininfor` VALUES (980, 'admin', '47.236.14.136', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-10 14:42:30');
INSERT INTO `sys_logininfor` VALUES (981, 'admin', '47.236.14.136', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-10 14:42:30');
INSERT INTO `sys_logininfor` VALUES (982, 'admin', '47.236.14.136', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-10 14:42:31');
INSERT INTO `sys_logininfor` VALUES (983, 'admin', '120.86.87.55', '东莞市', 'Chrome', 'Android 14', '0', '账号或密码不正确', '2024-12-10 15:38:53');
INSERT INTO `sys_logininfor` VALUES (984, 'admin', '120.86.87.55', '东莞市', 'Chrome', 'Android 14', '0', 'login success', '2024-12-10 15:39:10');
INSERT INTO `sys_logininfor` VALUES (985, 'admin', '114.92.35.230', '上海市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-10 15:58:20');
INSERT INTO `sys_logininfor` VALUES (986, 'admin', '1.202.116.43', '北京市', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-12-11 02:26:46');
INSERT INTO `sys_logininfor` VALUES (987, 'admin', '171.8.232.208', '郑州市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-11 03:06:00');
INSERT INTO `sys_logininfor` VALUES (988, 'admin', '185.220.236.11', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-11 06:46:46');
INSERT INTO `sys_logininfor` VALUES (989, 'admin', '185.220.236.11', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-11 06:46:49');
INSERT INTO `sys_logininfor` VALUES (990, 'admin', '185.220.236.11', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-11 07:21:21');
INSERT INTO `sys_logininfor` VALUES (991, 'admin', '185.220.236.11', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-11 07:21:22');
INSERT INTO `sys_logininfor` VALUES (992, 'admin', '185.220.236.11', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-11 07:21:23');
INSERT INTO `sys_logininfor` VALUES (993, 'admin', '103.151.172.31', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-11 14:00:01');
INSERT INTO `sys_logininfor` VALUES (994, 'admin', '31.56.47.22', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-12-12 07:27:49');
INSERT INTO `sys_logininfor` VALUES (995, 'admin', '64.32.13.208', '', 'Chrome', 'Linux x86_64', '0', 'login success', '2024-12-13 07:03:40');
INSERT INTO `sys_logininfor` VALUES (996, 'admin', '195.245.229.139', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-12-13 07:25:58');
INSERT INTO `sys_logininfor` VALUES (997, 'admin', '223.100.29.154', '沈阳市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-13 08:06:34');
INSERT INTO `sys_logininfor` VALUES (998, 'admin', '45.150.166.3', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-13 08:27:17');
INSERT INTO `sys_logininfor` VALUES (999, 'admin', '27.38.193.157', '深圳市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-13 16:18:59');
INSERT INTO `sys_logininfor` VALUES (1000, 'admin', '38.181.77.139', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-14 20:11:11');
INSERT INTO `sys_logininfor` VALUES (1001, 'admin', '120.237.55.194', '茂名市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-15 02:34:59');
INSERT INTO `sys_logininfor` VALUES (1002, 'admin', '27.38.193.157', '深圳市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-15 06:23:09');
INSERT INTO `sys_logininfor` VALUES (1003, 'admin', '203.75.191.91', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-16 02:12:50');
INSERT INTO `sys_logininfor` VALUES (1004, 'admin', '43.154.26.196', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-17 01:57:56');
INSERT INTO `sys_logininfor` VALUES (1005, 'admin', '183.220.107.199', '成都市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-17 08:39:14');
INSERT INTO `sys_logininfor` VALUES (1006, 'admin', '183.220.107.199', '成都市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-17 08:49:08');
INSERT INTO `sys_logininfor` VALUES (1007, 'admin', '218.17.138.194', '深圳市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-17 10:00:08');
INSERT INTO `sys_logininfor` VALUES (1008, 'admin', '137.175.91.36', '', 'Chrome', 'Linux x86_64', '0', 'login success', '2024-12-18 04:15:15');
INSERT INTO `sys_logininfor` VALUES (1009, 'admin', '116.211.87.243', '武汉市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-18 05:32:25');
INSERT INTO `sys_logininfor` VALUES (1010, 'admin', '222.85.232.58', '贵阳市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-18 05:38:16');
INSERT INTO `sys_logininfor` VALUES (1011, 'admin', '117.28.187.60', '厦门市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-18 08:48:15');
INSERT INTO `sys_logininfor` VALUES (1012, 'admin', '59.61.82.172', '厦门市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-18 09:37:21');
INSERT INTO `sys_logininfor` VALUES (1013, 'admin', '61.136.109.109', '新乡市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-19 02:02:16');
INSERT INTO `sys_logininfor` VALUES (1014, 'admin', '27.186.193.239', '保定市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-19 03:10:43');
INSERT INTO `sys_logininfor` VALUES (1015, 'admin', '103.17.98.225', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-12-19 03:34:19');
INSERT INTO `sys_logininfor` VALUES (1016, 'admin', '154.64.227.57', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-19 11:10:29');
INSERT INTO `sys_logininfor` VALUES (1017, 'admin', '154.64.227.57', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-19 11:10:49');
INSERT INTO `sys_logininfor` VALUES (1018, 'admin', '54.254.158.238', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-19 14:24:39');
INSERT INTO `sys_logininfor` VALUES (1019, 'admin', '103.99.74.231', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-20 01:58:13');
INSERT INTO `sys_logininfor` VALUES (1020, 'admin', '103.99.74.231', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-20 01:59:36');
INSERT INTO `sys_logininfor` VALUES (1021, 'admin', '46.232.48.5', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-20 05:19:02');
INSERT INTO `sys_logininfor` VALUES (1022, 'admin', '196.247.24.151', '', 'Chrome', 'Intel Mac OS X 10_15_7', '0', 'login success', '2024-12-20 09:37:25');
INSERT INTO `sys_logininfor` VALUES (1023, 'admin', '219.78.196.66', '', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-20 13:55:54');
INSERT INTO `sys_logininfor` VALUES (1024, 'admin', '27.38.193.73', '深圳市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-20 14:02:40');
INSERT INTO `sys_logininfor` VALUES (1025, 'admin', '36.28.83.76', '杭州市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-21 06:26:22');
INSERT INTO `sys_logininfor` VALUES (1026, 'admin', '36.28.83.76', '杭州市', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-21 06:27:52');
INSERT INTO `sys_logininfor` VALUES (1027, 'admin', '36.28.83.76', '杭州市', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-21 06:27:57');
INSERT INTO `sys_logininfor` VALUES (1028, 'admin', '119.236.144.204', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-21 08:44:27');
INSERT INTO `sys_logininfor` VALUES (1029, 'admin', '119.236.144.204', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-21 08:44:33');
INSERT INTO `sys_logininfor` VALUES (1030, 'admin', '119.236.144.204', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-21 08:44:49');
INSERT INTO `sys_logininfor` VALUES (1031, 'admin', '223.160.128.24', '北京市', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-21 10:33:48');
INSERT INTO `sys_logininfor` VALUES (1032, 'admin', '223.160.128.24', '北京市', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-21 10:33:58');
INSERT INTO `sys_logininfor` VALUES (1033, 'admin', '223.160.128.24', '北京市', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-21 10:34:09');
INSERT INTO `sys_logininfor` VALUES (1034, 'admin', '223.160.128.24', '北京市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-21 10:34:17');
INSERT INTO `sys_logininfor` VALUES (1035, 'admin', '223.160.128.24', '北京市', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-21 10:35:14');
INSERT INTO `sys_logininfor` VALUES (1036, 'admin', '223.160.128.24', '北京市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-21 10:35:21');
INSERT INTO `sys_logininfor` VALUES (1037, 'admin', '119.236.144.204', '', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-21 12:56:19');
INSERT INTO `sys_logininfor` VALUES (1038, 'admin', '223.160.128.24', '北京市', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-21 14:02:06');
INSERT INTO `sys_logininfor` VALUES (1039, 'admin', '223.160.128.24', '北京市', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-21 14:02:14');
INSERT INTO `sys_logininfor` VALUES (1040, 'admin', '223.160.128.24', '北京市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-21 14:02:24');
INSERT INTO `sys_logininfor` VALUES (1041, 'admin', '223.160.128.24', '北京市', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-21 14:04:10');
INSERT INTO `sys_logininfor` VALUES (1042, 'admin', '223.160.128.24', '北京市', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-21 14:04:54');
INSERT INTO `sys_logininfor` VALUES (1043, 'admin', '223.160.128.24', '北京市', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-21 14:05:08');
INSERT INTO `sys_logininfor` VALUES (1044, 'admin', '223.160.128.24', '北京市', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2024-12-21 14:05:50');
INSERT INTO `sys_logininfor` VALUES (1045, 'admin', '223.160.128.24', '北京市', 'Chrome', 'Windows 10', '0', 'login success', '2024-12-21 14:07:06');
INSERT INTO `sys_logininfor` VALUES (1046, 'admin', '::1', '北京市', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2025-07-17 05:41:03');
INSERT INTO `sys_logininfor` VALUES (1047, 'admin', '::1', '北京市', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2025-07-17 05:41:08');
INSERT INTO `sys_logininfor` VALUES (1048, 'admin', '::1', '北京市', 'Chrome', 'Windows 10', '0', 'login success', '2025-07-17 05:41:15');
INSERT INTO `sys_logininfor` VALUES (1049, 'admin', '::1', '北京市', 'Chrome', 'Windows 10', '0', '账号或密码不正确', '2025-07-17 05:42:06');
INSERT INTO `sys_logininfor` VALUES (1050, 'admin', '::1', '北京市', 'Chrome', 'Windows 10', '0', 'login success', '2025-07-17 05:42:14');

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `menu_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `menu_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '菜单名称',
  `parent_id` bigint(20) NULL DEFAULT NULL COMMENT '父菜单ID',
  `order_num` int(11) NULL DEFAULT NULL COMMENT '显示顺序',
  `url` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '请求地址',
  `target` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '打开方式（menuItem页签 menuBlank新窗口）',
  `menu_type` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单类型（M目录 C菜单 F按钮）',
  `visible` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单状态（0显示 1隐藏）',
  `perms` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '权限标识',
  `icon` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单图标',
  `create_by` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '删除标记',
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1195 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '菜单权限表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, '系统管理', 0, 1, '#', '', 'M', '0', '', 'fa fa-gear', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统管理目录', '0', NULL);
INSERT INTO `sys_menu` VALUES (2, '系统监控', 0, 2, '#', '', 'M', '0', '', 'fa fa-video-camera', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统监控目录', '0', NULL);
INSERT INTO `sys_menu` VALUES (3, '系统工具', 0, 3, '#', 'menuItem', 'M', '0', '', 'fa fa-gavel', 'admin', '2024-07-08 02:06:56', 'admin', '2024-07-08 02:06:56', '系统工具目录', '0', NULL);
INSERT INTO `sys_menu` VALUES (4, '业务管理', 0, 1, '#', 'menuItem', 'M', '0', NULL, 'fa fa-newspaper-o', 'admin', '2019-12-02 16:39:09', 'admin', NULL, '', '0', NULL);
INSERT INTO `sys_menu` VALUES (5, '实例演示', 0, 5, '#', 'menuItem', 'M', '0', '', 'fa fa-address-card', 'admin', '2024-08-19 11:43:12', 'admin', '2024-08-19 11:43:12', '系统工具目录', '0', NULL);
INSERT INTO `sys_menu` VALUES (6, '表单演示', 5, 1, '#', 'menuItem', 'M', '0', '', 'fa fa-vcard-o', 'admin', '2024-07-09 12:34:09', 'admin', '2024-07-09 12:34:09', '表单演示', '0', NULL);
INSERT INTO `sys_menu` VALUES (7, '表格演示', 5, 2, '#', 'menuItem', 'M', '0', '', 'fa fa-book', 'admin', '2024-07-09 12:35:32', 'admin', '2024-07-09 12:35:32', '表格演示', '0', NULL);
INSERT INTO `sys_menu` VALUES (8, '弹框演示', 5, 3, '#', '', 'M', '0', '', '', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '弹框演示', '0', NULL);
INSERT INTO `sys_menu` VALUES (9, '操作演示', 5, 4, '#', 'menuItem', 'M', '0', '', '', 'admin', '2024-07-08 12:07:28', 'admin', '2024-07-08 12:07:28', '操作演示', '0', NULL);
INSERT INTO `sys_menu` VALUES (10, '报表演示', 5, 5, '#', 'menuItem', 'M', '0', '', '', 'admin', '2024-07-08 17:05:23', 'admin', '2024-07-08 17:05:23', '报表演示', '0', NULL);
INSERT INTO `sys_menu` VALUES (11, '图标演示', 5, 6, '#', 'menuItem', 'M', '0', '', 'fa fa-asterisk', 'admin', '2024-07-29 15:18:33', 'admin', '2024-07-29 15:18:33', '图标演示', '0', NULL);
INSERT INTO `sys_menu` VALUES (12, '栅格演示', 6, 2, 'demo/form/grid', '', 'C', '0', 'demo:grid:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (13, '下拉框', 6, 3, 'demo/form/select', '', 'C', '0', 'demo:select:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (14, '时间轴', 6, 4, 'demo/form/timeline', '', 'C', '0', 'demo:timeline:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (15, '基本表单', 6, 5, 'demo/form/basic', '', 'C', '0', 'demo:basic:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (16, '卡片列表', 6, 6, 'demo/form/cards', '', 'C', '0', 'demo:cards:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (17, '功能扩展', 6, 7, 'demo/form/jasny', '', 'C', '0', 'demo:jasny:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (18, '拖动排序', 6, 8, 'demo/form/sortable', '', 'C', '0', 'demo:sortable:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (19, '选项卡&面板', 6, 9, 'demo/form/tabs_panels', '', 'C', '0', 'demo:tabs_panels:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (20, '表单校验', 6, 10, 'demo/form/validate', '', 'C', '0', 'demo:validate:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (21, '表单向导', 6, 11, 'demo/form/wizard', '', 'C', '0', 'demo:wizard:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (22, '文件上传', 6, 12, 'demo/form/upload', '', 'C', '0', 'demo:upload:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (23, '日期和时间', 6, 13, 'demo/form/datetime', '', 'C', '0', 'demo:datetime:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (24, '富文本编辑器', 6, 14, 'demo/form/summernote', '', 'C', '0', 'demo:summernote:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (25, '左右互选', 6, 15, 'demo/form/duallistbox', '', 'C', '0', 'demo:duallistbox:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (27, '按钮演示', 6, 1, 'demo/form/button', 'menuItem', 'C', '0', 'demo:button:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', '', '2020-02-04 08:46:48', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (28, '数据汇总', 7, 2, 'demo/table/footer', '', 'C', '0', 'demo:footer:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (29, '组合表头', 7, 3, 'demo/table/groupHeader', '', 'C', '0', 'demo:groupHeader:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (31, '记住翻页', 7, 5, 'demo/table/remember', '', 'C', '0', 'demo:remember:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (32, '跳转指定页', 7, 6, 'demo/table/pageGo', '', 'C', '0', 'demo:pageGo:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (33, '查询参数', 7, 7, 'demo/table/params', '', 'C', '0', 'demo:params:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (35, '点击加载表格', 7, 9, 'demo/table/button', '', 'C', '0', 'demo:button:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (36, '表格冻结列', 7, 10, 'demo/table/fixedColumns', '', 'C', '0', 'demo:fixedColumns:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (37, '触发事件', 7, 11, 'demo/table/event', '', 'C', '0', 'demo:event:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (38, '细节视图', 7, 12, 'demo/table/detail', '', 'C', '0', 'demo:detail:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (39, '父子视图', 7, 13, 'demo/table/child', '', 'C', '0', 'demo:child:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (40, '图片预览', 7, 14, 'demo/table/image', '', 'C', '0', 'demo:image:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (41, '动态增删改查', 7, 15, 'demo/table/curd', '', 'C', '0', 'demo:curd:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (42, '表格拖曳', 7, 16, 'demo/table/recorder', '', 'C', '0', 'demo:recorder:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (43, '行内编辑', 7, 17, 'demo/table/editable', '', 'C', '0', 'demo:editable:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (44, '其它操作', 7, 18, 'demo/table/other', '', 'C', '0', 'demo:other:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (45, '查询条件', 7, 1, 'demo/table/search', '', 'C', '0', 'demo:search:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (46, '弹层组件', 8, 2, 'demo/modal/layer', '', 'C', '0', 'demo:layer:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (47, '弹层表格', 8, 3, 'demo/modal/table', '', 'C', '0', 'demo:table:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (48, '模态窗口', 8, 1, 'demo/modal/dialog', '', 'C', '0', 'demo:dialog:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (49, '其他操作', 9, 2, 'demo/operate/other', '', 'C', '0', 'demo:other:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (50, '表格操作', 9, 1, 'demo/operate/table', '', 'C', '0', 'demo:table:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (51, 'Peity', 10, 2, 'demo/report/metrics', '', 'C', '0', 'demo:metrics:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (52, 'SparkLine', 10, 3, 'demo/report/peity', '', 'C', '0', 'demo:peity:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (53, '图表组合', 10, 4, 'demo/report/sparkline', '', 'C', '0', 'demo:sparkline:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (54, '百度Echarts', 10, 1, 'demo/report/echarts', '', 'C', '0', 'demo:echarts:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (55, 'Glyphicons', 11, 2, 'demo/icon/glyphicons', '', 'C', '0', 'demo:glyphicons:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (56, 'Font Awesome', 11, 1, 'demo/icon/fontawesome', '', 'C', '0', 'demo:fontawesome:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (100, '用户管理', 1, 1, 'system/user', 'menuItem', 'C', '0', 'system:user:view', 'fa fa-user-o', 'admin', '2024-09-01 03:20:50', 'admin', '2024-09-01 03:20:50', '用户管理菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (101, '角色管理', 1, 2, 'system/role', 'menuItem', 'C', '0', 'system:role:view', 'fa fa-user-circle', 'admin', '2024-09-01 03:31:16', 'admin', '2024-09-01 03:31:16', '角色管理菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (102, '菜单管理', 1, 3, 'system/menu', 'menuItem', 'C', '0', 'system:menu:view', 'fa fa-map-o', 'admin', '2024-09-01 03:19:59', 'admin', '2024-09-01 03:19:59', '菜单管理菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (103, '组织管理', 1, 4, 'system/dept', 'menuItem', 'C', '0', 'system:dept:view', 'fa fa-object-group', 'admin', '2024-09-01 03:21:22', 'admin', '2024-09-01 03:21:22', '部门管理菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (104, '岗位管理', 1, 5, 'system/post', 'menuItem', 'C', '0', 'system:post:view', 'fa fa-vcard-o', 'admin', '2024-09-01 03:27:34', 'admin', '2024-09-01 03:27:34', '岗位管理菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (105, '字典管理', 1, 6, 'system/dict', 'menuItem', 'C', '0', 'system:dict:view', 'fa fa-newspaper-o', 'admin', '2024-09-01 03:26:28', 'admin', '2024-09-01 03:26:28', '字典管理菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (106, '参数设置', 1, 7, 'system/config', 'menuItem', 'C', '0', 'system:config:view', 'fa fa-asterisk', 'admin', '2024-09-01 03:26:54', 'admin', '2024-09-01 03:26:54', '参数设置菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (109, '在线用户', 2, 1, 'monitor/online', '', 'C', '0', 'monitor:online:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '在线用户菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (112, '服务监控', 2, 3, 'monitor/server', '', 'C', '0', 'monitor:server:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '服务监控菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (113, '表单构建', 3, 1, 'tool/build', '', 'C', '0', 'tool:build:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '表单构建菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (114, '代码生成', 3, 2, 'tool/gen', 'menuItem', 'M', '0', 'tool:gen:view', 'fa fa-sticky-note-o', 'admin', '2024-07-07 00:25:44', 'admin', '2024-07-07 00:25:44', '代码生成菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (115, '系统接口', 3, 3, 'tool/swagger', '', 'C', '0', 'tool:swagger:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统接口菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (500, '操作日志', 2, 4, 'monitor/operlog', '', 'C', '0', 'monitor:operlog:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '操作日志菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (501, '登录日志', 2, 5, 'monitor/logininfor', '', 'C', '0', 'monitor:logininfor:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '登录日志菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (1000, '用户查询', 100, 1, '#', '', 'F', '0', 'system:user:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1001, '用户新增', 100, 2, '#', '', 'F', '0', 'system:user:add', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1002, '用户修改', 100, 3, '#', '', 'F', '0', 'system:user:edit', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1003, '用户删除', 100, 4, '#', '', 'F', '0', 'system:user:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1004, '用户导出', 100, 5, '#', '', 'F', '0', 'system:user:export', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1005, '用户导入', 100, 6, '#', '', 'F', '0', 'system:user:import', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1006, '重置密码', 100, 7, '#', '', 'F', '0', 'system:user:resetPwd', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1007, '角色查询', 101, 1, '#', '', 'F', '0', 'system:role:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1008, '角色新增', 101, 2, '#', '', 'F', '0', 'system:role:add', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1009, '角色修改', 101, 3, '#', '', 'F', '0', 'system:role:edit', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1010, '角色删除', 101, 4, '#', '', 'F', '0', 'system:role:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1011, '角色导出', 101, 5, '#', '', 'F', '0', 'system:role:export', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1012, '菜单查询', 102, 1, '#', '', 'F', '0', 'system:menu:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1013, '菜单新增', 102, 2, '#', '', 'F', '0', 'system:menu:add', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1014, '菜单修改', 102, 3, '#', '', 'F', '0', 'system:menu:edit', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1015, '菜单删除', 102, 4, '#', '', 'F', '0', 'system:menu:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1016, '部门查询', 103, 1, '#', '', 'F', '0', 'system:dept:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1017, '部门新增', 103, 2, '#', '', 'F', '0', 'system:dept:add', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1018, '部门修改', 103, 3, '#', '', 'F', '0', 'system:dept:edit', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1019, '部门删除', 103, 4, '#', '', 'F', '0', 'system:dept:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1020, '岗位查询', 104, 1, '#', '', 'F', '0', 'system:post:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1021, '岗位新增', 104, 2, '#', '', 'F', '0', 'system:post:add', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1022, '岗位修改', 104, 3, '#', '', 'F', '0', 'system:post:edit', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1023, '岗位删除', 104, 4, '#', '', 'F', '0', 'system:post:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1024, '岗位导出', 104, 5, '#', '', 'F', '0', 'system:post:export', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1025, '字典查询', 105, 1, '#', '', 'F', '0', 'system:dict:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1026, '字典新增', 105, 2, '#', '', 'F', '0', 'system:dict:add', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1027, '字典修改', 105, 3, '#', '', 'F', '0', 'system:dict:edit', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1028, '字典删除', 105, 4, '#', '', 'F', '0', 'system:dict:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1029, '字典导出', 105, 5, '#', '', 'F', '0', 'system:dict:export', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1030, '参数查询', 106, 1, '#', '', 'F', '0', 'system:config:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1031, '参数新增', 106, 2, '#', '', 'F', '0', 'system:config:add', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1032, '参数修改', 106, 3, '#', '', 'F', '0', 'system:config:edit', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1033, '参数删除', 106, 4, '#', '', 'F', '0', 'system:config:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1034, '参数导出', 106, 5, '#', '', 'F', '0', 'system:config:export', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1039, '操作查询', 500, 1, '#', '', 'F', '0', 'monitor:operlog:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1040, '操作删除', 500, 2, '#', '', 'F', '0', 'monitor:operlog:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1041, '详细信息', 500, 3, '#', '', 'F', '0', 'monitor:operlog:detail', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1042, '日志导出', 500, 4, '#', '', 'F', '0', 'monitor:operlog:export', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1043, '登录查询', 501, 1, '#', '', 'F', '0', 'monitor:logininfor:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1044, '登录删除', 501, 2, '#', '', 'F', '0', 'monitor:logininfor:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1045, '日志导出', 501, 3, '#', '', 'F', '0', 'monitor:logininfor:export', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1046, '账户解锁', 501, 4, '#', '', 'F', '0', 'monitor:logininfor:unlock', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1047, '在线查询', 109, 1, '#', '', 'F', '0', 'monitor:online:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1048, '批量强退', 109, 2, '#', '', 'F', '0', 'monitor:online:batchForceLogout', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1049, '单条强退', 109, 3, '#', '', 'F', '0', 'monitor:online:forceLogout', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1057, '生成查询', 114, 1, '#', '', 'F', '0', 'tool:gen:list', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1058, '生成修改', 114, 2, '#', '', 'F', '0', 'tool:gen:edit', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1059, '生成删除', 114, 3, '#', '', 'F', '0', 'tool:gen:remove', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1060, '预览代码', 114, 4, '#', '', 'F', '0', 'tool:gen:preview', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1061, '生成代码', 114, 5, '#', '', 'F', '0', 'tool:gen:code', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1062, '字典详情', 105, 5, '#', '', 'F', '0', 'system:dict:detail', '#', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1064, '监控平台', 4, 1, 'https://lostvip.com', 'menuItem', 'C', '0', '/biz/task', 'fa fa-globe', '', '2024-08-22 02:54:41', '', '2024-08-22 02:54:41', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1065, '搜索自动补全', 6, 100, 'demo/form/autocomplete', 'menuItem', 'C', '', '', 'fa fa-sticky-note-o', '', '2023-09-18 10:15:49', '', '2023-09-18 10:17:03', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1066, '发电量重算', 4, 100, 'iot/hisData/toWizard', 'menuItem', 'C', '', '', 'fa fa-sticky-note-o', '', '2023-09-18 21:42:44', '', '2023-09-19 21:36:52', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1124, '计划任务', 4, 1, 'https://lostvip.com/crontab', 'menuItem', 'C', '0', 'biz:task:view', 'fa fa-calendar', 'admin', '2024-09-01 02:04:25', 'admin', '2024-09-01 02:04:25', '计划任务菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (1143, '自定义视图', 7, 100, '/demo/table/customView', 'menuItem', 'C', '0', '', 'fa fa-bookmark-o', '', '2024-07-14 04:55:36', '', '2024-07-14 04:55:36', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1144, '子表格', 7, 2, '/demo/table/subdata', 'menuItem', 'C', '0', '', 'fa fa-sticky-note-o', '', '2024-07-14 04:51:14', '', '2024-07-14 04:51:14', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1145, '设备管理', 4, 1, '/iot_dev/device/preListTree', 'menuItem', 'C', '0', '/biz/task', 'fa fa-sticky-note-o', 'admin', '2024-08-22 02:54:23', 'admin', '2024-08-22 02:54:23', 'device菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (1146, 'device新增', 1145, 2, '#', 'menuItem', 'F', '0', 'iot_dev:device:edit', '#', 'admin', '2024-08-22 02:55:26', 'admin', '2024-08-22 02:55:26', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1147, 'device查询', 1145, 1, '#', NULL, 'F', '0', 'iot_dev:device:list', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1148, 'device修改', 1145, 3, '#', NULL, 'F', '0', 'iot_dev:device:edit', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1149, 'device删除', 1145, 4, '#', NULL, 'F', '0', 'iot_dev:device:remove', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1150, 'device导出', 1145, 5, '#', NULL, 'F', '0', 'iot_dev:device:export', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1151, '产品管理', 4, 1, '/iot_dev/product', 'menuItem', 'C', '0', 'iot_dev:product:view', 'fa fa-sticky-note-o', 'admin', '2024-07-21 13:41:14', 'admin', '2024-07-21 13:41:14', 'product菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (1152, 'product新增', 1151, 2, '#', NULL, 'F', '0', 'iot_dev:product:add', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1153, 'product查询', 1151, 1, '#', NULL, 'F', '0', 'iot_dev:product:list', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1154, 'product修改', 1151, 3, '#', NULL, 'F', '0', 'iot_dev:product:edit', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1155, 'product删除', 1151, 4, '#', NULL, 'F', '0', 'iot_dev:product:remove', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1156, 'product导出', 1151, 5, '#', NULL, 'F', '0', 'iot_dev:product:export', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1157, 'realtime', 4, 1, '/iot_data/realtime', NULL, 'C', '0', 'iot_data:realtime:view', 'fa fa-sticky-note-o', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', 'realtime菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (1158, 'realtime新增', 1157, 2, '#', NULL, 'F', '0', 'iot_data:realtime:add', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1159, 'realtime查询', 1157, 1, '#', NULL, 'F', '0', 'iot_data:realtime:list', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1160, 'realtime修改', 1157, 3, '#', NULL, 'F', '0', 'iot_data:realtime:edit', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1161, 'realtime删除', 1157, 4, '#', NULL, 'F', '0', 'iot_data:realtime:remove', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1162, 'realtime导出', 1157, 5, '#', NULL, 'F', '0', 'iot_data:realtime:export', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1170, '定时任务', 2, 100, 'http://lostvip.com:8090/', 'menuItem', 'C', '0', '', 'fa fa-calendar-check-o', '', '2024-09-01 03:08:05', '', '2024-09-01 03:08:05', '', '0', 0);
INSERT INTO `sys_menu` VALUES (1171, '字典类型', 4, 1, '/robvi/type', NULL, 'C', '0', 'robvi:type:view', 'fa fa-sticky-note-o', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '字典类型菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (1172, '字典类型新增', 1171, 2, '#', NULL, 'F', '0', 'robvi:type:add', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1173, '字典类型查询', 1171, 1, '#', NULL, 'F', '0', 'robvi:type:list', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1174, '字典类型修改', 1171, 3, '#', NULL, 'F', '0', 'robvi:type:edit', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1175, '字典类型删除', 1171, 4, '#', NULL, 'F', '0', 'robvi:type:remove', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1176, '字典类型导出', 1171, 5, '#', NULL, 'F', '0', 'robvi:type:export', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1177, '字典类型', 4, 1, '/robvi/type', NULL, 'C', '0', 'robvi:type:view', 'fa fa-sticky-note-o', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '字典类型菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (1178, '字典类型新增', 1177, 2, '#', NULL, 'F', '0', 'robvi:type:add', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1179, '字典类型查询', 1177, 1, '#', NULL, 'F', '0', 'robvi:type:list', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1180, '字典类型修改', 1177, 3, '#', NULL, 'F', '0', 'robvi:type:edit', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1181, '字典类型删除', 1177, 4, '#', NULL, 'F', '0', 'robvi:type:remove', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1182, '字典类型导出', 1177, 5, '#', NULL, 'F', '0', 'robvi:type:export', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1183, '字典类型', 4, 1, '/robvi/type', NULL, 'C', '0', 'robvi:type:view', 'fa fa-sticky-note-o', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '字典类型菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (1184, '字典类型新增', 1183, 2, '#', NULL, 'F', '0', 'robvi:type:add', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1185, '字典类型查询', 1183, 1, '#', NULL, 'F', '0', 'robvi:type:list', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1186, '字典类型修改', 1183, 3, '#', NULL, 'F', '0', 'robvi:type:edit', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1187, '字典类型删除', 1183, 4, '#', NULL, 'F', '0', 'robvi:type:remove', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1188, '字典类型导出', 1183, 5, '#', NULL, 'F', '0', 'robvi:type:export', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1189, '字典类型', 4, 1, '/robvi/type', NULL, 'C', '0', 'robvi:type:view', 'fa fa-sticky-note-o', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '字典类型菜单', '0', NULL);
INSERT INTO `sys_menu` VALUES (1190, '字典类型新增', 1189, 2, '#', NULL, 'F', '0', 'robvi:type:add', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1191, '字典类型查询', 1189, 1, '#', NULL, 'F', '0', 'robvi:type:list', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1192, '字典类型修改', 1189, 3, '#', NULL, 'F', '0', 'robvi:type:edit', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1193, '字典类型删除', 1189, 4, '#', NULL, 'F', '0', 'robvi:type:remove', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);
INSERT INTO `sys_menu` VALUES (1194, '字典类型导出', 1189, 5, '#', NULL, 'F', '0', 'robvi:type:export', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '', '0', NULL);

-- ----------------------------
-- Table structure for sys_notice
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice`;
CREATE TABLE `sys_notice`  (
  `notice_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '公告ID',
  `notice_title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '公告标题',
  `notice_type` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '公告类型（1通知 2公告）',
  `notice_content` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '公告内容',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '公告状态（0正常 1关闭）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`notice_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '通知公告表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_notice
-- ----------------------------

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log`  (
  `oper_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '模块标题',
  `business_type` int(11) NULL DEFAULT NULL COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '方法名称',
  `request_method` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '请求方式',
  `operator_type` int(11) NULL DEFAULT NULL COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作人员',
  `dept_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '部门名称',
  `oper_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '请求URL',
  `oper_ip` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '主机地址',
  `oper_location` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作地点',
  `oper_param` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '请求参数',
  `json_result` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '返回参数',
  `status` int(11) NULL DEFAULT NULL COMMENT '操作状态（0正常 1异常）',
  `error_msg` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '错误消息',
  `oper_time` datetime(0) NULL DEFAULT NULL COMMENT '操作时间',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `create_by` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `del_flag` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 796 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '操作日志记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------
INSERT INTO `sys_oper_log` VALUES (491, '操作日志管理', 3, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/operlog/clean', '::1', '内网IP', '\"all\"', '{\"code\":0,\"msg\":\"操作成功\",\"data\":2,\"otype\":3}', 0, '', '2020-03-24 13:37:47', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (492, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 14:14:53', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (493, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 14:15:09', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (494, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":12}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [* * * * * *]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 14:33:58', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (495, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [* * * * * *]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 14:42:52', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (496, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":12}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [0/10 * * * * ?]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 14:51:19', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (497, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":12}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [0/10 * * * * ?]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 14:55:48', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (498, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [0/10 * * * * ?]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 14:57:27', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (499, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [0 30 * * * *]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:00:14', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (500, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [0 30 * * * *]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:04:08', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (501, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [*/5 * * * * *]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:05:22', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (502, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [*/5 * * * * *]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:08:47', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (503, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:12:22', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (504, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"任务已存在\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:14:10', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (505, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"任务已存在\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:15:57', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (506, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"任务已存在\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:15:58', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (507, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":12}', '{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:16:45', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (508, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:20:31', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (509, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"任务已存在\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:20:58', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (510, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:22:31', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (511, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"任务已存在\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:38:23', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (512, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:39:28', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (513, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 15:49:23', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (514, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 15:51:20', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (515, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"任务已存在\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:51:20', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (516, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 15:52:20', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (517, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 15:54:12', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (518, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 15:57:02', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (519, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 15:57:12', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (520, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"任务已存在\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:57:21', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (521, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:01:34', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (522, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:03:34', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (523, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:03:46', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (524, '定时任务管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/edit', '::1', '内网IP', '{\"JobName\":\"test1\",\"JobParams\":\"\",\"JobGroup\":\"DEFAULT\",\"JobId\":10,\"InvokeTarget\":\"test1\",\"CronExpression\":\"1 * * * * \",\"MisfirePolicy\":\"2\",\"Concurrent\":\"1\",\"Status\":\"0\",\"Remark\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2020-03-24 16:11:12', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (525, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:11:15', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (526, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:11:31', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (527, '定时任务管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/edit', '::1', '内网IP', '{\"JobName\":\"test1\",\"JobParams\":\"\",\"JobGroup\":\"DEFAULT\",\"JobId\":10,\"InvokeTarget\":\"test1\",\"CronExpression\":\"1 * * * * \",\"MisfirePolicy\":\"1\",\"Concurrent\":\"1\",\"Status\":\"0\",\"Remark\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2020-03-24 16:12:46', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (528, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:14:47', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (529, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:18:06', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (530, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:18:14', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (531, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"任务已存在\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 16:18:20', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (532, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:18:29', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (533, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:20:56', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (534, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:27:49', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (535, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:20:36', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (536, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:44:53', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (537, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:46:52', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (538, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:51:59', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (539, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:52:24', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (540, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:53:27', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (541, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:55:51', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (542, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:56:55', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (543, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:58:04', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (544, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:58:29', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (545, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:59:59', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (546, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 19:01:46', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (547, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 19:02:18', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (548, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 19:04:44', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (549, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 19:05:02', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (550, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 19:06:01', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (551, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 19:30:13', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (552, '生成代码', 3, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/remove', '::1', '内网IP', '{\"Ids\":\"37,38,39,40,41\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2020-03-24 22:09:41', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (553, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-24 22:09:45', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (554, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"sql: converting argument $1 type: unsupported type []string, a slice of string\",\"data\":null,\"otype\":1}', 1, '', '2020-03-24 22:09:45', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (555, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-24 22:09:49', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (556, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"sql: converting argument $1 type: unsupported type []string, a slice of string\",\"data\":null,\"otype\":1}', 1, '', '2020-03-24 22:09:49', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (557, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 16:27:02', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (558, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"sql: converting argument $1 type: unsupported type []string, a slice of string\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 16:27:02', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (559, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:35:07', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (560, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"sql: converting argument $1 type: unsupported type []string, a slice of string\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:35:07', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (561, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:40:25', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (562, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"sql: converting argument $1 type: unsupported type []string, a slice of string\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:40:25', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (563, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:40:29', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (564, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"sql: converting argument $1 type: unsupported type []string, a slice of string\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:40:29', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (565, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:41:19', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (566, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"sql: converting argument $1 type: unsupported type []string, a slice of string\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:41:19', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (567, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"sys_config\"}', '{\"code\":500,\"msg\":\"sql: converting argument $1 type: unsupported type []string, a slice of string\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:42:16', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (568, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"sys_config\"}', '{\"code\":500,\"msg\":\"Error 1054: Unknown column \'table_id\' in \'field list\'\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:43:50', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (569, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-26 19:38:33', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (570, '生成代码', 3, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/remove', '::1', '内网IP', '{\"Ids\":\"1\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2020-03-26 19:41:15', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (571, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-26 19:43:19', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (572, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-26 19:43:24', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (573, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-26 20:13:33', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (574, '生成代码', 3, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/remove', '::1', '内网IP', '{\"Ids\":\"2,3,4\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2020-03-26 20:35:09', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (575, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-26 20:35:20', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (576, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 21:32:29', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (577, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 21:33:36', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (578, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 21:35:11', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (579, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 21:36:12', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (580, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 21:37:19', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (581, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 21:37:47', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (582, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 21:42:54', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (583, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 21:49:13', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (584, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 22:11:02', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (585, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 22:12:45', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (586, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 22:17:06', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (587, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 22:17:57', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (588, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-26 22:18:46', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (589, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-26 22:19:02', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (590, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:27:16', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (591, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:28:30', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (592, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:28:43', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (593, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:29:05', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (594, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:29:39', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (595, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:36:59', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (596, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":500,\"msg\":\"json: cannot unmarshal string into Go struct field Entity.column_id of type int64\",\"data\":null,\"otype\":2}', 1, '', '2020-03-27 09:44:01', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (597, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:44:52', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (598, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":500,\"msg\":\"json: cannot unmarshal number into Go struct field Entity.is_insert of type string\",\"data\":null,\"otype\":2}', 1, '', '2020-03-27 09:46:43', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (599, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:49:14', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (600, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:49:30', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (601, '生成代码', 0, 'GET', 'GET', 1, 'admin', '运维部门', '/tool/gen/genCode', '::1', '内网IP', '{\"tableId\":5}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-27 11:12:03', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (602, '生成代码', 0, 'GET', 'GET', 1, 'admin', '运维部门', '/tool/gen/genCode', '::1', '内网IP', '{\"tableId\":5}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-27 11:16:07', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (603, '生成代码', 0, 'GET', 'GET', 1, 'admin', '运维部门', '/tool/gen/genCode', '::1', '内网IP', '{\"tableId\":5}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-27 11:17:39', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (604, '生成代码', 0, 'GET', 'GET', 1, 'admin', '运维部门', '/tool/gen/genCode', '::1', '内网IP', '{\"tableId\":5}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-27 11:19:14', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (605, '生成代码', 0, 'GET', 'GET', 1, 'admin', '运维部门', '/tool/gen/genCode', '::1', '内网IP', '{\"tableId\":5}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-27 12:03:03', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (606, '生成代码', 0, 'GET', 'GET', 1, 'admin', '运维部门', '/tool/gen/genCode', '::1', '内网IP', '{\"tableId\":5}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-27 21:01:05', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (607, '部门管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/dept/edit', '::1', '内网IP', '{\"DeptId\":112,\"ParentId\":100,\"DeptName\":\"4535345\",\"OrderNum\":1,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2020-03-28 20:57:32', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (608, '部门管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/dept/edit', '::1', '内网IP', '{\"DeptId\":112,\"ParentId\":100,\"DeptName\":\"3242343243243242\",\"OrderNum\":1,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2020-03-28 20:57:37', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (609, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '运维部门', '/system/menu/remove', '::1', '内网IP', '{\"id\":1035}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2020-05-05 14:18:01', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (610, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '运维部门', '/system/menu/remove', '::1', '内网IP', '{\"id\":1038}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2020-05-05 14:18:14', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (611, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '运维部门', '/system/menu/remove', '::1', '内网IP', '{\"id\":1037}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2020-05-05 14:18:22', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (612, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '运维部门', '/system/menu/remove', '::1', '内网IP', '{\"id\":1036}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2020-05-05 14:18:28', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (613, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '运维部门', '/system/menu/remove', '::1', '内网IP', '{\"id\":107}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2020-05-05 14:18:33', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (614, '菜单管理', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/add', '127.0.0.1', '内网IP', '{\"ParentId\":0,\"MenuType\":\"C\",\"MenuName\":\"baidu\",\"OrderNum\":1,\"Url\":\"http://www.baidu.com\",\"Icon\":\"fa fa-asl-interpreting\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1063,\"otype\":1}', 0, '', '2021-06-02 00:00:56', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (615, '角色管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/edit', '127.0.0.1', '内网IP', '{\"RoleId\":1,\"RoleName\":\"管理员\",\"RoleKey\":\"admin\",\"RoleSort\":\"1\",\"Status\":\"0\",\"Remark\":\"管理员\",\"MenuIds\":\"1,100,1001,1004,1005,101,1007,1008,1009,1010,1011,102,1012,1013,103,1016,1017,1018,1019,104,1020,1021,1022,1023,1024,105,1025,1026,1029,106,1030,1031,1034,4,2,109,1047,1048,1049,110,1050,1051,1052,1053,1054,1055,1056,112,500,1039,1040,1041,1042,501,1043,1044,1045,1046,3,113,114,1057,1058,1059,1060,1061,115\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2021-06-02 00:01:45', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (616, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1063,\"ParentId\":0,\"MenuType\":\"C\",\"MenuName\":\"baidu\",\"OrderNum\":1,\"Url\":\"http://www.baidu.com\",\"Icon\":\"fa fa-asl-interpreting\",\"Target\":\"menuBlank\",\"Perms\":\"\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2021-06-02 00:02:48', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (617, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1063,\"ParentId\":0,\"MenuType\":\"C\",\"MenuName\":\"baidu\",\"OrderNum\":1,\"Url\":\"http://www.baidu.com\",\"Icon\":\"fa fa-asl-interpreting\",\"Target\":\"menuBlank\",\"Perms\":\"tool:swagger:view\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2021-06-02 00:04:21', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (618, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1063,\"ParentId\":0,\"MenuType\":\"C\",\"MenuName\":\"baidu\",\"OrderNum\":1,\"Url\":\"http://www.baidu.com\",\"Icon\":\"fa fa-asl-interpreting\",\"Target\":\"menuBlank\",\"Perms\":\"tool:swagger:view\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2021-06-02 00:05:04', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (619, '用户强退', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/online/forceLogout', '::1', '内网IP', '{\"sessionId\":\"UPZMETVQLVEVUUOZ4CEJXCOQQ5BBRZKEBGV4KLRCR6F52RYGQGYA\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2021-06-19 23:09:27', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (620, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '::1', '内网IP', '{\"MenuId\":1063,\"ParentId\":0,\"MenuType\":\"C\",\"MenuName\":\"baidu\",\"OrderNum\":1,\"Url\":\"http://www.baidu.com\",\"Icon\":\"fa fa-asl-interpreting\",\"Target\":\"menuItem\",\"Perms\":\"tool:swagger:view\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2021-06-19 23:11:01', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (621, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '运维部门', '/system/menu/remove', '::1', '内网IP', '{\"id\":1063}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2021-06-19 23:11:55', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (622, '菜单管理', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/add', '::1', '内网IP', '{\"ParentId\":4,\"MenuType\":\"C\",\"MenuName\":\"baidu\",\"OrderNum\":1,\"Url\":\"http://www.baidu.com\",\"Icon\":\"fa fa-bank\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1064,\"otype\":1}', 0, '', '2021-06-19 23:12:45', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (623, '生成代码', 0, 'GET', 'GET', 1, 'admin', '运维部门', '/tool/gen/genCode', '::1', '内网IP', '{\"tableId\":5}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2021-06-19 23:15:27', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (624, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '::1', '内网IP', '{\"MenuId\":104,\"ParentId\":1,\"MenuType\":\"C\",\"MenuName\":\"岗位管理\",\"OrderNum\":5,\"Url\":\"/system/post\",\"Icon\":\"fa fa-balance-scale\",\"Target\":\"menuItem\",\"Perms\":\"system:post:view\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2021-06-19 23:27:38', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (625, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '::1', '内网IP', '{\"MenuId\":6,\"ParentId\":5,\"MenuType\":\"M\",\"MenuName\":\"表单演示\",\"OrderNum\":1,\"Url\":\"#\",\"Icon\":\"fa fa-asterisk\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2021-06-19 23:28:01', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (626, '角色管理', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/export', '::1', '内网IP', '{\"RoleName\":\"\",\"Status\":\"\",\"RoleKey\":\"\",\"DataScope\":\"\",\"BeginTime\":\"\",\"EndTime\":\"\",\"PageNum\":0,\"PageSize\":0,\"OrderByColumn\":\"\",\"IsAsc\":\"\"}', '{\"code\":0,\"msg\":\"1624120150167529600.xls\",\"data\":null,\"otype\":0}', 0, '', '2021-06-20 00:29:10', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (627, '参数管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/config/edit', '127.0.0.1', '内网IP', '{\"ConfigId\":1,\"ConfigName\":\"主框架页-默认皮肤样式名称\",\"ConfigKey\":\"sys.index.skinName\",\"ConfigValue\":\"skin-blue\",\"ConfigType\":\"Y\",\"Remark\":\"蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2021-06-20 13:51:55', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (628, '修改用户', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/user/edit', '127.0.0.1', '内网IP', '{\"UserId\":1,\"UserName\":\"超级管理员\",\"Phonenumber\":\"15888888881\",\"Email\":\"dd122221111d@163.com\",\"DeptId\":110,\"Sex\":\"0\",\"Status\":\"0\",\"RoleIds\":\"\",\"PostIds\":\"\",\"Remark\":\"管理员111111\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2021-06-20 13:52:46', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (629, '角色管理', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/export', '127.0.0.1', '内网IP', '{\"RoleName\":\"\",\"Status\":\"\",\"RoleKey\":\"\",\"DataScope\":\"\",\"BeginTime\":\"2022-11-18\",\"EndTime\":\"2022-11-20\",\"PageNum\":0,\"PageSize\":0,\"OrderByColumn\":\"\",\"IsAsc\":\"\"}', '{\"code\":0,\"msg\":\"1668755539757563400.xls\",\"data\":null,\"otype\":0}', 0, '', '2022-11-18 15:12:20', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (630, '角色管理', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/export', '127.0.0.1', '内网IP', '{\"RoleName\":\"\",\"Status\":\"\",\"RoleKey\":\"\",\"DataScope\":\"\",\"BeginTime\":\"2022-11-18\",\"EndTime\":\"2022-11-20\",\"PageNum\":0,\"PageSize\":0,\"OrderByColumn\":\"\",\"IsAsc\":\"\"}', '{\"code\":0,\"msg\":\"1668755548153720100.xls\",\"data\":null,\"otype\":0}', 0, '', '2022-11-18 15:12:28', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (631, '角色管理', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/export', '127.0.0.1', '内网IP', '{\"RoleName\":\"\",\"Status\":\"\",\"RoleKey\":\"\",\"DataScope\":\"\",\"BeginTime\":\"2022-11-18\",\"EndTime\":\"2022-11-20\",\"PageNum\":0,\"PageSize\":0,\"OrderByColumn\":\"\",\"IsAsc\":\"\"}', '{\"code\":0,\"msg\":\"1668755559891987300.xls\",\"data\":null,\"otype\":0}', 0, '', '2022-11-18 15:12:40', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (632, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":103,\"ParentId\":1,\"MenuType\":\"C\",\"MenuName\":\"部门管理\",\"OrderNum\":4,\"Url\":\"/system/dept\",\"Icon\":\"fa fa-bank\",\"Target\":\"menuItem\",\"Perms\":\"system:dept:view\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-14 16:54:48', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (633, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":103,\"ParentId\":1,\"MenuType\":\"C\",\"MenuName\":\"部门管理\",\"OrderNum\":4,\"Url\":\"/system/dept\",\"Icon\":\"#\",\"Target\":\"menuItem\",\"Perms\":\"system:dept:view\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-14 16:55:37', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (634, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":104,\"ParentId\":1,\"MenuType\":\"C\",\"MenuName\":\"岗位管理\",\"OrderNum\":5,\"Url\":\"/system/post\",\"Icon\":\"#\",\"Target\":\"menuItem\",\"Perms\":\"system:post:view\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-14 16:55:51', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (635, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1064,\"ParentId\":4,\"MenuType\":\"C\",\"MenuName\":\"baidu\",\"OrderNum\":1,\"Url\":\"https://www.robvi.com\",\"Icon\":\"fa fa-bank\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-14 22:29:22', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (636, '保存头像', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/user/profile/updateAvatar', '127.0.0.1', '内网IP', '{\"userid\":1}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2023-09-14 23:00:39', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (637, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1064,\"ParentId\":4,\"MenuType\":\"C\",\"MenuName\":\"baidu\",\"OrderNum\":1,\"Url\":\"https://www.robvi.com\",\"Icon\":\"fa fa-globe\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-14 23:01:56', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (638, '角色管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/edit', '127.0.0.1', '内网IP', '{\"RoleId\":3,\"RoleName\":\"培训员\",\"RoleKey\":\"dfgdfg\",\"RoleSort\":\"0\",\"Status\":\"0\",\"Remark\":\"\",\"MenuIds\":\"1,100,1001,1002,1003,1004,1005,1006,101,1007,1008,1009,1010,1011,102,1012,1013,1014,1015,103,1016,1017,1018,1019,104,1020,1021,1022,1023,1024,105,1025,1026,1027,1028,1029,106,1030,1031,1032,1033,1034,2,109,1047,1048,1049,110,1050,1051,1052,1053,1054,1055,1056,112,500,1039,1040,1041,1042,501,1043,1044,1045,1046,3,113,114,1057,1058,1059,1060,1061,115\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-14 23:14:51', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (639, '角色管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/edit', '127.0.0.1', '内网IP', '{\"RoleId\":3,\"RoleName\":\"厂家运维人员\",\"RoleKey\":\"oper\",\"RoleSort\":\"0\",\"Status\":\"0\",\"Remark\":\"\",\"MenuIds\":\"1,100,1001,1002,1003,1004,1005,1006,101,1007,1008,1009,1010,1011,102,1012,1013,1014,1015,103,1016,1017,1018,1019,104,1020,1021,1022,1023,1024,105,1025,1026,1027,1028,1029,106,1030,1031,1032,1033,1034,2,109,1047,1048,1049,110,1050,1051,1052,1053,1054,1055,1056,112,500,1039,1040,1041,1042,501,1043,1044,1045,1046,3,113,114,1057,1058,1059,1060,1061,115\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 10:59:21', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (640, '角色管理', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/authDataScope', '127.0.0.1', '内网IP', '{\"RoleId\":3,\"RoleName\":\"厂家运维人员\",\"RoleKey\":\"oper\",\"DataScope\":\"1\",\"DeptIds\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-09-15 10:59:47', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (641, '新增用户', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/system/user/add', '127.0.0.1', '内网IP', '{\"UserName\":\"qy1156010001\",\"Phonenumber\":\"18518276786\",\"Email\":\"uhvs@163.com\",\"LoginName\":\"qy1156010001\",\"Password\":\"admin123\",\"DeptId\":110,\"Sex\":\"0\",\"Status\":\"0\",\"RoleIds\":\"3\",\"PostIds\":\"\",\"Remark\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":2,\"otype\":1}', 0, '', '2023-09-15 11:01:47', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (642, '角色管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/edit', '127.0.0.1', '内网IP', '{\"RoleId\":3,\"RoleName\":\"厂家运维人员\",\"RoleKey\":\"oper\",\"RoleSort\":\"0\",\"Status\":\"0\",\"Remark\":\"\",\"MenuIds\":\"4,1064\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:03:23', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (643, '部门管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":100,\"ParentId\":0,\"DeptName\":\"DPC\",\"OrderNum\":0,\"Leader\":\"admin\",\"Phone\":\"15888888888\",\"Email\":\"110@qq.com\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:34:01', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (644, '岗位管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/post/edit', '127.0.0.1', '内网IP', '{\"PostId\":2,\"PostName\":\"项目经理\",\"PostCode\":\"se\",\"PostSort\":2,\"Status\":\"0\",\"Remark\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:34:52', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (645, '岗位管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/post/edit', '127.0.0.1', '内网IP', '{\"PostId\":1,\"PostName\":\"经理\",\"PostCode\":\"manager\",\"PostSort\":1,\"Status\":\"0\",\"Remark\":\"manager\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:35:40', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (646, '岗位管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/post/edit', '127.0.0.1', '内网IP', '{\"PostId\":2,\"PostName\":\"运维\",\"PostCode\":\"operator\",\"PostSort\":2,\"Status\":\"0\",\"Remark\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:36:02', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (647, '部门管理', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/system/dept/add', '127.0.0.1', '内网IP', '{\"ParentId\":100,\"DeptName\":\"江苏铁塔\",\"OrderNum\":1,\"Leader\":\"江苏铁塔\",\"Phone\":\"\",\"Email\":\"\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}', 0, '', '2023-09-15 11:38:29', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (648, '部门管理', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/system/dept/add', '127.0.0.1', '内网IP', '{\"ParentId\":100,\"DeptName\":\"爱客\",\"OrderNum\":1,\"Leader\":\"\",\"Phone\":\"\",\"Email\":\"\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}', 0, '', '2023-09-15 11:39:06', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (649, '部门管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":113,\"ParentId\":100,\"DeptName\":\"江苏铁塔\",\"OrderNum\":2,\"Leader\":\"江苏铁塔\",\"Phone\":\"\",\"Email\":\"\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:39:15', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (650, '部门管理', 3, 'GET', 'GET', 1, 'admin', '运维部门', '/system/dept/remove', '127.0.0.1', '内网IP', '{\"id\":113}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2023-09-15 11:39:38', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (651, '部门管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":110,\"ParentId\":100,\"DeptName\":\"江苏铁塔\",\"OrderNum\":1,\"Leader\":\"\",\"Phone\":\"\",\"Email\":\"\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":500,\"msg\":\"部门名称已存在\",\"data\":null,\"otype\":2}', 1, '', '2023-09-15 11:40:27', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (652, '部门管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":110,\"ParentId\":100,\"DeptName\":\"江苏铁塔\",\"OrderNum\":1,\"Leader\":\"\",\"Phone\":\"\",\"Email\":\"\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":500,\"msg\":\"部门名称已存在\",\"data\":null,\"otype\":2}', 1, '', '2023-09-15 11:41:02', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (653, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":110,\"ParentId\":100,\"DeptName\":\"江苏铁塔\",\"OrderNum\":1,\"Leader\":\"\",\"Phone\":\"\",\"Email\":\"\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:51:28', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (654, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":112,\"ParentId\":100,\"DeptName\":\"动力源运维\",\"OrderNum\":1,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:54:14', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (655, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":112,\"ParentId\":100,\"DeptName\":\"动力源运维\",\"OrderNum\":100,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:54:34', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (656, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":112,\"ParentId\":100,\"DeptName\":\"动力源运维\",\"OrderNum\":0,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:54:43', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (657, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":110,\"ParentId\":100,\"DeptName\":\"江苏铁塔\",\"OrderNum\":100,\"Leader\":\"\",\"Phone\":\"\",\"Email\":\"\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:54:50', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (658, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":114,\"ParentId\":100,\"DeptName\":\"爱客\",\"OrderNum\":100,\"Leader\":\"\",\"Phone\":\"\",\"Email\":\"\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:55:13', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (659, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":112,\"ParentId\":100,\"DeptName\":\"动力源运维\",\"OrderNum\":10,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:55:24', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (660, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:39:59', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (661, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:40:00', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (662, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:40:21', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (663, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:40:22', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (664, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:40:40', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (665, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:40:40', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (666, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:43:15', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (667, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:43:16', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (668, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:47:40', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (669, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:47:40', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (670, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:47:46', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (671, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:47:46', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (672, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:48:25', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (673, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:48:26', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (674, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:48:56', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (675, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:48:57', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (676, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误tables未选中\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 19:03:47', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (677, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 19:03:48', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (678, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误tables未选中\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 19:10:59', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (679, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 19:10:59', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (680, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"sys_oper_log\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-09-17 19:16:00', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (681, '菜单管理', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/menu/add', '127.0.0.1', '内网IP', '{\"ParentId\":6,\"MenuType\":\"C\",\"MenuName\":\"搜索自动补全\",\"OrderNum\":100,\"Url\":\"/demo/form/autocomplete\",\"Icon\":\"\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1065,\"otype\":1}', 0, '', '2023-09-18 10:15:49', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (682, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1065,\"ParentId\":6,\"MenuType\":\"C\",\"MenuName\":\"搜索自动补全\",\"OrderNum\":0,\"Url\":\"/demo/form/autocomplete\",\"Icon\":\"\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-18 10:17:03', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (683, '菜单管理', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/menu/add', '127.0.0.1', '内网IP', '{\"ParentId\":4,\"MenuType\":\"C\",\"MenuName\":\"发电量重算\",\"OrderNum\":100,\"Url\":\"iot/hisData/toWizard\",\"Icon\":\"\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1066,\"otype\":1}', 0, '', '2023-09-18 21:42:44', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (684, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1066,\"ParentId\":4,\"MenuType\":\"C\",\"MenuName\":\"发电量重算\",\"OrderNum\":100,\"Url\":\"iot/hisData/toWizard\",\"Icon\":\"fa fa-bars\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-19 21:32:53', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (685, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1066,\"ParentId\":4,\"MenuType\":\"C\",\"MenuName\":\"发电量重算\",\"OrderNum\":100,\"Url\":\"iot/hisData/toWizard\",\"Icon\":\"fa fa-sticky-note-o\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-19 21:36:52', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (686, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1064,\"ParentId\":4,\"MenuType\":\"C\",\"MenuName\":\"监控平台\",\"OrderNum\":1,\"Url\":\"http://iciom.dpc.com.cn/\",\"Icon\":\"fa fa-globe\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-19 21:44:36', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (687, '修改用户', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/user/edit', '127.0.0.1', '内网IP', '{\"UserId\":2,\"UserName\":\"qy1156010001\",\"Phonenumber\":\"18518276786\",\"Email\":\"uhvs@163.com\",\"DeptId\":110,\"Sex\":\"0\",\"Status\":\"0\",\"RoleIds\":\"3\",\"PostIds\":\"\",\"Remark\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-19 22:01:29', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (688, '登陆日志管理', 3, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/monitor/logininfor/remove', '36.5.112.29', '合肥市', '{\"Ids\":\"651\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}', 0, '', '2023-09-21 11:43:20', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (689, '字典数据管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/system/dict/data/edit', '127.0.0.1', '内网IP', '{\"DictCode\":1,\"DictLabel\":\"男\",\"DictValue\":\"0\",\"DictType\":\"\",\"DictSort\":1,\"CssClass\":\"\",\"ListClass\":\"default\",\"IsDefault\":\"Y\",\"Status\":\"0\",\"Remark\":\"[[*{remark}]]\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-21 21:45:51', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (690, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":112,\"ParentId\":100,\"DeptName\":\"动力源运维\",\"OrderNum\":10,\"Leader\":\"lostvip\",\"Phone\":\"\",\"Email\":\"331641539@qq.com\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-21 21:49:43', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (691, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/menu/edit', '10.0.0.2', '', '{\"MenuId\":1064,\"ParentId\":4,\"MenuType\":\"C\",\"MenuName\":\"监控平台\",\"OrderNum\":1,\"Url\":\"https://lostvip.com\",\"Icon\":\"fa fa-globe\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-10-17 14:10:12', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (692, '岗位管理', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/post/add', '10.0.0.2', '', '{\"PostName\":\"xx\",\"PostCode\":\"sd\",\"PostSort\":3,\"Status\":\"0\",\"Remark\":\"\"}', '{\"code\":500,\"msg\":\"操作失败\",\"data\":3,\"otype\":1}', 1, '', '2023-10-21 17:02:59', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (693, '岗位管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/post/edit', '10.0.0.2', '', '{\"PostId\":2,\"PostName\":\"运维\",\"PostCode\":\"operator\",\"PostSort\":3,\"Status\":\"0\",\"Remark\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-10-21 17:03:15', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (694, '岗位管理', 3, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/post/remove', '10.0.0.2', '', '{\"Ids\":\"3\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2023-10-21 17:03:27', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (695, '岗位管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/post/edit', '10.0.0.2', '', '{\"PostId\":2,\"PostName\":\"运维\",\"PostCode\":\"operator\",\"PostSort\":2,\"Status\":\"0\",\"Remark\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-10-21 17:03:38', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (696, '生成代码', 3, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/remove', '192.168.100.103', '', '{\"Ids\":\"5,6\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2023-10-29 19:30:00', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (697, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/importTable', '192.168.100.103', '', '{\"tables\":\"his_patient\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-29 21:43:56', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (698, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/importTable', '192.168.100.103', '', '{\"tables\":\"his_patient\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-29 21:44:26', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (699, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/importTable', '192.168.100.103', '', '{\"tables\":\"his_patient\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-29 21:45:07', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (700, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/importTable', '192.168.100.103', '', '{\"tables\":\"his_patient\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-29 21:46:40', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (701, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/importTable', '192.168.100.103', '', '{\"tables\":\"his_patient\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-29 21:53:11', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (702, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/importTable', '192.168.100.103', '', '{\"tables\":\"his_patient\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-29 21:55:17', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (703, '生成代码', 3, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/remove', '192.168.100.103', '', '{\"Ids\":\"12\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2023-10-29 22:04:34', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (704, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/importTable', '192.168.100.103', '', '{\"tables\":\"his_patient\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-29 22:05:23', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (705, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/dept/edit', '10.0.0.2', '', '{\"DeptId\":112,\"ParentId\":100,\"DeptName\":\"公司运维\",\"OrderNum\":10,\"Leader\":\"lostvip\",\"Phone\":\"18788996255\",\"Email\":\"331641539@qq.com\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-10-29 22:30:16', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (706, '生成代码', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/edit', '192.168.100.103', '', '{\"tableName\":\"gen_table\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2023-10-30 19:15:33', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (707, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 19:17:13', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (708, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 19:36:31', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (709, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 19:41:41', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (710, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 19:42:32', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (711, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 19:43:06', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (712, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 19:48:25', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (713, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 19:57:49', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (714, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:04:56', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (715, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:05:21', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (716, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:08:31', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (717, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:14:52', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (718, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:19:06', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (719, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:22:39', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (720, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:27:28', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (721, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:29:05', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (722, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:44:11', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (723, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:48:28', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (724, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:48:44', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (725, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:51:51', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (726, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:56:53', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (727, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:57:09', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (728, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:20:16', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (729, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:20:39', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (730, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:21:34', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (731, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:21:48', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (732, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:28:21', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (733, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:32:31', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (734, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:32:46', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (735, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:34:06', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (736, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:34:17', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (737, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:36:18', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (738, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:36:27', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (739, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/menu/edit', '10.0.0.2', '', '{\"MenuId\":100,\"ParentId\":1,\"MenuType\":\"C\",\"MenuName\":\"用户管理\",\"OrderNum\":1,\"Url\":\"system/user\",\"Icon\":\"fa fa-address-book\",\"Target\":\"menuItem\",\"Perms\":\"system:user:view\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-11-04 16:49:17', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (740, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/importTable', '192.168.230.1', '', '{\"tables\":\"his_patient\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-11-18 10:23:30', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (741, '生成代码', 3, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/remove', '192.168.230.1', '', '{\"Ids\":\"14\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2023-11-18 12:15:15', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (742, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/system/menu/remove', '192.168.230.1', '', '{\"id\":1070}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-18 21:37:25', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (743, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/system/menu/remove', '192.168.230.1', '', '{\"id\":1069}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-18 21:37:38', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (744, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.230.1', '', '{\"tableId\":13}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-18 21:44:12', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (745, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/system/menu/remove', '192.168.230.1', '', '{\"id\":1068}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-18 21:55:18', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (746, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/system/menu/remove', '192.168.230.1', '', '{\"id\":1088}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-18 22:49:21', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (747, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/system/menu/remove', '192.168.230.1', '', '{\"id\":1100}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-18 22:50:26', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (748, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/system/menu/remove', '192.168.230.1', '', '{\"id\":1094}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-18 22:50:48', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (749, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/system/menu/remove', '192.168.230.1', '', '{\"id\":1106}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-19 13:23:57', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (750, '生成代码', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/edit', '192.168.230.1', '', '{\"tableName\":\"his_patient\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 1, '', '2023-11-19 14:10:53', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (751, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '192.168.230.1', '', '{\"tableId\":13}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-19 14:11:08', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (752, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '192.168.230.1', '', '{\"tableId\":13}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-19 14:25:36', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (753, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/system/menu/remove', '192.168.230.1', '', '{\"id\":1082}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-19 14:27:26', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (754, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/system/menu/remove', '192.168.230.1', '', '{\"id\":1118}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-19 14:27:32', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (755, '导出Excel', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/user/export', '10.0.0.2', '', '{\"LoginName\":\"\",\"Status\":\"\",\"Phonenumber\":\"\",\"BeginTime\":\"\",\"EndTime\":\"\",\"DeptId\":0,\"PageNum\":0,\"PageSize\":0,\"SortName\":\"create_time\",\"SortOrder\":\"desc\",\"Ancestors\":\"\",\"TenantId\":0}', '{\"code\":200,\"msg\":\"1700483125461160200.xls\",\"data\":null,\"otype\":0}', 1, '', '2023-11-20 20:25:25', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (756, '参数管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/config/edit', '10.0.0.2', '', '{\"ConfigId\":1,\"ConfigName\":\"主框架页-默认皮肤样式名称\",\"ConfigKey\":\"sys.index.skinName\",\"ConfigValue\":\"skin-red\",\"ConfigType\":\"Y\",\"Remark\":\"蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 1, '', '2023-11-23 11:09:44', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (757, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '192.168.230.1', '', '{\"tableId\":13}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-24 18:26:22', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (758, '生成代码', 3, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/remove', '192.168.230.1', '', '{\"Ids\":\"13\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-24 18:29:45', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (759, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '192.168.230.1', '', '{\"tables\":\"his_patient\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-24 18:30:08', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (760, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '192.168.230.1', '', '{\"tableId\":15}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-24 18:33:04', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (761, '生成代码', 3, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/remove', '192.168.230.1', '', '{\"Ids\":\"15\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-24 18:55:11', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (762, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '192.168.230.1', '', '{\"tables\":\"his_patient\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-24 18:55:48', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (763, '生成代码', 3, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/remove', '192.168.230.1', '', '{\"Ids\":\"16\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-24 18:57:34', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (764, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '192.168.230.1', '', '{\"tables\":\"his_patient\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-24 18:58:55', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (765, '生成代码', 3, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/remove', '192.168.230.1', '', '{\"Ids\":\"17\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-24 19:09:42', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (766, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '192.168.230.1', '', '{\"tables\":\"his_patient\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-24 19:10:23', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (767, '生成代码', 3, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/remove', '192.168.230.1', '', '{\"Ids\":\"18\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-24 19:14:10', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (768, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '192.168.230.1', '', '{\"tables\":\"his_patient\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-24 19:17:05', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (769, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '192.168.230.1', '', '{\"tables\":\"his_patient\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-24 19:17:09', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (770, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '192.168.230.1', '', '{\"tables\":\"dpc_task\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-24 19:20:49', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (771, '生成代码', 3, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/remove', '192.168.230.1', '', '{\"Ids\":\"19,20,22\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-24 19:23:00', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (772, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '192.168.230.1', '', '{\"tables\":\"his_patient\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-26 16:01:04', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (773, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '192.168.230.1', '', '{\"tableId\":23}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-26 16:01:12', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (774, '生成代码', 3, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/remove', '192.168.230.1', '', '{\"Ids\":\"23\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-26 16:09:05', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (775, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '192.168.230.1', '', '{\"tables\":\"his_patient\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-26 16:09:22', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (776, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '192.168.230.1', '', '{\"tableId\":24}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-26 16:09:29', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (777, '生成代码', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/edit', '192.168.230.1', '', '{\"tableName\":\"his_patient\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 1, '', '2023-11-26 16:12:47', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (778, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '192.168.230.1', '', '{\"tableId\":24}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-26 16:13:10', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (779, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '192.168.230.1', '', '{\"tables\":\"dpc_task\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-26 16:27:25', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (780, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '192.168.230.1', '', '{\"tableId\":25}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-26 16:27:49', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (781, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '192.168.230.1', '', '{\"tableId\":25}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-26 16:54:06', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (782, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '192.168.230.1', '', '{\"tableId\":25}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-26 17:06:38', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (783, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '192.168.230.1', '', '{\"tableId\":25}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-26 17:07:27', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (784, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '192.168.230.1', '', '{\"tableId\":25}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-26 17:16:13', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (785, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '192.168.230.1', '', '{\"tableId\":25}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-26 17:20:22', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (786, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '192.168.230.1', '', '{\"tableId\":25}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-26 17:24:48', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (787, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '192.168.230.1', '', '{\"tableId\":24}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-26 17:26:15', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (788, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/menu/edit', '192.168.230.1', '', '{\"MenuId\":1112,\"ParentId\":4,\"MenuType\":\"C\",\"MenuName\":\"患者基本信息\",\"OrderNum\":1,\"Url\":\"/biz/patient\",\"Icon\":\"fa fa-address-book-o\",\"Target\":\"menuItem\",\"Perms\":\"biz:patient:view\",\"Visible\":\"0\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 1, '', '2023-11-26 17:29:29', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (789, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '10.0.0.2', '', '{\"tableId\":25}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-27 22:57:54', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (790, '生成代码', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/edit', '10.0.0.2', '', '{\"tableName\":\"dpc_task\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 1, '', '2023-11-27 23:01:37', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (791, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '10.0.0.2', '', '{\"tableId\":25}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-27 23:01:41', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (792, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '10.0.0.2', '', '{\"tables\":\"sys_post\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-30 14:16:22', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (793, '参数管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/config/edit', '10.0.0.2', '', '{\"ConfigId\":3,\"ConfigName\":\"主框架页-侧边栏主题\",\"ConfigKey\":\"sys.index.sideTheme\",\"ConfigValue\":\"theme-blue\",\"ConfigType\":\"Y\",\"Remark\":\"深黑主题theme-dark，浅色主题theme-light，深蓝主题theme-blue\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 1, '', '2023-11-30 20:05:02', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (794, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '10.0.0.2', '', '{\"tables\":\"sys_dict_type\"}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-12-01 15:09:35', NULL, NULL, 0);
INSERT INTO `sys_oper_log` VALUES (795, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/tool/gen/genCode', '10.0.0.2', '', '{\"tableId\":27}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-12-01 15:09:40', NULL, NULL, 0);

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `post_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '岗位名称',
  `post_sort` int(11) NOT NULL COMMENT '显示顺序',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '删除标记',
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '岗位信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES (1, 'manager', '经理', 1, '0', 'admin', '2018-03-16 11:33:00', '', '2023-09-15 11:35:40', 'manager', 0, '0');
INSERT INTO `sys_post` VALUES (2, 'operator', '运维', 2, '0', 'admin', '2018-03-16 11:33:00', 'admin', '2023-10-21 17:03:37', '', 0, '0');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色名称',
  `role_key` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色权限字符串',
  `role_sort` int(11) NOT NULL COMMENT '显示顺序',
  `data_scope` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色状态（0正常 1停用）',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '删除标记',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  PRIMARY KEY (`role_id`) USING BTREE,
  UNIQUE INDEX `idx_roleKey`(`role_key`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '管理员', 'admin1', 1, '1', '0', '0', 'admin', '2024-12-18 05:34:22', '', '2024-12-18 05:34:22', '管理员', NULL);
INSERT INTO `sys_role` VALUES (3, '厂家运维人员', 'oper', 0, '1', '0', '0', 'admin', '2024-12-19 03:34:31', 'admin', '2024-12-19 03:34:31', '', NULL);

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept`  (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `dept_id` bigint(20) NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`, `dept_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色和部门关联表' ROW_FORMAT = Dynamic;

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
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色和菜单关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES (1, 1);
INSERT INTO `sys_role_menu` VALUES (1, 2);
INSERT INTO `sys_role_menu` VALUES (1, 3);
INSERT INTO `sys_role_menu` VALUES (1, 4);
INSERT INTO `sys_role_menu` VALUES (1, 5);
INSERT INTO `sys_role_menu` VALUES (1, 6);
INSERT INTO `sys_role_menu` VALUES (1, 7);
INSERT INTO `sys_role_menu` VALUES (1, 8);
INSERT INTO `sys_role_menu` VALUES (1, 9);
INSERT INTO `sys_role_menu` VALUES (1, 10);
INSERT INTO `sys_role_menu` VALUES (1, 11);
INSERT INTO `sys_role_menu` VALUES (1, 12);
INSERT INTO `sys_role_menu` VALUES (1, 13);
INSERT INTO `sys_role_menu` VALUES (1, 14);
INSERT INTO `sys_role_menu` VALUES (1, 15);
INSERT INTO `sys_role_menu` VALUES (1, 16);
INSERT INTO `sys_role_menu` VALUES (1, 17);
INSERT INTO `sys_role_menu` VALUES (1, 18);
INSERT INTO `sys_role_menu` VALUES (1, 19);
INSERT INTO `sys_role_menu` VALUES (1, 20);
INSERT INTO `sys_role_menu` VALUES (1, 21);
INSERT INTO `sys_role_menu` VALUES (1, 22);
INSERT INTO `sys_role_menu` VALUES (1, 23);
INSERT INTO `sys_role_menu` VALUES (1, 24);
INSERT INTO `sys_role_menu` VALUES (1, 25);
INSERT INTO `sys_role_menu` VALUES (1, 27);
INSERT INTO `sys_role_menu` VALUES (1, 28);
INSERT INTO `sys_role_menu` VALUES (1, 29);
INSERT INTO `sys_role_menu` VALUES (1, 31);
INSERT INTO `sys_role_menu` VALUES (1, 32);
INSERT INTO `sys_role_menu` VALUES (1, 33);
INSERT INTO `sys_role_menu` VALUES (1, 35);
INSERT INTO `sys_role_menu` VALUES (1, 36);
INSERT INTO `sys_role_menu` VALUES (1, 37);
INSERT INTO `sys_role_menu` VALUES (1, 38);
INSERT INTO `sys_role_menu` VALUES (1, 39);
INSERT INTO `sys_role_menu` VALUES (1, 40);
INSERT INTO `sys_role_menu` VALUES (1, 41);
INSERT INTO `sys_role_menu` VALUES (1, 42);
INSERT INTO `sys_role_menu` VALUES (1, 43);
INSERT INTO `sys_role_menu` VALUES (1, 44);
INSERT INTO `sys_role_menu` VALUES (1, 45);
INSERT INTO `sys_role_menu` VALUES (1, 46);
INSERT INTO `sys_role_menu` VALUES (1, 47);
INSERT INTO `sys_role_menu` VALUES (1, 48);
INSERT INTO `sys_role_menu` VALUES (1, 49);
INSERT INTO `sys_role_menu` VALUES (1, 50);
INSERT INTO `sys_role_menu` VALUES (1, 51);
INSERT INTO `sys_role_menu` VALUES (1, 52);
INSERT INTO `sys_role_menu` VALUES (1, 53);
INSERT INTO `sys_role_menu` VALUES (1, 54);
INSERT INTO `sys_role_menu` VALUES (1, 55);
INSERT INTO `sys_role_menu` VALUES (1, 56);
INSERT INTO `sys_role_menu` VALUES (1, 100);
INSERT INTO `sys_role_menu` VALUES (1, 101);
INSERT INTO `sys_role_menu` VALUES (1, 102);
INSERT INTO `sys_role_menu` VALUES (1, 103);
INSERT INTO `sys_role_menu` VALUES (1, 104);
INSERT INTO `sys_role_menu` VALUES (1, 105);
INSERT INTO `sys_role_menu` VALUES (1, 106);
INSERT INTO `sys_role_menu` VALUES (1, 109);
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
INSERT INTO `sys_role_menu` VALUES (1, 1057);
INSERT INTO `sys_role_menu` VALUES (1, 1058);
INSERT INTO `sys_role_menu` VALUES (1, 1059);
INSERT INTO `sys_role_menu` VALUES (1, 1060);
INSERT INTO `sys_role_menu` VALUES (1, 1061);
INSERT INTO `sys_role_menu` VALUES (1, 1064);
INSERT INTO `sys_role_menu` VALUES (1, 1065);
INSERT INTO `sys_role_menu` VALUES (1, 1066);
INSERT INTO `sys_role_menu` VALUES (1, 1124);
INSERT INTO `sys_role_menu` VALUES (1, 1145);
INSERT INTO `sys_role_menu` VALUES (1, 1146);
INSERT INTO `sys_role_menu` VALUES (1, 1147);
INSERT INTO `sys_role_menu` VALUES (1, 1148);
INSERT INTO `sys_role_menu` VALUES (1, 1149);
INSERT INTO `sys_role_menu` VALUES (1, 1150);
INSERT INTO `sys_role_menu` VALUES (1, 1151);
INSERT INTO `sys_role_menu` VALUES (1, 1152);
INSERT INTO `sys_role_menu` VALUES (1, 1153);
INSERT INTO `sys_role_menu` VALUES (1, 1154);
INSERT INTO `sys_role_menu` VALUES (1, 1155);
INSERT INTO `sys_role_menu` VALUES (1, 1156);
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
INSERT INTO `sys_role_menu` VALUES (3, 1);
INSERT INTO `sys_role_menu` VALUES (3, 4);
INSERT INTO `sys_role_menu` VALUES (3, 100);
INSERT INTO `sys_role_menu` VALUES (3, 101);
INSERT INTO `sys_role_menu` VALUES (3, 102);
INSERT INTO `sys_role_menu` VALUES (3, 103);
INSERT INTO `sys_role_menu` VALUES (3, 104);
INSERT INTO `sys_role_menu` VALUES (3, 105);
INSERT INTO `sys_role_menu` VALUES (3, 106);
INSERT INTO `sys_role_menu` VALUES (3, 1000);
INSERT INTO `sys_role_menu` VALUES (3, 1001);
INSERT INTO `sys_role_menu` VALUES (3, 1002);
INSERT INTO `sys_role_menu` VALUES (3, 1003);
INSERT INTO `sys_role_menu` VALUES (3, 1004);
INSERT INTO `sys_role_menu` VALUES (3, 1005);
INSERT INTO `sys_role_menu` VALUES (3, 1006);
INSERT INTO `sys_role_menu` VALUES (3, 1007);
INSERT INTO `sys_role_menu` VALUES (3, 1008);
INSERT INTO `sys_role_menu` VALUES (3, 1009);
INSERT INTO `sys_role_menu` VALUES (3, 1010);
INSERT INTO `sys_role_menu` VALUES (3, 1011);
INSERT INTO `sys_role_menu` VALUES (3, 1012);
INSERT INTO `sys_role_menu` VALUES (3, 1013);
INSERT INTO `sys_role_menu` VALUES (3, 1014);
INSERT INTO `sys_role_menu` VALUES (3, 1015);
INSERT INTO `sys_role_menu` VALUES (3, 1016);
INSERT INTO `sys_role_menu` VALUES (3, 1017);
INSERT INTO `sys_role_menu` VALUES (3, 1018);
INSERT INTO `sys_role_menu` VALUES (3, 1019);
INSERT INTO `sys_role_menu` VALUES (3, 1020);
INSERT INTO `sys_role_menu` VALUES (3, 1021);
INSERT INTO `sys_role_menu` VALUES (3, 1022);
INSERT INTO `sys_role_menu` VALUES (3, 1023);
INSERT INTO `sys_role_menu` VALUES (3, 1024);
INSERT INTO `sys_role_menu` VALUES (3, 1025);
INSERT INTO `sys_role_menu` VALUES (3, 1026);
INSERT INTO `sys_role_menu` VALUES (3, 1027);
INSERT INTO `sys_role_menu` VALUES (3, 1028);
INSERT INTO `sys_role_menu` VALUES (3, 1029);
INSERT INTO `sys_role_menu` VALUES (3, 1030);
INSERT INTO `sys_role_menu` VALUES (3, 1031);
INSERT INTO `sys_role_menu` VALUES (3, 1032);
INSERT INTO `sys_role_menu` VALUES (3, 1033);
INSERT INTO `sys_role_menu` VALUES (3, 1034);
INSERT INTO `sys_role_menu` VALUES (3, 1062);
INSERT INTO `sys_role_menu` VALUES (3, 1064);
INSERT INTO `sys_role_menu` VALUES (4, 5);
INSERT INTO `sys_role_menu` VALUES (4, 6);
INSERT INTO `sys_role_menu` VALUES (4, 7);
INSERT INTO `sys_role_menu` VALUES (4, 8);
INSERT INTO `sys_role_menu` VALUES (4, 9);
INSERT INTO `sys_role_menu` VALUES (4, 10);
INSERT INTO `sys_role_menu` VALUES (4, 11);
INSERT INTO `sys_role_menu` VALUES (4, 12);
INSERT INTO `sys_role_menu` VALUES (4, 13);
INSERT INTO `sys_role_menu` VALUES (4, 14);
INSERT INTO `sys_role_menu` VALUES (4, 15);
INSERT INTO `sys_role_menu` VALUES (4, 16);
INSERT INTO `sys_role_menu` VALUES (4, 17);
INSERT INTO `sys_role_menu` VALUES (4, 18);
INSERT INTO `sys_role_menu` VALUES (4, 19);
INSERT INTO `sys_role_menu` VALUES (4, 20);
INSERT INTO `sys_role_menu` VALUES (4, 21);
INSERT INTO `sys_role_menu` VALUES (4, 22);
INSERT INTO `sys_role_menu` VALUES (4, 23);
INSERT INTO `sys_role_menu` VALUES (4, 24);
INSERT INTO `sys_role_menu` VALUES (4, 25);
INSERT INTO `sys_role_menu` VALUES (4, 27);
INSERT INTO `sys_role_menu` VALUES (4, 28);
INSERT INTO `sys_role_menu` VALUES (4, 29);
INSERT INTO `sys_role_menu` VALUES (4, 31);
INSERT INTO `sys_role_menu` VALUES (4, 32);
INSERT INTO `sys_role_menu` VALUES (4, 33);
INSERT INTO `sys_role_menu` VALUES (4, 35);
INSERT INTO `sys_role_menu` VALUES (4, 36);
INSERT INTO `sys_role_menu` VALUES (4, 37);
INSERT INTO `sys_role_menu` VALUES (4, 38);
INSERT INTO `sys_role_menu` VALUES (4, 39);
INSERT INTO `sys_role_menu` VALUES (4, 40);
INSERT INTO `sys_role_menu` VALUES (4, 41);
INSERT INTO `sys_role_menu` VALUES (4, 42);
INSERT INTO `sys_role_menu` VALUES (4, 43);
INSERT INTO `sys_role_menu` VALUES (4, 44);
INSERT INTO `sys_role_menu` VALUES (4, 45);
INSERT INTO `sys_role_menu` VALUES (4, 46);
INSERT INTO `sys_role_menu` VALUES (4, 47);
INSERT INTO `sys_role_menu` VALUES (4, 48);
INSERT INTO `sys_role_menu` VALUES (4, 49);
INSERT INTO `sys_role_menu` VALUES (4, 50);
INSERT INTO `sys_role_menu` VALUES (4, 51);
INSERT INTO `sys_role_menu` VALUES (4, 52);
INSERT INTO `sys_role_menu` VALUES (4, 53);
INSERT INTO `sys_role_menu` VALUES (4, 54);
INSERT INTO `sys_role_menu` VALUES (4, 55);
INSERT INTO `sys_role_menu` VALUES (4, 56);
INSERT INTO `sys_role_menu` VALUES (4, 1065);
INSERT INTO `sys_role_menu` VALUES (4, 1143);
INSERT INTO `sys_role_menu` VALUES (4, 1144);

-- ----------------------------
-- Table structure for sys_tenant
-- ----------------------------
DROP TABLE IF EXISTS `sys_tenant`;
CREATE TABLE `sys_tenant`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '商户名称',
  `address` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系地址',
  `manager` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '负责人',
  `phone` varchar(18) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  `start_time` datetime(0) NULL DEFAULT NULL COMMENT '起租时间',
  `end_time` datetime(0) NULL DEFAULT NULL COMMENT '结束时间',
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '安全邮箱',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_tenant
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `dept_id` bigint(20) NULL DEFAULT NULL COMMENT '部门ID',
  `login_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '登录账号',
  `user_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户昵称',
  `user_type` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户类型（00系统用户）',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户邮箱',
  `phonenumber` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机号码',
  `sex` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户性别（0男 1女 2未知）',
  `avatar` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像路径',
  `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码',
  `salt` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '盐加密',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '帐号状态（0正常 1停用）',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '删除标记',
  `login_ip` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '最后登陆IP',
  `login_date` datetime(0) NULL DEFAULT NULL COMMENT '最后登陆时间',
  `create_by` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建日期',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新日期',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 110, 'admin', '超级管理员', '00', 'dd122221111d@163.com', '15888888881', '0', '/upload/admin1694703638771106900.png', '45afb3995fb72f114e98469a3652194a', 'HF44Ku', '0', '0', '127.0.0.1', '2020-01-13 13:20:40', 'admin', '2024-12-21 06:27:15', 'admin', '2021-06-20 13:52:46', '管理员111111', 0);
INSERT INTO `sys_user` VALUES (2, 110, 'test', '测试帐号', '', 'uhvs@163.com', '18518276786', '0', '', '49efab53d90058dd6086aed4158a5bbe', 'YYYYYd', '0', '0', '', NULL, 'admin', '2023-09-15 11:01:46', 'admin', '2023-09-19 22:01:29', '', 0);

-- ----------------------------
-- Table structure for sys_user_online
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_online`;
CREATE TABLE `sys_user_online`  (
  `session_id` varchar(250) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户会话id',
  `login_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '登录名',
  `dept_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '登录名',
  `ipaddr` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `login_location` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `browser` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `os` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `start_timestamp` datetime(0) NULL DEFAULT NULL COMMENT '开始时间',
  `last_access_time` datetime(0) NULL DEFAULT NULL COMMENT '上次访问',
  `expire_time` bigint(20) NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`session_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '在线用户记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_online
-- ----------------------------
INSERT INTO `sys_user_online` VALUES ('6de1f46577c2c02bb92a3cbeecf8718a', '超级管理员', '', '[::1]:52146', '北京市', 'Chrome', 'Windows 10', 'on_line', '2025-07-17 05:42:13', '2025-07-17 05:42:13', 1440, '2025-07-17 05:42:13');

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post`  (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `post_id` bigint(20) NOT NULL COMMENT '岗位ID',
  `del_flag` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
  PRIMARY KEY (`user_id`, `post_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户与岗位关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户和角色关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES (2, 3);

-- ----------------------------
-- Table structure for system_metrics
-- ----------------------------
DROP TABLE IF EXISTS `system_metrics`;
CREATE TABLE `system_metrics`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `data` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `timestamp` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of system_metrics
-- ----------------------------

-- ----------------------------
-- Table structure for test1
-- ----------------------------
DROP TABLE IF EXISTS `test1`;
CREATE TABLE `test1`  (
  `table_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
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
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`table_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '代码生成业务表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of test1
-- ----------------------------

-- ----------------------------
-- Table structure for test2
-- ----------------------------
DROP TABLE IF EXISTS `test2`;
CREATE TABLE `test2`  (
  `column_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_id` bigint(20) NULL DEFAULT NULL COMMENT '归属表编号',
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
  `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`column_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1231 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '代码生成业务表字段' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of test2
-- ----------------------------
INSERT INTO `test2` VALUES (539, 12, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (540, 12, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (541, 12, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (542, 12, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (543, 12, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (544, 12, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (545, 12, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (546, 12, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (547, 12, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (548, 12, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (549, 12, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (550, 12, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (551, 12, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (552, 12, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (553, 12, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (554, 12, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (555, 12, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (556, 12, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (557, 12, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (558, 12, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (559, 12, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (560, 12, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (561, 12, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (562, 12, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (563, 12, 'create_time', '创建时间', 'date', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (564, 12, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (565, 12, 'update_time', '更新时间', 'date', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (566, 12, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (567, 12, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (568, 13, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (569, 13, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (570, 13, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (571, 13, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (572, 13, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (573, 13, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (574, 13, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (575, 13, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (576, 13, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (577, 13, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (578, 13, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (579, 13, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (580, 13, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'select', 'sys_user_sex', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (581, 13, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (582, 13, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (583, 13, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (584, 13, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (585, 13, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (586, 13, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (587, 13, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (588, 13, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (589, 13, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '0', '0', 'LIKE', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (590, 13, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (591, 13, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (592, 13, 'create_time', '创建时间', 'date', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (593, 13, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (594, 13, 'update_time', '更新时间', 'date', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (595, 13, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (596, 13, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '0', '1', 'EQ', 'select', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (597, 14, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (598, 14, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (599, 14, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (600, 14, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (601, 14, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (602, 14, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (603, 14, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (604, 14, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (605, 14, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (606, 14, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (607, 14, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (608, 14, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (609, 14, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (610, 14, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (611, 14, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (612, 14, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (613, 14, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (614, 14, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (615, 14, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (616, 14, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (617, 14, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (618, 14, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (619, 14, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (620, 14, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (621, 14, 'create_time', '创建时间', 'date', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (622, 14, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (623, 14, 'update_time', '更新时间', 'date', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (624, 14, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (625, 14, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (626, 15, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (627, 15, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (628, 15, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (629, 15, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (630, 15, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (631, 15, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (632, 15, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (633, 15, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (634, 15, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (635, 15, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (636, 15, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (637, 15, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (638, 15, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (639, 15, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (640, 15, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (641, 15, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (642, 15, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (643, 15, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (644, 15, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (645, 15, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (646, 15, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (647, 15, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (648, 15, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (649, 15, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (650, 15, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (651, 15, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (652, 15, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (653, 15, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (654, 15, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (655, 16, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (656, 16, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (657, 16, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (658, 16, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (659, 16, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (660, 16, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (661, 16, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (662, 16, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (663, 16, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (664, 16, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (665, 16, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (666, 16, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (667, 16, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (668, 16, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (669, 16, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (670, 16, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (671, 16, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (672, 16, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (673, 16, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (674, 16, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (675, 16, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (676, 16, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (677, 16, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (678, 16, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (679, 16, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (680, 16, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (681, 16, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (682, 16, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (683, 16, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (684, 17, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (685, 17, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (686, 17, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (687, 17, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (688, 17, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (689, 17, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (690, 17, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (691, 17, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (692, 17, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (693, 17, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (694, 17, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (695, 17, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (696, 17, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (697, 17, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (698, 17, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (699, 17, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (700, 17, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (701, 17, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (702, 17, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (703, 17, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (704, 17, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (705, 17, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (706, 17, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (707, 17, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (708, 17, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (709, 17, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (710, 17, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (711, 17, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (712, 17, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (713, 18, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (714, 18, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (715, 18, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (716, 18, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (717, 18, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (718, 18, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (719, 18, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (720, 18, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (721, 18, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (722, 18, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (723, 18, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (724, 18, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (725, 18, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (726, 18, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (727, 18, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (728, 18, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (729, 18, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (730, 18, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (731, 18, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (732, 18, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (733, 18, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (734, 18, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '0', 'EQ', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (735, 18, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (736, 18, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (737, 18, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (738, 18, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (739, 18, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (740, 18, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (741, 18, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (742, 19, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (743, 19, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (744, 19, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (745, 19, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (746, 19, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (747, 19, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (748, 19, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (749, 19, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (750, 19, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (751, 19, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (752, 19, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (753, 19, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (754, 20, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (755, 19, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (756, 20, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (757, 19, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (758, 20, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (759, 19, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (760, 20, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (761, 19, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (762, 19, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (763, 20, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (764, 19, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (765, 20, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (766, 19, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (767, 20, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (768, 19, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (769, 20, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (770, 19, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (771, 20, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (772, 19, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '0', 'EQ', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (773, 20, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (774, 19, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (775, 19, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (776, 20, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (777, 19, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (778, 20, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (779, 19, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (780, 20, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (781, 19, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (782, 20, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (783, 19, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (784, 20, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (785, 19, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (786, 20, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (787, 20, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (788, 20, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (789, 20, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (790, 20, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (791, 20, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (792, 20, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '0', 'EQ', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (793, 20, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (794, 20, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (795, 20, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (796, 20, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (797, 20, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (798, 20, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (799, 20, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (801, 22, 'id', '', 'int(11)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (802, 22, 'username', '', 'varchar(255)', 'string', 'Username', 'username', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (803, 22, 'password', '', 'varchar(255)', 'string', 'Password', 'password', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (804, 22, 'taskContent', '', 'varchar(255)', 'string', 'TaskContent', 'taskContent', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (805, 22, 'startDate', '', 'date', 'Time', 'StartDate', 'startDate', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (806, 22, 'endDate', '', 'date', 'Time', 'EndDate', 'endDate', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'datetime', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (807, 22, 'workDays', '', 'int(11)', 'int64', 'WorkDays', 'workDays', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (808, 22, 'autoSubmit', '', 'tinyint(1)', 'int', 'AutoSubmit', 'autoSubmit', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (809, 22, 'status', '', 'varchar(255)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'radio', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (810, 22, 'sort', '', 'int(11)', 'int64', 'Sort', 'sort', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (811, 22, 'updateTime', '', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'datetime', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (812, 23, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (813, 23, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (814, 23, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (815, 23, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (816, 23, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (817, 23, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (818, 23, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (819, 23, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (820, 23, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (821, 23, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (822, 23, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (823, 23, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (824, 23, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (825, 23, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (826, 23, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (827, 23, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (828, 23, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (829, 23, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (830, 23, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (831, 23, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (832, 23, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (833, 23, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (834, 23, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'LIKE', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (835, 23, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (836, 23, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (837, 23, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (838, 23, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (839, 23, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (840, 23, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 29, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (841, 24, 'id', '', 'bigint(32)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (842, 24, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (843, 24, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (844, 24, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (845, 24, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (846, 24, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (847, 24, 'idcard_path', '身份证照片', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (848, 24, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (849, 24, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (850, 24, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (851, 24, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (852, 24, 'family_id', '家庭ID', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (853, 24, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'select', 'sys_user_sex', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (854, 24, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (855, 24, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (856, 24, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (857, 24, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (858, 24, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (859, 24, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (860, 24, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (861, 24, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (862, 24, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (863, 24, 'del_flag', '删除标识1删除0未删除', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'LIKE', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (864, 24, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '1', '0', '0', '0', 'LIKE', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (865, 24, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (866, 24, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '1', '1', '0', '0', 'LIKE', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (867, 24, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '1', '1', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (868, 24, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (869, 25, 'id', 'ID', 'int(11)', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (870, 25, 'username', '工号', 'varchar(16)', 'string', 'Username', 'username', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (871, 25, 'password', '密码', 'varchar(32)', 'string', 'Password', 'password', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (872, 25, 'prj_code', '项  目  号', 'varchar(16)', 'string', 'PrjCode', 'prjCode', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (873, 25, 'task_content', '任务内容', 'varchar(64)', 'string', 'TaskContent', 'taskContent', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (874, 25, 'start_date', '开始日期', 'date', 'Time', 'StartDate', 'startDate', '0', '0', '0', '1', '1', '1', '1', 'GTE', 'datetime', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (875, 25, 'end_date', '结束日期', 'date', 'Time', 'EndDate', 'endDate', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'datetime', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (876, 25, 'work_days', '本月工时', 'int(11)', 'int64', 'WorkDays', 'workDays', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (877, 25, 'auto_submit', '自动提交', 'char(1)', 'string', 'AutoSubmit', 'autoSubmit', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (878, 25, 'status', '任务状态，ready running end', 'varchar(16)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'radio', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (879, 25, 'sort', '排序，大的优先', 'int(11)', 'int64', 'Sort', 'sort', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (880, 25, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (881, 25, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (882, 26, 'post_id', '岗位ID', 'bigint(20)', 'int64', 'PostId', 'postId', '1', '1', '0', '1', '0', '1', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (883, 26, 'post_code', '岗位编码', 'varchar(64)', 'string', 'PostCode', 'postCode', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (884, 26, 'post_name', '岗位名称', 'varchar(50)', 'string', 'PostName', 'postName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (885, 26, 'post_sort', '显示顺序', 'int(11)', 'int64', 'PostSort', 'postSort', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (886, 26, 'status', '状态（0正常 1停用）', 'char(1)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'radio', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (887, 26, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (888, 26, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (889, 26, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (890, 26, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (891, 26, 'remark', '备注', 'varchar(500)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (892, 26, 'tenant_id', '租户id', 'bigint(20)', 'int64', 'TenantId', 'tenantId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (893, 27, 'dict_id', '字典主键', 'bigint(20)', 'int64', 'DictId', 'dictId', '1', '1', '0', '1', '0', '1', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (894, 27, 'dict_name', '字典名称', 'varchar(100)', 'string', 'DictName', 'dictName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (895, 27, 'dict_type', '字典类型', 'varchar(100)', 'string', 'DictType', 'dictType', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'select', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (896, 27, 'status', '状态（0正常 1停用）', 'char(1)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'radio', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (897, 27, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (898, 27, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (899, 27, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (900, 27, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (901, 27, 'remark', '备注', 'varchar(500)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (902, 28, 'notice_id', '公告ID', 'int(11)', 'int', 'NoticeId', 'noticeId', '1', '1', '0', '0', '0', '1', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (903, 28, 'notice_title', '公告标题', 'varchar(50)', 'string', 'NoticeTitle', 'noticeTitle', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (904, 28, 'notice_type', '公告类型（1通知 2公告）', 'char(1)', 'string', 'NoticeType', 'noticeType', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'select', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (905, 28, 'notice_content', '公告内容', 'varchar(2000)', 'string', 'NoticeContent', 'noticeContent', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (906, 28, 'status', '公告状态（0正常 1关闭）', 'char(1)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'radio', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (907, 28, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (908, 28, 'create_time', '创建时间', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (909, 28, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (910, 28, 'update_time', '更新时间', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (911, 28, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (912, 29, 'created', '创建时间', 'bigint(20)', 'int64', 'Created', 'created', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (913, 29, 'modified', '更新时间', 'bigint(20)', 'int64', 'Modified', 'modified', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (914, 29, 'id', '主键', 'varchar(255)', 'string', 'Id', 'id', '1', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (915, 29, 'cloud_device_id', '云设备ID', 'varchar(255)', 'string', 'CloudDeviceId', 'cloudDeviceId', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (916, 29, 'cloud_product_id', '云产品ID', 'varchar(255)', 'string', 'CloudProductId', 'cloudProductId', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (917, 29, 'cloud_instance_id', '云实例ID', 'varchar(255)', 'string', 'CloudInstanceId', 'cloudInstanceId', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (918, 29, 'drive_instance_id', '驱动实例ID', 'varchar(255)', 'string', 'DriveInstanceId', 'driveInstanceId', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (919, 29, 'name', '名字', 'varchar(255)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (920, 29, 'status', '设备状态', 'varchar(50)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'radio', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (921, 29, 'description', '描述', 'text', 'string', 'Description', 'description', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'textarea', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (922, 29, 'product_id', '产品ID', 'varchar(255)', 'string', 'ProductId', 'productId', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (923, 29, 'secret', '密钥', 'varchar(255)', 'string', 'Secret', 'secret', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (924, 29, 'platform', '平台名称', 'varchar(50)', 'string', 'Platform', 'platform', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (925, 29, 'install_location', '安装地址', 'varchar(255)', 'string', 'InstallLocation', 'installLocation', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (926, 29, 'last_sync_time', '最后一次同步时间', 'bigint(20)', 'int64', 'LastSyncTime', 'lastSyncTime', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (927, 29, 'last_online_time', '最后一次在线时间', 'bigint(20)', 'int64', 'LastOnlineTime', 'lastOnlineTime', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (928, 30, 'id', '主键', 'varchar(255)', 'string', 'Id', 'id', '1', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (929, 30, 'cloud_device_id', '云设备ID', 'varchar(255)', 'string', 'CloudDeviceId', 'cloudDeviceId', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (930, 30, 'cloud_product_id', '云产品ID', 'varchar(255)', 'string', 'CloudProductId', 'cloudProductId', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (931, 30, 'cloud_instance_id', '云实例ID', 'varchar(255)', 'string', 'CloudInstanceId', 'cloudInstanceId', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (932, 30, 'drive_instance_id', '驱动实例ID', 'varchar(255)', 'string', 'DriveInstanceId', 'driveInstanceId', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (933, 30, 'name', '名字', 'varchar(255)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (934, 30, 'status', '设备状态', 'varchar(50)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'radio', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (935, 30, 'description', '描述', 'text', 'string', 'Description', 'description', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'textarea', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (936, 30, 'product_id', '产品ID', 'varchar(255)', 'string', 'ProductId', 'productId', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (937, 30, 'secret', '密钥', 'varchar(255)', 'string', 'Secret', 'secret', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (938, 30, 'platform', '平台名称', 'varchar(50)', 'string', 'Platform', 'platform', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (939, 30, 'install_location', '安装地址', 'varchar(255)', 'string', 'InstallLocation', 'installLocation', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (940, 30, 'last_sync_time', '最后一次同步时间', 'bigint(20)', 'int64', 'LastSyncTime', 'lastSyncTime', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (941, 30, 'last_online_time', '最后一次在线时间', 'bigint(20)', 'int64', 'LastOnlineTime', 'lastOnlineTime', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (942, 30, 'del_flag', '删除标记', 'tinyint(1)', 'int', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (943, 30, 'create_time', '创建日期', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (944, 30, 'update_time', '更新日期', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (945, 30, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (946, 30, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (947, 31, 'id', '', 'bigint(20)', 'int64', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (948, 31, 'product_id', '产品ID', 'bigint(20)', 'int64', 'ProductId', 'productId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (949, 31, 'cloud_device_id', '云设备ID', 'varchar(32)', 'string', 'CloudDeviceId', 'cloudDeviceId', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (950, 31, 'cloud_product_id', '云产品ID', 'varchar(32)', 'string', 'CloudProductId', 'cloudProductId', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (951, 31, 'cloud_instance_id', '云实例ID', 'varchar(32)', 'string', 'CloudInstanceId', 'cloudInstanceId', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (952, 31, 'drive_instance_id', '驱动实例ID', 'varchar(32)', 'string', 'DriveInstanceId', 'driveInstanceId', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (953, 31, 'name', '名字', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (954, 31, 'status', '设备状态', 'varchar(50)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'radio', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (955, 31, 'description', '描述', 'text', 'string', 'Description', 'description', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'textarea', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (956, 31, 'secret', '密钥', 'varchar(32)', 'string', 'Secret', 'secret', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (957, 31, 'platform', '平台名称', 'varchar(50)', 'string', 'Platform', 'platform', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (958, 31, 'install_location', '安装地址', 'varchar(32)', 'string', 'InstallLocation', 'installLocation', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (959, 31, 'last_sync_time', '最后一次同步时间', 'datetime', 'time.Time', 'LastSyncTime', 'lastSyncTime', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'datetime', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (960, 31, 'last_online_time', '最后一次在线时间', 'datetime', 'time.Time', 'LastOnlineTime', 'lastOnlineTime', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'datetime', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (961, 31, 'del_flag', '删除标记', 'tinyint(1)', 'int', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (962, 31, 'create_time', '创建日期', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (963, 31, 'update_time', '更新日期', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (964, 31, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (965, 31, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (966, 32, 'id', '主键', 'bigint(20)', 'int64', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (967, 32, 'name', '名字', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (968, 32, 'key', '产品标识', 'varchar(32)', 'string', 'Key', 'key', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (969, 32, 'cloud_product_id', '云产品ID', 'varchar(32)', 'string', 'CloudProductId', 'cloudProductId', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (970, 32, 'cloud_instance_id', '云实例ID', 'varchar(32)', 'string', 'CloudInstanceId', 'cloudInstanceId', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (971, 32, 'platform', '平台', 'varchar(32)', 'string', 'Platform', 'platform', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (972, 32, 'protocol', '协议', 'varchar(32)', 'string', 'Protocol', 'protocol', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (973, 32, 'node_type', '节点类型', 'varchar(32)', 'string', 'NodeType', 'nodeType', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'select', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (974, 32, 'net_type', '网络类型', 'varchar(32)', 'string', 'NetType', 'netType', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'select', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (975, 32, 'data_format', '数据类型', 'varchar(32)', 'string', 'DataFormat', 'dataFormat', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (976, 32, 'last_sync_time', '最后一次同步时间', 'bigint(20)', 'int64', 'LastSyncTime', 'lastSyncTime', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (977, 32, 'factory', '工厂名称', 'varchar(32)', 'string', 'Factory', 'factory', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (978, 32, 'description', '描述', 'text', 'string', 'Description', 'description', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'textarea', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (979, 32, 'status', '产品状态', 'varchar(32)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'radio', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (980, 32, 'extra', '扩展字段', 'varchar(32)', 'string', 'Extra', 'extra', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (981, 32, 'del_flag', '删除标记', 'tinyint(1)', 'int', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (982, 32, 'create_time', '创建日期', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (983, 32, 'update_time', '更新日期', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (984, 32, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (985, 32, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (986, 33, 'service_name', '', 'longtext', 'string', 'ServiceName', 'serviceName', '0', '0', '1', '1', '1', '1', '0', 'EQ', 'textarea', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (987, 33, 'type', '', 'bigint(20)', 'int64', 'Type', 'type', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (988, 33, 'level', '', 'bigint(20)', 'int64', 'Level', 'level', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (989, 33, 'time', '', 'bigint(20)', 'int64', 'Time', 'time', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (990, 33, 'content', '', 'longtext', 'string', 'Content', 'content', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'textarea', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (991, 33, 'del_flag', '删除标记', 'tinyint(1)', 'int', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (992, 33, 'create_time', '创建日期', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (993, 33, 'update_time', '更新日期', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (994, 33, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (995, 33, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (996, 34, 'id', '', 'varchar(191)', 'string', 'Id', 'id', '1', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (997, 34, 'name', '', 'longtext', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'textarea', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (998, 34, 'device_id', '', 'varchar(191)', 'string', 'DeviceId', 'deviceId', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (999, 34, 'alert_type', '', 'longtext', 'string', 'AlertType', 'alertType', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1000, 34, 'alert_level', '', 'longtext', 'string', 'AlertLevel', 'alertLevel', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'textarea', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1001, 34, 'status', '', 'longtext', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '0', 'EQ', 'radio', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1002, 34, 'condition', '', 'longtext', 'string', 'Condition', 'condition', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'textarea', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1003, 34, 'sub_rule', '', 'longtext', 'string', 'SubRule', 'subRule', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'textarea', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1004, 34, 'notify', '', 'longtext', 'string', 'Notify', 'notify', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'textarea', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1005, 34, 'silence_time', '', 'bigint(20)', 'int64', 'SilenceTime', 'silenceTime', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1006, 34, 'description', '', 'longtext', 'string', 'Description', 'description', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'textarea', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1007, 34, 'del_flag', '删除标记', 'tinyint(1)', 'int', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1008, 34, 'create_time', '创建日期', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1009, 34, 'update_time', '更新日期', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1010, 34, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1011, 34, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1012, 35, 'id', '主键', 'int(11)', 'int', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1013, 35, 'product_id', '产品ID', 'int(11)', 'int', 'ProductId', 'productId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1014, 35, 'code', '标识符', 'varchar(32)', 'string', 'Code', 'code', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1015, 35, 'name', '名字', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1016, 35, 'description', '描述', 'varchar(64)', 'string', 'Description', 'description', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1017, 35, 'require', '是否必须', 'tinyint(1)', 'int', 'Require', 'require', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1018, 35, 'call_type', '调用方式', 'varchar(50)', 'string', 'CallType', 'callType', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'select', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1019, 35, 'input_params', '输入参数', 'text', 'string', 'InputParams', 'inputParams', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'textarea', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1020, 35, 'output_params', '输入参数', 'text', 'string', 'OutputParams', 'outputParams', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'textarea', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1021, 35, 'tag', '标签', 'varchar(50)', 'string', 'Tag', 'tag', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1022, 35, 'system', '系统内置', 'tinyint(1)', 'int', 'System', 'system', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1023, 35, 'del_flag', '删除标记', 'tinyint(1)', 'int', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1024, 35, 'create_time', '创建日期', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1025, 35, 'update_time', '更新日期', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1026, 35, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1027, 35, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1028, 36, 'id', '主键', 'varchar(32)', 'string', 'Id', 'id', '1', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1029, 36, 'product_id', '产品ID', 'varchar(32)', 'string', 'ProductId', 'productId', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1030, 36, 'event_type', '事件类型', 'varchar(32)', 'string', 'EventType', 'eventType', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'select', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1031, 36, 'code', '标识符', 'varchar(32)', 'string', 'Code', 'code', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1032, 36, 'name', '名字', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1033, 36, 'description', '描述', 'text', 'string', 'Description', 'description', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'textarea', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1034, 36, 'require', '是否必须', 'tinyint(1)', 'int', 'Require', 'require', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1035, 36, 'output_params', '输入参数', 'text', 'string', 'OutputParams', 'outputParams', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'textarea', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1036, 36, 'tag', '标签', 'varchar(50)', 'string', 'Tag', 'tag', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1037, 36, 'system', '系统内置', 'tinyint(1)', 'int', 'System', 'system', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1038, 36, 'del_flag', '删除标记', 'tinyint(1)', 'int', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1039, 36, 'create_time', '创建日期', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1040, 36, 'update_time', '更新日期', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1041, 36, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1042, 36, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1043, 37, 'id', '主键', 'int(11)', 'int', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1044, 37, 'product_id', '产品ID', 'int(11)', 'int', 'ProductId', 'productId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1045, 37, 'name', '名字', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1046, 37, 'code', '标识符', 'varchar(32)', 'string', 'Code', 'code', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1047, 37, 'access_mode', '读写模型', 'varchar(50)', 'string', 'AccessMode', 'accessMode', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1048, 37, 'require', '是否必须', 'tinyint(1)', 'int', 'Require', 'require', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1049, 37, 'type_spec', '属性物模型详情', 'longtext', 'string', 'TypeSpec', 'typeSpec', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'textarea', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1050, 37, 'description', '描述', 'longtext', 'string', 'Description', 'description', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'textarea', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1051, 37, 'tag', '标签', 'varchar(50)', 'string', 'Tag', 'tag', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1052, 37, 'system', '系统内置', 'tinyint(1)', 'int', 'System', 'system', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1053, 37, 'del_flag', '删除标记', 'tinyint(1)', 'int', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1054, 37, 'create_time', '创建日期', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1055, 37, 'update_time', '更新日期', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1056, 37, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1057, 37, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1058, 38, 'id', '主键', 'int(11)', 'int', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1059, 38, 'name', '名字', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1060, 38, 'description', '描述', 'longtext', 'string', 'Description', 'description', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'textarea', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1061, 38, 'status', '状态', 'varchar(50)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'radio', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1062, 38, 'filter', '', 'longtext', 'string', 'Filter', 'filter', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'textarea', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1063, 38, 'data_resource_id', '资源ID', 'int(11)', 'int', 'DataResourceId', 'dataResourceId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1064, 38, 'del_flag', '删除标记', 'tinyint(1)', 'int', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1065, 38, 'create_time', '创建日期', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1066, 38, 'update_time', '更新日期', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1067, 38, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1068, 38, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1069, 39, 'id', '主键', 'int(11)', 'int', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1070, 39, 'product_id', '产品ID', 'int(11)', 'int', 'ProductId', 'productId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1071, 39, 'code', '标识符', 'varchar(32)', 'string', 'Code', 'code', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1072, 39, 'tag', '标签', 'varchar(50)', 'string', 'Tag', 'tag', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1073, 39, 'name', '名字', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1074, 39, 'call_type', '调用方式', 'varchar(50)', 'string', 'CallType', 'callType', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'select', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1075, 39, 'input_params', '输入参数', 'text', 'string', 'InputParams', 'inputParams', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'textarea', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1076, 39, 'output_params', '输入参数', 'text', 'string', 'OutputParams', 'outputParams', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'textarea', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1077, 39, 'remark', '描述', 'varchar(64)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1078, 39, 'del_flag', '删除标记', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1079, 39, 'create_time', '创建日期', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1080, 39, 'update_time', '更新日期', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1081, 39, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1082, 39, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1083, 39, 'tenant_id', '租户id', 'bigint(20)', 'int64', 'TenantId', 'tenantId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1084, 40, 'id', '主键', 'varchar(32)', 'string', 'Id', 'id', '1', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1085, 40, 'product_id', '产品ID', 'varchar(32)', 'string', 'ProductId', 'productId', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1086, 40, 'event_type', '事件类型', 'varchar(32)', 'string', 'EventType', 'eventType', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'select', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1087, 40, 'code', '标识符', 'varchar(32)', 'string', 'Code', 'code', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1088, 40, 'tag', '标签', 'varchar(50)', 'string', 'Tag', 'tag', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1089, 40, 'name', '名字', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1090, 40, 'remark', '描述', 'longtext', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'textarea', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1091, 40, 'output_params', '输入参数', 'text', 'string', 'OutputParams', 'outputParams', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'textarea', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1092, 40, 'del_flag', '删除标记', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'LIKE', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1093, 40, 'create_time', '创建日期', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1094, 40, 'update_time', '更新日期', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1095, 40, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1096, 40, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1097, 40, 'tenant_id', '租户id', 'bigint(20)', 'int64', 'TenantId', 'tenantId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1098, 41, 'id', '主键', 'int(11)', 'int', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1099, 41, 'product_id', '产品ID', 'int(11)', 'int', 'ProductId', 'productId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1100, 41, 'name', '名字', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1101, 41, 'code', '标识符', 'varchar(32)', 'string', 'Code', 'code', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1102, 41, 'tag', '标签', 'varchar(50)', 'string', 'Tag', 'tag', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1103, 41, 'access_mode', '读写模型', 'varchar(50)', 'string', 'AccessMode', 'accessMode', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1104, 41, 'type', '数据类型', 'varchar(50)', 'string', 'Type', 'type', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'select', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1105, 41, 'unit', '单位', 'varchar(50)', 'string', 'Unit', 'unit', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1106, 41, 'remark', '描述', 'longtext', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'textarea', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1107, 41, 'del_flag', '删除标记', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1108, 41, 'create_time', '创建日期', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1109, 41, 'update_time', '更新日期', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1110, 41, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1111, 41, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1112, 41, 'tenant_id', '租户id', 'bigint(20)', 'int64', 'TenantId', 'tenantId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1113, 42, 'id', '主键', 'int(11)', 'int', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1114, 42, 'product_id', '产品ID', 'int(11)', 'int', 'ProductId', 'productId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1115, 42, 'code', '标识符', 'varchar(32)', 'string', 'Code', 'code', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1116, 42, 'tag', '标签', 'varchar(50)', 'string', 'Tag', 'tag', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1117, 42, 'name', '名字', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1118, 42, 'call_type', '调用方式', 'varchar(50)', 'string', 'CallType', 'callType', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'select', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1119, 42, 'input_params', '输入参数', 'text', 'string', 'InputParams', 'inputParams', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'textarea', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1120, 42, 'output_params', '输入参数', 'text', 'string', 'OutputParams', 'outputParams', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'textarea', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1121, 42, 'remark', '描述', 'varchar(64)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1122, 42, 'del_flag', '删除标记', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1123, 42, 'create_time', '创建日期', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1124, 42, 'update_time', '更新日期', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1125, 42, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1126, 42, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1127, 42, 'tenant_id', '租户id', 'bigint(20)', 'int64', 'TenantId', 'tenantId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1128, 43, 'id', '主键', 'mediumint(9)', 'int', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1129, 43, 'product_id', '产品ID', 'int(11)', 'int', 'ProductId', 'productId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1130, 43, 'event_type', '事件类型', 'varchar(32)', 'string', 'EventType', 'eventType', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'select', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1131, 43, 'code', '标识符', 'varchar(32)', 'string', 'Code', 'code', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1132, 43, 'tag', '标签', 'varchar(50)', 'string', 'Tag', 'tag', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1133, 43, 'name', '名字', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1134, 43, 'remark', '描述', 'longtext', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'textarea', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1135, 43, 'output_params', '输入参数', 'text', 'string', 'OutputParams', 'outputParams', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'textarea', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1136, 43, 'del_flag', '删除标记', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'LIKE', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1137, 43, 'create_time', '创建日期', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1138, 43, 'update_time', '更新日期', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1139, 43, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1140, 43, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1141, 43, 'tenant_id', '租户id', 'bigint(20)', 'int64', 'TenantId', 'tenantId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1142, 44, 'id', '主键', 'bigint(20)', 'int64', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1143, 44, 'product_id', '产品ID', 'int(11)', 'int', 'ProductId', 'productId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1144, 44, 'code', '标识符', 'varchar(32)', 'string', 'Code', 'code', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1145, 44, 'tag', '标签', 'varchar(50)', 'string', 'Tag', 'tag', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1146, 44, 'name', '名字', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1147, 44, 'call_type', '调用方式', 'varchar(50)', 'string', 'CallType', 'callType', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'select', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1148, 44, 'input_params', '输入参数', 'text', 'string', 'InputParams', 'inputParams', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'textarea', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1149, 44, 'output_params', '输入参数', 'text', 'string', 'OutputParams', 'outputParams', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'textarea', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1150, 44, 'remark', '描述', 'varchar(64)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1151, 44, 'del_flag', '删除标记', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1152, 44, 'create_time', '创建日期', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1153, 44, 'update_time', '更新日期', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1154, 44, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1155, 44, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1156, 44, 'tenant_id', '租户id', 'bigint(20)', 'int64', 'TenantId', 'tenantId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1157, 45, 'id', '主键', 'bigint(20)', 'int64', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1158, 45, 'product_id', '产品ID', 'int(11)', 'int', 'ProductId', 'productId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1159, 45, 'event_type', '事件类型', 'varchar(32)', 'string', 'EventType', 'eventType', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'select', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1160, 45, 'code', '标识符', 'varchar(32)', 'string', 'Code', 'code', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1161, 45, 'tag', '标签', 'varchar(50)', 'string', 'Tag', 'tag', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1162, 45, 'name', '名字', 'varchar(32)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1163, 45, 'remark', '描述', 'longtext', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'textarea', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1164, 45, 'output_params', '输入参数', 'text', 'string', 'OutputParams', 'outputParams', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'textarea', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1165, 45, 'del_flag', '删除标记', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'LIKE', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1166, 45, 'create_time', '创建日期', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1167, 45, 'update_time', '更新日期', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1168, 45, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1169, 45, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1170, 45, 'tenant_id', '租户id', 'bigint(20)', 'int64', 'TenantId', 'tenantId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1171, 46, 'dict_code', '字典编码', 'bigint(20)', 'int64', 'DictCode', 'dictCode', '1', '1', '0', '0', '0', '1', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1172, 46, 'dict_sort', '字典排序', 'int(11)', 'int', 'DictSort', 'dictSort', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1173, 46, 'dict_label', '字典标签', 'varchar(100)', 'string', 'DictLabel', 'dictLabel', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1174, 46, 'dict_value', '字典键值', 'varchar(100)', 'string', 'DictValue', 'dictValue', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1175, 46, 'dict_type', '字典类型', 'varchar(100)', 'string', 'DictType', 'dictType', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'select', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1176, 46, 'css_class', '样式属性（其他样式扩展）', 'varchar(100)', 'string', 'CssClass', 'cssClass', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1177, 46, 'list_class', '表格回显样式', 'varchar(100)', 'string', 'ListClass', 'listClass', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1178, 46, 'is_default', '是否默认（Y是 N否）', 'char(1)', 'string', 'IsDefault', 'isDefault', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1179, 46, 'status', '状态（0正常 1停用）', 'char(1)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'radio', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1180, 46, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1181, 46, 'create_time', '创建时间', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1182, 46, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1183, 46, 'update_time', '更新时间', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1184, 46, 'remark', '备注', 'varchar(500)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1185, 47, 'dict_id', '字典主键', 'bigint(20)', 'int64', 'DictId', 'dictId', '1', '1', '0', '0', '0', '1', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1186, 47, 'dict_name', '字典名称', 'varchar(100)', 'string', 'DictName', 'dictName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1187, 47, 'dict_type', '字典类型', 'varchar(100)', 'string', 'DictType', 'dictType', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'select', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1188, 47, 'status', '状态（0正常 1停用）', 'char(1)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'radio', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1189, 47, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1190, 47, 'create_time', '创建时间', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1191, 47, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'LIKE', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1192, 47, 'update_time', '更新时间', 'datetime', 'time.Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1193, 47, 'remark', '备注', 'varchar(500)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1194, 48, 'info_id', '访问ID', 'bigint(20)', 'int64', 'InfoId', 'infoId', '1', '1', '0', '0', '0', '1', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1195, 48, 'login_name', '登录账号', 'varchar(50)', 'string', 'LoginName', 'loginName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1196, 48, 'ipaddr', '登录IP地址', 'varchar(50)', 'string', 'Ipaddr', 'ipaddr', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1197, 48, 'login_location', '登录地点', 'varchar(255)', 'string', 'LoginLocation', 'loginLocation', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1198, 48, 'browser', '浏览器类型', 'varchar(50)', 'string', 'Browser', 'browser', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1199, 48, 'os', '操作系统', 'varchar(50)', 'string', 'Os', 'os', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1200, 48, 'status', '登录状态（0成功 1失败）', 'char(1)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'radio', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1201, 48, 'msg', '提示消息', 'varchar(255)', 'string', 'Msg', 'msg', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1202, 48, 'login_time', '访问时间', 'datetime', 'time.Time', 'LoginTime', 'loginTime', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'datetime', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1203, 49, 'session_id', '用户会话id', 'varchar(250)', 'string', 'SessionId', 'sessionId', '1', '0', '0', '0', '0', '1', '0', 'LIKE', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1204, 49, 'login_name', '登录名', 'varchar(64)', 'string', 'LoginName', 'loginName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1205, 49, 'dept_name', '登录名', 'varchar(64)', 'string', 'DeptName', 'deptName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1206, 49, 'ipaddr', '', 'varchar(64)', 'string', 'Ipaddr', 'ipaddr', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1207, 49, 'login_location', '', 'varchar(64)', 'string', 'LoginLocation', 'loginLocation', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1208, 49, 'browser', '', 'varchar(64)', 'string', 'Browser', 'browser', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1209, 49, 'os', '', 'varchar(64)', 'string', 'Os', 'os', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1210, 49, 'status', '', 'varchar(64)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'radio', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1211, 49, 'start_timestamp', '开始时间', 'datetime', 'time.Time', 'StartTimestamp', 'startTimestamp', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'datetime', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1212, 49, 'last_access_time', '上次访问', 'datetime', 'time.Time', 'LastAccessTime', 'lastAccessTime', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'datetime', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1213, 49, 'expire_time', '', 'bigint(20)', 'int64', 'ExpireTime', 'expireTime', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1214, 49, 'create_time', '创建时间', 'datetime', 'time.Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1215, 50, 'oper_id', '日志主键', 'bigint(20)', 'int64', 'OperId', 'operId', '1', '1', '0', '0', '0', '1', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1216, 50, 'title', '模块标题', 'varchar(50)', 'string', 'Title', 'title', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1217, 50, 'business_type', '业务类型（0其它 1新增 2修改 3删除）', 'int(11)', 'int', 'BusinessType', 'businessType', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1218, 50, 'method', '方法名称', 'varchar(100)', 'string', 'Method', 'method', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1219, 50, 'request_method', '请求方式', 'varchar(10)', 'string', 'RequestMethod', 'requestMethod', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1220, 50, 'operator_type', '操作类别（0其它 1后台用户 2手机端用户）', 'int(11)', 'int', 'OperatorType', 'operatorType', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'select', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1221, 50, 'oper_name', '操作人员', 'varchar(50)', 'string', 'OperName', 'operName', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1222, 50, 'dept_name', '部门名称', 'varchar(50)', 'string', 'DeptName', 'deptName', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1223, 50, 'oper_url', '请求URL', 'varchar(255)', 'string', 'OperUrl', 'operUrl', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1224, 50, 'oper_ip', '主机地址', 'varchar(50)', 'string', 'OperIp', 'operIp', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1225, 50, 'oper_location', '操作地点', 'varchar(255)', 'string', 'OperLocation', 'operLocation', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1226, 50, 'oper_param', '请求参数', 'varchar(2000)', 'string', 'OperParam', 'operParam', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1227, 50, 'json_result', '返回参数', 'varchar(2000)', 'string', 'JsonResult', 'jsonResult', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1228, 50, 'status', '操作状态（0正常 1异常）', 'int(11)', 'int', 'Status', 'status', '0', '0', '1', '1', '1', '1', '0', 'EQ', 'radio', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1229, 50, 'error_msg', '错误消息', 'varchar(2000)', 'string', 'ErrorMsg', 'errorMsg', '0', '0', '0', '1', '1', '1', '0', 'LIKE', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `test2` VALUES (1230, 50, 'oper_time', '操作时间', 'datetime', 'time.Time', 'OperTime', 'operTime', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'datetime', '', 16, 'admin', NULL, '', NULL);

-- ----------------------------
-- Table structure for thing_model_template
-- ----------------------------
DROP TABLE IF EXISTS `thing_model_template`;
CREATE TABLE `thing_model_template`  (
  `created` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `modified` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `category_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '品类key',
  `category_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '品类名字',
  `thing_model_json` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '物模型信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of thing_model_template
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名字',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `lang` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '语言',
  `gateway_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密钥',
  `open_api_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密钥',
  INDEX `idx_user_username`(`username`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('admin', '$2a$10$s34GbctUs6CV9pDbUiXBjOEXiqwh.uEWeehHkLPDCc5KnuW0TWL0y', 'en', 'd8133e69-143b-40e5-a725-773b36c8f304', '70369c02-b658-43b5-a169-a7ce0f53cbd8');

SET FOREIGN_KEY_CHECKS = 1;
