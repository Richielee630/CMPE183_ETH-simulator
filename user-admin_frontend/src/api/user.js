import request from '@/utils/request'

export function login(params) {
  return request({
    url: '/api/login',
    method: 'post',
    params
  })
}
export function register(params) {
  return request({
    url: '/api/register',
    method: 'post',
    params
  })
}
export function getInfo(params) {
  return request({
    url: '/api/system/userinfo',
    method: 'post',
    params
  })
}
export function recharge(params) {
  return request({
    url: '/api/system/recharge',
    method: 'post',
    params
  })
}
export function transfer(params) {
  return request({
    url: '/api/system/transfer',
    method: 'post',
    params
  })
}
export function gettx(params) {
  return request({
    url: '/api/system/gettx',
    method: 'post',
    params
  })
}
export function setContracts(params) {
  return request({
    url: '/api/system/set',
    method: 'post',
    params
  })
}
export function getContracts(params) {
  return request({
    url: '/api/system/get',
    method: 'post',
    params
  })
}
export function logout() {
  return request({
    url: '/user/logout',
    method: 'post'
  })
}
