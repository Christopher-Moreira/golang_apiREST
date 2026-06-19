CREATE TABLE openings (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    role TEXT,
    company TEXT,
    location TEXT,
    remote BOOLEAN,
    link TEXT,
    salary BIGINT
);

CREATE INDEX idx_openings_deleted_at ON openings (deleted_at);
