<template>
  <a-layout class="container">
    <NavPrc></NavPrc>
    <a-layout>
      <a-layout-header class="headerBtn">
        <HeaderPrc></HeaderPrc>
      </a-layout-header>
      <a-layout-content>
        <router-view :key="$route.path"></router-view>
      </a-layout-content>
      <a-layout-footer>
        <FooterPrc></FooterPrc>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
import NavPrc from '../components/admin/NavPrc.vue'
import FooterPrc from '../components/admin/FooterPrc.vue'
import HeaderPrc from '../components/admin/HeaderPrc.vue'

export default {
  name: 'AdminPrc',
  components: { NavPrc, FooterPrc, HeaderPrc }
}
</script>

<style scoped>
.container {
  height: 100%;
}
.headerBtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>
