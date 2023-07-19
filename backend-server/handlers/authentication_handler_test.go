package handlers

//func Test_AuthenticationHandler(t *testing.T) {
//
//	router := setupRouter()
//
//	json := "{\n    \"username\": \"testUser\", \n    \"password\": \"strongPassword\",\n    \"client_id\": \"testClient\",\n    \"client_secret\": \"clientSecret\",\n    \"grant_type\": \"dbckhbd\"\n}"
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("POST", "/login", strings.NewReader(json))
//
//	router.ServeHTTP(w, req)
//
//	fmt.Println(w.Code, w.Body)
//
//}
//
//func setupRouter() *gin.Engine {
//	gin.SetMode(gin.TestMode)
//	r := gin.Default()
//	r.POST("/login", func(c *gin.Context) {
//		c.String(200, "pong")
//	})
//	return r
//}
