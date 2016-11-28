// Package services contains docker command routines
package services

import (
	"bufio"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/user/goerror"
	"github.com/googollee/go-socket.io"
)

//type Package struct {
	//Description string
//}


// Client list of all clients connected
var Clients []socketio.Socket

func AddClient(clienttoadd socketio.Socket) {
	Clients = append(Clients, clienttoadd)
}

func RemoveClient(clienttoremove socketio.Socket) {
	for i, client := range Clients {
		if client == clienttoremove {
			Clients = append(Clients[:i], Clients[i+1:]...)
			break
		}
	}
}
// flag for parse function
var whitespace bool

// USE: with strings.FieldsFunc
// FUNC: info string[] = strings.FieldsFunc(string, parse)
// function parse removes multiple whitespace from strings.
func parse(c rune) bool {
	if c == ' ' && whitespace {
		whitespace = true
		return true
	} else if c != ' ' {
		whitespace = false
		return false
	} else {
		whitespace = true
	}
	return false
}

// Docker Container structure
type container struct {
	command string
	created string
	id string
	img string
	name string
	ports string
	status string
}

// Docker Image structure
type image struct {
	repository string
	tag string
	imageId string
	created string
	size string
}

var desc map[string]string

// GetImages function returns the current images available and installed
func GetImages() (string, *goerror.GoError) {
	html := ""
	connected := false
	if hasConnection() {
		connected = true
	}
	// docker build current directory
	cmdName := "docker"
	cmdArgs := []string{"images"}

	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		gerr := goerror.FromError(err)
		gerr.SetMsg("Error creating StdoutPipe for command")
		return html, gerr
	}
	scanner := bufio.NewScanner(cmdReader)
	err = cmd.Start()
	if err != nil {
		gerr := goerror.FromError(err)
		//gerr.SetMsg("Error running ImagesI()()magesmagesmages")
		return html, gerr
	}
	fmt.Printf("scanner.Scan()\n")
	//go func() {packageD
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Printf("docker images out | %s\n", line[:9])
		if line[:9] == "dockeros2" {
			info := strings.Fields(line[10:])
			fmt.Printf("docker images out | %s\n", info[0])
			app := info[0]
			html = html + "<div class=\"item\">" +
				"<div style=\"margin: 5px 20px; width: 60px; float: left; color: rgb(51, 51, 255);\">" +
				"<img src=\"/packages/" + app + "/" + app + ".png\" alt=\"icon\" style=\"margin: 0px 20px; width: 40px; vertical-align: middle;\">" +
				"</div>" +
				"<div style=\"margin: 12px 0px; overflow: hidden; width: 120px; float: left; color: rgb(51, 51, 255);\">" +
				"<span style=\"font-weight:	bold;\">" + app + "</span>" +
				"</div>" +
				"<div style=\"margin: 12px 20px; width: 400px; float: left;\">" +
				desc[app] + "<br>" +
				"</div>" +
				"<div style=\"margin: 10px 100px 10px 0px; float: right;\">"
			if connected {
				html = html +
					"<button type=\"button\" stStartyle=\"font-sizStartStarte: 14px;\" onclick=\"alert('Hello World!')\">Uninstall</button>"
			} else {
				html = html +
					"<button type=\"button\" disabled style=\"font-size: 14px;\" onclick=\"alert('Hello World!')\">Uninstall</button>"
			}
			html = html +
				"</div>" +
				"<br style=\"clear: both;\">" +
				"</div>"
		}
	}
	err = cmd.Wait()
	if err != nil {
		gerr := goerror.FromError(err)
		//gerr.SetMsg("Error running command")
		return html, gerr
	}
	return html, goerror.New("")
}

// GetImages function returns the current images available and installed
func GetPackageLists() {
	desc = make(map[string]string)
	// read packages.lst
	desc["cathode"] = "Cathode Terminal presents an old cathode screen image. For amusement."
	desc["torbrowser"] = "The Tor Browser offers protectection and anonimity whist surfing the internet."
}

