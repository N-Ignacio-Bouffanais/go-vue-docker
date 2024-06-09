<script setup lang="ts">
import Dashboard from './components/Dashboard.vue';

</script>

<template>
  <header>
    <nav class="flex h-24 bg-[#141b2d]">
      <ul class="flex w-[90dvw] mx-auto justify-between items-center">
        <li class="flex items-center text-4xl font-bold text-white">
          Undock Awaited
        </li>
        <li>
          <button class="text-white">
            <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 512 512">
              <path fill="currentColor"
                d="M264 480A232 232 0 0 1 32 248c0-94 54-178.28 137.61-214.67a16 16 0 0 1 21.06 21.06C181.07 76.43 176 104.66 176 136c0 110.28 89.72 200 200 200c31.34 0 59.57-5.07 81.61-14.67a16 16 0 0 1 21.06 21.06C442.28 426 358 480 264 480" />
            </svg>
          </button>
        </li>
      </ul>
    </nav>
  </header>

  <main class="flex justify-center py-6">
    <Dashboard />
  </main>
  <footer>
    <p class="text-white font-bold text-lg">Designed & Built by Nicolas Bouffanais</p>
  </footer>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Roboto:wght@400;500;700&display=swap');
html {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-size: 10px;
  font-family: 'Roboto', 'Helvetica';
}

main {
  width: 100dvw;
  background-color: #141b2d;
  height: 100dvh;
}
footer{
  background-color: #131416;
  width: 100dvw;
  height: 12rem;
  padding: 3rem 7vw;
  margin: 0 auto;
}
::-webkit-scrollbar {
  width: 1rem;
  height: 0.3rem;
}

::-webkit-scrollbar-button {
  background-color: #666;
}

::-webkit-scrollbar-thumb {
  background-color: #666;
  border-radius: 3px;
}

::-webkit-scrollbar-track-piece {
  background-color: #000;
}
</style>
