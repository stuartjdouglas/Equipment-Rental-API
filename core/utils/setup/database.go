package setup

import (
	"log"
	"github.com/remony/Equipment-Rental-API/core/config"
)

func Start(context config.Context) {
	log.Println("Setting up database")
	setupdb(context)
}

/*
				--NOTE--

	Due to the current implementation of the mysql driver,
	it does not support files, and/or sequential inserts
	so for each table and procedure to be created it has
	to be done via separate calls.

 */

// Sets up Database
func setupdb(db config.Context) {
	// Create the tables
	createTables(db)
	// Create the procedures
	createProcedures(db)
	log.Println("DB is fully installed")
}

// Create the procedures
func createProcedures(db config.Context) {
	// Call each function
	createLikeProcedure(db)
	createUnLikeProcedure(db)
	createGetLikesProcedure(db)
	createAddCommentProcedure(db)
	createEditCommentProcedure(db)
	createHasUserReviewedListingProcedure(db)
	createDisableCommentsProcedure(db)
	createEnableCommentsProcedure(db)
	DisableCommentsRequireAuthProcedure(db)
	EnableCommentsRequireAuthProcedure(db)
	ApproveCommentProcedure(db)
	GetCommentsProcedure(db)
	GetCommentProcedure(db)
	GetOwnerCommentsProcedure(db)
	DeleteCommentProcedure(db)
	GetPushNotificationIDsOfUserProcedure(db)
	GetPushNotificationIDsOfProductProcedure(db)
	addPushNotificationRegIDProcedure(db)
	registerProcedure(db)
	removeUserProcedure(db)
	doesUserExistProcedure(db)
	getDigestProcedure(db)
	loginProcedure(db)
	addImageProcedure(db)
	AddAnotherImageProcedure(db)
	imageExistsProcedure(db)
	createProductProcedure(db)
	EditProductProcedure(db)
	removeImageProcedure(db)
	removeProductProcedure(db)
	getListingOfTagProcedure(db)
	searchFilterByTagProcedure(db)
	getListingProcedure(db)
	getProductProcedure(db)
	addTagProcedure(db)
	removeTagProcedure(db)
	GetTagsProcedure(db)
	getImageProcedure(db)
	RequestToBorrowItemProcedure(db)
	OwnerGetProductRequestsProcedure(db)
	UserGetOngoingRequestsProcedure(db)
	OwnerGetRequestsProcedure(db)
	OwnerDropRequestProcedure(db)
	CancelRequestProcedure(db)
	GetRequestStatusProcedure(db)
	RentFromRequestProcedure(db)
	RentItemProcedure(db)
	ReturnItemProcedure(db)
	ReturnItemAsOwnerProcedure(db)
	checkItemAvailabilityProcedure(db)
	getPagedProductsProcedure(db)
	getMostRecentPagedProductsProcedure(db)
	getRecentlyUpdatedPagedProductsProcedure(db)
	getRandomPagedProductsProcedure(db)
	getMostLikedPagedProductsProcedure(db)
	getRentedProductsProcedure(db)
	getCurrentlyRentingProductsProcedure(db)
	getRentalsDueThreeDaysProcedure(db)
	getUsernameProcedure(db)
	checkProductAvailabilityProcedure(db)
	checkAuthedProductAvailabilityProcedure(db)
	getHolderProcedure(db)
	getOwnerProductsProcedure(db)
	CheckProductAvailabilityOwnerProcedure(db)
	isOwnerProcedure(db)
	ownerProductStatusProcedure(db)
	getUserIDofTokenProcedure(db)
	getIndexProcedure(db)
	updateSiteProcedure(db)
	GetUnAuthorizedProductsProcedure(db)
	AuthorizeProductProcedure(db)
	GetUserRoleProcedure(db)
	DeleteImageProcedure(db)
	getUsersProcedure(db)
	removeUserAsAdminProcedure(db)
	ChangeUserRoleProcedure(db)
	getMostUsedTagsProcedure(db)
	searchByTagProcedure(db)
}

// Create tables
func createTables(db config.Context) {
	createUsersTable(db)
	createTokenTable(db)
	createPostsTable(db)
	createImagesTable(db)
	createProductsTable(db)
	createProductHasImagesTable(db)
	createHasTable(db)
	createUserRentProductTable(db)
	createUserRequestsProductTable(db)
	createTagsTable(db)
	createProductHasTagsTable(db)
	createSiteTable(db)
	createPushTokensTable(db)
	createUsersHasPushTokensTable(db)
	createCommmentsTable(db)
	createProductHasCommentsTable(db)
	createLikesTable(db)
	createProductHasLikesTable(db)
}

