<script setup>
import { ref } from 'vue';

const isOpen = ref(true);

const navItems = ref([
  {
    name: "Sequence Diagram",
    path: "/",
    type: 'sequence'
  },
  
  {
    name: "Class Diagram",
    path: "/class-diagram",
    type: "class"
  } , 

  {
    name: "Flowchart",
    path: "/flowchart",
    type: "flowchart"
  } ,
  {
    name : 'Recent',
    path : '/recent',
    type : 'recent'
  } ,
  {
    name: "About",
    path: "/about",
    type: "about"
  } , 
  
]);

const toggleSidebar = () => {
  isOpen.value = !isOpen.value;
};

defineExpose({ isOpen });
</script>


<template>

  <button
    class="sidebar-toggle"
    :class="{ 'toggle-shifted': isOpen }"
    @click="toggleSidebar"
    :title="isOpen ? 'Close sidebar' : 'Open sidebar'"
  >
    <span class="toggle-icon" :class="{ 'rotated': !isOpen }">&#9776;</span>
  </button>

  <!-- Sidebar -->
  <aside class="sidebar" :class="{ 'sidebar-closed': !isOpen }">
    <div class="sidebar-header">
      <span class="sidebar-brand">SR Automate</span>
    </div>
    <ul>
      <li v-for="item in navItems" :key="item.name">
        <router-link :to="item.path">
          <span class="nav-label">{{ item.name }}</span>
        </router-link>
      </li>
    </ul>
  </aside>
</template>


<style scoped>
.sidebar {
  width: 220px;
  height: 100vh;
  background: linear-gradient(180deg, #1e293b 0%, #0f172a 100%);
  position: fixed;
  top: 0;
  left: 0;
  z-index: 1000;
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1), box-shadow 0.3s ease;
  box-shadow: 4px 0 24px rgba(0, 0, 0, 0.15);
  overflow: hidden;
}

.sidebar-closed {
  transform: translateX(-220px);
}

.sidebar-header {
  padding: 1.5rem 1.25rem 1rem;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
}

.sidebar-brand {
  font-family: 'Google Sans', 'Inter', sans-serif;
  font-size: 1.15rem;
  font-weight: 700;
  background: linear-gradient(135deg, #42b883, #63dda8);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  letter-spacing: 0.02em;
}

.sidebar ul {
  list-style: none;
  padding: 0.75rem 0;
  margin: 0;
}

.sidebar li {
  margin: 0.2rem 0;
}

.sidebar li a {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.7rem 1.25rem;
  color: #94a3b8;
  text-decoration: none;
  font-family: 'Google Sans', 'Inter', sans-serif;
  font-size: 0.92rem;
  font-weight: 500;
  border-radius: 8px;
  margin: 0 0.5rem;
  transition: background 0.2s ease, color 0.2s ease, transform 0.15s ease;
}

.sidebar li a:hover {
  background: rgba(66, 184, 131, 0.1);
  color: #e2e8f0;
  transform: translateX(4px);
}

.sidebar li a.router-link-exact-active {
  background: rgba(66, 184, 131, 0.15);
  color: #42b883;
  font-weight: 600;
}

.nav-icon {
  font-size: 1.1rem;
  flex-shrink: 0;
}

.nav-label {
  white-space: nowrap;
}

/* Toggle button */
.sidebar-toggle {
  position: fixed;
  top: 0.85rem;
  left: 0.75rem;
  z-index: 1100;
  width: 38px;
  height: 38px;
  border: none;
  border-radius: 10px;
  background: rgba(30, 41, 59, 0.9);
  backdrop-filter: blur(8px);
  color: #42b883;
  font-size: 1.2rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: left 0.3s cubic-bezier(0.4, 0, 0.2, 1),
              background 0.2s ease,
              transform 0.2s ease,
              box-shadow 0.2s ease;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.15);
}

.sidebar-toggle.toggle-shifted {
  left: 230px;
}

.sidebar-toggle:hover {
  background: rgba(66, 184, 131, 0.2);
  transform: scale(1.08);
  box-shadow: 0 4px 16px rgba(66, 184, 131, 0.25);
}

.sidebar-toggle:active {
  transform: scale(0.95);
}

.toggle-icon {
  display: inline-block;
  transition: transform 0.3s ease;
  line-height: 1;
}

.toggle-icon.rotated {
  transform: rotate(90deg);
}
</style>