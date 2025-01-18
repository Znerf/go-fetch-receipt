CREATE TABLE IF NOT EXISTS `receipts` (
    `id`              BINARY(16)     NOT NULL,
    `retailer`        VARCHAR(255)   NOT NULL,
    `purchase_date`    VARCHAR(255)           NOT NULL,
    `purchase_time`    VARCHAR(255)           NOT NULL,
    `total`           VARCHAR(255)            NOT NULL,
    PRIMARY KEY (`id`)
);


CREATE TABLE IF NOT EXISTS `items` (
    `id`                   BINARY(16)              NOT NULL,
    `receipt_id`           BINARY(16)              NOT NULL,
    `short_description`    VARCHAR(255)            NOT NULL,
    `price`                VARCHAR(255)          NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `items_ibfk_receipts` FOREIGN KEY (`receipt_id`)
    REFERENCES `receipts` (`id`) ON DELETE CASCADE
);
