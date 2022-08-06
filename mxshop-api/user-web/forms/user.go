package forms

type PassWordLoginForm struct {
	// 手机号码格式有规范可寻 自定义validator
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	//`"required,min=3,max=10"`中间务必不要加上空格
	PassWord  string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

type RegisterForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	PassWord  string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Code string `form:"code" json:"code" binding:"required,min=6,max=6"`
}
