import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";
import NProgress from "nprogress";
import "nprogress/nprogress.css";

export const Layout = () => import("@/layout/index.vue");

export const constantRoutes: Array<RouteRecordRaw> = [
  {
    path: "/test",
    component: () => import("@/views/test/index.vue"),
  },
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
    path: "/system",
    component: Layout,
    redirect: "/system/backup",
    name: "系统管理",
    meta: { title: "系统管理", hidden: false},
    children: [
      {
        path: "backup",
        name: "数据备份",
        component: () => import("@/views/system/backup/index.vue"),
        meta: { title: "数据备份", hidden: false, keepAlive: true },
      }
    ]
  },
  {
    path: "/",
    component: Layout,
    redirect: "/dashboard",
    children: [
      
      {
        path: "exam",
        redirect: "/exam/list",
        meta: { title: "数据管理", hidden: true, keepAlive: true },
        children: [
          {
            path: "list",
            name: "考试管理",
            component: () => import("@/views/exam/list/index.vue"),
            meta: { title: "考试管理", hidden: false, keepAlive: true },
          },

          {
            path: "detail",
            name: "成绩查询",
            component: () => import("@/views/exam/detail/index.vue"),
            meta: { title: "成绩查询", hidden: false,keepAlive: true },
            props: route => ({exan_id: route.query.exan_id, student_id: route.query.student_id}),
          },
          {
            path: "verify",
            name: "成绩核验",
            component: () => import("@/views/exam/verify/index.vue"),
            meta: { title: "成绩核验", hidden: false,keepAlive: true },
            props: route => ({exan_id: route.query.exan_id, student_id: route.query.student_id}),
          },
        ]
      }, 
      {
      path: "user",
      redirect: "/user/student",
      meta: { title: "用户管理", hidden: false,keepAlive: true },
      children: [
        {
          path: "list",
          name: "用户管理",
          component: () => import("@/views/user/list/index.vue"),
        },
        
        {
          path: "student",
          name: "学生管理",
          component: () => import("@/views/user/student/index.vue"),
          meta: { title: "考生管理", hidden: false,keepAlive: true },
        },
        {
          path: "teacher",
          name: "教师管理",
          component: () => import("@/views/user/teacher/index.vue"),
          meta: { title: "教师管理", hidden: false,keepAlive: true },
        },
      ]
      },
      {
        path: "dashboard",
        component: () => import("@/views/dashboard/index.vue"),
        name: "Dashboard",
        meta: { title: "dashboard", icon: "homepage", affix: true, keepAlive: true,},
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



export const constantRoutes1: RouteRecordRaw[] = [

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
  {
    path: "/exam",
    component: Layout,
    redirect: "/exam/manage",
    meta: {
      title: "考试管理",
      icon: "system",
      hidden: false,
      roles: ["admin"],
      keepAlive: true,
    },
    children: [
      {
        path: "manage",
        name: "manage",
        component: () => import("@/views/exam/manager/index.vue"),
        meta: {
          title: "考试管理",
          icon: "user",
          hidden: false,
          roles: ["admin"],
          keepAlive: true,
        },
      },
      {
        path: "grade",
        name: "grade",
        component: () => import("@/views/exam/detail/index.vue"),
        meta: {
          title: "成绩查询",
          icon: "user",
          idden: false,
          roles: ["admin", "student"],
          keepAlive: true,
        }
      },
      {
        path: "verify",
        name: "verify",
        component: () => import("@/views/exam/verify/index.vue"),
        meta: {
          title: "成绩核验",
          icon: "user",
          idden: false,
          roles: ["admin", "student"],
          keepAlive: true,
        }
      },
      {
        path: "grade/:exam_id",
        component: () => import("@/views/exam/detail/index.vue"),
        props: true,
        name: "grade",
        meta: {
          title: "成绩查询",
          icon: "user",
          hidden: false,
          roles: ["admin", "student"],
          keepAlive: true,
        },
      },
      {
        path: "verify/:examID",
        component: () => import("@/views/exam/verify/index.vue"),
        props: true,
        name: "verify",
        meta: {
          title: "成绩核验",
          icon: "user",
          hidden: false,
          roles: ["admin", "student"],
          keepAlive: true,
        }
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

NProgress.configure({ showSpinner: false }); 
router.beforeEach(async () =>{
  NProgress.start();

})

router.afterEach(() => {
  NProgress.done();
});

export default router;
