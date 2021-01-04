<template>
  <div class="block max-w-2xl mx-auto my-4">
    <h1 class="ml-2 text-2xl font-semibold text-gray-300">Education</h1>
    <div v-for="(education, index) in educations" :key="index" class="block max-w-3xl py-2 mx-auto text-gray-100">
      <div class="flex flex-col">
        <p class="ml-6 text-lg">
        <span class="font-semibold">{{ education.school }}</span>, <span class="font-semibold">{{ education.start }} - {{ education.graduation }}</span>, <span>{{ education.title }}</span>
        </p>
        <div class="block mx-auto">
          <p class="max-w-xl py-2 text-sm">{{ education.description }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

interface Education {
  school: String,
  title: String,
  start: String,
  graduation: String,
  description: String
}

const Education = defineComponent({
  data() {
    return {
      educations: [{
        school: "Tekniska Yrkeshögskolan",
        title: ".NET Developer",
        start: "2020",
        graduation: "2022",
        description: "Ipsam similique atque et voluptatem quia accusantium consequatur iusto. Sunt natus dolorem est. Cum ea non voluptatem atque. Doloribus eos harum minima. Adipisci maiores omnis incidunt eos vel quia qui.…"
      },
      {
        school: "LBS Kreativa Gymnasiet",
        title: "Upper Secondary School - Diploma",
        start: "2017",
        graduation: "2020",
        description: "During my school time we focused on both product development and project management, technologies and tools used during studies was Go, C#, Vue, Postgres, Docker"
      }] as Array<Education>

    }
  },

  created() {
    console.log(this.educations)
  }
});

export default Education;
</script>
