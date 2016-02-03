USE honoursproject;
SELECT *
FROM users;

SELECT *
FROM tokens;

SELECT *
FROM images;

SELECT COUNT(*)
FROM products;

SELECT *
FROM products
LIMIT 2, 2;

SELECT EXISTS(SELECT 1
              FROM images
              WHERE file_name LIKE '%A2ocy2zhxM%');

SELECT *
FROM images
WHERE products_id = "5";

SELECT *
FROM images
WHERE file_name = "a_lemon.jpg";

INSERT INTO images (file_name, title, date_added, orignal_name) VALUES ();

SELECT EXISTS(SELECT 1
              FROM tokens
              WHERE token = "cb8609cf-daf3-4da0-85e7-908ecddca4af" AND active = 1 AND NOW() <= date_expires
);
SELECT EXISTS(SELECT 1
              FROM tokens
              WHERE token = "4fa49c8e-30a1-495d-b700-181eafc3589c" AND active = 1 AND date_expires <= "2015-12-22");


UPDATE tokens
SET active = 1
WHERE idenf = "8a27272a-a5ab-4703-a244-a0c5a475cfd6";

DROP TABLE IF EXISTS `honoursproject`.`images`;
DROP TABLE IF EXISTS `honoursproject`.`posts`;
DROP TABLE IF EXISTS `honoursproject`.`tokens`;
DROP TABLE IF EXISTS `honoursproject`.`has`;
DROP TABLE IF EXISTS `honoursproject`.`products`;
DROP TABLE IF EXISTS `honoursproject`.`users`;

ALTER TABLE `users` DROP COLUMN Salt;
CREATE TABLE IF NOT EXISTS `honoursproject`.`users` (
  `id`              INT          NOT NULL AUTO_INCREMENT,
  `username`        VARCHAR(45)  NOT NULL,
  `password`        VARCHAR(250) NOT NULL,
  `email`           VARCHAR(45)  NOT NULL,
  `first_name`      VARCHAR(45)  NULL     DEFAULT 'first_name',
  `last_name`       VARCHAR(45)  NULL     DEFAULT 'last_name',
  `location`        VARCHAR(45)  NULL     DEFAULT 'unknown',
  `bio`             VARCHAR(140) NULL     DEFAULT 'Please describe me',
  `date_registered` DATE         NOT NULL,
  `karma`           INT          NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  INDEX `username` (`username` ASC)
)
  ENGINE = InnoDB;


