<template>
    <div class="app-container">
        <el-card width="100%">
            <el-form :inline="true">
                <el-form-item label="学号">
                    <el-input v-model="student_id" />
                </el-form-item>
                <el-form-item label="姓名">
                    <el-input v-model="name" />
                </el-form-item>
                <el-form-item>
                    <el-form-item><el-button type="primary" :icon="Search"
                            @click="handleSearch">搜索</el-button></el-form-item>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="handleAdd">添加</el-button>
                    <el-button type="primary" @click="handleImport">导入</el-button>
                </el-form-item>
            </el-form>
            <el-dialog v-model="dialogImportVisible" title="导入考生">
                <el-card>
                    <template #header>
                        <div class="card-header">
                        
                            <el-upload width="50%" ref="upload" class="upload-demo" action="" :limit="1" :auto-upload="false"
                                accept=".xlsx,.xls" :on-preview="fileChange" :on-exceed="handleExceed">
                                <template #trigger>
                                    <el-button type="primary">选择学生信息文件</el-button>
                                </template>
                                
                                
                            </el-upload>
                            <el-button class="button" type="success">导入</el-button>
                        </div>
                    </template>

                    <el-table :data="importStudents">
                        <el-table-column prop="student_id" label="学号">
                        </el-table-column>
                        <el-table-column prop="name" label="姓名">
                        </el-table-column>
                        <el-table-column prop="gender" label="性别">
                        </el-table-column>
                        <el-table-column prop="email" label="邮箱">
                        </el-table-column>
                        <el-table-column prop="phone" label="电话">
                        </el-table-column>
                    </el-table>
                </el-card>
            </el-dialog>



            <el-divider></el-divider>
            <el-table :data="filteredStudents">
                <el-table-column prop="student_id" label="学号">
                </el-table-column>
                <el-table-column prop="name" label="姓名">
                </el-table-column>
                <el-table-column prop="gender" label="性别">
                </el-table-column>
                <el-table-column prop="email" label="邮箱">
                </el-table-column>
                <el-table-column prop="phone" label="电话">
                </el-table-column>
                <el-table-column label="操作" width="220">
                    <template #default="{ row }">
                        <el-button type="primary" size="small" @click="handleCheck(row)">查看考试</el-button>
                        <el-button type="primary" size="small" @click="handleEdit(row)">编辑</el-button>
                        <el-button type="danger" size="small" @click="handleDelete(row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-card>

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
                    <el-button type="primary" @click="addFormSubmit(formRef)">提交</el-button>
                    <el-button @click="dialogVisible = false">取消</el-button>
                </el-form-item>

            </el-form>
        </el-dialog>

    </div>
</template>

<script setup lang="ts">
import { Student } from '@/api/student/types';
import { getAllStudentApi, createStudentApi, updateStudentApi, deleteStudentApi } from '@/api/student/index';
import { ElDialog, ElMessageBox, FormInstance, UploadFile, UploadInstance, UploadProps, UploadRawFile, genFileId } from 'element-plus';
import router from '@/router';
import { Search } from '@element-plus/icons-vue';
import { read, utils } from 'xlsx'
const students = ref()



function getStudents() {
    getAllStudentApi().then(({ data }) => {
        students.value = data
        filteredStudents.value = data
    })
}
onMounted(() => {
    getStudents()
})

const importStudents = ref<Student[]>()
const dialogImportVisible = ref(false)

function handleImport() {
    dialogImportVisible.value = true
}

function loadFile(file: File) {
    const reader = new FileReader()
    reader.onload = (e) => {
        const data = e.target!.result
        const workbook = read(data, { type: 'binary' })
        const sheetNames = workbook.SheetNames
        const worksheet = workbook.Sheets[sheetNames[0]]
        const json = utils.sheet_to_json(worksheet, { header: ["student_id", "name", "gender", "email", "phone"] })

        const stus = new Array<Student>()
        json.slice(1).forEach((e) => {
            stus.push(e as Student)
        })
        importStudents.value = stus

    }
    reader.readAsBinaryString(file)
}
function fileChange(file: UploadFile) {

}

const upload = ref<UploadInstance>()
const handleExceed: UploadProps['onExceed'] = (files) => {
    upload.value!.clearFiles()
    const file = files[0] as UploadRawFile
    file.uid = genFileId()
    upload.value!.handleStart(file)
    loadFile(file)
}

const student_id = ref('')
const name = ref('')
const filteredStudents = ref()
function handleSearch() {
    filteredStudents.value = students.value.filter((student: Student) => {
        return student.student_id.includes(student_id.value) && student.name.includes(name.value)
    })
}
var operate = 0

function handleEdit(row: Student) {
    addForm.value = row
    dialogVisible.value = true
    operate = 1
}

function handleCheck(row: Student) {
    router.push({
        path: '/exam/detail',
        query: {
            student_id: row.student_id
        }
    })
}
function handleDelete(row: Student) {
    ElMessageBox.confirm('此操作将删除该考生, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        deleteStudentApi(row.user_id).then(res => {
            getStudents()
        })
    }).catch(() => {
        console.log('取消删除')
    })

}

function handleAdd() {
    dialogVisible.value = true
    operate = 0
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
            if (operate == 0) {
                createStudentApi(addForm.value).then(res => {
                    getStudents()
                })
            } else {
                updateStudentApi(addForm.value).then(res => {
                    getStudents()
                })
            }
            dialogVisible.value = false
            addForm.value = {} as Student
        } else {
            console.log('error submit!!')
            return false
        }
    })
}
</script>

<style>

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>