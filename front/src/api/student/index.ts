import request from "@/utils/request"
import type { AxiosPromise } from 'axios'
import type { Student } from "./types"


export function getStudentByIdApi(student_id: string) :AxiosPromise<Student> {
    return request({
        url: `student/${student_id}`,
        method: 'get',
    })
}

export function getStudentByNameApi(name: string) :AxiosPromise<Student> {
    return request({
        url: `student/query`,
        method: 'get',
        params: name
    })
}

export function createStudentApi(student: Student) :AxiosPromise<Student> {
    return request({
        url: `student/create`,
        method: 'post',
        data: student
    })
}

export function updateStudentApi(student: Student) :AxiosPromise<Student> {
    return request({
        url: `student/update`,
        method: 'post',
        data: student
    })
}

export function deleteStudentApi(student_id: string) :AxiosPromise<Student> {
    return request({
        url: `student/${student_id}`,
        method: 'delete',
    })
}

export function getAllStudentApi() :AxiosPromise<Student[]> {
    return request({
        url: `student/list`,
        method: 'get',
    })
}