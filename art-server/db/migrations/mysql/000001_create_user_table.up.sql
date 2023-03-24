CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `realname` varchar(255) CHARACTER SET utf8 NOT NULL,
  `username` varchar(255) CHARACTER SET utf8 NOT NULL,
  `avatar` varchar(255) CHARACTER SET utf8 NOT NULL,
  `company` varchar(255) CHARACTER SET utf8 COLLATE utf8_persian_ci DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `bio` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `img` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `join_date` datetime DEFAULT NULL,
  `verify_type` varchar(255) DEFAULT NULL,
  `verify_name` varchar(255) DEFAULT NULL,
  `is_creator` tinyint(4) DEFAULT '0',
  `follower_num` int(11) DEFAULT '0',
  `following_num` int(11) DEFAULT '0',
  `links` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;