// GoHtml function returns the current docker states
func GoHtml(param []string) (string, *goerror.GoError) {
	var cmdArgs []string
	html := ""

	fmt.Print("GoHtml() - param[0] = " + param[0] + "\n")
	switch param[0] {
		case "containers":
			{
				cmdArgs = []string{"ps", "-a"}
				out, gerr := docker(cmdArgs)
				if !gerr.IsNil() {
					gerr := goerror.New("servicesos2.go:146:ERROR:500:Docker command error - " + param[0])
					return html, gerr
				}
				for idx, line := range(out) {
					if idx != 0 {
						fmt.Print("GoHtml::idx = %d\n", idx)
						html = html + getContainer(line)
					}
				}
			}
		case "images":
			{
				cmdArgs = param
				out, gerr := docker(cmdArgs)
				if !gerr.IsNil() {
					gerr := goerror.New("servicesos2.go:161:ERROR:500:Docker command error - " + param[0])
					return html, gerr
				}
				for idx, line := range(out) {
					if idx != 0 {
						html = html + getImage(line)
					}
				}
			}
		case "info":
			{
				cmdArgs = param
				out, gerr := docker(cmdArgs)
				if !gerr.IsNil() {
					gerr := goerror.New("servicesos2.go:175:ERROR:500:Docker command error - " + param[0])
					return html, gerr
				}
				html = "<div style=\"margin: 0px 20px 20px 20px;\">" +
						"<a href=\"docker.html\">Back</a>" +
						"</div>" +
						"<div style=\"margin: 0px 20px 20px 60px;\"><pre>"
				for idx, line := range(out) {
					fmt.Printf("GoHtml::idx = %d\n", idx)
					fmt.Printf("line = " + line + "\n")
					html = html + line + "\n"
				}
				html = html + "</pre></div>"
			}	
		case "inspect":
			{
				cmdArgs = param
				out, gerr := docker(cmdArgs)
				if !gerr.IsNil() {
					gerr := goerror.New("servicesos2.go:194:ERROR:500:Docker command error - " + param[0])
					return html, gerr
				}
				for idx, line := range(out) {
					fmt.Printf("GoHtml::idx = %d\n", idx)
					fmt.Printf("line = " + line + "\n")
					html = html + getInspect(line)
				}
			}
		case "logging":
			{
				cmdArgs = []string{"ps", "-a"}
				out, gerr := docker(cmdArgs)
				if !gerr.IsNil() {
					gerr := goerror.New("servicesos2.go:208:ERROR:500:Docker command error - " + param[0])
					return html, gerr
				}				
				for idx, line := range(out) {
					if idx != 0 {
						html = html + getLogging(line)
					}
				}
			}
			case "logs":
			{
				cmdArgs = param
				out, gerr := docker(cmdArgs)
				if !gerr.IsNil() {
					gerr := goerror.New("servicesos2.go:222:ERROR:500:Docker command error - " + param[0])
					return html, gerr
				}
				html = "<div style=\"margin: 0px 20px 20px 20px;\">" +
						"<a href=\"docker.html\">Back</a>" +
						"</div>" +
						"<div style=\"margin: 0px 20px 20px 60px;\"><pre>"		
				for idx, line := range(out) {
					if idx != 0 {
						html = html + line + "\n"
					}
				}
				html = html + "</pre></div>"
			}
		case "pull":
			{
				cmdArgs = param
				out, gerr := docker(cmdArgs)
				if !gerr.IsNil() {
					gerr := goerror.New("servicesos2.go:241:ERROR:500:Docker command error - " + param[0])
					return html, gerr
				}
				html = "<div style=\"margin: 0px 5px 5px 5px;\"><pre>"
				for idx, line := range(out) {
					fmt.Printf("GoHtml::idx = %d\n", idx)
					fmt.Printf("line = " + line + "\n")
					html = html + line + "\n"
				}
				html = html + "</pre></div>"
			}
		case "push":
			{
				cmdArgs = param
				out, gerr := docker(cmdArgs)

				html = "<div style=\"margin: 0px 5px 5px 5px;\"><pre>"
				for idx, line := range(out) {
					fmt.Printf("GoHtml::idx = %d\n", idx)
					fmt.Printf("line = " + line + "\n")
					html = html + line + "\n"
				}
				html = html + "</pre></div>"
				if !gerr.IsNil() {
					gerr := goerror.New("servicesos2.go:257:ERROR:500:Docker command error - " + param[0])
					return html, gerr
				}
			}
		default:
			{
				gerr := goerror.New("servicesos2.go:270:ERROR:500:Unrecognised param value - " + param[0])
				return html, gerr
			}				
	}
	return html, goerror.New("")
}
	

