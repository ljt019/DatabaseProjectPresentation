DROP VIEW IF EXISTS OperationalNeedsRepair;
CREATE VIEW OperationalNeedsRepair AS
SELECT ex.Name AS ExhibitName, ex.CurrentStatus
FROM Exhibit ex
WHERE ex.CurrentStatus IN ('Operational', 'Needs Repair');