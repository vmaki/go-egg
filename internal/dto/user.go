package dto

type UserRegisterReq struct {
	Username   string `json:"username" binding:"required"`
	Phone      string `json:"phone" binding:"required,is-phone-exist"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Age        uint8  `json:"age" binding:"gte=1,lte=120"`
}
