import httpx
import fastapi
import uvicorn
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel

app = fastapi.FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

class DiagramRequest(BaseModel):
    prompt: str

@app.post("/generate")
async def generate_diagram(request: DiagramRequest):
    async with httpx.AsyncClient() as client:
        try:
            response = await client.post(
                "http://localhost:11434/api/chat",
                json={
                    "model": "gemma3:1b",
                    "messages": [
                        {
                            "role": "system",
                            "content": "You are a Mermaid expert. Return ONLY the code for a sequenceDiagram. Do not use markdown backticks."
                        },
                        {"role": "user", "content": request.prompt}
                    ],
                    "stream": False
                },
                timeout=30.0
            )
            response_json = response.json()
            ai_reply = response_json['message']['content'].strip()
            
            # Remove any markdown fluff
            clean_code = ai_reply.replace("```mermaid", "").replace("```", "").strip()
            
            return {"mermaid_code": clean_code}
        except Exception as e:
            print(f"Backend Error: {e}")
            return {"error": str(e), "mermaid_code": "sequenceDiagram\nAI->>User: Error generating diagram"}

if __name__ == "__main__":
    uvicorn.run(app, host="127.0.0.1", port=8080)