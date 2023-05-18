<template>
    <div class="app-container">
        <el-form :inline="true">
            <el-form-item>
                <el-button type="primary" @click="handleResume">
                    恢复
                </el-button>
            </el-form-item>
            <el-form-item>
                <el-button type="success" @click="handleBackup">
                    备份
                </el-button>
            </el-form-item>
        </el-form>
        <el-card>
            <el-upload class="upload-demo" drag action="" :http-request="upload" accept=".sql"
                multiple>
                <el-icon class="el-icon--upload"><upload-filled /></el-icon>
                <div class="el-upload__text">
                    拖拽数据文件至此处或者 <em>点击选择文件</em>
                </div>
                <template #tip>
                    <div class="el-upload__tip">
                        请上传 <em>数据文件</em>
                    </div>
                </template>
            </el-upload>
        </el-card>
    </div>
</template>
  
<script setup lang="ts">
import { UploadFilled } from '@element-plus/icons-vue'
import { uploadApi, confimReume, backupApi } from '@/api/system/index'
import { ElMessage, ElMessageBox, UploadRequestOptions } from 'element-plus'
import router from '@/router';
import request from '@/utils/request';

function handleBackup() {
    ElMessageBox.confirm('是否确定备份数据', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        window.open(request.defaults.baseURL+'/system/backup', '_self')
    }).catch(() => {
        ElMessage({
            type: 'info',
            message: '已取消备份'
        });
    });
    
}
function handleResume() {
    ElMessageBox.confirm('此操作将恢复数据库，是否继续？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        confimReume().then(() => {
            ElMessage({
                type: 'success',
                message: '恢复成功!'
            });
        }).catch(() => {
            ElMessage.error({
                type: 'info',
                message: '恢复失败'
            });
        });
    }).catch(() => {
        ElMessage({
            type: 'info',
            message: '已取消恢复'
        });
    });
}
function upload(options: UploadRequestOptions) {
    return new Promise((resolve, reject) => {
        const formData = new FormData();
    formData.append("file",options.file);
       uploadApi(formData).then(() => {
           resolve(true)
       }).catch(() => {
           reject(false)
       })
    })

}

</script>
  