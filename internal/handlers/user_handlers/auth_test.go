package user_handlers

//func TestUserHandler_Signup(t *testing.T) {
//	// Create a new user handler with a mock repository and logger
//	repo := &mocks.Repository{}
//	logger := logger2.NewLogger(logger2.Config{})
//	userHandler := &UserHandler{repo: repo, logger: logger}
//
//	// Create a new HTTP request with a JSON-encoded user DTO in the body
//	user := entities.UserDto{
//		User: entities.User{
//			DefaultModel: entities.DefaultModel{},
//			Address:      "0x12kg209934nnv237",
//			NetworkName:  "",
//			NetworkURL:   "",
//			ChainID:      "",
//			RefreshKey:   "",
//		},
//	}
//	requestBody, _ := json.Marshal(user)
//	request := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(requestBody))
//
//	// Create a new HTTP response recorder to capture the response
//	responseRecorder := httptest.NewRecorder()
//
//	// Call the Signup function with the request and response recorder
//	userHandler.Signup(responseRecorder, request)
//
//	// Check that the response status code is 200 OK
//	if responseRecorder.Code != http.StatusOK {
//		t.Errorf("expected status code %d but got %d", http.StatusOK, responseRecorder.Code)
//	}
//
//	// Check that the response body contains the expected data
//	expectedResponseBody := `Data: {"name":"John Doe","email":"john.doe@example.com","password":"password123"}`
//	if responseRecorder.Body.String() != expectedResponseBody {
//		t.Errorf("expected response body %q but got %q", expectedResponseBody, responseRecorder.Body.String())
//	}
//}
