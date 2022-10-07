CREATE TABLE `xxx_template` (
                                      `xxx_template_id` bigint(20) unsigned NOT NULL  AUTO_INCREMENT COMMENT '自增 ID',
                                      `customer_id` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'gid',
                                      `name` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'Template名称',
                                      `analysis_type` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'media_type_id:媒体类型,media_flag:媒体范围,sentiment_type_id:情感类型,link_type_id:文章类型,clustering:文章去重,precise:精准度,chart_type:图表类型,engagement_num:互动高于,subject_frequency:主体词频',
                                      `created_uid` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'uid',
                                      `updated_uid` varchar(255) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT 'uid',
                                      `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '记录是否被删除,默认 0',
                                      `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建的时间 (自动生成)',
                                      `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录的更新时间 (自动生成)',
                                      PRIMARY KEY (`xxx_template_id`),
                                      KEY `IDX_name` (`name`),
                                      KEY `IDX_customer_id` (`customer_id`),
                                      KEY `IDX_created_uid` (`created_uid`),
                                      KEY `IDX_deleted` (`deleted`),
                                      KEY `IDX_created_time` (`created_time`)

) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ;

CREATE TABLE `xxx_template_layout` (
                                      `id` bigint(20) unsigned NOT NULL  AUTO_INCREMENT COMMENT '自增 ID',
                                      `xxx_template_id` bigint(20) unsigned NOT NULL  DEFAULT 0 COMMENT 'xxx表主键',
                                      `analysis_type` varchar(50) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'media_type_id:媒体类型,media_flag:媒体范围,sentiment_type_id:情感类型,link_type_id:文章类型,clustering:文章去重,precise:精准度,chart_type:图表类型,engagement_num:互动高于,subject_frequency:主体词频',
                                      `sort` tinyint(1) NOT NULL DEFAULT '1' COMMENT '顺序 從 1 開始',
                                      `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '记录是否被删除,默认 0',
                                      `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建的时间 (自动生成)',
                                      `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录的更新时间 (自动生成)',
                                      PRIMARY KEY (`id`),
                                      KEY `IDX_xxx_template_id` (`xxx_template_id`),
                                      KEY `IDX_deleted` (`deleted`),
                                      KEY `IDX_created_time` (`created_time`)

) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='报告模板设置表';



CREATE TABLE `xxx_setup` (
                                      `xxx_setup_id` bigint(20) unsigned NOT NULL  AUTO_INCREMENT COMMENT '自增 ID',
                                      `xxx_setup_version` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '版本',
                                      `customer_id` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'gid',
                                      `enable` tinyint(1) NOT NULL DEFAULT '1' COMMENT '啓用狀態',
                                      `status` char(50) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT  '状态: draft,done',
                                      `xxx_type` char(50) COLLATE utf8mb4_bin  NOT NULL DEFAULT 'real-time' COMMENT  '实时:real-time, 日报:daily',
                                      `name` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '报告名称',
                                      `lang` char(50) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT  'cn, en-zh, zh, en-cn',
                                      `title_time` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'title 时间',
                                      `background_img` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '背景',
                                      `logo_img` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '{"bucket":"common-filestore-prod","key":"xxx/2021-07-02/bjrnd3/autofolder0702183630.xlsx"}',
                                      `start_time` char(50) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT '实时報告 数据范围 开始时间: 2022-08-28 00:00:00.000',
                                      `end_time` char(50) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT '实时報告 数据范围 结束时间: 2022-08-28 23:59:59.999',
                                      `scheduling_type` char(50) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT '發送类型: manual,auto',
                                      `data_range_start_date_type` char(50) COLLATE utf8mb4_bin COMMENT '日报 数据范围: yesterday, today',
                                      `data_range_start_time` char(50) COLLATE utf8mb4_bin COMMENT '日报 数据范围: 00:00',
                                      `data_range_end_date_type` char(50) COLLATE utf8mb4_bin COMMENT '日报 数据范围: yesterday, today',
                                      `data_range_end_time` char(50) COLLATE utf8mb4_bin COMMENT '日报 数据范围: 00:00',
                                      `scheduling_frequency` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '发送頻次 多筆: Mon,Tue,Wed,Thur,Fri,Sat,Sun,',
                                      `notification_frequency` char(128) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT '发送时间: 09:00',
                                      `notification_type` char(50) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT '发送渠道 多筆: email,weixin,',
                                      `notification_setting` json   COMMENT
'{"wechat": [12, 13],
"email": {"sender_info": {"sender_name": "xxx Information Limited", "sender_address": "<no-reply@xxx.com>", "sender_duplicate": false}, "email_option": ["topic"], "show_keyword": true, "attachment_format": "word", "email_picture_option": "", "email_template_topic": "xxxOne News", "email_template_content": "", "recipient_address_list": ["janetleung@xxx.com"]}',
                                      `created_uid` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'uid',
                                      `updated_uid` varchar(255) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT 'uid',
                                      `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '记录是否被删除,默认 0',
                                      `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建的时间 (自动生成)',
                                      `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录的更新时间 (自动生成)',
                                      PRIMARY KEY (`xxx_setup_id`),
                                      KEY `IDX_name` (`name`),
                                      KEY `IDX_customer_id` (`customer_id`),
                                      KEY `IDX_created_uid` (`created_uid`),
                                      KEY `IDX_status` (`status`),
                                      KEY `IDX_deleted` (`deleted`),
                                      KEY `IDX_created_time` (`created_time`)

) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='报告设置表';


CREATE TABLE `xxx_setup_layout` (
                                      `id` bigint(20) unsigned NOT NULL  AUTO_INCREMENT COMMENT '自增 ID',
                                      `xxx_setup_id` bigint(20) unsigned NOT NULL  DEFAULT 0 COMMENT 'xxx表主键',
                                      `xxx_setup_version` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '版本',
                                      `analysis_type` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'media_type_id:媒体类型,media_flag:媒体范围,sentiment_type_id:情感类型,link_type_id:文章类型,clustering:文章去重,precise:精准度,chart_type:图表类型,engagement_num:互动高于,subject_frequency:主体词频',
                                      `chart_type` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '图表类型',
                                      `name` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '配置名称',
                                      `folder_id` varchar(1024) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '数据源 多筆: 10001,10002,',
                                      `filter` json   COMMENT
'{media_type_id:媒体类型,media_flag:媒体范围,sentiment_type_id:情感类型,link_type_id:文章类型,clustering:文章去重,engagement_cnt:互动高于,subject_precise:精准度,subject_term_frequency:主体词频,adv_flag:广告,
sort: TOP文章列表排序, size: TOP文章數量, document_field_ids: TOP文章列表字段選擇',
                                      `sort` tinyint(1) NOT NULL DEFAULT '0' COMMENT '顺序 從 1 開始',
                                      `customer_id` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'gid',
                                      `created_uid` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'uid',
                                      `updated_uid` varchar(255) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT 'uid',
                                      `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建的时间 (自动生成)',
                                      `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录的更新时间 (自动生成)',
                                      PRIMARY KEY (`id`),
                                      KEY `IDX_xxx_setup_id` (`xxx_setup_id`),
                                      KEY `IDX_deleted` (`deleted`),
                                      KEY `IDX_name` (`name`),
                                      KEY `IDX_analysis_type` (`analysis_type`),
                                      KEY `IDX_created_time` (`created_time`)

) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ;


CREATE TABLE `xxx_result` (
                                      `xxx_result_id` bigint(20) unsigned NOT NULL  AUTO_INCREMENT COMMENT '自增 ID',
                                      `xxx_setup_id` bigint(20) unsigned NOT NULL  DEFAULT 0 COMMENT 'xxx表主键',
                                      `xxx_setup_version` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '版本 xxx setup version',
                                      `xxx_generation_date` char(50) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT '数据开始时间: 2021-09-29 21:00:00',
                                      `status` varchar(50) COLLATE utf8mb4_bin  NOT NULL DEFAULT 'pending' COMMENT '报告数据生成阶段: pending starting processing generating done failing',
                                      `error_info` varchar(500) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '失败信息，status != failing 时该字段为空',
                                      `beluga_task_id` bigint(20) unsigned NOT NULL  DEFAULT 0 COMMENT 'beluga任务ID',
                                      `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '记录是否被删除,默认 0',
                                      `created_uid` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'uid',
                                      `updated_uid` varchar(255) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT 'uid',
                                      `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建的时间 (自动生成)',
                                      `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录的更新时间 (自动生成)',
                                      PRIMARY KEY (`xxx_result_id`),
                                      KEY `IDX_xxx_result_id` (`xxx_setup_id`),
                                      KEY `IDX_xxx_generation_date` (`xxx_generation_date`),
                                      KEY `IDX_deleted` (`deleted`),
                                      KEY `IDX_created_time` (`created_time`)

) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='报告结果表';


