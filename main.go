package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"
	"log"
	"strings"
	"time"
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Package struct {
	PackageCd string
	IsOption int
	PackageNm string
	PackagePrice string
}

type Color struct {
	SAPExterioColorCode string
	ExtColorLongDesc string
}

type Vehicle struct {
	Vin string
	ModelNm string
	TrimDesc string
	ModelCd string
	Price string
	ExteriorColorCd string
	InteriorColorCd string
	DrivetrainDesc string
	TransmissionDesc string
	TotalPackages int
	TotalOptions int
	Packages []Package
	Colors []Color
	PlannedDeliveryDate string
	InventoryStatus string
}

type DealerInfo struct {
	DealerCd string
	DealerNm string
	DealerEmail string
	DealerUrl string
	Address1 string
	Address2 string
	City string
	State string
	Zip string
	Phone string
	Fax string
	Region string
	Latitude float32
	Longitude float32
	Distance float32
	ShopperAssurance string
	IsPMADealer int
	Vehicles []Vehicle
}

type Data struct {
	ModelYear int
	YrSerCd string
	DealerInfo []DealerInfo
}

type Response struct {
	Status string
	Data []Data
}

func main() {
	for range time.Tick(time.Minute * 1) {
		fmt.Println("***************Running search**************")
        search()
		fmt.Println("***************Finished search**************")
    }
}

