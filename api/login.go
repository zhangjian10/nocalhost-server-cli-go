package api

func Login() {
	getV1Request().SetBody(`{"email":"","password":""}`).Post("/v1/login")
}
