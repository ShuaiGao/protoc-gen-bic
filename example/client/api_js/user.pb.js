// Code generated by protoc-gen-bic. DO NOT EDIT.
// versions: 1.4.2-old

import request from '@/utils/axiosReq'

/**
 * 获取用户列表
 * @param { Object } params            - RequestUsers 
 * @param { ?number } params.page      - 页码
 * @param { ?number } params.pageSize  - 每页数量
 * @return { Promise } Promise object - ResponseUsers
 */
export function GetUsers(params) {
  return request({
    url: "/users/v1/",
    method: "GET",
    params,
  })
} 

/**
 * 添加用户
 * @param { Object } data            - User 
 * @param { ?number } data.id        - 主键ID
 * @param { ?string } data.username  - 用户名
 * @param { ?string } data.email     - 邮箱
 * @return { Promise } Promise object - ResponseNil
 */
export function PostUsers(data) {
  return request({
    url: "/users/v1/",
    method: "POST",
    data,
  })
} 

/**
 * 获取用户详情
 * @Param { number } id - 
 * @param { Object } params    - RequestNil 
 * @return { Promise } Promise object - User
 */
export function GetUserInfo(id, params) {
  return request({
    url: "/users/v1/" + id + "/",
    method: "GET",
    params,
  })
} 

/**
 * 用户下载
 * @Param { number } id - 
 * @param { Object } params    - RequestNil 
 * @return { Promise } Promise object - ResponseNil
 */
export function GetUserInfoDownload(id, params) {
  return request({
    url: "/users/v1/" + id + "/",
    method: "GET",
    params,
  })
} 

/**
 * 空参数示例1
 * @param { Object } params    - CommonNil 
 * @return { Promise } Promise object - User
 */
export function GetCommonNil() {
  return request({
    url: "/users/v1/common/nil/",
    method: "GET",
  })
} 

/**
 * 空参数示例2
 * @param { Object } params        - RequestPostEmpty 
 * @param { ?string } params.name  - 
 * @return { Promise } Promise object - Empty
 */
export function PostEmpty(params) {
  return request({
    url: "/users/v1/empty/",
    method: "GET",
    params,
  })
} 

