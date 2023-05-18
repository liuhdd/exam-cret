import request from '@/utils/request'
import type { AxiosPromise } from 'axios'
import type { ExamResult, Grade, Exam } from './types'

export function showExamApi(exam_id: string, student_id: string): AxiosPromise<ExamResult> {
    return request({
        url: '/exam/show',
        method: 'get',
        params: {exam_id: exam_id, student_id: student_id}
    })
}

export function queryGradeApi(grade: Grade): AxiosPromise<Grade[]> {
    return request({
        url: '/exam/grades',
        method: 'post',
        data: grade
    })
}
export function getStudentGrade(student_id: string): AxiosPromise<Grade[]> {
    return request({
        url: `/exam/grade/${student_id}`,
        method: 'get',
    })
}

export function getExamListApi(): AxiosPromise<Exam[]> {
    return request({
        url: '/exam/list',
        method: 'get'
    })
}
export function createExamApi(exam: Exam): AxiosPromise<Exam> {
    return request({
        url: '/exam/create',
        method: 'post',
        data: exam
    })
}

export function deleteExamApi(exam_id: string): AxiosPromise<Exam> {
    return request({
        url: `/exam/${exam_id}`,
        method: 'delete'
    })
}

export function queryExamApi(exam: Exam): AxiosPromise<Exam[]> {
    return request({
        url: '/exam/query',
        method: 'post',
        data: exam
    })
}
