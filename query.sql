-- name: CreatePanel :one
INSERT INTO panel (identifier, dsc_panel, num_serie, active, ctrl_bright )
VALUES ($1, $2, $3, $4, $5)
RETURNING idpanel;

-- name: GetPanelByIdentifier :one
SELECT idpanel --, identifier, dsc_panel, num_serie, active, ctrl_bright, dthr_ins, dthr_alt
FROM panel
WHERE identifier = $1;


-- name: UpdatePanel :exec
UPDATE panel
SET identifier = $2, dsc_panel = $3, num_serie = $4, active = $5, ctrl_bright = $6, dthr_alt = $7
WHERE idpanel = $1;

--------------------sinc-------------------------------
-- name: CreateSinc :one
INSERT INTO sinc (idpanel, tag, data,  sinc)
VALUES ($1, $2, $3, $4)
RETURNING idsinc;

-- name: GetSincsByPanelID :many
SELECT idsinc, idpanel, tag, data, dthr_ins, sinc, dthr_sinc
FROM sinc
WHERE idpanel = $1;

-- name: UpdateSincStatus :exec
UPDATE sinc
SET sinc = $2, dthr_sinc = $3, data = $4
WHERE idsinc = $1;


--------------------users-------------------------------
-- name: CreateUser :one
INSERT INTO users (iduser, name, login, pass, active, level)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING iduser;

-- name: UpdateUser :exec
UPDATE users
SET name = $2, login = $3, pass = $4, active = $5, level = $6, dthr_alt = CURRENT_TIMESTAMP
WHERE iduser = $1;

-- name: GetUserByID :one
SELECT iduser, name, login, pass, active, level, dthr_ins, dthr_alt
FROM users
WHERE iduser = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE iduser = $1;

-- name: GetActiveUsers :many
SELECT iduser, name, login, pass, active, level, dthr_ins, dthr_alt
FROM users
WHERE active = 1;

-- name: GetUsersByLevel :many
SELECT iduser, name, login, pass, active, level, dthr_ins, dthr_alt
FROM users
WHERE level = $1;

-- name: GetUserByLoginAndPassword :one
SELECT iduser, name, login, pass, active, level, dthr_ins, dthr_alt
FROM users
WHERE login = $1 AND pass = $2;

----------------------panel_status-----------------------------
-- name: CreatePanelStatus :exec
INSERT INTO panel_status (idstatus, idpanel, status)
VALUES ($1, $2, $3);

-- name: GetPanelStatusById :one
SELECT idstatus, idpanel, status, dthr_ins
FROM panel_status
WHERE idstatus = $1;

-- name: UpdatePanelStatus :exec
UPDATE panel_status
SET status = $2, dthr_ins = DEFAULT
WHERE idstatus = $1;


--------------------bright_lum-----------------------------
-- name: CreateBrightLum :exec
INSERT INTO bright_lum (idlum, luminosity, bright)
VALUES ($1, $2, $3);

-- name: GetBrightLumById :one
SELECT idlum, luminosity, bright, dthr_ins, dthr_alt
FROM bright_lum
WHERE idlum = $1;

-- name: UpdateBrightLum :exec
UPDATE bright_lum
SET luminosity = $2, bright = $3, dthr_alt = DEFAULT
WHERE idlum = $1;


--------------------bright_time-----------------------------
-- name: CreateBrightTime :exec
INSERT INTO bright_time (idtime, time, bright)
VALUES ($1, $2, $3);

-- name: GetBrightTimeById :one
SELECT idtime, time, bright, dthr_ins, dthr_alt
FROM bright_time
WHERE idtime = $1;

-- name: UpdateBrightTime :exec
UPDATE bright_time
SET time = $2, bright = $3, dthr_alt = DEFAULT
WHERE idtime = $1;


---------------------msg-----------------------------
-- name: CreateMsg :exec
INSERT INTO msg (msg, dsc)
VALUES ($1, $2);

-- name: GetMsgById :one
SELECT msg, dsc, dthr_ins, dthr_alt
FROM msg
WHERE msg = $1;

-- name: UpdateMsg :exec
UPDATE msg
SET dsc = $2, dthr_alt = DEFAULT
WHERE msg = $1;


---------------------msg_pag-----------------------------
-- name: CreateMsgPag :exec
INSERT INTO msg_pag (msg, page, data, time_ms, active)
VALUES ($1, $2, $3, $4, $5);

-- name: GetMsgPagByMsgAndPage :one
SELECT msg, page, data, time_ms, active, dthr_ins, dthr_alt
FROM msg_pag
WHERE msg = $1 AND page = $2;

-- name: UpdateMsgPag :exec
UPDATE msg_pag
SET data = $3, time_ms = $4, active = $5, dthr_alt = DEFAULT
WHERE msg = $1 AND page = $2;



-------------------fun-------------------------------
-- name: CreateFun :exec
INSERT INTO fun (idfun, dsc, fun_on, fun_off)
VALUES ($1, $2, $3, $4);

-- name: GetFunById :one
SELECT idfun, dsc, fun_on, fun_off, dthr_ins, dthr_alt
FROM fun
WHERE idfun = $1;

-- name: UpdateFun :exec
UPDATE fun
SET dsc = $2, fun_on = $3, fun_off = $4, dthr_alt = DEFAULT
WHERE idfun = $1;



-------------------user_log-------------------------------
-- name: CreateUserLog :exec
INSERT INTO user_log (idlog, iduser, action, complete)
VALUES ($1, $2, $3, $4);

-- name: GetUserLogById :one
SELECT idlog, iduser, action, complete, dthr_ins
FROM user_log
WHERE idlog = $1;

-- name: ListUserLogs :many
SELECT idlog, iduser, action, complete, dthr_ins
FROM user_log
ORDER BY dthr_ins DESC;
