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
      <el-descriptions-item v-if="question.question_type == 1">
        <template #label>选项</template>
        <el-tag v-for="option in question.options?.split('&&')" :key="option.option_id">
          {{ option }}
        </el-tag>
      </el-descriptions-item>
    </el-descriptions>

    <el-collapse>
      <el-collapse-item title="答题记录" name="1">
        <el-table :data="actions" border stripe style="width: 100%">
          <el-table-column prop="action_id" label="记录码" width="180"></el-table-column>
          <el-table-column prop="answer" label="答案" width="180"></el-table-column>
          <el-table-column prop="action_time" sortable width="180" label="操作时间">
            <template v-slot="{ row }">
              {{ formatTime(row.action_time) }}
            </template>
          </el-table-column>
        </el-table>
      </el-collapse-item>
      <el-collapse-item title="评分记录" name="2">
        <el-table :data="scores" border stripe style="width: 100%">
          <el-table-column prop="action_id" label="记录码" width="180"></el-table-column>
          <el-table-column prop="score" label="得分" width="180"></el-table-column>
          <el-table-column prop="scored_by" label="评分人" width="180"></el-table-column>
          <el-table-column prop="scored_time" sortable label="打分时间" width="180">
            <template v-slot="{ row }">
              {{ formatTime(row.scored_time) }}
            </template>
          </el-table-column>
        </el-table>
      </el-collapse-item>

    </el-collapse>


  </el-dialog>
</template>

<script setup lang="ts">
import { listActionsApi } from '@/api/action'
import { queryQuestionScore } from '@/api/score';

const props = defineProps(['question', 'student_id', 'exam_id', 'display'])
const emit = defineEmits(['update:display', 'update:question', 'update:student_id', 'update:exam_id'])
const student_id = props.student_id
const exam_id = props.exam_id
const showDialog = ref(false)
const question = ref()
watch(() => props.question, (newVal) => {

  if (newVal) {
    showDialog.value = true
    question.value = props.question
    getActions()
  }
})

function handleClose() {
  showDialog.value = false
  emit('update:question', null)
  actions.value = []
}

const scores = ref()
const actions = ref()
function getActions() {
  listActionsApi(exam_id, student_id, props.question.question_id).then(res => {
    actions.value = res.data
  })
  queryQuestionScore(exam_id, student_id, props.question.question_id).then(res => {
    scores.value = res.data
  })
}

function formatTime(time: number) {
  return new Date(time).toLocaleString()
}
</script>