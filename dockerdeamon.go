package main

import (
	"flag"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"github.com/googollee/go-socket.io"
	"github.com/user/dockerclient/webcontroller"
	"github.com/user/dockerclient/webservices"
	am "github.com/user/dockerclient/accountmanager"
	"github.com/user/dockerclient/services"
	
)

type httpError struct {
	status  int
	message string
}

func (e *httpError) Error() string {
	return fmt.Sprintf("%d - %s", e.status, e.message)
}

type page struct {
	Title string
	Body  []byte
}

const webPath string = "/home/richard/DockerClient"

func errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	var status int
	//var message string
	if ae, ok := err.(*httpError); ok {
		status = ae.status
		//message = ae.message
	}
	w.WriteHeader(status)
	switch status {
	case http.StatusBadRequest:
		fmt.Fprint(w, "custom 400")
	case http.StatusUnauthorized:
		fmt.Fprint(w, "custom 401")
	case http.StatusNotFound:
		fmt.Fprint(w, "custom 404")
	}
}

// regexp used to validate the requested path for security purposes
var validPath = regexp.MustCompile("^/([a-z.-])+/([a-zA-Z0-9./-]+)$")

/****************************************************************************************/
/*  the getAbsPath function accepts a file path as a string constant and validates the  */
/*  path. The path is converted to an absolute path, and returned to calling function   */
/****************************************************************************************/
func getAbsPath(path string) (string, error) {
	// validate the path
	m := validPath.FindStringSubmatch(webPath + path)
	// if path is not valid return URL is invalid error
	if m == nil {
		fmt.Print("match = nil\n")
		err := &httpError{400, "URL is invalid"}
		return "", err
	}
	//return the valid absolute path and no errors
	fmt.Printf("valid path: %s\n", m[0])
	return m[0], nil
}

/****************************************************************************************/
/*  the getPage function accepts a file pathname as a string constant and converts it   */
/*  to an absolute path. The file is retrieved from storage and new Page object is      */
/*  created and returned to calling function                                            */
/****************************************************************************************/
func getPage(urlpath string) (*page, error) {
	fmt.Printf("GET, %s\n", urlpath)
	// get the absolute filepath
	filepath, err := getAbsPath(urlpath)
	// if the getAbsPath routine fails return the error
	if err != nil { // Invalid URL path
		return nil, err
	}
	// else fill in the placeholders to construct the Page Object
	title := path.Base(filepath)
	body, err := ioutil.ReadFile(filepath)
	// if the io routine fails return the error
	if err != nil {
		err := &httpError{404, "File not found"}
		return nil, err
	}
	// else return the Page Object with no errors
	return &page{Title: title, Body: body}, nil
}

/****************************************************************************************/
/*  the getDataHandler function queries the incoming http request for the URL query     */
/*  parameters, and passes the parameters to the appropriate routine handler defined    */
/*  by the index parameter.                                                             */
/****************************************************************************************/
func getDataHandler(w http.ResponseWriter, r *http.Request) {
	// get the current-user cookie from the http request
	//var user string

	//user = "guest"

	command := r.URL.Query().Get("cmd")
	//fmt.Printf("cmd = %q\n", command)
	params := r.URL.Query().Get("params")
	fmt.Printf("params = %q\n", params)
	parameters := strings.Fields(params)
	if command != "" {
		switch command {
		case "docker":
			{
				var data string
				fmt.Printf("len(parameters) = %d\n", len(parameters))
				if len(parameters) > 0 {
					out, gerr := services.GoDocker(parameters)
					if !gerr.IsNil() {
						fmt.Print("gerr.PrintCodeError()..............\n")
						gerr.PrintCodeError()
					}
					for idx, line := range(out) {
						data = data + line
						fmt.Printf("return idx: %d\n", idx)
					}
					fmt.Printf("return data: '%s'\n", data)
					fmt.Fprintf(w, "%s", data)
				}
			}
		case "html":
			{
				html, gerr := services.GoHtml(parameters)
				if !gerr.IsNil() {
					fmt.Print("gerr.PrintCodeError()..............\n")
					gerr.PrintCodeError()
				}
				fmt.Fprintf(w, "%s", html)
			}	
		case "connected":
			{
				resp := "false"
				_, err := http.Get("http://hub.docker.com/")
				if err == nil {
					resp = "true"
				}
				fmt.Fprintf(w, "%s", resp)

			}
		default:
			fmt.Printf("Unrecognised index value : " + command)
		}
	}
	form := r.URL.Query().Get("form")
	fmt.Printf("form = %q\n", form)
	if form != "" {
		switch form {
		case "clubdetails":
			{
			}
		default:
			fmt.Printf("Unrecognised form value : " + form)
		}
	}
}

