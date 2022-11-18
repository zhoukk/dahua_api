package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/zhoukk/dahua_api"
)

func main() {
	var host string
	var user string
	var pass string
	flag.StringVar(&host, "h", "192.168.1.108", "ip camera host")
	flag.StringVar(&user, "u", "admin", "ip camera username")
	flag.StringVar(&pass, "p", "123456", "ip camera password")
	flag.Parse()

	c := dahua_api.NewClient(host, user, pass)

	arg := url.Values{}
	ret := make(map[string]string)
	err := errors.New("")

	err = c.CGI("magicBox.cgi", "getSerialNo", nil, ret)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("%+v\n", ret)
	}

	arg = url.Values{}
	ret = make(map[string]string)
	arg.Add("name", "NAS")
	err = c.CGI("configManager.cgi", "getConfig", arg, ret)
	if err != nil {
		log.Println(err)
	} else {
		for k, v := range ret {
			if strings.HasPrefix(k, "table.NAS[0]") {
				log.Println(k, "=", v)
			}
		}
	}

	arg = url.Values{}
	ret = make(map[string]string)
	arg.Add("NAS[0].Name", "FTP1")
	arg.Add("NAS[0].Enable", "true")
	arg.Add("NAS[0].Protocol", "FTP")
	arg.Add("NAS[0].Address", "192.168.1.123")
	arg.Add("NAS[0].Port", "21")
	arg.Add("NAS[0].UserName", "KFTPD")
	arg.Add("NAS[0].Password", "123456")
	arg.Add("NAS[0].Directory", "KFTPD")
	err = c.CGI("configManager.cgi", "setConfig", arg, ret)
	if err != nil {
		log.Println(err)
	} else {
		for k, v := range ret {
			log.Println(k, "=", v)
		}
	}

	arg = url.Values{}
	ret = make(map[string]string)
	arg.Add("name", "RecordStoragePoint")
	err = c.CGI("configManager.cgi", "getConfig", arg, ret)
	if err != nil {
		log.Println(err)
	} else {
		for k, v := range ret {
			if strings.HasPrefix(k, "table.RecordStoragePoint[0]") {
				log.Println(k, "=", v)
			}
		}
	}

	arg = url.Values{}
	ret = make(map[string]string)
	arg.Add("RecordStoragePoint[0].TimingSnapShot.FTP", "true")
	err = c.CGI("configManager.cgi", "setConfig", arg, ret)
	if err != nil {
		log.Println(err)
	} else {
		for k, v := range ret {
			log.Println(k, "=", v)
		}
	}

	arg = url.Values{}
	ret = make(map[string]string)
	arg.Add("name", "NTP")
	err = c.CGI("configManager.cgi", "getConfig", arg, ret)
	if err != nil {
		log.Println(err)
	} else {
		for k, v := range ret {
			if strings.HasPrefix(k, "table.NTP") {
				log.Println(k, "=", v)
			}
		}
	}

	arg = url.Values{}
	ret = make(map[string]string)
	arg.Add("NTP.Enable", "true")
	arg.Add("NTP.Address", "clock.isc.org")
	arg.Add("NTP.Port", "123")
	arg.Add("NTP.TimeZone", "13")
	arg.Add("NTP.UpdatePeriod", "10")

	err = c.CGI("configManager.cgi", "setConfig", arg, ret)
	if err != nil {
		log.Println(err)
	} else {
		for k, v := range ret {
			log.Println(k, "=", v)
		}
	}

	arg = url.Values{}
	ret = make(map[string]string)
	arg.Add("name", "Snap")
	err = c.CGI("configManager.cgi", "getConfig", arg, ret)
	if err != nil {
		log.Println(err)
	} else {
		for k, v := range ret {
			if strings.HasPrefix(k, "table.Snap[0]") {
				log.Println(k, "=", v)
			}
		}
	}

	arg = url.Values{}
	ret = make(map[string]string)
	arg.Add("Snap[0].HolidayEnable", "false")
	for wd := 0; wd <= 7; wd++ {
		arg.Add(fmt.Sprintf("Snap[0].TimeSection[%d][0]", wd), "1 00:00:00-23:59:59")
	}

	err = c.CGI("configManager.cgi", "setConfig", arg, ret)
	if err != nil {
		log.Println(err)
	} else {
		for k, v := range ret {
			log.Println(k, "=", v)
		}
	}

	arg = url.Values{}
	ret = make(map[string]string)
	arg.Add("name", "Encode[0].SnapFormat")
	err = c.CGI("configManager.cgi", "getConfig", arg, ret)
	if err != nil {
		log.Println(err)
	} else {
		for k, v := range ret {
			if strings.HasPrefix(k, "table.Encode[0].SnapFormat[0].Video") {
				log.Println(k, "=", v)
			}
		}
	}

	arg = url.Values{}
	ret = make(map[string]string)
	arg.Add("Encode[0].SnapFormat[0].Video.FPS", fmt.Sprintf("%f", 1.0/3))

	err = c.CGI("configManager.cgi", "setConfig", arg, ret)
	if err != nil {
		log.Println(err)
	} else {
		for k, v := range ret {
			log.Println(k, "=", v)
		}
	}

}
