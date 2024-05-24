CREATE TABLE IF NOT EXISTS Exhibit (
    Name VARCHAR(100) PRIMARY KEY,
    Description TEXT,
    BuildingLocation VARCHAR(100),
    ExhibitImage VARCHAR(255),
    DocumentationLink VARCHAR(255),
    CurrentStatus VARCHAR(50)
);
