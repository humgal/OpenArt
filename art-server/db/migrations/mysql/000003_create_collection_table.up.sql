CREATE TABLE `collection` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `img` varchar(255) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL,
  `creator` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`,`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;