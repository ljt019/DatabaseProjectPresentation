SELECT ex.Name AS ExhibitName, n.Note, n.Date
FROM Exhibit ex
JOIN Note n ON ex.Name = n.EXHIBIT_Name
WHERE ex.Name = 'Moon Chair';