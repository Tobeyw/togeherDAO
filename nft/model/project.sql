CREATE TABLE `project`
(
    `project_id` varchar(255) NOT NULL COMMENT 'project id',
    `project_token` varchar(255) NOT NULL COMMENT 'project token',
    PRIMARY KEY(`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;