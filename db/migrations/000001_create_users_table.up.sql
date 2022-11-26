create table users(
    id int PRIMARY KEY,
    name TEXT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    avatar TEXT NULL,
    role TEXT NULL,
    created_at timestamp NULL,
    updated_at timestamp NULL
);