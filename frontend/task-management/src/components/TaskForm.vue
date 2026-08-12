<script setup>
const title=ref('')
const descriptipon=ref('')
const complete=ref(false)
async function GetTask(){
    const response=await fetch("http://localhost:8080/api/tasks", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            title: title.value,
            description: descriptipon.value,
            complete: complete.value
        })
    });
    const data=await response.json()
    console.log(data)
}
</script>
