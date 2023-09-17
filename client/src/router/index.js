import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import ConnectWallet from '@/components/ConnectWallet'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/wallet',
      name: 'ConnectWallet',
      component: ConnectWallet
    },
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    }
  ]
})
