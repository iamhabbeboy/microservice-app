<template>
  <div class="container">
    <div class="py-3">
      <h3 class="title">Microservice Demo <span class="fs-6 fw-normal">(Built with Vue, Go, Laravel, Kafka and MongoDB)</span></h3>
      <div>
        <div class="row align-items-start my-2">
          <div class="col">
            <h4 class="font-bold">Payload</h4>
            <div class="py-4 px-4 payload">
              <div>
                <h6>Login</h6>
                Valid details: <span class="badge rounded-pill text-bg-secondary">email: admin@gmail.com, password: verify</span> <span class="badge rounded-pill text-bg-secondary">email: user@gmail.com, password: user</span>
                <div class="mb-3">
                  <input type="email" v-model="auth.email" class="form-control" placeholder="name@example.com">
                </div>
                 <div class="mb-3">
                  <input type="password" v-model="auth.password" class="form-control" placeholder="password">
                </div>
              </div>
              <div>
                <div class="mb-3">
                  <label for="log" class="form-label">Enter Log sample</label>
                  <textarea class="form-control" id="log" v-model="log" rows="3">This is a log sample from frontend</textarea>
                </div>
              </div>
            </div>
          </div>
          <div class="col">
            <h4 class="font-bold">Output</h4>
               <div class="w-full border overflow-auto break-words output-h">
                <div v-show="isLoading">
                  <img src="./assets/loader.svg" width="100" height="100"/>
                  <p class="color-secondary">please wait</p>
                </div>
                <pre class="">
                {{response}}
                </pre>
            </div>
          </div>
        </div>
        <div class="mt-10">
          <button class=" p-3 btn btn-success rounded-md text-white" @click="handleAuth">Test Auth</button>
          <button class="mx-1 p-3 btn btn-primary rounded-md  text-white" @click="handleLog">Test Logger</button>
          <button class="p-3 btn btn-dark rounded-md text-white" @click="handleGetLogs">Get Logs</button>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped>
.title {
  color: #666;
  border-bottom: 3px solid #DDD;
}

.payload {
  background: #EEE;
  height:500px;
}

.output-h {
  height:500px;
}
</style>
<script>
 import axios from 'axios'
 import { ref, reactive } from 'vue'
export default{
  setup() {
    const response = ref()
    const log = ref('')
    const isLoading = ref(false)
    const auth = reactive({
      email: '',
      password: '',
    })
    const handleAuth = async () => {
      if(auth.email == '' || auth.password == '') {
        return alert("kindly enter email and password ")
      }
      log.value = ''
      const payload = {
        action: "auth",
        auth: {
          email: auth.email,
          password: auth.password
        }
      }
      isLoading.value = true
      response.value = undefined
      try {
        const res = await axios.post('http://localhost:8083', payload)
        isLoading.value = false
        response.value = res.data
        auth.email = ''
        auth.password = ''
      }catch(error) {
        response.value = error
      }
    }

    const handleLog = async () => {
      if(log.value == '') {
        return alert('Kindly enter a log')
      }
      auth.email = ''
      auth.password = ''
      const payload = {
        action: "log",
        log: {
          name: "log",
          data: log.value,
        }
      }
      isLoading.value = true
      response.value = undefined
      try {
        const res = await axios.post('http://localhost:8083', payload)
        isLoading.value = false
        response.value = res.data
        log.value = ''
      } catch(error) {
         response.value = error
      }
    }

    const handleGetLogs = async () => {
      const payload = {
        action: "logs",
        logs: {
        }
      }
        isLoading.value = true
        response.value = undefined
      try {
        const res = await axios.post('http://localhost:8083', payload)
        isLoading.value = false
        response.value = res.data
      }catch(error) {
        response.value = error
      }
    }

    return {
      log,
      auth,
      isLoading,
      response,
      handleAuth,
      handleLog,
      handleGetLogs,
    }
  },
}
</script>
