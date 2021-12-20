package systemDTO

type SystemUserInfoDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Sex      int    `json:"sex"`
	Avatar   string `json:"avatar"`
	Remark   string `json:"remark"`
}
