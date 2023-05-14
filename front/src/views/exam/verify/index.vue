<template>
  <el-form :inline="true" class="search">
    <el-form-item>
      <el-input v-model="exam_id" placeholder="请输入考试ID" suffix-icon="el-icon-search" style="width: 200px"></el-input>
    </el-form-item>
    <el-form-item>
      <el-input v-model="student_id" placeholder="请输入学号" suffix-icon="el-icon-search" style="width: 200px"></el-input>
    </el-form-item>
    <el-form-item><el-button slot="append" @click="handleSearch">核验</el-button></el-form-item>
  </el-form>

  <el-card>
    <el-table :data="quesitons" style="width: 100%">
      <el-table-column label="问题" prop="question_id"></el-table-column>
      <el-table-column label="答案">
        <template #default="{ row }">
          {{ row.actions.answer }}
        </template>
      </el-table-column>
      <el-table-column label="分数" prop="score" ></el-table-column>
      <el-table-column label="评分人" prop="scored_by" ></el-table-column>
      <el-table-column label="评分时间" prop="scored_time"></el-table-column>
      <el-table-column label="操作记录">
        <template #default="scope">
          <el-table :data="scope.row.actions" style="width: 100%" :tree-props="{children: 'children', hasChildren: 'hasChildren'}" border>
            <el-table-column label="操作" prop="action_id"></el-table-column>
            <el-table-column label="答案" prop="answer"></el-table-column>
            <el-table-column label="操作时间" prop="action_time"></el-table-column>
          </el-table>
        </template>
      </el-table-column>

    </el-table>
  </el-card>
</template>
  
<script setup lang="ts">
import { ExamProcess, QuestionInfo, ActionInfo, ExamResult } from "@/api/exam/types";
import { verifyExamApi, showExamApi } from "@/api/exam";
const props = defineProps(['student_id', 'exam_id'])
const student_id = ref(props.student_id)
const exam_id = ref(props.student_id)
const quesitons = ref<Questions[]>()
const ok = ref(false)
onMounted(() => {
  handleVerify()
})

function handleSearch() {
  handleVerify()
}



function handleVerify() {
  getExamResult(exam_id.value, student_id.value)
}


function getExamResult(exam_id: string, student_id: string) {
  return new Promise((resolve, reject) => {
    showExamApi(exam_id, student_id).then((res) => {
      verifyExam(res.data)
      resolve(res.data)
    }).catch((err) => {
      reject(err)
    })
  })
}

interface Questions{
  question_id: string
  score: number
  scored_by: string
  scored_time: number
  actions: Actions
}

interface Actions{
  action_id: string
  answer: string
  action_time: number
  hasChildren?: boolean
  children?: Actions[]
}
function verifyExam(examResult: ExamResult) {

  return new Promise((resolve, reject) => {
    verifyExamApi(examResult).then((res) => {
      var q = new Array<Questions>(res.data.process.questions.length)
    
      res.data.process.questions.forEach((e: QuestionInfo, i) => {
        var actions = e.actions.slice(1)
        
        var a = new Array<Actions>(actions.length)
          actions.forEach((e,i)=> {
            a[i] = {
              action_id: e.action_id,
              answer: e.answer,
              action_time: e.action_time
            }
        })
        q[i] = {
          question_id: e.question_id,
          score: e.score,
          scored_by: e.scored_by,
          scored_time: e.scored_time,
          actions: {
            action_id: e.actions[0].action_id,
            answer: e.actions[0].answer,
            action_time: e.actions[0].action_time,
            hasChildren: true,
            children: actions
          }
        }
      });
      quesitons.value = q
      ok.value = res.data.ok
    
      
      resolve(res.data)
    }).catch((err) => {
      reject(err)
    })
  })
}
</script>