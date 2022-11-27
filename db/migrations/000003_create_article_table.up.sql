create table article(
    id int PRIMARY KEY,
    user_id int NULL,
    judul TEXT NULL,
    short_descriptions TEXT NULL,
    descriptions TEXT NULL,
    slug TEXT NULL,
    point int NULL,
    approve int NULL,
    created_at timestamp NULL,
    updated_at timestamp NULL
);

ALTER TABLE `article` CHANGE `id` `id` INT(11) NOT NULL AUTO_INCREMENT;

ALTER TABLE `article` ADD INDEX(`user_id`);
