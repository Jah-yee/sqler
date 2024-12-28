-- $${{define "user.create"}}$$
INSERT INTO users (name, email, password, is_sys_admin, created_at)
VALUES (
    $${{.Bind .Args.Name}}$$,
    $${{.Bind .Args.Email}}$$,
    $${{.Bind .Args.Password}}$$,
    $${{.Bind .Args.IsSysAdmin}}$$,
    CURRENT_TIMESTAMP
);
-- $${{end}}$$

-- $${{define "user.findByEmail"}}$$
SELECT id, name, email, password, is_sys_admin, created_at FROM users WHERE email = $${{.Bind .Args}}$$ LIMIT 1;
-- $${{end}}$$


-- $${{define "user.findById"}}$$
SELECT id, name, email, password, is_sys_admin, created_at FROM users WHERE id = $${{.Bind .Args}}$$ LIMIT 1;
-- $${{end}}$$

-- $${{define "user.acquireAccessToken"}}$$
INSERT INTO user_access_tokens(value, user_id, expires_at, user_agent, ip_address, created_at, updated_at)
VALUES (
    $${{.Bind .Args.value}}$$,
    $${{.Bind .Args.user_id}}$$,
    $${{.Bind .Args.expires_at}}$$,
    $${{.Bind .Args.user_agent}}$$,
    $${{.Bind .Args.ip_address}}$$,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
-- $${{end}}$$