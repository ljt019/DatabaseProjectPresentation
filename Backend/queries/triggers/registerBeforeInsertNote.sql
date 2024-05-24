DELIMITER $$

CREATE TRIGGER BeforeInsertNote
BEFORE INSERT ON Note
FOR EACH ROW
BEGIN
    IF NEW.NoteID IS NULL THEN
        SET NEW.NoteID = UUID();
    END IF;
END$$

DELIMITER ;