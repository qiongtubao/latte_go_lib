package xhr

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	lib "github.com/qiongtubao/latte_go_lib"
)

type Xhr struct {
	typeName string
	url      string
	data     map[string]interface{}
	headers  map[string]string
	client   *http.Client
}

func (xhr Xhr) Query(querys map[string]interface{}) Xhr {
	for key, value := range querys {
		xhr.data[key] = value
	}
	return xhr
}

func (xhr Xhr) Send() (string, error) {
	var array = []string{}
	var i = 0
	if xhr.typeName == http.MethodGet {
		for k, v := range xhr.data {
			str, err := lib.ToJson(v)
			if err == nil {
				array[i] = k + "=" + str
			}
		}
	}
	var reader io.Reader
	if xhr.typeName == http.MethodPost {
		bytesData, err := json.Marshal(xhr.data)
		if err == nil {
			reader = bytes.NewReader(bytesData)
		}

	}
	req, _ := http.NewRequest(xhr.typeName, xhr.url+strings.Join(array, "&"), reader)

	for key, value := range xhr.headers {
		req.Header.Set(key, value)
	}
	res, err := xhr.client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
func (xhr Xhr) Set(key string, value string) Xhr {
	xhr.headers[key] = value
	return xhr
}

func (xhr Xhr) Use(callback func(Xhr) Xhr) Xhr {
	return callback(xhr)
}
func Get(url string) (Xhr, error) {
	// data := map[string]string{}
	// headers := map[string]string{}

	client := http.Client{}
	data := map[string]interface{}{}
	headers := map[string]string{}
	xhr := Xhr{http.MethodGet, url, data, headers, &client}
	return xhr, nil
}

func Post(url string) Xhr {
	// data := map[string]string{}
	// headers := map[string]string{}
	client := http.Client{}
	data := map[string]interface{}{}
	headers := map[string]string{}
	xhr := Xhr{http.MethodPost, url, data, headers, &client}
	return xhr
}
