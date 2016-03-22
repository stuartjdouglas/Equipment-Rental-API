package setup

import (
	"github.com/remony/Equipment-Rental-API/core/config"
	"log"
)

func createLikeProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `like`(u_token VARCHAR(240), p_id VARCHAR(240))  BEGIN  DECLARE pid INT;  DECLARE uid INT;  select user_id into uid from tokens where token = u_token;  select id into pid from products where product_id = p_id;  INSERT INTO likes(`like`, date_added) VALUES(1, NOW());  INSERT INTO products_has_likes(products_id, users_id, likes_id) VALUES(pid, uid, LAST_INSERT_ID());  END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("Like Procedure Created")
}

func createUnLikeProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `unLike`(u_token VARCHAR(240), p_id VARCHAR(240))  BEGIN    DECLARE pid INT;    DECLARE uid INT;    DECLARE lid INT;    select user_id into uid from tokens where token = u_token;    select id into pid from products where product_id = p_id;    select likes_id into lid from products_has_likes where users_id = uid and products_id = pid;    delete from products_has_likes where likes_id = lid and products_id = pid and users_id = uid;    delete from likes where id = lid;  END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("Unlike Procedure Created")
}

func createGetLikesProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `getLikes`(p_id VARCHAR(240), u_token VARCHAR(240))  BEGIN    DECLARE pid INT;    DECLARE ilike BOOl;    DECLARE uid INT;    select id into pid from products where product_id = p_id;    select user_id into uid from tokens where token = u_token;    select exists(select * from products_has_likes where products_id = pid and users_id = uid) into ilike;    select SUM(`like`) as likes, ilike as liked  from products_has_likes    left join likes on products_has_likes.likes_id = likes.id      LEFT JOIN products on products_has_likes.products_id = products.id    where products_id = pid    GROUP BY `like`;  END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("Get Likes Procedure Created")
}

func createAddCommentProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `AddComment`(u_token VARCHAR(240), p_id VARCHAR(240), u_comment VARCHAR(140), requiresApproval BOOL, u_rating int)  BEGIN    DECLARE uid INT;    DECLARE pid INT;    DECLARE requires_authoriz BOOL;    SELECT user_id into uid FROM tokens where token = u_token;    select id, comments_require_approval into pid, requires_authoriz from products where product_id = p_id;    if (requires_authoriz OR requiresApproval) THEN      insert into comments(comment, date_added, date_updated, author, ident, authorized, rating) VALUES(u_comment, NOW(), NOW(), uid, UUID(), FALSE, u_rating);    ELSE      insert into comments(comment, date_added, date_updated, author, ident, rating) VALUES(u_comment, NOW(), NOW(), uid, UUID(), u_rating);    END IF;    insert into products_has_comments(products_id, comments_id, users_id) VALUES(pid, LAST_INSERT_ID(), uid);    select ident from products_has_comments      left join comments on products_has_comments.comments_id = comments.id    where products_has_comments.comments_id = LAST_INSERT_ID();  END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("AddComment Procedure Created")
}

func createEditCommentProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `EditComment`(u_token VARCHAR(240), c_id VARCHAR(240), u_comment VARCHAR(140), u_rating int)  BEGIN    DECLARE uid INT;    DECLARE requires_authoriz BOOL;    SELECT user_id into uid FROM tokens where token = u_token;    UPDATE comments SET comment = u_comment, rating = u_rating where ident = c_id and author = uid;    select ident from products_has_comments      left join comments on products_has_comments.comments_id = comments.id    where ident = c_id;  END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("EditComment Procedure Created")
}

func createHasUserReviewedListingProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `HasUserReviewedListing`(u_token VARCHAR(240), p_id VARCHAR(240))  BEGIN    DECLARE uid INT;    DECLARE pid INT;    SELECT user_id into uid FROM tokens where token = u_token;    select id into pid from products where product_id = p_id;    SELECT EXISTS(select comment from products_has_comments              left join comments on products_has_comments.comments_id = comments.id              left join products on products_has_comments.products_id = products.id           WHERE products_id = pid and users_id = users_id);  END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("HasUserReviewedListing Procedure Created")
}

func createDisableCommentsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `DisableComments`(p_id VARCHAR(240))  BEGIN    UPDATE products SET enable_comments = false WHERE product_id = p_id;  END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("DisableComments Procedure Created")
}

func createEnableCommentsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `EnableComments`(p_id VARCHAR(240))  BEGIN    UPDATE products SET enable_comments = true WHERE product_id = p_id;  END")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("EnableComments Procedure Created")
}

func DisableCommentsRequireAuthProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `DisableCommentsRequireAuth`(p_id VARCHAR(240)) BEGIN UPDATE products SET comments_require_approval = false WHERE product_id = p_id; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("DisableCommentsRequireAuth Procedure Created")
}

func EnableCommentsRequireAuthProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `EnableCommentsRequireAuth`(p_id VARCHAR(240)) BEGIN UPDATE products SET comments_require_approval = true WHERE product_id = p_id; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("EnableCommentsRequireAuth Procedure Created")
}
func ApproveCommentProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `ApproveComment`(c_id VARCHAR(240)) BEGIN DECLARE cid INT; select id into cid from comments where ident = c_id; UPDATE comments SET authorized = true where id = cid; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("ApproveComment Procedure Created")
}

func GetCommentsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `GetComments`(p_id VARCHAR(240)) BEGIN DECLARE pid INT; DECLARE enabled BOOL; select id , enable_comments into pid, enabled from products where product_id = p_id; select comment, username, md5(email) as gravatar, comments.date_added, `comments`.date_updated, ident as indentifier, comments.authorized, rating from products_has_comments left JOIN comments on products_has_comments.comments_id = comments.id LEFT JOIN users on products_has_comments.users_id = users.id left join products on products_has_comments.products_id = products.id WHERE products_id = pid and comments.authorized = true and products.enable_comments = true ORDER BY date_added ASC; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("GetComments Procedure Created")
}
func GetCommentProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `GetComment`(c_id VARCHAR(240)) BEGIN select comment, username, md5(email) as gravatar, comments.date_added, `comments`.date_updated, ident as indentifier, comments.authorized, rating from products_has_comments left JOIN comments on products_has_comments.comments_id = comments.id LEFT JOIN users on products_has_comments.users_id = users.id WHERE ident = c_id; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("GetComment Procedure Created")
}
func GetOwnerCommentsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `GetOwnerComments`(p_id VARCHAR(240)) BEGIN DECLARE pid INT; DECLARE enabled BOOL; select id , enable_comments into pid, enabled from products where product_id = p_id; select comment, username, md5(email) as gravatar, comments.date_added, `comments`.date_updated, ident as indentifier, comments.authorized, rating from products_has_comments left JOIN comments on products_has_comments.comments_id = comments.id LEFT JOIN users on products_has_comments.users_id = users.id left join products on products_has_comments.products_id = products.id WHERE products_id = pid and products.enable_comments = true ORDER BY date_added ASC; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("GetOwnerComments Procedure Created")
}
func DeleteCommentProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `DeleteComment`(u_token VARCHAR(240), comment_id VARCHAR(240)) BEGIN DECLARE cid INT; DECLARE uid INT; select user_id into uid from tokens where token = u_token; select id into cid from comments where ident = comment_id; delete from products_has_comments where users_id = uid and comments_id = cid; delete from comments where ident = comment_id; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("DeleteComment Procedure Created")
}
func GetPushNotificationIDsOfUserProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `GetPushNotificationIDsOfUser`(u_name VARCHAR(240)) BEGIN DECLARE uid INT; SELECT id into uid from users where username = u_name; SELECT IFNULL(username,\" \"), IFNULL(GROUP_CONCAT(CONCAT(reqid) SEPARATOR ', '),\" \") as reqid, IFNULL(type,\" \") FROM users_has_push_tokens LEFT JOIN push_tokens ON users_has_push_tokens.push_tokens_id = push_tokens.id LEFT JOIN users ON users_has_push_tokens.users_id = users.id WHERE username = u_name; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("GetPushNotificationIDsOfUser Procedure Created")
}
func GetPushNotificationIDsOfProductProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `GetPushNotificationIDsOfProduct`(p_id VARCHAR(240)) BEGIN DECLARE uid INT; DECLARE pid INT; SELECT users.id, products.id into uid, pid FROM user_rent_product LEFT JOIN users ON user_rent_product.users_id = users.id LEFT JOIN products ON user_rent_product.products_id = products.id WHERE product_id = p_id; SELECT username, GROUP_CONCAT(CONCAT(reqid) SEPARATOR ', ') as reqid, type FROM users_has_push_tokens LEFT JOIN push_tokens ON users_has_push_tokens.push_tokens_id = push_tokens.id LEFT JOIN users ON users_has_push_tokens.users_id = users.id WHERE users_id = uid; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("GetPushNotificationIDsOfProduct Procedure Created")
}
func addPushNotificationRegIDProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `addPushNotificationRegID`(u_token VARCHAR(240), p_regid VARCHAR(240), p_type TEXT) BEGIN DECLARE uid INT; DECLARE pid INT; DECLARE r_exists BOOL; SELECT EXISTS(select id from push_tokens where reqid = p_regid) into r_exists; if (r_exists) THEN SELECT FALSE; ELSE SELECT user_id into uid FROM tokens where token = u_token; INSERT INTO push_tokens (type, reqid) VALUES(p_type, p_regid); INSERT INTO users_has_push_tokens (users_id, push_tokens_id) VALUES(uid, LAST_INSERT_ID()); SELECT TRUE; END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("addPushNotificationRegID Procedure Created")
}
func registerProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `register`(u_name VARCHAR(240), u_password VARCHAR(240), u_email VARCHAR(240), u_firstname VARCHAR(240), u_lastname VARCHAR(240), p_date_of_birth DATETIME) BEGIN INSERT INTO users (username, password, email, first_name, last_name, location, date_registered, role, date_of_birth) VALUES (u_name, u_password, u_email, u_firstname, u_lastname, \"null\", NOW(), \"user\", p_date_of_birth); END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("register Procedure Created")
}
func removeUserProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `removeUser`(u_name VARCHAR(240)) BEGIN DECLARE UID INT; SELECT id INTO UID FROM users WHERE username = u_name; DELETE FROM tokens WHERE user_id = UID; DELETE FROM users WHERE username = u_name; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("removeUser Procedure Created")
}
func doesUserExistProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `doesUserExist`(u_name VARCHAR(240)) BEGIN SELECT EXISTS(SELECT username FROM users WHERE username = u_name); END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("doesUserExist Procedure Created")
}
func getDigestProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `getDigest`(u_name VARCHAR(240)) BEGIN SELECT password FROM users WHERE username = u_name; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getDigest Procedure Created")
}
func loginProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `login`(u_name VARCHAR(240), u_token VARCHAR(240), u_idenf VARCHAR(240)) BEGIN DECLARE userid INT; SELECT id INTO userid FROM users WHERE username = u_name; INSERT INTO tokens (token, user_id, date_generated, date_expires, idenf, active) VALUES (u_token, userid, NOW(), NOW() + INTERVAL 7 DAY, u_idenf, TRUE); SELECT TRUE AS success, username, md5(email) AS gravatar, u_token AS token, NOW() + INTERVAL 7 DAY AS expiry, role FROM users WHERE username = u_name; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("login Procedure Created")
}

func addImageProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE addImage(i_name VARCHAR(240), i_title VARCHAR(240), i_original VARCHAR(240), u_token VARCHAR(240)) BEGIN DECLARE userid INT; SELECT user_id INTO userid FROM tokens WHERE token = u_token; INSERT INTO images (file_name, title, date_added, original_name) VALUES (i_name, i_title, NOW(), i_original); END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("addImage Procedure Created")
}
func AddAnotherImageProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE AddAnotherImage(i_name VARCHAR(240), i_title VARCHAR(240), i_original VARCHAR(240), u_token VARCHAR(240), p_id VARCHAR(240)) BEGIN DECLARE userid INT; DECLARE pid INT; select id into pid from products where product_id = p_id; SELECT user_id INTO userid FROM tokens WHERE token = u_token; INSERT INTO images (file_name, title, date_added, original_name) VALUES (i_name, i_title, NOW(), i_original); INSERT INTO products_has_images(products_id, images_id) VALUES (pid, LAST_INSERT_ID()); END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("AddAnotherImage Procedure Created")
}
func imageExistsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `imageExists`(code VARCHAR(240)) BEGIN SELECT EXISTS(SELECT 1 FROM images WHERE file_name LIKE CONCAT('%', code, '%')); END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("imageExists Procedure Created")
}
func createProductProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE createProduct(product_name VARCHAR(240), product_id VARCHAR(240), date_added DATETIME, date_updated DATETIME, product_description VARCHAR(240), product_rental_period_limit VARCHAR(240), product_image_id VARCHAR(240), owner_id INT, p_condition VARCHAR(240), requires_approval BOOL, n_content TEXT) BEGIN DECLARE imgid INT; SELECT id INTO imgid FROM images WHERE file_name = product_image_id ORDER BY date_added DESC; if (requires_approval) THEN INSERT INTO products (product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, ownerid, `condition`, content) VALUES (product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, owner_id, p_condition, n_content); ELSE INSERT INTO products (product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, ownerid, `condition`, authorized, content) VALUES (product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, owner_id, p_condition, TRUE, n_content); END IF; SET @last_id = LAST_INSERT_ID(); INSERT INTO has (users_id, products_id, status) VALUES (owner_id, @last_id, 0); INSERT INTO products_has_images (products_id, images_id) VALUES (@last_id, imgid); END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("createProduct Procedure Created")
}
func EditProductProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `EditProduct` (p_id VARCHAR(240), p_name VARCHAR(240), p_description VARCHAR(240), p_rental_period_limit VARCHAR(240), p_condition VARCHAR(240), comments_enabled BOOL, comments_require_approvala BOOL, n_content TEXT) BEGIN UPDATE products SET product_name = p_name, product_description = p_description, product_rental_period_limit = p_rental_period_limit, date_updated = NOW(), `condition` = p_condition, enable_comments = comments_enabled, comments_require_approval = comments_require_approvala, content = n_content WHERE product_id = p_id; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("EditProduct Procedure Created")
}

func removeImageProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE removeImage(p_id VARCHAR(240)) BEGIN DECLARE pid INT; SELECT id INTO pid FROM products WHERE product_id = p_id; DELETE from products_has_images where products_id = pid limit 1; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("removeImage Procedure Created")
}
func removeProductProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE removeProduct(u_token VARCHAR(240), p_id VARCHAR(240)) BEGIN DECLARE uid INT; DECLARE pid INT; DECLARE iid INT; SELECT user_id INTO uid FROM tokens WHERE token = u_token; SELECT id INTO pid FROM products WHERE product_id = p_id; SELECT images_id INTO iid FROM products_has_images WHERE products_id = pid; DELETE FROM has WHERE users_id = uid AND products_id = pid; DELETE FROM has WHERE products_id = pid; DELETE FROM products_has_tags WHERE products_id = pid; DELETE FROM products_has_likes WHERE products_id = pid; DELETE FROM user_rent_product WHERE products_id = pid; DELETE FROM products_has_images WHERE products_id = pid; DELETE from products_has_comments WHERE products_id = pid; DELETE FROM products WHERE id = pid; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("removeProduct Procedure Created")
}
func getListingOfTagProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `getListingOfTag`(s_tag VARCHAR(240), start INT, count INT) BEGIN SELECT product_id AS id, products.product_name AS name, products.product_description AS description, products.date_added, products.date_updated, products.product_rental_period_limit AS time_period, has.products_id AS image_id, username AS username, md5(email) AS gravatar, age_rating FROM has LEFT OUTER JOIN products ON has.products_id = products.id LEFT OUTER JOIN users ON has.users_id = users.id LEFT OUTER JOIN products_has_tags ON products.id = products_has_tags.products_id LEFT OUTER JOIN tags ON products_has_tags.tags_id = tags.id WHERE tag = s_tag ORDER BY products.date_updated DESC LIMIT START, COUNT; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getListingOfTag Procedure Created")
}
func searchByTagProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `searchByTag`(s_tag TEXT) BEGIN SELECT * FROM products_has_tags LEFT JOIN tags ON products_has_tags.tags_id = tags.id WHERE tag LIKE CONCAT(\" % \", s_tag, \" % \"); END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("searchByTag Procedure Created")
}
func searchFilterByTagProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `searchFilterByTag`(s_tag TEXT, start INT, count INT) BEGIN SELECT product_id AS id, products.product_name AS name, products.product_description AS description, products.date_added, products.date_updated, products.product_rental_period_limit AS time_period, has.products_id AS image_id, username AS username, md5(email) AS gravatar FROM has LEFT OUTER JOIN products ON has.products_id = products.id LEFT OUTER JOIN users ON has.users_id = users.id LEFT OUTER JOIN products_has_tags ON products.id = products_has_tags.products_id LEFT OUTER JOIN tags ON products_has_tags.tags_id = tags.id WHERE tag LIKE CONCAT(\" % \", s_tag, \" % \") GROUP BY product_id ORDER BY products.date_updated DESC LIMIT START, COUNT; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("searchFilterByTag Procedure Created")
}
func getListingProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE getListing() BEGIN SELECT username AS username, md5(email) AS gravatar, product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, products.id AS id, content FROM products LEFT JOIN has ON products.id = has.products_id LEFT JOIN users ON has.users_id = users.id ORDER BY date_updated DESC; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getListing Procedure Created")
}
func getProductProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE getProduct(pid VARCHAR(240)) BEGIN DECLARE p_id INT; DECLARE tags TEXT; SELECT id INTO p_id FROM products WHERE product_id = pid; SELECT GROUP_CONCAT(CONCAT(tag) SEPARATOR ', ') INTO tags FROM products_has_tags LEFT JOIN tags ON products_has_tags.tags_id = tags.id WHERE products_id = p_id; IF (tags IS NULL) THEN SET tags = \"no tags\"; END IF; SELECT product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, username, products.id AS id, tags, `condition`, enable_comments as comments_enabled, comments_require_approval as comments_require_approval, content, age_rating FROM has LEFT JOIN users ON has.users_id = users.id LEFT JOIN products ON has.products_id = products.id WHERE product_id = pid; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getProduct Procedure Created")
}
func addTagProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `addTag`(p_id VARCHAR(240), p_tag VARCHAR(240)) BEGIN DECLARE pid INT; DECLARE tag_exists INT; DECLARE tag_relation_exists BOOL; SELECT id INTO pid FROM products WHERE product_id = p_id; SELECT id INTO tag_exists FROM tags WHERE tag = p_tag; SELECT EXISTS(SELECT * FROM products_has_tags WHERE products_id = pid AND tags_id = tag_exists) INTO tag_relation_exists; IF (tag_exists IS NULL) THEN INSERT INTO tags (tag) VALUES (p_tag); SELECT id INTO tag_exists FROM tags WHERE tag = p_tag; IF (tag_relation_exists IS FALSE) THEN INSERT INTO products_has_tags (products_id, tags_id) VALUES (pid, tag_exists); END IF; ELSE IF (tag_relation_exists IS FALSE) THEN INSERT INTO products_has_tags (products_id, tags_id) VALUES (pid, tag_exists); END IF; END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("addTag Procedure Created")
}
func removeTagProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `removeTag`(p_id VARCHAR(240), p_tag VARCHAR(240)) BEGIN DECLARE pid INT; DECLARE tid INT; SELECT id INTO pid FROM products WHERE product_id = p_id; SELECT id INTO tid FROM tags WHERE tag = p_tag; DELETE FROM products_has_tags WHERE tags_id = tid AND products_id = pid; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("removeTag Procedure Created")
}
func GetTagsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `GetTags`(pid VARCHAR(240)) BEGIN DECLARE p_id INT; SELECT id INTO p_id FROM products WHERE product_id = pid; SELECT tags.tag FROM products_has_tags LEFT JOIN tags ON products_has_tags.tags_id = tags.id WHERE products_id = p_id; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("GetTags Procedure Created")
}
func getImageProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE getImage(pid INT) BEGIN SELECT file_name, title, date_added FROM products_has_images LEFT JOIN images ON products_has_images.images_id = images.id WHERE products_id = pid ORDER BY date_added ASC; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getImage Procedure Created")
}
func RequestToBorrowItemProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE RequestToBorrowItem(u_pid VARCHAR(240), u_token VARCHAR(240)) BEGIN DECLARE uid INT; DECLARE pid INT; SELECT user_id INTO uid FROM tokens WHERE token = u_token; SELECT id INTO pid FROM products WHERE product_id = u_pid; IF (SELECT EXISTS(SELECT * FROM users_requests_products WHERE users_id = uid AND products_id = pid)) THEN SELECT FALSE, \"null\", NOW(); ELSE INSERT INTO users_requests_products (products_id, users_id, date_requested) VALUES (pid, uid, NOW()); SELECT TRUE, u_pid, NOW(); END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("RequestToBorrowItem Procedure Created")
}
func OwnerGetProductRequestsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `OwnerGetProductRequests`(u_token VARCHAR(240), u_pid VARCHAR(240)) BEGIN DECLARE uid INT; DECLARE pid INT; SELECT user_id INTO uid FROM tokens WHERE token = u_token; SELECT id INTO pid FROM products WHERE product_id = u_pid; SELECT username, md5(email) as gravatar, date_requested from users_requests_products LEFT JOIN users ON users_requests_products.users_id = users.id where products_id = pid ORDER BY date_requested DESC; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("OwnerGetProductRequests Procedure Created")
}
func UserGetOngoingRequestsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `UserGetOngoingRequests`(u_token VARCHAR(240), step int, count int) BEGIN DECLARE uid int; select user_id into uid from tokens where token = u_token; SELECT product_id AS id, product_name AS name, product_description AS description, date_requested AS date_requested, products_id AS image_id, username AS owner FROM users_requests_products LEFT OUTER JOIN products ON users_requests_products.products_id = products.id LEFT OUTER JOIN users ON users_id = users.id WHERE users_requests_products.date_requested <= NOW() AND users_id = uid ORDER BY users_requests_products.date_requested DESC LIMIT step, count; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("UserGetOngoingRequests Procedure Created")
}
func OwnerGetRequestsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `OwnerGetRequests`(u_token VARCHAR(240), step int, count int) BEGIN DECLARE uid int; DECLARE requests int; select user_id into uid from tokens where token = u_token; SELECT username AS username, md5(email) AS gravatar, product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, products.id AS id, COUNT(date_requested) as requests FROM users_requests_products LEFT JOIN products ON users_requests_products.products_id = products.id LEFT JOIN has ON products.id = has.products_id LEFT JOIN users ON has.users_id = users.id WHERE ownerid = uid GROUP BY product_id ORDER BY date_updated DESC LIMIT step, count; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("OwnerGetRequests Procedure Created")
}
func OwnerDropRequestProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `OwnerDropRequest`(u_pid VARCHAR(240) , t_username VARCHAR(240)) BEGIN DECLARE uid INT; DECLARE pid INT; SELECT id INTO uid FROM users WHERE username = t_username; SELECT id INTO pid FROM products WHERE product_id = u_pid; DELETE FROM users_requests_products WHERE products_id = pid AND users_id = uid; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("OwnerDropRequest Procedure Created")
}
func CancelRequestProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `CancelRequest`(p_id VARCHAR(240), u_token VARCHAR(240)) BEGIN DECLARE uid INT; DECLARE pid INT; SELECT user_id INTO uid FROM tokens WHERE token = u_token; SELECT id INTO pid FROM products WHERE product_id = p_id; DELETE FROM users_requests_products WHERE products_id = pid AND users_id = uid; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("CancelRequest Procedure Created")
}
func GetRequestStatusProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE GetRequestStatus(p_id VARCHAR(240), u_token VARCHAR(240)) BEGIN DECLARE uid INT; DECLARE pid INT; DECLARE pname VARCHAR(240); DECLARE d_requested DATETIME; DECLARE numofreq int; DECLARE owner bool; SELECT user_id INTO uid FROM tokens WHERE token = u_token; SELECT id INTO pid FROM products WHERE product_id = p_id; SELECT product_name, date_requested INTO pname, d_requested FROM users_requests_products LEFT JOIN users ON users_requests_products.users_id = users.id LEFT JOIN products ON users_requests_products.products_id = products.id WHERE products_id = pid AND users_id = uid; SELECT EXISTS( SELECT * FROM has LEFT JOIN users ON has.users_id = users.id LEFT JOIN products ON has.products_id = products.id WHERE product_id = p_id AND users_id = uid) INTO owner; if (owner IS TRUE) THEN select count(*) into numofreq from users_requests_products where products_id = pid; ELSE SET numofreq = 0; END IF; IF (pname IS NULL) THEN SELECT FALSE AS requested, \"null\" AS product_title, NOW() AS date_requested, numofreq as requests; ELSE SELECT TRUE AS requested, pname AS product_title, d_requested AS date_requested, numofreq as requests; END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("GetRequestStatus Procedure Created")
}
func RentFromRequestProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE RentFromRequest(u_pid VARCHAR(240), usrname VARCHAR(240)) BEGIN DECLARE userid INT; DECLARE days INT; DECLARE pid INT; SELECT id INTO userid FROM users WHERE username = usrname; SELECT id, product_rental_period_limit INTO pid, days FROM products WHERE product_id = u_pid; DELETE FROM users_requests_products WHERE users_id = userid AND products_id = pid; INSERT INTO user_rent_product (products_id, users_id, date_received, date_due) VALUES (pid, userid, NOW(), DATE_ADD(NOW(), INTERVAL days DAY)); END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("RentFromRequest Procedure Created")
}

func RentItemProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE RentItem(u_pid VARCHAR(240), usrname VARCHAR(240)) BEGIN DECLARE userid INT; DECLARE days INT; DECLARE pid INT; SELECT id INTO userid FROM users WHERE username = usrname; SELECT id, product_rental_period_limit INTO pid, days FROM products WHERE product_id = u_pid; INSERT INTO user_rent_product (products_id, users_id, date_received, date_due) VALUES (pid, userid, NOW(), DATE_ADD(NOW(), INTERVAL days DAY)); END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("RentItem Procedure Created")
}
func ReturnItemProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE ReturnItem(o_token VARCHAR(240), product VARCHAR(240)) BEGIN DECLARE productid INT; DECLARE u_id VARCHAR(240); DECLARE tmp_u_id VARCHAR(240); SELECT user_id INTO tmp_u_id FROM tokens WHERE token = o_token; SELECT users_id INTO u_id FROM user_rent_product LEFT JOIN products ON user_rent_product.products_id = products.id LEFT JOIN users ON ownerid = users.id WHERE users_id = tmp_u_id AND product_id = product; SELECT id INTO productid FROM products WHERE product_id = product; DELETE FROM user_rent_product WHERE users_id = u_id AND products_id = productid; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("ReturnItem Procedure Created")
}
func ReturnItemAsOwnerProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE ReturnItemAsOwner(o_token VARCHAR(240), product VARCHAR(240)) BEGIN DECLARE pid INT; DECLARE uid INT; DECLARE tmp_u_id INT; SELECT user_id INTO tmp_u_id FROM tokens WHERE token = o_token; SELECT products_id, users_id INTO pid, uid FROM user_rent_product LEFT JOIN products ON user_rent_product.products_id = products.id WHERE product_id = product AND ownerid = tmp_u_id; DELETE FROM user_rent_product WHERE users_id = uid AND products_id = pid; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("ReturnItemAsOwner Procedure Created")
}
func checkItemAvailabilityProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `checkItemAvailability`(product VARCHAR(240), usrname VARCHAR(240)) BEGIN DECLARE due_date DATETIME; SELECT date_due INTO due_date FROM user_rent_product LEFT OUTER JOIN products ON user_rent_product.products_id = products.id LEFT OUTER JOIN users ON user_rent_product.users_id = users.id WHERE products.product_id = product ORDER BY products.date_updated DESC; IF (due_date > NOW()) THEN SELECT FALSE, due_date; ELSE SELECT TRUE AS Available, NOW(); END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("checkItemAvailability Procedure Created")
}
func getPagedProductsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE getPagedProducts(step INT, count INT, sorting bool) BEGIN if (sorting) THEN SELECT product_id AS id, product_name AS name, product_description AS description, products.date_added, products.date_updated, product_rental_period_limit AS time_period, products.id AS image_id, username AS username, md5(email) AS gravatar, `condition`, content, COALESCE(sum(likes.`like`), 0) as likes, age_rating FROM has LEFT OUTER JOIN products ON has.products_id = products.id LEFT OUTER JOIN users ON has.users_id = users.id LEFT JOIN products_has_likes ON products.id = products_has_likes.products_id LEFT JOIN likes ON products_has_likes.likes_id = likes.id WHERE visable = TRUE AND authorized = TRUE GROUP BY products.product_id ORDER BY products.date_added DESC LIMIT step, COUNT; ELSE SELECT product_id AS id, product_name AS name, product_description AS description, products.date_added, products.date_updated, product_rental_period_limit AS time_period, products.id AS image_id, username AS username, md5(email) AS gravatar, `condition`, content, COALESCE(sum(likes.`like`), 0) as likes, age_rating FROM has LEFT OUTER JOIN products ON has.products_id = products.id LEFT OUTER JOIN users ON has.users_id = users.id LEFT JOIN products_has_likes ON products.id = products_has_likes.products_id LEFT JOIN likes ON products_has_likes.likes_id = likes.id WHERE visable = TRUE AND authorized = TRUE GROUP BY products.product_id ORDER BY products.date_added ASC LIMIT step, COUNT; END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getPagedProducts Procedure Created")
}
func getMostRecentPagedProductsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE getMostRecentPagedProducts(step INT, count INT) BEGIN SELECT product_id AS id, product_name AS name, product_description AS description, products.date_added, date_updated, product_rental_period_limit AS time_period, products.id AS image_id, username AS username, md5(email) AS gravatar, `condition`, content, COALESCE(sum(likes.`like`), 0) as likes, age_rating FROM has LEFT OUTER JOIN products ON has.products_id = products.id LEFT OUTER JOIN users ON has.users_id = users.id LEFT JOIN products_has_likes ON products.id = products_has_likes.products_id LEFT JOIN likes ON products_has_likes.likes_id = likes.id WHERE visable = TRUE AND authorized = TRUE ORDER BY products.date_added DESC LIMIT step, COUNT; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getMostRecentPagedProducts Procedure Created")
}
func getRecentlyUpdatedPagedProductsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE getRecentlyUpdatedPagedProducts(step INT, count INT, t_order bool) BEGIN if (t_order) THEN SELECT product_id AS id, product_name AS name, product_description AS description, products.date_added, date_updated, product_rental_period_limit AS time_period, products.id AS image_id, username AS username, md5(email) AS gravatar, `condition`, content, COALESCE(sum(likes.`like`), 0) as likes, age_rating FROM has LEFT OUTER JOIN products ON has.products_id = products.id LEFT OUTER JOIN users ON has.users_id = users.id LEFT JOIN products_has_likes ON products.id = products_has_likes.products_id LEFT JOIN likes ON products_has_likes.likes_id = likes.id WHERE visable = TRUE AND authorized = TRUE GROUP BY products.product_id ORDER BY products.date_updated DESC LIMIT step, COUNT; ELSE SELECT product_id AS id, product_name AS name, product_description AS description, products.date_added, date_updated, product_rental_period_limit AS time_period, products.id AS image_id, username AS username, md5(email) AS gravatar, `condition`, content, COALESCE(sum(likes.`like`), 0) as likes, age_rating FROM has LEFT OUTER JOIN products ON has.products_id = products.id LEFT OUTER JOIN users ON has.users_id = users.id LEFT JOIN products_has_likes ON products.id = products_has_likes.products_id LEFT JOIN likes ON products_has_likes.likes_id = likes.id WHERE visable = TRUE AND authorized = TRUE GROUP BY products.product_id ORDER BY products.date_updated ASC LIMIT step, COUNT; END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getRecentlyUpdatedPagedProducts Procedure Created")
}
func getRandomPagedProductsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE getRandomPagedProducts(step INT, count INT) BEGIN SELECT product_id AS id, product_name AS name, product_description AS description, products.date_added, date_updated, product_rental_period_limit AS time_period, products.id AS image_id, username AS username, md5(email) AS gravatar, `condition`, content, COALESCE(sum(likes.`like`), 0) as likes, age_rating FROM has LEFT OUTER JOIN products ON has.products_id = products.id LEFT OUTER JOIN users ON has.users_id = users.id LEFT JOIN products_has_likes ON products.id = products_has_likes.products_id LEFT JOIN likes ON products_has_likes.likes_id = likes.id WHERE visable = TRUE AND authorized = TRUE GROUP BY products.product_id ORDER BY RAND() LIMIT step, COUNT; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getRandomPagedProducts Procedure Created")
}
func getMostLikedPagedProductsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE getMostLikedPagedProducts(step INT, count INT, sortOrder bool) BEGIN if (sortOrder) THEN SELECT product_id AS id, product_name AS name, product_description AS description, products.date_added, products.date_updated, product_rental_period_limit AS time_period, products.id AS image_id, username AS username, md5(email) AS gravatar, `condition`, content, COALESCE(sum(likes.`like`), 0) as likes, age_rating FROM has LEFT JOIN products ON has.products_id = products.id LEFT JOIN users ON has.users_id = users.id LEFT JOIN products_has_likes ON products.id = products_has_likes.products_id LEFT JOIN likes ON products_has_likes.likes_id = likes.id WHERE visable = TRUE AND authorized = TRUE GROUP BY products.product_id ORDER BY COALESCE(sum(likes.`like`), 0) DESC LIMIT step, COUNT; ELSE SELECT product_id AS id, product_name AS name, COALESCE(sum(likes.`like`), 0) as likes, product_description AS description, products.date_added, products.date_updated, product_rental_period_limit AS time_period, products.id AS image_id, username AS username, md5(email) AS gravatar, `condition`, content, COALESCE(sum(likes.`like`), 0) as likes, age_rating FROM has LEFT JOIN products ON has.products_id = products.id LEFT JOIN users ON has.users_id = users.id LEFT JOIN products_has_likes ON products.id = products_has_likes.products_id LEFT JOIN likes ON products_has_likes.likes_id = likes.id WHERE visable = TRUE AND authorized = TRUE GROUP BY products.product_id ORDER BY COALESCE(sum(likes.`like`), 0) ASC LIMIT step, COUNT; END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getMostLikedPagedProducts Procedure Created")
}
func getRentedProductsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE getRentedProducts(username VARCHAR(240), step INT, count INT) BEGIN SELECT product_id AS id, product_name AS name, product_description AS description, date_added, date_updated, product_rental_period_limit AS time_period, product_image_id AS image_id, username AS owner FROM user_rent_product LEFT OUTER JOIN products ON user_rent_product.products_id = products.id WHERE user_rent_product.date_due < NOW() ORDER BY products.date_updated DESC LIMIT step, COUNT; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getRentedProducts Procedure Created")
}
func getCurrentlyRentingProductsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE getCurrentlyRentingProducts(u_name VARCHAR(240), step INT, count INT) BEGIN SELECT product_id AS id, product_name AS name, product_description AS description, date_due AS due_date, date_received AS received_date, products_id AS image_id, username AS owner FROM user_rent_product LEFT OUTER JOIN products ON user_rent_product.products_id = products.id LEFT OUTER JOIN users ON user_rent_product.users_id = users.id WHERE username = u_name ORDER BY user_rent_product.date_due ASC LIMIT step, COUNT; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getCurrentlyRentingProducts Procedure Created")
}
func getRentalsDueThreeDaysProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE getRentalsDueThreeDays() BEGIN SELECT product_id AS id, product_name AS name, product_description AS description, date_due AS due_date, date_received AS received_date, products_id AS image_id, username AS owner FROM user_rent_product LEFT OUTER JOIN products ON user_rent_product.products_id = products.id LEFT OUTER JOIN users ON user_rent_product.users_id = users.id WHERE user_rent_product.date_due < DATE_ADD(CURDATE(), INTERVAL +3 DAY) ORDER BY user_rent_product.date_due ASC; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getRentalsDueThreeDays Procedure Created")
}
func getUsernameProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE getUsername(usrtoken VARCHAR(240)) BEGIN DECLARE userid INT; SELECT user_id INTO userid FROM tokens WHERE token = usrtoken; SELECT username FROM users WHERE id = userid; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getUsername Procedure Created")
}
func checkProductAvailabilityProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `checkProductAvailability`(product VARCHAR(240)) BEGIN DECLARE due_date DATETIME; DECLARE active_state BOOLEAN; SELECT date_due INTO due_date FROM user_rent_product LEFT OUTER JOIN products ON user_rent_product.products_id = products.id LEFT OUTER JOIN users ON user_rent_product.users_id = users.id WHERE products.product_id = product ORDER BY products.date_updated DESC; IF (due_date > NOW()) THEN SELECT FALSE AS available, due_date AS due_date; ELSE SELECT TRUE AS available, NOW() AS due_date; END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("checkProductAvailability Procedure Created")
}
func checkAuthedProductAvailabilityProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `checkAuthedProductAvailability`(product VARCHAR(240)) BEGIN DECLARE due_date DATETIME; DECLARE taken_date DATETIME; DECLARE user_name VARCHAR(240); SELECT date_due, date_received, username INTO due_date, taken_date, user_name FROM user_rent_product LEFT JOIN products ON user_rent_product.products_id = products.id LEFT JOIN users ON user_rent_product.users_id = users.id WHERE products.product_id = product ORDER BY date_received DESC LIMIT 1; IF (user_name != \" \") THEN IF (due_date > NOW()) THEN SELECT FALSE AS available, due_date, taken_date, user_name AS username; ELSE SELECT TRUE AS available, NOW() AS due_date, NOW() AS taken_date, user_name AS username; END IF; ELSE SELECT TRUE AS available, NOW(), NOW(), \"null\"; END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("checkAuthedProductAvailability Procedure Created")
}
func getHolderProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `getHolder`(p_id VARCHAR(240)) BEGIN DECLARE uid INT; DECLARE pid INT; SELECT id into pid from products WHERE product_id = p_id; SELECT username, md5(email) from user_rent_product LEFT JOIN users ON user_rent_product.users_id = users.id WHERE products_id = pid; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getHolder Procedure Created")
}
func getOwnerProductsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE getOwnerProducts(u_token VARCHAR(240), step INT, count INT) BEGIN DECLARE usrname VARCHAR(240); SELECT username INTO usrname FROM tokens LEFT OUTER JOIN users ON tokens.user_id = users.id WHERE token = u_token; SELECT product_id AS id, product_name AS name, product_description AS description, date_added, date_updated, product_rental_period_limit AS time_period, products_id AS image_id, username AS username, md5(email) AS gravatar, `condition`, enable_comments as comments_enabled, comments_require_approval as comments_require_approval FROM has LEFT OUTER JOIN products ON has.products_id = products.id LEFT OUTER JOIN users ON has.users_id = users.id WHERE users.username = usrname ORDER BY date_updated DESC LIMIT step, COUNT; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getOwnerProducts Procedure Created")
}
func CheckProductAvailabilityOwnerProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE CheckProductAvailabilityOwner(o_token VARCHAR(240), p_id VARCHAR(240)) BEGIN DECLARE due_date DATETIME; DECLARE active_state BOOLEAN; SELECT date_due, active INTO due_date, active_state FROM user_rent_product LEFT OUTER JOIN products ON user_rent_product.products_id = products.id LEFT OUTER JOIN users ON user_rent_product.users_id = users.id WHERE products.product_id = p_id ORDER BY products.date_updated DESC; IF (active_state = 1) THEN IF (due_date > NOW()) THEN SELECT FALSE AS available, due_date AS due_date; ELSE SELECT TRUE AS available, NOW() AS due_date; END IF; ELSE SELECT TRUE AS available, NOW() AS due_date; END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("CheckProductAvailabilityOwner Procedure Created")
}
func isOwnerProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE isOwner(u_token VARCHAR(240), p_id VARCHAR(240)) BEGIN DECLARE u_id INT; SELECT user_id INTO u_id FROM tokens WHERE token = u_token; SELECT EXISTS( SELECT * FROM has LEFT JOIN users ON has.users_id = users.id LEFT JOIN products ON has.products_id = products.id WHERE product_id = p_id AND users_id = u_id) AS owner; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("isOwner Procedure Created")
}
func ownerProductStatusProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE ownerProductStatus(u_token VARCHAR(240), p_id VARCHAR(240)) BEGIN DECLARE isOwner BOOL; DECLARE u_id INT; SELECT user_id INTO u_id FROM tokens WHERE token = u_token; SELECT EXISTS( SELECT * FROM has LEFT JOIN users ON has.users_id = users.id LEFT JOIN products ON has.products_id = products.id WHERE product_id = p_id AND users_id = u_id) INTO isOwner; IF (isOwner) THEN CALL checkAuthedProductAvailability(p_id); ELSE CALL checkProductAvailability(p_id); END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("ownerProductStatus Procedure Created")
}
func getUserIDofTokenProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE getUserIDofToken(u_token VARCHAR(240)) BEGIN SELECT user_id FROM tokens WHERE token = u_token; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getUserIDofToken Procedure Created")
}
func getIndexProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE getIndex() BEGIN SELECT Title AS title, description AS description FROM Site WHERE id = 1; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getIndex Procedure Created")
}
func updateSiteProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE updateSite(s_title VARCHAR(240), s_description VARCHAR(240), u_token VARCHAR(240)) BEGIN DECLARE urole VARCHAR(240); DECLARE uid int; select user_id into uid from tokens where token = u_token; select role into urole from users where id = uid; if (urole = \"admin\") THEN UPDATE Site SET Title = s_title, Description = s_description WHERE id = 1; select \"true\"; ELSE select \"false\"; END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("undefined Procedure Created")
}
func GetUnAuthorizedProductsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE GetUnAuthorizedProducts(step INT, count INT) BEGIN SELECT product_id AS id, product_name AS name, product_description AS description, date_added, date_updated, product_rental_period_limit AS time_period, products_id AS image_id, username AS username, md5(email) AS gravatar FROM has LEFT OUTER JOIN products ON has.products_id = products.id LEFT OUTER JOIN users ON has.users_id = users.id WHERE authorized = FALSE ORDER BY products.date_updated DESC LIMIT step, COUNT; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("GetUnAuthorizedProducts Procedure Created")
}
func AuthorizeProductProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `AuthorizeProduct`(p_id VARCHAR(240), u_token VARCHAR(240)) BEGIN DECLARE urole VARCHAR(240); DECLARE uid int; select user_id into uid from tokens where token = u_token; select role into urole from users where id = uid; if (urole = \"admin\") THEN update products SET authorized = 1 where product_id = p_id; select \"true\"; ELSE select \"false\"; END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("AuthorizeProduct Procedure Created")
}
func GetUserRoleProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `GetUserRole`(u_token VARCHAR(240)) BEGIN DECLARE uid INT; select user_id into uid from tokens where token = u_token; select username, role from users where id = uid; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("GetUserRole Procedure Created")
}
func DeleteImageProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `DeleteImage`(image_title VARCHAR(240)) BEGIN DECLARE iid INT; select id into iid from images where title = image_title limit 1; delete from products_has_images where images_id = iid; delete from images where id = iid; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("DeleteImage Procedure Created")
}
func getUsersProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `getUsers`(u_token VARCHAR(240)) BEGIN DECLARE uid INT; DECLARE token_expires DATETIME; DECLARE isAdmin BOOL; select user_id, date_expires into uid, token_expires from tokens where token = u_token; select exists(select role from users where id = uid AND role = 'admin') into isAdmin; if (token_expires > NOW() AND isAdmin) THEN select username, md5(email) as gravatar, date_registered, email, role from users; ELSE select \"nope\", \"nope\", \"nope\", \"nope\", \"nope\"; END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getUsers Procedure Created")
}
func removeUserAsAdminProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `removeUserAsAdmin`(u_name VARCHAR(240), u_token VARCHAR(240)) BEGIN DECLARE auid INT; DECLARE token_expires DATETIME; DECLARE isAdmin BOOL; DECLARE UID INT; select user_id, date_expires into auid, token_expires from tokens where token = u_token; select exists(select role from users where id = auid AND role = \"admin\") into isAdmin; SELECT id INTO UID FROM users WHERE username = u_name; IF (isAdmin) THEN DELETE FROM tokens WHERE user_id = UID; DELETE FROM users WHERE username = u_name; SELECT \"user deleted\"; ELSE SELECT \"user not deleted\"; END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("removeUserAsAdmin Procedure Created")
}
func ChangeUserRoleProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `ChangeUserRole`(c_username VARCHAR(240), n_role VARCHAR(240), u_token VARCHAR(240)) BEGIN DECLARE uid INT; DECLARE isAdmin BOOL; select user_id into uid from tokens where token = u_token; select exists(select role from users where id = uid AND role = 'admin') into isAdmin; if (isAdmin) THEN UPDATE users SET role = n_role where username = c_username; END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("ChangeUserRole Procedure Created")
}
func getMostUsedTagsProcedure(db config.Context) {
	stmt, err := db.Session.Exec("CREATE PROCEDURE `getMostUsedTags` (start INT, count INT, sortOrder bool) BEGIN if (sortOrder) THEN select tag, count(tag) from products_has_tags left JOIN tags ON products_has_tags.tags_id = tags.id GROUP BY tag ORDER BY count(tag) DESC LIMIT start, count; ELSE select tag, count(tag) from products_has_tags left JOIN tags ON products_has_tags.tags_id = tags.id GROUP BY tag ORDER BY count(tag) ASC LIMIT start, count; END IF; END;")
	if err != nil {
		log.Println(err)
	}
	_ = stmt
	log.Println("getMostUsedTags Procedure Created")
}