CREATE TABLE IF NOT EXISTS Sponsorship (
    SponsorName VARCHAR(100),
    StartDate DATE,
    EndDate DATE,
    EXHIBIT_Name VARCHAR(100),
    FOREIGN KEY (EXHIBIT_Name) REFERENCES Exhibit(Name),
    PRIMARY KEY (SponsorName, EXHIBIT_Name)
);