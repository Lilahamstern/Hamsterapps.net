<template>
  <div class="h-screen">
    <Header :name='name'/>
    <Employments />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from './components/Header.vue'
import Employments from './components/Employments.vue'

const App = defineComponent({
  components: {
    Header,
    Employments
  },
  data() {
    return {
      name: "Leo Rönnebro"
    }
  },
});

export default App;
</script>
