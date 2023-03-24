CREATE TABLE `item` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `tag` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `upload_url` varchar(255) NOT NULL,
  `sales_status` int(11) NOT NULL,
  `price` decimal(10,2) NOT NULL,
  `creator_id` int(11) NOT NULL,
  `creator` varchar(255) NOT NULL,
  `create_date` datetime DEFAULT NULL,
  `collection_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;