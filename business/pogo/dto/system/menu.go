package systemDTO

type MenuInfo struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	Path      string `json:"path"`
	Redirect  string `json:"redirect"`
	Query     string `json:"query"`
	Component string `json:"component"`
	Order     int    `json:"order"`
	Status    int    `json:"status"`
	Visible   int    `json:"visible"`
	ParentId  uint   `json:"parent_id"`
	Remark    string `json:"remark"`
	Perm      string `json:"perm"`
	Type      int    `json:"type"`
}
