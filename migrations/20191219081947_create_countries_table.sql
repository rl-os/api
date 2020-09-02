-- migrate:up
create table countries
(
    id serial not null,
    code varchar not null,
    name varchar not null
);

comment on table countries is 'contains all country codes and names ';

create unique index countries_code_uindex
    on countries (code);

create unique index countries_id_uindex
    on countries (id);

create unique index countries_name_uindex
    on countries (name);

alter table countries
    add constraint countries_pk
        primary key (id);


-- INSERTING DATA
INSERT INTO countries (id, code, name) VALUES (0, '-', 'None');
INSERT INTO countries (code, name) VALUES ('FI', 'Finland');
INSERT INTO countries (code, name) VALUES ('RU', 'Russia');
INSERT INTO countries (code, name) VALUES ('BY', 'Belarus');
INSERT INTO countries (code, name) VALUES ('UA', 'Ukraine');
INSERT INTO countries (code, name) VALUES ('MD', 'Moldova');
INSERT INTO countries (code, name) VALUES ('LT', 'Lithuania');
INSERT INTO countries (code, name) VALUES ('LV', 'Latvia');
INSERT INTO countries (code, name) VALUES ('EE', 'Estonia');
INSERT INTO countries (code, name) VALUES ('KZ', 'Kazakhstan');
INSERT INTO countries (code, name) VALUES ('GE', 'Georgia');
INSERT INTO countries (code, name) VALUES ('UZ', 'Uzbekistan');
INSERT INTO countries (code, name) VALUES ('KP', 'North Korea');
INSERT INTO countries (code, name) VALUES ('MN', 'Mongolia');
INSERT INTO countries (code, name) VALUES ('VN', 'Vietnam');
INSERT INTO countries (code, name) VALUES ('CN', 'China');
INSERT INTO countries (code, name) VALUES ('CU', 'Cuba');
INSERT INTO countries (code, name) VALUES ('AL', 'Albania');
INSERT INTO countries (code, name) VALUES ('JP', 'Japan');
INSERT INTO countries (code, name) VALUES ('BA', 'Bosnia and Herzegovina');
INSERT INTO countries (code, name) VALUES ('PL', 'Poland');
INSERT INTO countries (code, name) VALUES ('BG', 'Bulgaria');
INSERT INTO countries (code, name) VALUES ('RO', 'Romania');
INSERT INTO countries (code, name) VALUES ('CZ', 'Czech Republic');
INSERT INTO countries (code, name) VALUES ('HU', 'Hungary');
INSERT INTO countries (code, name) VALUES ('SK', 'Slovakia');
INSERT INTO countries (code, name) VALUES ('AZ', 'Azerbaijan');
INSERT INTO countries (code, name) VALUES ('AM', 'Armenia');
INSERT INTO countries (code, name) VALUES ('KG', 'Kyrgyzstan');
INSERT INTO countries (code, name) VALUES ('IE', 'Ireland');
INSERT INTO countries (code, name) VALUES ('KR', 'South Korea');
INSERT INTO countries (code, name) VALUES ('ME', 'Montenegro');
INSERT INTO countries (code, name) VALUES ('MK', 'North Macedonia');
INSERT INTO countries (code, name) VALUES ('TJ', 'Tajikistan');
INSERT INTO countries (code, name) VALUES ('TM', 'Turkmenistan');
INSERT INTO countries (code, name) VALUES ('AF', 'Afghanistan');
INSERT INTO countries (code, name) VALUES ('GB', 'United Kingdom');
INSERT INTO countries (code, name) VALUES ('ES', 'Spain');
INSERT INTO countries (code, name) VALUES ('RS', 'Serbia');
INSERT INTO countries (code, name) VALUES ('GR', 'Greece');
INSERT INTO countries (code, name) VALUES ('SE', 'Sweden');
INSERT INTO countries (code, name) VALUES ('TR', 'Turkey');
INSERT INTO countries (code, name) VALUES ('NO', 'Norway');
INSERT INTO countries (code, name) VALUES ('HR', 'Croatia');
INSERT INTO countries (code, name) VALUES ('SI', 'Slovenia');
INSERT INTO countries (code, name) VALUES ('DE', 'Germany');
INSERT INTO countries (code, name) VALUES ('AT', 'Austria');
INSERT INTO countries (code, name) VALUES ('LU', 'Luxembourg');
INSERT INTO countries (code, name) VALUES ('IT', 'Italy');
INSERT INTO countries (code, name) VALUES ('NL', 'Netherlands');
INSERT INTO countries (code, name) VALUES ('CH', 'Switzerland');
INSERT INTO countries (code, name) VALUES ('DK', 'Denmark');
INSERT INTO countries (code, name) VALUES ('FR', 'France');
INSERT INTO countries (code, name) VALUES ('BE', 'Belgium');
INSERT INTO countries (code, name) VALUES ('ET', 'Egypt');
INSERT INTO countries (code, name) VALUES ('TN', 'Tunisia');
INSERT INTO countries (code, name) VALUES ('DZ', 'Algeria');
INSERT INTO countries (code, name) VALUES ('MA', 'Morocco');
INSERT INTO countries (code, name) VALUES ('PT', 'Portugal');
INSERT INTO countries (code, name) VALUES ('IL', 'Israel');
INSERT INTO countries (code, name) VALUES ('IR', 'Iran');
INSERT INTO countries (code, name) VALUES ('SY', 'Syria');
INSERT INTO countries (code, name) VALUES ('LB', 'Lebanon');
INSERT INTO countries (code, name) VALUES ('IQ', 'Iraq');
INSERT INTO countries (code, name) VALUES ('ZA', 'South Africa');
INSERT INTO countries (code, name) VALUES ('AU', 'Australia');
INSERT INTO countries (code, name) VALUES ('CA', 'Canada');
INSERT INTO countries (code, name) VALUES ('US', 'United States');

-- migrate:down
drop index countries_code_uindex;
drop index countries_id_uindex;
drop index countries_name_uindex;
drop table countries;
