package Models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type RecaptchaResponse struct {
	Success bool `json:"success"`
}

// [TODO:]Check this...
var captchaPrivateKey string = os.Getenv("GOOGLE_CAPTCHA_SECRET")

const recaptchaServerName = "https://www.google.com/recaptcha/api/siteverify"

func ChaptchaCheck(response string) (r RecaptchaResponse) {
	resp, err := http.PostForm(recaptchaServerName, url.Values{"secret": {captchaPrivateKey}, "response": {response}})
	defer resp.Body.Close()
	if err != nil {
		r.Success = false
		return r
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.Success = false
		return r
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		r.Success = false
		return r
	}
	return r
}

func CaptchaConfirm(response string) (result bool) {
	result = ChaptchaCheck(response).Success
	return
}
