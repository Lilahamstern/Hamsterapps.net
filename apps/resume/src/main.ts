import { createApp } from 'vue'
import App from './App.vue'
import './index.css'

// Fontawesome
import { library, dom } from '@fortawesome/fontawesome-svg-core'
import { faVuejs } from '@fortawesome/free-brands-svg-icons'
import { fas } from '@fortawesome/free-solid-svg-icons'
library.add(faVuejs, fas)

// Transform <i> to <svg>
dom.watch()

createApp(App).mount('#app')

