CREATE TABLE `payment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `item_id` int(11) NOT NULL,
  `onsale_type` tinyint(4) DEFAULT NULL,
  `price` decimal(10,2) NOT NULL,
  `service_fee` decimal(10,2) DEFAULT NULL,
  `creator` varchar(255) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL,
  `pay_status` tinyint(4) DEFAULT NULL,
  `pay_user` varchar(255) DEFAULT NULL,
  `pay_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
