<script setup>
import { nextTick, onMounted, ref } from 'vue';
import mermaid from 'mermaid';
import { SequenceDiagramFactory } from '@/factory/diagramFactory';
import axios from 'axios';


// initialize mermaid theme and security level
mermaid.initialize({ 
  startOnLoad: false, 
  theme: 'default',
  securityLevel: 'loose',
});






// define expert from sequence diagram factory variable  
const expert = SequenceDiagramFactory.sequence;

const userPrompt = ref("");
const resultDiagram = ref("");
const loading = ref(false);
const apiData = ref(null);




// ================================================= test connection to pocketbase ======================================================== //




const testConnection = async() =>{

  try {

  const response = await fetch('http://localhost:8090/api/test' , {
    credentials: 'include' 
  });

  if (!response.ok) throw new Error("Connection to PocketBase failed.");
  if(response.status === 200) console.log("Connection to PocketBase successful!");
  
  const json = await response.json();
  apiData.value = json.pbid;

  console.log("Your PBID from PocketBase is:", apiData.value);

  

  

  } catch (error) {
    console.error('Error:', error);
    apiData.value = "Error: " + error.message;
  }
  
}

onMounted(() => {
  testConnection();
});

// ============================================== generate sequence diagram ======================================================== //


const generateDiagram = async () => {
  if (!userPrompt.value) return alert("Please enter a prompt");
  
  loading.value = true;
  resultDiagram.value = ""; // Reset

  try {
    // 1. Direct Fetch to Ollama
    const response = await fetch('http://localhost:11434/api/chat', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        model: "gemma3:1b",
        messages: [
          {
            role: expert.role,
            content: expert.systemPrompt,
            Rules: expert.fixer,
          },
          {
            role: "user",
            content: `Generate a sequence diagram for: ${userPrompt.value}`
          }
        ],
        stream: false
      })
    });

    if (!response.ok) throw new Error("Connection to Ollama failed.");

    const data = await response.json();
    let aiReply = data.message.content.trim();

    
    aiReply = aiReply.replace(/```mermaid/g, "").replace(/```/g, "");
    
    const startIdx = aiReply.indexOf(expert.header);
    if (startIdx === -1) throw new Error("Invalid diagram format.");


    const finalCleanCode = expert.fixer(aiReply.substring(startIdx));
    resultDiagram.value = finalCleanCode;

    await nextTick();
    const element = document.getElementById('mermaid-box');
    
    if (element) {
      element.removeAttribute('data-processed');

      element.innerHTML = finalCleanCode;
      
    
      await mermaid.run({ nodes: [element] });
    }

  } catch (error) {
    console.error('Error:', error);
    resultDiagram.value = "Error: " + error.message;
  } finally {
    loading.value = false;
  }
};

// ============================================== copy to clipboard ======================================================== //

const copyToClipboard = async () => {
   if (!resultDiagram.value) return;

   try {

    await navigator.clipboard.writeText(resultDiagram.value);
    alert("Diagram copied to clipboard");

   } catch (error) {

    console.error('Error:', error);
    alert("Failed to copy diagram to clipboard");

   }
}

// ============================================== copy image to clipboard ======================================================== //

