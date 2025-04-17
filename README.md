# 🧾 JD Generator - Automatic Job Description PDF Generator

JD Generator is a backend system that allows for the automatic generation of professional **Job Description (JD)** PDF files from input JSON data. The system uses Go for the API request handling and Python combined with LaTeX to render high-quality PDFs.

## 📌 Features

- Receives job description data via API
- Compiles data to LaTeX and generates PDF
- Returns the file or a download link to the PDF
- Supports custom templates
- Can handle asynchronous processing and send real-time status updates (WebSocket)

## 🏗️ Overall Architecture

Client
↓
Go Backend (REST API)
↓
Python Module (Render LaTeX → PDF)
↓
Storage (Local / S3)
↓
Client receives link or file


## Additional: Redis can be integrated for caching, queueing, or job status tracking.

## 🛠️ Technologies Used

| Component          | Technology           |
|-------------------|----------------------|
| API backend       | Go                   |
| PDF generator     | Python + LaTeX       |
| Storage           | Local / S3           |
| Caching / Queue   | Redis *(optional)* |
| Real-time status  | WebSocket *(optional)* |

## 🚀 Installation Guide

### 1. Clone repository

```bash
git clone [https://github.com/yourusername/jd-generator.git](https://github.com/yourusername/jd-generator.git)
cd jd-generator
```
2. Directory Structure
```
.
├── backend/          # Go backend
├── generator/        # Python LaTeX generator
├── templates/        # LaTeX templates
├── output/           # PDF output
└── README.md
```
3. Build backend (Go)
```
cd backend
go build -o jd-server main.go
```
4. Install Python dependencies
```

cd generator
python -m venv venv
source venv/bin/activate
pip install -r requirements.txt
```
Note: LaTeX (TexLive) needs to be installed on your system separately.

5. Run the application
```
./jd-server
```
Once the server is running, send a request to /generate with the following JSON payload:
JSON

```
{
  "position": "Backend Engineer",
  "location": "Hanoi",
  "responsibilities": [
    "Develop APIs using Go",
    "Integrate Redis and WebSocket"
  ],
  "requirements": [
    "2+ years of Go experience",
    "Understanding of distributed systems"
  ]
}
```
6. Example
After processing, you will receive:

File PDF: output/backend_engineer_2025.pdf

Or a link: http://localhost:8080/files/backend_engineer_2025.pdf

7. 📦 Future Expansions
- UI for uploading & previewing templates
- Dashboard for tracking job status
- Digital signature integration
- Batch JD generation system from Excel/CSV files
- OpenAI/LLM support for automatic description writing
8. 💬 Contact
If you have any questions or would like to contribute, please open an issue or contact us directly.