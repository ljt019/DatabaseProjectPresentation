DELIMITER $$

CREATE TRIGGER AfterInsertIssue
AFTER INSERT ON Issue
FOR EACH ROW
BEGIN
    UPDATE Exhibit
    SET CurrentStatus = 'Needs Repair'
    WHERE Name = NEW.EXHIBIT_Name;
END$$

DELIMITER ;