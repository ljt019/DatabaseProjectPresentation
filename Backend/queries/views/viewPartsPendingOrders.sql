DROP VIEW IF EXISTS PartsPendingOrders;
CREATE VIEW PartsPendingOrders AS
SELECT p.Name AS PartName, p.NumInStock, p.ExpectedDeliveyDate, ex.Name AS ExhibitName
FROM Part p
JOIN Exhibit ex ON p.EXHIBIT_Name = ex.Name
WHERE p.PendingOrder = TRUE;