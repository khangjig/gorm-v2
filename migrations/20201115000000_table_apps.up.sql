CREATE TABLE apps
(
    `id`         BIGINT(20)   NOT NULL AUTO_INCREMENT,
    `signature`  TEXT DEFAULT NULL,
    `secret_key` TEXT DEFAULT NULL,
    `name`       VARCHAR(255) NOT NULL,
    `created_at` DATETIME     NOT NULL,
    `updated_at` DATETIME     NOT NULL,

    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
