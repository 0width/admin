package systemBO

type MenuCommon struct {
	Name      string `json:"name" binding:"required" required_err:"菜单(组件名称)不能为空"`
	Title     string `json:"title" binding:"required" required_err:"页面标题不能为空"`
	Icon      string `json:"icon"`
	Path      string `json:"path" binding:"required" required_err:"Path不能为空"`
	Redirect  string `json:"redirect"`
	Query     string `json:"query"`
	Component string `json:"component" binding:"required_if=Type 1" err:"组件不能为空"`
	Order     int    `json:"order"`
	Visible   int    `json:"visible"`
	ParentId  uint   `json:"parent_id"`
	Remark    string `json:"remark"`
	Perm      string `json:"perm" binding:"required_unless=Type 0" err:"权限不能为空"`
	Type      int    `json:"type"`
}

type AddMenuInfo struct {
	MenuCommon
}

type EditMenuInfo struct {
	MenuCommon
	ID uint `json:"id" binding:"required" required_err:"菜单ID不能为空"`
}
