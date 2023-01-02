SELECT `name` AS `customers`
FROM `customers`
         LEFT JOIN `orders` ON `customers`.`id` = `orders`.`customerId`
WHERE `orders`.`id` IS NULL;