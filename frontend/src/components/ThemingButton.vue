<script lang="ts" setup>
import { onMounted } from "vue";

const themes = [
  "light",
  "dark",
  "cupcake",
  "bumblebee",
  "emerald",
  "corporate",
  "synthwave",
  "retro",
  "cyberpunk",
  "valentine",
  "halloween",
  "garden",
  "forest",
  "aqua",
  "lofi",
  "pastel",
  "fantasy",
  "wireframe",
  "black",
  "luxury",
  "dracula",
  "cmyk",
  "autumn",
  "business",
  "acid",
  "lemonade",
  "night",
  "coffee",
  "winter",
];

onMounted(() => {
  setupTheme();
});

function getAllThemes() {
  return themes;
}

function setupTheme(): void {
  if (!localStorage.getItem("theme")) {
    localStorage.setItem("theme", "light");
  }

  setTheme(getCurrentTheme());
}

function setTheme(themeName: string): void {
  document.documentElement.setAttribute("data-theme", themeName);
  localStorage.setItem("theme", themeName);
}

function getCurrentTheme(): string {
  return localStorage.getItem("theme") ?? "light";
}
</script>

<template>
  <button class="" onclick="themeModal.showModal()">
    <i class="fa fa-pen-nib" />
  </button>

  <dialog id="themeModal" class="modal">
    <div class="modal-box w-11/12 max-w-5xl">
      <h3 class="font-bold text-lg mb-2">Wähle dein bevorzugtes Design!</h3>
      <div>
        <div class="rounded-box grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5">
          <div v-for="t in getAllThemes()" :key="t"
            class="border-base-content/20 hover:border-base-content/40 overflow-hidden rounded-lg border outline outline-2 outline-offset-2 outline-transparent"
            @click="setTheme(t)">
            <div :data-theme="t" class="bg-base-100 text-base-content w-full cursor-pointer font-sans">
              <div class="grid grid-cols-5 grid-rows-3">
                <div class="bg-base-200 col-start-1 row-span-2 row-start-1"></div>
                <div class="bg-base-300 col-start-1 row-start-3"></div>
                <div class="bg-base-100 col-span-4 col-start-2 row-span-3 row-start-1 flex flex-col gap-1 p-2">
                  <div class="font-bold">{{ t }}</div>
                  <div class="flex flex-wrap gap-1">
                    <div class="bg-primary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                      <div class="text-primary-content text-sm font-bold">
                        A
                      </div>
                    </div>
                    <div class="bg-secondary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                      <div class="text-secondary-content text-sm font-bold">
                        A
                      </div>
                    </div>
                    <div class="bg-accent flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                      <div class="text-accent-content text-sm font-bold">A</div>
                    </div>
                    <div class="bg-neutral flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                      <div class="text-neutral-content text-sm font-bold">
                        A
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="modal-action">
        <form method="dialog">
          <button class="btn">Schließen</button>
        </form>
      </div>
    </div>
  </dialog>
</template>
