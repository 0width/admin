import request from '@/utils/request'

export function menuList() {
  return request({
    url: '/system/menu/list',
    method: 'get'
  })
}

// 查询菜单详细
export function getMenu(menuId) {
  return request({
    url: '/system/menu/' + menuId,
    method: 'get'
  })
}
// 新增菜单
export function addMenu(data) {
  return request({
    url: '/system/menu/add',
    method: 'post',
    data: data
  })
}

// 修改菜单
export function editMenu(data) {
  return request({
    url: '/system/menu/edit',
    method: 'put',
    data: data
  })
}

// 删除菜单
export function delMenu(menuId) {
  return request({
    url: '/system/menu/' + menuId,
    method: 'delete'
  })
}
