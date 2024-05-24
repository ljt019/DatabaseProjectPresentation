DELIMITER $$

CREATE TRIGGER BeforeInsertJotForm
BEFORE INSERT ON JotForm
FOR EACH ROW
BEGIN
    DECLARE count INT;
    SELECT COUNT(*) INTO count FROM JotForm WHERE SubmissionID = NEW.SubmissionID;
    IF count > 0 THEN
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'Duplicate submission ID';
    END IF;
END$$

DELIMITER ;