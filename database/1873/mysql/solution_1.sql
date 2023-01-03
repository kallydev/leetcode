SELECT `employee_id`, if(`employee_id` % 2 = 1 AND left(`name`, 1) != 'M', `salary`, 0) AS `bonus`
FROM `employees`
ORDER BY `employee_id`;
