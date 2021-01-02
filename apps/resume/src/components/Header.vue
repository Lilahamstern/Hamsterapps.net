<template>
  <div class="m-3 text-gray-50">

    <div class="flex flex-row justify-center mt-3">
      <span v-for='link in links' :key="link.name" class="flex flex-row text-xl">
        <a :href='link.url' class="link"><i :class="link.icon" class="mr-4"></i></a>
      </span>
    </div>

    <div class="flex flex-row justify-center my-3">
      <div>
        <h2 class="pb-3 text-3xl">{{ name }}</h2>
        <div class="block mx-auto rounded-lg shadow-lg w-28 h-28">
          <img src="../assets/profile.jpg" alt="Profile picture" class="rounded-lg">
        </div>
      </div>
    </div>

    <div class="flex flex-col justify-center w-full text-sm text-center">
      <p class="py-1">
        <i class="mr-2 fas fa-map-marker-alt"></i>
        {{ address }}
      </p>
      <p class="py-1">
        <i class="mr-2 fas fa-envelope"></i>
        <a :href="mail">{{ email }}</a>
      </p>
      <p class="py-1">
        <i class="mr-2 fas fa-phone"></i>
        {{ phone }}
      </p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

interface Link {
  url: String,
  name: String,
  icon: String
}

const Header = defineComponent({
  data() {
    return {
      phone: '+46 (0)73 - 834 45 66',
      address: 'Sannegårdsgatan 20, 417 61 Gothenenburg',
      mail: 'mailto:leo.ronnebro@nerdyhamster.net',
      links: [{
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
      }] as Array<Link>
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
      let mail = this.mail.split(':');
      return mail[1];
    }
  },
});

export default Header;
</script>
