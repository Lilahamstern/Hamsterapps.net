<template>
  <div class="h-screen">
    <Header :name='name'/>
    <Employments />
    <Projects />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from './components/Header.vue'
import Employments from './components/Employments.vue'
import Projects from './components/Projects.vue'

const App = defineComponent({
  components: {
    Header,
    Employments,
    Projects,
  },
  data() {
    return {
      name: "Leo Rönnebro"
    }
  },
});

export default App;
</script>
