<template>
  <h1 class="pb-2 ml-2 text-2xl font-semibold text-center text-gray-300">Projects</h1>
  <div v-for="(project, index) in projects" :key="index" class="block max-w-3xl mx-auto text-gray-100">
    <div class="flex flex-col">
      <p class="text-lg text-center">
        <span class="font-semibold">{{ project.role }} - {{ project.title }}</span>, <span class="font-bold">{{ project.type }}</span>, <span class="font-semibold">{{ project.started }} - {{ project.ended }}</span>
      </p>
      <div class="block mx-auto">
        <p class="max-w-xl py-2 text-sm">{{ project.description }}</p>
        <p class="text-sm">Technologies: <span>{{ getTechnologies( project.technologies ) }} </span> </p>
        <a :href="project.link">Source code</a>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

interface Project {
  role: String,
  title: String,
  type: String,
  started: String,
  ended: String,
  description: String,
  technologies: Array<String>,
  link?: String
}

const Projects  = defineComponent({

  data() {
    return {
      projects: [{
        role: "Creator",
        title: "Nerdyhamster.net",
        type: "Personal Website",
        started: "2019",
        ended: "Present",
        description: "Nerdyhamster.net is my personal website where I will have a project portfoilio, etc etc etc etcet dwadaklwjdalkwdjl djawdlajdl daljwdajwdlka wdlkawd aklwdj alkwdj alkwdjl",
        technologies: ["Docker", "Git", "Ansible", "Vue"],
        link: "https://github.com/TheNerdyHamster/Nerdyhamster.net"
      }] as Array<Project>

    }
  },
  methods: {
    getTechnologies: function(technologies: Array<String>): String {
      if (technologies === null) {
        return "";
      }

      return technologies.join(", ");
    }
  }
});

export default Projects;
</script>
