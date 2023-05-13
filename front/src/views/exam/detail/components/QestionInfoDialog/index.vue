<template>
    <el-dialog :visible.sync="showDialog" title="Questions 详细信息" align-center>
      <div v-if="currentExamResult">
        <p>Exam ID: {{ currentExamResult.exam_id }}</p>
        <p>Student ID: {{ currentExamResult.student_id }}</p>
        <ul>
          <li v-for="(question, index) in currentExamResult.questions" :key="index">
            <p>Question ID: {{ question.question_id }}</p>
            <p>Answer: {{ question.answer }}</p>
            <p>Score: {{ question.score }}</p>
          </li>
        </ul>
      </div>
    </el-dialog>
</template>

<script setup lang="ts">
    const props = defineProps(['currentExamResult'])
    const showDialog = ref(false)
    watch(() => props.currentExamResult, (newVal) => {
        if (newVal) {
            showDialog.value = true
        }
    })
</script>