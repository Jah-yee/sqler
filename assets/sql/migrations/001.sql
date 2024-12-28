CREATE TABLE IF NOT EXISTS migrations (
    id SERIAL PRIMARY KEY,
    filename TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    migrated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL, -- bcrypt
    is_sys_admin BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS user_access_tokens (
    id BIGSERIAL PRIMARY KEY,
    value TEXT NOT NULL UNIQUE,
    user_id BIGINT NOT NULL,
    expires_at TIMESTAMP,
    user_agent TEXT NOT NULL,
    ip_address INET NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS data_sources (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    driver TEXT NOT NULL,
    connection_string TEXT NOT NULL,
    creator_user_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS projects (
    id BIGSERIAL PRIMARY KEY,
    parent_project_id BIGINT,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS project_data_sources (
    project_id BIGINT NOT NULL,
    data_source_id BIGINT NOT NULL,

    PRIMARY KEY (project_id, data_source_id)
);

CREATE TABLE IF NOT EXISTS project_members (
    project_id BIGINT NOT NULL,
    member_user_id BIGINT NOT NULL,
    role TEXT NOT NULL, -- admin, maintainer, contributor, viewer

    PRIMARY KEY (project_id, member_user_id)
);

CREATE TABLE IF NOT EXISTS queries (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS query_revisions (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    query_id BIGINT NOT NULL,
    parent_query_revision_id BIGINT,
    data_source_id BIGINT NOT NULL,
    code TEXT NOT NULL,
    note TEXT, -- e.g., reason for creating this revision
    creator_user_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS query_revision_promotions (
    id BIGSERIAL PRIMARY KEY,
    query_id BIGINT NOT NULL,
    query_revision_id BIGINT NOT NULL,
    creator_user_id BIGINT NOT NULL,
    note TEXT, -- e.g., reason for promotion
    created_at TIMESTAMP NOT NULL
);
