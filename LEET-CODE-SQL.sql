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