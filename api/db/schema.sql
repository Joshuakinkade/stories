CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(128) UNIQUE NOT NULL,
    display_name VARCHAR(128) UNIQUE NOT NULL,
    slug VARCHAR(128) UNIQUE,
    password TEXT NOT NULL,
    roles VARCHAR(32)[],
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE roles (
    id UUID PRIMARY KEY,
    name VARCHAR(32) UNIQUE NOT NULL,
    description text,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE user_roles (
    user_id UUID NOT NULL,
    role_id UUID NOT NULL,
    PRIMARY KEY (user_id, role_id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_user_roles_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_user_roles_role_id FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);

CREATE TABLE registration_codes (
    code VARCHAR(32) PRIMARY KEY,
    user_id UUID,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_registration_codes_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
