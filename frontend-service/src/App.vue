<template>
  <div class="">
    <div class="w-10/12 mx-auto mt-10">
      <h1 class="text-2xl border-b-2">Microservice with Kafka</h1>
      <div>
        <div class="flex mt-4">
          <div class="w-6/12 h-64">
            <h3 class="font-bold">Payload</h3>
            <div class="w-full h-full bg-gray-200">
                <pre class="whitespace-pre-line">{
                  action: "auth"
                }
                </pre>
            </div>
          </div>
          <div class="w-6/12">
            <h3 class="font-bold">Output</h3>
               <div class="w-full h-full border">
                <pre class="whitespace-pre-line">
                {{response}}
                </pre>
            </div>
          </div>
        </div>
        <div class="mt-10">
           <button class="mr-3 p-3 bg-green-600 rounded-md text-white hover:bg-green-500" @click="handleAuth">Test Auth</button>
          <button class="mr-3 p-3 bg-blue-600 rounded-md text-white hover:bg-blue-500">Test Broker</button>
          <button class="mr-3 p-3 bg-purple-600 rounded-md  text-white hover:bg-purple-500" @click="handleLog">Test Logger</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
 import axios from 'axios'
 import { ref } from 'vue'
export default{
  setup() {
    const response = ref()
    const handleAuth = async () => {
      const payload = {
        action: "auth",
        auth: {
          email: "admin@gmail.com",
          password: "verify"
        }
      }
      const res = await axios.post('http://localhost:8083', payload)
      response.value = res.data
      console.log(res.data)
    }

    const handleLog = async () => {
      const payload = {
        action: "log",
        log: {
          name: "frontend",
          data: "This is a log"
        }
      }
      const res = await axios.post('http://localhost:8083', payload)
      response.value = res.data
      console.log(res.data)
    }

    return {
      response,
      handleAuth,
      handleLog,
    }
  },
}
</script>
