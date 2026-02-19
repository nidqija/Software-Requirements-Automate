<script setup>
import { onMounted, ref } from 'vue';
import mermaid from 'mermaid';
import axios from 'axios';


const recentDiagrams = ref([]);
const loading = ref(true);
const error = ref(null);


const api = axios.create({
  baseURL: 'http://localhost:8090/api',
  withCredentials: true, 
});


const formatDate = (dateString) => {

  const options = { 
    year: 'numeric',
    month: 'long',
    day: 'numeric' 
};

  return new Date(dateString).toLocaleDateString(undefined, options);

};

// fetching recent diagrams from backend API upon successful generation 

const fetchRecentDiagrams = async () => {
  try {
    const response = await api.get('http://localhost:8090/api/fetch-diagrams');
    recentDiagrams.value = response.data.diagrams || [];
  } catch (err) {
    error.value = 'Failed to load recent diagrams.';
    console.error(err);
  } finally {
    loading.value = false;
  }
};

// Fetch recent diagrams when the component is mounted (reloaded)

onMounted(() => {
  fetchRecentDiagrams();
});

// Function to copy text to clipboard

const copyToClipboard = (text) => {
  navigator.clipboard.writeText(text)
    .then(() => {
      alert('Diagram code copied to clipboard!');
    })
    .catch((err) => {
      console.error('Failed to copy text: ', err);
    });
};


// Function to download image

const downloadImage = async (imageUrl , diagramType , extension = 'svg') =>{
   try {
    const response = await fetch(imageUrl);
    const blob = await response.blob();
    const url = window.URL.createObjectURL(blob);

    const link = document.createElement('a');
    link.href = url;

    const filename = `diagram_${diagramType}_${Date.now()}.${extension}`;

    link.setAttribute('download', filename);
    document.body.appendChild(link);
    link.click();

    link.remove();
    window.URL.revokeObjectURL(url);
   } catch(err){
    console.error('Failed to download image: ', err);
    alert('Failed to download image. Please try again.');
   }
}




</script>



<template>
  <main class="recent">
  <section class="hero">
  <div class="recent-view">
    <h1 class="title">Recent Diagrams</h1>
    <p>Here you can view and manage your recently created diagrams.</p>
   
    <div v-if="loading" class="loading">
       <p>Loading recent diagrams...</p> 
    </div>

   


    <!-- Placeholder for recent diagrams list -->
    <div class="recent-diagrams">
      <h5 class="recentTitle">Recent</h5>
       <div v-if="!loading && recentDiagrams.length == 0" class="no-data">
         <p class="no-diagram">No recent diagrams found. Start creating your first diagram!</p>
       </div>

       <div class="content-grid" v-else>
         <div class="card" v-for="diagram in recentDiagrams" :key="diagram.id">
          <div>
           <p style="padding: 10px; text-align: start;">{{ formatDate(diagram.created) }}</p>
           <p style="text-align: start; padding: 10px;">{{ diagram.user_prompt }}</p>
           <p class="diagram-type">{{ diagram.diagram_type }}</p>
           <div class="dropdown flex justify-content-end" style="position: absolute; top: 10px; right: 10px;">
             <button class="dropdown-toggle" type="button" data-bs-toggle="dropdown" aria-expanded="false">
              ...
             </button>
           <ul class="dropdown-menu">
              <li><a class="dropdown-item" href="#" @click="copyToClipboard(diagram.diagram_code)">Copy diagram code</a></li>
              <li><a class="dropdown-item" href="#" @click="copyToClipboard(diagram.user_prompt)">Copy user prompt</a></li>
              <li><a class="dropdown-item" href="#" @click="downloadImage(`http://localhost:8090/api/files/srauto_diagrams/${diagram.id}/${diagram.diagram_png}`, diagram.diagram_type)">Download Diagram Image</a></li>
          </ul>
          </div>
           <img 
              :src="`http://localhost:8090/api/files/srauto_diagrams/${diagram.id}/${diagram.diagram_png}`"
              alt="Sequence Diagram"
              style="width: 80%; margin-top: 1rem; border-radius: 4px;"
            />
          

          </div>
         
         
         
         
         
    
       </div>
    </div>
  </div>
  </div>
  </section>
  </main>
</template>


