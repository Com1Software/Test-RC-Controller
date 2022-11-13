package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {
	xip := "localhost"
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			xip = fmt.Sprintf("%s", ipv4)
		}
	}
	fmt.Printf("Host IP : %s\n", xip)
	http.HandleFunc("/", RC_Controller)
	http.HandleFunc("/register/", Register)
	http.HandleFunc("/connect/", Connect)
	http.HandleFunc("/disconnect/", Disconnect)
	http.HandleFunc("/test/", TestServer)
	http.ListenAndServe(":8080", nil)
}

func TestServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "RC-Controller %s   %d \n", r.URL.Path[1:], len(r.URL.Path))
	xip := "localhost"
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			xip = fmt.Sprintf("%s", ipv4)
		}
	}
	fmt.Fprintf(w, "Host IP : %s\n", xip)
	fmt.Fprintf(w, "RawQuery : %s\n", r.URL.RawQuery)

}

func RC_Controller(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "RC-Controller  \n")
	xip := "localhost"
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			xip = fmt.Sprintf("%s", ipv4)
		}
	}
	fmt.Fprintf(w, "Host IP : %s\n\n", xip)
	fmt.Fprintf(w, "/register?IPAddress \n")
	fmt.Fprintf(w, "/connect?IPAddress \n")
	fmt.Fprintf(w, "/disconnect?IPAddress \n")

}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Register Vehicle\n")

}

func Connect(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Connect Vehicle\n")

}

func Disconnect(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Disconnect Vehicle\n")

}
