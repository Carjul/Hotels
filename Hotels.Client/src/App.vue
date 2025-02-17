<template>
  <router-view></router-view>
</template>

<script setup lang="ts">
import { onBeforeMount, onMounted } from 'vue';
import { useUsuarioStore } from './store/User';
const handleBeforeUnload = (e: BeforeUnloadEvent) => {
  e.preventDefault();
  e.returnValue = '';
  return '';
};

onBeforeMount(() => {
  // Add event listener for page reload
  window.addEventListener('beforeunload', handleBeforeUnload);

});

const storeUser = useUsuarioStore()
onMounted(() => {
  var ls: string | null  = localStorage.getItem("user");
  if (ls) {
    let user = JSON.parse(ls)
    storeUser.usuario= user;
  };
})
</script>