-- name: CheckRegistrationCode :one
SELECT code FROM registration_codes WHERE code = $1 AND expires_at > CURRENT_TIMESTAMP LIMIT 1;

-- name: CheckEmail :one
SELECT id FROM users WHERE email = $1;

-- name: CreateUser :exec
INSERT INTO users (id, email, password, display_name) VALUES ($1, $2, $3, $4);

-- name: SetDisplayName :exec
UPDATE users SET display_name = $2, slug = $3 WHERE id = $1;

-- name: CreateRole :exec
INSERT INTO roles (name, description) VALUES ($1, $2);

-- name: CreateUserRole :exec
INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2);

-- name: CreateRegistrationCode :exec
INSERT INTO registration_codes (code, expires_at) VALUES ($1, $2);
