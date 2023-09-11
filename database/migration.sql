CREATE TABLE `users`
(
    id   bigint auto_increment,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `users` (`name`)
VALUES ('田中 太郎'),
       ('佐々木 花子');