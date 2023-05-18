import request from '@/utils/request'
import type { AxiosPromise } from 'axios'


export function uploadApi(file) {
    return request({
        url: '/system/resume',
        method: 'post',
        headers: {
            'Content-Type': 'multipart/form-data'
        },
        data: file
    })
}

export function confimReume() {
    return request({
        url: '/system/resume/confirm',
        method: 'post',
    })
}

export function backupApi() {
    return request({
        url: '/system/backup',
        method: 'get',
    })
}