// GoDocker function returns the current docker images installed and running
func GoDocker(cmdArgs []string) ([]string, *goerror.GoError) {
	cmdName := "docker"
	var cmdOut []string
	fmt.Print("cmdArgs[0] = " + cmdArgs[0] + "\n")
	switch cmdArgs[0] {
		case "images":
			{
			}
		case "info":
			{
			}	
		case "inspect":
			{
			}
		case "logs":
			{
				//mt.Printf("cmdArgs.length = %d\n", len(cmdArgs))
				/*for idx, arg := range(cmdArgs) {
					fmt.Printf("cmdArg[%d] = '%s'\n", idx, arg)
				}*/
				//cmdOut, gerr := goLogs(cmdArgs)
				//return cmdOut, gerr
			}	
		case "ps":
			{
			}
		case "pull":
			{
			}
		case "push":
			{
			}
		case "create":
			{
				if len(cmdArgs) < 2 {
					gerr := goerror.New("servicesos2.go:314:ERROR:500:Not enough params - " + cmdArgs[0])
					return cmdOut, gerr
				}
			}
		case "start":
			{
				if len(cmdArgs) < 2 {
					gerr := goerror.New("servicesos2.go:321:ERROR:500:Not enough params - " + cmdArgs[0])
					return cmdOut, gerr
				}
			}
		case "stop":
			{
				if len(cmdArgs) < 2 {
					gerr := goerror.New("servicesos2.go:328:ERROR:500:Not enough params - " + cmdArgs[0])
					return cmdOut, gerr
				}
			}
		case "rm":
			{
				if len(cmdArgs) < 2 {
					gerr := goerror.New("servicesos2.go:1335:ERROR:500:Not enough params - " + cmdArgs[0])
					return cmdOut, gerr
				}
			}
		case "rmi":
			{
				if len(cmdArgs) < 2 {
					gerr := goerror.New("servicesos2.go:342:ERROR:500:Not enough params - " + cmdArgs[0])
					return cmdOut, gerr
				}
			}
		default:
			{
				gerr := goerror.New("servicesos2.go:348:ERROR:500:Unrecognised param value - " + cmdArgs[0])
				return cmdOut, gerr
			}
	}
	
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		gerr := goerror.FromError(err)
		gerr.SetMsg("Error creating StdoutPipe for command")
		cmdOut[0] = "Error cmd.Reader():" + fmt.Sprint(err)
		
		return cmdOut, gerr
	}
	if cmdArgs[0] == "logs" {
		cmd.Stderr = cmd.Stdout
	}
	scanner := bufio.NewScanner(cmdReader)
	
	err = cmd.Start()
	if err != nil {
		gerr := goerror.FromError(err)
		cmdOut[0] = "Error cmd.Start():" + fmt.Sprint(err)
		//gerr.SetMsg("Error running ImagesI()()magesmagesmages")
		return cmdOut, gerr
	}

	fmt.Printf("scanner.Scan()\n")
	go func() {
		//idx := 0
		for scanner.Scan() {

					//fmt.Printf("GoDocker::idx = %d\n", idx)
					cmdOut = append(cmdOut, scanner.Text())
		}
	}()
	err = cmd.Wait()
	if err != nil {
		gerr := goerror.FromError(err)
		//gerr.SetMsg("Error running command")
		cmdOut[0] = "Error cmd.Wait():" + fmt.Sprint(err)
		return cmdOut, gerr
	}
	fmt.Printf("cmdOut.length = %d\n", len(cmdOut))
	return cmdOut, goerror.New("")
}

