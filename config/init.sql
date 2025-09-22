CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

drop table if EXISTS notes;

CREATE TABLE notes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title TEXT NOT NULL,
    content TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_done BOOLEAN NOT NULL DEFAULT FALSE
);

select * from notes;