func search() {
	url := "https://www.hyundaiusa.com/var/hyundai/services/inventory/vehicleList.json?zip=20151&year=2022&model=Palisade&radius=100"
	req, err := http.NewRequest("GET", url, nil)
	if (err != nil) {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:100.0) Gecko/20100101 Firefox/100.0")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://www.hyundaiusa.com/us/en/inventory-search/vehicles-list?model=Palisade&year=2022")
	req.Header.Set("Cookie", os.ExpandEnv("check=true; mbox=PC#2db0d0e5186040b2b60860b333e1e83a.34_0#1717852103|session#c944696918124c6ca03b36d063402d03#1654608130; utag_main=v_id:01813b74cb370001b0c63105cc1905054003200f00942$_sn:2$_se:5$_ss:0$_st:1654609103190$vapi_domain:hyundaiusa.com$ses_id:1654606269912%3Bexp-session$_pn:3%3Bexp-session; _bamls_usid=7e8de4a2-edfc-42c4-8a13-26501f9e0fe6; _gcl_au=1.1.1961220274.1654559919; AMCV_C3BCE0154FA24300A4C98A1%40AdobeOrg=1585540135%7CMCIDTS%7C19150%7CvVersion%7C4.4.0%7CMCMID%7C82738761952502289264142552394239761200%7CMCAID%7CNONE%7CMCOPTOUT-1654613581s%7CNONE; __atuvc=7%7C23; __atssc=google%3B2; _ga=GA1.2.590258494.1654559921; _gid=GA1.2.447526977.1654559921; s_ppvl=h%253A%2520shopping%2520tools%253A%2520search%2520inventory%253A%2520results%253A%2520Palisade%2C34%2C34%2C565%2C1440%2C565%2C1440%2C900%2C2%2CP; s_ppv=h%253A%2520shopping%2520tools%253A%2520search%2520inventory%253A%2520results%253A%2520Palisade%2C28%2C24%2C565%2C1440%2C565%2C1440%2C900%2C2%2CP; _tt_enable_cookie=1; _ttp=2f112f51-939a-49eb-9f42-ceddd2cea2f8; _scid=c50128bf-8b32-4588-8f90-5384a855e91f; _pin_unauth=dWlkPVpEWTJPRFl5TXpjdE1XSXpPUzAwWTJVMkxUazFZbU10T0RaaVlXTmlNVEZtTnpaaw; AMCV_3C3BCE0154FA24300A4C98A1%40AdobeOrg=1585540135%7CMCMID%7C10196674858856126213137868773484345617%7CMCAID%7CNONE%7CMCOPTOUT-1654613470s%7CNONE%7CMCAAMLH-1655211070%7C7%7CMCAAMB-1655211070%7Cj8Odv6LonN4r3an7LhD3WZrU1bUpAkFkkiY1ncBR96t2PTI%7CvVersion%7C4.4.0; s_ecid=MCMID%7C82738761952502289264142552394239761200; _clck=155m3jj|1|f24|0; AMCVS_3C3BCE0154FA24300A4C98A1%40AdobeOrg=1; AMCVS_C3BCE0154FA24300A4C98A1%40AdobeOrg=1; _clsk=40stlb|1654607182904|4|1|l.clarity.ms/collect; c_m=www.google.comNatural%20Search; s_channelstack=%5B%5B%27Natural%2520Search%27%2C%271654559933701%27%5D%5D; s_cc=true; _hjSessionUser_1596276=eyJpZCI6IjllMTU3YWEzLTNiM2MtNWM1Mi04ZTE3LWM3YjUzMjdhOGQ2NyIsImNyZWF0ZWQiOjE2NTQ1NTk5MzM0MTQsImV4aXN0aW5nIjp0cnVlfQ==; OptanonConsent=isGpcEnabled=0&datestamp=Tue+Jun+07+2022+09%3A08%3A23+GMT-0400+(Eastern+Daylight+Time)&version=6.30.0&isIABGlobal=false&hosts=&consentId=5441ee58-4226-4ba3-8222-c53df8d734dd&interactionCount=2&landingPath=NotLandingPage&groups=C0001%3A1%2CC0002%3A1%2CC0003%3A1%2CC0004%3A1%2CBG18%3A1&AwaitingReconsent=false&geolocation=US%3BVA; _sctr=1|1654488000000; LPVID=EyMDBiNzY0M2VjNjQ2YWFm; LPSID-41916303=1Zb-b0-fR2qMzVgF9KpkQg; _aeaid=967fa051-fca4-4217-a219-411c42182094; ipe_s=4c9d3b22-fd04-dd3e-38a1-a9d19c224005; IPE_LandingTime=1654559935216; ipe.322.pageViewedCount=1; ipe.322.pageViewedDay=158; ipe_322_fov=%7B%22numberOfVisits%22%3A1%2C%22sessionId%22%3A%224c9d3b22-fd04-dd3e-38a1-a9d19c224005%22%2C%22expiry%22%3A%222022-07-06T23%3A58%3A55.217Z%22%2C%22lastVisit%22%3A%222022-06-07T12%3A53%3A03.706Z%22%7D; aelastsite=cPse%2F4ilq18K5PO0nmpeINtYqinYaz9cECulfv7mhMaUn0ylDJlpA0ixb3hz4I9T; aelreadersettings=%7B%22c_big%22%3A0%2C%22rg%22%3A0%2C%22memph%22%3A0%2C%22contrast_setting%22%3A0%2C%22colorshift_setting%22%3A0%2C%22text_size_setting%22%3A0%2C%22space_setting%22%3A0%2C%22font_setting%22%3A0%2C%22k%22%3A0%2C%22k_disable_default%22%3A0%2C%22hlt%22%3A0%2C%22disable_animations%22%3A0%2C%22display_alt_desc%22%3A0%7D; aeatstartmessage=true; s_sq=%5B%5BB%5D%5D; _dpm_id.5888=14719644-ca98-431a-b8a3-14a41df9bb92.1654560035.2.1654606382.1654560098.29cd953f-871c-4548-8237-dc57c188f426; OptanonAlertBoxClosed=2022-06-07T00:01:18.585Z; re-evaluation=true; mboxEdgeCluster=34; __atuvs=629f49bd3c6f0ec5002; s_ppn=h%3A%20shopping%20tools%3A%20search%20inventory%3A%20results%3A%20Palisade; _dpm_ses.5888=*; _hjIncludedInSessionSample=0; _hjSession_1596276=eyJpZCI6IjI2ZTE1MzZhLTdjMTktNDcyYy05OTY5LTQzYTNjOTM2NDFiNSIsImNyZWF0ZWQiOjE2NTQ2MDYyNzE1MzEsImluU2FtcGxlIjpmYWxzZX0=; _hjIncludedInPageviewSample=1; _hjAbsoluteSessionInProgress=0; _uetsid=9d181d20e5f411eca2837f5917b44efb; _uetvid=9d1841a0e5f411ecb83b09b74c24908a"))
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Te", "trailers")
	resp, err := http.DefaultClient.Do(req)
	if (err != nil) {
		log.Fatal(err)
	}

	// get response body
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// convert to struct
    var response Response
    err = json.Unmarshal(body, &response)
    if err != nil {
        log.Fatal(err)
    }

	// look for vehicles
	for _, data := range response.Data {
		for _, info := range data.DealerInfo {
			// filter vehicles down to only those that meet our condition
			var goodVehicles []Vehicle
			for _, vehicle := range info.Vehicles {
				if meetsCondition(vehicle) {
					goodVehicles = append(goodVehicles, vehicle)
				}
			}

			// details of each good vehicle
			// send a text for any vehicle that has not had their vin logged yet
			for _, vehicle := range goodVehicles {
				var body string
				body += "https://www.hyundaiusa.com/us/en/inventory-search/vehicles-list?model=Palisade&year=2022\n"
				body += "Name: " + info.DealerNm + "\n"
				body += "Distance: " + fmt.Sprintf("%f", info.Distance) + "\n"
				body += "Phone: " + info.Phone + "\n"

				v, _ := PrettyStruct(vehicle)
				body += v
				fmt.Println(body)

				// if the vin has not already been reported then send the message and log the vin
				if (!vinLogged(vehicle.Vin)) {
					sendText(body, os.Getenv("VEHICLE_SEARCH_TO_PHONE_NUMNBER"))
					writeVin(vehicle.Vin)
				} else {
					fmt.Println(vehicle.Vin, " already exists in log")
				}
			}
		}
	}
}

func writeVin(vin string) {
	vin += "\n"
	f, err := os.OpenFile("vins.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(vin); err != nil {
		log.Println(err)
	}
}

func vinLogged(vin string) bool {
	// read the whole file at once
    b, err := ioutil.ReadFile("vins.log")
    if err != nil {
        panic(err)
    }
    s := string(b)
    // //check whether s contains substring text
    return strings.Contains(s, vin)
}

func sendText(body string, to string) {
	from := os.Getenv("VEHICLE_SEARCH_FROM_PHONE_NUMNBER")

	accountSid := os.Getenv("TWILO_ACCOUNT_SID")
	authToken := os.Getenv("TWILO_AUTH_TOKEN")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(body)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}

func meetsCondition(vehicle Vehicle) bool {
	// WDN is for a beige interior
	if (vehicle.InteriorColorCd == "WDN" && vehicle.TrimDesc == "CALLIGRAPHY") {
		return true
	}

	return false;
}

func PrettyStruct(data interface{}) (string, error) {
    val, err := json.MarshalIndent(data, "", "    ")
    if err != nil {
        return "", err
    }
    return string(val), nil
}
