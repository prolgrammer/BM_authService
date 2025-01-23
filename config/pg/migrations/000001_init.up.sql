create table users
(
    id UUID default gen_random_uuid() primary key,
    name varchar(64) not null,
    surname varchar(64) not null,
    patronymic varchar(64) not null,
    phone_number varchar(15) not null
);


create table accounts
(
    id serial primary key,
    email varchar(255) not null,
    password varchar(100) not null,
    registration_date date not null,
    is_verified boolean default false not null,
    role_ varchar(10) default 'USER' not null,
    user_id UUID references users(id) on delete cascade
);

