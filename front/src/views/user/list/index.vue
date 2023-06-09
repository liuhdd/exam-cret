<template>
    <div class="app-container">
        <el-card width="100%">
            <el-form :inline="true">
                <el-form-item label="用户名">
                    <el-input v-model="username" />
                </el-form-item>
                <el-form-item>
                    <el-form-item><el-button type="primary" :icon="Search" @click="handleSearch">搜索</el-button></el-form-item>
                </el-form-item>
                <el-form-item>
                <el-button type="primary" @click="handleAdd">新增</el-button>
            </el-form-item>
            </el-form>
            <el-divider></el-divider>
            <el-table :data="filteredUsers">
                <el-table-column prop="username" label="用户名">
                </el-table-column>
                <el-table-column prop="role" label="角色" 
                :filters="[
                    { text: 'admin', value: 'admin' },
                    { text: 'teacher', value: 'teacher' },
                    { text: 'student', value: 'student' },
                ]"
                :filter-method="filterTag"
                >
                    <template #default="scope">
                        <el-tag v-if="scope.row.role" :type="scope.row.role === 'student' ? '' : 'success'"
                            disable-transitions>
                            {{ scope.row.role
                            }}</el-tag>
                    </template>
                </el-table-column>
                <el-table-column label="操作">
                    <template #default="{ row }">
                        <el-button type="primary" size="small" @click="handleEdit(row)">编辑</el-button>
                        <el-button :type="row.state === 1? 'warning' : 'success'" size="small" @click="">
                            {{ row.state === 1 ? '禁用' : '启用' }}
                        </el-button>
                        <el-button type="primary" size="small" @click="handleReset(row)">
                            重置密码
                        </el-button>
                        <el-button type="danger" size="small" @click="handleDelete(row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-card>

        <el-dialog v-model="dialogVisible" title="新增用户">
            <el-form ref="formRef" :model="addForm" :rules="rules" label-width="80px" style="width: 400px">
                <el-form-item label="用户名" prop="username">
                    <el-input v-model.lazy="addForm.username"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="password">
                    <el-input v-model.lazy="addForm.password" type="password"></el-input>
                </el-form-item>
                <el-form-item label="角色" prop="role">
                    <el-radio-group v-model.lazy="addForm.role">
                        <el-radio label="student">学生</el-radio>
                        <el-radio label="teacher">教师</el-radio>
                        <el-radio label="admin">管理员</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="addFormSubmit(formRef)">提交</el-button>
                    <el-button @click="dialogVisible = false">取消</el-button>
                </el-form-item>

            </el-form>
        </el-dialog>

    </div>
</template>

<script setup lang="ts">
import { Student } from '@/api/student/types';
import { listUserApi, deleteUserApi, resetUserApi } from '@/api/auth';
import { FormInstance } from 'element-plus';
import { User } from '@/api/auth/types';
import { Search } from '@element-plus/icons-vue'
const users = ref()

function getStudents() {
    listUserApi().then(res => {
        users.value = res.data
        filteredUsers.value = res.data
    })
}
onMounted(() => {
    getStudents()
})

function handleReset(row: User) {

    ElMessageBox.confirm('此操作将重置用户密码，是否继续？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        resetUserApi(row.username).then(res => {
            ElMessage({
                type: 'success',
                message: '重置成功'
            });
        })
    }).catch(() => {
        ElMessage({
            type: 'info',
            message: '已取消重置'
        });
    });
}

const username = ref('')
const filteredUsers = ref()
function handleSearch() {
    filteredUsers.value = users.value.filter((user: User) => {
        return user.username.includes(username.value)
    })
}

function filterTag(value: string, row: User) {
    return row.role === value
}

function handleEdit(row: Student) {
    addForm.value = row
    dialogVisible.value = true
}

function handleDelete(row: User) {
    deleteUserApi(row.username).then(res => {
        getStudents()
    })
}

function handleAdd() {
    dialogVisible.value = true

}

// 表单
const dialogVisible = ref(false)
const addForm = ref({} as Student)
const formRef = ref<FormInstance>()
const rules = {
    user_id: [{ required: true, message: '请输入用户ID', trigger: 'blur' }],
    student_id: [{ required: true, message: '请输入学号', trigger: 'blur' }],
    username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
    password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
    role: [{ required: true, message: '请选择角色', trigger: 'change' }],
    name: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
    gender: [{ required: true, message: '请选择性别', trigger: 'change' }],
    email: [{ required: false, message: '请输入邮箱', trigger: 'blur' }],
    phone: [{ required: false, message: '请输入电话', trigger: 'blur' }]
}

function addFormSubmit(formEl: FormInstance | undefined) {

    if (!formEl) return
    formEl.validate((valid) => {
        if (valid) {

            dialogVisible.value = false
            addForm.value = {} as Student
        } else {
            console.log('error submit!!')
            return false
        }
    })
}
</script>