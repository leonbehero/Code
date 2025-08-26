# Task Master: Visual HTTP API Debugging Tool (React + Go Backend)

## Persona
**Role**: Full-stack development expert  
**Skills**:  
- Frontend: React, Tailwind, TypeScript  
- Visual drag-and-drop editors: React Flow  
- Backend: Go (Gin framework preferred)  
**Mindset**: Build a modular, extensible, developer-friendly tool, similar to n8n + Postman hybrid.

---

## Vision
Create a **visual HTTP API debugging tool** with a **drag-and-drop workflow editor**.  
- Simplify API testing and debugging.  
- Modular design for extensibility.  
- Clean UI inspired by n8n and Postman.  

---

## Objectives
1. Provide a **workflow editor** where users can drag and connect nodes.  
2. Support **HTTP request debugging** with configurable requests (URL, method, headers, body).  
3. Enable **execution flow** across nodes with real-time feedback.  
4. Allow **workflow persistence** (import/export JSON).  
5. Offer a **logging system** to track executions.  
6. Use **Go backend** for API proxy and workflow execution.  

---

## Tasks

### Task 1: Frontend (React + TypeScript + React Flow)
- Build a drag-and-drop workflow canvas.  
- Node types: HTTP Request, JSON Parser, Condition, Data Transformer, Log Output.  
- Node configuration side panel.  
- Execution log panel at bottom.  

### Task 2: Backend (Go + Gin)
- Implement REST API endpoints:  
  - `POST /execute` → Execute single node  
  - `POST /workflow/run` → Execute workflow  
  - `GET /workflow/:id` → Fetch workflow config  
  - `POST /workflow/save` → Save workflow  
- Proxy HTTP requests with support for headers and body.  
- Return full response (status, headers, body).  
- Provide logging API.  

### Task 3: Workflow Persistence
- Enable saving workflows as JSON.  
- Enable importing/exporting workflows.  

### Task 4: Logging & Debugging
- Record logs for each execution.  
- Display logs in frontend panel.  

### Task 5: Extended Features (Future)
- Plugin system: add new node types (WebSocket, DB Query, Delay Timer).  
- Variables/Context passing between nodes.  
- Multi-tab workflows.  

---

## Deliverables
1. **Frontend Code**: React + TypeScript + Tailwind + React Flow  
2. **Backend Code**: Go + Gin REST API  
3. **README**: Setup instructions (run Go server + React client)  
4. **Example Workflow**:  
   - HTTP Request → JSON Parser → Log Output  

---

## Suggested Project Name
- **FlowAPI**  
- **ReqFlow**  
- **DebugFlow**  