const copyImageToClipboard = async () => {
  if (!resultDiagram.value) return;

  const svgElement = document.querySelector("#mermaid-box svg");
  if (!svgElement) return;

  try {
    const scaleFactor = 3;
    const padding = 80;

    // 1. Get bounding box of the actual diagram content
    const bBox = svgElement.getBBox();

    const contentWidth = bBox.width + padding;
    const contentHeight = bBox.height + padding;

    // 2. Clone SVG and set explicit viewBox + dimensions so the
    //    Image element renders at the exact size we expect
    const svgClone = svgElement.cloneNode(true);
    svgClone.setAttribute("viewBox", `${bBox.x - padding / 2} ${bBox.y - padding / 2} ${contentWidth} ${contentHeight}`);
    svgClone.setAttribute("width", contentWidth);
    svgClone.setAttribute("height", contentHeight);

    // 3. Serialize the corrected SVG
    const svgData = new XMLSerializer().serializeToString(svgClone);
    const svgBlob = new Blob([svgData], { type: "image/svg+xml;charset=utf-8" });
    const url = URL.createObjectURL(svgBlob);

    // 4. Draw onto a high-res canvas
    const canvas = document.createElement("canvas");
    const ctx = canvas.getContext("2d");
    canvas.width = contentWidth * scaleFactor;
    canvas.height = contentHeight * scaleFactor;
    ctx.scale(scaleFactor, scaleFactor);

   // ... inside copyImageToClipboard, after creating the 'img' constant

const img = new Image();
// 1. Set crossOrigin to Anonymous to prevent tainting
img.crossOrigin = "anonymous"; 

img.onload = () => {
  ctx.fillStyle = "white";
  ctx.fillRect(0, 0, canvas.width, canvas.height);
  
  // Use the scaled content dimensions
  ctx.drawImage(img, 0, 0, contentWidth * scaleFactor, contentHeight * scaleFactor);

  try {
    canvas.toBlob((blob) => {
      if (!blob) throw new Error("Blob creation failed");
      
      const item = new ClipboardItem({ "image/png": blob });
      navigator.clipboard.write([item]).then(() => {
        alert("Full HD Diagram copied!");
        URL.revokeObjectURL(url);
      });
    }, "image/png", 1.0);
  } catch (err) {
    console.error("Export failed:", err);
    alert("Canvas tainted. Try using a different browser or downloading as SVG.");
  }
};

// 2. Encode the SVG string properly using encodeURIComponent
// This is often more reliable than a raw Blob URL for canvas drawing
const encodedData = encodeURIComponent(svgData);
img.src = `data:image/svg+xml;charset=utf-8,${encodedData}`;
} catch (error) {
    console.error('Copy Error:', error);
    alert("Error copying image");
  }
}


</script>

<template>
  <main class="home">
    <section class="hero">
      <h1 class="title">SR Automate</h1>
      <p class="subtitle">Create your sequence diagram in minutes!</p>
       <p v-if="apiData" class="user-status">
      Session active: {{ apiData.substring(0, 100) }}...
     </p>
    </section>

    <section class="content-grid">
     
      <div class="card">
        <h2>Prompt</h2>
        <textarea id="prompt" v-model="userPrompt" placeholder="Enter your user story or functional requirement"></textarea>
        <button id="generate" @click="generateDiagram">Generate</button>
      </div>
    </section>

    <section class="result-section">
      <div class="card result-card">
        <h2 class="title">Result</h2>

        <div v-if="resultDiagram && !loading" class="action-bar">
          <button class="btn-copy" @click="copyToClipboard">
            Copy to Clipboard
          </button>
           <button class="btn-copy" @click="copyImageToClipboard">
            Copy Image to Clipboard
          </button> 
          
        </div>

        <div class="result-box">
          <p v-if="loading" class="subtitle">Generating your diagram...</p>
          
          <div id="mermaid-box" class="mermaid">
            {{ resultDiagram }}
          </div>
        </div>
      </div>
    </section>

    <section class="tech-stack">
      <h3>Powered By</h3>
      <div class="chips">
        <span>FastAPI</span>
        <span>Vue 3</span>
        <span>OpenAI</span>
        <span>Mermaid.js</span>
      </div>
    </section>
  </main>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Google+Sans:ital,opsz,wght@0,17..18,400..700;1,17..18,400..700&family=Hanken+Grotesk:ital,wght@0,100..900;1,100..900&family=Lato:ital,wght@0,100;0,300;0,400;0,700;0,900;1,100;1,300;1,400;1,700;1,900&family=League+Spartan:wght@100..900&family=Playfair+Display:ital,wght@0,400..900;1,400..900&display=swap');


.title {
  font-family: 'Google Sans', sans-serif;
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