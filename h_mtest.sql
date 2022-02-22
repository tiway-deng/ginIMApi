/*
 Navicat Premium Data Transfer

 Source Server         : local-docker
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : h_mtest

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 22/02/2022 17:14:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for im_article
-- ----------------------------
DROP TABLE IF EXISTS `im_article`;
CREATE TABLE `im_article`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '笔记ID',
  `user_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `class_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分类ID',
  `tags_id` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '笔记关联标签',
  `title` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '笔记标题',
  `abstract` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '笔记摘要',
  `image` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '笔记首图',
  `is_asterisk` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否星标笔记[0:否;1:是]',
  `status` tinyint(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '笔记状态[1:正常;2:已删除]',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '添加时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '最后一次更新时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '笔记删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_id_class_id_title`(`user_id`, `class_id`, `title`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户笔记表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_article_annex
-- ----------------------------
DROP TABLE IF EXISTS `im_article_annex`;
CREATE TABLE `im_article_annex`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '文件ID',
  `user_id` int(0) UNSIGNED NOT NULL COMMENT '上传文件的用户ID',
  `article_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '笔记ID',
  `file_suffix` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '文件后缀名',
  `file_size` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件大小（单位字节）',
  `save_dir` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件保存地址（相对地址）',
  `original_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '原文件名',
  `status` tinyint(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '附件状态[1:正常;2:已删除]',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '附件上传时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '附件删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_id_article_id`(`user_id`, `article_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '笔记附件信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_article_class
-- ----------------------------
DROP TABLE IF EXISTS `im_article_class`;
CREATE TABLE `im_article_class`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '笔记分类ID',
  `user_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `class_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '分类名',
  `sort` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
  `is_default` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '默认分类[1:是;0:不是]',
  `created_at` int(0) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_id_sort`(`user_id`, `sort`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '笔记分类表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of im_article_class
-- ----------------------------
INSERT INTO `im_article_class` VALUES (1, 1, '我的笔记', 1, 1, 1641810323);

-- ----------------------------
-- Table structure for im_article_detail
-- ----------------------------
DROP TABLE IF EXISTS `im_article_detail`;
CREATE TABLE `im_article_detail`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '笔记详情ID',
  `article_id` int(0) UNSIGNED NOT NULL COMMENT '笔记ID',
  `md_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'Markdown 内容',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'Markdown 解析HTML内容',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unique_article_id`(`article_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '笔记详情表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_article_tags
-- ----------------------------
DROP TABLE IF EXISTS `im_article_tags`;
CREATE TABLE `im_article_tags`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '笔记标签ID',
  `user_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `tag_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '标签名',
  `sort` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
  `created_at` int(0) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '笔记标签表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_chat_records
-- ----------------------------
DROP TABLE IF EXISTS `im_chat_records`;
CREATE TABLE `im_chat_records`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '聊天记录ID',
  `source` tinyint(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '消息来源[1:好友消息;2:群聊消息]',
  `msg_type` tinyint(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '消息类型[1:文本消息;2:文件消息;3:系统提示好友入群消息或系统提示好友退群消息;4:会话记录转发]',
  `user_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '发送消息的用户ID[0:代表系统消息]',
  `receive_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '接收消息的用户ID或群聊ID',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '文本消息',
  `is_revoke` tinyint(0) NOT NULL DEFAULT 0 COMMENT '是否撤回消息[0:否;1:是]',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '发送消息的时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_userid_receiveid`(`user_id`, `receive_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户聊天记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of im_chat_records
-- ----------------------------
INSERT INTO `im_chat_records` VALUES (1, 1, 1, 2, 1, '12122', 0, '2022-01-10 18:33:50');
INSERT INTO `im_chat_records` VALUES (2, 1, 1, 2, 1, '诶从而', 0, '2022-01-10 18:33:52');
INSERT INTO `im_chat_records` VALUES (3, 1, 1, 2, 1, 'oiei', 0, '2022-01-10 18:33:56');
INSERT INTO `im_chat_records` VALUES (4, 1, 1, 2, 1, '212', 0, '2022-01-10 18:34:15');
INSERT INTO `im_chat_records` VALUES (5, 1, 1, 2, 1, '223', 0, '2022-01-10 18:34:43');
INSERT INTO `im_chat_records` VALUES (6, 1, 1, 1, 2, '2111', 0, '2022-01-10 18:34:53');
INSERT INTO `im_chat_records` VALUES (7, 1, 1, 1, 2, '888', 0, '2022-01-10 18:34:55');
INSERT INTO `im_chat_records` VALUES (8, 1, 5, 2, 1, NULL, 0, '2022-01-12 14:01:28');
INSERT INTO `im_chat_records` VALUES (9, 1, 1, 2, 1, '121', 0, '2022-01-18 09:51:28');
INSERT INTO `im_chat_records` VALUES (10, 1, 1, 1, 2, '11', 0, '2022-02-22 16:16:32');
INSERT INTO `im_chat_records` VALUES (11, 1, 1, 1, 2, '343434', 0, '2022-02-22 16:16:43');

-- ----------------------------
-- Table structure for im_chat_records_code
-- ----------------------------
DROP TABLE IF EXISTS `im_chat_records_code`;
CREATE TABLE `im_chat_records_code`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '入群或退群通知ID',
  `record_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '消息记录ID',
  `user_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上传文件的用户ID',
  `code_lang` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '代码片段类型(如：php,java,python)',
  `code` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '代码片段内容',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_recordid`(`record_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户聊天记录_代码块消息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of im_chat_records_code
-- ----------------------------
INSERT INTO `im_chat_records_code` VALUES (1, 8, 2, 'less', '545455225', '2022-01-12 14:01:28');

-- ----------------------------
-- Table structure for im_chat_records_delete
-- ----------------------------
DROP TABLE IF EXISTS `im_chat_records_delete`;
CREATE TABLE `im_chat_records_delete`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '聊天删除记录ID',
  `record_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '聊天记录ID',
  `user_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_record_user_id`(`record_id`, `user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户聊天记录_删除记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_chat_records_file
-- ----------------------------
DROP TABLE IF EXISTS `im_chat_records_file`;
CREATE TABLE `im_chat_records_file`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '文件ID',
  `record_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '消息记录ID',
  `user_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上传文件的用户ID',
  `file_source` tinyint(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '文件来源[1:用户上传;2:表情包]',
  `file_type` tinyint(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '消息类型[1:图片;2:视频;3:文件]',
  `save_type` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件保存方式（0:本地 1:第三方[阿里OOS、七牛云] ）',
  `original_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '原文件名',
  `file_suffix` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '文件后缀名',
  `file_size` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件大小（单位字节）',
  `save_dir` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '文件保存地址（相对地址/第三方网络地址）',
  `is_delete` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件是否已删除[0:否;1:已删除]',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_record_id`(`record_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户聊天记录_文件消息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_chat_records_forward
-- ----------------------------
DROP TABLE IF EXISTS `im_chat_records_forward`;
CREATE TABLE `im_chat_records_forward`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '合并转发ID',
  `record_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '消息记录ID',
  `user_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '转发用户ID',
  `records_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '转发的聊天记录ID，多个用\',\'分割',
  `text` json NOT NULL COMMENT '记录快照',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '转发时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_id_records_id`(`user_id`, `records_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户聊天记录_转发信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_chat_records_invite
-- ----------------------------
DROP TABLE IF EXISTS `im_chat_records_invite`;
CREATE TABLE `im_chat_records_invite`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '入群或退群通知ID',
  `record_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '消息记录ID',
  `type` tinyint(0) NOT NULL DEFAULT 1 COMMENT '通知类型[1:入群通知;2:自动退群;3:管理员踢群]',
  `operate_user_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作人的用户ID(邀请人)',
  `user_ids` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户ID，多个用\',\'分割',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_recordid`(`record_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户聊天记录_入群或退群消息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_emoticon
-- ----------------------------
DROP TABLE IF EXISTS `im_emoticon`;
CREATE TABLE `im_emoticon`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '表情分组ID',
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '表情分组名称',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '图片地址',
  `created_at` int(0) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '表情包' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_emoticon_details
-- ----------------------------
DROP TABLE IF EXISTS `im_emoticon_details`;
CREATE TABLE `im_emoticon_details`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '表情包ID',
  `emoticon_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '表情分组ID',
  `user_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID（0：代码系统表情包）',
  `describe` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '表情关键字描述',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '表情链接',
  `file_suffix` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '文件后缀名',
  `file_size` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件大小（单位字节）',
  `created_at` int(0) UNSIGNED NULL DEFAULT 0 COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '聊天表情包' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_file_split_upload
-- ----------------------------
DROP TABLE IF EXISTS `im_file_split_upload`;
CREATE TABLE `im_file_split_upload`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '临时文件ID',
  `file_type` tinyint(0) UNSIGNED NOT NULL DEFAULT 2 COMMENT '数据类型[1:合并文件;2:拆分文件]',
  `user_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上传的用户ID',
  `hash_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '临时文件hash名',
  `original_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '原文件名',
  `split_index` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '当前索引块',
  `split_num` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '总上传索引块',
  `save_dir` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '文件的临时保存路径',
  `file_ext` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '文件后缀名',
  `file_size` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '临时文件大小',
  `is_delete` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件是否已被删除[0:否;1:是]',
  `upload_at` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '文件上传时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_id_hash_name`(`user_id`, `hash_name`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文件拆分上传' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_group
-- ----------------------------
DROP TABLE IF EXISTS `im_group`;
CREATE TABLE `im_group`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '群ID',
  `creator_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者ID(群主ID)',
  `group_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '群名称',
  `profile` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '群介绍',
  `avatar` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '群头像',
  `max_num` smallint(0) UNSIGNED NOT NULL DEFAULT 200 COMMENT '最大群成员数量',
  `is_overt` tinyint(0) NOT NULL DEFAULT 0 COMMENT '是否公开可见[0:否;1:是;]',
  `is_mute` tinyint(0) NOT NULL DEFAULT 0 COMMENT '是否全员禁言 [0:否;1:是;]，提示:不包含群主或管理员',
  `is_dismiss` tinyint(0) NOT NULL DEFAULT 0 COMMENT '是否已解散[0:否;1:是;]',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `dismissed_at` datetime(0) NULL DEFAULT NULL COMMENT '解散时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '聊天群组表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_group_member
-- ----------------------------
DROP TABLE IF EXISTS `im_group_member`;
CREATE TABLE `im_group_member`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '群成员ID',
  `group_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '群ID',
  `user_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `leader` tinyint(0) NOT NULL COMMENT '成员属性[0:普通成员;1:管理员;2:群主;]',
  `is_mute` tinyint(0) NOT NULL DEFAULT 0 COMMENT '是否禁言[0:否;1:是;]',
  `is_quit` tinyint(0) NOT NULL DEFAULT 0 COMMENT '是否退群[0:否;1:是;]',
  `user_card` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '群名片',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '入群时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '退群时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_group_id_user_id`(`group_id`, `user_id`) USING BTREE,
  INDEX `idx_user_id_group_id`(`user_id`, `group_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '聊天群组成员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_group_notice
-- ----------------------------
DROP TABLE IF EXISTS `im_group_notice`;
CREATE TABLE `im_group_notice`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '群公告ID',
  `group_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '群组ID',
  `creator_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者用户ID',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '公告标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '公告内容',
  `is_top` tinyint(0) NOT NULL DEFAULT 0 COMMENT '是否置顶[0:否;1:是;]',
  `is_delete` tinyint(0) NOT NULL DEFAULT 0 COMMENT '是否删除[0:否;1:是;]',
  `is_confirm` tinyint(0) NOT NULL DEFAULT 0 COMMENT '是否需群成员确认公告[0:否;1:是;]',
  `confirm_users` json NULL COMMENT '已确认成员',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_group_id_is_delete_is_top_updated_at`(`group_id`, `is_delete`, `is_top`, `updated_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '群组公告表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_migrations
-- ----------------------------
DROP TABLE IF EXISTS `im_migrations`;
CREATE TABLE `im_migrations`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of im_migrations
-- ----------------------------
INSERT INTO `im_migrations` VALUES (1, '2020_11_04_152602_create_users_table', 1);
INSERT INTO `im_migrations` VALUES (2, '2020_11_04_153238_create_article_table', 1);
INSERT INTO `im_migrations` VALUES (3, '2020_11_04_153251_create_article_annex_table', 1);
INSERT INTO `im_migrations` VALUES (4, '2020_11_04_153304_create_article_class_table', 1);
INSERT INTO `im_migrations` VALUES (5, '2020_11_04_153316_create_article_detail_table', 1);
INSERT INTO `im_migrations` VALUES (6, '2020_11_04_153327_create_article_tags_table', 1);
INSERT INTO `im_migrations` VALUES (7, '2020_11_04_153337_create_emoticon_table', 1);
INSERT INTO `im_migrations` VALUES (8, '2020_11_04_153347_create_emoticon_details_table', 1);
INSERT INTO `im_migrations` VALUES (9, '2020_11_04_153358_create_file_split_upload_table', 1);
INSERT INTO `im_migrations` VALUES (10, '2020_11_04_153408_create_user_login_log_table', 1);
INSERT INTO `im_migrations` VALUES (11, '2020_11_04_153421_create_users_chat_list_table', 1);
INSERT INTO `im_migrations` VALUES (12, '2020_11_04_153431_create_chat_records_table', 1);
INSERT INTO `im_migrations` VALUES (13, '2020_11_04_153442_create_chat_records_file_table', 1);
INSERT INTO `im_migrations` VALUES (14, '2020_11_04_153453_create_chat_records_delete_table', 1);
INSERT INTO `im_migrations` VALUES (15, '2020_11_04_153504_create_chat_records_forward_table', 1);
INSERT INTO `im_migrations` VALUES (16, '2020_11_04_153516_create_chat_records_invite_table', 1);
INSERT INTO `im_migrations` VALUES (17, '2020_11_04_153529_create_chat_records_code_table', 1);
INSERT INTO `im_migrations` VALUES (18, '2020_11_04_153541_create_users_emoticon_table', 1);
INSERT INTO `im_migrations` VALUES (19, '2020_11_04_153553_create_users_friends_table', 1);
INSERT INTO `im_migrations` VALUES (20, '2020_11_04_153605_create_users_friends_apply_table', 1);
INSERT INTO `im_migrations` VALUES (21, '2020_11_04_153616_create_group_table', 1);
INSERT INTO `im_migrations` VALUES (22, '2020_11_04_153626_create_group_member_table', 1);
INSERT INTO `im_migrations` VALUES (23, '2020_11_04_153636_create_group_notice_table', 1);

-- ----------------------------
-- Table structure for im_user_login_log
-- ----------------------------
DROP TABLE IF EXISTS `im_user_login_log`;
CREATE TABLE `im_user_login_log`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '登录日志ID',
  `user_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `ip` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '登录地址IP',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '登录时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户登录日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_users
-- ----------------------------
DROP TABLE IF EXISTS `im_users`;
CREATE TABLE `im_users`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `mobile` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `nickname` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户头像地址',
  `gender` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户性别[0:未知;1:男;2:女]',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户密码',
  `motto` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户座右铭',
  `email` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户邮箱',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '注册时间',
  `status` tinyint(1) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_mobile`(`mobile`) USING BTREE,
  UNIQUE INDEX `users_mobile_unique`(`mobile`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of im_users
-- ----------------------------
INSERT INTO `im_users` VALUES (1, '13533333333', '13533333333', '', 0, 'e10adc3949ba59abbe56e057f20f883e', '', '', '2022-01-10 18:25:23', 1, '2022-01-18 09:49:10');
INSERT INTO `im_users` VALUES (2, '13522222222', '13522222222', '', 0, 'e10adc3949ba59abbe56e057f20f883e', '', '', '2022-01-10 18:25:23', 1, '2022-01-18 09:49:31');
INSERT INTO `im_users` VALUES (3, '13570081693', 'tiway', '', 0, 'e10adc3949ba59abbe56e057f20f883e', '', '', '2022-01-10 18:25:23', 1, '2022-01-18 09:49:31');
INSERT INTO `im_users` VALUES (4, '13566696969', '13588', '', 0, '0192023a7bbd73250516f069df18b500', '', '', '2022-02-22 17:04:31', 0, NULL);
INSERT INTO `im_users` VALUES (5, '18798272054', '18798272054', '', 0, '0192023a7bbd73250516f069df18b500', '', '', '2022-02-22 17:05:06', 1, NULL);

-- ----------------------------
-- Table structure for im_users_chat_list
-- ----------------------------
DROP TABLE IF EXISTS `im_users_chat_list`;
CREATE TABLE `im_users_chat_list`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '聊天列表ID',
  `type` tinyint(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '聊天类型[1:好友;2:群聊]',
  `uid` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `friend_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '朋友的用户ID',
  `group_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '聊天分组ID',
  `status` int(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态[0:已删除;1:正常]',
  `is_top` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否置顶[0:否;1:是]',
  `not_disturb` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否消息免打扰[0:否;1:是]',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_uid_type_friend_id_group_id`(`uid`, `friend_id`, `group_id`, `type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户聊天列表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of im_users_chat_list
-- ----------------------------
INSERT INTO `im_users_chat_list` VALUES (1, 1, 2, 1, 0, 1, 1, 0, '2022-01-10 18:33:41', '2022-01-12 13:39:04');
INSERT INTO `im_users_chat_list` VALUES (2, 1, 1, 2, 0, 1, 0, 0, '2022-01-10 18:33:50', '2022-01-10 18:33:50');

-- ----------------------------
-- Table structure for im_users_emoticon
-- ----------------------------
DROP TABLE IF EXISTS `im_users_emoticon`;
CREATE TABLE `im_users_emoticon`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '表情包收藏ID',
  `user_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `emoticon_ids` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '表情包ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `users_emoticon_user_id_unique`(`user_id`) USING BTREE,
  INDEX `idx_user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户收藏表情包' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_users_friends
-- ----------------------------
DROP TABLE IF EXISTS `im_users_friends`;
CREATE TABLE `im_users_friends`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '关系ID',
  `user1` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户1(user1 一定比 user2 小)',
  `user2` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户2(user1 一定比 user2 小)',
  `user1_remark` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '好友备注',
  `user2_remark` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '好友备注',
  `active` tinyint(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '主动邀请方[1:user1;2:user2]',
  `status` tinyint(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '好友状态[0:已解除好友关系;1:好友状态]',
  `agree_time` datetime(0) NOT NULL COMMENT '成为好友时间',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user1_user2`(`user1`, `user2`) USING BTREE,
  INDEX `idx_user2_user1`(`user2`, `user1`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户好友关系表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of im_users_friends
-- ----------------------------
INSERT INTO `im_users_friends` VALUES (1, 1, 2, '13522222222', '4545', 1, 1, '2022-01-10 18:33:35', '2022-01-10 18:33:35');

-- ----------------------------
-- Table structure for im_users_friends_apply
-- ----------------------------
DROP TABLE IF EXISTS `im_users_friends_apply`;
CREATE TABLE `im_users_friends_apply`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '申请ID',
  `user_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '申请人ID',
  `friend_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '被申请人',
  `status` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '申请状态[0:等待处理;1:已同意]',
  `remarks` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '申请人备注信息',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '申请时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '处理时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_id`(`user_id`) USING BTREE,
  INDEX `idx_friend_id`(`friend_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户添加好友申请表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of im_users_friends_apply
-- ----------------------------
INSERT INTO `im_users_friends_apply` VALUES (1, 1, 2, 1, '4545', '2022-01-10 18:30:49', '2022-01-10 18:33:35');

SET FOREIGN_KEY_CHECKS = 1;
