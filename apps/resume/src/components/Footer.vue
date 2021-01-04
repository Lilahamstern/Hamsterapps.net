<template>
  <div class="flex justify-around w-full max-w-2xl mx-auto text-gray-300">
    <div class="w-full">
      <h1 class="pb-4 ml-2 text-2xl font-semibold">References</h1>
      <div>
        <p class="ml-6 text-lg">Provided uppon request.</p>
      </div>
    </div>
    <div class="w-full text-right">
      <h1 class="pb-4 mr-2 text-2xl font-semibold">Languages</h1>
      <div>
        <p class="mr-6 text-lg">Swedish - <span>Spoken, Written</span></p>
        <p class="mr-6 text-lg">English - <span>Spoken, Written</span></p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

const Footer  = defineComponent({

  data() {
    return {
    }
  },
  methods: {
  }
});

export default Footer;
</script>
