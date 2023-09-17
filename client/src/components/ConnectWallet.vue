<template>
    <div>
        <h1>Connect Wallet</h1>
        <p>Connect your wallet to continue</p>
        <button v-if="!isConnectedToMetaMask" @click="connectWeb3">Connect Wallet</button>
        <div v-else>
          <p>{{ userAddress }}</p>
          <p>{{ chainId }}</p>
          <button></button>
        </div>
    </div>
</template>
<script>

import {ref} from 'vue'
import {RequestParams} from '../config/request'
import axios from 'axios'

export default {
  name: 'ConnectWallet',
  setup () {
    const chainId = ref(0)
    return {
      chainId
    }
  },
  data () {
    return {
      userAddress: '',
      isConnectedToMetaMask: false
    }
  },
  init () {
    this.connectWeb3()
  },
  methods: {
    connectWeb3: function () {
      // Check if MetaMask is installed
      console.log(window.ethereum.isConnected())
      if (window.ethereum) {
        console.log(window.ethereum)
        window.ethereum.request({ method: 'eth_requestAccounts' })
          .then((res) => {
            console.log('result', res)
            this.userAddress = res[0]
            this.getUserAccount()
            this.isConnectedToMetaMask = true

            this.registerServer({
              address: this.userAddress,
              chainId: this.chainId
            })
          })
          .catch((err) => {
            console.error(err)
          })
      } else {
        // MetaMask is not installed, handle accordingly
        console.error('MetaMask is not installed.')
      }
    },
    getUserAccount: async function () {
      window.ethereum.request({
        method: 'eth_chainId'
      }).then(res => {
        this.chainId = res
      }).catch(err => {
        console.error(err)
      })
    },
    registerServer (data) {
      axios.post(RequestParams.host + RequestParams.path.register, {
        data
      }, {
        headers: {
          'Access-Control-Allow-Origin': '*'
        }
      })
        .then(res => {
          console.log(res)
        })
        .catch(err => {
          console.error(err)
        })
    }
  }
}
</script>
