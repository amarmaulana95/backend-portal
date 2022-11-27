create table article_images(
    id int PRIMARY KEY,
    article_id int NULL,
    file_name TEXT NULL,
    is_primary TINYINT NULL,
    created_at timestamp NULL,
    updated_at timestamp NULL
);

ALTER TABLE `article_images` CHANGE `id` `id` INT(11) NOT NULL AUTO_INCREMENT;

ALTER TABLE `article_images` ADD INDEX(`article_id`);