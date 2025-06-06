-- 176.中 第二高的薪水

-- Employee 表：
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | id          | int  |
-- | salary      | int  |
-- +-------------+------+
-- id 是这个表的主键。
-- 表的每一行包含员工的工资信息。
 

-- 查询并返回 Employee 表中第二高的 不同 薪水 。如果不存在第二高的薪水，查询应该返回 null(Pandas 则返回 None) 。

-- 查询结果如下例所示。

 

-- 示例 1：

-- 输入：
-- Employee 表：
-- +----+--------+
-- | id | salary |
-- +----+--------+
-- | 1  | 100    |
-- | 2  | 200    |
-- | 3  | 300    |
-- +----+--------+
-- 输出：
-- +---------------------+
-- | SecondHighestSalary |
-- +---------------------+
-- | 200                 |
-- +---------------------+
-- 示例 2：

-- 输入：
-- Employee 表：
-- +----+--------+
-- | id | salary |
-- +----+--------+
-- | 1  | 100    |
-- +----+--------+
-- 输出：
-- +---------------------+
-- | SecondHighestSalary |
-- +---------------------+
-- | null                |
-- +---------------------+

-- 方法1:
select (select distinct salary from Employee order by salary desc limit 1,1) as SecondHighestSalary; -- distinct避免仅存在两条相同salary的场景

-- 方法2:
select ifnull((select distinct salary from Employee order by salary desc limit 1,1),null) as SecondHighestSalary

-- 方法3:
SELECT MAX(salary) AS SecondHighestSalary
FROM Employee
WHERE salary < (
    SELECT MAX(salary)
    FROM Employee
);