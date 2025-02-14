<script setup lang="ts">
import { onMounted } from 'vue';
import Nav from '../../components/sideBar/Nav.vue';
import { useRouter } from 'vue-router';

const router = useRouter();


onMounted(() => {
  if (localStorage.getItem("token") === null) router.push("/login");

})

</script>

<template>
  <div class="flex flex-row justify-start">
    <Nav></Nav>
  <div class="flex flex-col w-5/6 p-5">
    <router-view></router-view>
  </div>
  </div>
  
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}

.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}

.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
