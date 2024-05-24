CREATE TABLE IF NOT EXISTS Part (
    Name VARCHAR(100) PRIMARY KEY,
    Link VARCHAR(255),
    NumInStock INT,
    PurchaseOrderCreated BOOLEAN,
    PendingOrder BOOLEAN,
    ExpectedDeliveyDate DATE,
    EXHIBIT_Name VARCHAR(100),
    FOREIGN KEY (EXHIBIT_Name) REFERENCES Exhibit(Name)
);