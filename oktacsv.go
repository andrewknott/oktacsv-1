package main

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"net/url"
	//"strconv"
	"os"
	//"strings"
	"strings"
)

var OktaEvent []struct {
	EventID   string `json:"eventId"`
	SessionID string `json:"sessionId"`
	RequestID string `json:"requestId"`
	Published string `json:"published"`
	Action    struct {
			  Message    string `json:"message"`
			  Categories []string `json:"categories"`
			  ObjectType string `json:"objectType"`
			  RequestURI string `json:"requestUri"`
		  } `json:"action"`
	Actors    []struct {
		ID          string `json:"id"`
		DisplayName string `json:"displayName"`
		Login       string `json:"login,omitempty"`
		ObjectType  string `json:"objectType"`
		IPAddress   string `json:"ipAddress,omitempty"`
	} `json:"actors"`
	Targets   []struct {
		ID          string `json:"id"`
		DisplayName string `json:"displayName"`
		Login       string `json:"login"`
		ObjectType  string `json:"objectType"`
	} `json:"targets"`
}

func main() {

	if len(os.Args) != 1 {

		OktaOrg := strings.ToLower(os.Args[1]) //"https://hardCodeYourOktaOrg.oktapreview.com"
		OktaKey := os.Args[2] //"Your key"

		if (  ! strings.HasPrefix(OktaOrg, "https://")) {
			fmt.Fprintln(os.Stderr, "Your Okta Org should begin with https://")
			os.Exit(0);
		}

		fmt.Fprintln(os.Stderr, "OktaCSV by Patrick McDowell pmcdowell@okta.com")

		fmt.Fprintln(os.Stderr, "   ___  _   _         ___ _____   __")
		fmt.Fprintln(os.Stderr, "  / _ \\| |_| |_ __ _ / __/ __\\ \\ / /")
		fmt.Fprintln(os.Stderr, " | (_) | / /  _/ _` | (__\\__ \\\\ V /")
		fmt.Fprintln(os.Stderr, "  \\___/|_\\_\\\\__\\__,_|\\___|___/ \\_/")

		fmt.Fprintln(os.Stderr, "\nOktaCSV is checking to see what time it is in OktaLand, and starting to follow the System Log")
		fmt.Fprintln(os.Stderr, "This can take a few seconds, but it is worth the wait")

		lastEvent := ReturnTimeLastEvent(OktaOrg, OktaKey)

		i := 1
		for {
			i += 1
			duration := time.Second * 1
			time.Sleep(duration)
			events := GetOktaEvent(OktaOrg, OktaKey, "filter=published%20gt%20%22" + lastEvent + "%22")
			OktaEvent = nil
			json.Unmarshal([]byte (events), &OktaEvent)

			if (OktaEvent != nil && len (OktaEvent) !=0  ) {
				for v := range OktaEvent {
					fmt.Println(OktaEvent[v].Published + "," + OktaEvent[v].Action.Message)

				}

				lastEvent = OktaEvent[len(OktaEvent) - 1].Published
				OktaEvent = nil

			}
		}
	} else {
		fmt.Println("Usage: oktacsv OktaOrg OktaKey")
		fmt.Println("   ___  _   _         ___ _____   __")
		fmt.Println("  / _ \\| |_| |_ __ _ / __/ __\\ \\ / /")
		fmt.Println(" | (_) | / /  _/ _` | (__\\__ \\\\ V /")
		fmt.Println("  \\___/|_\\_\\\\__\\__,_|\\___|___/ \\_/\n")
		fmt.Println("OktaCSV by Patrick McDowell pmcdowell@okta.com")
	}

}

func ReturnTimeLastEvent(OktaOrg string, OktaKey string) string {

	url := OktaOrg + "/api/v1/events?limit=1&filter=published%20gt%20%222017-12-03T05%3A20%3A48.000Z%22"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "SSWS " + OktaKey)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "690b5379-d5f0-3cff-b1a9-a6a89bc40af4")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	date := string(res.Header.Get("Date"))

	t, err := time.Parse(time.RFC1123, date)

	if err != nil {
		fmt.Println("parse error", err.Error())
	}

	threeHours := time.Hour * 0
	newTime := t.Add(threeHours) // 7 hours actually

	returnString := newTime.Format("2006-01-02T15:04:05") + ".000Z"

	fmt.Fprintln(os.Stderr, "Wait for Events after this Published Date:" + returnString + ". Events take some time to hit the Event Log")

	return returnString
}

func GetOktaEvent(OktaOrg string, OktaKey string, arguments string) []byte {

	url := OktaOrg + "/api/v1/events?" + arguments

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "SSWS " + OktaKey)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "fcd54dc9-bd3b-bdbf-f99a-47272d773855")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body
}

func UrlEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}






