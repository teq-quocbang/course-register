CREATE TABLE IF NOT EXISTS `class` (
  `id` VARCHAR(50),
  `start_time` TIMESTAMP,
  `end_time` TIMESTAMP,
  `course_id` varchar(50),
  `semester_id` varchar(50),
  `credits` INT,
  `max_slot` INT,
  `current_slot` INT,
  `can_cancel` BOOLEAN,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT, 
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` INT,

  PRIMARY KEY (`id`),
  FOREIGN KEY (`semester_id`) REFERENCES semester(id),
  FOREIGN KEY (`course_id`) REFERENCES course(id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;