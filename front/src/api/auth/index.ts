import request from '@/utils/request'
import type { AxiosPromise } from 'axios'
import type { LoginData, LoginResult } from './types'

export function loginApi(data : LoginData): AxiosPromise<LoginResult> {
    return request({
        url: '/api/auth/login',
        method: 'post',
        params: data
    })
}

export function logoutApi() {
    return request({
      url: '/api/v1/auth/logout',
      method: 'post'
    });
}