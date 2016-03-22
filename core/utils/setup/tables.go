package setup

import (
	"github.com/remony/Equipment-Rental-API/core/config"
	"log"
)

func createUsersTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`users` (`id` int(11) NOT NULL AUTO_INCREMENT, `username` varchar(45) NOT NULL, `password` varchar(250) NOT NULL, `email` varchar(45) NOT NULL,`first_name` varchar(45) DEFAULT 'first_name', `last_name` varchar(45) DEFAULT 'last_name', `location` varchar(45) DEFAULT 'unknown', `bio` varchar(140) DEFAULT 'Please describe me', `date_registered` date NOT NULL, `karma` int(11) NOT NULL DEFAULT '0', `role` varchar(240) NOT NULL, `date_of_birth` datetime NOT NULL, PRIMARY KEY (`id`),KEY `username` (`username`)) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("User Table Created")
}

func createTokenTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`tokens` (`id` int(11) NOT NULL AUTO_INCREMENT,`token` varchar(250) NOT NULL,`date_generated` date NOT NULL,`date_expires` date NOT NULL,`user_id` int(11) NOT NULL,`idenf` varchar(250) NOT NULL,`active` tinyint(1) NOT NULL DEFAULT '1',PRIMARY KEY (`id`),KEY `fk_token_user_idx` (`user_id`),CONSTRAINT `fk_token_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION) ENGINE=InnoDB AUTO_INCREMENT=706 DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("Token Table Created")
}

func createPostsTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`posts` (`id` int(11) NOT NULL AUTO_INCREMENT,`title` varchar(140) DEFAULT NULL,`slug` varchar(140) DEFAULT NULL,`author` varchar(45) DEFAULT NULL,`content` varchar(140) DEFAULT NULL,`date_created` date DEFAULT NULL,`date_edited` date DEFAULT NULL,`users_id` int(11) NOT NULL,PRIMARY KEY (`id`),KEY `fk_posts_users1_idx` (`users_id`),CONSTRAINT `fk_posts_users1` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION) ENGINE=InnoDB DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("Posts Table Created")
}

func createImagesTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`images` (`id` int(11) NOT NULL AUTO_INCREMENT,`file_name` varchar(256) NOT NULL,`title` varchar(256) NOT NULL,`date_added` datetime NOT NULL,`original_name` varchar(256) NOT NULL DEFAULT 'Null',PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=334 DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("Images Table Created")
}

func createProductsTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`products` (`id` int(11) NOT NULL AUTO_INCREMENT,`product_name` varchar(240) NOT NULL,`product_id` varchar(240) NOT NULL,`date_added` datetime NOT NULL,`date_updated` datetime NOT NULL,`product_description` varchar(240) NOT NULL,`product_rental_period_limit` varchar(240) NOT NULL,`ownerid` int(11) NOT NULL,`condition` varchar(240) NOT NULL,`authorized` tinyint(1) NOT NULL DEFAULT '0',`visable` tinyint(1) DEFAULT '1',`enable_comments` tinyint(1) NOT NULL DEFAULT '1',`comments_require_approval` tinyint(1) NOT NULL DEFAULT '0',`content` text NOT NULL,`age_rating` int(11) NOT NULL DEFAULT '0',PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=153 DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("Products Table Created")
}

func createProductHasImagesTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`products_has_images` (`products_id` int(11) NOT NULL,`images_id` int(11) NOT NULL,PRIMARY KEY (`products_id`,`images_id`),KEY `fk_products_has_images_images1_idx` (`images_id`),KEY `fk_products_has_images_products1_idx` (`products_id`),CONSTRAINT `fk_products_has_images_images1` FOREIGN KEY (`images_id`) REFERENCES `images` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,CONSTRAINT `fk_products_has_images_products1` FOREIGN KEY (`products_id`) REFERENCES `products` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION) ENGINE=InnoDB DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("ProductHasImages Table Created")
}

func createCommmentsTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`comments` (`id` int(11) NOT NULL AUTO_INCREMENT,`comment` text NOT NULL,`date_added` datetime NOT NULL,`date_updated` datetime NOT NULL,`author` int(11) NOT NULL,`ident` varchar(240) NOT NULL,`authorized` tinyint(1) NOT NULL DEFAULT '1',`rating` int(11) NOT NULL DEFAULT '3',PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=139 DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("Comements Table Created")
}

func createHasTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`has` (`users_id` int(11) NOT NULL,`products_id` int(11) NOT NULL,`status` int(11) NOT NULL DEFAULT '0',PRIMARY KEY (`users_id`,`products_id`),KEY `fk_users_has_products_products1_idx` (`products_id`),KEY `fk_users_has_products_users1_idx` (`users_id`),CONSTRAINT `fk_users_has_products_products1` FOREIGN KEY (`products_id`) REFERENCES `products` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,CONSTRAINT `fk_users_has_products_users1` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION) ENGINE=InnoDB DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("Has Table Created")
}

func createLikesTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`likes` (`id` int(11) NOT NULL AUTO_INCREMENT,`like` tinyint(1) NOT NULL,`date_added` datetime NOT NULL,PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=134 DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("Likes Table Created")
}

func createProductHasCommentsTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`products_has_comments` (`products_id` int(11) NOT NULL,`comments_id` int(11) NOT NULL,`users_id` int(11) NOT NULL,PRIMARY KEY (`products_id`,`comments_id`,`users_id`),KEY `fk_products_has_comments_comments1_idx` (`comments_id`),KEY `fk_products_has_comments_products1_idx` (`products_id`),KEY `fk_products_has_comments_users1_idx` (`users_id`),CONSTRAINT `fk_products_has_comments_comments1` FOREIGN KEY (`comments_id`) REFERENCES `comments` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,CONSTRAINT `fk_products_has_comments_products1` FOREIGN KEY (`products_id`) REFERENCES `products` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,CONSTRAINT `fk_products_has_comments_users1` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION) ENGINE=InnoDB DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("products_has_comments Table Created")
}

func createProductHasLikesTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`products_has_likes` (`likes_id` int(11) NOT NULL,`products_id` int(11) NOT NULL,`users_id` int(11) NOT NULL,PRIMARY KEY (`likes_id`,`products_id`,`users_id`),KEY `fk_likes_has_products_products1_idx` (`products_id`),KEY `fk_likes_has_products_likes1_idx` (`likes_id`),KEY `fk_likes_has_products_users1_idx` (`users_id`),CONSTRAINT `fk_likes_has_products_likes1` FOREIGN KEY (`likes_id`) REFERENCES `likes` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,CONSTRAINT `fk_likes_has_products_products1` FOREIGN KEY (`products_id`) REFERENCES `products` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,CONSTRAINT `fk_likes_has_products_users1` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION) ENGINE=InnoDB DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("Products_has_likes Table Created")
}

func createProductHasTagsTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`products_has_tags` (`products_id` int(11) NOT NULL,`tags_id` int(11) NOT NULL,PRIMARY KEY (`products_id`,`tags_id`),KEY `fk_products_has_tags_tags1_idx` (`tags_id`),KEY `fk_products_has_tags_products1_idx` (`products_id`),CONSTRAINT `fk_products_has_tags_products1` FOREIGN KEY (`products_id`) REFERENCES `products` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,CONSTRAINT `fk_products_has_tags_tags1` FOREIGN KEY (`tags_id`) REFERENCES `tags` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION) ENGINE=InnoDB DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("Products_has_tags Table Created")
}

func createPushTokensTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`push_tokens` (`id` int(11) NOT NULL AUTO_INCREMENT,`type` text NOT NULL,`reqid` varchar(240) NOT NULL,PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("Push_tokens Table Created")
}
func createSiteTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`Site` (`id` int(11) NOT NULL,`Title` varchar(45) DEFAULT NULL,`Description` varchar(140) DEFAULT NULL,PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("Site Table Created")
}
func createTagsTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`tags` (`id` int(11) NOT NULL AUTO_INCREMENT,`tag` varchar(240) DEFAULT NULL,PRIMARY KEY (`id`),UNIQUE KEY `tag_UNIQUE` (`tag`)) ENGINE=InnoDB AUTO_INCREMENT=233 DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("Tags Table Created")
}
func createUserRentProductTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`user_rent_product` (`users_id` int(11) NOT NULL,`products_id` int(11) NOT NULL,`date_received` datetime NOT NULL,`date_due` datetime NOT NULL,PRIMARY KEY (`users_id`,`products_id`),KEY `fk_users_has_products_products2_idx` (`products_id`),KEY `fk_users_has_products_users2_idx` (`users_id`),CONSTRAINT `fk_users_has_products_products2` FOREIGN KEY (`products_id`) REFERENCES `products` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,CONSTRAINT `fk_users_has_products_users2` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION) ENGINE=InnoDB DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("User_rent_product Table Created")
}
func createUsersHasPushTokensTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`users_has_push_tokens` (`users_id` int(11) NOT NULL,`push_tokens_id` int(11) NOT NULL,PRIMARY KEY (`users_id`,`push_tokens_id`),KEY `fk_users_has_push_tokens_push_tokens1_idx` (`push_tokens_id`),KEY `fk_users_has_push_tokens_users1_idx` (`users_id`),CONSTRAINT `fk_users_has_push_tokens_push_tokens1` FOREIGN KEY (`push_tokens_id`) REFERENCES `push_tokens` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,CONSTRAINT `fk_users_has_push_tokens_users1` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION) ENGINE=InnoDB DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("Users_has_push_tokens Table Created")
}
func createUserRequestsProductTable(db config.Context) {
	stmt, err := db.Session.Prepare("CREATE TABLE IF NOT EXISTS `honoursproject`.`users_requests_products` (`users_id` int(11) NOT NULL,`products_id` int(11) NOT NULL,`date_requested` datetime NOT NULL,PRIMARY KEY (`users_id`,`products_id`),KEY `fk_users_has_products_products3_idx` (`products_id`),KEY `fk_users_has_products_users3_idx` (`users_id`),CONSTRAINT `fk_users_has_products_products3` FOREIGN KEY (`products_id`) REFERENCES `products` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,CONSTRAINT `fk_users_has_products_users3` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION) ENGINE=InnoDB DEFAULT CHARSET=latin1")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("User_requests_Product Table Created")
}
