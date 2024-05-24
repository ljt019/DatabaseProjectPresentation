CREATE TABLE IF NOT EXISTS Issue (
    IssueID VARCHAR(50) PRIMARY KEY,
    SubmitterName VARCHAR(100),
    BuildingLocation VARCHAR(100),
    RepairDescription TEXT,
    PriorityLevel VARCHAR(50),
    SubmitDate DATE,
    Deadline DATE,
    CompletedDate DATE,
    CompletedStatus VARCHAR(50),
    EXHIBIT_Name VARCHAR(100),
    EMPLOYEE_AccountID VARCHAR(50),
    JOTFORM_SubmissionID VARCHAR(50),
    FOREIGN KEY (EXHIBIT_Name) REFERENCES Exhibit(Name),
    FOREIGN KEY (EMPLOYEE_AccountID) REFERENCES Employee(AccountID),
    FOREIGN KEY (JOTFORM_SubmissionID) REFERENCES JotForm(SubmissionID)
);