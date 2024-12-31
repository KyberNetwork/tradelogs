ALTER TABLE txorigin ADD COLUMN type TEXT;
UPDATE txorigin SET type = 'cow';