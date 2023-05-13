<template>
  <div>

    <el-form :inline="true" class="search">
      <el-form-item>
        <el-input v-model="student_id" placeholder="请输入学生ID" suffix-icon="el-icon-search" style="width: 200px"></el-input>
      </el-form-item>
      <el-form-item>
        <el-input v-model="exam_id" placeholder="请输入考试ID" suffix-icon="el-icon-search" style="width: 200px"></el-input>
      </el-form-item>
      <el-form-item><el-button slot="append" @click="handleFilter">搜索</el-button></el-form-item>
    </el-form>

    <el-table :data="filteredExamResults" highlight-current-row border>
      <el-table-column prop="exam_id" label="Exam ID"></el-table-column>
      <el-table-column prop="student_id" label="Student ID"></el-table-column>

      <el-table-column label="操作" fixed="right" width="220">
        <template #default="scope">
          <el-button type="primary" link size="small" @click="handleCheck(scope.row)"><i-ep-delete />查看详情</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
  
<script lang="ts">

import { ExamResult, QuestionResult } from "@/api/exam/types"
import { useRouter, useRoute } from 'vue-router'

export default {
  name: "ExamResultsTable",
  props: {
    examResultsData: {
      type: Array<ExamResult>,
      required: true,
    },
  },
  setup(props) {
    const student_id = ref("");
    const exam_id = ref("");
    const examResults = ref(props.examResultsData);
    // 使用 computed 计算筛选后的数据
    const filteredExamResults = computed(() => {
      return examResults.value.filter((examResult: ExamResult) =>
        examResult.student_id.includes(student_id.value)
        && examResult.exam_id.includes(exam_id.value)
      );
    });

    const handleCheck = (row: ExamResult) => {
      const router = useRouter();

    }

    const handleFilter = () => {
      filteredExamResults.value;
    };
    return {
      examResults,
      handleCheck,
      handleFilter,
      filteredExamResults,
      student_id,
      exam_id,
    };
  },
};
</script>