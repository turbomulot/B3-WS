FROM python:3.11-slim

# éviter les fichiers .pyc et activer le log direct
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

# dossier de travail dans le conteneur
WORKDIR /app

# installer dépendances système (si besoin pour pgpy)
RUN apt-get update && apt-get install -y \
    gcc libpq-dev \
    && rm -rf /var/lib/apt/lists/*

# copier et installer les dépendances python
COPY backend/requirements.txt /app/
RUN pip install --no-cache-dir -r requirements.txt

# copier le code backend
COPY backend/ /app/

# lancer le serveur avec Daphne (ASGI)
CMD ["daphne", "-b", "0.0.0.0", "-p", "8000", "project.asgi:application"]
