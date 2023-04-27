<template>
  <div>

    <el-form :inline="true" class="search">
      <el-form-item>
        <el-input v-model="searchText" placeholder="请输入学生ID" suffix-icon="el-icon-search" style="width: 200px"></el-input>
      </el-form-item>
      <el-form-item><el-button slot="append" @click="handleFilter">搜索</el-button></el-form-item>
    </el-form>

    <el-table :data="filteredExamResults" highlight-current-row border>
      <el-table-column prop="exam_id" label="Exam ID"></el-table-column>
      <el-table-column prop="student_id" label="Student ID"></el-table-column>
      <el-table-column label="Questions">
        <template #default="data">
          <el-table :data="data.row.questions">
            <el-table-column prop="question_id" label="Question ID"></el-table-column>
            <el-table-column prop="answer" label="Answer"></el-table-column>
            <el-table-column prop="score" label="Score"></el-table-column>
          </el-table>
        </template>
      </el-table-column>

      <el-table-column label="操作" fixed="right" width="220">
        <template #default="scope">
          <!-- <el-button type="primary" link size="small" @click="verify(scope.row.id)"
            v-hasPerm="['sys:user:edit']"><i-ep-edit />验证</el-button> -->
          <el-button type="primary" link size="small" @click="handleCheck(scope.row.id)"><i-ep-delete />查看详情</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
  
<script lang="ts">

import { ExamResult, QuestionResult } from "@/api/exam/types"
export default {
  name: "ExamResultsTable",
  props: {
    examResultsData: {
      type: Array<ExamResult>,
      required: true,
    },
  },
  setup(props) {
    const searchText = ref("");
    const examResults = ref(props.examResultsData);
    // 使用 computed 计算筛选后的数据
    const filteredExamResults = computed(() => {
      return examResults.value.filter((examResult: ExamResult) =>
        examResult.student_id.includes(searchText.value)
      );
    });

    const handleCheck = (id: number) => {
      //todo impliment
    }

    const handleFilter = () => {
      filteredExamResults.value;
    };
    return {
      examResults,
      handleCheck,
      handleFilter,
      filteredExamResults,
      searchText,
    };
  },
};
</script>