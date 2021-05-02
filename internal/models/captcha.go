package models

// Response object that contains all of our captcha attributes.
// swagger:response NewCaptcha
type Captcha struct {
	Image string `json:"image"`
	Code  string `json:"code"`
}
