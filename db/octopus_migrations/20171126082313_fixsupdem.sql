
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE IF EXISTS sup_src;
DROP TABLE IF EXISTS sup_dem_src;
DROP TABLE IF EXISTS demand_report;
DROP TABLE IF EXISTS supplier_report;
DROP TABLE IF EXISTS exchange_report;
CREATE TABLE `sup_src` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `supplier` varchar(25) DEFAULT NULL,
  `source` varchar(50) DEFAULT NULL,
  `time_id` int(11) DEFAULT NULL,
  `request_in` int(11) NOT NULL DEFAULT '0',
  `ad_in` int(11) DEFAULT NULL COMMENT 'total imp comes into exchange',
  `bid_out` float NOT NULL DEFAULT '0',
  `ad_out` int(11) DEFAULT NULL,
  `ad_deliver` int(11) DEFAULT NULL COMMENT 'total show count of winner request',
  `bid_deliver` float NOT NULL,
  `profit` float NOT NULL,
  `click` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sup_src_index_gp` (`time_id`,`supplier`,`source`),
  CONSTRAINT `sup_src_time_table_id_fk` FOREIGN KEY (`time_id`) REFERENCES `time_table` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4;
CREATE TABLE `sup_dem_src` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `demand` varchar(25) DEFAULT NULL,
  `supplier` varchar(25) DEFAULT NULL,
  `source` varchar(50) DEFAULT NULL,
  `time_id` int(11) DEFAULT NULL,
  `bid_in` float DEFAULT NULL,
  `request_out` int(11) DEFAULT NULL,
  `ad_out` int(11) DEFAULT NULL,
  `bid_win` float DEFAULT NULL,
  `ad_win` int(11) DEFAULT NULL,
  `ad_in` int(11) DEFAULT NULL,
  `ad_deliver` int(11) DEFAULT NULL,
  `bid_deliver` float DEFAULT NULL,
  `profit` float NOT NULL,
  `click` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sup_dem_src_index_name` (`time_id`,`demand`,`supplier`,`source`),
  CONSTRAINT `sup_dem_src_time_table_id_fk` FOREIGN KEY (`time_id`) REFERENCES `time_table` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=963 DEFAULT CHARSET=utf8mb4;
CREATE TABLE `demand_report` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `demand` varchar(25) NOT NULL,
  `target_date` date NOT NULL,
  `ad_out` int(11) NOT NULL DEFAULT '0' COMMENT 'total ad send to demand from exchange',
  `ad_in` int(11) DEFAULT '0' COMMENT 'total ad comes from demand',
  `ad_win` int(11) NOT NULL DEFAULT '0' COMMENT 'total win ad from demand',
  `ad_deliver` int(11) DEFAULT '0',
  `bid_deliver` float NOT NULL DEFAULT '0',
  `profit` float NOT NULL DEFAULT '0',
  `click` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `demand_report_demand_time` (`demand`,`target_date`),
  KEY `idx_demand_report_target_date` (`target_date`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;
CREATE TABLE `supplier_report` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `supplier` varchar(25) NOT NULL,
  `target_date` date NOT NULL,
  `ad_in` int(11) NOT NULL DEFAULT '0',
  `ad_out` int(11) NOT NULL DEFAULT '0',
  `ad_deliver` int(11) NOT NULL DEFAULT '0',
  `earn` float NOT NULL,
  `click` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `supplier_report_unique_gp` (`supplier`,`target_date`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;
CREATE TABLE `exchange_report` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `target_date` date NOT NULL,
  `supplier_ad_in` int(11) NOT NULL DEFAULT '0',
  `supplier_ad_out` int(11) NOT NULL DEFAULT '0',
  `demand_ad_in` int(11) NOT NULL DEFAULT '0',
  `demand_ad_out` int(11) NOT NULL DEFAULT '0',
  `earn` float NOT NULL DEFAULT '0',
  `spent` float NOT NULL DEFAULT '0',
  `income` float NOT NULL DEFAULT '0',
  `click` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `exchange_date_uindex` (`target_date`)
) ENGINE=InnoDB AUTO_INCREMENT=113 DEFAULT CHARSET=utf8mb4;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS sup_src;
DROP TABLE IF EXISTS sup_dem_src;
DROP TABLE IF EXISTS demand_report;
DROP TABLE IF EXISTS supplier_report;
DROP TABLE IF EXISTS exchange_report;
