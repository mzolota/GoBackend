<script>
import { ref, onMounted, onBeforeUnmount } from 'vue'
import './style.css'  // globalni CSS

export default {
  name: 'Chat',
  setup() {
    const ws = ref(null)             // WebSocket instanca
    const messages = ref([])         // Lista primljenih poruka
    const newMessage = ref('')       // Poruka za slanje

    // Funkcija za slanje poruka
    const sendMessage = () => {
      if (newMessage.value.trim() !== '' && ws.value.readyState === WebSocket.OPEN) {
        ws.value.send(newMessage.value)
        newMessage.value = ''
      }
    }

    // Pokretanje WebSocket konekcije
    onMounted(() => {
      ws.value = new WebSocket('ws://localhost:8080/ws')

      ws.value.onmessage = (event) => {
        messages.value.push(event.data)
      }

      ws.value.onopen = () => {
        console.log('Connected to WebSocket server')
      }

      ws.value.onclose = () => {
        console.log('Disconnected from WebSocket server')
      }

      ws.value.onerror = (err) => {
        console.error('WebSocket error:', err)
      }
    })

    // Zatvaranje konekcije prilikom odlaska sa stranice
    onBeforeUnmount(() => {
      if (ws.value) ws.value.close()
    })

    return {
      messages,
      newMessage,
      sendMessage,
    }
  },
}
</script>

<template>
  <div class="p-6">
    <h1 class="text-3xl font-bold mb-4">Live Chat</h1>

    <div class="chat-box border p-4 h-64 overflow-y-auto mb-4">
      <div v-for="(msg, index) in messages" :key="index">
        {{ msg }}
      </div>
    </div>

    <input
      v-model="newMessage"
      @keyup.enter="sendMessage"
      type="text"
      placeholder="Upiši poruku..."
      class="border p-2 w-full mb-2"
    />
    <button @click="sendMessage" class="bg-blue-500 text-white p-2 rounded">
      Pošalji
    </button>
  </div>
</template>



