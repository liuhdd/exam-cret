<template>
    <div class="app-container">
        <el-card width="100%">
            <el-form :inline="true">
                <el-form-item label="职工号">
                    <el-input v-model="teacher_id" />
                </el-form-item>
                <el-form-item label="姓名">
                    <el-input v-model="name" />
                </el-form-item>
                <el-form-item>
                    <el-form-item><el-button type="primary" :icon="Search" @click="handleSearch">搜索</el-button></el-form-item>
                </el-form-item>
                <el-form-item>
                <el-button type="primary" @click="handleAdd">添加</el-button>
            </el-form-item>
            </el-form>
            <el-divider></el-divider>
            <el-table :data="filteredTeachers">
                <el-table-column prop="teacher_id" label="职工号">
                </el-table-column>
                <el-table-column prop="name" label="姓名">
                </el-table-column>
                <el-table-column prop="gender" label="性别">
                </el-table-column>
                <el-table-column prop="email" label="邮箱">
                </el-table-column>
                <el-table-column prop="phone" label="电话">
                </el-table-column>
                <el-table-column label="操作">
                    <template #default="{ row }">
                        <el-button type="primary" size="small" @click="handleEdit(row)">编辑</el-button>
                        <el-button type="danger" size="small" @click="handleDelete(row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-card>

        <el-dialog v-model="dialogVisible" title="新增教师">
            <el-form ref="formRef" :model="addForm" :rules="rules" label-width="80px" style="width: 400px">
                <el-form-item label="职工号" prop="teacher_id">
                    <el-input v-model.lazy="addForm.teacher_id"></el-input>
                </el-form-item>
                <el-form-item label="姓名" prop="name">
                    <el-input v-model.lazy="addForm.name"></el-input>
                </el-form-item>
                <el-form-item label="性别" prop="gender">
                    <el-radio-group v-model.lazy="addForm.gender">
                        <el-radio label="男">男</el-radio>
                        <el-radio label="女">女</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="邮箱" prop="email">
                    <el-input v-model.lazy="addForm.email"></el-input>
                </el-form-item>
                <el-form-item label="电话" prop="phone">
                    <el-input v-model.lazy="addForm.phone"></el-input>
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
import { Teacher } from '@/api/teacher/types';
import { createTeacherApi, deleteTeacherApi, updateTeacherApi, getAllTeacherApi } from '@/api/teacher/index';
import { FormInstance } from 'element-plus';
import { Search } from '@element-plus/icons-vue';

const teachers = ref()

function getTeachers() {
    getAllTeacherApi().then(res => {
        teachers.value = res.data
        filteredTeachers.value = res.data
    })
}
onMounted(() => {
    getTeachers()
})


const teacher_id = ref('')
const name = ref('')
const filteredTeachers = ref()
function handleSearch() {
    filteredTeachers.value = teachers.value.filter((teacher: Teacher) => {
        return teacher.teacher_id.indexOf(teacher_id.value)==0 && teacher.name.indexOf(name.value)==0
    })
}
var operate = 0

function handleEdit(row: Teacher) {
    addForm.value = row
    dialogVisible.value = true
    operate = 1
}

function handleDelete(row: Teacher) {
    deleteTeacherApi(row.user_id).then(res => {
        getTeachers()
    })
}

function handleAdd() {
    dialogVisible.value = true
    operate = 0
}

// 表单
const dialogVisible = ref(false)
const addForm = ref({} as Teacher)
const formRef = ref<FormInstance>()
const rules = {
    user_id: [{ required: true, message: '请输入用户ID', trigger: 'blur' }],
    teacher_id: [{ required: true, message: '请输入职工号', trigger: 'blur' }],
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
            if (operate == 0) {
                createTeacherApi(addForm.value).then(res => {
                    getTeachers()
                })
            } else {
                updateTeacherApi(addForm.value).then(res => {
                    getTeachers()
                })
            }
            dialogVisible.value = false
            addForm.value = {} as Teacher
        } else {
            console.log('error submit!!')
            return false
        }
    })
}
</script>