// gologs function returns the current docker images installed and running
func goLogs(cmdArgs []string) ([]string, *goerror.GoError) {
	cmdName := "docker"
	var cmdOut []string
	fmt.Print("cmdArgs[0] = " + cmdArgs[0] + "\n")
	switch cmdArgs[0] {
		case "logs":
			{
				fmt.Printf("cmdArgs.length = %d\n", len(cmdArgs))
				for idx, arg := range(cmdArgs) {
					fmt.Printf("cmdArg[%d] = '%s'\n", idx, arg)
				}	
			}	
		default:
			{
				gerr := goerror.New("servicesos2.go:411:ERROR:500:Unrecognised param value - " + cmdArgs[0])
				return cmdOut, gerr
			}
	}
	out, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if err != nil {
		gerr := goerror.FromError(err)
		//gerr.SetMsg("Error running command")
		cmdOut = append(cmdOut, "Error cmd.Wait():" + fmt.Sprint(err))
		return cmdOut, gerr
	}
	cmdOut = append(cmdOut, string(out[:]))
	fmt.Printf("goLogs cmdOut.length = %d\n", len(cmdOut))
	return cmdOut, goerror.New("")
}

// docker function runs the docker cli command with arguments from cmdArgs array
func docker(cmdArgs []string) ([]string, *goerror.GoError) {
	var cmdOut []string
	
	fmt.Print("cmdArgs[0] = " + cmdArgs[0] + "\n")
	
	cmd := exec.Command("docker", cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		gerr := goerror.FromError(err)
		gerr.SetMsg("Error creating StdoutPipe for command")
		cmdOut[0] = "Error cmd.Reader():" + fmt.Sprint(err)
		
		return cmdOut, gerr
	}
	if cmdArgs[0] == "logs" {
		cmd.Stderr = cmd.Stdout
	}
	scanner := bufio.NewScanner(cmdReader)
	
	err = cmd.Start()
	if err != nil {
		gerr := goerror.FromError(err)
		cmdOut[0] = "Error cmd.Start():" + fmt.Sprint(err)
		//gerr.SetMsg("Error running ImagesI()()magesmagesmages")
		return cmdOut, gerr
	}

	fmt.Printf("scanner.Scan()\n")
	go func() {
		for scanner.Scan() {
			switch cmdArgs[0] {
				case "pull":
				{
					for _, client := range Clients {
						client.Emit("pullmsg", string(scanner.Text()[:]))
					}
					cmdOut = append(cmdOut, scanner.Text())
				}
				case "push":
				{
					for _, client := range Clients {
						client.Emit("pushmsg", scanner.Text())
					}
					cmdOut = append(cmdOut, scanner.Text())
				}
				default:
				{
			//fmt.Printf("docker::idx = %d\n", idx)
					cmdOut = append(cmdOut, scanner.Text())
					fmt.Printf("scanner.Text() = " + scanner.Text() + "\n")
				}
			}
		}
	}()
	err = cmd.Wait()
	if err != nil {
		gerr := goerror.FromError(err)
		//gerr.SetMsg("Error running command")
		cmdOut[0] = "Error cmd.Wait():" + fmt.Sprint(err)
		return cmdOut, gerr
	}
	fmt.Printf("cmdOut.length = %d\n", len(cmdOut))
	return cmdOut, goerror.New("")
}


func hasConnection() bool {
	_, err := http.Get("http://hub.docker.com/")
	if err == nil {
		return true
	}
	return false
}

