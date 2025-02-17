<script setup lang="ts">
import { useHotelsStore } from '../../../store/Hotel';

const {  hotelsUser, GetHotel } = useHotelsStore()


</script>
<template>
  <h1 class="text-3xl font-bold underline">
    Hello hotels!
  </h1>



  <div v-for="item in hotelsUser">
    <p class="text-3xl font-bold "> {{ item.name }}</p>
    <button class="btn w-16 h-12 bg-blue-400" v-on:click="GetHotel(item._id)">ir</button>
  </div>
</template>
<style scoped></style>