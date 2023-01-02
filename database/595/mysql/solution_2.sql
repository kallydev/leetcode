SELECT `name`, `area`, `population`
FROM `world`
WHERE `area` >= 3000000
UNION
SELECT `name`, `area`, `population`
FROM `world`
WHERE `population` >= 25000000;
