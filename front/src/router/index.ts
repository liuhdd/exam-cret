import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";

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
    children: [
      {
        path: "backup",
        component: () => import("@/views/system/backup/index.vue"),
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
        children: [
          {
            path: "list",
            component: () => import("@/views/exam/list/index.vue"),
          },
          {
            path: "detail/:exam_id",
            component: () => import("@/views/exam/detail/index.vue"),
            props: true,
          },
          {
            path: "detail",
            component: () => import("@/views/exam/detail/index.vue")
          },
          {
            path: "verify",
            component: () => import("@/views/exam/verify/index.vue"),
            props: true
          },
          {
            path: "verify/:examID",
            component: () => import("@/views/exam/verify/index.vue"),
            props: true
          }
        ]
      }, 
      {
      path: "user",
      redirect: "/user/list",
      children: [
        {
          path: "list",
          component: import("@/views/user/list/index.vue"),
        },
        
        {
          path: "student",
          component: import("@/views/user/student/index.vue"),
        },
        {
          path: "teacher",
          component: import("@/views/user/teacher/index.vue"),
        },
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


export default router;
