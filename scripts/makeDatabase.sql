DROP DATABASE shop_inventory;
CREATE DATABASE shop_inventory;
USE shop_inventory;

CREATE TABLE item (
    barcode VARCHAR(30) NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    price INT NOT NULL
);

CREATE TABLE shop(
    shopid INT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    longitude TEXT NOT NULL,
    latitude TEXT NOT NULL
);

CREATE TABLE shop_item(
    shopid INT NOT NULL,
    barcode VARCHAR(30) NOT NULL,
    quantity INT NOT NULL,
    PRIMARY KEY (shopid,barcode),
    CONSTRAINT FK_shopid FOREIGN KEY (shopid)
    REFERENCES shop(shopid),
    CONSTRAINT FK_barcode FOREIGN KEY (barcode)
    REFERENCES item(barcode)
);
