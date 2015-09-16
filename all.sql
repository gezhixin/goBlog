DROP TABLE IF EXISTS `t_user`;
CREATE TABLE IF NOT EXISTS `tbl_tag` (
        `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `name` varchar(20) NOT NULL DEFAULT '' ,
        `count` bigint NOT NULL DEFAULT '0'
    ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
CREATE INDEX `tbl_tag_name` ON `tbl_tag` (`name`);