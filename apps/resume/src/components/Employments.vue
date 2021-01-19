<template>
  <div class="block max-w-2xl mx-auto my-4">
    <h1 class="ml-2 text-2xl font-semibold text-gray-300">Employments</h1>
    <div v-for="(employment, index) in employments" :key="index" class="block py-2 mx-auto text-gray-100">
      <div class="flex flex-col">
        <p class="ml-6 text-lg">
          <span class="font-semibold">{{ employment.title }} - {{ employment.type }}</span>, <span class="text-base font-bold text-gray-400">{{ employment.company }}</span>, <span class="font-semibold">{{ employment.started }} - {{ employment.ended }}</span>
        </p>
        <div class="block mx-auto">
          <p class="max-w-xl py-2 text-sm">{{ employment.description }}</p>
        </div>
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
        description: "Techinal support at webinars such as Lindholmen Software Development Day, A Great Day 2020 and Lindholmen Open Days. I do also provide service for equipment and upgrade to newer. Helping colleagues with techinal questions and problem that can occur during a day."
      }] as Array<Employement>

    }
  }
});

export default Employments;
</script>
