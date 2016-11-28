// Package webservices supplies common services for web http requests
package webservices

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"github.com/satori/go.uuid"
)

type _client struct {
	loggedIn bool
}

var sessionStore map[string]_client
var storageMutex sync.RWMutex

// InitAuthentication initialises authentication webservice
func InitAuthentication() {
	createSessionStore()
}
// CreateSessionStore initialises cookie client map
func createSessionStore() {
	sessionStore = make(map[string]_client)
}
// Authenticate returns the current status of the client request
func Authenticate(w *http.ResponseWriter, r *http.Request) (loggedIn bool, err error) {
	// get the session cookie from the http request
	cookie, err := r.Cookie("session")
	// fmt.Printf("Authenticate cookie: %s\n", cookie)
	// check for errors reading session cookie
	if err != nil {
		if err == http.ErrNoCookie { // if error No Cookie just discard err and continue
			// fmt.Printf("Authenticate no cookie!\n")
			err = nil
		} else { // else return error to web server
			//http.SetCookie(*w, cookie)
			fmt.Printf("Error %s\n", err)
			return false, err
		}
	}

	var present bool
	var client _client

	if cookie != nil { // if cookie is not nil
		// fmt.Printf("Authenticate cookie != nil: %s\n", cookie)
		// return the client associated with the cookie value from session store
		storageMutex.RLock()
		client, present = sessionStore[cookie.Value]
		storageMutex.RUnlock()
		if present == false {
			client = _client{false}
			storageMutex.Lock()
			sessionStore[cookie.Value] = client
			storageMutex.Unlock()
			present = true
		}
	} else {
		present = false
	}
	/*if present == true {
		fmt.Printf("Authentication present = true.\n")
	} else {
		fmt.Printf("Authentication present = false.\n")
	}*/
	if present == false { // else if no cookie then create a new one
		// fmt.Print("Authenticate cookie present = false\n")
		// fmt.Printf("Authentication creating new cookie.\n")
		value := uuid.NewV4().String()
		// fmt.Printf("Authentication value = %s.\n", value)
		cookie = &http.Cookie{
			Name:  "session",
			Value: value,
		}
		// fmt.Printf("Authentication cookie.Value = %s.\n", cookie.Value)
		client = _client{false}
		storageMutex.Lock()
		sessionStore[cookie.Value] = client
		storageMutex.Unlock()
	}
	//if client.loggedIn == true {
		//fmt.Printf("Authentication return loggedIn = true.\n")
	//} else {
		//fmt.Printf("Authentication return loggedIn = false.\n")
	//}
	// fmt.Printf("Authentication cookie.Name = %s.\n", cookie.Name)
	// fmt.Printf("Authentication cookie.Value = %s.\n", cookie.Value)
	http.SetCookie(*w, cookie)
	return client.loggedIn, err
}
// UpdateLoginStatus updates the session store client record loggedIn status
func UpdateLoginStatus(cookieValue string, loggedIn bool)(error){
	var err error = nil
	if cookieValue != "" { // if cookie is not nil
		// fmt.Printf("Authenticate cookie != nil: %s\n", cookie)
		// return the client associated with the cookie value from session store
		var present bool
		var client _client

		storageMutex.RLock()
		client, present = sessionStore[cookieValue]
		storageMutex.RUnlock()
		if present == true {
			client.loggedIn = loggedIn
			storageMutex.Lock()
			sessionStore[cookieValue] = client
			storageMutex.Unlock()
		}else{
			err = errors.New("ERROR: Client not authenticated.")
		}
	}else{
		err = errors.New("ERROR: No authentication code.")
	}
	return err
}
