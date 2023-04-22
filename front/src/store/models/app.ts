import { defineStore } from "pinia"
import { useStorage } from "@vueuse/core"

import defaultSettings from "@/settings"

import zhCn from "element-plus/es/locale/lang/zh-cn"
import en from "element-plus/es/locale/lang/en"

import { reactive, computed } from "vue"

export const useAppStore = defineStore("app", ()=>{
    
    const device = useStorage("device", "desktop");
    const size = useStorage<any>("size", defaultSettings.size);
    const language = useStorage("language", defaultSettings.language);
    const sidebarStatus = useStorage("sidebarStatus", "closed");

    const sidebar = reactive({
        opened: sidebarStatus.value !== "closed",
        withoutAnimation: false,
    })
    const locale = computed(() => {
        if (language?.value == "en") {
          return en;
        } else {
          return zhCn;
        }
    })

    function toggleSidebar(withoutAnimation: boolean) {
        sidebar.opened = !sidebar.opened;
        sidebar.withoutAnimation = withoutAnimation;
        if (sidebar.opened) {
          sidebarStatus.value = "opened";
        } else {
          sidebarStatus.value = "closed";
        }
      }
    
      function closeSideBar(withoutAnimation: boolean) {
        sidebar.opened = false;
        sidebar.withoutAnimation = withoutAnimation;
        sidebarStatus.value = "closed";
      }
    
      function openSideBar(withoutAnimation: boolean) {
        sidebar.opened = true;
        sidebar.withoutAnimation = withoutAnimation;
        sidebarStatus.value = "opened";
      }
    
      function toggleDevice(val: string) {
        device.value = val;
      }
    
      function changeSize(val: string) {
        size.value = val;
      }

      function changeLanguage(val: string) {
        language.value = val;
      }

      return {
        device,
        sidebar,
        language,
        locale,
        size,
        toggleDevice,
        changeSize,
        changeLanguage,
        toggleSidebar,
        closeSideBar,
        openSideBar,
      }
})