BEGIN;
    ALTER TABLE accounts
    ADD COLUMN IF NOT EXISTS user_id CHARACTER VARYING(100);
COMMIT;

