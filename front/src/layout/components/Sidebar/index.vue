<script setup lang="ts">
import { useRoute } from 'vue-router';

import AppLink from './Link.vue';
import Logo from './Logo.vue';

import { useSettingsStore } from '@/store/models/settings';
import { useAppStore } from '@/store/models/app';
import { storeToRefs } from 'pinia';
import variables from '@/styles/variables.module.scss';
import { dynamicRoutes } from '@/router/index'
const settingsStore = useSettingsStore();
const appStore = useAppStore();

const routes = dynamicRoutes
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
        <el-sub-menu index="exam">
          <template #title>
            <svg-icon icon-class="user" />
            考试管理
          </template>
          <app-link to="/exam/show">
            <el-menu-item index="/show" teleported>
              <span>
                结果查询
              </span>
            </el-menu-item>
          </app-link>
          <app-link to="/exam/verify">
            <el-menu-item index="/show" teleported>
              <span>
                结果核验
              </span>
            </el-menu-item>
          </app-link>
        </el-sub-menu>
      </el-menu>
    </el-scrollbar>
  </div>
</template>