/****************************************************************************************/
/*  the getHandler function queries the incoming http request for the URL path of the   */
/*  static file, retrieves the file from storage and writes it to the http response     */
/*  and terminates the response, which() is then returnhtml := ""html := ""ed by the server. All index files  */
/*  are stored in the index folder created in the same directory as the go server       */
/****************************************************************************************/
func getHandler(w http.ResponseWriter, r *http.Request) {
	// direct all root urls to index.html (login page)
	if r.URL.Path == "/" {
		r.URL.Path = "/index.html"
	}
	// check for missing .html extension
	if filepath.Ext(r.URL.Path) == "" {
		r.URL.Path = r.URL.Path + ".html"
	}
	
	// authenticate client
	loggedIn, err := webservices.Authenticate(&w, r)
	if err == nil {
		if loggedIn == true {
			//fmt.Printf("loggedIn = false\n")
			r.URL.Path = "/services" + r.URL.Path
			p, err := getPage(r.URL.Path)
			if err != nil {
				errorHandler(w, r, err)
				return
			}
			if path.Ext(r.URL.Path) == ".css" {
				w.Header().Add("Content-Type", "text/css")
			}
			fmt.Fprintf(w, "%s", p.Body)
			return
		} else {
			r.URL.Path = "/login" + r.URL.Path
			p, err := getPage(r.URL.Path)
			if err != nil {
				if path.Ext(r.URL.Path) == ".html" {
					r.URL.Path = "/login/index.html"
					p, err := getPage(r.URL.Path)
					if err != nil {
						fmt.Printf("not loggedIn error page %s\n", r.URL.Path)
						errorHandler(w, r, err)
						return
					}
					if path.Ext(r.URL.Path) == ".css" {
						w.Header().Add("Content-Type", "text/css")
					}
					fmt.Fprintf(w, "%s", p.Body)
					return
				} else {
					errorHandler(w, r, err)
					return
				}
			}
			if path.Ext(r.URL.Path) == ".css" {
				w.Header().Add("Content-Type", "text/css")
			}
			fmt.Fprintf(w, "%s", p.Body)
		}
	} else {
		// write error to client page
		err := &httpError{401, "Unauthorized"}
		errorHandler(w, r, err)
	}

	// it is possible to use templating to create the final html file
	//if path.Base(r.URL.Path) == "/html/index.html" {
	//t, _ :=template.ParseFiles("view.html")
	//t.Execute(w, p)
	//}
	//fmt.Print("Ext = " + path.Ext(r.UR()()().Path) + "\n")
	// correct response header content type for css files
	return
}

/****************************************************************************************/
/*  the getLoginHandler function queries the incoming http request for the URL login     */
/*  parameters, and passes the parameters to the appropriate routine handler defined    */
/*  by the index parameter.                                                             */
/****************************************************************************************/
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var response string
	var err error
	var loggedIn bool
	loggedIn, err = webservices.Authenticate(&w, r)
	if err == nil {
		if loggedIn != true {
			//fmt.Printf("loggedIn = false\n")
			response, err = am.Login(&w, r)
			if err != nil {
				// need error recovery here to give a smoother user experience
				// redirect with  html error page
				return
			}
		}
	//loginStatus := strings.Split(response, ":")
	//if loginStatus[0] == "OK" {
		// redirect to frontpage
		fmt.Fprintf(w, "%s", response)
		//fmt.Fprintln(w, "Correct Password.")
	//} else {
		//fmt.Fprintf(w, "%s", response)
		//fmt.Fprintln(w, "Wrong Password.")
	//}
	} else {
		// need error recovery here to give a smoother user experience
		// redirect with  html error page
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	var response string
	var err error
	var loggedIn bool
	loggedIn, err = webservices.Authenticate(&w, r)
	if err == nil {
		if loggedIn == true {
			//fmt.Printf("loggedIn = false\n")
			response, err = am.Logout(&w, r)
			if err != nil {
				// need error recovery here to give a smoother user experience
				// redirect with  html error page
				return
			}
		}
	//loginStatus := strings.Split(response, ":")
	//if loginStatus[0] == "OK" {
		// redirect to frontpage
		fmt.Fprintf(w, "%s", response)
		//fmt.Fprintln(w, "Correct Password.")
	//} else {
		//fmt.Fprintf(w, "%s", response)
		//fmt.Fprintln(w, "Wrong Password.")
	//}
	} else {
		// need error recovery here to give a smoother user experience
		// redirect with  html error page
	}
}

