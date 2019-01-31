package main

// env GOOS=linux GOARCH=amd64 go build

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// GondulData is
type GondulData struct {
	Source string    `json:"src"`
	Meta   MetaData  `json:"metadata"`
	Lease  LeaseInfo `json:"data"`
}

// MetaData contains some data :D
type MetaData struct {
	Server string `json:"server"`
}

// LeaseInfo is the data sendt to Gondul
type LeaseInfo struct {
	ClientIP   string    `json:"clientip"`
	ClientMac  string    `json:"clientmac"`
	ClientName string    `json:"clientname,omitempty"`
	LeaseTime  int       `json:"leasetime,omitempty"`
	CircuitID  string    `json:"circuitid,omitempty"`
	Time       time.Time `json:"time"`
}

// Flags
var (
	clientIP    = flag.String("ip", "", "Client IP")
	clientMac   = flag.String("mac", "", "Client MAC")
	clientName  = flag.String("clientname", "", "Client Name")
	leaseTime   = flag.Int("lease", 0, "Lease time")
	circuitID   = flag.String("circuit", "", "Circuit ID from Option 82")
	gondulURL   = flag.String("gondul", "http://gondul.tg.lol/api/write/dhcp", "Gondul URL")
	debugFlag   = flag.Bool("d", false, "Print debug info")
	hostname, _ = os.Hostname()
)

func main() {
	// Parse flags
	flag.Parse()

	// Exit if missing IP or MAC
	if *clientIP == "" || *clientMac == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Create data struct
	d := GondulData{
		Source: "dhcp",
		Meta: MetaData{
			Server: hostname,
		},
		Lease: LeaseInfo{
			ClientIP:   *clientIP,
			ClientMac:  validateMac(*clientMac),
			ClientName: *clientName,
			LeaseTime:  *leaseTime,
			Time:       time.Now(),
		},
	}

	// Send data to Gondul
	postData(d, *gondulURL)

	// Save to logfile for local debug
	saveLog(d, *gondulURL)

	if *debugFlag == true {
		debug(d)
	}

	os.Exit(0)
}

// postData send LeaseInfo to the Gondul API
func postData(data GondulData, endpoint string) {

	d, err := json.Marshal(data)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(d))
	req.Header.Set("Content-Type", "application/json")

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		os.Exit(1)
	}

}

// validateMac fixes a bug where dhcpd cuts leading 0's from the mac in each octet
func validateMac(mac string) string {
	m := strings.Split(mac, ":")
	for i := range m {
		if len(m[i]) < 2 {
			m[i] = fmt.Sprintf("0%s", m[i])
		}
	}
	validMac := strings.Join(m, ":")
	return validMac
}

// saveLog is for local debugging
func saveLog(d GondulData, endpoint string) {
	logLine := fmt.Sprintf(
		"Host: %v API: %v IP: %v MAC: %v NAME: %v CIRCUIT: %v",
		hostname,
		endpoint,
		d.Lease.ClientIP,
		d.Lease.ClientMac,
		d.Lease.ClientName,
		d.Lease.CircuitID,
	)
	f, err := os.OpenFile("/tmp/gondul-dhcp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(logLine + "\n")
	f.Close()
}

func debug(d GondulData) {
	json, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(json))
}
