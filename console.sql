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

DROP TABLE user_rent_product;

CREATE TABLE IF NOT EXISTS `honoursproject`.`user_rent_product` (
  `users_id` INT NOT NULL COMMENT '',
  `products_id` INT NOT NULL COMMENT '',
  `date_received` DATETIME NOT NULL COMMENT '',
  `date_due` DATETIME NOT NULL COMMENT '',
  `active` BOOLEAN NOT NULL COMMENT '',
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

call createProduct("item3","something3","2015-12-27","2015-12-27","something",7,0,1);

INSERT INTO products (product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, product_image_id, owner_id)
values ("item","something3","2015-12-27","2015-12-27","something",7,0,1);


SELECT * from has
  LEFT OUTER JOIN products ON has.products_id = products.id
  LEFT OUTER JOIN users ON has.users_id = users.id
  GROUP BY users.username;

call RentItem ("something3", "remon");
DROP PROCEDURE RentItem;

CREATE PROCEDURE RentItem (product VARCHAR(240), usrname VARCHAR(240))
  BEGIN
    DECLARE userid INT;
    DECLARE days INT;
    DECLARE productid INT;

    SELECT id INTO userid FROM users WHERE username = usrname;
    SELECT id, product_rental_period_limit INTO productid, days FROM products WHERE product_id = product LIMIT 1;

    INSERT INTO user_rent_product (products_id, users_id, date_received, date_due, active) VALUES (productid, userid, NOW(), DATE_ADD(CURDATE(), INTERVAL days DAY), TRUE);
  END;

DROP PROCEDURE ReturnItem;

CALL ReturnItem("5f7f2bfb-c8f8-44e6-b05d-f8e59adc1722", "2d85d358-736a-405a-b6ac-9c0fe5af29ff");

CREATE PROCEDURE ReturnItem (o_token VARCHAR(240), product VARCHAR(240))
  BEGIN
    DECLARE productid INT;
    DECLARE u_id VARCHAR(240);
    DECLARE tmp_u_id VARCHAR(240);

     select user_id into tmp_u_id from tokens
      where token = o_token;

    select users_id INTO u_id from user_rent_product
    LEFT JOIN products ON user_rent_product.products_id = products.id
    LEFT JOIN users ON owner_id = users.id
    WHERE users_id = tmp_u_id AND product_id = product;

    SELECT id INTO productid FROM products WHERE product_id = product;
    DELETE FROM user_rent_product WHERE users_id = u_id AND products_id = productid;
  END;

select user_id from tokens
    left JOIN users ON tokens.user_id = users.id
    where token = "43c4a78f-7fb2-449d-8011-7914926e4cc3";

DROP PROCEDURE ReturnItemAsOwner;
CALL ReturnItemAsOwner("43c4a78f-7fb2-449d-8011-7914926e4cc3", "2d85d358-736a-405a-b6ac-9c0fe5af29ff");

CREATE PROCEDURE ReturnItemAsOwner (o_token VARCHAR(240), product VARCHAR(240))
  BEGIN
    DECLARE productid INT;
    DECLARE tmp_u_id INT;
    DECLARE u_id VARCHAR(240);

    select user_id into tmp_u_id from tokens
    where token = o_token;

#     select tmp_u_id;

    select users_id INTO u_id from user_rent_product
    LEFT JOIN products ON user_rent_product.products_id = products.id
    LEFT JOIN users ON owner_id = users.id
    WHERE owner_id = tmp_u_id AND product_id = product;

    SELECT id INTO productid FROM products WHERE product_id = product;

    DELETE FROM user_rent_product WHERE users_id = u_id AND products_id = productid;
  END;


SELECT id FROM products WHERE product_id = "something" LIMIT 1;
 SELECT product_rental_period_limit FROM products WHERE product_id = "something" LIMIT 1;
select product_rental_period_limit from products where product_id = "something";

DROP PROCEDURE checkItemAvailability;
CALL checkItemAvailability("works so well", "remon");

CREATE PROCEDURE `checkItemAvailability`(product VARCHAR(240), usrname VARCHAR(240))
BEGIN
  DECLARE due_date DATETIME;
  DECLARE activestate BOOLEAN;

  SELECT date_due, active INTO due_date, activestate FROM user_rent_product
    LEFT OUTER JOIN products ON user_rent_product.products_id = products.id
    LEFT OUTER JOIN users ON user_rent_product.users_id = users.id
    WHERE products.product_id = product
    ORDER BY products.date_updated DESC;

    if (due_date > NOW() || activestate) THEN
      select FALSE , due_date;
      ELSE
        select TRUE as Available, NOW();
    END IF;
  END;

DROP PROCEDURE getPagedProducts;
CALL getPagedProducts(0, 6);

CREATE PROCEDURE getPagedProducts (step INT, count INT)
  BEGIN
    SELECT product_id as id, product_name as name, product_description as description, date_added, date_updated, product_rental_period_limit as time_period, product_image_id as image_id, username as owner from has
    LEFT OUTER JOIN products ON has.products_id = products.id
    LEFT OUTER JOIN users ON has.users_id = users.id
    ORDER BY products.date_updated DESC LIMIT step, count;
  END;

DROP PROCEDURE getRentedProducts;
CALL getRentedProducts("remon", 0, 1);

CREATE PROCEDURE getRentedProducts (username VARCHAR(240), step INT, count INT)
  BEGIN
    select product_id as id, product_name as name, product_description as description, date_added, date_updated, product_rental_period_limit as time_period, product_image_id as image_id, username as owner from user_rent_product
    LEFT OUTER JOIN products ON user_rent_product.products_id = products.id
    WHERE user_rent_product.date_due < NOW()
    ORDER BY products.date_updated DESC LIMIT step, count;
  END;

DROP PROCEDURE getCurrentlyRentingProducts;
CALL getCurrentlyRentingProducts("remon", 0, 2);

CREATE PROCEDURE getCurrentlyRentingProducts (username VARCHAR(240), step INT, count INT)
  BEGIN
    select product_id as id, product_name as name, product_description as description, date_due as due_date, date_received as received_date, product_image_id as image_id, username as owner from user_rent_product
    LEFT OUTER JOIN products ON user_rent_product.products_id = products.id
    WHERE user_rent_product.date_due > NOW()
    ORDER BY products.date_updated DESC LIMIT step, count;
  END;


DROP PROCEDURE getUsername;
CALL getUsername("94a17bfa-6c49-4398-8155-137f07612f7d");

CREATE PROCEDURE getUsername (usrtoken VARCHAR(240))
  BEGIN
    DECLARE userid INT;
    SELECT user_id INTO userid FROM tokens WHERE token = usrtoken;
    SELECT username FROM users WHERE id = userid;
  END;

DROP PROCEDURE getCurrentlyRentingProducts;
CALL getCurrentlyRentingProducts("remon", 0, 1);

SELECT user_id from tokens where token = "94a17bfa-6c49-4398-8155-137f07612f7d";

DROP PROCEDURE checkProductAvailability;
CALL checkProductAvailability("works so well");

CREATE PROCEDURE `checkProductAvailability`(product VARCHAR(240))
BEGIN
  DECLARE due_date DATETIME;
  DECLARE active_state BOOLEAN;

  SELECT date_due, active INTO due_date, active_state FROM user_rent_product
    LEFT OUTER JOIN products ON user_rent_product.products_id = products.id
    LEFT OUTER JOIN users ON user_rent_product.users_id = users.id
    WHERE products.product_id = product
    ORDER BY products.date_updated DESC;

    if (active_state = 1) THEN
      if (due_date > NOW()) THEN
        select FALSE as available, due_date as due_date;
      ELSE
        select TRUE as available, NOW() as due_date;
      END IF;
    ELSE
      select TRUE as available, NOW() as due_date;
    END IF;
  END;

DROP PROCEDURE checkAuthedProductAvailability;
CALL checkAuthedProductAvailability("40cdb44f-3bf4-4d71-9299-0f0887417731");
CALL checkAuthedProductAvailability("Windos");

CREATE PROCEDURE `checkAuthedProductAvailability`(product VARCHAR(240))
BEGIN
    DECLARE due_date DATETIME;
    DECLARE taken_date DATETIME;
    DECLARE user_name VARCHAR(240);
    DECLARE active_state BOOLEAN;

  SELECT date_due, date_received, username, active INTO due_date, taken_date, user_name, active_state FROM user_rent_product
    LEFT JOIN products ON user_rent_product.products_id = products.id
    LEFT JOIN users ON user_rent_product.users_id = users.id
    WHERE products.product_id = product
    ORDER BY date_received DESC
    LIMIT 1;
if (user_name) THEN
    if (active_state = 1) THEN
      if (due_date > NOW()) THEN
        select FALSE as available, due_date, taken_date, user_name as username;
      ELSE
        select TRUE as available, NOW() as due_date, NOW() as taken_date, user_name as username;
      END IF;
    ELSE
      SET user_name = "nil";
      select TRUE as available, NOW() as due_date, NOW() as taken_date, user_name as username;
    END IF;
ELSE
  select FALSE as available, due_date, taken_date, user_name as username;
END IF;
  END;

DROP PROCEDURE  getOwnerProducts;
CALL getOwnerProducts("94a17bfa-6c49-4398-8155-137f07612f7d", 0, 15);

CREATE PROCEDURE getOwnerProducts(u_token VARCHAR(240), step INT, count INT)
BEGIN
  DECLARE usrname VARCHAR(240);
  SELECT username into usrname from tokens
  LEFT OUTER JOIN users on tokens.user_id = users.id
  where token = u_token;
  select  product_id as id, product_name as name, product_description as description, date_added, date_updated, product_rental_period_limit as time_period, product_image_id as image_id, username as owner from has
      LEFT OUTER JOIN products ON has.products_id = products.id
      LEFT OUTER JOIN users ON has.users_id = users.id
      where users.username = usrname
      ORDER BY date_updated DESC
      LIMIT step, count;
  END;

DROP PROCEDURE CheckProductAvailabilityOwner;
CALL CheckProductAvailabilityOwner("94a17bfa-6c49-4398-8155-137f07612f7d", "4ded9e43-174f-4203-a48b-58f34dc9b90b");

CREATE PROCEDURE CheckProductAvailabilityOwner(o_token VARCHAR(240), p_id VARCHAR(240))
  BEGIN
    DECLARE due_date DATETIME;
    DECLARE active_state BOOLEAN;

    SELECT date_due, active INTO due_date, active_state FROM user_rent_product
    LEFT OUTER JOIN products ON user_rent_product.products_id = products.id
    LEFT OUTER JOIN users ON user_rent_product.users_id = users.id
    WHERE products.product_id = p_id
    ORDER BY products.date_updated DESC;

    if (active_state = 1) THEN
      if (due_date > NOW()) THEN
        select FALSE as available, due_date as due_date;
      ELSE
        select TRUE as available, NOW() as due_date;
      END IF;
    ELSE
      select TRUE as available, NOW() as due_date;
    END IF;
  END;


DROP PROCEDURE isOwner;
CALL isOwner("94a17bfa-6c49-4398-8155-137f07612f7d", "4ded9e43-174f-4203-a48b-58f34dc9b90b");

CREATE PROCEDURE isOwner(u_token VARCHAR(240), p_id VARCHAR(240))
  BEGIN
    DECLARE u_id INT;
    SELECT user_id INTO u_id FROM tokens
    WHERE token = u_token;

    SELECT EXISTS(
      SELECT * FROM has
      LEFT JOIN users ON has.users_id = users.id
      LEFT JOIN products ON has.products_id = products.id
      WHERE product_id = p_id AND users_id = u_id) as owner;
  END;

DROP PROCEDURE ownerProductStatus;
CALL ownerProductStatus("194b1286-4585-4ce6-8897-94c3cd9473e9", "4ded9e43-174f-4203-a48b-58f34dc9b90b");

CREATE PROCEDURE ownerProductStatus(u_token VARCHAR(240), p_id VARCHAR(240))
  BEGIN
    DECLARE isOwner BOOL;
    DECLARE u_id INT;


    SELECT user_id INTO u_id FROM tokens
    WHERE token = u_token;

    SELECT EXISTS(
      SELECT * FROM has
      LEFT JOIN users ON has.users_id = users.id
      LEFT JOIN products ON has.products_id = products.id
      WHERE product_id = p_id AND users_id = u_id) INTO isOwner;

    IF (isOwner) THEN
      CALL checkAuthedProductAvailability(p_id);
    ELSE
      SELECT "Failed";
    END IF;
  END;

DROP PROCEDURE getUserIDofToken;
CALL getUserIDofToken("94a17bfa-6c49-4398-8155-137f07612f7d");

CREATE PROCEDURE getUserIDofToken(u_token VARCHAR(240))
  BEGIN
  SELECT user_id FROM tokens
  WHERE token = u_token;
  END;