import { login, logout, getInfo, menuList } from '@/api/user'
import { getToken, setToken, removeToken } from '@/utils/auth'
import { resetRouter } from '@/router'
import Layout from '@/layout'
import { loadComponent } from '@/utils/aync_import'

const getDefaultState = () => {
  return {
    token: getToken(),
    name: '',
    avatar: '',
    menus: false
  }
}

const state = getDefaultState()

const mutations = {
  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState())
  },
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  },
  SET_MENU: (state) => {
    state.menus = true
  }
}

const table = () => import('@/views/table/index')
const m = new Map()
m.set('table', table)

const actions = {
  // user login
  login({ commit }, userInfo) {
    const { username, password } = userInfo
    return new Promise((resolve, reject) => {
      login({ username: username.trim(), password: password }).then(response => {
        const { data } = response
        commit('SET_TOKEN', data.token)
        setToken(data.token)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },
  menuList({ commit, state }) {
    return new Promise((resolve, reject) => {
      menuList().then(response => {
        const { data } = response
        if (!data) {
          return reject('获取菜单失败')
        }
        const routes = []
        while (data.length > 0) {
          const c = data.shift()
          let cComponent = null
          if (c.parent_id === 0) {
            cComponent = Layout
          } else {
            cComponent = loadComponent(c.component)
          }
          const r = {
            id: c.id,
            path: c.path,
            hidden: c.visible,
            component: cComponent,
            redirect: c.redirect,
            name: c.name,
            meta: {
              title: c.title,
              icon: c.icon
            }
          }
          if (c.parent_id === 0) {
            routes.push(r)
          } else {
            routes.forEach((v, i, arr) => {
              if (c.parent_id === v.id) {
                if (!v.children) {
                  v['children'] = []
                }
                v.children.push(r)
              }
            })
          }
        }
        // 放入默认路由， 解决动态路由刷新的时候跳转404的问题
        routes.push({
          path: '*',
          redirect: '/404',
          hidden: true
        })
        resolve(routes)
        console.log(routes)
      })
    })
  },
  // get user info
  getInfo({ commit, state }) {
    return new Promise((resolve, reject) => {
      getInfo(state.token).then(response => {
        const { data } = response

        if (!data) {
          return reject('Verification failed, please Login again.')
        }

        const { name, avatar } = data

        commit('SET_NAME', name)
        commit('SET_AVATAR', avatar)
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },

  // user logout
  logout({ commit, state }) {
    return new Promise((resolve, reject) => {
      logout(state.token).then(() => {
        removeToken() // must remove  token  first
        resetRouter()
        commit('RESET_STATE')
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // remove token
  resetToken({ commit }) {
    return new Promise(resolve => {
      removeToken() // must remove  token  first
      commit('RESET_STATE')
      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

