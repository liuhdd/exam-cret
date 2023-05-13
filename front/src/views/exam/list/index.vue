<template>
  <div class="app-container">
    <exam-results-table ref="tab" :exam-results-data="examResultsData" />
  </div>
</template>
  
<script lang="ts">
import ExamResultsTable from "@/components/ExamResultsTable/index.vue"
import { reactive } from "vue";


export default {
  name: "App",
  components: {
    ExamResultsTable,
  },
  setup() {
    
    const examResultsData = reactive([
      {
        exam_id: "1",
        student_id: "A00001",
        questions: [
          {
            question_id: "1",
            answer: "A",
            score: 100,
          },
          {
            question_id: "2",
            answer: "B",
            score: 90,
          },
        ],
      },
      {
        exam_id: "2",
        student_id: "A00002",
        questions: [
          {
            question_id: "1",
            answer: "B",
            score: 80,
          },
          {
            question_id: "2",
            answer: "A",
            score: 95,
          },
        ],
      },
    ])
    
    return {
      examResultsData,
    };
  },
};
</script>