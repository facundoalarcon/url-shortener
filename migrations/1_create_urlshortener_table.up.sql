CREATE TABLE IF NOT EXISTS `urlshortener` (
    `id`             int(11) NOT NULL AUTO_INCREMENT,
    `original_url`   varchar(150),
    `short`          varchar(150),
    `short_url`      varchar(150),
    `created_at`     DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     DATETIME ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);