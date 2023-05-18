<script setup lang="ts">
import { storeToRefs } from "pinia";
import { useRoute, useRouter } from "vue-router";
import { useAppStore } from "@/store/models/app";
import { useTagsViewStore } from "@/store/models/tagsView";
import { useUserStore } from "@/store/models/user";
import { ElMessageBox } from "element-plus";
import { User } from "@/api/auth/types";

import {updateUserApi} from "@/api/auth"


const appStore = useAppStore();
const tagsViewStore = useTagsViewStore();
const userStore = useUserStore();

const route = useRoute();
const router = useRouter();

const { device } = storeToRefs(appStore); // 设备类型：desktop-宽屏设备 || mobile-窄屏设备

function toggleSideBar() {
  appStore.toggleSidebar(true);
}





function handlerSetPwd(){
  ElMessageBox.prompt('请输入密码', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    inputPattern: /\S/,
    inputErrorMessage: '密码不能为空',
    inputType: 'password',
    showClose: false,
    closeOnClickModal: false,
    closeOnPressEscape: false,
    closeOnHashChange: false,
    center: true
  }).then(({ value }) => {
    updateUserApi({ username: userStore.username,password: value } as User)
      .then(() => {
        ElMessage.success('修改成功')
      })
      .catch(() => {
        ElMessage.error('修改失败')
      })
  })
}


// 注销
function logout() {
  ElMessageBox.confirm("确定注销并退出系统吗？", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    userStore
      .logout()
      .then(() => {
        tagsViewStore.delAllViews();
      })
      .then(() => {
        router.push(`/login?redirect=${route.fullPath}`);
      });
  });
}
</script>

<template>
  <!-- 顶部导航栏 -->
  <div class="navbar">
    <!-- 左侧面包屑 -->
    <div class="flex">
      <hamburger
        :is-active="appStore.sidebar.opened"
        @toggleClick="toggleSideBar"
      />
      <breadcrumb />
    </div>

    <!-- 右侧导航设置 -->
    <div class="flex">
      <!-- 导航栏设置(窄屏隐藏)-->

      <div v-if="device !== 'mobile'" class="flex items-center">
        {{ userStore.username }}
        <!--全屏 -->
        <screenfull class="navbar-setting-item" />
        <!-- 布局大小 -->
        <el-tooltip content="布局大小" effect="dark" placement="bottom">
          <size-select class="navbar-setting-item" />
        </el-tooltip>
        <!--语言选择-->
        <lang-select class="navbar-setting-item" />
      </div>

      <!-- 用户头像 -->
      <el-dropdown trigger="click">
        <div class="flex justify-center items-center mx-2">
          <el-avatar class="w-[40px] h-[40px] rounded-lg"  src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png" />
          
          <i-ep-caret-bottom class="w-3 h-3" />
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <router-link to="/">
              <el-dropdown-item>{{ $t("navbar.dashboard") }}</el-dropdown-item>
            </router-link>
            <el-dropdown-item divided @click="handlerSetPwd">
              修改密码
            </el-dropdown-item>
            <el-dropdown-item divided @click="logout">
              {{ $t("navbar.logout") }}
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.navbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 50px;
  background-color: #fff;
  box-shadow: 0 0 1px #0003;

  .navbar-setting-item {
    display: inline-block;
    width: 30px;
    height: 50px;
    line-height: 50px;
    color: #5a5e66;
    text-align: center;
    cursor: pointer;

    &:hover {
      background: rgb(249 250 251 / 100%);
    }
  }
}
</style>
