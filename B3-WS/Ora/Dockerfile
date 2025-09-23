# Dockerfile à la racine de Ora/
FROM python:3.11-slim

WORKDIR /app

# Copier requirements et installer les dépendances
COPY backend/requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt daphne channels channels_redis

# Copier le backend complet
COPY backend/ /app/

# Exposer le port
EXPOSE 8000

# Commande par défaut (peut être surchargée par docker-compose)
CMD ["daphne", "-b", "0.0.0.0", "-p", "8000", "project.asgi:application"]
