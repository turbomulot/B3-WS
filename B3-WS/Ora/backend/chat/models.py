from django.db import models

class Peer(models.Model):
    user_id = models.CharField(max_length=200, unique=True)
    public_key = models.TextField()
    created_at = models.DateTimeField(auto_now_add=True)
