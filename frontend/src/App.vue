<template>
  <div id="app" class="container mt-5">
    <div class="row">
      <div v-for="(developer, index) in developers" :key="index" class="col">
        <div class="card">
          <div class="card-header">
            <h2 class="card-title">{{ developer.name }}</h2>
          </div>
          <div class="card-body">
            <div v-for="(task, index) in developer.tasks" :key="index" class="card my-2">
              <div class="card-body">
                <h3 class="card-title">{{ task.name }}</h3>
                <p class="card-text">{{ task.duration }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="mt-3">
      <p>Toplam süre: {{ totalWeeks }} hafta</p>
    </div>
  </div>
</template>


<script>
import axios from 'axios';


export default {
  name: 'App',
  data() {
    return {
      developers: [
      ],
      totalWeeks: null
    }
  },
  created() {
    console.log(process.env.VUE_APP_SERVICE_URL);
    axios.get(`${process.env.VUE_APP_SERVICE_URL}/tasks`)
      .then(response => {
        this.developers = response.data.developers;
        this.totalWeeks = response.data.totalWeek;
      })
      .catch(error => {
        console.error(error);
      });
  },
}
</script>

<style scoped>
.developers {
  display: flex;
  justify-content: space-around;
}
.developer {
  border: 1px solid #000;
  margin: 10px;
  padding: 10px;
  max-width: 200px;
  max-height: 600px;
  overflow-y: auto;
}

.card-body {
  max-height: 300px; /* Maksimum yükseklik */
  overflow-y: auto; /* Dikey kaydırma çubuğu */
}

.task {
  background-color: #f9f9f9;
  margin: 5px 0;
  padding: 5px;
}
.week-info {
  margin-top: 20px;
  font-size: 20px;
  font-weight: bold;
  color: #333;
}
</style>
