<template>
  <el-dialog v-model="showDialog" align-center :before-close="handleClose" width="80%">
    <el-descriptions title="题目详细信息" :column="4" border>
      <el-descriptions-item>
        <template #label>编号</template>
        {{ question.question_id }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>题目</template>
        {{ question.content }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>作答</template>
        {{ question.answer }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label><span>得分</span></template>
        {{ question.score }}
      </el-descriptions-item>
    </el-descriptions>

    <el-collapse>
      <el-collapse-item title="答题记录">
        <el-table :data="actions" border stripe style="width: 100%">
          <el-table-column prop="action_id" label="记录码" width="180"></el-table-column>
          <el-table-column prop="answer" label="答案" width="180"></el-table-column>
          <el-table-column sortable width="180" label="操作时间">
            <template v-slot="{row}">
              {{ formatTime(row.action_time) }}
            </template>
          </el-table-column>
        </el-table>
      </el-collapse-item>
    </el-collapse>
  </el-dialog>
</template>

<script setup lang="ts">
import { listActionsApi } from '@/api/action'
const props = defineProps(['question', 'student_id', 'exam_id', 'display'])
const emit = defineEmits(['update:display', 'update:question'])
const student_id = props.student_id
const exam_id = props.exam_id
const showDialog = ref(false)
watch(() => props.question, (newVal) => {
  if (newVal) {
    showDialog.value = true
    getActions()
  }
})

function handleClose() {
  showDialog.value = false
  emit('update:question', null)
  actions.value = []
  
}

const actions = ref()
function getActions() {
  listActionsApi(exam_id, student_id, props.question.question_id).then(res => {
    actions.value = res.data
  })
}

function formatTime(time: number) {
  return new Date(time).toLocaleString()
}
</script>