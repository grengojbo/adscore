package adscore

import (
	"testing"

	"github.com/jackc/pgx"
	. "github.com/smartystreets/goconvey/convey"
)

// var Base64={_keyStr:"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=",encode:function(e){var t="";var n,r,i,s,o,u,a;var f=0;e=Base64._utf8_encode(e);while(f<e.length){n=e.charCodeAt(f++);r=e.charCodeAt(f++);i=e.charCodeAt(f++);s=n>>2;o=(n&3)<<4|r>>4;u=(r&15)<<2|i>>6;a=i&63;if(isNaN(r)){u=a=64}else if(isNaN(i)){a=64}t=t+this._keyStr.charAt(s)+this._keyStr.charAt(o)+this._keyStr.charAt(u)+this._keyStr.charAt(a)}return t},decode:function(e){var t="";var n,r,i;var s,o,u,a;var f=0;e=e.replace(/[^A-Za-z0-9\+\/\=]/g,"");while(f<e.length){s=this._keyStr.indexOf(e.charAt(f++));o=this._keyStr.indexOf(e.charAt(f++));u=this._keyStr.indexOf(e.charAt(f++));a=this._keyStr.indexOf(e.charAt(f++));n=s<<2|o>>4;r=(o&15)<<4|u>>2;i=(u&3)<<6|a;t=t+String.fromCharCode(n);if(u!=64){t=t+String.fromCharCode(r)}if(a!=64){t=t+String.fromCharCode(i)}}t=Base64._utf8_decode(t);return t},_utf8_encode:function(e){e=e.replace(/\r\n/g,"\n");var t="";for(var n=0;n<e.length;n++){var r=e.charCodeAt(n);if(r<128){t+=String.fromCharCode(r)}else if(r>127&&r<2048){t+=String.fromCharCode(r>>6|192);t+=String.fromCharCode(r&63|128)}else{t+=String.fromCharCode(r>>12|224);t+=String.fromCharCode(r>>6&63|128);t+=String.fromCharCode(r&63|128)}}return t},_utf8_decode:function(e){var t="";var n=0;var r=c1=c2=0;while(n<e.length){r=e.charCodeAt(n);if(r<128){t+=String.fromCharCode(r);n++}else if(r>191&&r<224){c2=e.charCodeAt(n+1);t+=String.fromCharCode((r&31)<<6|c2&63);n+=2}else{c2=e.charCodeAt(n+1);c3=e.charCodeAt(n+2);t+=String.fromCharCode((r&15)<<12|(c2&63)<<6|c3&63);n+=3}}return t}}
//

// func init() {
//   _, file, _, _ := runtime.Caller(1)
//   apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
//   beego.TestBeegoInit(apppath)
// }

var (
	paramsYes         = "MTkyLjE2OC44OC4zM3wwMTIzLjQ1NjcuODlBQi5DREVG"
	paramsYesPort     = "MTkyLjE2OC44OC4zMzo2NTc4MHwwMTIzLjQ1NjcuODlBQi5DREVG"
	paramsNo          = "MTkyLjE2OC44OC4zM3wwMTIzLjQ1NjcuODlBQi5ERROR"
	paramsIpv6        = "OjoxfDAxMjMuNDU2Ny44OUFCLkNERUY="
	paramsNoIpv6      = "Wzo6MV18MDEyMy40NTY3Ljg5QUIuQ0RFRg=="
	paramsYesIpv6Port = "Wzo6MV06NTU1Njd8MDEyMy40NTY3Ljg5QUIuQ0RFRg=="
	paramsNone        = "none"
	paramsNoMac       = "MTkyLjE2OC44OC4zM3wwMTIzLjQ1NjcuODlBQi5DREVGMTE="
	paramsNoIp        = "MTkyLjE2OC44OC4zMzQ0fDAxMjMuNDU2Ny44OUFCLkNERUY="
	clientMac         = pgx.NullString{String: "01:23:45:67:89:ab:cd:ef", Valid: true}
	clientIp          = pgx.NullString{String: "192.168.88.33", Valid: true}
	clientIpV6        = pgx.NullString{String: "::1", Valid: true}
	srcIp             = "192.168.88.33"
	srcNoIp           = "192.168.88.257"
	srcIpV6           = "::1"
	srcNoIpV6         = "[::1]"
	srcIpPort         = "192.168.88.33:64657"
	srcIpV6Port       = "[::1]:64657"
)

