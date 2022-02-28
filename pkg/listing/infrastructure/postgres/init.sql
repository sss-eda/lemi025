CREATE TYPE mode AS ENUM ('FLASH', 'PC', 'FLASH + PC');

CREATE TABLE IF NOT EXISTS instrument (
    id SERIAL PRIMARY KEY,
    config
    time
    coefficients1
    coefficients2
    gps
    flash
    dac
);

CREATE TABLE IF NOT EXISTS config (
    instrument_id INTEGER,
    sn SMALLINT NOT NULL,
    FOREIGN KEY (instrument_ID) REFERENCES instrument (id)
);

CREATE TABLE IF NOT EXISTS reading (
    id BIGSERIAL NOT NULL,
    instrument_id INT NOT NULL,
    time TIMESTAMPTZ NOT NULL,

    PRIMARY KEY (id),
);

CREATE TABLE IF NOT EXISTS temperature (
    time TIMESTAMPTZ NOT NULL,
    instrument_id INT NOT NULL,

);