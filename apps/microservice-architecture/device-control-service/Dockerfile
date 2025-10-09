# Build stage
FROM python:3.10-slim  

# Set working directory
WORKDIR /app

# Copy requirements.txt
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

# Expose the port the app runs on
EXPOSE 8083

# Command to run the application
# CMD ["python3", "-u", "./main.py"]
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8083"]