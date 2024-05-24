DROP VIEW IF EXISTS ExhibitDetails;
CREATE VIEW ExhibitDetails AS
SELECT 
    ex.Name AS ExhibitName, 
    ex.Description, 
    ex.BuildingLocation, 
    ex.ExhibitImage, 
    ex.DocumentationLink, 
    ex.CurrentStatus, 
    n.Note, 
    n.Date, 
    s.SponsorName, 
    s.StartDate, 
    s.EndDate
FROM Exhibit ex
JOIN Note n ON ex.Name = n.EXHIBIT_Name
JOIN Sponsorship s ON ex.Name = s.EXHIBIT_Name;