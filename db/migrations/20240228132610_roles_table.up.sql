CREATE TABLE IF NOT EXISTS roles
(
    id            serial primary key not null,
    name          int not null,
    title         varchar(255) not null,
);