DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_catalog.pg_database WHERE datname = '${DB_NAME}') THEN
        CREATE DATABASE '${DB_NAME}'
            WITH OWNER = '${DB_USER}'
            ENCODING = 'UTF8'
            LC_COLLATE = 'en_US.UTF-8'
            LC_CTYPE = 'en_US.UTF-8'
            TABLESPACE = pg_default
            CONNECTION LIMIT = -1;
    END IF;
END $$;