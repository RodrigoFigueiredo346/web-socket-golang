create TABLE panel (
    idpanel SERIAL PRIMARY KEY,
    identifier VARCHAR(50) NOT NULL,
    dscpanel VARCHAR(50) NOT NULL,
    num_serie VARCHAR(50) NOT NULL,
    active INT DEFAULT 1, 
    ctrl_bright INT DEFAULT 1, 
    dthr_ins TIMESTAMP,
    dthr_alt TIMESTAMP
);

CREATE TABLE sinc3 (
    idsinc SERIAL PRIMARY KEY,
    idpanel VARCHAR(50) NOT NULL,
    tag VARCHAR(50) NOT NULL,
    "data" TEXT NULL,
    dthr_ins TIMESTAMP DEFAULT TIME.NOW(), 
    sinc INT DEFAULT 0 NOT NULL,
    dthr_sinc TIMESTAMP DEFAULT TIME.NOW()
);
