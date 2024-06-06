<script setup lang="ts">

</script>

<template>
  <header>
    <nav class="flex h-20 bg-[#141627]">
      <ul class="flex w-[90dvw] mx-auto justify-between items-center">
        <li class="flex items-center text-4xl font-bold text-white">
          Undock Awaited
        </li>
        <li>
          <button class="bg-red-400"></button>
        </li>
      </ul>
    </nav>
  </header>

  <main>
    <div class="grid">

    </div>
  </main>
  <footer>
    <p class="text-white font-bold text-lg">Designed & Built by Nicolas Bouffanais</p>
  </footer>
</template>

<style>
html {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-size: 10px;
}

main {
  width: 100dvw;
  background-color: #141627;
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
