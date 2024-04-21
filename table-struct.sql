
CREATE TABLE `inverter_data` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `day_energy` decimal(10,2) DEFAULT NULL,
  `fac` decimal(10,2) DEFAULT NULL,
  `iac` decimal(10,2) DEFAULT NULL,
  `idc` decimal(10,2) DEFAULT NULL,
  `pac` decimal(10,2) DEFAULT NULL,
  `total_energy` decimal(12,2) DEFAULT NULL,
  `uac` decimal(10,2) DEFAULT NULL,
  `udc` decimal(10,2) DEFAULT NULL,
  `year_energy` decimal(12,2) DEFAULT NULL,
  `datetime` timestamp NOT NULL DEFAULT current_timestamp(),
  `device_id` tinyint(4) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;