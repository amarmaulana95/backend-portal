create table article(
    id int PRIMARY KEY,
    user_id int NULL,
    judul TEXT NULL,
    short_descriptions TEXT NULL,
    descriptions TEXT NULL,
    slug TEXT NULL,
    created_at timestamp NULL,
    updated_at timestamp NULL
);
