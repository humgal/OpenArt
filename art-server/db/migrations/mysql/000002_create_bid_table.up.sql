CREATE TABLE `bid` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `item_id` int(11) NOT NULL,
  `service_fee` decimal(10,2) DEFAULT NULL,
  `total` decimal(10,2) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `bid_date` datetime DEFAULT NULL,
  `item_price_type` tinyint(4) DEFAULT NULL,
  `onsale_type` tinyint(4) DEFAULT NULL,
  `bid_end_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;