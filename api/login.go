package api

import "fmt"

func Login(email string, password string) {
	resp, err := getV1Request().
		SetBody(map[string]interface{}{"email": email, "password": password}).
		Post("/login")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
