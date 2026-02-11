<script setup>
import { nextTick, ref } from 'vue';
import mermaid from 'mermaid';


mermaid.initialize({ 
  startOnLoad: false, 
  theme: 'default',
  securityLevel: 'loose',
});

const userPrompt = ref("");
const resultDiagram = ref("");
const loading = ref(false);

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
            role: "system",
            content: `You are a Mermaid generator. 
           Rules:
           1. Start with 'sequenceDiagram'.
           2. EVERY command must be on a NEW LINE.
           3. Use 'alt', 'else', and 'end' for logic, but put them on their OWN lines.
           4. NO backticks. NO explanations.

           Example:
           sequenceDiagram
           User->>System: Login
           alt success
           System->>User: OK
           else failure
           System->>User: Error
           end`
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
    
    const startIndex = aiReply.indexOf("sequenceDiagram");
    if (startIndex === -1) {
      throw new Error("AI failed to provide a valid sequenceDiagram start.");
    }


    let cleanCode = aiReply.substring(startIndex).trim();
    let fixedCode = cleanCode
  .replace(/Systemelse/g, 'System\nelse')
  .replace(/endelse/g, 'end\nelse')
  .replace(/deactivateSystem/g, 'deactivate System')
  .replace(/activateSystem/g, 'activate System');

  resultDiagram.value = fixedCode;


  
    await nextTick();
    const element = document.getElementById('mermaid-box');
    
    if (element) {
      element.removeAttribute('data-processed');

      element.innerHTML = fixedCode;
      
    
      await mermaid.run({ nodes: [element] });
    }

  } catch (error) {
    console.error('Error:', error);
    resultDiagram.value = "Error: " + error.message;
  } finally {
    loading.value = false;
  }
};





</script>

<template>
  <main class="home">
    <section class="hero">
      <h1 class="title">SR Automate</h1>
      <p class="subtitle">Human Prompt -> Sequence Diagram. In minutes</p>
    </section>

    <section class="content-grid">
      <div class="card">
        <h2>Prompt</h2>
        <textarea id="prompt" v-model="userPrompt" placeholder="Enter your user story or functional requirement"></textarea>
        <button id="generate" @click="generateDiagram">Generate</button>
      </div>

      <div class="card">
  <h2 class="title">Result</h2>
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
  max-width: 900px;
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