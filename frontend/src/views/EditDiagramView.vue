<script setup>
import { ref , onMounted } from 'vue';
import { useRoute } from 'vue-router';
import axios from 'axios';




const route = useRoute();
const diagramId = route.params.id; // Get the diagram ID from the route parameters
const diagramDetails = ref([]);


const api = axios.create({
  baseURL: 'http://localhost:8090/api',
  withCredentials: true, 
});



const fetchDiagramDetails = async (diagramId) => {
    try {
        const response = await api.get(`/fetch-diagram-by-id`, {
            params: { diagramId: diagramId }
        });

        if (response.data && response.data.diagram){
            diagramDetails.value = response.data.diagram;
        } else {
            alert('Diagram details not found.');
        }
        

    } catch (err) {
        console.error('Failed to fetch diagram details:', err);
        alert('Failed to load diagram details. Please try again.');
    }
}


onMounted(() =>{
    if (diagramId) {
        fetchDiagramDetails(diagramId);
    }
})


</script>



<template>
    <div class="container text-center p-5">
    <div class="edit-diagram-view">
        <h1 class="mb-5">Edit Diagram</h1>
        <p class="mb-5">Diagram ID: {{ diagramDetails.id }}</p>
        <pre style="color: black;">{{ diagramDetails.diagram_code }}</pre>
        <!-- Additional UI for editing the diagram can be added here -->
      

    </div>
    </div>
</template>



<style scoped>


.edit-diagram-view {
  padding: 20px;
  color: black;
}
</style>

