const baseCleaner = (code) => {
    // 1. Remove Markdown code fences immediately
    let clean = code
        .replace(/```mermaid/g, "")
        .replace(/```/g, "")
        .trim();

    let lines = clean.split('\n');
    let result = [];

    for (let line of lines) {
        let trimmed = line.trim();

        // 2. Skip truly empty lines
        if (!trimmed) continue;

        // 3. Skip lines that are CLEARLY just numbered lists from the AI's prose
        // We check for "1. " at the start, but ONLY if the line doesn't contain Mermaid symbols
        if (/^\d+\.\s/.test(trimmed) && !trimmed.includes("->") && !trimmed.includes(":")) {
            continue;
        }

        // 4. LESS STRICT: If the line isn't a known "garbage" pattern, keep it.
        // This ensures things like 'loop', 'opt', or complex labels don't get deleted.
        // We only exclude lines that look like conversational English (no symbols/keywords).
        const isMermaidConstruct = 
            /[-=>]/.test(trimmed) ||                // Contains arrows or relationships
            /[:{}]/.test(trimmed) ||                // Contains labels or blocks
            /^(sequenceDiagram|classDiagram|stateDiagram|erDiagram|gantt|pie|flowchart|graph)/i.test(trimmed) ||
            /^(participant|actor|note|alt|else|end|loop|opt|rect|critical|break|activate|deactivate)/i.test(trimmed) ||
            /^(class|style|callback|click|link|subgraph)/i.test(trimmed);

        if (isMermaidConstruct) {
            result.push(line);
        }
    }

    return result.join('\n');
};


const SequenceDiagramFactory = {
    sequence: {
        role: 'system',
        header: 'sequenceDiagram',
        systemPrompt: `
        You are a Mermaid generator for sequence diagrams. 
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
        fixer: (raw) => {
            let code = baseCleaner(raw);
            return code
                .replace(/Systemelse/g, 'System\nelse')
                .replace(/endelse/g, 'end\nelse')
                .replace(/deactivateSystem/g, 'deactivate System')
                .replace(/activateSystem/g, 'activate System')
                .replace(/^\s*\d+[\.\)]\s*/gm, '');
        }


    },
}

const ClassDiagramFactory = {
    class: {
        role: 'system',
        header: 'classDiagram',
        systemPrompt: `
        You are a Mermaid generator for class diagrams. 
           Rules:
           1. Start with 'classDiagram'.
           2. EVERY command must be on a NEW LINE.
           3. Use 'alt', 'else', and 'end' for logic, but put them on their OWN lines.
           4. NO backticks. NO explanations.
           5. Relationships: <|-- (Inheritance), *-- (Composition), o-- (Aggregation).
           6. Use visibility markers: + (public), - (private).
           7. Define classes using 'class ClassName { ... }'.

           Example:
           classDiagram
           class User {
             +String name
             +String email
             +login()
           }

           class Order {
             +String orderId
             +Date date
             +calculateTotal()
           }

           User "1" *-- "many" Order : places
        ` ,
        fixer: (code) => {
            return code
                .replace(/Systemelse/g, 'System\nelse')
                .replace(/endelse/g, 'end\nelse')
                .replace(/deactivateSystem/g, 'deactivate System')
                .replace(/activateSystem/g, 'activate System')
                .replace(/^\s*\d+[\.\)]\s*/gm, '');
        }


    },
}



const FlowchartFactory = {
    flowchart: {
        role: 'system',
        header: 'flowChart',
        systemPrompt: `
        You are a Mermaid generator for flowcharts. 
            Rules:
            1. Start with 'flowchart TD'.
            2. EVERY command must be on a NEW LINE.
            3. Use Nodes with Shapes:
               - [Rectangle] for Actions
               - {Diamond} for Decisions
               - ([Pill]) for Start/End
            4. Connections use '-->' or '-- text -->'.
            5. NO backticks. NO explanations.

            Example:
            flowchart TD
            A([Start]) --> B{Is Logged In?}
            B -- Yes --> C[Go to Dashboard]
            B -- No --> D[Go to Login Page]
            D --> E([End])
        ` ,
        fixer: (raw) => {
    let code = baseCleaner(raw);
    return code
        // This regex finds 'flowchart TD' (or LR) and ensures there's a newline after it
        .replace(/(flowchart\s+[A-Z]{2})([^\n])/i, '$1\n$2') 
        .replace(/^\s*\d+[\.\)]\s*/gm, '');
}


    },
}





export { SequenceDiagramFactory, ClassDiagramFactory , FlowchartFactory };