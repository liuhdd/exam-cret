<template>
  <div class="app-container">
  <el-card>
    <el-form :inline="true" class="search">
      <el-form-item v-if="role != 'student'">
        <el-input v-model="student_id" placeholder="请输入学号" suffix-icon="el-icon-search" style="width: 200px"></el-input>
      </el-form-item>
      <el-form-item v-if="role != 'student'">
        <el-input v-model="student_name" placeholder="请输入学号" suffix-icon="el-icon-search" style="width: 200px"></el-input>
      </el-form-item>
      <el-form-item>
        <el-input v-model="exam_name" placeholder="请输入考试科目" suffix-icon="el-icon-search" style="width: 200px"></el-input>
      </el-form-item>
      <el-form-item><el-button type="primary" :icon="Search" @click="handleSearch">搜索</el-button></el-form-item>
    </el-form>

    <el-divider></el-divider>
    <el-table :data="filterGrades" highlight-current-row border>
      <el-table-column label="学号" prop="student_id"></el-table-column>
      <el-table-column label="姓名" prop="student_name"></el-table-column>
      <el-table-column label="考试场次编号" prop="exam_id"></el-table-column>
      <el-table-column label="考试科目" prop="exam_name"></el-table-column>
      <el-table-column label="成绩" prop="grade"></el-table-column>
      <el-table-column label="开始时间" width="180">
          <template v-slot="{ row }">
            {{ formatTime(row.begin_time) }}
          </template>
        </el-table-column>
        <el-table-column label="结束时间" width="180">
          <template v-slot="{ row }">
            {{ formatTime(row.begin_time) }}
          </template>
        </el-table-column>
      <el-table-column label="操作" fixed="right" width="220">
        <template #default="scope">
          <el-button type="primary" @click="handleCheck(scope.row)">查看详情</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
</div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { Grade} from "@/api/exam/types"
import { queryGradeApi, getStudentGrade } from '@/api/exam/index'
import { useUserStore } from '@/store/models/user'
import router from '@/router'
import { Search } from '@element-plus/icons-vue'

const prop = defineProps(['exam_id'])
const student_id = ref('student111' || useUserStore().username)
const exam_id = ref(prop.exam_id || 'exam111')
const grades = ref<Grade[]>()
const role = useUserStore().role
const exam_name = ref('')
const student_name = ref('')
const filterGrades = ref<Grade[]>()
function handleSearch() {
  filterGrades.value = grades.value?.filter(item => {
    return item.exam_name.includes(exam_name.value) && 
    item.student_id.includes(student_id.value) &&
    item.student_name.includes(student_name.value)
  })
}

onMounted(() => {
  showGrades()
})
function showGrades() {

  queryGradeApi({
    exam_id: exam_id.value,
    student_id: student_id.value,
  } as Grade).then(res => {
    grades.value = res.data
    filterGrades.value = res.data
  })
}

function handleCheck(row) {
  router.push({
    path: '/exam/verify',
    query: {
      exam_id: row.exam_id,
      student_id: row.student_id
    }
  })
}

function formatTime(time) {
  return new Date(time).toLocaleString()
}

</script>