DELIMITER $$

CREATE TRIGGER BeforeInsertIssueStatus
BEFORE INSERT ON Issue
FOR EACH ROW
BEGIN
    IF NEW.CompletedStatus IS NULL THEN
        SET NEW.CompletedStatus = 'Pending';
    END IF;
END$$

DELIMITER ;
