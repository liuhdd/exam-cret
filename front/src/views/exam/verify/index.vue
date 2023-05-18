<template>
  <div class="app-container">
  <el-card>
    <el-form :inline="true" class="search">
      <el-form-item>
        <el-input v-model="exam_id" placeholder="请输入考试ID" suffix-icon="el-icon-search" style="width: 200px"></el-input>
      </el-form-item>
      <el-form-item>
        <el-input v-if="role != 'student'" v-model="student_id" placeholder="请输入学号" suffix-icon="el-icon-search" style="width: 200px"></el-input>
      </el-form-item>
      <el-form-item><el-button slot="append" @click="handleSearch">搜索</el-button></el-form-item>
    </el-form>

    <el-divider></el-divider>
    <el-descriptions border> 
      <el-descriptions-item label="考试科目">{{ examResult?.exam_name }}</el-descriptions-item>
      <el-descriptions-item label="开始时间">{{ formatTime(examResult?.begin_time) }}</el-descriptions-item>
      <el-descriptions-item label="结束时间">{{ formatTime(examResult?.end_time) }}</el-descriptions-item>
      <el-descriptions-item label="分数">{{ examResult?.grade }}</el-descriptions-item>
    </el-descriptions> 
    <el-divider></el-divider>
    <el-table :data="questions" highlight-current-row border>
      <el-table-column prop="question_id" label="题号"></el-table-column>
      <el-table-column prop="answer" label="回答"></el-table-column>
      <el-table-column prop="score" label="得分"></el-table-column>
      <el-table-column label="操作" fixed="right" width="220">
        <template #default="scope">
          <el-button type="primary"  @click="handleCheck(scope.row)">查看详情</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
  </div>
  <question-info-dialog v-model:display="show" v-model:question="question" v-model:exam_id="exam_id" v-model:student_id="student_id" />
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ExamResult, QuestionResult } from "@/api/exam/types"
import QuestionInfoDialog from './components/QestionInfoDialog/index.vue'
import { showExamApi } from '@/api/exam/index'
import { useUserStore } from '@/store/models/user'

const prop = defineProps(['exam_id', "student_id"])
const student_id = ref(prop.student_id || '1000000011')
const exam_id = ref(prop.exam_id || 'EX001')
function handleSearch() {
  showExam()
}
const show = ref(false)
const role = useUserStore().role
onMounted(() => {

  if(role == "student"){
    student_id.value = useUserStore().username
  }
  showExam()
})
const question_id = ref("")
watch(question_id, (newVal) => {
  if (newVal) {
    handleFilter()
  }
})


const examResult = ref<ExamResult>()
const questions = ref()
function showExam() {
  showExamApi(exam_id.value, student_id.value).then((res) => {
    examResult.value = res.data
    questions.value = examResult.value.questions
  })
}
const filteredQuestionResults = computed(() => {
  return questions.value.filter((questionResult: QuestionResult) =>
    questionResult.question_id.includes(question_id.value)
  )
})
const handleFilter = () => {
  filteredQuestionResults.value
}


const question = ref()
function handleCheck(row: QuestionResult) {
  question.value = row
  show.value = true
}

function formatTime(time) {
  return new Date(time).toLocaleString()
}

</script>