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