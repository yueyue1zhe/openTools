package captchautil

type Request struct {
	CaptchaId    string `json:"captcha_id"`
	CaptchaValue string `json:"captcha_value"`
}
