DROP VIEW IF EXISTS EmployeeIssues;
CREATE VIEW EmployeeIssues AS
SELECT e.FirstName, e.LastName, i.IssueID, i.RepairDescription, i.PriorityLevel
FROM Employee e
JOIN Issue i ON e.AccountID = i.EMPLOYEE_AccountID;