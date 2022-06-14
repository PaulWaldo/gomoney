package http

// func New()
// const Path = "/account"

// func (cont Controller) AddRoutes(r *gin.Engine) {
// 	r.Group(Path)
// 	{
// 		r.GET("/", cont.handleAccountCreate)

// 	}
// }

// func (cont Controller) handleAccountCreate(c *gin.Context) {
// 	t, err := template.ParseFiles("internal/html/accounts.gohtml")
//     if err != nil {
//         panic(err)
//     }

// 	var request createAccountRequest
// 	c.Bind(&request)
// 	acct := models.Account{Type: request.AccountType, Name: request.Name}
// 	cont.db.Create(acct)
// 	c.HTML(http.StatusOK, "accounts.gohtml", )
// 	// acctType, err := domain.AccountTypeFromString(request.AccountType)
// 	// res := controller.db.Where(&activationToken).First(&activationToken)

// 	// if err != nil {
// 	// 	http.Error(w, fmt.Sprintf("Bad Request: %s", err), http.StatusBadRequest)
// 	// 	return
// 	// }
// 	// id, err := a.svc.Create(request.Name, acctType)
// 	// if err != nil {
// 	// 	http.Error(w, fmt.Sprintf("Internal Server Error: %s", err), http.StatusInternalServerError)
// 	// 	return
// 	// }
// 	// response := createAccountResponse{ID: id}
// 	// encoder := json.NewEncoder(w)
// 	// encoder.Encode(response)
// }