func getImage(line string) (string) {
	var html string
	var img image
	//fmt.Printf("getImage() | %s\n", line)
	info := strings.FieldsFunc(line, parse)
	if len(info) == 5 {
		img = image{repository: info[0], tag: info[1], imageId: info[2], created: info[3], size: info[4]}
	} else {
		return html
	}
	img.repository = strings.TrimRight(img.repository, " ")
	img.tag = strings.TrimRight(img.tag, " ")
	img.imageId = strings.TrimRight(img.imageId, " ")
	img.created = strings.TrimRight(img.created, " ")
	img.size = strings.TrimRight(img.size, " ")
	//fmt.Printf("docker images out | %s\n", info[0])
	var source = img.repository + ":" + img.tag;
	html = html + "<div class=\"image-item\">" +
	"<div style=\"margin: 2px 20px; width: 50px; float: left; color: rgb(51, 51, 255);\">" +
	"<img onclick=\"showInfo('" + img.imageId + "'); return false\" src=\"/images/image.png\" alt=\"icon\" style=\"margin: 0px 20px; width: 40px; vertical-align: middle;\">" +
	"</div>" +
	"<div id=\"" + img.imageId + ".repo\"style=\"margin: 12px 0px; overflow: hidden; white-space: nowrap; width: 200px; float: left; color: rgb(51, 51, 255);\">" +
	"<span style=\"font-weight:	bold;\">" + img.repository + "</span>" +
	"</div>" +
	"<div style=\"margin: 12px 10px; overflow: hidden; white-space: nowrap; width: 80px; float: left; color: grey;\">" +
	"<span id=\"" + img.imageId + ".tag\" style=\"font-weight:	bold;\">" + img.tag + "</span>" +
	"</div>" +
	"<div style=\"margin: 12px 10px; overflow: hidden; width: 80px; float: left;\">" + img.size +
	"</div>" +
	"<div style=\"margin: 12px 10px; overflow: hidden; white-space: nowrap; width: 120px; float: left;\">" +
	"<span id=\"" + img.imageId + ".status\" style=\"font-weight:	bold;\">" + img.created + "</span>" +
	"</div>"
	html = html +
		"<div style=\"margin: 10px 50px 10px 0px; float: right; color: rgb(51, 51, 255);\">" +
		"<img id=\"" + img.imageId + ".wait\" src=\"../images/wait2.gif\" alt=\"wait\" style=\"display: none; height: 20px; width 60px; vertical-align: middle;\">" +
		"<img onclick=\"showContainerDialog('" + source + "'); return false\" src=\"/images/container.png\" alt=\"create\" title=\"create container\" style=\"margin: 0px 5px; height: 32px; width: 32px; vertical-align: middle;\">" +
		"<img onclick=\"showImagePushDialog('" + source + "'); return false\" src=\"/images/pull-image.png\" alt=\"push\" title=\"push image\" style=\"margin: 0px 5px; height: 28px; width: 28px; vertical-align: middle;\">" +
		"<img onclick=\"removeImage('" + img.imageId + "'); return false\" src=\"/images/remove-image.png\" alt=\"remove\" title=\"remove image\" style=\"margin: 0px 5px; height: 32px; width: 32px; vertical-align: middle;\">" +
		"</div>"
	html = html +
		"<br style=\"clear: both;\">" +
		"</div>"
	return html
}

func getInfo(line string) (string) {
	var html string
	//fmt.Printf("getInspect() | %s\n", line)
	if line == "[" {
		html = "<div style=\"margin: 0px 20px 20px 20px;\">" +
			"<a href=\"containers.html\">Back</a>" +
			"</div>" +
			"<div><pre>"
	} else if line == "]" {
		html = "</pre></div>"
	} else {
		if line == "    {" || line == "    }" {
			return html
		} else {
			html = line + "\n"
		}
	}
	return html
}
	

func getInspect(line string) (string) {
	var html string
	
	//fmt.Printf("getInspect() | %s\n", line)
	if line == "[" {
		html = "<div style=\"margin: 0px 20px 20px 20px;\">" +
			"<a href=\"containers.html\">Back</a>" +
			"</div>" +
			"<div><pre>"
	} else if line == "]" {
		html = "</pre></div>"
	} else {
		if line == "    {" || line == "    }" {
			return html
		} else {
			html = line + "\n"
		}
	}
	return html
}

