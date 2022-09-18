BEGIN;
    CREATE TABLE IF NOT EXISTS accounts (
        id UUID NOT NULL DEFAULT UUID_GENERATE_V4(),
        email CHARACTER VARYING(100) NOT NULL,
        bank CHARACTER VARYING(100) NOT NULL,
        account_number CHARACTER VARYING(100) NOT NULL,
        password TEXT NOT NULL,
        created_at TIMESTAMP WITHOUT TIME ZONE,
        updated_at TIMESTAMP WITHOUT TIME ZONE,
        deleted_at TIMESTAMP WITHOUT TIME ZONE,
        CONSTRAINT accounts_pkey PRIMARY KEY(id),
        CONSTRAINT accounts_unique_bank_account_number UNIQUE (bank, account_number, deleted_at)
    );

    CREATE INDEX IF NOT EXISTS idx_accounts_deleted_at ON accounts USING btree (deleted_at ASC NULLS LAST);

COMMIT;
