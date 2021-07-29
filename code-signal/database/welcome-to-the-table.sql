-- 1: projectList https://app.codesignal.com/arcade/db/welcome-to-the-table/RXErLMFkXFkM4MpYY
CREATE PROCEDURE projectList()
BEGIN
select project_name,team_lead,income FROM Projects order by internal_id ASC limit 10;
END