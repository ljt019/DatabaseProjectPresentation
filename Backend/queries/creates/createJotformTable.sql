CREATE TABLE IF NOT EXISTS JotForm (
    SubmissionID VARCHAR(50) PRIMARY KEY,
    SubmitterName VARCHAR(100),
    BuildingLocation VARCHAR(100),
    ExhibitName VARCHAR(100),
    RepairDescription TEXT,
    SuggestedPriorityLevel VARCHAR(50),
    SubmitDate DATE,
    FOREIGN KEY (ExhibitName) REFERENCES Exhibit(Name)
);