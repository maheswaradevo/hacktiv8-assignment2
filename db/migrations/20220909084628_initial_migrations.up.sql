CREATE TABLE IF NOT EXISTS `order`(
	order_id INT NOT NULL AUTO_INCREMENT,
	customer_name VARCHAR (255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    PRIMARY KEY (order_id)
);

CREATE TABLE IF NOT EXISTS `item`(
    item_id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    item_code VARCHAR(25) NOT NULL,
    description TEXT NOT NULL,
    quantity INT NOT NULL,
    order_id INT NOT null,
    FOREIGN KEY (order_id) REFERENCES `order`(order_id)
);