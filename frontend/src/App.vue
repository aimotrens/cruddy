<script setup lang="ts">
import ThemingButton from "./components/ThemingButton.vue";
import Config from "@/Config";
</script>

<template>
  <div class="bg-base-200 mb-10">
    <div
      class="mx-16 md:mx-24 lg:mx-48 xl:mx-96 h-24 flex items-center justify-between"
    >
      <div class="">
        <h1 class="text-3xl">{{ Config.Env.HEADER_NAME }}</h1>
      </div>

      <div class="text-lg">
        <ThemingButton />
      </div>
    </div>
  </div>

  <div class="mx-16 md:mx-24 lg:mx-48 xl:mx-96">
    <div class="text-left bg-base-00 p-4 rounded-box">
      <router-view />
    </div>
  </div>
</template>
