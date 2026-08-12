<script setup>
const response=await fetch(  "http://localhost:8080/api/tasks");
const task=await response.json()
console.log(task)
</script>
    <template>
    <div v-for="task in tasks" :key="task.id">
  <h3>{{ task.title }}</h3>
  <p>{{ task.description }}</p>
</div>
</template>
