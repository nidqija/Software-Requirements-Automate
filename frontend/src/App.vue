<script setup>
import { ref } from 'vue'
import { RouterLink, RouterView } from 'vue-router'
import SideBarView from '@/components/SideBarView.vue'

const sidebarRef = ref(null)
</script>

<template>
  <div class="layout">
    <SideBarView ref="sidebarRef" />
    <main
      class="content"
      :class="{ 'content-expanded': sidebarRef && !sidebarRef.isOpen }"
    >
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
.layout {
  display: flex;
  min-height: 100vh;
}

.content {
  margin-left: 220px;
  width: 100%;
  transition: margin-left 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.content-expanded {
  margin-left: 0;
}
</style>
