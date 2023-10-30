ALTER TABLE `register` ADD UNIQUE `uidx_account_semester_class_course`
  (`account_id`, `semester_id`, `class_id`, `course_id`);
