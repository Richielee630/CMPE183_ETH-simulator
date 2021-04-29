import Cookies from 'js-cookie'

const TokenKey = 'vue_admin_template_token'
const UserId = 'vue_admin_template_user_id'
const UserType = 'vue_admin_template_user_type'

export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}

export function getUserId() {
  return Cookies.get(UserId)
}

export function setUserId(userId) {
  return Cookies.set(UserId, userId)
}
export function removeUserId() {
  return Cookies.remove(UserId)
}

export function getUserType() {
  return Cookies.get(UserType)
}

export function setUserType(userType) {
  return Cookies.set(UserType, userType)
}
export function removeUserType() {
  return Cookies.remove(UserType)
}
