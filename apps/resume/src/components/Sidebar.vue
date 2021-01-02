<template>
  <div class="pt-10 text-center text-gray-100 bg-gray-900">
    <div class="block m-auto mb-10 rounded-full shadow-lg w-44 h-44">
      <img src="../assets/profile.jpg" alt="Profile picture" class="rounded-full">
    </div>

    <h2 class="pb-5 text-3xl">{{ name }}</h2>

    <div class="block w-1/2 pb-5 mx-auto text-left">
      <h3 class="pb-3 text-2xl text-center">Contact</h3>
      <p>{{ address }}</p>
      <p>{{ phone }}</p>
      <p><a :href='mail'>{{ email }}</a></p>
    </div>

    <p v-for='link in links' :key="link.name" class="text-lg">
      <i :class="link.icon" class="mr-4"></i>
      <a :href='link.url' class="link">{{ link.name }}</a>
    </p>

  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

interface Link {
  url: String,
  name: String,
  icon: String
}

const Sidebar = defineComponent({
  data() {
    return {
      phone: '+46 (0)73 - 834 45 66',
      address: 'Sannegårdsgatan 20, 417 61 Gothenburg',
      mail: 'mailto:leo.ronnebro@nerdyhamster.net',
      links: new Array<Link>()
    }
  },
  props: {
    name: {
      type: String,
      required: true
    }
  },
  computed: {
    email(): String {
      let s = this.mail.split(':');
      return s[1];
    }

  },
  created() {
    this.links = [{
      url: "https://github.com/TheNerdyHamster",
      name: "Github",
      icon: "fab fa-github"
    },
    {
      url: "https://linkedin.com",
      name: "Linkedin",
      icon: "fab fa-linkedin"
    },
    {
      url: "https://nerdyhamster.net",
      name: "Website",
      icon: "fas fa-globe"
    }
    ]
  }
});

export default Sidebar;
</script>
