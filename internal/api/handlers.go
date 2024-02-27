package api

func getLatestUpdatedRequestHandler(ctx *gin.Context) {
	var request Request
	// code := ctx.Query("code")
	// code2 := url.QueryEscape(code)
	db.Order("updated DESC").First(&request)
	ctx.JSON(http.StatusOK, &request)
}