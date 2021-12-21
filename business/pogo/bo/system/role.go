package systemBO

type RoleInfo struct {
	ID      uint   `json:"id"`
	Name    string `json:"name" binding:"required" required_err:"角色名称不能为空"`
	Sort    int    `json:"sort"`
	Remark  string `json:"remark"`
	MenuIds []uint `json:"menuIds"`
}
