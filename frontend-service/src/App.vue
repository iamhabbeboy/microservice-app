<template>
  <div class="">
    <div class="w-10/12 mx-auto mt-10">
      <h1 class="text-2xl border-b-2">Microservice with Kafka</h1>
      <div>
        <div class="flex mt-4">
          <div class="w-6/12 h-64">
            <h3 class="font-bold">Payload</h3>
            <div class="w-full h-full bg-gray-200">
                {
                  action: "broker"
                }
            </div>
          </div>
          <div class="w-6/12">
            <h3 class="font-bold">Output</h3>
               <div class="w-full h-full bg-gray-400">
                sdlfkdjf
            </div>
          </div>
        </div>
        <div class="mt-10">
           <button class="mr-3 p-3 bg-green-600 rounded-md text-white hover:bg-green-500" @click="handleAuth">Test Auth</button>
          <button class="mr-3 p-3 bg-blue-600 rounded-md text-white hover:bg-blue-500">Test Broker</button>
          <button class="mr-3 p-3 bg-purple-600 rounded-md  text-white hover:bg-purple-500">Test Logger</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
 import axios from 'axios'
export default{
  setup() {
    const handleAuth = async () => {
      const payload = {
        action: "auth",
        auth: {
          email: "admin@gmail.com",
          password: "verysecret"
        }
      }
      const res = await axios.post('http://localhost:8083', payload)
      console.log(res)
    }

    return {
      handleAuth,
    }
  },
}
</script>
