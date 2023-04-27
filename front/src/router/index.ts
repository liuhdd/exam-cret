import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";

export const Layout = () => import("@/layout/index.vue");

export const constantRoutes: Array<RouteRecordRaw> = [
  {
    path: "/redirect",
    component: Layout,
    meta: { hidden: true },
    children: [
      {
        path: "/redirect/:path(.*)",
        component: () => import("@/views/redirect/index.vue"),
      },
    ],
  },
  {
    path: "/login",
    component: () => import("@/views/login/index.vue"),
    meta: { hidden: true },
  },
  {
    path: "/",
    component: Layout,
    redirect: "/dashboard",
    children: [
      {
        path: "exam",
        redirect: "/exam/show",
        children: [
          {
            path: "show",
            component: () => import("@/views/exam/show/index.vue"),
          },
          {
            path: "detail",
            component: () => import("@/views/exam/detail/index.vue")
          }
        ]
      }, 
      {
        path: "dashboard",
        component: () => import("@/views/dashboard/index.vue"),
        name: "Dashboard",
        meta: { title: "dashboard", icon: "homepage", affix: true },
      },
      {
        path: "401",
        component: () => import("@/views/error-page/401.vue"),
        meta: { hidden: true },
      },
      {
        path: "404",
        component: () => import("@/views/error-page/404.vue"),
        meta: { hidden: true },
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: constantRoutes,
  scrollBehavior: () => ({ left: 0, top: 0 }),
});

export function resetRouter() {
  router.replace({ path: "/login" });
  location.reload();
}

export const dynamicRoutes = [
  {
    path: "/exam",
    component: Layout,
    redirect: "/exam/show",
    meta: {
      title: "考试管理",
      icon: "system",
      hidden: false,
      roles: ["ADMIN"],
      keepAlive: true,
    },
    children: [
      {
        path: "show",
        name: "show",
        meta: {
          title: "结果查询",
          icon: "user",
          hidden: false,
          roles: ["ADMIN"],
          keepAlive: true,
        },
      },
      {
        path: "detail",
        component: "exam/index",
        name: "detail",
        meta: {
          title: "结果查询z",
          icon: "user",
          hidden: false,
          roles: ["ADMIN"],
          keepAlive: true,
        },
      },
    ],
  },
  {
    path: "/api",
    component: "Layout",
    meta: {
      title: "接口",
      icon: "api",
      hidden: false,
      roles: ["ADMIN"],
      keepAlive: true,
    },
    children: [
      {
        path: "apidoc",
        component: "demo/apidoc",
        name: "apidoc",
        meta: {
          title: "接口文档",
          icon: "api",
          hidden: false,
          roles: ["ADMIN"],
          keepAlive: true,
        },
      },
    ],
  },
  {
    path: "/external-link",
    component: "Layout",
    redirect: "noredirect",
    meta: {
      title: "外部链接",
      icon: "link",
      hidden: false,
      roles: ["ADMIN"],
      keepAlive: true,
    },
    children: [
      {
        path: "https://www.cnblogs.com/haoxianrui/p/17331952.html",
        meta: {
          title: "document",
          icon: "document",
          hidden: false,
          roles: ["ADMIN"],
          keepAlive: true,
        },
      },
    ],
  },
  {
    path: "/multi-level-menu",
    component: "Layout",
    redirect: "/nested/level1/level2",
    meta: {
      title: "多级菜单",
      icon: "multi_level",
      hidden: false,
      roles: ["ADMIN"],
      keepAlive: true,
    },
    children: [
      {
        path: "nested_level1_index",
        component: "nested/level1/index",
        redirect: "/nested/level1/level2",
        meta: {
          title: "菜单一级",
          icon: "",
          hidden: false,
          roles: ["ADMIN"],
          keepAlive: true,
        },
        children: [
          {
            path: "nested_level1_level2_index",
            component: "nested/level1/level2/index",
            redirect: "/nested/level1/level2/level3",
            meta: {
              title: "菜单二级",
              icon: "",
              hidden: false,
              roles: ["ADMIN"],
              keepAlive: true,
            },
            children: [
              {
                path: "nested_level1_level2_level3_index1",
                component: "nested/level1/level2/level3/index1",
                name: "nested_level1_level2_level3_index1",
                meta: {
                  title: "菜单三级-1",
                  icon: "",
                  hidden: false,
                  roles: ["ADMIN"],
                  keepAlive: true,
                },
              },
              {
                path: "nested_level1_level2_level3_index2",
                component: "nested/level1/level2/level3/index2",
                name: "nested_level1_level2_level3_index2",
                meta: {
                  title: "菜单三级-2",
                  icon: "",
                  hidden: false,
                  roles: ["ADMIN"],
                  keepAlive: true,
                },
              },
            ],
          },
        ],
      },
    ],
  },
  {
    path: "/demo",
    component: "Layout",
    meta: {
      title: "组件封装",
      icon: "menu",
      hidden: false,
      roles: ["ADMIN"],
      keepAlive: true,
    },
    children: [
      {
        path: "wangEditor",
        component: "demo/wangEditor",
        name: "wangEditor",
        meta: {
          title: "富文本编辑器",
          icon: "",
          hidden: false,
          roles: ["ADMIN"],
          keepAlive: true,
        },
      },
      {
        path: "uploader",
        component: "demo/uploader",
        name: "uploader",
        meta: {
          title: "上传组件",
          icon: "",
          hidden: false,
          roles: ["ADMIN"],
          keepAlive: true,
        },
      },
      {
        path: "IconSelector",
        component: "demo/IconSelector",
        name: "IconSelector",
        meta: {
          title: "图标选择器",
          icon: "",
          hidden: false,
          roles: ["ADMIN"],
          keepAlive: true,
        },
      },
    ],
  },
];
export default router;
