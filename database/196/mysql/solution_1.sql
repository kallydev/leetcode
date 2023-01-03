DELETE `left`
FROM `person` AS `left`
         INNER JOIN `person` AS `right`
WHERE `left`.`email` = `right`.`email`
  AND `left`.`id` > `right`.`id`;
