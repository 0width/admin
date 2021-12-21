package systemDTO

type RoleInfo struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Sort   int    `json:"sort"`
	Remark string `json:"remark"`
}
