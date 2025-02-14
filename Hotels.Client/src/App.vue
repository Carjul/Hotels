<template>
  <router-view></router-view>
</template>
<script setup>
import { onMounted } from 'vue';
import { useUsuarioStore } from './store/User';
const storeUser = useUsuarioStore()

onMounted(() => {
  if (localStorage.getItem("user")) {
    let x = JSON.parse(localStorage.getItem("user"))
    console.log(x)
    storeUser.usuario=x;
  };

})

</script>