/****************************************************************************************/
/*  the handler function queries the i()()()ncoming http request for pre-allocated URL paths  */
/*  any other paths are served as static files using GET, POST and PUT methods          */
/****************************************************************************************/
func handler(w http.ResponseWriter, r *http.Request) {
	// first check for pre-allocated URL paths and pass to appropriate handler functions
	//fmt.Printf("Request: %q\n", html.EscapeString(r.URL.Path))
	//fmt.Printf("Request: %q\n", html.EscapeString(r.URL.RawQuery))
	if r.URL.Path == "/getData" {
		//fmt.Printf("getData, %q\n", r.URL.RawQuery)
		getDataHandler(w, r)
		return
	} else if r.URL.Path == "/login" {
		fmt.Printf("login, %q\n", r.URL.RawQuery)
		//fmt.Fprintf(w, "%s", string("Login Page"))
		loginHandler(w, r)
		return
	} else if r.URL.Path == "/logout" {
		fmt.Printf("login, %q\n", r.URL.RawQuery)
		//fmt.Fprintf(w, "%s", string("Login Page"))
		logoutHandler(w, r)
		return
	}
	// the http request method is selected and the response/request methods
	// are passed to the relevant handler function
	if r.Method == "GET" || r.Method == "get" {
		fmt.Printf("GET, %q\n", html.EscapeString(r.URL.Path))
		getHandler(w, r)
	} else if r.Method == "POST" || r.Method == "post" {
		fmt.Printf("POST, %q\n", html.EscapeString(r.URL.Path))
		//if r.URL.Path == "/postData" {
		//postDataHandler(w, r)
		//return
		//} else if r.URL.Path == "/upload" {
		//uploadHandler(w, r)
		//return
		//}()()()()
	} else if r.Method == "PUT" || r.Method == "put" {
		fmt.Printf("PUT, %q\n", html.EscapeString(r.URL.Path))
	} else {
		http.Error(w, "Invalid request method.\n", 405)
	}
}

/****************************************************************************************/
/*  the main function just has one route to the route handler function                  */
/*  its main task is to setup the go server                                             */
/****************************************************************************************/
func main() {
	flag.StringVar(&am.Admin.Name, "admin", "", "admin name")
	flag.StringVar(&am.Admin.Password, "password", "", "admin password")
	flag.Parse()
	fmt.Println("adminName:", am.Admin.Name)
   fmt.Println("adminPassword:", am.Admin.Password)
   
   // Initialise socketio
   var err error
   ioserver,err := socketio.NewServer(nil)
   if err != nil {
		log.Fatal(err)
	}

	ioserver.On("connection", func(client socketio.Socket) {
		services.AddClient(client)
		
		client.On("disconnection", func() {
			services.RemoveClient(client)
		})
	})
   ioserver.On("error", func(so socketio.Socket, err error) {
		
	})
   
   
	// Initialise authentication module
	webservices.InitAuthentication()
	// initialise server
	server := webcontroller.Init()
	http.Handle("/socket.io/", ioserver)
	http.Handle("/", server.Run(handler))
	fmt.Print("Server started using port : 8081\n")
	// the server is started with a logger to log fatal errors
	// it is a blocking procedure
	// at the moment all requests are passed to the root handler function
	log.Fatal(http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil))
	
}
