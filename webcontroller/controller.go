// Package webcontroller supplies common services for web http requests
package webcontroller

import (
	"fmt"
	"net/http"

	"github.com/user/dockerclient/webservices"
)

// ServiceFunc is webservice function definition
type ServiceFunc func(http.HandlerFunc) http.HandlerFunc
// Controller structure holds array of webservice functions
type Controller struct {
	services []ServiceFunc
}

// AddService function appends webservice function to services array
func (c Controller) AddService(service ServiceFunc) Controller {
	newServices := make([]ServiceFunc, 0, len(c.services) + 1)
	fmt.Printf("AddService services size = %d\n", len(c.services))
	newServices = append(newServices, c.services...)	// ... notation explodes the array to be appended
	newServices = append(newServices, service)
	return Controller{newServices}
}
// Run function executes the sequence of web services added to the Controller
// returning the http.HandlerFunc passed in
func (c Controller) Run(fn http.HandlerFunc) http.HandlerFunc {
	fmt.Printf("Run services size = %d\n", len(c.services))
	for n := range c.services {
		fmt.Printf("n = %d\n", n)
		fn = c.services[len(c.services) - 1 - n](fn)
	}
	return fn
}
// Init function creates the Controller and adds the webservices in order of execution
func Init() Controller {
	c := Controller{}
	c = c.AddService(webservices.Log)
	c = c.AddService(webservices.Recovery)
	return c
}
