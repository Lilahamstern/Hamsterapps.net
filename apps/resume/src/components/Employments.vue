<template>
  <div v-for="(employment, index) in employments" :key="index" class="text-gray-100">
    <div class="flex flex-col justify-items-center">
      <p class="text-lg">
        <span class="font-semibold">{{ employment.title }} - {{ employment.type }}</span>, <span class="font-bold">{{ employment.company }}</span>, <span class="font-semibold">{{ employment.started }} - {{ employment.ended }}</span>
      </p>
      <p class="w-3/4 p-2">{{ employment.description }}</p>
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
