-- +goose Up
-- +goose StatementBegin

CREATE EXTENSION pgcrypto;

CREATE SCHEMA IF NOT EXISTS test;

CREATE TABLE IF NOT EXISTS test.contact
(
    id           uuid         DEFAULT gen_random_uuid()      NOT NULL
    CONSTRAINT pk_contact
    PRIMARY KEY,
    created_at   timestamp,
    modified_at  timestamp,
    name         varchar(50)  DEFAULT '':: character varying NOT NULL,
    surname      varchar(100) DEFAULT '':: character varying NOT NULL,
    patronymic   varchar(100) DEFAULT '':: character varying NOT NULL,
    email        varchar(250),
    phone_number varchar(50),
    age          smallint
    CONSTRAINT age_check
    CHECK ((age >= 0) AND (age <= 200)),
    gender       smallint,
    is_archived  boolean      DEFAULT FALSE                  NOT NULL
    );

CREATE TABLE IF NOT EXISTS test."group"
(
    id            uuid      DEFAULT gen_random_uuid() NOT NULL
    CONSTRAINT pk_group
    PRIMARY KEY,
    created_at    timestamp DEFAULT CURRENT_TIMESTAMP,
    modified_at   timestamp DEFAULT CURRENT_TIMESTAMP,
    name          varchar(250)                        NOT NULL,
    description   varchar(1000),
    contact_count bigint    DEFAULT 0                 NOT NULL,
    is_archived   boolean   DEFAULT FALSE             NOT NULL
    );

CREATE TABLE IF NOT EXISTS test.contact_in_group
(
    id          uuid      DEFAULT gen_random_uuid() NOT NULL
    CONSTRAINT pk_contact_in_group
    PRIMARY KEY,
    created_at  timestamp DEFAULT CURRENT_TIMESTAMP,
    modified_at timestamp DEFAULT CURRENT_TIMESTAMP,
    contact_id  uuid                                not null
    constraint fk_contact_id
    references test.contact,
    group_id    uuid                                not null
    constraint fk_group_id
    references test."group"
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS test.contact_in_group;

DROP TABLE IF EXISTS test.contact;

DROP TABLE IF EXISTS test.group;

DROP SCHEMA IF EXISTS test;

-- +goose StatementEnd
