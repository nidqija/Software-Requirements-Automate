![alt text](image.png)

# SR Automate 🚀 
**Human Prompt → Sequence Diagram , Class Diagram , Flowchart**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![Vue Version](https://img.shields.io/badge/Vue.js-3.x-4FC08D?style=for-the-badge&logo=vuedotjs)](https://vuejs.org)
[![PocketBase](https://img.shields.io/badge/PocketBase-v0.22+-B8DBE4?style=for-the-badge&logo=pocketbase)](https://pocketbase.io)


SR Automate is an AI-driven tool designed to bridge the gap between abstract software requirements and technical documentation. By leveraging LLMs, it transforms natural language prompts into structured, renderable diagrams.

---

## 📌 Problem Statement
The transition from a business requirement to a technical design is often the most friction-heavy phase of the SDLC:

* **The "Blank Page" Bottleneck:** Engineers spend excessive time manually drafting initial diagrams, delaying architectural discussions.
* **Syntax Complexity:** Tools like Mermaid.js require learning specific syntax, creating a barrier for non-technical stakeholders.
* **Documentation Drift:** Updating visual diagrams is tedious, leading to outdated docs that misalign teams.

## 🎯 Our Mission
**SR Automate** acts as a **Technical Translator** to:
1.  **Accelerate Velocity:** Generate a 90% complete architectural starting point in seconds.
2.  **Democratize Design:** Enable Product Managers to generate technical visuals without mastering code.
3.  **Ensure Consistency:** Keep natural language and technical syntax in perfect sync.

---

## 🛠️ Tech Stack

| Component     | Technology |
| :------------ | :--- |
| **Frontend** | Vue 3 (Vite), Bootstrap 5.2, Axios |
| **Backend** | Go (PocketBase Framework) |
| **Database** | Embedded SQLite (via PocketBase) |
| **LLM** | Ollama (Gemma3) |
| **Rendering** | Mermaid.js |

---

## 🔄 Technical Workflow



1.  **Input:** User enters a natural language requirement (e.g., *"User signs up via Google"*).
2.  **Processing:** The Go backend sends the prompt to **Ollama** with specialized system instructions.
3.  **Generation:** The LLM returns valid **Mermaid.js** syntax.
4.  **Rendering:** The Vue.js frontend reactively renders the syntax into an interactive SVG diagram.

---

## 📂 Project Structure

```text
SR_Automate/
├── backend/             # Go (PocketBase) Server
│   ├── main.go          # API Routes & AI Controller
│   └── pb_data/         # Database & Migrations
└── frontend/            # Vue 3 Web App
    ├── src/
    │   ├── factory/     # Mermaid syntax templates
    │   ├── views/       # EditDiagramView.vue, Home.vue
    │   └── api/         # Axios configuration
    └── package.json

```

--- 

## 📚 Documentation
- **Software Requirements**: https://docs.google.com/document/d/1AJgbtDQgTli-4sOeMn2EfwC8PVFq6g-l/edit#heading=h.k2hqyx1pgi0a

---