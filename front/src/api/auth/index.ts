import request from '@/utils/request'
import type { AxiosPromise } from 'axios'
import type { LoginData, LoginResult, User } from './types'

export function loginApi(data : LoginData) : AxiosPromise<LoginResult> {
    return request({
        url: '/user/login',
        method: 'post',
        data: data,
    })
}

export function logoutApi() {
    return request({
      url: '/user/logout',
      method: 'post'
    });
}

export function listUserApi(): AxiosPromise<User[]> {
    return request({
        url: '/user/list',
        method: 'get'
    })
}

export function deleteUserApi(username: string) {
    return request({
        url: `/user/${username}`,
        method: 'delete'
    })
}