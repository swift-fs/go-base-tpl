-- hotgo自动生成菜单权限SQL 通常情况下只在首次生成代码时自动执行一次
-- 如需再次执行请先手动删除生成的菜单权限和SQL文件：C:\Users\Administrator\Desktop\gosrc\hotgo_dev\server\storage\data\generate\test_category_menu.sql
-- Version: 2.13.1
-- Date: 2024-04-08 23:26:32
-- Link https://github.com/bufanyun/hotgo

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;

--
-- 数据库： `hotgo`
--

-- --------------------------------------------------------

--
-- 插入表中的数据 `gc_admin_menu`
--


SET @now := now();


-- 菜单目录
INSERT INTO `gc_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, '2366', '测试分类', 'testCategory', '/testCategory', '', '1', '/develop/generateDemo/testCategory/index', '', '', 'ParentLayout', '1', '', '0', '0', '', '0', '0', '0', '3', 'tr_2097 tr_2366 ', '1000', '', '1', @now, @now);


SET @dirId = LAST_INSERT_ID();


-- 菜单页面
-- 列表
INSERT INTO `gc_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @dirId, '测试分类列表', 'testCategoryIndex', 'index', '', '2', '', '/testCategory/list', '', '/testCategory/index', '1', 'testCategory', '0', '0', '', '0', '1', '0', '4', CONCAT('tr_2097 tr_2366 tr_', @dirId,' '), '10', '', '1', @now, @now);


SET @listId = LAST_INSERT_ID();

-- 详情
INSERT INTO `gc_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '测试分类详情', 'testCategoryView', '', '', '3', '', '/testCategory/view', '', '', '1', '', '0', '0', '', '0', '1', '0', '5', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId,' '), '10', '', '1', @now, @now);


-- 菜单按钮

-- 编辑
INSERT INTO `gc_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '编辑/新增测试分类', 'testCategoryEdit', '', '', '3', '', '/testCategory/edit', '', '', '1', '', '0', '0', '', '0', '1', '0', '5', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId,' '), '20', '', '1', @now, @now);


SET @editId = LAST_INSERT_ID();

-- 获取最大排序
INSERT INTO `gc_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @editId, '获取测试分类最大排序', 'testCategoryMaxSort', '', '', '3', '', '/testCategory/maxSort', '', '', '1', '', '0', '0', '', '0', '0', '0', '6', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId, ' tr_', @editId,' '), '30', '', '1', @now, @now);


-- 删除
INSERT INTO `gc_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '删除测试分类', 'testCategoryDelete', '', '', '3', '', '/testCategory/delete', '', '', '1', '', '0', '0', '', '0', '0', '0', '5', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId,' '), '40', '', '1', @now, @now);





COMMIT;