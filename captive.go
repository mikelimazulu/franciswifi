package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/hsanjuan/go-captive"
)

func getLocalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Name == "wlan0" {
			addrs, err := iface.Addrs()
			if err != nil {
				return "", err
			}
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				if ip != nil && ip.To4() != nil {
					return ip.String(), nil
				}
			}
		}
	}
	return "", fmt.Errorf("could not determine local IP")
}

func loginHandler(r *http.Request) bool {
	return true
}

func main() {
	ip, err := getLocalIP()
	if err != nil {
		fmt.Println("Error getting local IP:", err)
		os.Exit(1)
	}

	proxy := &captive.Portal{
		LoginPath:           "/login",
		PortalDomain:        ip,
		AllowedBypassPortal: false,
		WebPath:             "staticContentFolder",
		LoginHandler:        loginHandler,
	}

	go func() {
		http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
			ok := loginHandler(r)
			if ok {
				w.WriteHeader(http.StatusAccepted)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
		})
		fs := http.FileServer(http.Dir("staticContentFolder"))
		http.Handle("/", fs)
		log.Fatal(http.ListenAndServe(":9080", nil))
	}()

	err = proxy.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
