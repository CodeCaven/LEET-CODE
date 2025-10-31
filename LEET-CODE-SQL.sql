-- 175: Combine 2 Tables
-- Use a LEFT JOIN for nulls
SELECT p.firstName, p.lastName, a.city, a.state
FROM Person p
LEFT JOIN Address a
ON p.personId = a.personId;

-- 176: Second Highest Salary
SELECT max(salary) as SecondHighestSalary
FROM employee
WHERE employee.salary < (
SELECT max(salary)
FROM employee
)

-- 178: Rank Scores
select s1.Score, count(distinct s2.Score) as "Rank" from Scores s1, Scores s2
where s2.Score >= s1.Score
group by s1.Id
order by Score DESC;

-- 584: Customer Referee
select name
from Customer
where referee_id != 2 or referee_id is null;

-- 182: Duplicate Emails
SELECT email FROM Person
GROUP BY email HAVING COUNT(email) > 1;

--184: Department Highest Salary
SELECT d.name as "Department", e.name as "Employee", e.salary as "Salary"
FROM Employee e, Department d
WHERE e.departmentId = d.id
AND
(e.DepartmentId, e.Salary) in
(SELECT e.DepartmentId, max(e.Salary) FROM Employee e GROUP BY e.DepartmentId);

-- 180: Consecutive Numbers
-- Example uses Common Table Expression and Window Functions
With cte_num AS (Select Num AS V1, LEAD(num,1) Over (order by id) AS V2, LEAD(num,2) Over (order by id) AS V3 From Logs)
Select V1 AS ConsecutiveNums from cte_num
where V1=V2
And V2=V3;