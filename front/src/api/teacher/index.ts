import request from "@/utils/request"
import type { AxiosPromise } from 'axios'
import type { Teacher } from "./types"

export function getTeacherByIdApi(teacher_id: string) :AxiosPromise<Teacher> {
    return request({
        url: `teacher/${teacher_id}`,
        method: 'get',
    })
}

export function getTeacherByNameApi(name: string) :AxiosPromise<Teacher> {
    return request({
        url: `teacher/query`,
        method: 'get',
        params: name
    })
}

export function createTeacherApi(teacher: Teacher) :AxiosPromise<Teacher> {
    return request({
        url: `teacher/create`,
        method: 'post',
        data: teacher
    })
}

export function updateTeacherApi(teacher: Teacher) :AxiosPromise<Teacher> {
    return request({
        url: `teacher/update`,
        method: 'post',
        data: teacher
    })
}

export function deleteTeacherApi(teacher_id: string) :AxiosPromise<Teacher> {
    return request({
        url: `teacher/${teacher_id}`,
        method: 'delete',
    })
}

export function getAllTeacherApi(): AxiosPromise<Teacher[]> {
    return request({
        url: `teacher/list`,
        method: 'get',
    })
}