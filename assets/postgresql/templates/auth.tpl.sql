-- $${{define "auth.tokenCreate"}}$$
INSERT INTO auth_tokens (owner_type, owner_id, expires_at, user_agent, ip_address, created_at, updated_at)
VALUES (
   $${{.Bind .Args.OwnerType}}$$,
   $${{.Bind .Args.OwnerID}}$$,
   $${{.Bind .Args.ExpiresAt}}$$,
   $${{.Bind .Args.UserAgent}}$$,
   $${{.Bind .Args.IPAddress}}$$,
   CURRENT_TIMESTAMP,
   CURRENT_TIMESTAMP
) RETURNING id;
-- $${{end}}$$
