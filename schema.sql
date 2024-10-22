-- tabela de usuários
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


-- tabela que armazenará os logs
CREATE TABLE sinc (
    idsinc SERIAL PRIMARY KEY,
    idpanel VARCHAR(50) NOT NULL,
    tag VARCHAR(50) NOT NULL,
    "data" TEXT NULL,
    dthr_ins TIMESTAMP DEFAULT TIME.NOW(), 
    sinc INT DEFAULT 0 NOT NULL,
    dthr_sinc TIMESTAMP DEFAULT TIME.NOW()
);


-- tabela de usuários com acesso ao app
CREATE TABLE user_panel (
    iduser VARCHAR(50) PRIMARY KEY,  -- Identificador único do usuário
    name VARCHAR(10) NOT NULL,       -- Nome do usuário
    pass VARCHAR(10) NOT NULL,       -- Hash da senha do usuário
    active INT DEFAULT 1,            -- Estado ativo do usuário (1 = ativo, 0 = inativo)
    dthr_ins TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Data/hora de inserção
    dthr_alt TIMESTAMP DEFAULT CURRENT_TIMESTAMP  
);

-- tabela de usuários com acesso ao portal - frontend
CREATE TABLE users (
    iduser VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    login VARCHAR(100) NOT NULL UNIQUE,
    pass VARCHAR(100) NOT NULL,
    active INT DEFAULT 1,
    level INT NOT NULL,
    dthr_ins TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    dthr_alt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela que salva os status dos equipamentos (panel_status)
CREATE TABLE panel_status (
    idstatus VARCHAR(50) PRIMARY KEY,
    idpanel VARCHAR(50) NOT NULL,
    status TEXT NOT NULL,
    dthr_ins TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela de configuração de bright sensor (luminosidade do ambiente)
CREATE TABLE bright_lum (
    idlum VARCHAR(50) PRIMARY KEY,
    luminosity INT NOT NULL,
    bright INT NOT NULL,
    dthr_ins TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    dthr_alt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela de configuração de bright hour (controle por horário)
CREATE TABLE bright_time (
    idtime VARCHAR(50) PRIMARY KEY,
    time TIME NOT NULL,
    bright INT NOT NULL,
    dthr_ins TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    dthr_alt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela de mensagens (msg)
CREATE TABLE msg (
    msg INT PRIMARY KEY,  -- Valores de 0 - 19
    dsc VARCHAR(15) NOT NULL,  -- Estrutura JSON da mensagem
    dthr_ins TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    dthr_alt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela de páginas (msg_pag)
CREATE TABLE msg_pag (
    msg INT NOT NULL,
    page INT NOT NULL,  -- Valor de 1 - 8
    data TEXT NOT NULL,  -- HEXA da mensagem
    time_ms INT NOT NULL,  -- Tempo de exibição da página
    active INT DEFAULT 1 NOT NULL,
    dthr_ins TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    dthr_alt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (msg, page)
);

-- Tabela de configuração do acionamento dos ventiladores (fun)
CREATE TABLE fun (
    idfun VARCHAR(50) PRIMARY KEY,
    dsc VARCHAR(50) NOT NULL,
    fun_on INT NOT NULL,  -- Temperatura para ligar
    fun_off INT NOT NULL,  -- Temperatura para desligar
    dthr_ins TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    dthr_alt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela de logs de usuários do sistema Web (user_log)
CREATE TABLE user_log (
    idlog VARCHAR(50) PRIMARY KEY,
    iduser VARCHAR(50) NOT NULL,
    action VARCHAR(50) NOT NULL,  -- Ações como INSERT, LIST, UPDATE, DELETE, LOGIN, CMD
    complete TEXT NOT NULL,
    dthr_ins TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
