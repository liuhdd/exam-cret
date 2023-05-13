import request from '@/utils/request'
import type { AxiosPromise } from 'axios'
import type { ExamResult, VerificationResult } from './types'

export function showExamApi(exam_id: string, student_id: string): AxiosPromise<ExamResult> {
    return request({
        url: '/exam/show',
        method: 'get',
        params: {exam_id, student_id}
    })
}

export function verifyExamApi(examResult: ExamResult): AxiosPromise<VerificationResult> {
    return request({
        url: '/exam/verify',
        method: 'post',
        data: examResult
    })
}