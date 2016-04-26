/**********************************************************
 * Author        : nfalse
 * Email         : nfalse@163.com
 * Last modified : 2016-04-26 10:28
 * Filename      : zoomeye.go
 * Description   :
 * *******************************************************/
package zoomeye // zoomeye
import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const uagent string = "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.95 Safari/537.36 SE 2.X MetaSr 1.0"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Token struct {
	AccessToken string `json:"access_token"`
}
type HostAnswer struct {
	Matches []struct {
		Geoinfo struct {
			City struct {
				GeonameID int `json:"geoname_id"`
				Names     struct {
					ZhCn string `json:"zh-CN"`
					En   string `json:"en"`
				} `json:"names"`
			} `json:"city"`
			Country struct {
				GeonameID int    `json:"geoname_id"`
				Code      string `json:"code"`
				Names     struct {
					ZhCn string `json:"zh-CN"`
					En   string `json:"en"`
				} `json:"names"`
			} `json:"country"`
			Isp       string `json:"isp"`
			Continent struct {
				GeonameID int    `json:"geoname_id"`
				Code      string `json:"code"`
				Names     struct {
					ZhCn string `json:"zh-CN"`
					En   string `json:"en"`
				} `json:"names"`
			} `json:"continent"`
			Subdivisions struct {
				GeonameID int    `json:"geoname_id"`
				Code      string `json:"code"`
				Names     struct {
					ZhCn string `json:"zh-CN"`
					En   string `json:"en"`
				} `json:"names"`
			} `json:"subdivisions"`
			Location struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"location"`
			Organization string `json:"organization"`
			Aso          string `json:"aso"`
			Asn          int    `json:"asn"`
		} `json:"geoinfo"`
		IP       string `json:"ip"`
		Portinfo struct {
			Product   string `json:"product"`
			Hostname  string `json:"hostname"`
			Service   string `json:"service"`
			Os        string `json:"os"`
			Extrainfo string `json:"extrainfo"`
			Version   string `json:"version"`
			Device    string `json:"device"`
			Banner    string `json:"banner"`
			Port      int    `json:"port"`
		} `json:"portinfo"`
		Timestamp string `json:"timestamp"`
	} `json:"matches"`
	Facets struct {
		App []struct {
			Count    int    `json:"count"`
			App      string `json:"app"`
			Versions []struct {
				Count   int    `json:"count"`
				Version string `json:"version"`
			} `json:"versions"`
		} `json:"app"`
		Os []struct {
			Count int    `json:"count"`
			Os    string `json:"os"`
		} `json:"os"`
	} `json:"facets"`
	Total int `json:"total"`
}
type WebtAnswer struct {
	Matches []struct {
		Geoinfo struct {
			City struct {
				Names struct {
					ZhCn string `json:"zh-CN"`
					En   string `json:"en"`
				} `json:"names"`
			} `json:"city"`
			Asn      int `json:"asn"`
			Location struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"location"`
			Continent struct {
				Code  string `json:"code"`
				Names struct {
					ZhCn string `json:"zh-CN"`
					En   string `json:"en"`
				} `json:"names"`
			} `json:"continent"`
			Country struct {
				Code  string `json:"code"`
				Names struct {
					ZhCn string `json:"zh-CN"`
					En   string `json:"en"`
				} `json:"names"`
			} `json:"country"`
		} `json:"geoinfo"`
		CheckTime string   `json:"check_time"`
		Language  []string `json:"language"`
		Title     string   `json:"title"`
		IP        []string `json:"ip"`
		Plugin    []struct {
			Version string `json:"version"`
			Based   string `json:"based"`
			Name    string `json:"name"`
			Chinese string `json:"chinese"`
		} `json:"plugin"`
		Db []struct {
			Version interface{} `json:"version"`
			Name    string      `json:"name"`
			Chinese string      `json:"chinese"`
		} `json:"db"`
		Site     string `json:"site"`
		Headers  string `json:"headers"`
		Keywords string `json:"keywords"`
		Webapp   []struct {
			URL     string `json:"url"`
			Version string `json:"version"`
			Name    string `json:"name"`
			Chinese string `json:"chinese"`
		} `json:"webapp"`
		Domains     []string `json:"domains"`
		Description string   `json:"description"`
	} `json:"matches"`
	Total int `json:"total"`
}

func GetToken(url, user, password string) (token Token, err error) {
	var _user User
	_user.Username = user
	_user.Password = password
	jbody, err := json.Marshal(_user)
	if err != nil {
		return
	}
	ibody := bytes.NewBuffer([]byte(jbody))
	reqest, err := http.NewRequest("POST", url+"/user/login", ibody)
	reqest.Header.Set("User-Agent", uagent)
	client := &http.Client{}
	response, err := client.Do(reqest)
	defer response.Body.Close()
	if err != nil {
		return
	}
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &token)
	if err != nil {
		return
	}
	return
}
func ConditionGet(url, search_type, condition string, token Token) (result []byte, err error) {
	reqest, err := http.NewRequest("GET", url+"/"+search_type+"/search?"+condition, nil)
	if err != nil {
		return
	}
	reqest.Header.Set("Authorization", "JWT "+token.AccessToken)
	reqest.Header.Set("User-Agent", uagent)
	client := &http.Client{}
	response, err := client.Do(reqest)
	defer response.Body.Close()
	if err != nil {
		return
	}
	result, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	return
}
func HostGet(url, condition string, token Token) (answer HostAnswer, err error) {
	result, err := ConditionGet(url, "host", condition, token)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &answer)
	if err != nil {
		return
	}
	return
}
func WebGet(url, condition string, token Token) (answer WebtAnswer, err error) {
	result, err := ConditionGet(url, "web", condition, token)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &answer)
	if err != nil {
		return
	}
	return
}
