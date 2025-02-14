<template>
  <div class="min-h-screen bg-gradient-to-br from-indigo-500 via-purple-500 to-pink-500">
    <div class="container mx-auto px-4 h-screen flex items-center justify-center">
      <div class="max-w-4xl mx-auto text-center">
        <!-- Hero Content -->
        <h1 class="text-5xl md:text-6xl font-bold text-white mb-8 animate-fade-in">
          Gestión Hotelera Simplificada
        </h1>
        
        <p class="text-xl md:text-2xl text-white/90 mb-12 max-w-2xl mx-auto">
          Optimiza la gestión de tu hotel con nuestra plataforma intuitiva y profesional
        </p>

        <!-- CTA Button -->
        <router-link 
          to="/home/dashboard" 
          class="inline-flex items-center px-8 py-4 text-lg font-semibold text-indigo-600 bg-white rounded-lg shadow-lg hover:bg-indigo-50 transition duration-300 transform hover:scale-105"
        >
          <span>Comenzar ahora</span>
          <svg 
            xmlns="http://www.w3.org/2000/svg" 
            class="h-6 w-6 ml-2" 
            fill="none" 
            viewBox="0 0 24 24" 
            stroke="currentColor"
          >
            <path 
              stroke-linecap="round" 
              stroke-linejoin="round" 
              stroke-width="2" 
              d="M13 7l5 5m0 0l-5 5m5-5H6" 
            />
          </svg>
        </router-link>

        <!-- Features Grid -->
        <div class="grid md:grid-cols-3 gap-8 mt-20">
          <div class="bg-white/10 backdrop-blur-lg rounded-xl p-6 text-white">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mx-auto mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
            <h3 class="text-xl font-semibold mb-2">Rápido y Eficiente</h3>
            <p class="text-white/80">Gestiona reservas y habitaciones en tiempo real</p>
          </div>

          <div class="bg-white/10 backdrop-blur-lg rounded-xl p-6 text-white">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mx-auto mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
            </svg>
            <h3 class="text-xl font-semibold mb-2">Seguro y Confiable</h3>
            <p class="text-white/80">Datos protegidos con la última tecnología</p>
          </div>

          <div class="bg-white/10 backdrop-blur-lg rounded-xl p-6 text-white">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mx-auto mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
            </svg>
            <h3 class="text-xl font-semibold mb-2">Fácil de Usar</h3>
            <p class="text-white/80">Interfaz intuitiva y amigable</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.animate-fade-in {
  animation: fadeIn 1s ease-out;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>