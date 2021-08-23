-- 1: projectList https://app.codesignal.com/arcade/db/welcome-to-the-table/RXErLMFkXFkM4MpYY
CREATE PROCEDURE projectList()
BEGIN
select project_name,team_lead,income FROM Projects order by internal_id ASC limit 10;
END

-- 2: countriesSelection https://app.codesignal.com/arcade/db/welcome-to-the-table/jLeSZGMvaEhSJnEsS
CREATE PROCEDURE countriesSelection()
BEGIN
select * from countries where continent = 'Africa';
END

-- 3: monthlyScholarships https://app.codesignal.com/arcade/db/welcome-to-the-table/6eMusSHbdjavds7Cq
CREATE PROCEDURE monthlyScholarships()
BEGIN
select id,scholarship/12 as scholarship from scholarships order by id asc;
END

-- 4: projectsTeam https://app.codesignal.com/arcade/db/welcome-to-the-table/J8JfCzFnr4cPMQgZ6
CREATE PROCEDURE projectsTeam()
BEGIN
-- select distinct name from projectLog order by name; จริงๆควรเป็นอันนี้มากกว่า
-- ใช้ distinct จะไว้หาที่ duplicate อย่าวเดียว
-- ถ้าใช้ group by เหมาะกับการทำ aggregate functions
-- ซึ่งโจทย์ถามคำตอบแค่ "ลบตัวซ้ำ" เพราะงันควรเลือก distinct
select name from projectLog group by name order by name;
END

-- 5: automaticNotifications https://app.codesignal.com/arcade/db/welcome-to-the-table/kaxWej78qgdGHy8mr/
CREATE PROCEDURE automaticNotifications()
SELECT email
FROM users
WHERE LOWER(role) NOT IN ("admin", "premium")

ORDER BY email;
