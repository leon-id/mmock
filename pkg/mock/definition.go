package mock

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

type Values map[string][]string

type Cookies map[string]string

type HttpHeaders struct {
	Headers Values  `json:"headers"`
	Cookies Cookies `json:"cookies"`
}

type Request struct {
	Scheme                string `json:"scheme"`
	Host                  string `json:"host"`
	Port                  string `json:"port"`
	Method                string `json:"method"`
	Path                  string `json:"path"`
	QueryStringParameters Values `json:"queryStringParameters"`
	Fragment              string `json:"fragment"`
	HttpHeaders
	Body string `json:"body"`
}

type Response struct {
	StatusCode int `json:"statusCode"`
	HttpHeaders
	Body string `json:"body"`
}

type Callback struct {
	Delay  Delay  `json:"delay"`
	Method string `json:"method"`
	Url    string `json:"url"`
	HttpHeaders
	Body    string `json:"body"`
	Timeout Delay  `json:"timeout"`
}

type Scenario struct {
	Name          string   `json:"name"`
	RequiredState []string `json:"requiredState"`
	NewState      string   `json:"newState"`
}

type Delay struct {
	time.Duration
}

func (d *Delay) UnmarshalJSON(data []byte) (err error) {
	var (
		v interface{}
		s string
	)
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v.(type) {
	case float64:
		s = fmt.Sprintf("%ds", int(v.(float64)))
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println("! DEPRECATION NOTICE:                                        !")
		fmt.Println("! Please use a time unit (m,s,ms) to define the delay value. !")
		fmt.Println("! Ex: \"delay\":\"1s\" instead \"delay\":1                   !")
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	case string:
		s = v.(string)
	case interface{}:
		map_, _ := v.(map[string]interface{})
		s = fmt.Sprintf("%dns", int(map_["Duration"].(float64)))
		// log.Printf("interface found! %s:, casted as %s", v, s)
	default:
		log.Printf("%s: invalid value for delay, setting 0sec delay instead of %v", v, reflect.TypeOf(v))
		s = "0s"
	}

	d.Duration, err = time.ParseDuration(s)
	return err
}

type Control struct {
	Priority     int      `json:"priority"`
	Delay        Delay    `json:"delay"`
	Crazy        bool     `json:"crazy"`
	Scenario     Scenario `json:"scenario"`
	ProxyBaseURL string   `json:"proxyBaseURL"`
	WebHookURL   string   `json:"webHookURL"`
}

//Definition contains the user mock config
type Definition struct {
	URI         string
	Description string   `json:"description"`
	Request     Request  `json:"request"`
	Response    Response `json:"response"`
	Callback    Callback `json:"callback"`
	Control     Control  `json:"control"`
}
