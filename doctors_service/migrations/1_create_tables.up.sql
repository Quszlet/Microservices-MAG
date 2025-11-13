CREATE TABLE doctors (
    id serial primary key,
    name varchar(50) not null,
    surname varchar(50) not null,
    patronymic varchar(50),
    specialty varchar(50) not null,
    clinic varchar(50) not null,
    status boolean not null
);