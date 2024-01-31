<script lang="ts" setup>
import { onMounted, ref, watch } from "vue";
import { useRoute } from "vue-router";
import { backendApiGet } from "@/lib/BackendApi";
import CruddyApi from "@/lib/CruddyApi";
import type FilesystemItem from "@/dto/FilesystemItem";
import moment from "moment";
import DeleteModal from "@/components/DeleteModal.vue";
import MoveModal from "@/components/MoveModal.vue";
import CopyModal from "@/components/CopyModal.vue";

const dirContent = ref<FilesystemItem[]>();
const route = useRoute();

onMounted(() => {
  loadItems();
});

watch(
  () => route.query.path,
  () => {
    loadItems();
  },
);

async function loadItems() {
  await backendApiGet(CruddyApi.List((route.query.path as string) ?? "/"), {
    onSuccess: (data) => {
      const items = data as FilesystemItem[];

      items.sort((a: FilesystemItem, b: FilesystemItem) => {
        if (a.isDir && !b.isDir) {
          return -1;
        }
        if (!a.isDir && b.isDir) {
          return 1;
        }
        if (a.name < b.name) {
          return -1;
        }
        if (a.name > b.name) {
          return 1;
        }
        return 0;
      });

      dirContent.value = items;
    },
  });
}

function getByteUnit(bytes: number) {
  const units = ["Byte", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];
  let l = 0;
  let n = bytes || 0;

  while (n >= 1024 && ++l) {
    n /= 1024;
  }

  return `${n.toFixed(n >= 10 || l < 1 ? 0 : 1)} ${units[l]}`;
}

function generateBreadcrumbs(path: string) {
  const parts = path.split("/");
  const breadcrumbs = [{ name: "/", path: "/" }];

  for (let i = 0; i < parts.length; i++) {
    const part = parts[i];

    if (part === "") {
      continue;
    }

    breadcrumbs.push({
      name: part,
      path: parts.slice(0, i + 1).join("/"),
    });
  }

  return breadcrumbs;
}

function showModal(modalID: string) {
  const modal = document.getElementById(modalID) as HTMLDialogElement;
  modal?.showModal();
}

function onItemOperationError(error: string) {
  alert(error);
}
</script>

<template>
  <div class="breadcrumbs">
    <ul>
      <li v-for="p in generateBreadcrumbs((route.query.path as string) ?? '/')" :key="p.path">
        <router-link :to="{ name: 'Home', query: { path: p.path } }">
          {{ p.name }}
        </router-link>
      </li>
    </ul>
  </div>

  <hr class="my-5" />

  <table class="table">
    <tr>
      <th class=""></th>
      <th class="w-full">Name</th>
      <th class=""></th>
      <th class="">Größe</th>
      <th class="">Änderung</th>
    </tr>

    <tr v-for="item in dirContent" :key="item.name" class="hover">
      <td class="">
        <i class="fa" :class="{ 'fa-file': !item.isDir, 'fa-folder': item.isDir }" />
      </td>

      <td class="">
        <router-link v-if="item.isDir" :to="{ name: 'Home', query: { path: item.filePath } }" class="hover:underline">
          {{ item.name }}
        </router-link>

        <a v-if="!item.isDir" :href="CruddyApi.GetFullUrl(CruddyApi.Download(item.filePath))" target="_blank"
          rel="noopener noreferrer" class="hover:underline">
          {{ item.name }}
        </a>
      </td>

      <td class="">
        <div class="dropdown">
          <label tabindex="0" class="px-3 m-1 cursor-pointer"><i class="fa fa-ellipsis-vertical" /></label>
          <ul tabindex="0" class="dropdown-content z-[1] menu p-2 shadow bg-base-100 rounded-box w-52">
            <li>
              <button @click="showModal('moveModal' + item.name)" disabled>
                <i class="fa fa-file-export" />
                Verschieben
              </button>
            </li>
            <li>
              <button @click="showModal('copyModal' + item.name)" disabled>
                <i class="fa fa-copy" />
                Kopieren
              </button>
            </li>
            <li>
              <button @click="showModal('deleteModal' + item.name)">
                <i class="fa fa-trash" />
                Löschen
              </button>
            </li>
          </ul>
        </div>
      </td>

      <td class="whitespace-nowrap">
        <span v-if="!item.isDir">{{ getByteUnit(item.size) }}</span>
      </td>

      <td class="whitespace-nowrap">
        <span v-if="!item.isDir">{{
          moment(item.changed).format("DD.MM.YYYY HH:mm")
        }}</span>
      </td>

      <MoveModal :modalID="'moveModal' + item.name" :item="item" :onMoved="loadItems" :onError="onItemOperationError" />
      <CopyModal :modalID="'copyModal' + item.name" :item="item" :onCopied="loadItems" :onError="onItemOperationError" />
      <DeleteModal :modalID="'deleteModal' + item.name" :item="item" :onDeleted="loadItems"
        :onError="onItemOperationError" />
    </tr>
  </table>
</template>
