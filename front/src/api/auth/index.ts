import request from '@/utils/request'
import type { AxiosPromise } from 'axios'
import type { LoginData, LoginResult } from './types'

export function loginApi(data : LoginData): AxiosPromise<LoginResult> {
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