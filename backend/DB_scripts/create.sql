\c rsoi_persons

create table if not exists Persons (
    id serial primary key,
    first_name varchar(600),
    last_name varchar(600),
    age int
);