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
INSERT INTO `gen_table` VALUES (13, 'his_patient', '患者基本信息', 'HisPatient', 'crud', 'robvi', 'mywork', 'patient', '患者基本信息', 'lv', '', 'admin', '2023-10-29 22:05:20', 'admin', '2023-10-30 19:15:28', '');

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
INSERT INTO `sys_config` VALUES (1, '主框架页-默认皮肤样式名称', 'system.index.skinName', 'skin-blue', 'Y', 'admin', '2018-03-16 11:33:00', '', '2021-06-20 13:51:55', '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow');
INSERT INTO `sys_config` VALUES (2, '用户管理-账号初始密码', 'system.user.initPassword', '123456', 'Y', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '初始化密码 123456');
INSERT INTO `sys_config` VALUES (3, '主框架页-侧边栏主题', 'system.index.sideTheme', 'theme-dark', 'Y', 'admin', '2018-03-16 11:33:00', '', '2020-02-05 10:46:28', '深黑主题theme-dark，浅色主题theme-light，深蓝主题theme-blue');
INSERT INTO `sys_config` VALUES (4, '静态资源网盘存储', 'system.resource.url', '/static', 'Y', 'admin', '2020-02-18 20:10:33', '', '2020-03-23 20:51:39', 'public目录下的静态资源存储到OSS/COS等网盘，将键值设为/public表示本地');

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
INSERT INTO `sys_menu` VALUES (45, '查询条件', 7, 1, 'demo/table/lv_sql', '', 'C', '0', 'demo:lv_sql:view', 'fa fa-sticky-note-o', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '');
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
  `login_ip` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '最后登录IP',
  `login_date` datetime(0) NULL DEFAULT NULL COMMENT '最后登录时间',
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
