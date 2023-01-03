WITH `ids` AS (SELECT *, rank() OVER (PARTITION BY `email` ORDER BY `id`) AS `rank` FROM `person`)
DELETE `left`
FROM `person` AS `left`
         RIGHT JOIN `ids` AS `right` ON `left`.`id` = `right`.`id`
WHERE `right`.`rank` > 1;
