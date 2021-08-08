-- Link: https://leetcode.com/problems/nth-highest-salary/

CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  SET N=N-1;
  RETURN (
      # Write your MySQL query statement below.
      select e.Salary
      from Employee e
      group by e.Salary
      order by e.Salary desc
      limit 1 offset N
  );
END