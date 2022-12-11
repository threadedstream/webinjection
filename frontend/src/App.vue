<template>
  <MDBNavbar expand="lg" light bg="light" container>
    <MDBNavbarBrand href="#">WebInjection</MDBNavbarBrand>
    <MDBNavbarToggler
        @click="collapse1 = !collapse1"
        target="#navbarSupportedContent"
    ></MDBNavbarToggler>
    <MDBCollapse v-model="collapse1" id="navbarSupportedContent">
      <MDBNavbarNav class="mb-2 mb-lg-0">
        <MDBNavbarItem to="/" active>
          Home
        </MDBNavbarItem>
        <MDBNavbarItem to="/tasks" active>
          Tasks
        </MDBNavbarItem>
      </MDBNavbarNav>
    </MDBCollapse>
  </MDBNavbar>
  <router-view></router-view>
</template>

<script setup lang="ts">
import {
  MDBNavbar,
  MDBNavbarToggler,
  MDBNavbarBrand,
  MDBNavbarNav,
  MDBNavbarItem,
  MDBCollapse,
} from 'mdb-vue-ui-kit'

import { ref } from 'vue'

const collapse1 = ref(false)
const dropdown1 = ref(false)

</script>

<script lang="ts">
  export default {
    name: 'App'
  }
</script>

<style scoped>
  #app {
    font-family: Roboto, Helvetica, Arial, sans-serif;
  }
</style>
