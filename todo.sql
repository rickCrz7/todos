drop table if exists todos;

create table owners (
    id varchar(50) primary key,
    name varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null default now()
);
create table todos (
    id varchar(50) primary key,
    title varchar(255) not null,
    completed boolean default false,
    created_at timestamp not null,
    updated_at timestamp not null default now(),
    owner_id varchar(50) not null references owners(id)
);

INSERT INTO owners (id, name, created_at) VALUES ('1', 'Alice', now());
INSERT INTO owners (id, name, created_at) VALUES ('2', 'Bob', now());

INSERT INTO todos (id, title, created_at, owner_id) VALUES ('1', 'Buy milk', now(), '1');
INSERT INTO todos (id, title, created_at, owner_id) VALUES ('2', 'Buy eggs', now(), '1');
INSERT INTO todos (id, title, created_at, owner_id) VALUES ('3', 'Buy bread', now(), '2');
INSERT INTO todos (id, title, created_at, owner_id) VALUES ('4', 'Buy butter', now(), '2');
```