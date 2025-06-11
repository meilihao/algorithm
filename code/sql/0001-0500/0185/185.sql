-- 185.困 部门工资前三高的所有员工

-- 表: Employee

-- +--------------+---------+
-- | Column Name  | Type    |
-- +--------------+---------+
-- | id           | int     |
-- | name         | varchar |
-- | salary       | int     |
-- | departmentId | int     |
-- +--------------+---------+
-- id 是该表的主键列(具有唯一值的列)。
-- departmentId 是 Department 表中 ID 的外键（reference 列）。
-- 该表的每一行都表示员工的ID、姓名和工资。它还包含了他们部门的ID。
 

-- 表: Department

-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | name        | varchar |
-- +-------------+---------+
-- id 是该表的主键列(具有唯一值的列)。
-- 该表的每一行表示部门ID和部门名。
 

-- 公司的主管们感兴趣的是公司每个部门中谁赚的钱最多。一个部门的 高收入者 是指一个员工的工资在该部门的 不同 工资中 排名前三 。

-- 编写解决方案，找出每个部门中 收入高的员工 。

-- 以 任意顺序 返回结果表。

-- 返回结果格式如下所示。

 

-- 示例 1:

-- 输入: 
-- Employee 表:
-- +----+-------+--------+--------------+
-- | id | name  | salary | departmentId |
-- +----+-------+--------+--------------+
-- | 1  | Joe   | 85000  | 1            |
-- | 2  | Henry | 80000  | 2            |
-- | 3  | Sam   | 60000  | 2            |
-- | 4  | Max   | 90000  | 1            |
-- | 5  | Janet | 69000  | 1            |
-- | 6  | Randy | 85000  | 1            |
-- | 7  | Will  | 70000  | 1            |
-- +----+-------+--------+--------------+
-- Department  表:
-- +----+-------+
-- | id | name  |
-- +----+-------+
-- | 1  | IT    |
-- | 2  | Sales |
-- +----+-------+
-- 输出: 
-- +------------+----------+--------+
-- | Department | Employee | Salary |
-- +------------+----------+--------+
-- | IT         | Max      | 90000  |
-- | IT         | Joe      | 85000  |
-- | IT         | Randy    | 85000  |
-- | IT         | Will     | 70000  |
-- | Sales      | Henry    | 80000  |
-- | Sales      | Sam      | 60000  |
-- +------------+----------+--------+
-- 解释:
-- 在IT部门:
-- - Max的工资最高
-- - 兰迪和乔都赚取第二高的独特的薪水
-- - 威尔的薪水是第三高的

-- 在销售部:
-- - 亨利的工资最高
-- - 山姆的薪水第二高
-- - 没有第三高的工资，因为只有两名员工

-- 不使用窗口函数:
SELECT
    d.Name AS 'Department', e1.Name AS 'Employee', e1.Salary
FROM
    Employee e1
        JOIN
    Department d ON e1.DepartmentId = d.Id
WHERE
    3 > (SELECT
            COUNT(DISTINCT e2.Salary)
        FROM
            Employee e2
        WHERE
            e2.Salary > e1.Salary
                AND e1.DepartmentId = e2.DepartmentId
        )
;
-- 使用了一个相关子查询来确定每个员工在其部门中的薪资排名, 对于主查询中的每一个员工(e1)，子查询计算在同一个部门中薪资高于该员工的不同薪资值的数量:
-- 排名逻辑: 条件 3 > COUNT(...) 确保只选择排名前三的员工
-- - 如果子查询返回0，表示没有比该员工薪资更高的员工 → 该员工是部门最高薪
-- - 如果子查询返回1，表示有1个更高的薪资 → 该员工是部门第二高薪
-- - 如果子查询返回2，表示有2个更高的薪资 → 该员工是部门第三高薪
-- 使用 DISTINCT e2.Salary 确保计算的是不同的薪资值，而不是员工人数

-- 使用窗口函数:
select Department, Employee, Salary
from (
    select d.name Department, ee.name Employee, ee.salary Salary, dense_rank() over(partition by departmentId order by salary desc) ranks
    from Employee ee
    left join Department d
    on ee.departmentId = d.id
) t
where ranks <= 3