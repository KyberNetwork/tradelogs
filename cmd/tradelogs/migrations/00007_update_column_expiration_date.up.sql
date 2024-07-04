UPDATE tradelogs SET expiration_date = 0 WHERE expiration_date IS NULL;
ALTER TABLE tradelogs ALTER COLUMN expiration_date SET DEFAULT 0;
ALTER TABLE tradelogs ALTER COLUMN expiration_date SET NOT NULL;
