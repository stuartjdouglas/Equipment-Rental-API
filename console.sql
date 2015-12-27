use honoursproject;
select * from users;

select * from tokens;

select * from images;

SELECT COUNT(*) from products;

select * from products limit 2, 2;

SELECT EXISTS (SELECT 1 FROM images WHERE file_name LIKE '%A2ocy2zhxM%');

select * from images where products_id = "5";

select * FROM images where file_name = "a_lemon.jpg";

insert into images (file_name, title, date_added, orignal_name) values ();

SELECT EXISTS (SELECT 1 FROM tokens WHERE token = "cb8609cf-daf3-4da0-85e7-908ecddca4af" AND active = 1 AND NOW() <= date_expires
);
SELECT EXISTS (SELECT 1 FROM tokens WHERE token = "4fa49c8e-30a1-495d-b700-181eafc3589c" AND active = 1 AND date_expires <= "2015-12-22");


UPDATE tokens SET active = 1 WHERE idenf = "8a27272a-a5ab-4703-a244-a0c5a475cfd6";

DROP TABLE IF EXISTS `honoursproject`.`images` ;
DROP TABLE IF EXISTS `honoursproject`.`posts` ;
DROP TABLE IF EXISTS `honoursproject`.`tokens` ;
DROP TABLE IF EXISTS `honoursproject`.`has` ;
DROP TABLE IF EXISTS `honoursproject`.`products` ;
DROP TABLE IF EXISTS `honoursproject`.`users` ;

CREATE TABLE IF NOT EXISTS `honoursproject`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '',
  `username` VARCHAR(45) NOT NULL COMMENT '',
  `password` VARCHAR(250) NOT NULL COMMENT '',
  `email` VARCHAR(45) NOT NULL COMMENT '',
  `first_name` VARCHAR(45) NULL DEFAULT 'first_name' COMMENT '',
  `last_name` VARCHAR(45) NULL DEFAULT 'last_name' COMMENT '',
  `location` VARCHAR(45) NULL DEFAULT 'unknown' COMMENT '',
  `bio` VARCHAR(140) NULL DEFAULT 'Please describe me' COMMENT '',
  `date_registered` DATE NOT NULL COMMENT '',
  PRIMARY KEY (`id`)  COMMENT '',
  INDEX `username` (`username` ASC)  COMMENT '')
ENGINE = InnoDB;


CREATE TABLE IF NOT EXISTS `honoursproject`.`tokens` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '',
  `token` VARCHAR(250) NOT NULL COMMENT '',
  `date_generated` DATE NOT NULL COMMENT '',
  `date_expires` DATE NOT NULL COMMENT '',
  `user_id` INT NOT NULL COMMENT '',
  `idenf` VARCHAR(250) NOT NULL COMMENT '',
  `active` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '',
  PRIMARY KEY (`id`)  COMMENT '',
  INDEX `fk_token_user_idx` (`user_id` ASC)  COMMENT '',
  CONSTRAINT `fk_token_user`
    FOREIGN KEY (`user_id`)
    REFERENCES `honoursproject`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `honoursproject`.`posts` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '',
  `title` VARCHAR(140) NULL COMMENT '',
  `slug` VARCHAR(140) NULL COMMENT '',
  `author` VARCHAR(45) NULL COMMENT '',
  `content` VARCHAR(140) NULL COMMENT '',
  `date_created` DATE NULL COMMENT '',
  `date_edited` DATE NULL COMMENT '',
  `users_id` INT NOT NULL COMMENT '',
  PRIMARY KEY (`id`)  COMMENT '',
  INDEX `fk_posts_users1_idx` (`users_id` ASC)  COMMENT '',
  CONSTRAINT `fk_posts_users1`
    FOREIGN KEY (`users_id`)
    REFERENCES `honoursproject`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `honoursproject`.`images` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '',
  `file_name` VARCHAR(256) NOT NULL COMMENT '',
  `title` VARCHAR(256) NOT NULL COMMENT '',
  `date_added` DATETIME NOT NULL COMMENT '',
  `original_name` VARCHAR(256) NOT NULL DEFAULT 'Null' COMMENT '',
  `users_id` INT NOT NULL COMMENT '',
  PRIMARY KEY (`id`)  COMMENT '',
  INDEX `fk_images_users1_idx` (`users_id` ASC)  COMMENT '',
  CONSTRAINT `fk_images_users1`
    FOREIGN KEY (`users_id`)
    REFERENCES `honoursproject`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;





