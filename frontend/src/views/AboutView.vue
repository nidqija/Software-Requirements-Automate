<script setup>
// You can import assets or components here
</script>

<template>
  <main class="about-container">
    <section class="hero">
      <h1>About SR Automate</h1>
      <p class="subtitle">Bridging the gap between human thought and technical architecture.</p>
    </section>

    <section class="content-grid">
      <div class="card">
        <h2>The Mission</h2>
        <p>
          Creating sequence diagrams manually is time-consuming and prone to human error. 
          <strong>SR Automate</strong> uses advanced AI to interpret natural language 
          requirements and instantly generate precise, renderable technical diagrams.
        </p>
      </div>

      <div class="card">
        <h2>How It Works</h2>
        <ol>
          <li><strong>Input:</strong> Paste your user story or functional requirement.</li>
          <li><strong>Analysis:</strong> Our FastAPI backend processes the logic using LLMs.</li>
          <li><strong>Visualization:</strong> The system generates structured code to render your diagram.</li>
        </ol>
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
.about-container {
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