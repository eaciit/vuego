import Vue from 'vue'
import VueRouter from 'vue-router'

import App from './App.vue'

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    //redirect: '/products'
    component: App
  },
/*  {
    path: '/products',
    component: ProductList
  },
  {
    path: '/products/:id',
    component: Product
  }
*/
];

const router = new VueRouter({
  routes
});

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
