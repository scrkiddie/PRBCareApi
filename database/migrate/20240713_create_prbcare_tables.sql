CREATE TYPE status_pasien_enum AS ENUM ('aktif', 'selesai');
CREATE TYPE status_pengambilan_obat_enum AS ENUM ('menunggu', 'diambil', 'batal');
CREATE TYPE status_kontrol_balik_enum AS ENUM ('menunggu', 'selesai', 'batal');

CREATE TABLE admin_super
(
    id       SERIAL PRIMARY KEY,
    username VARCHAR(50)  NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE admin_puskesmas
(
    id             SERIAL PRIMARY KEY,
    nama_puskesmas VARCHAR(100) NOT NULL,
    telepon        VARCHAR(15)  NOT NULL UNIQUE,
    alamat         TEXT         NOT NULL,
    username       VARCHAR(50)  NOT NULL UNIQUE,
    password       VARCHAR(255) NOT NULL
);

CREATE TABLE admin_apotek
(
    id          SERIAL PRIMARY KEY,
    nama_apotek VARCHAR(100) NOT NULL,
    telepon     VARCHAR(15)  NOT NULL UNIQUE,
    alamat      TEXT         NOT NULL,
    username    VARCHAR(50)  NOT NULL UNIQUE,
    password    VARCHAR(255) NOT NULL
);

CREATE TABLE pengguna
(
    id               SERIAL PRIMARY KEY,
    token_perangkat  VARCHAR(255),
    nama_lengkap     VARCHAR(100) NOT NULL,
    telepon          VARCHAR(15)  NOT NULL UNIQUE,
    telepon_keluarga VARCHAR(15)  NOT NULL,
    alamat           TEXT         NOT NULL,
    username         VARCHAR(50)  NOT NULL UNIQUE,
    password         VARCHAR(255) NOT NULL
);

CREATE TABLE obat
(
    id             SERIAL PRIMARY KEY,
    id_admin_apotek INT REFERENCES admin_apotek (id) NOT NULL,
    nama_obat      VARCHAR(100) NOT NULL,
    jumlah         INT          NOT NULL
);

CREATE TABLE pasien
(
    id                 SERIAL PRIMARY KEY,
    no_rekam_medis     VARCHAR(50)                         NOT NULL,
    id_pengguna        INT REFERENCES pengguna (id) NOT NULL,
    id_admin_puskesmas INT REFERENCES admin_puskesmas (id) NOT NULL,
    berat_badan        DECIMAL(5, 2)                       NOT NULL,
    tinggi_badan       DECIMAL(5, 2)                       NOT NULL,
    tekanan_darah      VARCHAR(20)                         NOT NULL,
    denyut_nadi        INT                                 NOT NULL,
    hasil_lab          TEXT,
    hasil_ekg          TEXT,
    tanggal_periksa    BIGINT                              NOT NULL,
    status             status_pasien_enum                  NOT NULL
);

CREATE TABLE kontrol_balik
(
    id              SERIAL PRIMARY KEY,
    id_pasien       INT REFERENCES pasien (id) NOT NULL,
    tanggal_kontrol BIGINT                     NOT NULL,
    status          status_kontrol_balik_enum  NOT NULL
);

CREATE TABLE pengambilan_obat
(
    id                  SERIAL PRIMARY KEY,
    resi                VARCHAR(50)                      NOT NULL,
    id_pasien           INT REFERENCES pasien (id) NOT NULL,
    id_obat             INT REFERENCES obat (id) NOT NULL,
    jumlah              INT                              NOT NULL,
    tanggal_pengambilan BIGINT                           NOT NULL,
    status              status_pengambilan_obat_enum     NOT NULL
);
