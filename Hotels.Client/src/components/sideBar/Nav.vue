<template>
  <!-- Sidebar -->
  <div class="flex flex-col w-1/6">
    <!-- Sidebar -->
    <div class="fixed inset-y-0 left-0  bg-gradient-to-b from-indigo-500 to-indigo-800 text-white transition-all duration-300 transform shadow-lg">
      <!-- Logo -->
      <div class="flex items-center justify-center h-16 border-b border-indigo-500">
        <router-link to="/" class="flex items-center space-x-2 px-4">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" viewBox="0 0 24 24" fill="currentColor">
            <path d="M4 21V9l8-6 8 6v12h-6v-7h-4v7H4z M12 3.5l6 4.5v2H6v-2l6-4.5z" />
          </svg>
          <span class="text-xl font-bold">Hotels Manager</span>
        </router-link>
      </div>

      <!-- Navigation Items -->
      <nav class="mt-8">
        <router-link to="/home/dashboard" class="flex items-center px-6 py-3 text-gray-100 hover:bg-indigo-700 transition-colors duration-200"
          :class="{ 'bg-indigo-700': $route.path === '/home/dashboard' }">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
          </svg>
          <span>Dashboard</span>
        </router-link>

        <router-link to="/home/hotels" class="flex items-center px-6 py-3 text-gray-100 hover:bg-indigo-700 transition-colors duration-200"
          :class="{ 'bg-indigo-700': $route.path === '/home/hotels' }">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
          </svg>
          <span>Hotels</span>
        </router-link>

        <router-link to="/home/rooms" class="flex items-center px-6 py-3 text-gray-100 hover:bg-indigo-700 transition-colors duration-200"
          :class="{ 'bg-indigo-700': $route.path === '/home/rooms' }">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
          </svg>
          <span>Rooms</span>
        </router-link>

        <!-- Profile and Logout Section -->
        <div class="absolute bottom-0 w-full border-t border-indigo-500">
          <div class="px-6 py-4">
            <div class="flex items-center space-x-3 mb-4">
              <div class="h-8 w-8 rounded-full bg-indigo-500 flex items-center justify-center">
                <img class="h-8 w-8 rounded-full" :src="usuario.imagen" alt="img">
              </div>
              <span class="text-sm font-medium">{{ usuario.nombre }}</span>
            </div>
            <button @click="Logout" 
              class="w-full flex items-center px-4 py-2 text-gray-100 hover:bg-indigo-700 rounded transition-colors duration-200">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
              </svg>
              <span>Logout</span>
            </button>
          </div>
        </div>
      </nav>
    </div>
  </div>
</template>

<script setup lang="ts">
import {useUsuarioStore} from '../../store/User'
const {LogOut, usuario} = useUsuarioStore()

function Logout() {
  LogOut();
}
</script>