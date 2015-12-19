use honoursproject;
select * from users;

select * from tokens;

select * from images;

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
  `product_rented_to` VARCHAR(240) NOT NULL COMMENT '',
  `users_id` INT NOT NULL COMMENT '',
  `owner_id` INT NULL COMMENT '',
  PRIMARY KEY (`id`, `users_id`)  COMMENT '',
  INDEX `fk_products_users1_idx` (`users_id` ASC)  COMMENT '',
  CONSTRAINT `fk_products_users1`
    FOREIGN KEY (`users_id`)
    REFERENCES `honoursproject`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


select * from images;

select * from images where file_name = "15b2459a-705a-4831-92fb-d69cba3ed3eb.gif";


select * from images ORDER BY date_added DESC;

DROP PROCEDURE imageExists;



# Create a procedure which checks if an image exists or not
 CREATE PROCEDURE `imageExists`(code VARCHAR(240))
    BEGIN
      SELECT EXISTS (SELECT 1 FROM images WHERE file_name LIKE CONCAT('%', code , '%'));
    END;




SELECT EXISTS (SELECT 1 FROM images WHERE file_name LIKE "%qKLgSgpS4m%");


CALL imageExists('LT7yTmKe9g');