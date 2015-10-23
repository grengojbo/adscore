package adscore

import (
	"net"
	"strings"

	"encoding/base64"

	"github.com/jackc/pgx"
)

// Parsing Base64 encode client IP and Mac
func ParseParams(params string) (clientIp pgx.NullString, ipV4 bool, clientMac pgx.NullString) {
	ipV4 = false
	clientIp = pgx.NullString{Valid: false}
	clientMac = pgx.NullString{Valid: false}
	b, _ := base64.StdEncoding.DecodeString(params)
	s := strings.SplitN(string(b), "|", 2)
	if len(s) != 2 {
		return clientIp, ipV4, clientMac
	}

	clientIp, ipV4 = GetIP(s[0])
	umac, err := net.ParseMAC(s[1])
	if err == nil {
		clientMac = pgx.NullString{String: umac.String(), Valid: true}
	}
	return clientIp, ipV4, clientMac
}

// Return valid IP
func GetIP(s string) (clientIp pgx.NullString, ipV4 bool) {
	ipV4 = false
	clientIp = pgx.NullString{Valid: false}

	ip, _, err := net.SplitHostPort(s)
	if err == nil {
		clientIp = pgx.NullString{String: ip, Valid: true}
		if net.ParseIP(ip).To4() != nil {
			ipV4 = true
		}
	} else {
		ip := net.ParseIP(s)
		if ip != nil {
			clientIp = pgx.NullString{String: ip.String(), Valid: true}
		}
		if ip.To4() != nil {
			ipV4 = true
		}
	}
	return clientIp, ipV4
}
