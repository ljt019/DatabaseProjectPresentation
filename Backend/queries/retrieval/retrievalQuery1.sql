SELECT i.IssueID, i.RepairDescription, i.PriorityLevel, e.FirstName, e.LastName
FROM Issue i
JOIN Employee e ON i.EMPLOYEE_AccountID = e.AccountID
WHERE i.CompletedStatus = 'Pending';
