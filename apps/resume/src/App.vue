<template>
  <div class="h-screen">
    <Header :name='name'/>
    <Employments />
    <Projects />
    <Educations />
    <Footer />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from './components/Header.vue'
import Employments from './components/Employments.vue'
import Projects from './components/Projects.vue'
import Educations from './components/Education.vue'
import Footer from './components/Footer.vue'

const App = defineComponent({
  components: {
    Header,
    Employments,
    Projects,
    Educations,
    Footer,
  },
  data() {
    return {
      name: "Leo Rönnebro"
    }
  },
});

export default App;
</script>
