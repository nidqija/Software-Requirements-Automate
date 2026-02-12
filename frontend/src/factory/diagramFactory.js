
const baseCleaner = (code) => {
  return code
    .replace(/```mermaid/g, "")
    .replace(/```/g, "")
    // Remove numbered lists (like 1. 2. 3.) that break Mermaid
    .split('\n')
    .filter(line => !/^\d+\./.test(line.trim())) 
    .join('\n')
    .trim();
};


const SequenceDiagramFactory= {
    sequence :{
        role : 'system' ,
        header : 'sequenceDiagram' ,
        systemPrompt : `
        You are a Mermaid generator. 
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
           end
        ` ,
        fixer : (raw) => {
            let code = baseCleaner(raw);
            return code
            .replace(/Systemelse/g, 'System\nelse')
            .replace(/endelse/g, 'end\nelse')
            .replace(/deactivateSystem/g, 'deactivate System')
            .replace(/activateSystem/g, 'activate System');
        }


    } , 
}

const ClassDiagramFactory= {
    class :{
        role : 'system',
        header : 'classDiagram' ,
        systemPrompt : `
        You are a Mermaid generator. 
           Rules:
           1. Start with 'classDiagram'.
           2. EVERY command must be on a NEW LINE.
           3. Use 'alt', 'else', and 'end' for logic, but put them on their OWN lines.
           4. NO backticks. NO explanations.

           Example:
           classDiagram
           User->>System: Login
           alt success
           System->>User: OK
           else failure
           System->>User: Error
           end
        ` ,
        fixer : (code) => {
            return code
            .replace(/Systemelse/g, 'System\nelse')
            .replace(/endelse/g, 'end\nelse')
            .replace(/deactivateSystem/g, 'deactivate System')
            .replace(/activateSystem/g, 'activate System');
        }


    } , 
}

export { SequenceDiagramFactory , ClassDiagramFactory }