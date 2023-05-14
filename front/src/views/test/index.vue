<template>
    <el-button type="primary" @click="dialogVisible = true">新增</el-button>

    <el-dialog v-model="dialogVisible" title="新增考生">
        <el-form ref="formRef" :model="addForm" :rules="rules" label-width="80px" style="width: 400px">
            <el-form-item label="学号" prop="student_id">
                <el-input v-model.lazy="addForm.student_id"></el-input>
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
                <el-button type="primary" @click="addFormSubmit">提交</el-button>
                <el-button @click="dialogVisible = false">取消</el-button>
            </el-form-item>
        </el-form>
    </el-dialog>
</template>

<script setup lang="ts">
import { Student } from '@/api/student/types';
import { ElMessageBox, FormInstance } from 'element-plus';


const dialogVisible = ref(false)
const addForm = ref({} as Student)
const formRef = ref<FormInstance>()
const rules = {
    user_id: [{ required: true, message: '请输入用户ID', trigger: 'blur' }],
    student_id: [{required: true, message: '请输入学号', trigger: 'blur'}],
    username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
    password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
    role: [{ required: true, message: '请选择角色', trigger: 'change' }],
    name: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
    gender: [{ required: true, message: '请选择性别', trigger: 'change' }],
    email: [{ required: false, message: '请输入邮箱', trigger: 'blur' }],
    phone: [{ required: false, message: '请输入电话', trigger: 'blur' }]
}

function addFormSubmit(){

}
</script>