func getContainer(line string) (string) {
	var html string
	var cntnr container
	//fmt.Printf("getContainer() | %s\n", line)
	info := strings.FieldsFunc(line, parse)
	if len(info) == 7 {
		cntnr = container{id: info[0], img: info[1], command: info[2], created: info[3], status: info[4], ports: info[5], name: info[6]}
	}	else if len(info) == 6 {
		cntnr = container{id: info[0], img: info[1], command: info[2], created: info[3], status: info[4], ports: "", name: info[5]}
	}
	fmt.Printf("container id = '%s'\n", cntnr.id);
	fmt.Printf("container img = '%s'\n", cntnr.img);
	fmt.Printf("container cmd = '%s'\n", cntnr.command);
	fmt.Printf("container status = '%s'\n", cntnr.status);
	fmt.Printf("container ports = '%s'\n", cntnr.ports);
	fmt.Printf("container name = '%s'\n", cntnr.name);
	
	
	cntnr.id = strings.TrimRight(cntnr.id, " ")
	cntnr.img = strings.TrimRight(cntnr.img, " ")
	cntnr.command = strings.TrimRight(cntnr.command, " ")
	cntnr.status = strings.TrimRight(cntnr.status, " ")
	cntnr.ports = strings.TrimRight(cntnr.ports, " ")
	cntnr.name = strings.TrimRight(cntnr.name, " ")
	//fmt.Printf("getContainer.Status()\n")
	if cntnr.status[:6] == "Exited" {
		cntnr.status = "Stopped"
	} else if cntnr.status[:2] == "Up" {
		cntnr.status = "Started"
	}
	//fmt.Printf("docker images out | %s\n", info[0])
	html = html + "<div id=\"" + cntnr.id + "\" class=\"item\">" +
	"<div style=\"margin: 5px 0px 5px 30px; width: 40px; float: left; color: rgb(51, 51, 255);\">" +
	"<img onclick=\"showInfo('" + cntnr.name + "'); return slicesfalse\" src=\"/images/container.png\" alt=\"contnr\" style=\"margin: 0px; width: 40px; vertical-align: middle;\">" +
	"</div>" +
	"<div style=\"margin: 10px 10px; width: 24px; float: left; color: rgb(51, 51, 255);\">" +
	"<img onclick=\"showInspect('" + cntnr.id + "'); return false\" src=\"/images/info.png\" alt=\"info\" style=\"margin: 0px; width: 24px; vertical-align: middle;\">" +
	"</div>" +
	"<div style=\"margin: 12px 10px; overflow: hidden; width: 140px; float: left; color: rgb(51, 51, 255);\">" +
	"<span style=\"font-weight:	bold;\">" + cntnr.name + "</span>" +
	"</div>" +
	"<div style=\"margin: 12px 10px; overflow: hidden; width: 80px; float: left; color: grey;\">" +
	"<span id=\"" + cntnr.name + ".status\" style=\"font-weight:	bold;\">" + cntnr.status + "</span>" +
	"</div>" +
	"<div style=\"margin: 10px 40px 10px 0px; width: 50px; float: left;\">" +
	"<img id=\"" + cntnr.name + ".wait\" src=\"../images/wait2.gif\" alt=\"wait\" style=\"display: none; height: 20px; width 40px; vertical-align: middle;\">"
	if cntnr.status == "Started" {
		html = html +
			"<button type=\"button\" id=\"" + cntnr.name + ".stop\" style=\"display: block; font-size: 14px;\" onclick=\"stopContainer('" + cntnr.name + "')\">Stop</button>" +
			"<button type=\"button\" id=\"" + cntnr.name + ".start\" style=\"display: none; font-size: 14px;\" onclick=\"startContainer('" + cntnr.name + "')\">Start</button>"
	} else {
		html = html +
		"<button type=\"button\" id=\"" + cntnr.name + ".stop\" style=\"display: none; font-size: 14px;\" onclick=\"stopContainer('" + cntnr.name + "')\" >Stop</button>" +
		"<button type=\"button\" id=\"" + cntnr.name + ".start\" style=\"display: block; font-size: 14px;\" onclick=\"startContainer('" + cntnr.name + "')\" display: block>Start</button>"
	}
	html = html +
	"</div>" +
	//"<div style=\"margin: 12px 20px; overflow: hidden; width: 120px; float: left;\">" + cntnr.id + "<br>" +
	//"</div>" +slices
	"<div style=\"margin: 12px 20px; overflow: hidden; width: 240px; float: left;\">" + cntnr.img + "<br>" +
	"</div>"
	fmt.Printf("cntnr.status = '%s'", cntnr.status);
	if (cntnr.status == "Stopped" || cntnr.status == "Created") {
		html = html +
			"<div style=\"margin: 10px 50px 10px 0px; float: right;\">" +
			"<button type=\"button\" id=\"" + cntnr.name + ".remove\" style=\"display: block; font-size: 14px;\" onclick=\"removeContainer('" + cntnr.name + "')\">Remove</button>" +
			"</div>"
	} else {
		html = html +
			"<div style=\"margin: 10px 50px 10px 0px; float: right;\">" +
			"<button type=\"button\" id=\"" + cntnr.name + ".remove\" style=\"display: none; font-size: 14px;\" onclick=\"removeContainer('" + cntnr.name + "')\">Remove</button>" +
			"</div>"
	}

	html = html +
		"<br style=\"clear: both;\">" +
		"</div>"
	return html
}