<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Google+Sans:ital,opsz,wght@0,17..18,400..700;1,17..18,400..700&family=Hanken+Grotesk:ital,wght@0,100..900;1,100..900&family=Lato:ital,wght@0,100;0,300;0,400;0,700;0,900;1,100;1,300;1,400;1,700;1,900&family=League+Spartan:wght@100..900&family=Playfair+Display:ital,wght@0,400..900;1,400..900&display=swap');


.title {
  font-family: 'Google Sans', sans-serif;
}

.recentTitle {
  margin-top: 10rem;
  font-family: 'Google Sans', sans-serif;
  font-size: 1.5rem;
  color: #000000;
  margin-bottom: 1rem;
  text-align: start;
}

.no-diagram {
  font-family: 'Google Sans', sans-serif;
  font-size: 1.25rem;
  color: #64748b;
  font-style: italic;
  text-align: center;
  margin-top: 2rem;
}


.recent {
  max-width: 1200px;
  margin: 0 auto;
  padding: 4rem 2rem;
  font-family: 'Inter', sans-serif;
  color: #2c3e50;
  line-height: 1.6;
}

.diagram-type{
  text-align: center;
  padding: 10px ;
  background: linear-gradient(45deg, #42b883, #35495e);
  color: white;
  width: 20%;
  border-radius: 20px;
  margin-top: 10px;
}

.hero {
  text-align: center;
  margin-bottom: 4rem;
}

.hero h1 {
  font-size: 3rem;
  font-weight: 800;
  background: linear-gradient(45deg, #42b883, #35495e);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  margin-bottom: 1rem;
}

.btn-copy{
  font-size: 10px;
  padding: 10px;
}

.subtitle {
  font-size: 1.25rem;
  color: #64748b;
  font-family: 'Google Sans', sans-serif;
}

.content-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
  gap: 2rem;
  margin-bottom: 4rem;

}

.card {
  padding: 2rem;
  background: #f8fafc;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
  width: 100%;
}

.card h2 {
  color: #42b883;
  margin-bottom: 1rem;
}


.action-bar{
  margin-bottom: 1rem;
  
}

.result-section {
  margin-bottom: 4rem;
}

.result-card {
  width: 100%;
}

.result-box {
  min-height: 500px;
  max-height: 80vh;
  overflow: auto;
  padding: 1.5rem;
  background: #ffffff;
  border: 2px solid #e2e8f0;
  border-radius: 10px;
  display:block;
  position: relative;
}

.result-box .mermaid {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  min-height: 100%;
}

.result-box .mermaid svg {
  width: 100% !important;
  height: 100% !important;
  min-height: 460px;
}

textarea {
  width: 100%;
  min-height: 160px;
  padding: 1rem 1.2rem;
  border: 2px solid #e2e8f0;
  border-radius: 10px;
  background: #ffffff;
  color: #2c3e50;
  font-family: 'Google Sans', sans-serif;
  font-size: 0.95rem;
  line-height: 1.6;
  resize: vertical;
  outline: none;
  transition: border-color 0.3s ease, box-shadow 0.3s ease;
}

textarea::placeholder {
  color: #94a3b8;
  font-style: italic;
}

textarea:focus {
  border-color: #42b883;
  box-shadow: 0 0 0 4px rgba(66, 184, 131, 0.15);
}

button {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  margin-top: 1rem;
  padding: 0.75rem 2rem;
  background: linear-gradient(135deg, #42b883, #35495e);
  color: #ffffff;
  font-family: 'Google Sans', sans-serif;
  font-size: 0.95rem;
  font-weight: 700;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  letter-spacing: 0.02em;
  transition: transform 0.2s ease, box-shadow 0.2s ease, opacity 0.2s ease;
}

button:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(66, 184, 131, 0.35);
}

button:active {
  transform: translateY(0);
  opacity: 0.9;
}

.tech-stack {
  text-align: center;
  border-top: 1px solid #e2e8f0;
  padding-top: 3rem;
}

.chips {
  display: flex;
  justify-content: center;
  gap: 1rem;
  margin-top: 1rem;
}

.chips span {
  background: #35495e;
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 20px;
  font-size: 0.85rem;
  font-weight: 600;
}

ol {
  padding-left: 1.2rem;
}

ol li {
  margin-bottom: 0.5rem;
}
</style>