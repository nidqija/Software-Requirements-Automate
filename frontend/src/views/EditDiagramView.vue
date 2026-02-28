  <script setup>
  import { ref , onMounted, render , watch } from 'vue';
  import { useRoute } from 'vue-router';
  import axios from 'axios';
  import mermaid from 'mermaid';
  



  mermaid.initialize({ 
    startOnLoad: false, 
    theme: 'default',
    securityLevel: 'loose',
  });

  const route = useRoute();
  const diagramId = route.params.id; // Get the diagram ID from the route parameters
  const diagramDetails = ref({diagram_code: ''});
  const mermaidSvg = ref(''); 
  const isSaving = ref(false);


  
  const api = axios.create({
    baseURL: 'http://localhost:8090/api',
    withCredentials: true, 
  });



  const renderDiagram = async () => {
    if(!diagramDetails.value.diagram_code) return;

    try {
        const id = `mermaid-${Math.random().toString(36).substring(2, 9)}`;
        const { svg } = await mermaid.render(id, diagramDetails.value.diagram_code);
        mermaidSvg.value = svg;
    } catch (err) {
        console.error('Failed to render diagram:', err);
        alert('Failed to render diagram. Please check the diagram code for errors.');
    }
  }


  let debounceTimer;
  watch(() => diagramDetails.value.diagram_code, () => {
    clearTimeout(debounceTimer);
    debounceTimer = setTimeout(() => {
      renderDiagram();
    }, 500); 

  })


  const fetchDiagramDetails = async (id) => {
      try {
          const response = await api.get(`/fetch-diagram-by-id`, {
              params: { diagramId: id }
          });

          if (response.data && response.data.diagram){
              diagramDetails.value = response.data.diagram;
              renderDiagram();
          } else {
              alert('Diagram details not found.');
          }
          

      } catch (err) {
          console.error('Failed to fetch diagram details:', err);
          alert('Failed to load diagram details. Please try again.');
      }
  }

  const submitChanges = async () => {
      isSaving.value = true;
      try {
        const response = await api.post('/update-diagram', {
              diagramId: diagramId,
              diagramCode: diagramDetails.value.diagram_code,
          });

          if (response.data && response.data.status === "success") {
            alert('Diagram updated successfully!');
          } else {
            alert('Failed to update diagram.');
          }


      } catch (err) {
          console.error('Failed to update diagram:', err);
          alert('Failed to update diagram. Please try again.');
      } finally {
          isSaving.value = false;
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
          <p class="mb-5">Diagram Code:</p>
          <div class="row">
    <div class="col-sm-6">
      <form action="post" method="post">
      <div class="card">
        <div class="mb-3">
            <textarea class="form-control" style="height: 400px;" id="exampleFormControlTextarea1" rows="3" v-model="diagramDetails.diagram_code"></textarea>
          </div>
          <button 
  class="btn btn-primary" 
  type="button" 
  @click="submitChanges" 
  :disabled="isSaving"
>

  {{ isSaving ? 'Saving...' : 'Submit Changes' }}
</button>
      
      </div> 
        </form>
    </div>
    <div class="col-sm-6">
      <div class="card">
      <div class="card-body">

        <div id="mermaid-box" v-if="mermaidSvg" v-html="mermaidSvg" class="img-fluid">
        </div>
      
  <img 
      id="mermaid-box"
      v-else-if="diagramDetails.diagram_png" 
      :src="`http://127.0.0.1:8090/api/files/pbc_1294652289/${diagramDetails.id}/${diagramDetails.diagram_png}`" 
      class="img-fluid" 
      width="300px"
      alt="Diagram Preview"
    >

      </div>
          
      </div>
    </div>
  </div>
        
        

      </div>
      </div>
  </template>



  <style scoped>
  @import url('https://fonts.googleapis.com/css2?family=Google+Sans:ital,opsz,wght@0,17..18,400..700;1,17..18,400..700&family=Hanken+Grotesk:ital,wght@0,100..900;1,100..900&family=Lato:ital,wght@0,100;0,300;0,400;0,700;0,900;1,100;1,300;1,400;1,700;1,900&family=League+Spartan:wght@100..900&family=Playfair+Display:ital,wght@0,400..900;1,400..900&display=swap');


  .edit-diagram-view {
    padding: 20px;
    color: black;
  }



  .title {
    font-family: 'Google Sans', sans-serif;
  }

  .btnj {
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


  .home {
    max-width: 1200px;
    margin: 0 auto;
    padding: 4rem 2rem;
    font-family: 'Inter', sans-serif;
    color: #2c3e50;
    line-height: 1.6;
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
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 2rem;
    margin-bottom: 4rem;
  }

  .card {
    padding: 2rem;
    background: #f8fafc;
    border-radius: 12px;
    border: 1px solid #e2e8f0;
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

