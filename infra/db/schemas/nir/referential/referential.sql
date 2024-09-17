-- name: GetResult :one
SELECT referential_id, "name", "label", image_path, params, company_id, user_audit_id
FROM nir.referential
LIMIT 1;

-- name: GetResults :many
SELECT referential_id, "name", "label", image_path, params, company_id, user_audit_id
FROM nir.referential;