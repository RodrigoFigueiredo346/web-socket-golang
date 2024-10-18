-- name: CreatePanel :one
INSERT INTO panel (identifier, dscpanel, num_serie, active, ctrl_bright, dthr_ins, dthr_alt)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING idpanel;

-- name: GetPanelByIdentifier :one
SELECT idpanel --, identifier, dscpanel, num_serie, active, ctrl_bright, dthr_ins, dthr_alt
FROM panel
WHERE identifier = $1;


-- name: UpdatePanel :exec
UPDATE panel
SET identifier = $2, dscpanel = $3, num_serie = $4, active = $5, ctrl_bright = $6, dthr_alt = $7
WHERE idpanel = $1;

-- name: CreateSinc :one
INSERT INTO sinc3 (idpanel, tag, data,  sinc)
VALUES ($1, $2, $3, $4)
RETURNING idsinc;



-- name: GetSincsByPanelID :many
SELECT idsinc, idpanel, tag, data, dthr_ins, sinc, dthr_sinc
FROM sinc3
WHERE idpanel = $1;

-- name: UpdateSincStatus :exec
UPDATE sinc3
SET sinc = $2, dthr_sinc = $3, data = $4
WHERE idsinc = $1;
