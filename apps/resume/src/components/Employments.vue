<template>
  <h1 class="pb-2 ml-2 text-2xl font-semibold text-center text-gray-300">Employments</h1>
  <div v-for="(employment, index) in employments" :key="index" class="block max-w-3xl mx-auto text-gray-100">
    <div class="flex flex-col">
      <p class="text-lg text-center">
        <span class="font-semibold">{{ employment.title }} - {{ employment.type }}</span>, <span class="font-bold">{{ employment.company }}</span>, <span class="font-semibold">{{ employment.started }} - {{ employment.ended }}</span>
      </p>
      <div class="block mx-auto">
        <p class="max-w-xl py-2 text-sm">{{ employment.description }}</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

interface Employement {
  title: String,
  type: String,
  company: String,
  started: String,
  ended: String,
  description: String
}

const Employments = defineComponent({

  data() {
    return {
      employments: [{
        title: "IT Support",
        type: "Part time",
        company: "Lindholmen Science Park",
        started: "March 2020",
        ended: "Present",
        description: "Ipsam similique atque et voluptatem quia accusantium consequatur iusto. Sunt natus dolorem est. Cum ea non voluptatem atque. Doloribus eos harum minima. Adipisci maiores omnis incidunt eos vel quia qui.…"
      }] as Array<Employement>

    }
  }
});

export default Employments;
</script>
