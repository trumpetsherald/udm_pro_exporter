package unifiapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const namespace = "udm"

type APIBroker struct {
	URL      string
	Username string
	Password string

	csrfToken string
	client    *http.Client
}

func NewAPIBroker(url string, username string, password string) *APIBroker {
	api := APIBroker{URL: url, Username: username, Password: password}
	var err error

	jar, err := cookiejar.New(nil)

	if err != nil {
		log.Fatal(err)
	}

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	api.client = &http.Client{
		Transport: customTransport,
		Jar:       jar,
	}

	return &api
}

func (api *APIBroker) VerifyConnectivity() bool {
	if !api.IsLoggedIn() {
		loginStatus, err := api.Login()
		if err != nil || (loginStatus != 200 && loginStatus != 422) {
			log.Fatal(err)
		}
	} else {
		log.Println("Already logged in.")
	}

	logoutStatus, err := api.Logout()
	if err != nil || logoutStatus != 200 {
		log.Fatal(err)
	}

	return true
}

func (api *APIBroker) Login() (int, error) {
	loginMap, _ := json.Marshal(map[string]string{
		"username": api.Username,
		"password": api.Password,
	})

	payload := bytes.NewBuffer(loginMap)

	req, err := http.NewRequest("POST", api.URL+"api/auth/login", payload)
	if err != nil {
		log.Println("Error building login request.")
		return -1, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := api.client.Do(req)
	if err != nil {
		log.Println("Error executing login request.")
		return -1, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		log.Println("Login Request succeeded.")
	} else if resp.StatusCode == 403 {
		var loginError LoginAPIErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&loginError)
		if err != nil {
			log.Println("Error decoding stats response in JSON.")
		}

		if loginError.Message == "Invalid CSRF Token" {
			resp.StatusCode = 422 // Unprocessable Entity
			log.Println("Already logged in.")
		} else {
			log.Println("Login Request failed.")
		}
	}

	api.csrfToken = resp.Header.Get("X-Csrf-Token")

	return resp.StatusCode, nil
}

func (api *APIBroker) Logout() (int, error) {
	req, err := http.NewRequest("POST", api.URL+"api/auth/logout", nil)
	if err != nil {
		log.Println("Error building login request.")
		return -1, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Csrf-Token", api.csrfToken)

	resp, err := api.client.Do(req)
	if err != nil {
		log.Println("Error executing logout request.")
		return -1, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("Logout Request failed.")
	} else {
		log.Println("Logout Request succeeded.")
	}

	return resp.StatusCode, nil
}

func (api *APIBroker) IsLoggedIn() bool {
	result := false
	deviceURL, err := url.Parse(api.URL)
	if err != nil {
		log.Println("Error parsing URL.")
	}
	cookies := api.client.Jar.Cookies(deviceURL)

	if len(cookies) > 0 {
		for _, cookie := range cookies {
			if cookie.Name == "TOKEN" && len(cookie.Value) > 0 {
				result = true
				break
			}
		}
	}

	return result
}

func (api *APIBroker) Status() string {
	return "Status!!"
}

func (api *APIBroker) Health() string {
	return "Health"
}

func (api *APIBroker) Stats() (StatsAPIResponse, error) {
	var stats StatsAPIResponse

	req, err := http.NewRequest("GET", api.URL+"proxy/network/api/s/default/stat/sta", nil)
	if err != nil {
		log.Println("Error building stats request.")
		return stats, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := api.client.Do(req)
	if err != nil {
		log.Println("Error executing stats request.")
		return stats, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("Stats request failed.") //todo fix this
	}

	err = json.NewDecoder(resp.Body).Decode(&stats)
	if err != nil {
		log.Println("Error decoding stats response in JSON.")
		return stats, err
	}

	return stats, nil
}

func (api *APIBroker) SysInfo() (SysInfoAPIResponse, error) {
	var sysInfo SysInfoAPIResponse

	req, err := http.NewRequest("GET", api.URL+"proxy/network/api/s/default/stat/sysinfo", nil)
	if err != nil {
		log.Println("Error building sysinfo request.")
		return sysInfo, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := api.client.Do(req)
	if err != nil {
		log.Println("Error executing sysinfo request.")
		return sysInfo, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("Sysinfo request failed.") //todo fix this
	}

	err = json.NewDecoder(resp.Body).Decode(&sysInfo)
	if err != nil {
		log.Println("Error decoding sysinfo response in JSON.")
		return sysInfo, err
	}

	return sysInfo, nil
}

func (api *APIBroker) Device() (DeviceAPIResponse, error) {
	var device DeviceAPIResponse

	req, err := http.NewRequest("GET", api.URL+"proxy/network/api/s/default/stat/device", nil)
	if err != nil {
		log.Println("Error building device request.")
		return device, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := api.client.Do(req)
	if err != nil {
		log.Println("Error executing device request.")
		return device, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("Device request failed.") //todo fix this
	}

	err = json.NewDecoder(resp.Body).Decode(&device)
	if err != nil {
		log.Println("Error decoding device response in JSON.")
		return device, err
	}

	return device, nil
}

func (api *APIBroker) DeviceBasic() string {
	return "DeviceBasic!!"
}

func (api *APIBroker) CountryCodes() string {
	return "CountryCodes!!"
}
