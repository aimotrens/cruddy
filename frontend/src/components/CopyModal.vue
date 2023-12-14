<script lang="ts" setup>
import FilesystemItem from "@/dto/FilesystemItem";

const props = defineProps<{
  modalID: string;
  item: FilesystemItem;
  onCopied?: () => void;
  onError?: (message: string) => void;
}>();

async function copyItem() {
  alert(props.item.name);
}
</script>

<template>
  <dialog :id="modalID" class="modal">
    <div class="modal-box w-11/12 max-w-5xl">
      <h3 class="font-bold text-lg mb-2">Element löschen</h3>
      <div>Soll das Element {{ item.name }} nach ... kopiert werden?</div>
      <div class="modal-action">
        <form method="dialog">
          <button class="btn btn-primary mr-3" @click="copyItem">
            Verschieben
          </button>
          <button class="btn btn-secondary">Abbrechen</button>
        </form>
      </div>
    </div>
  </dialog>
</template>
