<template>
  <a-button type="danger" @click="loginOut">退出</a-button>
</template>

<script>
export default {
  name: 'HeaderPrc',
  methods: {
    loginOut() {
      window.sessionStorage.clear('token')
      this.$router.push('/login')
    }
  },
}
</script>

<style scoped>
</style>