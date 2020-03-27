drop table if exists userInfo;
create table `userInfo` (
    `id` BIGINT NOT NULL,
    `password` VARCHAR(32) DEFAULT '',
    `name` VARCHAR(32) DEFAULT '',
    `age` INT DEFAULT '0',
    `gender` INT DEFAULT '0',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;