drop table if exists todos;
create table todos (
    id serial primary key,
    title varchar(255) not null,
    completed boolean default false,
    created_at timestamp not null,
    updated_at timestamp not null default now()
);