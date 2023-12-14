<script lang="ts" setup>
import FilesystemItem from "@/dto/FilesystemItem";

import { backendApiPost } from "@/lib/BackendApi";
import CruddyApi from "@/lib/CruddyApi";

const props = defineProps<{
  modalID: string;
  item: FilesystemItem;
  onDeleted?: () => void;
  onError?: (message: string) => void;
}>();

async function deleteItem() {
  await backendApiPost(CruddyApi.Delete(props.item.filePath), {
    onSuccess: () => {
      props.onDeleted?.();
    },
    onError: (error?: Error) => {
      if (error) {
        props.onError?.(error.message);
      }
    },
  });
}
</script>

<template>
  <dialog :id="modalID" class="modal">
    <div class="modal-box w-11/12 max-w-5xl">
      <h3 class="font-bold text-lg mb-2">Element löschen</h3>
      <div>Soll das Element {{ item.name }} wirklich gelöscht werden?</div>
      <div class="modal-action">
        <form method="dialog">
          <button class="btn btn-primary mr-3" @click="deleteItem">
            Löschen
          </button>
          <button class="btn btn-secondary">Abbrechen</button>
        </form>
      </div>
    </div>
  </dialog>
</template>
