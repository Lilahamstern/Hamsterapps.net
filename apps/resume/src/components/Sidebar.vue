<template>
  <div class="text-center text-gray-100 bg-gray-900">
    <h2 class="text-3xl">{{ name }}</h2>
    <h3>Contact</h3>
    <p>Phone</p>
    <p>Mail</p>
    <p>Linkedin</p>
    <p>Website</p>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

const Sidebar = defineComponent({
  props: {
    name: {
      type: String,
      required: true
    }
  }
});

export default Sidebar;
</script>
