<template>
  <div class="app-container">
    <el-card>

      <el-form :inline="true">
        <el-card>
          <el-form-item>
            <el-autocomplete clearable :fetch-suggestions="querySearch"  v-model="exam_name" placeholder="考试科目" style="width: 200px"></el-autocomplete>
          </el-form-item>

          <el-form-item>
            <div class="block">
              <el-date-picker v-model="start_time" type="datetime" placeholder="开始时间" format="YYYY/MM/DD hh:mm:ss"
                value-format="x" />
            </div>
          </el-form-item>

          <el-form-item>
            <div class="block">
              <el-date-picker v-model="end_time" type="datetime" placeholder="结束时间" format="YYYY/MM/DD hh:mm:ss"
                value-format="x" />
            </div>
          </el-form-item>
          <el-form-item><el-button type="primary" :icon="Search" @click="handleSearch">搜索</el-button></el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleAdd">添加</el-button>
          </el-form-item>
        </el-card>
      </el-form>
      <el-table :data="filtedExams" border highlight-current-row stripe style="width: 100%">
        <el-table-column prop="exam_id" label="编号" width="180"></el-table-column>
        <el-table-column prop="exam_name" label="考试科目" width="180"></el-table-column>
        <el-table-column prop="paper" label="试卷" />
        <el-table-column prop="place" label="考试地点" />
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
        <el-table-column label="操作" fixed="right" width="300">
          <template v-slot="{ row }">
            <el-button text bg type="success" @click="handleStart(row)">开始考试</el-button>
            <el-button text bg type="primary" @click="handleDetail(row)">详情</el-button>
            <el-button text bg type="danger" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible">
      <el-form ref="formRef" :model="examForm" :rules="rules" label-width="100px">
        <el-form-item label="考试编号" prop="exam_id" required>
          <el-input v-model="examForm.exam_id"></el-input>
        </el-form-item>
        <el-form-item label="考试科目" prop="exam_name">
          <el-input v-model.lazy="examForm.exam_name"></el-input>
        </el-form-item>
        <el-form-item label="考试地点" required>
          <el-input v-model="examForm.place" />
        </el-form-item>
        <el-form-item label="试卷" required>
          <el-input v-model="examForm.paper" />
        </el-form-item>
        <el-form-item label="考试时间" required>
          <el-col :span="11">
            <el-form-item prop="begin_time">
              <el-date-picker  v-model="examForm.begin_time" type="datetime" placeholder="开始时间" value-format="x"
                style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col class="text-center" :span="2">
            <span class="text-gray-500">-</span>
          </el-col>
          <el-col :span="11">
            <el-form-item prop="end_time">
              <el-date-picker v-model="examForm.end_time" type="datetime" placeholder="结束时间" value-format="x"
                style="width: 100%" />
            </el-form-item>
          </el-col>
        </el-form-item>

        <el-form-item label="备注">
          <el-input v-model="examForm.exam_desc" type="textarea" />
        </el-form-item>


        <el-form-item>
          <el-button type="primary" @click="addFormSubmit(formRef)">提交</el-button>
          <el-button @click="dialogVisible = false">取消</el-button>
        </el-form-item>
      </el-form>


    </el-dialog>

    <el-dialog v-model="startVisible">
      haha
    </el-dialog>

  </div>
</template>
  
<script setup lang="ts">
import { getExamListApi, createExamApi, queryExamApi, deleteExamApi } from '@/api/exam'
import type { Exam } from '@/api/exam/types'
import router from '@/router'
import { useUserStore } from "@/store/models/user"
import {Search} from '@element-plus/icons-vue'


const exams = ref<Exam[]>()
const exam_id = ref("")
const exam_name = ref('')
const start_time = ref(0)
const end_time = ref(0)
const userStore = useUserStore()

const startVisible = ref(false)
function handleStart(row: Exam) {
  startVisible.value = true
}

interface Item{
  value: string
}
const filtedExams = ref<Exam[]>()
const querySearch = (queryString: string, cb: any) => {
  const names = exams.value?.map(i => {
    return {value: i.exam_name} as Item
  })
  
  const res = queryString?
   names?.filter((name) => name.value.includes(queryString)):
   names
   cb(res)
}
function handleSearch(){
  filtedExams.value = exams.value?.filter(exam => {
      if(start_time.value != 0){
        return exam.begin_time >= start_time.value
      }
      return true
  }).filter(exam=>{
    if(end_time.value != 0) {
      return exam.end_time <= end_time.value
    }
    return true
  }).filter(exam=>{
    return exam.exam_name.includes(exam_name.value)
  })
}

function getExamList() {
  getExamListApi().then((res) => {
    exams.value = res.data
    filtedExams.value = res.data
  })
}

const role = userStore.role
const dialogVisible = ref(false)
function handleAdd() {

  dialogVisible.value = true
}

const formRef = ref()
const examForm = reactive({} as Exam)
onMounted(() => {

  getExamList()
})
const rules = {
  exam_name: [{ required: true, message: '请输入考试科目', trigger: 'blur' }],
  begin_time: [{ required: true, message: '请选择开始时间', trigger: 'blur' }],
  end_time: [{ required: true, message: '请选择结束时间', trigger: 'blur' }]
}

function addFormSubmit(form) {
  if (!form) return
  form.validate((valid) => {
    if (valid) {
      createExamApi(examForm).then(res => {
        getExamList()
      }).finally(()=>{
        dialogVisible.value = false
      })
    }
  })
}



const handleDelete = (row) => {
  console.log(row)
  deleteExamApi(row.exam_id).then(res => {
    getExamList()
  })
}
function handleDetail(row) {
  exam_id.value = row.exam_id
  console.log(exam_id.value)
  if (role == 'student') {
    router.push({
      path: `/exam/verify/${exam_id.value}`,
    })
    return
  }
  router.push({
    path: `/exam/detail/${exam_id.value}`,
  })

}


function formatTime(time) {
  return new Date(time).toLocaleString()
}

</script>



<style>
.demo-datetime-picker .block {
  padding: 30px 0;
  text-align: center;
  border-right: solid 1px var(--el-border-color);
  flex: 1;
}
</style>