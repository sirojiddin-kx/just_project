CREATE TABLE IF NOT EXISTS profession (
    "id" uuid primary key,
    "name" varchar(255) not null
);

CREATE TYPE states AS ENUM(
    'datetime',
    'text',
    'number'
);

CREATE TABLE IF NOT EXISTS attribute (
    "id" uuid primary key,
    "name" varchar(255) not null,
    "current_state" states
);

CREATE TABLE IF NOT EXISTS position (
    "id" uuid primary key,
    "name" varchar(255),
    "profession_id" uuid,
    "company_id" uuid,
    constraint fk_profession foreign key("profession_id") references profession ("id"),
    constraint fk_company foreign key("company_id") references company ("id")
);

CREATE TABLE IF NOT EXISTS position_attribute (
    "id" uuid primary key,
    "attribute_id" uuid,
    "position_id" uuid,
    "value" varchar(200),
    constraint fk_attribute foreign key("attribute_id") references attribute ("id"),
    constraint fk_position foreign key("position_id") references position ("id") 
);

