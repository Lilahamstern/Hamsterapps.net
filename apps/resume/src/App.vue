<template>
  <div class="h-screen grid grid-cols-3">
    <Sidebar :name='name'/>
    <Main />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Sidebar from './components/Sidebar.vue'
import Main from './components/Main.vue'

const App = defineComponent({
  components: {
    Sidebar,
    Main
  },
  data() {
    return {
      name: "Leo Rönnebro"
    }
  },
});

export default App;
</script>
