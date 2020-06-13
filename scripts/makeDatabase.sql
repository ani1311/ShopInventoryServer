DROP DATABASE shop_inventory;
CREATE DATABASE shop_inventory;
USE shop_inventory;

CREATE TABLE item (
    barcode VARCHAR(30) NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    price INT NOT NULL
);

CREATE TABLE shop( 
    name VARCHAR(50) NOT NULL PRIMARY KEY,
    longitude TEXT NOT NULL,
    latitude TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE shop_item(
    shop_name VARCHAR(50) NOT NULL,
    barcode VARCHAR(30) NOT NULL,
    available BOOLEAN,
    PRIMARY KEY (shop_name,barcode),
    CONSTRAINT FK_shop_name FOREIGN KEY (shop_name)
    REFERENCES shop(name),
    CONSTRAINT FK_barcode FOREIGN KEY (barcode)
    REFERENCES item(barcode)
);

/* SET FOREIGN_KEY_CHECKS = 0; */
