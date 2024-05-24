SELECT ex.Name AS ExhibitName, s.SponsorName, s.StartDate, s.EndDate
FROM Exhibit ex
JOIN Sponsorship s ON ex.Name = s.EXHIBIT_Name;