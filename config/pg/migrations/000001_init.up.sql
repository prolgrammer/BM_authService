create table accounts
(
    id UUID default gen_random_uuid() primary key,
    email varchar(255) not null,
    password varchar(100) not null,
    registration_date date not null,
    is_verified boolean default false not null
)