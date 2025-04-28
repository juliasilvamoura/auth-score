DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'uuid-ossp') THEN
        CREATE EXTENSION "uuid-ossp";
        RAISE NOTICE 'Extension uuid-ossp installed successfully';
    ELSE
        RAISE NOTICE 'Extension uuid-ossp already exists';
    END IF;
END $$;