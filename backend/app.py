from fastapi import FastAPI
from pydantic import BaseModel
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI(title="AutoDev API")

# Allow Vue to call this API
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # later restrict
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# -------- Models --------

class GenerateRequest(BaseModel):
    sessionId: str
    prompt: str

class GenerateResponse(BaseModel):
    sessionId: str
    mermaid: str
    raw: dict

# -------- Fake in-memory DB --------

sessions = {}

# -------- Routes --------

@app.get("/")
def root():
    return {"message": "AutoDev backend running"}

@app.get("/health")
def health():
    return {"status": "ok"}

@app.post("/generate", response_model=GenerateResponse)
def generate(req: GenerateRequest):
    # Fake AI output for now
    fake_json = {
        "entities": [
            {
                "name": "User",
                "attributes": ["id", "name", "email"]
            },
            {
                "name": "Book",
                "attributes": ["id", "title", "author"]
            }
        ]
    }

    # Fake Mermaid
    mermaid = """
classDiagram
User {
  id
  name
  email
}
Book {
  id
  title
  author
}
User --> Book : borrows
"""

    # Store per session
    if req.sessionId not in sessions:
        sessions[req.sessionId] = []

    sessions[req.sessionId].append({
        "prompt": req.prompt,
        "json": fake_json,
        "mermaid": mermaid
    })

    return {
        "sessionId": req.sessionId,
        "mermaid": mermaid,
        "raw": fake_json
    }

@app.get("/session/{session_id}")
def get_session(session_id: str):
    return sessions.get(session_id, [])
