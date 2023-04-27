<template>
  <el-form :inline="true" class="search">
    <el-form-item>
      <el-input v-model="examId" placeholder="请输入考试ID" suffix-icon="el-icon-search" style="width: 200px"></el-input>
      <el-input v-model="studentId" placeholder="请输入学生ID" suffix-icon="el-icon-search" style="width: 200px"></el-input>
      <el-input v-model="questionId" placeholder="请输入题号" suffix-icon="el-icon-search" style="width: 200px"></el-input>
    </el-form-item>
    <el-form-item><el-button slot="append" @click="handleFilter">搜索</el-button></el-form-item>
  </el-form>
</template>

<script setup lang="ts">
import { ref } from 'vue';


const studentId = ref("")
const examId = ref("")
const questionId = ref("")

const handleFilter = () => {}



</script>