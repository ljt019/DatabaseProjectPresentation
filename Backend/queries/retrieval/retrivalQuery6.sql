SELECT e.FirstName, e.LastName
FROM Employee e
WHERE e.AccountID IN (
    SELECT i.EMPLOYEE_AccountID
    FROM Issue i
    JOIN Exhibit ex ON i.EXHIBIT_Name = ex.Name
    WHERE ex.CurrentStatus = 'Needs Repair'
);