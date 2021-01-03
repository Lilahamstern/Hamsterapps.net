<template>
  <div class="h-screen">
    <Header :name='name'/>
    <Employments />
    <Projects />
    <Educations />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from './components/Header.vue'
import Employments from './components/Employments.vue'
import Projects from './components/Projects.vue'
import Educations from './components/Education.vue'

const App = defineComponent({
  components: {
    Header,
    Employments,
    Projects,
    Educations,
  },
  data() {
    return {
      name: "Leo Rönnebro"
    }
  },
});

export default App;
</script>