CREATE TABLE IF NOT EXISTS `honoursproject`.`tokens` (
  `id`             INT          NOT NULL AUTO_INCREMENT
  COMMENT '',
  `token`          VARCHAR(250) NOT NULL
  COMMENT '',
  `date_generated` DATE         NOT NULL
  COMMENT '',
  `date_expires`   DATE         NOT NULL
  COMMENT '',
  `user_id`        INT          NOT NULL
  COMMENT '',
  `idenf`          VARCHAR(250) NOT NULL
  COMMENT '',
  `active`         TINYINT(1)   NOT NULL DEFAULT 1
  COMMENT '',
  PRIMARY KEY (`id`)
    COMMENT '',
  INDEX `fk_token_user_idx` (`user_id` ASC)
    COMMENT '',
  CONSTRAINT `fk_token_user`
  FOREIGN KEY (`user_id`)
  REFERENCES `honoursproject`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `honoursproject`.`posts` (
  `id`           INT          NOT NULL AUTO_INCREMENT
  COMMENT '',
  `title`        VARCHAR(140) NULL
  COMMENT '',
  `slug`         VARCHAR(140) NULL
  COMMENT '',
  `author`       VARCHAR(45)  NULL
  COMMENT '',
  `content`      VARCHAR(140) NULL
  COMMENT '',
  `date_created` DATE         NULL
  COMMENT '',
  `date_edited`  DATE         NULL
  COMMENT '',
  `users_id`     INT          NOT NULL
  COMMENT '',
  PRIMARY KEY (`id`)
    COMMENT '',
  INDEX `fk_posts_users1_idx` (`users_id` ASC)
    COMMENT '',
  CONSTRAINT `fk_posts_users1`
  FOREIGN KEY (`users_id`)
  REFERENCES `honoursproject`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `honoursproject`.`images` (
  `id`            INT          NOT NULL AUTO_INCREMENT,
  `file_name`     VARCHAR(256) NOT NULL,
  `title`         VARCHAR(256) NOT NULL,
  `date_added`    DATETIME     NOT NULL,
  `original_name` VARCHAR(256) NOT NULL DEFAULT 'Null',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `honoursproject`.`products` (
  `id`                          INT          NOT NULL AUTO_INCREMENT,
  `product_name`                VARCHAR(240) NOT NULL,
  `product_id`                  VARCHAR(240) NOT NULL,
  `date_added`                  DATETIME     NOT NULL,
  `date_updated`                DATETIME     NOT NULL,
  `product_description`         VARCHAR(240) NOT NULL,
  `product_rental_period_limit` VARCHAR(240) NOT NULL,
  `ownerid`                     INT          NOT NULL,
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `honoursproject`.`products_has_images` (
  `products_id` INT NOT NULL,
  `images_id`   INT NOT NULL,
  PRIMARY KEY (`products_id`, `images_id`),
  INDEX `fk_products_has_images_images1_idx` (`images_id` ASC),
  INDEX `fk_products_has_images_products1_idx` (`products_id` ASC),
  CONSTRAINT `fk_products_has_images_products1`
  FOREIGN KEY (`products_id`)
  REFERENCES `honoursproject`.`products` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_products_has_images_images1`
  FOREIGN KEY (`images_id`)
  REFERENCES `honoursproject`.`images` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `honoursproject`.`has` (
  `users_id`    INT NOT NULL,
  `products_id` INT NOT NULL,
  `status`      INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`users_id`, `products_id`),
  INDEX `fk_users_has_products_products1_idx` (`products_id` ASC),
  INDEX `fk_users_has_products_users1_idx` (`users_id` ASC),
  CONSTRAINT `fk_users_has_products_users1`
  FOREIGN KEY (`users_id`)
  REFERENCES `honoursproject`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_users_has_products_products1`
  FOREIGN KEY (`products_id`)
  REFERENCES `honoursproject`.`products` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB;

DROP TABLE user_rent_product;

CREATE TABLE IF NOT EXISTS `honoursproject`.`user_rent_product` (
  `users_id`      INT      NOT NULL,
  `products_id`   INT      NOT NULL,
  `date_received` DATETIME NOT NULL,
  `date_due`      DATETIME NOT NULL,
  PRIMARY KEY (`users_id`, `products_id`),
  INDEX `fk_users_has_products_products2_idx` (`products_id` ASC),
  INDEX `fk_users_has_products_users2_idx` (`users_id` ASC),
  CONSTRAINT `fk_users_has_products_users2`
  FOREIGN KEY (`users_id`)
  REFERENCES `honoursproject`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_users_has_products_products2`
  FOREIGN KEY (`products_id`)
  REFERENCES `honoursproject`.`products` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `honoursproject`.`users_requests_products` (
  `users_id`       INT         NOT NULL,
  `products_id`    INT         NOT NULL,
  `date_requested` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`users_id`, `products_id`),
  INDEX `fk_users_has_products_products3_idx` (`products_id` ASC),
  INDEX `fk_users_has_products_users3_idx` (`users_id` ASC),
  CONSTRAINT `fk_users_has_products_users3`
  FOREIGN KEY (`users_id`)
  REFERENCES `honoursproject`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_users_has_products_products3`
  FOREIGN KEY (`products_id`)
  REFERENCES `honoursproject`.`products` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB;


CREATE TABLE IF NOT EXISTS `honoursproject`.`tags` (
  `id`  INT          NOT NULL AUTO_INCREMENT,
  `tag` VARCHAR(240) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `tag_UNIQUE` (`tag` ASC)
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `honoursproject`.`products_has_tags` (
  `products_id` INT NOT NULL,
  `tags_id`     INT NOT NULL,
  PRIMARY KEY (`products_id`, `tags_id`),
  INDEX `fk_products_has_tags_tags1_idx` (`tags_id` ASC),
  INDEX `fk_products_has_tags_products1_idx` (`products_id` ASC),
  CONSTRAINT `fk_products_has_tags_products1`
  FOREIGN KEY (`products_id`)
  REFERENCES `honoursproject`.`products` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_products_has_tags_tags1`
  FOREIGN KEY (`tags_id`)
  REFERENCES `honoursproject`.`tags` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `honoursproject`.`Site` (
  `id`          INT          NOT NULL
  COMMENT '',
  `Title`       VARCHAR(45)  NULL
  COMMENT '',
  `Description` VARCHAR(140) NULL
  COMMENT '',
  PRIMARY KEY (`id`)
    COMMENT ''
)
  ENGINE = InnoDB;

#
#
#
#           Procedures
#
#
#


#
#     Register
#

DROP PROCEDURE register;

CALL register("lemon", "test", "lemon@lemondev.xyz", "lemon", "yamano");
CREATE PROCEDURE `register`(u_name      VARCHAR(240), u_password VARCHAR(240), u_email VARCHAR(240),
                            u_firstname VARCHAR(240), u_lastname VARCHAR(240))
  BEGIN
    INSERT INTO users (username, password, email, first_name, last_name, location, date_registered)
    VALUES (u_name, u_password, u_email, u_firstname, u_lastname, "null", NOW());
  END;

SELECT *
FROM users;

#
#   Remove user
#   > Later on we will want to limit this to only admins and the defined user by using there token
#
DROP PROCEDURE removeUser;
CALL removeUser("lemontest");

CREATE PROCEDURE `removeUser`(u_name VARCHAR(240))
  BEGIN
    DECLARE UID INT;
    SELECT id
    INTO UID
    FROM users
    WHERE username = u_name;
    DELETE FROM tokens
    WHERE user_id = UID;
    DELETE FROM users
    WHERE username = u_name;
  END;

#
#   Does User Exist
#

DROP PROCEDURE doesUserExist;

CALL doesUserExist("lemon");

CREATE PROCEDURE `doesUserExist`(u_name VARCHAR(240))
  BEGIN
    SELECT EXISTS(SELECT username
                  FROM users
                  WHERE username = u_name);
  END;

#
#   Get Digest
#
CREATE PROCEDURE `getDigest`(u_name VARCHAR(240))
  BEGIN
    SELECT password
    FROM users
    WHERE username = u_name;
  END;

#
#    Login
#
DROP PROCEDURE login;
CALL login('remon', 'bestToken', 'lookToken');

CREATE PROCEDURE `login`(u_name VARCHAR(240), u_token VARCHAR(240), u_idenf VARCHAR(240))
  BEGIN
    DECLARE userid INT;
    SELECT id
    INTO userid
    FROM users
    WHERE username = u_name;

    INSERT INTO tokens (token, user_id, date_generated, date_expires, idenf, active)
    VALUES (u_token, userid, NOW(), NOW() + INTERVAL 7 DAY, u_idenf, TRUE);

    SELECT
      TRUE                   AS success,
      username,
      md5(email)             AS gravatar,
      u_token                AS token,
      NOW() + INTERVAL 7 DAY AS expiry
    FROM users
    WHERE username = u_name;

  END;

DROP PROCEDURE addImage;

CALL addImage("image", "image", "image", "");
CREATE PROCEDURE addImage(i_name VARCHAR(240), i_title VARCHAR(240), i_original VARCHAR(240), u_token VARCHAR(240))
  BEGIN
    DECLARE userid INT;
    SELECT user_id
    INTO userid
    FROM tokens
    WHERE token = u_token;
    INSERT INTO images (file_name, title, date_added, original_name) VALUES (i_name, i_title, NOW(), i_original);
  END;
#
#   Image exists
#

DROP PROCEDURE imageExists;

CREATE PROCEDURE `imageExists`(code VARCHAR(240))
  BEGIN
    SELECT EXISTS(SELECT 1
                  FROM images
                  WHERE file_name LIKE CONCAT('%', code, '%'));
  END;
SELECT EXISTS(SELECT 1
              FROM images
              WHERE file_name LIKE "%1OxlR3nLip%");
CALL imageExists('1OxlR3nLip');

#
# Create Product
#
SELECT id
FROM images
WHERE file_name = 'EDH86AiKEx.jpg';

DROP PROCEDURE createProduct;
CALL createProduct("item3", "something3", "2015-12-27", "2015-12-27", "something", 7, 0, 16);

CREATE PROCEDURE createProduct(product_name                VARCHAR(240), product_id VARCHAR(240), date_added DATETIME,
                               date_updated                DATETIME, product_description VARCHAR(240),
                               product_rental_period_limit VARCHAR(240), product_image_id VARCHAR(240), owner_id INT)
  BEGIN
    DECLARE imgid INT;
    SELECT id
    INTO imgid
    FROM images
    WHERE file_name = product_image_id
    ORDER BY date_added DESC;
    INSERT INTO products (product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, ownerid)
    VALUES
      (product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, owner_id);
    SET @last_id = LAST_INSERT_ID();
    INSERT INTO has (users_id, products_id, status) VALUES (owner_id, @last_id, 0);
    INSERT INTO products_has_images (products_id, images_id) VALUES (@last_id, imgid);
  END;

#
#  Remove Product
#
DROP PROCEDURE removeProduct;
CALL removeProduct("4ecbc6df-0d66-40dc-ae91-d6d5488b4d7e", "9cc7a542-c22a-4855-8cdf-cc4b5cd4a13a");

CREATE PROCEDURE removeProduct(u_token VARCHAR(240), p_id VARCHAR(240))
  BEGIN
    DECLARE uid INT;
    DECLARE pid INT;
    DECLARE iid INT;
    SELECT user_id
    INTO uid
    FROM tokens
    WHERE token = u_token;
    SELECT id
    INTO pid
    FROM products
    WHERE product_id = p_id;
    SELECT images_id
    INTO iid
    FROM products_has_images
    WHERE products_id = pid;

    #     select iid;
    DELETE FROM has
    WHERE users_id = uid AND products_id = pid;
    DELETE FROM products_has_images
    WHERE products_id = pid;
    DELETE FROM images
    WHERE id = iid;
    DELETE FROM has
    WHERE products_id = pid;
    DELETE FROM user_rent_product
    WHERE products_id = pid;
    DELETE FROM products
    WHERE id = pid;

  END;

#
#   getListingOfTag
#
DROP PROCEDURE getListingOfTag;
CALL getListingOfTag("mobile", 0, 5);

CREATE PROCEDURE `getListingOfTag`(s_tag VARCHAR(240), start INT, count INT)
  BEGIN
    SELECT
      product_id                           AS id,
      products.product_name                AS name,
      products.product_description         AS description,
      products.date_added,
      products.date_updated,
      products.product_rental_period_limit AS time_period,
      has.products_id                      AS image_id,
      username                             AS username,
      md5(email)                           AS gravatar
    FROM has
      LEFT OUTER JOIN products ON has.products_id = products.id
      LEFT OUTER JOIN users ON has.users_id = users.id
      LEFT OUTER JOIN products_has_tags ON products.id = products_has_tags.products_id
      LEFT OUTER JOIN tags ON products_has_tags.tags_id = tags.id
    WHERE tag = s_tag
    ORDER BY products.date_updated DESC
    LIMIT START, COUNT;
  END;

#
#   searchFilterByTag
#

DROP PROCEDURE searchFilterByTag;
CALL searchFilterByTag("mobile", 0, 2);

CREATE PROCEDURE `searchFilterByTag`(s_tag TEXT, start INT, count INT)
  BEGIN
    SELECT
      product_id                           AS id,
      products.product_name                AS name,
      products.product_description         AS description,
      products.date_added,
      products.date_updated,
      products.product_rental_period_limit AS time_period,
      has.products_id                      AS image_id,
      username                             AS username,
      md5(email)                           AS gravatar
    FROM has
      LEFT OUTER JOIN products ON has.products_id = products.id
      LEFT OUTER JOIN users ON has.users_id = users.id
      LEFT OUTER JOIN products_has_tags ON products.id = products_has_tags.products_id
      LEFT OUTER JOIN tags ON products_has_tags.tags_id = tags.id
    WHERE tag LIKE CONCAT("%", s_tag, "%")
    GROUP BY product_id
    ORDER BY products.date_updated DESC
    LIMIT START, COUNT;
  END;

#
#   Get Listing
#
DROP PROCEDURE getListing;
CALL getListing();

CREATE PROCEDURE getListing()
  BEGIN
    SELECT
      username    AS username,
      md5(email)  AS gravatar,
      product_name,
      product_id,
      date_added,
      date_updated,
      product_description,
      product_rental_period_limit,
      products.id AS id
    FROM products
      LEFT JOIN has ON products.id = has.products_id
      LEFT JOIN users ON has.users_id = users.id
    ORDER BY date_updated DESC;
  END;

#
# Get Product
#
DROP PROCEDURE getProduct;
CALL getProduct("65400420-f002-4872-a358-72bcc34f0b30");

CREATE PROCEDURE getProduct(pid VARCHAR(240))
  BEGIN
    DECLARE p_id INT;
    DECLARE tags TEXT;

    SELECT id
    INTO p_id
    FROM products
    WHERE product_id = pid;

    SELECT GROUP_CONCAT(CONCAT(tag) SEPARATOR ', ')
    INTO tags
    FROM products_has_tags
      LEFT JOIN tags ON products_has_tags.tags_id = tags.id
    WHERE products_id = p_id;

    IF (tags IS NULL)
    THEN
      SET tags = "no tags";
    END IF;

    SELECT
      product_name,
      product_id,
      date_added,
      date_updated,
      product_description,
      product_rental_period_limit,
      username,
      products.id AS id,
      tags
    FROM has
      LEFT JOIN users ON has.users_id = users.id
      LEFT JOIN products ON has.products_id = products.id
    WHERE product_id = pid;
  END;

#
# Add Tag to product
#
DROP PROCEDURE addTag;
CALL addTag('fbf8a27c-da48-4300-b611-b427843c835e', 'google');

CREATE PROCEDURE `addTag`(p_id VARCHAR(240), p_tag VARCHAR(240))
  BEGIN
    DECLARE pid INT;
    DECLARE tag_exists INT;
    DECLARE tag_relation_exists BOOL;
    SELECT id
    INTO pid
    FROM products
    WHERE product_id = p_id;
    SELECT id
    INTO tag_exists
    FROM tags
    WHERE tag = p_tag;
    SELECT EXISTS(SELECT *
                  FROM products_has_tags
                  WHERE products_id = pid AND tags_id = tag_exists)
    INTO tag_relation_exists;
    #     SELECT tag_exists;

    IF (tag_exists IS NULL)
    THEN
      INSERT INTO tags (tag) VALUES (p_tag);
      SELECT id
      INTO tag_exists
      FROM tags
      WHERE tag = p_tag;
      IF (tag_relation_exists IS FALSE)
      THEN
        INSERT INTO products_has_tags (products_id, tags_id) VALUES (pid, tag_exists);
      END IF;
    ELSE
      IF (tag_relation_exists IS FALSE)
      THEN
        INSERT INTO products_has_tags (products_id, tags_id) VALUES (pid, tag_exists);
      END IF;
    END IF;
  END;

#
# Remove Tag
#

DROP PROCEDURE removeTag;
CALL removeTag('fbf8a27c-da48-4300-b611-b427843c835e', 'google');

CREATE PROCEDURE `removeTag`(p_id VARCHAR(240), p_tag VARCHAR(240))
  BEGIN
    DECLARE pid INT;
    DECLARE tid INT;
    SELECT id
    INTO pid
    FROM products
    WHERE product_id = p_id;
    SELECT id
    INTO tid
    FROM tags
    WHERE tag = p_tag;

    DELETE FROM products_has_tags
    WHERE tags_id = tid AND products_id = pid;
  END;

#
#  Get Tags
#
DROP PROCEDURE GetTags;
CALL GetTags("3c0bf3ad-4aae-4a3b-b2c8-5a6df691a3e7");

CREATE PROCEDURE `GetTags`(pid VARCHAR(240))
  BEGIN
    DECLARE p_id INT;

    SELECT id
    INTO p_id
    FROM products
    WHERE product_id = pid;

    SELECT tags.tag
    FROM products_has_tags
      LEFT JOIN tags ON products_has_tags.tags_id = tags.id
    WHERE products_id = p_id;
  END;

#
#  Get Image
#
CALL getImage(6);

CREATE PROCEDURE getImage(pid INT)
  BEGIN
    SELECT
      file_name,
      title,
      date_added
    FROM products_has_images
      LEFT JOIN images ON products_has_images.images_id = images.id
    WHERE products_id = pid;
  END;

#
#   RequestToBorrowItem
#

DROP PROCEDURE RequestToBorrowItem;
CALL RequestToBorrowItem("374c1e1a-12e8-49f0-8432-938f1594e2d4", "956e8553-b25c-45d2-896c-268ab0e7bacc");

CREATE PROCEDURE RequestToBorrowItem(u_pid VARCHAR(240), u_token VARCHAR(240))
  BEGIN
    DECLARE uid INT;
    DECLARE pid INT;

    SELECT user_id
    INTO uid
    FROM tokens
    WHERE token = u_token;

    SELECT id
    INTO pid
    FROM products
    WHERE product_id = u_pid;

    IF (SELECT EXISTS(SELECT *
                      FROM users_requests_products
                      WHERE users_id = uid AND products_id = pid))
    THEN
      SELECT
        FALSE,
        "null",
        NOW();
    ELSE
      INSERT INTO users_requests_products (products_id, users_id, date_requested) VALUES (pid, uid, NOW());
      SELECT
        TRUE,
        u_pid,
        NOW();
    END IF;


  END;

#
#  Get requests of product
#

DROP PROCEDURE OwnerGetProductRequests;
CALL OwnerGetProductRequests("1640049c-3930-4283-bde8-fb655ea70a5a", "374c1e1a-12e8-49f0-8432-938f1594e2d4");

CREATE PROCEDURE `OwnerGetProductRequests`(u_token VARCHAR(240), u_pid VARCHAR(240))
  BEGIN
    DECLARE uid INT;
    DECLARE pid INT;

    SELECT user_id
    INTO uid
    FROM tokens
    WHERE token = u_token;

    SELECT id
    INTO pid
    FROM products
    WHERE product_id = u_pid;

    SELECT username, md5(email) as gravatar, date_requested from users_requests_products
      LEFT JOIN users ON users_requests_products.users_id = users.id
      where products_id = pid
    ORDER BY date_requested DESC;
  END;


#
# Owner get items with requests
#

DROP PROCEDURE OwnerGetRequests;
CALL OwnerGetRequests("1640049c-3930-4283-bde8-fb655ea70a5a", 0, 5);

CREATE PROCEDURE `OwnerGetRequests`(u_token VARCHAR(240), step int, count int)
  BEGIN
    DECLARE uid int;
    DECLARE requests int;
    select user_id into uid from tokens where token = u_token;


#     select uid;
    SELECT
      username    AS username,
      md5(email)  AS gravatar,
      product_name,
      product_id,
      date_added,
      date_updated,
      product_description,
      product_rental_period_limit,
      products.id AS id,
      COUNT(date_requested) as requests
    FROM users_requests_products
      LEFT JOIN products ON users_requests_products.products_id = products.id
      LEFT JOIN has ON products.id = has.products_id
      LEFT JOIN users ON has.users_id = users.id
      WHERE ownerid = uid
      GROUP BY product_id

    ORDER BY date_updated DESC
    LIMIT step, count;
  END;

#
#  Owner Cancel Request
#

DROP PROCEDURE OwnerDropRequest;
CALL OwnerDropRequest();

CREATE PROCEDURE `OwnerDropRequest`(u_pid VARCHAR(240) , t_username VARCHAR(240))
  BEGIN
    DECLARE uid INT;
    DECLARE pid INT;

    SELECT id
    INTO uid
    FROM users
    WHERE username = t_username;

    SELECT id
    INTO pid
    FROM products
    WHERE product_id = u_pid;

    DELETE FROM users_requests_products
    WHERE products_id = pid AND users_id = uid;
  END;

#
#   Cancel Request
#

DROP PROCEDURE CancelRequest;
CALL CancelRequest("374c1e1a-12e8-49f0-8432-938f1594e2d4", "956e8553-b25c-45d2-896c-268ab0e7bacc");

CREATE PROCEDURE `CancelRequest`(p_id VARCHAR(240), u_token VARCHAR(240))
  BEGIN
    DECLARE uid INT;
    DECLARE pid INT;

    SELECT user_id
    INTO uid
    FROM tokens
    WHERE token = u_token;

    SELECT id
    INTO pid
    FROM products
    WHERE product_id = p_id;

    DELETE FROM users_requests_products
    WHERE products_id = pid AND users_id = uid;

  END;

#
#  GetRequestStatus
#

DROP PROCEDURE GetRequestStatus;
CALL GetRequestStatus("fbf749c8-010f-4bb1-aa10-7da3aca6ba0d", "674c99c7-da73-43f3-b8fe-1e6c96eedda7");

CREATE PROCEDURE GetRequestStatus(p_id VARCHAR(240), u_token VARCHAR(240))
  BEGIN
    DECLARE uid INT;
    DECLARE pid INT;
    DECLARE pname VARCHAR(240);
    DECLARE d_requested DATETIME;
    DECLARE numofreq int;
    DECLARE owner bool;
    SELECT user_id
    INTO uid
    FROM tokens
    WHERE token = u_token;

    SELECT id
    INTO pid
    FROM products
    WHERE product_id = p_id;

    SELECT
      product_name,
      date_requested
    INTO pname, d_requested
    FROM users_requests_products
      LEFT JOIN users ON users_requests_products.users_id = users.id
      LEFT JOIN products ON users_requests_products.products_id = products.id
    WHERE products_id = pid AND users_id = uid;

    SELECT EXISTS(
               SELECT *
               FROM has
                 LEFT JOIN users ON has.users_id = users.id
                 LEFT JOIN products ON has.products_id = products.id
               WHERE product_id = p_id AND users_id = uid) INTO owner;

    if (owner IS TRUE) THEN
      select count(*) into numofreq from users_requests_products where products_id = pid;
      ELSE
       SET numofreq = 0;
    END IF;



    IF (pname IS NULL)
    THEN
      SELECT
        FALSE  AS requested,
        "null" AS product_title,
        NOW()  AS date_requested,
        numofreq as requests;
    ELSE
      SELECT
        TRUE        AS requested,
        pname       AS product_title,
        d_requested AS date_requested,
        numofreq as requests;
    END IF;

  END;

#
#  RentFromRequest
#

DROP PROCEDURE RentFromRequest;
CALL RentFromRequest("fbf8a27c-da48-4300-b611-b427843c835e", "john");

CREATE PROCEDURE RentFromRequest(u_pid VARCHAR(240), usrname VARCHAR(240))
  BEGIN
    DECLARE userid INT;
    DECLARE days INT;
    DECLARE pid INT;

    SELECT id
    INTO userid
    FROM users
    WHERE username = usrname;
    SELECT
      id,
      product_rental_period_limit
    INTO pid, days
    FROM products
    WHERE product_id = u_pid;
    #     SELECT days;
    DELETE FROM users_requests_products
    WHERE users_id = userid AND products_id = pid;

    INSERT INTO user_rent_product (products_id, users_id, date_received, date_due)
    VALUES (pid, userid, NOW(), DATE_ADD(NOW(), INTERVAL days DAY));
  END;

#
# Rent Item
#

DROP PROCEDURE RentItem;
CALL RentItem("1cea1430-3a53-4e7a-9834-c52137ab8b5e", "remon");

CREATE PROCEDURE RentItem(u_pid VARCHAR(240), usrname VARCHAR(240))
  BEGIN
    DECLARE userid INT;
    DECLARE days INT;
    DECLARE pid INT;

    SELECT id
    INTO userid
    FROM users
    WHERE username = usrname;
    SELECT
      id,
      product_rental_period_limit
    INTO pid, days
    FROM products
    WHERE product_id = u_pid;
    #     SELECT days;
    INSERT INTO user_rent_product (products_id, users_id, date_received, date_due)
    VALUES (pid, userid, NOW(), DATE_ADD(NOW(), INTERVAL days DAY));
  END;

#
#   Return Item
#

DROP PROCEDURE ReturnItem;

CALL ReturnItem("14afe718-3b4d-4193-a8a0-d8401f9a4a01", "1cea1430-3a53-4e7a-9834-c52137ab8b5e");

CREATE PROCEDURE ReturnItem(o_token VARCHAR(240), product VARCHAR(240))
  BEGIN
    DECLARE productid INT;
    DECLARE u_id VARCHAR(240);
    DECLARE tmp_u_id VARCHAR(240);

    SELECT user_id
    INTO tmp_u_id
    FROM tokens
    WHERE token = o_token;

    SELECT users_id
    INTO u_id
    FROM user_rent_product
      LEFT JOIN products ON user_rent_product.products_id = products.id
      LEFT JOIN users ON ownerid = users.id
    WHERE users_id = tmp_u_id AND product_id = product;

    SELECT id
    INTO productid
    FROM products
    WHERE product_id = product;
    DELETE FROM user_rent_product
    WHERE users_id = u_id AND products_id = productid;
  END;

#
# Return Item As Owner
#

DROP PROCEDURE ReturnItemAsOwner;
CALL ReturnItemAsOwner("33c49783-f059-42d5-b5c2-6234eb8a5b78", "f4025ccb-3656-4975-a9ee-0bba89e085db");

CREATE PROCEDURE ReturnItemAsOwner(o_token VARCHAR(240), product VARCHAR(240))
  BEGIN
    DECLARE pid INT;
    DECLARE uid INT;
    DECLARE tmp_u_id INT;

    SELECT user_id
    INTO tmp_u_id
    FROM tokens
    WHERE token = o_token;

    SELECT
      products_id,
      users_id
    INTO pid, uid
    FROM user_rent_product
      LEFT JOIN products ON user_rent_product.products_id = products.id
    WHERE product_id = product AND ownerid = tmp_u_id;
    DELETE FROM user_rent_product
    WHERE users_id = uid AND products_id = pid;
  END;

#
# Check Item Availabibility
#

DROP PROCEDURE checkItemAvailability;
CALL checkItemAvailability("1cea1430-3a53-4e7a-9834-c52137ab8b5e", "remon");

CREATE PROCEDURE `checkItemAvailability`(product VARCHAR(240), usrname VARCHAR(240))
  BEGIN
    DECLARE due_date DATETIME;

    SELECT date_due
    INTO due_date
    FROM user_rent_product
      LEFT OUTER JOIN products ON user_rent_product.products_id = products.id
      LEFT OUTER JOIN users ON user_rent_product.users_id = users.id
    WHERE products.product_id = product
    ORDER BY products.date_updated DESC;

    IF (due_date > NOW())
    THEN
      SELECT
        FALSE,
        due_date;
    ELSE
      SELECT
        TRUE AS Available,
        NOW();
    END IF;
  END;

#
# Get Paged Products
#

DROP PROCEDURE getPagedProducts;
CALL getPagedProducts(0, 6);

CREATE PROCEDURE getPagedProducts(step INT, count INT)
  BEGIN
    SELECT
      product_id                  AS id,
      product_name                AS name,
      product_description         AS description,
      date_added,
      date_updated,
      product_rental_period_limit AS time_period,
      products_id                 AS image_id,
      username                    AS username,
      md5(email)                  AS gravatar
    FROM has
      LEFT OUTER JOIN products ON has.products_id = products.id
      LEFT OUTER JOIN users ON has.users_id = users.id
    ORDER BY products.date_updated DESC
    LIMIT step, COUNT;
  END;

# Get Rented Products

DROP PROCEDURE getRentedProducts;
CALL getRentedProducts("lemon", 0, 3);

CREATE PROCEDURE getRentedProducts(username VARCHAR(240), step INT, count INT)
  BEGIN
    SELECT
      product_id                  AS id,
      product_name                AS name,
      product_description         AS description,
      date_added,
      date_updated,
      product_rental_period_limit AS time_period,
      product_image_id            AS image_id,
      username                    AS owner
    FROM user_rent_product
      LEFT OUTER JOIN products ON user_rent_product.products_id = products.id
    WHERE user_rent_product.date_due < NOW()
    ORDER BY products.date_updated DESC
    LIMIT step, COUNT;
  END;

#
# Get Currently Renting Products
#

DROP PROCEDURE getCurrentlyRentingProducts;
CALL getCurrentlyRentingProducts("remon", 0, 3);

CREATE PROCEDURE getCurrentlyRentingProducts(u_name VARCHAR(240), step INT, count INT)
  BEGIN
    SELECT
      product_id          AS id,
      product_name        AS name,
      product_description AS description,
      date_due            AS due_date,
      date_received       AS received_date,
      products_id         AS image_id,
      username            AS owner
    FROM user_rent_product
      LEFT OUTER JOIN products ON user_rent_product.products_id = products.id
      LEFT OUTER JOIN users ON user_rent_product.users_id = users.id
    WHERE user_rent_product.date_due > NOW() AND username = u_name
    ORDER BY products.date_updated DESC
    LIMIT step, COUNT;
  END;

#
#  Get Username
#

DROP PROCEDURE getUsername;
CALL getUsername("94a17bfa-6c49-4398-8155-137f07612f7d");

CREATE PROCEDURE getUsername(usrtoken VARCHAR(240))
  BEGIN
    DECLARE userid INT;
    SELECT user_id
    INTO userid
    FROM tokens
    WHERE token = usrtoken;
    SELECT username
    FROM users
    WHERE id = userid;
  END;

# Get Product Availability

DROP PROCEDURE checkProductAvailability;
CALL checkProductAvailability("4c0bc251-bc9a-4a95-9612-a883bc6877ad");

CREATE PROCEDURE `checkProductAvailability`(product VARCHAR(240))
  BEGIN
    DECLARE due_date DATETIME;
    DECLARE active_state BOOLEAN;

    SELECT date_due
    INTO due_date
    FROM user_rent_product
      LEFT OUTER JOIN products ON user_rent_product.products_id = products.id
      LEFT OUTER JOIN users ON user_rent_product.users_id = users.id
    WHERE products.product_id = product
    ORDER BY products.date_updated DESC;


    IF (due_date > NOW())
    THEN
      SELECT
        FALSE    AS available,
        due_date AS due_date;
    ELSE
      SELECT
        TRUE  AS available,
        NOW() AS due_date;
    END IF;

  END;

#
#  Check Authed Product Availability
#

DROP PROCEDURE checkAuthedProductAvailability;
CALL checkAuthedProductAvailability("4c0bc251-bc9a-4a95-9612-a883bc6877ad");

CREATE PROCEDURE `checkAuthedProductAvailability`(product VARCHAR(240))
  BEGIN
    DECLARE due_date DATETIME;
    DECLARE taken_date DATETIME;
    DECLARE user_name VARCHAR(240);

    SELECT
      date_due,
      date_received,
      username
    INTO due_date, taken_date, user_name
    FROM user_rent_product
      LEFT JOIN products ON user_rent_product.products_id = products.id
      LEFT JOIN users ON user_rent_product.users_id = users.id
    WHERE products.product_id = product
    ORDER BY date_received DESC
    LIMIT 1;

    #   select due_date;
    IF (user_name != "")
    THEN
      IF (due_date > NOW())
      THEN
        SELECT
          FALSE     AS available,
          due_date,
          taken_date,
          user_name AS username;
      ELSE
        SELECT
          TRUE      AS available,
          NOW()     AS due_date,
          NOW()     AS taken_date,
          user_name AS username;
      END IF;
    ELSE
      SELECT
        TRUE AS available,
        NOW(),
        NOW(),
        "null";
    END IF;
  END;

#
# Get owner products
#

DROP PROCEDURE getOwnerProducts;
CALL getOwnerProducts("94a17bfa-6c49-4398-8155-137f07612f7d", 0, 15);

CREATE PROCEDURE getOwnerProducts(u_token VARCHAR(240), step INT, count INT)
  BEGIN
    DECLARE usrname VARCHAR(240);
    SELECT username
    INTO usrname
    FROM tokens
      LEFT OUTER JOIN users ON tokens.user_id = users.id
    WHERE token = u_token;
    SELECT
      product_id                  AS id,
      product_name                AS name,
      product_description         AS description,
      date_added,
      date_updated,
      product_rental_period_limit AS time_period,
      products_id                 AS image_id,
      username                    AS owner
    FROM has
      LEFT OUTER JOIN products ON has.products_id = products.id
      LEFT OUTER JOIN users ON has.users_id = users.id
    WHERE users.username = usrname
    ORDER BY date_updated DESC
    LIMIT step, COUNT;
  END;

#
#  Check Product Availability Owner
#

DROP PROCEDURE CheckProductAvailabilityOwner;
CALL CheckProductAvailabilityOwner("94a17bfa-6c49-4398-8155-137f07612f7d", "4ded9e43-174f-4203-a48b-58f34dc9b90b");

CREATE PROCEDURE CheckProductAvailabilityOwner(o_token VARCHAR(240), p_id VARCHAR(240))
  BEGIN
    DECLARE due_date DATETIME;
    DECLARE active_state BOOLEAN;

    SELECT
      date_due,
      active
    INTO due_date, active_state
    FROM user_rent_product
      LEFT OUTER JOIN products ON user_rent_product.products_id = products.id
      LEFT OUTER JOIN users ON user_rent_product.users_id = users.id
    WHERE products.product_id = p_id
    ORDER BY products.date_updated DESC;

    IF (active_state = 1)
    THEN
      IF (due_date > NOW())
      THEN
        SELECT
          FALSE    AS available,
          due_date AS due_date;
      ELSE
        SELECT
          TRUE  AS available,
          NOW() AS due_date;
      END IF;
    ELSE
      SELECT
        TRUE  AS available,
        NOW() AS due_date;
    END IF;
  END;

#
#  is Owner
#

select * from has
left join users ON has.users_id = users.id
  left join products on has.products_id = products.id
where product_id = "fbf749c8-010f-4bb1-aa10-7da3aca6ba0d";

DROP PROCEDURE isOwner;
CALL isOwner("f05a1fc2-2c05-4e88-92d1-b9a3a71b8dd2", "fbf749c8-010f-4bb1-aa10-7da3aca6ba0d");

CREATE PROCEDURE isOwner(u_token VARCHAR(240), p_id VARCHAR(240))
  BEGIN
    DECLARE u_id INT;
    SELECT user_id
    INTO u_id
    FROM tokens
    WHERE token = u_token;

    SELECT EXISTS(
               SELECT *
               FROM has
                 LEFT JOIN users ON has.users_id = users.id
                 LEFT JOIN products ON has.products_id = products.id
               WHERE product_id = p_id AND users_id = u_id) AS owner;
  END;

#
# Owner Product Status
#

DROP PROCEDURE ownerProductStatus;
CALL ownerProductStatus("899a7b0b-2e51-4488-a99a-b51b0e76f856", "4ded9e43-174f-4203-a48b-58f34dc9b90b");

CREATE PROCEDURE ownerProductStatus(u_token VARCHAR(240), p_id VARCHAR(240))
  BEGIN
    DECLARE isOwner BOOL;
    DECLARE u_id INT;


    SELECT user_id
    INTO u_id
    FROM tokens
    WHERE token = u_token;

    SELECT EXISTS(
        SELECT *
        FROM has
          LEFT JOIN users ON has.users_id = users.id
          LEFT JOIN products ON has.products_id = products.id
        WHERE product_id = p_id AND users_id = u_id)
    INTO isOwner;

    IF (isOwner)
    THEN
      CALL checkAuthedProductAvailability(p_id);
    ELSE
      CALL checkProductAvailability(p_id);
    END IF;
  END;

#
# get User ID of Token
#

DROP PROCEDURE getUserIDofToken;
CALL getUserIDofToken("94a17bfa-6c49-4398-8155-137f07612f7d");

CREATE PROCEDURE getUserIDofToken(u_token VARCHAR(240))
  BEGIN
    SELECT user_id
    FROM tokens
    WHERE token = u_token;
  END;

#
#   Get Index
#      Get the site index information
#

DROP PROCEDURE getIndex;

CREATE PROCEDURE getIndex()
  BEGIN
    SELECT
      Title       AS title,
      description AS description
    FROM Site
    WHERE id = 1;
  END;

DROP PROCEDURE updateSite;
CALL updateSite("lemon rental", "test");

#
#   Update Site
#   Update the meta data of the website
#

CREATE PROCEDURE updateSite(s_title VARCHAR(240), s_description VARCHAR(240))
  BEGIN
    UPDATE Site
    SET Title = s_title, Description = s_description
    WHERE id = 0;
  END;