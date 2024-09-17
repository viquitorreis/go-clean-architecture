-- name: GetResult :one
SELECT * FROM nir.result
where result_id = $1;

-- name: GetResults :many
SELECT * FROM nir.result
LIMIT 100;