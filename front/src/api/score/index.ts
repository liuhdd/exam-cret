import request from "@/utils/request"
import type { AxiosPromise } from 'axios'
import type { Score } from "./types"

export function queryQuestionScore(exam_id: string, student_id: string, question_id: string) :AxiosPromise<Score> {
    return request({
        url: "score/query",
        method: 'get',
        params: [exam_id, student_id, question_id]
    })
}
