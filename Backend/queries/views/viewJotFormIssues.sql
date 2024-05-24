DROP VIEW IF EXISTS JotFormIssues;
CREATE VIEW JotFormIssues AS
SELECT j.SubmissionID, j.SubmitterName, j.ExhibitName, j.RepairDescription, i.IssueID, i.PriorityLevel, i.CompletedStatus
FROM JotForm j
JOIN Issue i ON j.SubmissionID = i.JOTFORM_SubmissionID;