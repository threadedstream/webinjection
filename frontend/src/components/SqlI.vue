<template>
  <MDBInput label="Type in product name" type="text" v-model="input1"></MDBInput>
</template>

<script setup lang="ts">
  import {MDBInput} from 'mdb-vue-ui-kit'
  import {ref} from 'vue'

  const input1 = ref('')
</script>

<script lang="ts">
  export default {
    name: "SqlI"
  }
</script>

<style scoped>

</style>