CREATE TABLE IF NOT EXISTS `honoursproject`.`products` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '',
  `product_name` VARCHAR(240) NOT NULL COMMENT '',
  `product_id` VARCHAR(240) NOT NULL COMMENT '',
  `date_added` DATETIME NOT NULL COMMENT '',
  `date_updated` DATETIME NOT NULL COMMENT '',
  `product_description` VARCHAR(240) NOT NULL COMMENT '',
  `product_rental_period_limit` VARCHAR(240) NOT NULL COMMENT '',
  `product_image_id` VARCHAR(240) NOT NULL COMMENT '',
  `owner_id` INT NOT NULL COMMENT '',
  PRIMARY KEY (`id`)  COMMENT '')
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `honoursproject`.`has` (
  `users_id` INT NOT NULL COMMENT '',
  `products_id` INT NOT NULL COMMENT '',
  `status` INT NOT NULL DEFAULT 0 COMMENT '',
  PRIMARY KEY (`users_id`, `products_id`)  COMMENT '',
  INDEX `fk_users_has_products_products1_idx` (`products_id` ASC)  COMMENT '',
  INDEX `fk_users_has_products_users1_idx` (`users_id` ASC)  COMMENT '',
  CONSTRAINT `fk_users_has_products_users1`
    FOREIGN KEY (`users_id`)
    REFERENCES `honoursproject`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_users_has_products_products1`
    FOREIGN KEY (`products_id`)
    REFERENCES `honoursproject`.`products` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `honoursproject`.`user_rent_product` (
  `users_id` INT NOT NULL COMMENT '',
  `products_id` INT NOT NULL COMMENT '',
  `date_received` DATETIME NOT NULL COMMENT '',
  `date_due` DATETIME NOT NULL COMMENT '',
  PRIMARY KEY (`users_id`, `products_id`)  COMMENT '',
  INDEX `fk_users_has_products_products2_idx` (`products_id` ASC)  COMMENT '',
  INDEX `fk_users_has_products_users2_idx` (`users_id` ASC)  COMMENT '',
  CONSTRAINT `fk_users_has_products_users2`
    FOREIGN KEY (`users_id`)
    REFERENCES `honoursproject`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_users_has_products_products2`
    FOREIGN KEY (`products_id`)
    REFERENCES `honoursproject`.`products` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


CREATE TABLE IF NOT EXISTS `honoursproject`.`Site` (
  `id` INT NOT NULL COMMENT '',
  `Title` VARCHAR(45) NULL COMMENT '',
  `Description` VARCHAR(140) NULL COMMENT '',
  PRIMARY KEY (`id`)  COMMENT '')
ENGINE = InnoDB;


select * from images;

select * from images where file_name = "15b2459a-705a-4831-92fb-d69cba3ed3eb.gif";


select * from images ORDER BY date_added DESC;

DROP PROCEDURE imageExists;

 CREATE PROCEDURE `imageExists`(code VARCHAR(240))
    BEGIN
      SELECT EXISTS (SELECT 1 FROM images WHERE file_name LIKE CONCAT('%', code , '%'));
    END;
SELECT EXISTS (SELECT 1 FROM images WHERE file_name LIKE "%1OxlR3nLip%");
CALL imageExists('1OxlR3nLip');

select * from products;

CREATE PROCEDURE helloWorld()
  BEGIN
    select * from users;
  END;

DROP PROCEDURE createProduct;

CREATE PROCEDURE createProduct (product_name VARCHAR(240), product_id VARCHAR(240), date_added DATETIME, date_updated DATETIME, product_description VARCHAR(240), product_rental_period_limit VARCHAR(240), product_image_id VARCHAR(240), owner_id INT)
  BEGIN
    INSERT INTO products (product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, product_image_id, owner_id)
    values (product_name,product_id,date_added,date_updated,product_description,product_rental_period_limit,product_image_id,owner_id);
    INSERT INTO has (users_id, products_id, status) VALUES (owner_id, LAST_INSERT_ID(), 0);
  END;;

INSERT INTO has (users_id, products_id, status) VALUES (1, 1, 0);

call createProduct("item","something","2015-12-16","2015-12-16","something",7,0,1);

INSERT INTO products (product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, product_image_id, owner_id)
values ("item","something3","2015-12-27","2015-12-27","something",7,0,1);

DROP PROCEDURE getOwnerProducts;

CREATE PROCEDURE getOwnerProducts(username VARCHAR(240))
  BEGIN
    SELECT product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, product_image_id, username as owner from has
      LEFT OUTER JOIN products ON has.products_id = products.id
      LEFT OUTER JOIN users ON has.users_id = users.id
      where users.username = username
      AND has.status = 0;
  END;

call getOwnerProducts("remon");

SELECT * from has
  LEFT OUTER JOIN products ON has.products_id = products.id
  LEFT OUTER JOIN users ON has.users_id = users.id
  GROUP BY users.username;

call RentItem ("somethin2g", "remon");
DROP PROCEDURE RentItem;

CREATE PROCEDURE RentItem (product VARCHAR(240), username VARCHAR(240))
  BEGIN
    DECLARE userid INT;
    DECLARE days INT;
    DECLARE productid INT;

    SELECT id INTO userid FROM users WHERE username = username LIMIT 1;
    SELECT id, product_rental_period_limit INTO productid, days FROM products WHERE product_id = product;

    INSERT INTO user_rent_product (products_id, users_id, date_received, date_due) VALUES (productid, userid, NOW(), DATE_ADD(CURDATE(), INTERVAL days DAY));
  END;


SELECT id FROM products WHERE product_id = "something" LIMIT 1;
 SELECT product_rental_period_limit FROM products WHERE product_id = "something" LIMIT 1;
select product_rental_period_limit from products where product_id = "something";

DROP PROCEDURE checkItemAvailability;

CREATE PROCEDURE `checkItemAvailability`(product VARCHAR(240))
BEGIN
    DECLARE due_date DATETIME;
    select date_due into due_date from user_rent_product where products_id = 7;
    if (due_date < NOW()) THEN
      select TRUE , due_date;
      ELSE
        select FALSE as Available, due_date as date_;
    END IF;
  END;

DROP PROCEDURE getPagedProducts;
CALL getPagedProducts(0, 2);

CREATE PROCEDURE getPagedProducts (step INT, count INT)
  BEGIN
    SELECT product_id as id, product_name as name, product_description as description, product_rental_period_limit as time_period, product_image_id as image_id, username as owner from has
    LEFT OUTER JOIN products ON has.products_id = products.id
    LEFT OUTER JOIN users ON has.users_id = users.id;
  END;

DROP PROCEDURE getRentedProducts;
CALL getRentedProducts("remon");

CREATE PROCEDURE getRentedProducts (username VARCHAR(240))
  BEGIN
    select product_id as id, product_name as name, product_description as description, product_rental_period_limit as time_period, product_image_id as image_id, username as owner from user_rent_product
      LEFT OUTER JOIN products ON user_rent_product.products_id = products.id;
  END;

select id from users where username = 'remon';

SELECT
  *
FROM
  products
  JOIN has ON users.id = has.products_id;