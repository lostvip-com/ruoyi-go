/*
 Navicat Premium Data Transfer

 Source Server         : robnote.com
 Source Server Type    : MySQL
 Source Server Version : 50738
 Source Host           : robnote.com:13306
 Source Schema         : ruoyi-go

 Target Server Type    : MySQL
 Target Server Version : 50738
 File Encoding         : 65001

 Date: 19/11/2023 00:53:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '代码生成业务表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_table
-- ----------------------------
INSERT INTO `gen_table` VALUES (13, 'his_patient', '患者基本信息', 'HisPatient', 'crud', 'robvi', 'biz', 'patient', '患者基本信息', 'lv', '', 'admin', '2023-10-29 22:05:20', 'admin', '2023-10-30 19:15:28', '');

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
) ENGINE = InnoDB AUTO_INCREMENT = 626 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '代码生成业务表字段' ROW_FORMAT = Dynamic;

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
INSERT INTO `gen_table_column` VALUES (571, 13, 'head_url', '照片', 'varchar(256)', 'string', 'HeadUrl', 'headUrl', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (572, 13, 'idcard_path', '', 'varchar(128)', 'string', 'IdcardPath', 'idcardPath', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (573, 13, 'idcard', '证件号', 'varchar(18)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'LIKE', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (574, 13, 'bed_no', '床号', 'varchar(32)', 'string', 'BedNo', 'bedNo', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (575, 13, 'doctor_id', '责任医生Id', 'bigint(20)', 'int64', 'DoctorId', 'doctorId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (576, 13, 'org_id', '建档单位', 'bigint(20)', 'int64', 'OrgId', 'orgId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (577, 13, 'org_address', '建档单位地址', 'varchar(255)', 'string', 'OrgAddress', 'orgAddress', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (578, 13, 'org_establish', '建档单位', 'varchar(32)', 'string', 'OrgEstablish', 'orgEstablish', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (579, 13, 'family_id', '', 'bigint(32)', 'int64', 'FamilyId', 'familyId', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (580, 13, 'sex', '性别', 'varchar(1)', 'string', 'Sex', 'sex', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'select', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (581, 13, 'birth', '生日', 'varchar(10)', 'string', 'Birth', 'birth', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (582, 13, 'weight', '体重', 'float(10,0)', 'int64', 'Weight', 'weight', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 15, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (583, 13, 'height', '身高', 'float(10,0)', 'int64', 'Height', 'height', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 16, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (584, 13, 'nation', '民族', 'varchar(128)', 'string', 'Nation', 'nation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 17, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (585, 13, 'native_place', '籍贯', 'varchar(128)', 'string', 'NativePlace', 'nativePlace', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 18, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (586, 13, 'address', '现居地址', 'varchar(128)', 'string', 'Address', 'address', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 19, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (587, 13, 'occupation', '职业', 'varchar(32)', 'string', 'Occupation', 'occupation', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 20, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (588, 13, 'contactor_phone', '联系人手机号', 'varchar(11)', 'string', 'ContactorPhone', 'contactorPhone', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 21, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (589, 13, 'contactor_name', '联系人', 'varchar(128)', 'string', 'ContactorName', 'contactorName', '0', '0', '1', '1', '1', '1', '0', 'LIKE', 'input', '', 22, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (590, 13, 'del_flag', '', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 23, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (591, 13, 'create_by', '创建人', 'varchar(32)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 24, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (592, 13, 'create_time', '创建时间', 'date', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 25, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (593, 13, 'update_by', '更新者', 'varchar(32)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 26, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (594, 13, 'update_time', '更新时间', 'date', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 27, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (595, 13, 'remark', '备注信息', 'varchar(32)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 28, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (596, 13, 'dept_id', '建档单位', 'bigint(20)', 'int64', 'DeptId', 'deptId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 29, 'admin', NULL, '', NULL);
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

-- ----------------------------
-- Table structure for his_patient
-- ----------------------------
DROP TABLE IF EXISTS `his_patient`;
CREATE TABLE `his_patient`  (
  `id` bigint(32) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '姓名',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `head_url` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '照片',
  `idcard_path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `idcard` varchar(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '证件号',
  `bed_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '床号',
  `doctor_id` bigint(20) NULL DEFAULT NULL COMMENT '责任医生Id',
  `org_id` bigint(20) NULL DEFAULT NULL COMMENT '建档单位',
  `org_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '建档单位地址',
  `org_establish` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '建档单位',
  `family_id` bigint(32) NULL DEFAULT NULL,
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
  `del_flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0',
  `create_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `create_time` date NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `update_time` date NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  `dept_id` bigint(20) NULL DEFAULT NULL COMMENT '建档单位',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '患者基本信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of his_patient
-- ----------------------------
INSERT INTO `his_patient` VALUES (1, 'test', '17600196256', NULL, NULL, '63280120170516003X', '1', 1, NULL, '的撒范德萨春树暮云大本营', '的撒范德萨', NULL, '0', '2023-01-01', 12, 2, '汉', '的撒法第三', '浦东', '春树暮云', '17600196256', '121', '0', NULL, '2023-08-20', NULL, '2023-08-20', '夺花', NULL);
INSERT INTO `his_patient` VALUES (2, '秀梅', '17600196256', NULL, NULL, '63280120170516003X', '1', 1, NULL, '', '', NULL, NULL, '2017-05-16', 12, 1, '汉', '1212121', '浦东', '12121', '', '', '0', NULL, '2023-08-20', NULL, NULL, 'fdsfds', NULL);
INSERT INTO `his_patient` VALUES (3, '运营成本', '17600196256', NULL, NULL, '63280120170516003X', '1', 1, 103, '', '研发部门', NULL, NULL, '2017-05-16', 21, 21, '汉', '青海', '浦东', '', '', '', '0', 'admin', '2023-08-20', NULL, NULL, '士大夫噶但是', NULL);
INSERT INTO `his_patient` VALUES (4, '花木成畦手自栽大本营', '17600196256', NULL, NULL, '63280120170526003X', '1', 1, 103, '', '研发部门', NULL, NULL, '2017-05-16', 11, 21, '21', '21', '浦东', '121', '', '', '0', 'admin', '2023-08-20', NULL, NULL, '', NULL);

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
INSERT INTO `sys_config` VALUES (1, '主框架页-默认皮肤样式名称', 'sys.index.skinName', 'skin-blue', 'Y', 'admin', '2018-03-16 11:33:00', '', '2021-06-20 13:51:55', '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow');
INSERT INTO `sys_config` VALUES (2, '用户管理-账号初始密码', 'sys.user.initPassword', '123456', 'Y', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '初始化密码 123456');
INSERT INTO `sys_config` VALUES (3, '主框架页-侧边栏主题', 'sys.index.sideTheme', 'theme-dark', 'Y', 'admin', '2018-03-16 11:33:00', '', '2020-02-05 10:46:28', '深黑主题theme-dark，浅色主题theme-light，深蓝主题theme-blue');
INSERT INTO `sys_config` VALUES (4, '静态资源网盘存储', 'sys.resource.url', '/static', 'Y', 'admin', '2020-02-18 20:10:33', '', '2020-03-23 20:51:39', 'public目录下的静态资源存储到OSS/COS等网盘，将键值设为/public表示本地');

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint(20) NULL DEFAULT 0 COMMENT '父部门id',
  `ancestors` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `order_num` int(11) NULL DEFAULT 0 COMMENT '显示顺序',
  `leader` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `tenant_id` bigint(20) NULL DEFAULT 0 COMMENT '租户id',
  PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 115 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '部门表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (100, 0, '0', 'DPC', 0, 'admin', '15888888888', '110@qq.com', '0', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2023-09-15 11:34:01', 0);
INSERT INTO `sys_dept` VALUES (110, 100, '0,100', '江苏铁塔', 100, '', '', '', '0', '0', 'admin', '2019-12-02 17:07:02', 'admin', '2023-09-15 11:54:50', 0);
INSERT INTO `sys_dept` VALUES (111, 100, '0,100', '电影学院', 51, '曾尚兵1', '18788996255', 'ddd@163.com', '0', '2', 'admin', '2020-03-01 09:40:48', 'admin', '2020-03-01 09:40:55', 0);
INSERT INTO `sys_dept` VALUES (112, 100, '0,100', '公司运维', 10, 'lostvip', '18788996255', '331641539@qq.com', '0', '0', 'admin', '2020-03-21 16:30:26', 'admin', '2023-10-29 22:30:15', 0);
INSERT INTO `sys_dept` VALUES (113, 100, '0,100', '江苏铁塔', 2, '江苏铁塔', '', '', '0', '2', 'admin', '2023-09-15 11:38:29', 'admin', '2023-09-15 11:39:15', 0);
INSERT INTO `sys_dept` VALUES (114, 100, '0,100', '爱客', 100, '', '', '', '0', '0', 'admin', '2023-09-15 11:39:05', 'admin', '2023-09-15 11:55:13', 0);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `dict_sort` int(11) NULL DEFAULT 0 COMMENT '字典排序',
  `dict_label` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `css_class` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '表格回显样式',
  `is_default` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'N' COMMENT '是否默认（Y是 N否）',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '字典数据表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (1, 1, '男', '0', 'sys_user_sex', '', 'default', 'Y', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2023-09-21 21:45:51', '[[*{remark}]]');
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
  `dict_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `dict_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_id`) USING BTREE,
  UNIQUE INDEX `dict_type`(`dict_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '字典类型表' ROW_FORMAT = Dynamic;

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
) ENGINE = InnoDB AUTO_INCREMENT = 821 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '系统访问记录' ROW_FORMAT = Dynamic;

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

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `menu_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `menu_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '菜单名称',
  `parent_id` bigint(20) NULL DEFAULT 0 COMMENT '父菜单ID',
  `order_num` int(11) NULL DEFAULT 0 COMMENT '显示顺序',
  `url` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '#' COMMENT '请求地址',
  `target` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '打开方式（menuItem页签 menuBlank新窗口）',
  `menu_type` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '菜单类型（M目录 C菜单 F按钮）',
  `visible` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '菜单状态（0显示 1隐藏）',
  `perms` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '权限标识',
  `icon` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '#' COMMENT '菜单图标',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1118 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '菜单权限表' ROW_FORMAT = Dynamic;

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
INSERT INTO `sys_menu` VALUES (11, '图标演示', 5, 6, '#', '', 'M', '0', '', '', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '图标演示');
INSERT INTO `sys_menu` VALUES (12, '栅格演示', 6, 2, 'demo/form/grid', '', 'C', '0', 'demo:grid:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (13, '下拉框', 6, 3, 'demo/form/select', '', 'C', '0', 'demo:select:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (14, '时间轴', 6, 4, 'demo/form/timeline', '', 'C', '0', 'demo:timeline:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (15, '基本表单', 6, 5, 'demo/form/basic', '', 'C', '0', 'demo:basic:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (16, '卡片列表', 6, 6, 'demo/form/cards', '', 'C', '0', 'demo:cards:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (17, '功能扩展', 6, 7, 'demo/form/jasny', '', 'C', '0', 'demo:jasny:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (18, '拖动排序', 6, 8, 'demo/form/sortable', '', 'C', '0', 'demo:sortable:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (19, '选项卡&面板', 6, 9, 'demo/form/tabs_panels', '', 'C', '0', 'demo:tabs_panels:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (20, '表单校验', 6, 10, 'demo/form/validate', '', 'C', '0', 'demo:validate:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (21, '表单向导', 6, 11, 'demo/form/wizard', '', 'C', '0', 'demo:wizard:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (22, '文件上传', 6, 12, 'demo/form/upload', '', 'C', '0', 'demo:upload:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (23, '日期和时间', 6, 13, 'demo/form/datetime', '', 'C', '0', 'demo:datetime:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (24, '富文本编辑器', 6, 14, 'demo/form/summernote', '', 'C', '0', 'demo:summernote:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (25, '左右互选', 6, 15, 'demo/form/duallistbox', '', 'C', '0', 'demo:duallistbox:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (27, '按钮演示', 6, 1, 'demo/form/button', 'menuItem', 'C', '0', 'demo:button:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', '', '2020-02-04 08:46:48', '');
INSERT INTO `sys_menu` VALUES (28, '数据汇总', 7, 2, 'demo/table/footer', '', 'C', '0', 'demo:footer:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (29, '组合表头', 7, 3, 'demo/table/groupHeader', '', 'C', '0', 'demo:groupHeader:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (31, '记住翻页', 7, 5, 'demo/table/remember', '', 'C', '0', 'demo:remember:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (32, '跳转指定页', 7, 6, 'demo/table/pageGo', '', 'C', '0', 'demo:pageGo:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (33, '查询参数', 7, 7, 'demo/table/params', '', 'C', '0', 'demo:params:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (35, '点击加载表格', 7, 9, 'demo/table/button', '', 'C', '0', 'demo:button:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (36, '表格冻结列', 7, 10, 'demo/table/fixedColumns', '', 'C', '0', 'demo:fixedColumns:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (37, '触发事件', 7, 11, 'demo/table/event', '', 'C', '0', 'demo:event:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (38, '细节视图', 7, 12, 'demo/table/detail', '', 'C', '0', 'demo:detail:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (39, '父子视图', 7, 13, 'demo/table/child', '', 'C', '0', 'demo:child:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (40, '图片预览', 7, 14, 'demo/table/image', '', 'C', '0', 'demo:image:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (41, '动态增删改查', 7, 15, 'demo/table/curd', '', 'C', '0', 'demo:curd:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (42, '表格拖曳', 7, 16, 'demo/table/recorder', '', 'C', '0', 'demo:recorder:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (43, '行内编辑', 7, 17, 'demo/table/editable', '', 'C', '0', 'demo:editable:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (44, '其它操作', 7, 18, 'demo/table/other', '', 'C', '0', 'demo:other:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (45, '查询条件', 7, 1, 'demo/table/search', '', 'C', '0', 'demo:search:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (46, '弹层组件', 8, 2, 'demo/modal/layer', '', 'C', '0', 'demo:layer:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (47, '弹层表格', 8, 3, 'demo/modal/table', '', 'C', '0', 'demo:table:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (48, '模态窗口', 8, 1, 'demo/modal/dialog', '', 'C', '0', 'demo:dialog:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (49, '其他操作', 9, 2, 'demo/operate/other', '', 'C', '0', 'demo:other:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (50, '表格操作', 9, 1, 'demo/operate/table', '', 'C', '0', 'demo:table:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (51, 'Peity', 10, 2, 'demo/report/metrics', '', 'C', '0', 'demo:metrics:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (52, 'SparkLine', 10, 3, 'demo/report/peity', '', 'C', '0', 'demo:peity:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (53, '图表组合', 10, 4, 'demo/report/sparkline', '', 'C', '0', 'demo:sparkline:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (54, '百度Echarts', 10, 1, 'demo/report/echarts', '', 'C', '0', 'demo:echarts:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (55, 'Glyphicons', 11, 2, 'demo/icon/glyphicons', '', 'C', '0', 'demo:glyphicons:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (56, 'Font Awesome', 11, 1, 'demo/icon/fontawesome', '', 'C', '0', 'demo:fontawesome:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
INSERT INTO `sys_menu` VALUES (100, '用户管理', 1, 1, 'system/user', 'menuItem', 'C', '0', 'system:user:view', 'fa fa-address-book', 'admin', '2018-03-16 11:33:00', 'admin', '2023-11-04 16:49:16', '用户管理菜单');
INSERT INTO `sys_menu` VALUES (101, '角色管理', 1, 2, 'system/role', '', 'C', '0', 'system:role:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '角色管理菜单');
INSERT INTO `sys_menu` VALUES (102, '菜单管理', 1, 3, 'system/menu', '', 'C', '0', 'system:menu:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '菜单管理菜单');
INSERT INTO `sys_menu` VALUES (103, '部门管理', 1, 4, 'system/dept', 'menuItem', 'C', '0', 'system:dept:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2023-09-14 16:55:37', '部门管理菜单');
INSERT INTO `sys_menu` VALUES (104, '岗位管理', 1, 5, 'system/post', 'menuItem', 'C', '0', 'system:post:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2023-09-14 16:55:51', '岗位管理菜单');
INSERT INTO `sys_menu` VALUES (105, '字典管理', 1, 6, 'system/dict', '', 'C', '0', 'system:dict:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '字典管理菜单');
INSERT INTO `sys_menu` VALUES (106, '参数设置', 1, 7, 'system/config', '', 'C', '0', 'system:config:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '参数设置菜单');
INSERT INTO `sys_menu` VALUES (109, '在线用户', 2, 1, 'monitor/online', '', 'C', '0', 'monitor:online:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '在线用户菜单');
INSERT INTO `sys_menu` VALUES (110, '定时任务', 2, 2, 'monitor/job', '', 'C', '0', 'monitor:job:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '定时任务菜单');
INSERT INTO `sys_menu` VALUES (112, '服务监控', 2, 3, 'monitor/server', '', 'C', '0', 'monitor:server:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '服务监控菜单');
INSERT INTO `sys_menu` VALUES (113, '表单构建', 3, 1, 'tool/build', '', 'C', '0', 'tool:build:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '表单构建菜单');
INSERT INTO `sys_menu` VALUES (114, '代码生成', 3, 2, 'tool/gen', '', 'C', '0', 'tool:gen:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '代码生成菜单');
INSERT INTO `sys_menu` VALUES (115, '系统接口', 3, 3, 'tool/swagger', '', 'C', '0', 'tool:swagger:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统接口菜单');
INSERT INTO `sys_menu` VALUES (500, '操作日志', 2, 4, 'monitor/operlog', '', 'C', '0', 'monitor:operlog:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '操作日志菜单');
INSERT INTO `sys_menu` VALUES (501, '登录日志', 2, 5, 'monitor/logininfor', '', 'C', '0', 'monitor:logininfor:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '登录日志菜单');
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
INSERT INTO `sys_menu` VALUES (1064, '监控平台', 4, 1, 'https://lostvip.com', 'menuItem', 'C', '0', '', 'fa fa-globe', '', '2021-06-19 23:12:45', '', '2023-10-17 14:10:02', '');
INSERT INTO `sys_menu` VALUES (1065, '搜索自动补全', 6, 100, 'demo/form/autocomplete', 'menuItem', 'C', '', '', 'fa fa-sticky-note-o', '', '2023-09-18 10:15:49', '', '2023-09-18 10:17:03', '');
INSERT INTO `sys_menu` VALUES (1066, '发电量重算', 4, 100, 'iot/hisData/toWizard', 'menuItem', 'C', '', '', 'fa fa-sticky-note-o', '', '2023-09-18 21:42:44', '', '2023-09-19 21:36:52', '');
INSERT INTO `sys_menu` VALUES (1082, '患者基本信息', 4, 1, '/biz/patient', '', 'C', '0', 'biz:patient:view', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '患者基本信息菜单');
INSERT INTO `sys_menu` VALUES (1083, '患者基本信息新增', 1082, 2, '#', '', 'F', '0', 'biz:patient:add', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1084, '患者基本信息查询', 1082, 1, '#', '', 'F', '0', 'biz:patient:list', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1085, '患者基本信息修改', 1082, 3, '#', '', 'F', '0', 'biz:patient:edit', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1086, '患者基本信息删除', 1082, 4, '#', '', 'F', '0', 'biz:patient:remove', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1087, '患者基本信息导出', 1082, 5, '#', '', 'F', '0', 'biz:patient:export', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1106, '患者基本信息', 4, 1, '/biz/patient', '', 'C', '0', 'biz:patient:view', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '患者基本信息菜单');
INSERT INTO `sys_menu` VALUES (1107, '患者基本信息新增', 1106, 2, '#', '', 'F', '0', 'biz:patient:add', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1108, '患者基本信息查询', 1106, 1, '#', '', 'F', '0', 'biz:patient:list', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1109, '患者基本信息修改', 1106, 3, '#', '', 'F', '0', 'biz:patient:edit', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1110, '患者基本信息删除', 1106, 4, '#', '', 'F', '0', 'biz:patient:remove', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1111, '患者基本信息导出', 1106, 5, '#', '', 'F', '0', 'biz:patient:export', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1112, '患者基本信息', 4, 1, '/biz/patient', '', 'C', '0', 'biz:patient:view', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '患者基本信息菜单');
INSERT INTO `sys_menu` VALUES (1113, '患者基本信息新增', 1112, 2, '#', '', 'F', '0', 'biz:patient:add', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1114, '患者基本信息查询', 1112, 1, '#', '', 'F', '0', 'biz:patient:list', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1115, '患者基本信息修改', 1112, 3, '#', '', 'F', '0', 'biz:patient:edit', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1116, '患者基本信息删除', 1112, 4, '#', '', 'F', '0', 'biz:patient:remove', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');
INSERT INTO `sys_menu` VALUES (1117, '患者基本信息导出', 1112, 5, '#', '', 'F', '0', 'biz:patient:export', '#', 'admin', '2020-01-01 00:00:00', 'admin', '2020-01-01 00:00:00', '');

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
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log`  (
  `oper_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '模块标题',
  `business_type` int(11) NULL DEFAULT 0 COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '请求方式',
  `operator_type` int(11) NULL DEFAULT 0 COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '操作人员',
  `dept_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '请求URL',
  `oper_ip` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '操作地点',
  `oper_param` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '请求参数',
  `json_result` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '返回参数',
  `status` int(11) NULL DEFAULT 0 COMMENT '操作状态（0正常 1异常）',
  `error_msg` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '错误消息',
  `oper_time` datetime(0) NULL DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 749 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '操作日志记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------
INSERT INTO `sys_oper_log` VALUES (491, '操作日志管理', 3, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/operlog/clean', '::1', '内网IP', '\"all\"', '{\"code\":0,\"msg\":\"操作成功\",\"data\":2,\"otype\":3}', 0, '', '2020-03-24 13:37:47');
INSERT INTO `sys_oper_log` VALUES (492, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 14:14:53');
INSERT INTO `sys_oper_log` VALUES (493, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 14:15:09');
INSERT INTO `sys_oper_log` VALUES (494, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":12}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [* * * * * *]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 14:33:58');
INSERT INTO `sys_oper_log` VALUES (495, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [* * * * * *]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 14:42:52');
INSERT INTO `sys_oper_log` VALUES (496, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":12}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [0/10 * * * * ?]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 14:51:19');
INSERT INTO `sys_oper_log` VALUES (497, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":12}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [0/10 * * * * ?]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 14:55:48');
INSERT INTO `sys_oper_log` VALUES (498, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [0/10 * * * * ?]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 14:57:27');
INSERT INTO `sys_oper_log` VALUES (499, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [0 30 * * * *]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:00:14');
INSERT INTO `sys_oper_log` VALUES (500, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [0 30 * * * *]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:04:08');
INSERT INTO `sys_oper_log` VALUES (501, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [*/5 * * * * *]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:05:22');
INSERT INTO `sys_oper_log` VALUES (502, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"expected exactly 5 fields, found 6: [*/5 * * * * *]\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:08:47');
INSERT INTO `sys_oper_log` VALUES (503, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:12:22');
INSERT INTO `sys_oper_log` VALUES (504, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"任务已存在\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:14:10');
INSERT INTO `sys_oper_log` VALUES (505, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"任务已存在\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:15:57');
INSERT INTO `sys_oper_log` VALUES (506, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"任务已存在\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:15:58');
INSERT INTO `sys_oper_log` VALUES (507, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":12}', '{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:16:45');
INSERT INTO `sys_oper_log` VALUES (508, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:20:31');
INSERT INTO `sys_oper_log` VALUES (509, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"任务已存在\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:20:58');
INSERT INTO `sys_oper_log` VALUES (510, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:22:31');
INSERT INTO `sys_oper_log` VALUES (511, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"任务已存在\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:38:23');
INSERT INTO `sys_oper_log` VALUES (512, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"操作失败\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:39:28');
INSERT INTO `sys_oper_log` VALUES (513, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 15:49:23');
INSERT INTO `sys_oper_log` VALUES (514, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 15:51:20');
INSERT INTO `sys_oper_log` VALUES (515, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"任务已存在\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:51:20');
INSERT INTO `sys_oper_log` VALUES (516, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 15:52:20');
INSERT INTO `sys_oper_log` VALUES (517, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 15:54:12');
INSERT INTO `sys_oper_log` VALUES (518, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 15:57:02');
INSERT INTO `sys_oper_log` VALUES (519, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 15:57:12');
INSERT INTO `sys_oper_log` VALUES (520, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"任务已存在\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 15:57:21');
INSERT INTO `sys_oper_log` VALUES (521, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:01:34');
INSERT INTO `sys_oper_log` VALUES (522, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:03:34');
INSERT INTO `sys_oper_log` VALUES (523, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:03:46');
INSERT INTO `sys_oper_log` VALUES (524, '定时任务管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/edit', '::1', '内网IP', '{\"JobName\":\"test1\",\"JobParams\":\"\",\"JobGroup\":\"DEFAULT\",\"JobId\":10,\"InvokeTarget\":\"test1\",\"CronExpression\":\"1 * * * * \",\"MisfirePolicy\":\"2\",\"Concurrent\":\"1\",\"Status\":\"0\",\"Remark\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2020-03-24 16:11:12');
INSERT INTO `sys_oper_log` VALUES (525, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:11:15');
INSERT INTO `sys_oper_log` VALUES (526, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:11:31');
INSERT INTO `sys_oper_log` VALUES (527, '定时任务管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/edit', '::1', '内网IP', '{\"JobName\":\"test1\",\"JobParams\":\"\",\"JobGroup\":\"DEFAULT\",\"JobId\":10,\"InvokeTarget\":\"test1\",\"CronExpression\":\"1 * * * * \",\"MisfirePolicy\":\"1\",\"Concurrent\":\"1\",\"Status\":\"0\",\"Remark\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2020-03-24 16:12:46');
INSERT INTO `sys_oper_log` VALUES (528, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:14:47');
INSERT INTO `sys_oper_log` VALUES (529, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:18:06');
INSERT INTO `sys_oper_log` VALUES (530, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:18:14');
INSERT INTO `sys_oper_log` VALUES (531, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":500,\"msg\":\"任务已存在\",\"data\":null,\"otype\":0}', 1, '', '2020-03-24 16:18:20');
INSERT INTO `sys_oper_log` VALUES (532, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:18:29');
INSERT INTO `sys_oper_log` VALUES (533, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:20:56');
INSERT INTO `sys_oper_log` VALUES (534, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 16:27:49');
INSERT INTO `sys_oper_log` VALUES (535, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:20:36');
INSERT INTO `sys_oper_log` VALUES (536, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:44:53');
INSERT INTO `sys_oper_log` VALUES (537, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:46:52');
INSERT INTO `sys_oper_log` VALUES (538, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:51:59');
INSERT INTO `sys_oper_log` VALUES (539, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:52:24');
INSERT INTO `sys_oper_log` VALUES (540, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:53:27');
INSERT INTO `sys_oper_log` VALUES (541, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:55:51');
INSERT INTO `sys_oper_log` VALUES (542, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:56:55');
INSERT INTO `sys_oper_log` VALUES (543, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:58:04');
INSERT INTO `sys_oper_log` VALUES (544, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:58:29');
INSERT INTO `sys_oper_log` VALUES (545, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 18:59:59');
INSERT INTO `sys_oper_log` VALUES (546, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 19:01:46');
INSERT INTO `sys_oper_log` VALUES (547, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 19:02:18');
INSERT INTO `sys_oper_log` VALUES (548, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 19:04:44');
INSERT INTO `sys_oper_log` VALUES (549, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/stop', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"停止成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 19:05:02');
INSERT INTO `sys_oper_log` VALUES (550, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 19:06:01');
INSERT INTO `sys_oper_log` VALUES (551, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/job/start', '::1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-24 19:30:13');
INSERT INTO `sys_oper_log` VALUES (552, '生成代码', 3, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/remove', '::1', '内网IP', '{\"Ids\":\"37,38,39,40,41\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2020-03-24 22:09:41');
INSERT INTO `sys_oper_log` VALUES (553, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-24 22:09:45');
INSERT INTO `sys_oper_log` VALUES (554, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"sql: converting argument $1 type: unsupported type []string, a slice of string\",\"data\":null,\"otype\":1}', 1, '', '2020-03-24 22:09:45');
INSERT INTO `sys_oper_log` VALUES (555, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-24 22:09:49');
INSERT INTO `sys_oper_log` VALUES (556, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"sql: converting argument $1 type: unsupported type []string, a slice of string\",\"data\":null,\"otype\":1}', 1, '', '2020-03-24 22:09:49');
INSERT INTO `sys_oper_log` VALUES (557, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 16:27:02');
INSERT INTO `sys_oper_log` VALUES (558, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"sql: converting argument $1 type: unsupported type []string, a slice of string\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 16:27:02');
INSERT INTO `sys_oper_log` VALUES (559, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:35:07');
INSERT INTO `sys_oper_log` VALUES (560, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"sql: converting argument $1 type: unsupported type []string, a slice of string\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:35:07');
INSERT INTO `sys_oper_log` VALUES (561, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:40:25');
INSERT INTO `sys_oper_log` VALUES (562, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"sql: converting argument $1 type: unsupported type []string, a slice of string\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:40:25');
INSERT INTO `sys_oper_log` VALUES (563, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:40:29');
INSERT INTO `sys_oper_log` VALUES (564, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"sql: converting argument $1 type: unsupported type []string, a slice of string\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:40:29');
INSERT INTO `sys_oper_log` VALUES (565, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:41:19');
INSERT INTO `sys_oper_log` VALUES (566, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"sql: converting argument $1 type: unsupported type []string, a slice of string\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:41:19');
INSERT INTO `sys_oper_log` VALUES (567, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"sys_config\"}', '{\"code\":500,\"msg\":\"sql: converting argument $1 type: unsupported type []string, a slice of string\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:42:16');
INSERT INTO `sys_oper_log` VALUES (568, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"sys_config\"}', '{\"code\":500,\"msg\":\"Error 1054: Unknown column \'table_id\' in \'field list\'\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 17:43:50');
INSERT INTO `sys_oper_log` VALUES (569, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-26 19:38:33');
INSERT INTO `sys_oper_log` VALUES (570, '生成代码', 3, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/remove', '::1', '内网IP', '{\"Ids\":\"1\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2020-03-26 19:41:15');
INSERT INTO `sys_oper_log` VALUES (571, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-26 19:43:19');
INSERT INTO `sys_oper_log` VALUES (572, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-26 19:43:24');
INSERT INTO `sys_oper_log` VALUES (573, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-26 20:13:33');
INSERT INTO `sys_oper_log` VALUES (574, '生成代码', 3, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/remove', '::1', '内网IP', '{\"Ids\":\"2,3,4\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2020-03-26 20:35:09');
INSERT INTO `sys_oper_log` VALUES (575, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/importTable', '::1', '内网IP', '{\"tables\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-26 20:35:20');
INSERT INTO `sys_oper_log` VALUES (576, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 21:32:29');
INSERT INTO `sys_oper_log` VALUES (577, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 21:33:36');
INSERT INTO `sys_oper_log` VALUES (578, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 21:35:11');
INSERT INTO `sys_oper_log` VALUES (579, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 21:36:12');
INSERT INTO `sys_oper_log` VALUES (580, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 21:37:19');
INSERT INTO `sys_oper_log` VALUES (581, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 21:37:47');
INSERT INTO `sys_oper_log` VALUES (582, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 21:42:54');
INSERT INTO `sys_oper_log` VALUES (583, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 21:49:13');
INSERT INTO `sys_oper_log` VALUES (584, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 22:11:02');
INSERT INTO `sys_oper_log` VALUES (585, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 22:12:45');
INSERT INTO `sys_oper_log` VALUES (586, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 22:17:06');
INSERT INTO `sys_oper_log` VALUES (587, '生成代码', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/column/list', '::1', '内网IP', '{\"tableId\":0}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2020-03-26 22:17:57');
INSERT INTO `sys_oper_log` VALUES (588, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-26 22:18:46');
INSERT INTO `sys_oper_log` VALUES (589, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-26 22:19:02');
INSERT INTO `sys_oper_log` VALUES (590, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:27:16');
INSERT INTO `sys_oper_log` VALUES (591, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:28:30');
INSERT INTO `sys_oper_log` VALUES (592, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:28:43');
INSERT INTO `sys_oper_log` VALUES (593, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:29:05');
INSERT INTO `sys_oper_log` VALUES (594, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:29:39');
INSERT INTO `sys_oper_log` VALUES (595, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:36:59');
INSERT INTO `sys_oper_log` VALUES (596, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":500,\"msg\":\"json: cannot unmarshal string into Go struct field Entity.column_id of type int64\",\"data\":null,\"otype\":2}', 1, '', '2020-03-27 09:44:01');
INSERT INTO `sys_oper_log` VALUES (597, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:44:52');
INSERT INTO `sys_oper_log` VALUES (598, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":500,\"msg\":\"json: cannot unmarshal number into Go struct field Entity.is_insert of type string\",\"data\":null,\"otype\":2}', 1, '', '2020-03-27 09:46:43');
INSERT INTO `sys_oper_log` VALUES (599, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:49:14');
INSERT INTO `sys_oper_log` VALUES (600, '生成代码', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/tool/gen/edit', '::1', '内网IP', '{\"tableName\":\"sys_config\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2020-03-27 09:49:30');
INSERT INTO `sys_oper_log` VALUES (601, '生成代码', 0, 'GET', 'GET', 1, 'admin', '运维部门', '/tool/gen/genCode', '::1', '内网IP', '{\"tableId\":5}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-27 11:12:03');
INSERT INTO `sys_oper_log` VALUES (602, '生成代码', 0, 'GET', 'GET', 1, 'admin', '运维部门', '/tool/gen/genCode', '::1', '内网IP', '{\"tableId\":5}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-27 11:16:07');
INSERT INTO `sys_oper_log` VALUES (603, '生成代码', 0, 'GET', 'GET', 1, 'admin', '运维部门', '/tool/gen/genCode', '::1', '内网IP', '{\"tableId\":5}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-27 11:17:39');
INSERT INTO `sys_oper_log` VALUES (604, '生成代码', 0, 'GET', 'GET', 1, 'admin', '运维部门', '/tool/gen/genCode', '::1', '内网IP', '{\"tableId\":5}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-27 11:19:14');
INSERT INTO `sys_oper_log` VALUES (605, '生成代码', 0, 'GET', 'GET', 1, 'admin', '运维部门', '/tool/gen/genCode', '::1', '内网IP', '{\"tableId\":5}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-27 12:03:03');
INSERT INTO `sys_oper_log` VALUES (606, '生成代码', 0, 'GET', 'GET', 1, 'admin', '运维部门', '/tool/gen/genCode', '::1', '内网IP', '{\"tableId\":5}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2020-03-27 21:01:05');
INSERT INTO `sys_oper_log` VALUES (607, '部门管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/dept/edit', '::1', '内网IP', '{\"DeptId\":112,\"ParentId\":100,\"DeptName\":\"4535345\",\"OrderNum\":1,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2020-03-28 20:57:32');
INSERT INTO `sys_oper_log` VALUES (608, '部门管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/dept/edit', '::1', '内网IP', '{\"DeptId\":112,\"ParentId\":100,\"DeptName\":\"3242343243243242\",\"OrderNum\":1,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2020-03-28 20:57:37');
INSERT INTO `sys_oper_log` VALUES (609, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '运维部门', '/system/menu/remove', '::1', '内网IP', '{\"id\":1035}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2020-05-05 14:18:01');
INSERT INTO `sys_oper_log` VALUES (610, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '运维部门', '/system/menu/remove', '::1', '内网IP', '{\"id\":1038}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2020-05-05 14:18:14');
INSERT INTO `sys_oper_log` VALUES (611, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '运维部门', '/system/menu/remove', '::1', '内网IP', '{\"id\":1037}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2020-05-05 14:18:22');
INSERT INTO `sys_oper_log` VALUES (612, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '运维部门', '/system/menu/remove', '::1', '内网IP', '{\"id\":1036}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2020-05-05 14:18:28');
INSERT INTO `sys_oper_log` VALUES (613, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '运维部门', '/system/menu/remove', '::1', '内网IP', '{\"id\":107}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2020-05-05 14:18:33');
INSERT INTO `sys_oper_log` VALUES (614, '菜单管理', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/add', '127.0.0.1', '内网IP', '{\"ParentId\":0,\"MenuType\":\"C\",\"MenuName\":\"baidu\",\"OrderNum\":1,\"Url\":\"http://www.baidu.com\",\"Icon\":\"fa fa-asl-interpreting\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1063,\"otype\":1}', 0, '', '2021-06-02 00:00:56');
INSERT INTO `sys_oper_log` VALUES (615, '角色管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/edit', '127.0.0.1', '内网IP', '{\"RoleId\":1,\"RoleName\":\"管理员\",\"RoleKey\":\"admin\",\"RoleSort\":\"1\",\"Status\":\"0\",\"Remark\":\"管理员\",\"MenuIds\":\"1,100,1001,1004,1005,101,1007,1008,1009,1010,1011,102,1012,1013,103,1016,1017,1018,1019,104,1020,1021,1022,1023,1024,105,1025,1026,1029,106,1030,1031,1034,4,2,109,1047,1048,1049,110,1050,1051,1052,1053,1054,1055,1056,112,500,1039,1040,1041,1042,501,1043,1044,1045,1046,3,113,114,1057,1058,1059,1060,1061,115\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2021-06-02 00:01:45');
INSERT INTO `sys_oper_log` VALUES (616, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1063,\"ParentId\":0,\"MenuType\":\"C\",\"MenuName\":\"baidu\",\"OrderNum\":1,\"Url\":\"http://www.baidu.com\",\"Icon\":\"fa fa-asl-interpreting\",\"Target\":\"menuBlank\",\"Perms\":\"\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2021-06-02 00:02:48');
INSERT INTO `sys_oper_log` VALUES (617, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1063,\"ParentId\":0,\"MenuType\":\"C\",\"MenuName\":\"baidu\",\"OrderNum\":1,\"Url\":\"http://www.baidu.com\",\"Icon\":\"fa fa-asl-interpreting\",\"Target\":\"menuBlank\",\"Perms\":\"tool:swagger:view\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2021-06-02 00:04:21');
INSERT INTO `sys_oper_log` VALUES (618, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1063,\"ParentId\":0,\"MenuType\":\"C\",\"MenuName\":\"baidu\",\"OrderNum\":1,\"Url\":\"http://www.baidu.com\",\"Icon\":\"fa fa-asl-interpreting\",\"Target\":\"menuBlank\",\"Perms\":\"tool:swagger:view\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2021-06-02 00:05:04');
INSERT INTO `sys_oper_log` VALUES (619, '用户强退', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/monitor/online/forceLogout', '::1', '内网IP', '{\"sessionId\":\"UPZMETVQLVEVUUOZ4CEJXCOQQ5BBRZKEBGV4KLRCR6F52RYGQGYA\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2021-06-19 23:09:27');
INSERT INTO `sys_oper_log` VALUES (620, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '::1', '内网IP', '{\"MenuId\":1063,\"ParentId\":0,\"MenuType\":\"C\",\"MenuName\":\"baidu\",\"OrderNum\":1,\"Url\":\"http://www.baidu.com\",\"Icon\":\"fa fa-asl-interpreting\",\"Target\":\"menuItem\",\"Perms\":\"tool:swagger:view\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2021-06-19 23:11:01');
INSERT INTO `sys_oper_log` VALUES (621, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '运维部门', '/system/menu/remove', '::1', '内网IP', '{\"id\":1063}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2021-06-19 23:11:55');
INSERT INTO `sys_oper_log` VALUES (622, '菜单管理', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/add', '::1', '内网IP', '{\"ParentId\":4,\"MenuType\":\"C\",\"MenuName\":\"baidu\",\"OrderNum\":1,\"Url\":\"http://www.baidu.com\",\"Icon\":\"fa fa-bank\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1064,\"otype\":1}', 0, '', '2021-06-19 23:12:45');
INSERT INTO `sys_oper_log` VALUES (623, '生成代码', 0, 'GET', 'GET', 1, 'admin', '运维部门', '/tool/gen/genCode', '::1', '内网IP', '{\"tableId\":5}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2021-06-19 23:15:27');
INSERT INTO `sys_oper_log` VALUES (624, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '::1', '内网IP', '{\"MenuId\":104,\"ParentId\":1,\"MenuType\":\"C\",\"MenuName\":\"岗位管理\",\"OrderNum\":5,\"Url\":\"/system/post\",\"Icon\":\"fa fa-balance-scale\",\"Target\":\"menuItem\",\"Perms\":\"system:post:view\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2021-06-19 23:27:38');
INSERT INTO `sys_oper_log` VALUES (625, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '::1', '内网IP', '{\"MenuId\":6,\"ParentId\":5,\"MenuType\":\"M\",\"MenuName\":\"表单演示\",\"OrderNum\":1,\"Url\":\"#\",\"Icon\":\"fa fa-asterisk\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2021-06-19 23:28:01');
INSERT INTO `sys_oper_log` VALUES (626, '角色管理', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/export', '::1', '内网IP', '{\"RoleName\":\"\",\"Status\":\"\",\"RoleKey\":\"\",\"DataScope\":\"\",\"BeginTime\":\"\",\"EndTime\":\"\",\"PageNum\":0,\"PageSize\":0,\"OrderByColumn\":\"\",\"IsAsc\":\"\"}', '{\"code\":0,\"msg\":\"1624120150167529600.xls\",\"data\":null,\"otype\":0}', 0, '', '2021-06-20 00:29:10');
INSERT INTO `sys_oper_log` VALUES (627, '参数管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/config/edit', '127.0.0.1', '内网IP', '{\"ConfigId\":1,\"ConfigName\":\"主框架页-默认皮肤样式名称\",\"ConfigKey\":\"sys.index.skinName\",\"ConfigValue\":\"skin-blue\",\"ConfigType\":\"Y\",\"Remark\":\"蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2021-06-20 13:51:55');
INSERT INTO `sys_oper_log` VALUES (628, '修改用户', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/user/edit', '127.0.0.1', '内网IP', '{\"UserId\":1,\"UserName\":\"超级管理员\",\"Phonenumber\":\"15888888881\",\"Email\":\"dd122221111d@163.com\",\"DeptId\":110,\"Sex\":\"0\",\"Status\":\"0\",\"RoleIds\":\"\",\"PostIds\":\"\",\"Remark\":\"管理员111111\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2021-06-20 13:52:46');
INSERT INTO `sys_oper_log` VALUES (629, '角色管理', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/export', '127.0.0.1', '内网IP', '{\"RoleName\":\"\",\"Status\":\"\",\"RoleKey\":\"\",\"DataScope\":\"\",\"BeginTime\":\"2022-11-18\",\"EndTime\":\"2022-11-20\",\"PageNum\":0,\"PageSize\":0,\"OrderByColumn\":\"\",\"IsAsc\":\"\"}', '{\"code\":0,\"msg\":\"1668755539757563400.xls\",\"data\":null,\"otype\":0}', 0, '', '2022-11-18 15:12:20');
INSERT INTO `sys_oper_log` VALUES (630, '角色管理', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/export', '127.0.0.1', '内网IP', '{\"RoleName\":\"\",\"Status\":\"\",\"RoleKey\":\"\",\"DataScope\":\"\",\"BeginTime\":\"2022-11-18\",\"EndTime\":\"2022-11-20\",\"PageNum\":0,\"PageSize\":0,\"OrderByColumn\":\"\",\"IsAsc\":\"\"}', '{\"code\":0,\"msg\":\"1668755548153720100.xls\",\"data\":null,\"otype\":0}', 0, '', '2022-11-18 15:12:28');
INSERT INTO `sys_oper_log` VALUES (631, '角色管理', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/export', '127.0.0.1', '内网IP', '{\"RoleName\":\"\",\"Status\":\"\",\"RoleKey\":\"\",\"DataScope\":\"\",\"BeginTime\":\"2022-11-18\",\"EndTime\":\"2022-11-20\",\"PageNum\":0,\"PageSize\":0,\"OrderByColumn\":\"\",\"IsAsc\":\"\"}', '{\"code\":0,\"msg\":\"1668755559891987300.xls\",\"data\":null,\"otype\":0}', 0, '', '2022-11-18 15:12:40');
INSERT INTO `sys_oper_log` VALUES (632, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":103,\"ParentId\":1,\"MenuType\":\"C\",\"MenuName\":\"部门管理\",\"OrderNum\":4,\"Url\":\"/system/dept\",\"Icon\":\"fa fa-bank\",\"Target\":\"menuItem\",\"Perms\":\"system:dept:view\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-14 16:54:48');
INSERT INTO `sys_oper_log` VALUES (633, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":103,\"ParentId\":1,\"MenuType\":\"C\",\"MenuName\":\"部门管理\",\"OrderNum\":4,\"Url\":\"/system/dept\",\"Icon\":\"#\",\"Target\":\"menuItem\",\"Perms\":\"system:dept:view\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-14 16:55:37');
INSERT INTO `sys_oper_log` VALUES (634, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":104,\"ParentId\":1,\"MenuType\":\"C\",\"MenuName\":\"岗位管理\",\"OrderNum\":5,\"Url\":\"/system/post\",\"Icon\":\"#\",\"Target\":\"menuItem\",\"Perms\":\"system:post:view\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-14 16:55:51');
INSERT INTO `sys_oper_log` VALUES (635, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1064,\"ParentId\":4,\"MenuType\":\"C\",\"MenuName\":\"baidu\",\"OrderNum\":1,\"Url\":\"https://www.robvi.com\",\"Icon\":\"fa fa-bank\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-14 22:29:22');
INSERT INTO `sys_oper_log` VALUES (636, '保存头像', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/user/profile/updateAvatar', '127.0.0.1', '内网IP', '{\"userid\":1}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2023-09-14 23:00:39');
INSERT INTO `sys_oper_log` VALUES (637, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1064,\"ParentId\":4,\"MenuType\":\"C\",\"MenuName\":\"baidu\",\"OrderNum\":1,\"Url\":\"https://www.robvi.com\",\"Icon\":\"fa fa-globe\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-14 23:01:56');
INSERT INTO `sys_oper_log` VALUES (638, '角色管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/edit', '127.0.0.1', '内网IP', '{\"RoleId\":3,\"RoleName\":\"培训员\",\"RoleKey\":\"dfgdfg\",\"RoleSort\":\"0\",\"Status\":\"0\",\"Remark\":\"\",\"MenuIds\":\"1,100,1001,1002,1003,1004,1005,1006,101,1007,1008,1009,1010,1011,102,1012,1013,1014,1015,103,1016,1017,1018,1019,104,1020,1021,1022,1023,1024,105,1025,1026,1027,1028,1029,106,1030,1031,1032,1033,1034,2,109,1047,1048,1049,110,1050,1051,1052,1053,1054,1055,1056,112,500,1039,1040,1041,1042,501,1043,1044,1045,1046,3,113,114,1057,1058,1059,1060,1061,115\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-14 23:14:51');
INSERT INTO `sys_oper_log` VALUES (639, '角色管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/edit', '127.0.0.1', '内网IP', '{\"RoleId\":3,\"RoleName\":\"厂家运维人员\",\"RoleKey\":\"oper\",\"RoleSort\":\"0\",\"Status\":\"0\",\"Remark\":\"\",\"MenuIds\":\"1,100,1001,1002,1003,1004,1005,1006,101,1007,1008,1009,1010,1011,102,1012,1013,1014,1015,103,1016,1017,1018,1019,104,1020,1021,1022,1023,1024,105,1025,1026,1027,1028,1029,106,1030,1031,1032,1033,1034,2,109,1047,1048,1049,110,1050,1051,1052,1053,1054,1055,1056,112,500,1039,1040,1041,1042,501,1043,1044,1045,1046,3,113,114,1057,1058,1059,1060,1061,115\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 10:59:21');
INSERT INTO `sys_oper_log` VALUES (640, '角色管理', 0, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/authDataScope', '127.0.0.1', '内网IP', '{\"RoleId\":3,\"RoleName\":\"厂家运维人员\",\"RoleKey\":\"oper\",\"DataScope\":\"1\",\"DeptIds\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-09-15 10:59:47');
INSERT INTO `sys_oper_log` VALUES (641, '新增用户', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/system/user/add', '127.0.0.1', '内网IP', '{\"UserName\":\"qy1156010001\",\"Phonenumber\":\"18518276786\",\"Email\":\"uhvs@163.com\",\"LoginName\":\"qy1156010001\",\"Password\":\"admin123\",\"DeptId\":110,\"Sex\":\"0\",\"Status\":\"0\",\"RoleIds\":\"3\",\"PostIds\":\"\",\"Remark\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":2,\"otype\":1}', 0, '', '2023-09-15 11:01:47');
INSERT INTO `sys_oper_log` VALUES (642, '角色管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/role/edit', '127.0.0.1', '内网IP', '{\"RoleId\":3,\"RoleName\":\"厂家运维人员\",\"RoleKey\":\"oper\",\"RoleSort\":\"0\",\"Status\":\"0\",\"Remark\":\"\",\"MenuIds\":\"4,1064\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:03:23');
INSERT INTO `sys_oper_log` VALUES (643, '部门管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":100,\"ParentId\":0,\"DeptName\":\"DPC\",\"OrderNum\":0,\"Leader\":\"admin\",\"Phone\":\"15888888888\",\"Email\":\"110@qq.com\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:34:01');
INSERT INTO `sys_oper_log` VALUES (644, '岗位管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/post/edit', '127.0.0.1', '内网IP', '{\"PostId\":2,\"PostName\":\"项目经理\",\"PostCode\":\"se\",\"PostSort\":2,\"Status\":\"0\",\"Remark\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:34:52');
INSERT INTO `sys_oper_log` VALUES (645, '岗位管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/post/edit', '127.0.0.1', '内网IP', '{\"PostId\":1,\"PostName\":\"经理\",\"PostCode\":\"manager\",\"PostSort\":1,\"Status\":\"0\",\"Remark\":\"manager\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:35:40');
INSERT INTO `sys_oper_log` VALUES (646, '岗位管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/post/edit', '127.0.0.1', '内网IP', '{\"PostId\":2,\"PostName\":\"运维\",\"PostCode\":\"operator\",\"PostSort\":2,\"Status\":\"0\",\"Remark\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:36:02');
INSERT INTO `sys_oper_log` VALUES (647, '部门管理', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/system/dept/add', '127.0.0.1', '内网IP', '{\"ParentId\":100,\"DeptName\":\"江苏铁塔\",\"OrderNum\":1,\"Leader\":\"江苏铁塔\",\"Phone\":\"\",\"Email\":\"\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}', 0, '', '2023-09-15 11:38:29');
INSERT INTO `sys_oper_log` VALUES (648, '部门管理', 1, 'POST', 'POST', 1, 'admin', '运维部门', '/system/dept/add', '127.0.0.1', '内网IP', '{\"ParentId\":100,\"DeptName\":\"爱客\",\"OrderNum\":1,\"Leader\":\"\",\"Phone\":\"\",\"Email\":\"\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":1}', 0, '', '2023-09-15 11:39:06');
INSERT INTO `sys_oper_log` VALUES (649, '部门管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":113,\"ParentId\":100,\"DeptName\":\"江苏铁塔\",\"OrderNum\":2,\"Leader\":\"江苏铁塔\",\"Phone\":\"\",\"Email\":\"\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:39:15');
INSERT INTO `sys_oper_log` VALUES (650, '部门管理', 3, 'GET', 'GET', 1, 'admin', '运维部门', '/system/dept/remove', '127.0.0.1', '内网IP', '{\"id\":113}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2023-09-15 11:39:38');
INSERT INTO `sys_oper_log` VALUES (651, '部门管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":110,\"ParentId\":100,\"DeptName\":\"江苏铁塔\",\"OrderNum\":1,\"Leader\":\"\",\"Phone\":\"\",\"Email\":\"\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":500,\"msg\":\"部门名称已存在\",\"data\":null,\"otype\":2}', 1, '', '2023-09-15 11:40:27');
INSERT INTO `sys_oper_log` VALUES (652, '部门管理', 2, 'POST', 'POST', 1, 'admin', '运维部门', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":110,\"ParentId\":100,\"DeptName\":\"江苏铁塔\",\"OrderNum\":1,\"Leader\":\"\",\"Phone\":\"\",\"Email\":\"\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":500,\"msg\":\"部门名称已存在\",\"data\":null,\"otype\":2}', 1, '', '2023-09-15 11:41:02');
INSERT INTO `sys_oper_log` VALUES (653, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":110,\"ParentId\":100,\"DeptName\":\"江苏铁塔\",\"OrderNum\":1,\"Leader\":\"\",\"Phone\":\"\",\"Email\":\"\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:51:28');
INSERT INTO `sys_oper_log` VALUES (654, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":112,\"ParentId\":100,\"DeptName\":\"动力源运维\",\"OrderNum\":1,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:54:14');
INSERT INTO `sys_oper_log` VALUES (655, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":112,\"ParentId\":100,\"DeptName\":\"动力源运维\",\"OrderNum\":100,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:54:34');
INSERT INTO `sys_oper_log` VALUES (656, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":112,\"ParentId\":100,\"DeptName\":\"动力源运维\",\"OrderNum\":0,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:54:43');
INSERT INTO `sys_oper_log` VALUES (657, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":110,\"ParentId\":100,\"DeptName\":\"江苏铁塔\",\"OrderNum\":100,\"Leader\":\"\",\"Phone\":\"\",\"Email\":\"\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:54:50');
INSERT INTO `sys_oper_log` VALUES (658, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":114,\"ParentId\":100,\"DeptName\":\"爱客\",\"OrderNum\":100,\"Leader\":\"\",\"Phone\":\"\",\"Email\":\"\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:55:13');
INSERT INTO `sys_oper_log` VALUES (659, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":112,\"ParentId\":100,\"DeptName\":\"动力源运维\",\"OrderNum\":10,\"Leader\":\"曾尚兵\",\"Phone\":\"18788996255\",\"Email\":\"ddd@163.com\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-15 11:55:24');
INSERT INTO `sys_oper_log` VALUES (660, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:39:59');
INSERT INTO `sys_oper_log` VALUES (661, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:40:00');
INSERT INTO `sys_oper_log` VALUES (662, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:40:21');
INSERT INTO `sys_oper_log` VALUES (663, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:40:22');
INSERT INTO `sys_oper_log` VALUES (664, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:40:40');
INSERT INTO `sys_oper_log` VALUES (665, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:40:40');
INSERT INTO `sys_oper_log` VALUES (666, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:43:15');
INSERT INTO `sys_oper_log` VALUES (667, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:43:16');
INSERT INTO `sys_oper_log` VALUES (668, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:47:40');
INSERT INTO `sys_oper_log` VALUES (669, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:47:40');
INSERT INTO `sys_oper_log` VALUES (670, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:47:46');
INSERT INTO `sys_oper_log` VALUES (671, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:47:46');
INSERT INTO `sys_oper_log` VALUES (672, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:48:25');
INSERT INTO `sys_oper_log` VALUES (673, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:48:26');
INSERT INTO `sys_oper_log` VALUES (674, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:48:56');
INSERT INTO `sys_oper_log` VALUES (675, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 18:48:57');
INSERT INTO `sys_oper_log` VALUES (676, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误tables未选中\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 19:03:47');
INSERT INTO `sys_oper_log` VALUES (677, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 19:03:48');
INSERT INTO `sys_oper_log` VALUES (678, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"参数错误tables未选中\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 19:10:59');
INSERT INTO `sys_oper_log` VALUES (679, '生成代码', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"\"}', '{\"code\":500,\"msg\":\"请选择需要导入的表\",\"data\":null,\"otype\":1}', 1, '', '2023-09-17 19:10:59');
INSERT INTO `sys_oper_log` VALUES (680, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/tool/gen/importTable', '127.0.0.1', '内网IP', '{\"tables\":\"sys_oper_log\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-09-17 19:16:00');
INSERT INTO `sys_oper_log` VALUES (681, '菜单管理', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/menu/add', '127.0.0.1', '内网IP', '{\"ParentId\":6,\"MenuType\":\"C\",\"MenuName\":\"搜索自动补全\",\"OrderNum\":100,\"Url\":\"/demo/form/autocomplete\",\"Icon\":\"\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1065,\"otype\":1}', 0, '', '2023-09-18 10:15:49');
INSERT INTO `sys_oper_log` VALUES (682, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1065,\"ParentId\":6,\"MenuType\":\"C\",\"MenuName\":\"搜索自动补全\",\"OrderNum\":0,\"Url\":\"/demo/form/autocomplete\",\"Icon\":\"\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-18 10:17:03');
INSERT INTO `sys_oper_log` VALUES (683, '菜单管理', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/menu/add', '127.0.0.1', '内网IP', '{\"ParentId\":4,\"MenuType\":\"C\",\"MenuName\":\"发电量重算\",\"OrderNum\":100,\"Url\":\"iot/hisData/toWizard\",\"Icon\":\"\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1066,\"otype\":1}', 0, '', '2023-09-18 21:42:44');
INSERT INTO `sys_oper_log` VALUES (684, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1066,\"ParentId\":4,\"MenuType\":\"C\",\"MenuName\":\"发电量重算\",\"OrderNum\":100,\"Url\":\"iot/hisData/toWizard\",\"Icon\":\"fa fa-bars\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-19 21:32:53');
INSERT INTO `sys_oper_log` VALUES (685, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1066,\"ParentId\":4,\"MenuType\":\"C\",\"MenuName\":\"发电量重算\",\"OrderNum\":100,\"Url\":\"iot/hisData/toWizard\",\"Icon\":\"fa fa-sticky-note-o\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-19 21:36:52');
INSERT INTO `sys_oper_log` VALUES (686, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/menu/edit', '127.0.0.1', '内网IP', '{\"MenuId\":1064,\"ParentId\":4,\"MenuType\":\"C\",\"MenuName\":\"监控平台\",\"OrderNum\":1,\"Url\":\"http://iciom.dpc.com.cn/\",\"Icon\":\"fa fa-globe\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-19 21:44:36');
INSERT INTO `sys_oper_log` VALUES (687, '修改用户', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/user/edit', '127.0.0.1', '内网IP', '{\"UserId\":2,\"UserName\":\"qy1156010001\",\"Phonenumber\":\"18518276786\",\"Email\":\"uhvs@163.com\",\"DeptId\":110,\"Sex\":\"0\",\"Status\":\"0\",\"RoleIds\":\"3\",\"PostIds\":\"\",\"Remark\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-19 22:01:29');
INSERT INTO `sys_oper_log` VALUES (688, '登陆日志管理', 3, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/monitor/logininfor/remove', '36.5.112.29', '合肥市', '{\"Ids\":\"651\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}', 0, '', '2023-09-21 11:43:20');
INSERT INTO `sys_oper_log` VALUES (689, '字典数据管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/system/dict/data/edit', '127.0.0.1', '内网IP', '{\"DictCode\":1,\"DictLabel\":\"男\",\"DictValue\":\"0\",\"DictType\":\"\",\"DictSort\":1,\"CssClass\":\"\",\"ListClass\":\"default\",\"IsDefault\":\"Y\",\"Status\":\"0\",\"Remark\":\"[[*{remark}]]\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-21 21:45:51');
INSERT INTO `sys_oper_log` VALUES (690, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/system/dept/edit', '127.0.0.1', '内网IP', '{\"DeptId\":112,\"ParentId\":100,\"DeptName\":\"动力源运维\",\"OrderNum\":10,\"Leader\":\"lostvip\",\"Phone\":\"\",\"Email\":\"331641539@qq.com\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-09-21 21:49:43');
INSERT INTO `sys_oper_log` VALUES (691, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/menu/edit', '10.0.0.2', '', '{\"MenuId\":1064,\"ParentId\":4,\"MenuType\":\"C\",\"MenuName\":\"监控平台\",\"OrderNum\":1,\"Url\":\"https://lostvip.com\",\"Icon\":\"fa fa-globe\",\"Target\":\"menuItem\",\"Perms\":\"\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-10-17 14:10:12');
INSERT INTO `sys_oper_log` VALUES (692, '岗位管理', 1, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/post/add', '10.0.0.2', '', '{\"PostName\":\"xx\",\"PostCode\":\"sd\",\"PostSort\":3,\"Status\":\"0\",\"Remark\":\"\"}', '{\"code\":500,\"msg\":\"操作失败\",\"data\":3,\"otype\":1}', 1, '', '2023-10-21 17:02:59');
INSERT INTO `sys_oper_log` VALUES (693, '岗位管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/post/edit', '10.0.0.2', '', '{\"PostId\":2,\"PostName\":\"运维\",\"PostCode\":\"operator\",\"PostSort\":3,\"Status\":\"0\",\"Remark\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-10-21 17:03:15');
INSERT INTO `sys_oper_log` VALUES (694, '岗位管理', 3, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/post/remove', '10.0.0.2', '', '{\"Ids\":\"3\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2023-10-21 17:03:27');
INSERT INTO `sys_oper_log` VALUES (695, '岗位管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/post/edit', '10.0.0.2', '', '{\"PostId\":2,\"PostName\":\"运维\",\"PostCode\":\"operator\",\"PostSort\":2,\"Status\":\"0\",\"Remark\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-10-21 17:03:38');
INSERT INTO `sys_oper_log` VALUES (696, '生成代码', 3, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/remove', '192.168.100.103', '', '{\"Ids\":\"5,6\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2023-10-29 19:30:00');
INSERT INTO `sys_oper_log` VALUES (697, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/importTable', '192.168.100.103', '', '{\"tables\":\"his_patient\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-29 21:43:56');
INSERT INTO `sys_oper_log` VALUES (698, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/importTable', '192.168.100.103', '', '{\"tables\":\"his_patient\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-29 21:44:26');
INSERT INTO `sys_oper_log` VALUES (699, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/importTable', '192.168.100.103', '', '{\"tables\":\"his_patient\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-29 21:45:07');
INSERT INTO `sys_oper_log` VALUES (700, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/importTable', '192.168.100.103', '', '{\"tables\":\"his_patient\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-29 21:46:40');
INSERT INTO `sys_oper_log` VALUES (701, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/importTable', '192.168.100.103', '', '{\"tables\":\"his_patient\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-29 21:53:11');
INSERT INTO `sys_oper_log` VALUES (702, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/importTable', '192.168.100.103', '', '{\"tables\":\"his_patient\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-29 21:55:17');
INSERT INTO `sys_oper_log` VALUES (703, '生成代码', 3, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/remove', '192.168.100.103', '', '{\"Ids\":\"12\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2023-10-29 22:04:34');
INSERT INTO `sys_oper_log` VALUES (704, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/importTable', '192.168.100.103', '', '{\"tables\":\"his_patient\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-29 22:05:23');
INSERT INTO `sys_oper_log` VALUES (705, '部门管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/dept/edit', '10.0.0.2', '', '{\"DeptId\":112,\"ParentId\":100,\"DeptName\":\"公司运维\",\"OrderNum\":10,\"Leader\":\"lostvip\",\"Phone\":\"18788996255\",\"Email\":\"331641539@qq.com\",\"Status\":\"0\",\"TenantId\":0}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-10-29 22:30:16');
INSERT INTO `sys_oper_log` VALUES (706, '生成代码', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/edit', '192.168.100.103', '', '{\"tableName\":\"gen_table\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":2}', 0, '', '2023-10-30 19:15:33');
INSERT INTO `sys_oper_log` VALUES (707, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 19:17:13');
INSERT INTO `sys_oper_log` VALUES (708, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 19:36:31');
INSERT INTO `sys_oper_log` VALUES (709, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 19:41:41');
INSERT INTO `sys_oper_log` VALUES (710, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 19:42:32');
INSERT INTO `sys_oper_log` VALUES (711, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 19:43:06');
INSERT INTO `sys_oper_log` VALUES (712, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 19:48:25');
INSERT INTO `sys_oper_log` VALUES (713, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 19:57:49');
INSERT INTO `sys_oper_log` VALUES (714, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:04:56');
INSERT INTO `sys_oper_log` VALUES (715, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:05:21');
INSERT INTO `sys_oper_log` VALUES (716, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:08:31');
INSERT INTO `sys_oper_log` VALUES (717, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:14:52');
INSERT INTO `sys_oper_log` VALUES (718, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:19:06');
INSERT INTO `sys_oper_log` VALUES (719, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:22:39');
INSERT INTO `sys_oper_log` VALUES (720, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:27:28');
INSERT INTO `sys_oper_log` VALUES (721, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:29:05');
INSERT INTO `sys_oper_log` VALUES (722, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:44:11');
INSERT INTO `sys_oper_log` VALUES (723, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:48:28');
INSERT INTO `sys_oper_log` VALUES (724, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:48:44');
INSERT INTO `sys_oper_log` VALUES (725, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:51:51');
INSERT INTO `sys_oper_log` VALUES (726, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:56:53');
INSERT INTO `sys_oper_log` VALUES (727, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 20:57:09');
INSERT INTO `sys_oper_log` VALUES (728, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:20:16');
INSERT INTO `sys_oper_log` VALUES (729, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:20:39');
INSERT INTO `sys_oper_log` VALUES (730, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:21:34');
INSERT INTO `sys_oper_log` VALUES (731, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:21:48');
INSERT INTO `sys_oper_log` VALUES (732, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:28:21');
INSERT INTO `sys_oper_log` VALUES (733, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:32:31');
INSERT INTO `sys_oper_log` VALUES (734, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:32:46');
INSERT INTO `sys_oper_log` VALUES (735, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:34:06');
INSERT INTO `sys_oper_log` VALUES (736, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:34:17');
INSERT INTO `sys_oper_log` VALUES (737, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:36:18');
INSERT INTO `sys_oper_log` VALUES (738, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.100.103', '', '{\"tableId\":13}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-10-30 21:36:27');
INSERT INTO `sys_oper_log` VALUES (739, '菜单管理', 2, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/system/menu/edit', '10.0.0.2', '', '{\"MenuId\":100,\"ParentId\":1,\"MenuType\":\"C\",\"MenuName\":\"用户管理\",\"OrderNum\":1,\"Url\":\"system/user\",\"Icon\":\"fa fa-address-book\",\"Target\":\"menuItem\",\"Perms\":\"system:user:view\",\"Visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2023-11-04 16:49:17');
INSERT INTO `sys_oper_log` VALUES (740, '导入表结构', 0, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/importTable', '192.168.230.1', '', '{\"tables\":\"his_patient\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 0, '', '2023-11-18 10:23:30');
INSERT INTO `sys_oper_log` VALUES (741, '生成代码', 3, 'POST', 'POST', 1, 'admin', '江苏铁塔', '/go/tool/gen/remove', '192.168.230.1', '', '{\"Ids\":\"14\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 0, '', '2023-11-18 12:15:15');
INSERT INTO `sys_oper_log` VALUES (742, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/system/menu/remove', '192.168.230.1', '', '{\"id\":1070}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-18 21:37:25');
INSERT INTO `sys_oper_log` VALUES (743, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/system/menu/remove', '192.168.230.1', '', '{\"id\":1069}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-18 21:37:38');
INSERT INTO `sys_oper_log` VALUES (744, '生成代码', 0, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/tool/gen/genCode', '192.168.230.1', '', '{\"tableId\":13}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":0}', 1, '', '2023-11-18 21:44:12');
INSERT INTO `sys_oper_log` VALUES (745, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/system/menu/remove', '192.168.230.1', '', '{\"id\":1068}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-18 21:55:18');
INSERT INTO `sys_oper_log` VALUES (746, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/system/menu/remove', '192.168.230.1', '', '{\"id\":1088}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-18 22:49:21');
INSERT INTO `sys_oper_log` VALUES (747, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/system/menu/remove', '192.168.230.1', '', '{\"id\":1100}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-18 22:50:26');
INSERT INTO `sys_oper_log` VALUES (748, '菜单管理', 3, 'GET', 'GET', 1, 'admin', '江苏铁塔', '/go/system/menu/remove', '192.168.230.1', '', '{\"id\":1094}', '{\"code\":200,\"msg\":\"操作成功\",\"data\":null,\"otype\":3}', 1, '', '2023-11-18 22:50:48');

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
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint(20) NULL DEFAULT 0 COMMENT '租户id',
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '岗位信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES (1, 'manager', '经理', 1, '0', 'admin', '2018-03-16 11:33:00', '', '2023-09-15 11:35:40', 'manager', 0);
INSERT INTO `sys_post` VALUES (2, 'operator', '运维', 2, '0', 'admin', '2018-03-16 11:33:00', 'admin', '2023-10-21 17:03:37', '', 0);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色名称',
  `role_key` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色权限字符串',
  `role_sort` int(11) NOT NULL COMMENT '显示顺序',
  `data_scope` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1' COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色状态（0正常 1停用）',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色信息表' ROW_FORMAT = Dynamic;

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
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `dept_id` bigint(20) NULL DEFAULT NULL COMMENT '部门ID',
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
  `login_ip` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '最后登陆IP',
  `login_date` datetime(0) NULL DEFAULT NULL COMMENT '最后登陆时间',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `tenant_id` bigint(20) NULL DEFAULT 0 COMMENT '租户id',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户信息表' ROW_FORMAT = Dynamic;

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
  `sessionId` varchar(250) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户会话id',
  `login_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登录账号',
  `dept_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `ipaddr` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '操作系统',
  `status` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '在线状态on_line在线off_line离线',
  `start_timestamp` datetime(0) NULL DEFAULT NULL COMMENT 'session创建时间',
  `last_access_time` datetime(0) NULL DEFAULT NULL COMMENT 'session最后访问时间',
  `expire_time` int(11) NULL DEFAULT 0 COMMENT '超时时间，单位为分钟',
  PRIMARY KEY (`sessionId`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '在线用户记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_online
-- ----------------------------
INSERT INTO `sys_user_online` VALUES ('7IHPSW3T77IZ5EGASLEYR2TGROQCIWQQK2CPEDOLP4CQEPTL3QGQ', 'admin', '', '::1', '内网IP', 'Chrome', 'Windows 10', 'on_line', '2021-06-20 15:59:42', '2021-06-20 15:59:42', 1440);

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post`  (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `post_id` bigint(20) NOT NULL COMMENT '岗位ID',
  PRIMARY KEY (`user_id`, `post_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户与岗位关联表' ROW_FORMAT = Dynamic;

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

SET FOREIGN_KEY_CHECKS = 1;