func TestParseParams(t *testing.T) {
	Convey("Test Parsing Base64 encode client IP and Mac\n", t, func() {
		Convey("Valid client IpV4 and MAC", func() {
			ip, ipv4, mac := ParseParams(paramsYes)

			So(ip.String, ShouldEqual, clientIp.String)
			So(ip.Valid, ShouldBeTrue)
			So(ipv4, ShouldBeTrue)
			So(mac.String, ShouldEqual, clientMac.String)
			So(mac.Valid, ShouldBeTrue)

			ip, ipv4, mac = ParseParams(paramsYesPort)

			So(ip.String, ShouldEqual, clientIp.String)
			So(ip.Valid, ShouldBeTrue)
			So(ipv4, ShouldBeTrue)
			So(mac.String, ShouldEqual, clientMac.String)
			So(mac.Valid, ShouldBeTrue)
		})
		Convey("Valid client IpV6 and MAC", func() {
			ip, ipv4, mac := ParseParams(paramsIpv6)

			So(ip.String, ShouldEqual, clientIpV6.String)
			So(ip.Valid, ShouldBeTrue)
			So(ipv4, ShouldBeFalse)
			So(mac.String, ShouldEqual, clientMac.String)
			So(mac.Valid, ShouldBeTrue)

			ip, ipv4, mac = ParseParams(paramsYesIpv6Port)

			So(ip.String, ShouldEqual, clientIpV6.String)
			So(ip.Valid, ShouldBeTrue)
			So(ipv4, ShouldBeFalse)
			So(mac.String, ShouldEqual, clientMac.String)
			So(mac.Valid, ShouldBeTrue)
		})
		Convey("Valid only IP", func() {
			ip, ipv4, mac := ParseParams(paramsNoMac)

			So(ipv4, ShouldBeTrue)
			So(ip.Valid, ShouldBeTrue)
			So(mac.Valid, ShouldBeFalse)
			So(mac.String, ShouldEqual, "")
			So(ip.String, ShouldEqual, clientIp.String)
		})
		Convey("Valid only Mac", func() {
			ip, ipv4, mac := ParseParams(paramsNoIp)

			So(ipv4, ShouldBeFalse)
			So(ip.Valid, ShouldBeFalse)
			So(mac.Valid, ShouldBeTrue)
			So(mac.String, ShouldEqual, clientMac.String)
			So(ip.String, ShouldEqual, "")
		})
		Convey("Is not valid params", func() {
			ip, ipv4, mac := ParseParams(paramsNo)

			So(ipv4, ShouldBeTrue)
			So(ip.Valid, ShouldBeTrue)
			So(mac.Valid, ShouldBeFalse)
			So(mac.String, ShouldEqual, "")
			So(ip.String, ShouldEqual, clientIp.String)

			ip, ipv4, mac = ParseParams(paramsNone)

			So(ipv4, ShouldBeFalse)
			So(ip.Valid, ShouldBeFalse)
			So(mac.Valid, ShouldBeFalse)
			So(mac.String, ShouldEqual, "")
			So(ip.String, ShouldEqual, "")

			ip, ipv4, mac = ParseParams(paramsNoIpv6)

			So(ip.String, ShouldEqual, "")
			So(ip.Valid, ShouldBeFalse)
			So(ipv4, ShouldBeFalse)
			So(mac.String, ShouldEqual, clientMac.String)
			So(mac.Valid, ShouldBeTrue)

		})
	})

	Convey("Test client IP \n", t, func() {
		Convey("Valid client IP", func() {
			ip, ipv4 := GetIP(srcIp)
			So(ip.String, ShouldEqual, clientIp.String)
			So(ip.Valid, ShouldBeTrue)
			So(ipv4, ShouldBeTrue)

			ip, ipv4 = GetIP(srcIpPort)
			So(ip.String, ShouldEqual, clientIp.String)
			So(ip.Valid, ShouldBeTrue)
			So(ipv4, ShouldBeTrue)

			ip, ipv4 = GetIP(srcIpV6)
			So(ip.String, ShouldEqual, clientIpV6.String)
			So(ip.Valid, ShouldBeTrue)
			So(ipv4, ShouldBeFalse)

			ip, ipv4 = GetIP(srcIpV6Port)
			So(ip.String, ShouldEqual, clientIpV6.String)
			So(ip.Valid, ShouldBeTrue)
			So(ipv4, ShouldBeFalse)

		})

		Convey("Is not valid client IP", func() {
			ip, ipv4 := GetIP(srcNoIp)
			So(ip.String, ShouldEqual, "")
			So(ip.Valid, ShouldBeFalse)
			So(ipv4, ShouldBeFalse)

			ip, ipv4 = GetIP(srcNoIpV6)
			So(ip.String, ShouldEqual, "")
			So(ip.Valid, ShouldBeFalse)
			So(ipv4, ShouldBeFalse)

		})
	})
}
