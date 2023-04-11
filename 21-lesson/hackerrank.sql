-- 11
SELECT DISTINCT t.name
FROM employee t
WHERE t.name NOT LIKE 'A%' 
AND t.name NOT LIKE 'E%'
AND t.name NOT LIKE 'I%' 
AND t.name NOT LIKE 'O%' 
AND t.name NOT LIKE 'U%'
AND t.name NOT LIKE '%a' 
AND t.name NOT LIKE '%e' 
AND t.name NOT LIKE '%i' 
AND t.name NOT LIKE '%o' 
AND t.name NOT LIKE '%u';



--
SELECT NAME FROM STUDENTS WHERE Marks > 75 ORDER BY RIGHT(NAME, 3), ID ASC;
--select t.name from students t where t.marks>75 and t.name not LIKE '%bby' and t.name not LIKE '%ina' order by id asc;