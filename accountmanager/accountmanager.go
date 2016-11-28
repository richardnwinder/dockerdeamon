// Package accountmanager
package accountmanager

import (
	"crypto/subtle"
	"fmt"
	"net/http"
	"strings"

	"github.com/user/dockerclient/webservices"
)

// Login returns logged in client
func Login(w *http.ResponseWriter, r *http.Request) (response string, err error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		if err == http.ErrNoCookie {
			err = nil
		} else {
			http.SetCookie(*w, cookie)
			response = "ERROR:0010:No authorisation!"
			return response, err // return error no cookie to client
		}
	}

	// Open database to retrieve login password for user
	username := r.FormValue("username")
	password := r.FormValue("password")
	response, err = checkLogin(username, password)
	loginStatus := strings.Split(response, ":")
	if loginStatus[0] == "OK" {
		// login user
		webservices.UpdateLoginStatus(cookie.Value, true)
	}
	return response, err
}

// checkLogin validation routines
func checkLogin(admin string, password string) (response string, err error) {
	fmt.Printf("client: %s\n", admin)
	fmt.Printf("password: %s\n", password)
	if admin == Admin.Name {
		if subtle.ConstantTimeCompare([]byte(password), []byte(Admin.Password)) == 1 {
			redirectPath := "/docker.html"
			fmt.Printf("Redirecting to: %s\n", redirectPath)
			//response.setHeader("Location", redirectPath);
			response = "OK:0302:" + redirectPath
		} else {
			fmt.Printf("password incorrect - %s : %s\n", Admin.Password, password)
			response = "ERROR:0012:Incorrect admin name or password!"
		}
	} else {
		fmt.Printf("password incorrect - %s : %s\n", Admin.Password, password)
		response = "ERROR:0012:Incorrect admin name or password!"
	}
	return response, nil
}

func Logout(w *http.ResponseWriter, r *http.Request) (response string, err error) {
	cookie, err := r.Cookie("session")
	redirectPath := "/"
	fmt.Printf("Redirecting to: %s\n", redirectPath)
	webservices.UpdateLoginStatus(cookie.Value, false)
	response = "OK:0302:" + redirectPath
	return response, nil
}