CREATE TABLE `xxx_result_notification` (
                                      `id` bigint(20) unsigned NOT NULL  AUTO_INCREMENT COMMENT '自增 ID',
                                      `xxx_result_id` bigint(20) unsigned NOT NULL  DEFAULT 0 COMMENT 'xxx_result表主键',
                                      `status` varchar(50) COLLATE utf8mb4_bin  NOT NULL DEFAULT 'pending' COMMENT '报告发送阶段 pending starting done',
                                      `attachment_status` varchar(50) COLLATE utf8mb4_bin  NOT NULL DEFAULT 'pending' COMMENT '附件准备阶段 pending done',
                                      `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '记录是否被删除,默认 0',
                                      `created_uid` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'uid',
                                      `updated_uid` varchar(255) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT 'uid',
                                      `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建的时间 (自动生成)',
                                      `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录的更新时间 (自动生成)',
                                      PRIMARY KEY (`id`),
                                      KEY `IDX_xxx_result_id` (`xxx_result_id`),
                                      KEY `IDX_deleted` (`deleted`),
                                      KEY `IDX_created_time` (`created_time`)

) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='报告派发结果表';


CREATE TABLE `xxx_result_detail` (
                                      `xxx_result_detail_id` bigint(20) unsigned NOT NULL  AUTO_INCREMENT COMMENT '自增 ID',
                                      `xxx_result_id` bigint(20) unsigned NOT NULL  DEFAULT 0 COMMENT 'xxx_result表主键',
                                      `xxx_setup_layout_id` bigint(20) unsigned NOT NULL  DEFAULT 0 COMMENT 'xxx_setup_layout表主键',
                                      `result` json  COMMENT '报告结果',
                                      `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '记录是否被删除,默认 0',
                                      `created_uid` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'uid',
                                      `updated_uid` varchar(255) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT 'uid',
                                      `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建的时间 (自动生成)',
                                      `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录的更新时间 (自动生成)',
                                      PRIMARY KEY (`xxx_result_detail_id`),
                                      KEY `IDX_xxx_result_id` (`xxx_result_id`),
                                      KEY `IDX_xxx_setup_layout_id` (`xxx_setup_layout_id`),
                                      KEY `IDX_deleted` (`deleted`),
                                      KEY `IDX_created_time` (`created_time`)

) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='报告结果表detail';
