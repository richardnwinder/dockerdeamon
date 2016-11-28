// Package accountmanager
// contains maps for querying databases
package accountmanager

var Admin admin
// admin structure holds admin credentials
type admin struct {
	AccessKey   string
	Name   		string
	Password   	string
}
