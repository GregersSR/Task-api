DROP TABLE users IF EXISTS;
CREATE TABLE users (
    id BIGINT NOT NULL AUTO_INCREMENT,
    name VARCHAR(64) NOT NULL,
    email VARCHAR(64) NOT NULL,
    admin BOOLEAN NOT NULL,
    token VARCHAR(64) NOT NULL,
    PRIMARY KEY (id)
);

DROP TABLE devices IF EXISTS;
CREATE TABLE devices (
    id BIGINT NOT NULL AUTO_INCREMENT,
    name VARCHAR(64) NOT NULL,
    description VARCHAR(1024) NOT NULL,
    token VARCHAR(64) NOT NULL,
    PRIMARY KEY (id)
);

DROP TABLE tasks IF EXISTS;
CREATE TABLE tasks (
    id BIGINT NOT NULL AUTO_INCREMENT,
    title VARCHAR(64) NOT NULL,
    deviceid BIGINT NOT NULL,
    state SMALLINT,
    PRIMARY KEY (id),
    FOREIGN KEY (deviceid) REFERENCES devices(id)
);

DROP TABLE responses IF EXISTS;
CREATE TABLE responses (
    id BIGINT NOT NULL AUTO_INCREMENT,
    taskid BIGINT NOT NULL,
    state SMALLINT,
    PRIMARY KEY (id),
    FOREIGN KEY (taskid) REFERENCES tasks(id)
);

DROP TABLE controls IF EXISTS;
CREATE TABLE controls (
    userid BIGINT NOT NULL,
    deviceid BIGINT NOT NULL,
    FOREIGN KEY (userid) REFERENCES users(id),
    FOREIGN KEY (deviceid) REFERENCES devices(id)
);
