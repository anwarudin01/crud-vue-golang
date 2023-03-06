import { createRouter, createWebHistory } from 'vue-router';

const routes = [
  {
    path: '/',
    name: 'datatable.index',
    component: () => import('../views/datatable/index.vue'),
  },
  {
    path: '/create',
    name: 'datatable.create',
    component: () => import('../views/datatable/Create.vue'),
  },
  {
    path: '/edit/:id',
    name: 'datatable.edit',
    component: () => import('../views/datatable/Edit.vue'),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