func getLogging(line string) (string) {
	var html string
	var cntnr container
	//fmt.Printf("getMonitor() | %s\n", line)
	info := strings.FieldsFunc(line, parse)
	if len(info) == 7 {
		cntnr = container{id: info[0], img: info[1], command: info[2], created: info[3], status: info[4], ports: info[5], name: info[6]}
	}	else if len(info) == 6 {
		cntnr = container{id: info[0], img: info[1], command: info[2], created: info[3], status: info[4], ports: "", name: info[5]}
	}
	
	cntnr.id = strings.TrimRight(cntnr.id, " ")
	cntnr.img = strings.TrimRight(cntnr.img, " ")
	cntnr.command = strings.TrimRight(cntnr.command, " ")
	cntnr.status = strings.TrimRight(cntnr.status, " ")
	cntnr.ports = strings.TrimRight(cntnr.ports, " ")
	cntnr.name = strings.TrimRight(cntnr.name, " ")
	//fmt.Printf("getContainer.Status()\n")
	//fmt.Printf("docker images out | %s\n", info[0])
	html = html + "<div id=\"" + cntnr.id + "\" class=\"item\">" +
	"<div style=\"margin: 5px 0px 5px 30px; width: 40px; float: left; color: rgb(51, 51, 255);\">" +
	"<img onclick=\"showInfo('" + cntnr.name + "'); return slicesfalse\" src=\"/images/container.png\" alt=\"contnr\" style=\"margin: 0px; width: 40px; vertical-align: middle;\">" +
	"</div>" +
	"<div style=\"margin: 10px 10px; width: 24px; float: left; color: rgb(51, 51, 255);\">" +
	"<img onclick=\"showInspect('" + cntnr.id + "'); return false\" src=\"/images/info.png\" alt=\"info\" style=\"margin: 0px; width: 24px; vertical-align: middle;\">" +
	"</div>" +
	"<div style=\"margin: 12px 10px; overflow: hidden; width: 140px; float: left; color: rgb(51, 51, 255);\">" +
	"<span style=\"font-weight:	bold;\">" + cntnr.name + "</span>" +
	"</div>" +
	"<div style=\"margin: 12px 10px; overflow: hidden; width: 280px; float: left; color: grey;\">" +
	"<span id=\"" + cntnr.name + ".status\" style=\"font-weight:	bold;\">" + cntnr.status + "</span>" +
	"</div>" +
	"<div style=\"margin: 10px 40px 10px 0px; width: 50px; float: left;\">" +
	"<img id=\"" + cntnr.name + ".wait\" src=\"../images/wait2.gif\" alt=\"wait\" style=\"display: none; height: 20px; width 40px; vertical-align: middle;\">"
	html = html +
			"<button type=\"button\" id=\"" + cntnr.name + ".logs\" style=\"display: block; font-size: 14px;\" onclick=\"showLogging('" + cntnr.id + "')\">Logs</button>"
	html = html +
	"</div>" +
	//"<div style=\"margin: 12px 20px; overflow: hidden; width: 120px; float: left;\">" + cntnr.id + "<br>" +
	//"</div>" +slices
	"<div style=\"margin: 12px 20px; overflow: hidden; width: 240px; float: left;\">" + cntnr.img + "<br>" +
	"</div>"
	fmt.Printf("cntnr.status = '%s'", cntnr.status);
	html = html +
		"<br style=\"clear: both;\">" +
		"</div>"
	return html
}
