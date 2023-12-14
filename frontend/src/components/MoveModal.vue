<script lang="ts" setup>
import FilesystemItem from "@/dto/FilesystemItem";

// import { backendApiPost } from "@/lib/BackendApi";
// import CruddyApi from "@/lib/CruddyApi";

const props = defineProps<{
  modalID: string;
  item: FilesystemItem;
  onMoved?: () => void;
  onError?: (message: string) => void;
}>();

async function moveItem() {
  alert(props.item.name);
}
</script>

<template>
  <dialog :id="modalID" class="modal">
    <div class="modal-box w-11/12 max-w-5xl">
      <h3 class="font-bold text-lg mb-2">Element verschieben</h3>
      <div>
        Wohin soll das Element {{ item.name }} verschoben werden?
        <select class="select select-bordered w-full max-w-xs">
          <option disabled selected>Zielordner</option>
          <option>gubbel</option>
        </select>
      </div>
      <div class="modal-action">
        <form method="dialog">
          <button class="btn btn-primary mr-3" @click="moveItem">
            Verschieben
          </button>
          <button class="btn btn-secondary">Abbrechen</button>
        </form>
      </div>
    </div>
  </dialog>
</template>
