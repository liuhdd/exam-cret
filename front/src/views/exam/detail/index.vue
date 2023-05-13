<template>
  <el-card>
    <el-form :inline="true" class="search">
      <el-form-item>
        <el-input v-model="exam_id" placeholder="请输入考试ID" suffix-icon="el-icon-search" style="width: 200px"></el-input>
      </el-form-item>
      <el-form-item>
        <el-input v-model="student_id" placeholder="请输入学号" suffix-icon="el-icon-search" style="width: 200px"></el-input>
      </el-form-item>
      <el-form-item><el-button slot="append" @click="handleSearch">搜索</el-button></el-form-item>
    </el-form>

    <el-divider></el-divider>
    <el-form :inline="true" class="filter">
      <el-form-item>
        考试: {{ examResult.exam_id }}
      </el-form-item>
      <el-form-item>
       学号:  {{ examResult.student_id }}
      </el-form-item>
      <el-form-item>
        <el-input v-model="question_id" placeholder="请输入题号" suffix-icon="el-icon-search" style="width: 200px"></el-input>
      </el-form-item>
    </el-form>
    <el-divider></el-divider>
    <el-table :data="filteredQuestionResults" highlight-current-row border>

      <el-table-column prop="question_id" label="题目"></el-table-column>
      <el-table-column prop="answer" label="答案"></el-table-column>
      <el-table-column prop="score" label="得分"></el-table-column>
      <el-table-column label="操作" fixed="right" width="220">
        <template #default="scope">
          <el-button type="primary" link size="small" @click="handleCheck(scope.row)"><i-ep-delete />查看详情</el-button>
          <el-button type="primary" link size="small" @click="handleVerify(scope.row)"><i-ep-edit />验证</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
  <question-info-dialog :currentExamResult="examResult" :question_id="question_id" :showDialog="showDialog" />
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ExamResult, QuestionResult } from "@/api/exam/types"
import QuestionInfoDialog from './components/QestionInfoDialog/index.vue'
import { showExamApi } from '@/api/exam/index'
const props = defineProps(['examResult', 'student_id', 'exam_id'])
const student_id = ref(props.student_id)
const exam_id = ref(props.student_id)

function handleSearch() {
  showExam(exam_id.value, student_id.value)
}

const question_id = ref("")
watch(question_id, (newVal) => {
  if (newVal) {
    handleFilter()
  }
})
const questionResult = ref<QuestionResult>()
const examResult = ref<ExamResult>({
  exam_id: "1",
  student_id: "1",
  questions: [
    {
      question_id: "1",
      answer: "1",
      score: 1
    },
    {
      question_id: "2",
      answer: "2",
      score: 2
    }
  ]
})


function showExam(exam_id: string, student_id: string) {
  return new Promise((resolve, reject) => {
    showExamApi(exam_id, student_id).then((res) => {
      examResult.value = res.data
      resolve(res.data)
    }).catch((err) => {
      reject(err)
    })
  })
}
const filteredQuestionResults = computed(() => {
  return examResult.value.questions.filter((questionResult: QuestionResult) =>
    questionResult.question_id.includes(question_id.value)
  )
})

const handleFilter = () => {
  filteredQuestionResults.value
}



const showDialog = ref(false)
function handleCheck(row: QuestionResult) {
  questionResult.value = row
  showDialog.value = true
}


function handleVerify (row: QuestionResult) {
  questionResult.value = row
}

</script>