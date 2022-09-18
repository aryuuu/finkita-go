BEGIN;
    CREATE TABLE IF NOT EXISTS mutations (
        id UUID NOT NULL DEFAULT UUID_GENERATE_V4(),
        account_id UUID NOT NULL,
        email CHARACTER VARYING(100) NOT NULL,
        date TIMESTAMP WITHOUT TIME ZONE NOT NULL,
        description TEXT,
        type CHARACTER VARYING(10) NOT NULL,
        amount INTEGER NOT NULL,
        balance INTEGER NOT NULL,
        currency INTEGER NOT NULL,
        created_at TIMESTAMP WITHOUT TIME ZONE,
        updated_at TIMESTAMP WITHOUT TIME ZONE,
        deleted_at TIMESTAMP WITHOUT TIME ZONE,
        CONSTRAINT mutations_pkey PRIMARY KEY(id),
        CONSTRAINT fk_mutations_accounts FOREIGN KEY(account_id) REFERENCES accounts(id) ON UPDATE CASCADE ON DELETE CASCADE
    );

    CREATE INDEX IF NOT EXISTS idx_mutations_deleted_at ON accounts USING btree (deleted_at ASC NULLS LAST);
    CREATE INDEX IF NOT EXISTS idx_mutations_created_at ON accounts USING btree (created_at ASC NULLS LAST);

    CREATE TABLE IF NOT EXISTS stable_mutations (
        id UUID NOT NULL DEFAULT UUID_GENERATE_V4(),
        account_id UUID NOT NULL,
        email CHARACTER VARYING(100) NOT NULL,
        date TIMESTAMP WITHOUT TIME ZONE NOT NULL,
        description TEXT,
        type CHARACTER VARYING(10) NOT NULL,
        amount INTEGER NOT NULL,
        balance INTEGER NOT NULL,
        currency INTEGER NOT NULL,
        created_at TIMESTAMP WITHOUT TIME ZONE,
        updated_at TIMESTAMP WITHOUT TIME ZONE,
        deleted_at TIMESTAMP WITHOUT TIME ZONE,
        CONSTRAINT stable_mutations_pkey PRIMARY KEY(id),
        CONSTRAINT fk_stable_mutations_accounts FOREIGN KEY(account_id) REFERENCES accounts(id) ON UPDATE CASCADE ON DELETE CASCADE
    );

    CREATE INDEX IF NOT EXISTS idx_stable_mutations_deleted_at ON accounts USING btree (deleted_at ASC NULLS LAST);
    CREATE INDEX IF NOT EXISTS idx_stable_mutations_created_at ON accounts USING btree (created_at ASC NULLS LAST);

COMMIT;
