package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/client"
	"github.com/Kong/go-pdk/server"
)

const (
	Version    = "1.0.0"
	Priority   = 1000
	PluginName = "ip-tacker"
)

type Config struct {
	WhitelistIPs []string `json:"whitelist_ips"`
	BlacklistIPs []string `json:"blacklist_ips"`
}

func New() interface{} {
	return &Config{}
}
func Schema() map[string]interface{} {
	return map[string]interface{}{
		"whitelist_ips": map[string]interface{}{
			"type":     "array",
			"elements": map[string]interface{}{"type": "string"},
			"required": false,
		},
		"blacklist_ips": map[string]interface{}{
			"type":     "array",
			"elements": map[string]interface{}{"type": "string"},
			"required": false,
		},
	}
}

func (config *Config) Access(kong *pdk.PDK) error {
	clientIP, err := client.Client.GetIp(kong.Client)
	if err != nil {
		return fmt.Errorf("failed to get client IP: %v", err)
	}

	if len(config.BlacklistIPs) > 0 {
		if isIPInList(clientIP, config.BlacklistIPs) {
			kong.Response.SetStatus(403)
			kong.Response.SetHeader("Content-Type", "text/plain")
			kong.Response.Exit(http.StatusForbidden, []byte("IP is blacklisted"), nil)
			return nil
		}
	}
	if len(config.WhitelistIPs) > 0 {
		if !isIPInList(clientIP, config.WhitelistIPs) {
			kong.Response.SetStatus(403)
			kong.Response.SetHeader("Content-Type", "text/plain")
			kong.Response.Exit(http.StatusForbidden, []byte("IP is not whitelisted"), nil)
			return nil
		}
	}

	return nil
}

func isIPInList(clientIP string, ipList []string) bool {
	for _, ipStr := range ipList {
		if strings.Contains(ipStr, "/") {
			_, ipNet, err := net.ParseCIDR(ipStr)
			if err == nil && ipNet.Contains(net.ParseIP(clientIP)) {
				return true
			}
		} else {
			if net.ParseIP(ipStr) != nil && clientIP == ipStr {
				return true
			}
		}
	}
	return false
}

func main() {
	// Check if running in query mode
	if len(os.Args) > 1 && os.Args[1] == "-dump" {
		info := map[string]interface{}{
			"name":     PluginName,
			"version":  Version,
			"priority": Priority,
			"schema":   Schema(),
		}

		json, err := json.Marshal(info)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error encoding plugin info: %v", err)
			os.Exit(1)
		}
		fmt.Println(string(json))
		return
	}

	// Otherwise start the plugin server
	if err := server.StartServer(New, Version, Priority); err != nil {
		log.Printf("Error starting server for the plugin: %v,  %v", PluginName, err)
		panic(err)
	}
}
