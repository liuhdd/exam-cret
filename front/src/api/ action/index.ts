import request from "@/utils/request"
import type { AxiosPromise } from 'axios'
import type { Action } from "./types"
export function listActionsApi(exam_id: string, student_id: string, question_id: string): AxiosPromise<Action[]>{
    return request({
        url: '/action/question/'+exam_id+'/'+student_id+'/'+question_id,
        method: 'get'
    })
}

export function queryActionApi(action_id: string): AxiosPromise<Action> {
    return request({
        url: '/action/'+action_id,
        method: 'get'
    })
}

export function getActionsApi(exam_id: string, student_id:string): AxiosPromise<Action[]> {
    return request({
        url: '/action/list/'+ exam_id+'/'+student_id,
        method: 'get'
    })
}