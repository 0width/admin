import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/system/user/login',
    method: 'post',
    data
  })
}

export function userMenuList() {
  return request({
    url: '/system/user/menus',
    method: 'get'
  })
}

export function menuList() {
  return request({
    url: '/system/menu/list',
    method: 'get'
  })
}

export function getInfo() {
  return request({
    url: '/system/user/info',
    method: 'get'
  })
}

export function logout() {
  return request({
    url: '/system/user/logout',
    method: 'post'
  })
}
