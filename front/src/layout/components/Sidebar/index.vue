<script setup lang="ts">
import { useRoute } from 'vue-router';

import AppLink from './Link.vue';
import Logo from './Logo.vue';

import { useSettingsStore } from '@/store/models/settings';
import { useAppStore } from '@/store/models/app';
import { storeToRefs } from 'pinia';
import variables from '@/styles/variables.module.scss';
const settingsStore = useSettingsStore();
const appStore = useAppStore();
const { sidebarLogo } = storeToRefs(settingsStore);
const route = useRoute();
</script>

<template>
  <div :class="{ 'has-logo': sidebarLogo }">
    <logo v-if="sidebarLogo" :collapse="!appStore.sidebar.opened" />
    <el-scrollbar>
      <el-menu :default-active="route.path" :collapse="!appStore.sidebar.opened" :background-color="variables.menuBg"
        :text-color="variables.menuText" :active-text-color="variables.menuActiveText" :unique-opened="false"
        :collapse-transition="false" mode="vertical">
        <el-sub-menu index="user">
          <template #title>
            <svg-icon icon-class="user" />
            用户管理
          </template>
          <app-link to="/user/list">
            <el-menu-item index="/list" teleported>
              <span>
                用户查询
              </span>
            </el-menu-item>
          </app-link>
          <app-link to="/user/student">
            <el-menu-item index="/student" teleported>
              <span>
                考生管理
              </span>
            </el-menu-item>
          </app-link>
          <app-link to="/user/teacher">
            <el-menu-item index="/teacher" teleported>
              <span>
                教师管理
              </span>
            </el-menu-item>
          </app-link>
        </el-sub-menu>
        <el-sub-menu index="/exam">
          <template #title>
            <svg-icon icon-class="user" />
            考试管理
          </template>
          <app-link to="/exam/list">
            <el-menu-item index="/list" teleported>
              <span>
                考试查询
              </span>
            </el-menu-item>
          </app-link>
          <app-link to="/exam/detail">
            <el-menu-item index="/detail" teleported>
              <span>
                成绩查询
              </span>
            </el-menu-item>
          </app-link>
          <app-link to="/exam/verify">
            <el-menu-item index="/verify" teleported>
              <span>
                成绩核验
              </span>
            </el-menu-item>
          </app-link>
        </el-sub-menu>
        <el-sub-menu index="/system">
          <template #title>
            <svg-icon icon-class="user" />
            系统管理
          </template>
          <app-link to="/system/backup">
            <el-menu-item index="/role" teleported>
              <span>
                数据备份
              </span>
            </el-menu-item>
          </app-link>
          <app-link to="/system/permission">
            <el-menu-item index="/permission" teleported>
              <span>
                权限管理
              </span>
            </el-menu-item>
          </app-link>
          <app-link to="/system/log">
            <el-menu-item index="/log" teleported>
              <span>
                日志管理
              </span>
            </el-menu-item>
          </app-link>
        </el-sub-menu>
      </el-menu>
    </el-scrollbar>
  </div>
</template>
