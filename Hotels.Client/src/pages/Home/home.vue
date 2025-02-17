<script setup lang="ts">
import { onMounted } from 'vue';
import Nav from '../../components/sideBar/Nav.vue';
import { useRouter } from 'vue-router';
import { useHotelsStore } from '../../store/Hotel';
import { useUsuarioStore } from '../../store/User';
const { GetHotelsOfUser} = useHotelsStore()
const {usuario} = useUsuarioStore()
const router = useRouter();


onMounted(() => {
  if (localStorage.getItem("token") === null) {
    router.push("/login");
  } else{
    usuario.hotelId? GetHotelsOfUser(usuario.hotelId): null;
  }

})

</script>

<template>
  <div class="flex flex-col w-full">
    <Nav></Nav>
    <div class="flex flex-col w-full p-5">
      <router-view></router-view>
    </div>
  </div>

</template>

<style scoped>

</style>
