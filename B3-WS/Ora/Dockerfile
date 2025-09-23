FROM python:3.11-slim

# Installer les dépendances système utiles (ex : psycopg2, Pillow, etc.)
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    libpq-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copier requirements en premier (meilleur cache Docker)
COPY backend/requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

# Copier le backend complet
COPY backend/ /app/
