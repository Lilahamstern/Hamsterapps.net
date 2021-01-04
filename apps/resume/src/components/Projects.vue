<template>
  <div class="block max-w-2xl mx-auto my-4">
    <h1 class="ml-2 text-2xl font-semibold text-gray-300">Projects</h1>
    <div v-for="(project, index) in projects" :key="index" class="block max-w-3xl py-2 mx-auto text-gray-100">
      <div class="flex flex-col">
        <p class="ml-6 text-lg">
          <span class="font-semibold">{{ project.role }} - {{ project.title }}</span>, <span class="font-bold">{{ project.type }}</span>, <span class="font-semibold">{{ project.started }} - {{ project.ended }}</span>
        </p>
        <div class="block mx-auto">
          <p class="max-w-xl py-2 text-sm">{{ project.description }}</p>
          <p class="text-sm">Technologies: <span>{{ getTechnologies( project.technologies ) }} </span> </p>
          <a :href="project.link">Source code</a>
        </div>
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
      projects: [
        {
          role: "Maintiner",
          title: "Dotfiles",
          type: "",
          started: "October 2020",
          ended: "Present",
          description: "My current dotfiles does utilize ansible for installation and upgrade processes, this project is mostly mean't for my devices. But for other people to take a look to either get inspiration for their config files, or just take them.",
          technologies: ["Linux", "Ansible", "Git"],
          link: "https://github.com/TheNerdyHamster/Nerdyhamster.net",
        },
        {
          role: "Creator",
          title: "Nerdyhamster.net",
          type: "Personal Website",
          started: "June 2019",
          ended: "Present",
          description: "Nerdyhamster.net is my personal website where I will have a project portfoilio, etc etc etc etcet dwadaklwjdalkwdjl djawdlajdl daljwdajwdlka wdlkawd aklwdj alkwdj alkwdjl",
          technologies: ["Docker", "Git", "Ansible", "Vue"],
          link: "https://github.com/TheNerdyHamster/Nerdyhamster.net",
        },
        {
          role: "Developer",
          title: "CED",
          type: "Collaboration with WSP",
          started: "2019",
          ended: "Present",
          description: "CED is a project developed in collaboration with WSP where the goal is to develop an application adapted to new climate declaration that the Swedish government has given to the Housing Agency which till start to apply on 1 Jan 2021.",
          technologies: ["Docker", "Git", "Golang", "React", "Postgres"],
          link: "https://github.com/TheNerdyHamster/CED"
        },
        {
          role: "Creator",
          title: "Helmetstats",
          type: "PUBG Dev Challange Overwolf 2019",
          started: "February 2019",
          ended: "May 2019",
          description: "HelmentStats was planned to be my submission to the PUBG Dev Channalge Overwolf 2019, which did not end due to certain circumstances. Where players get live statics about their progress as well as match statistics",
          technologies: ["Typescript", "NodeJS", "Overwolf", "3rd-party API's"]
        },
        {
          role: "Developer",
          title: "Youtube Live Check",
          type: "Stream Game Labels",
          started: "2019",
          ended: "2019",
          description: "Youtube Live check was created for Stream Game Labels as a side project. Where Youtube API was used, to check when someone was last live. Technologies: NodeJS, Youtube API, Git",
          technologies: ["Javascript", "Youtube API"],
        },
        {
          role: "Developer",
          title: "Plugins, Website",
          type: "Granskogen",
          started: "April 2018",
          ended: "July 2018",
          description: "Developed both Minecraft plugins and website, to a medium Swedish Minecraft server.",
          technologies: ["Java", "Maven", "Bootstrap"],
          link: null,
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
