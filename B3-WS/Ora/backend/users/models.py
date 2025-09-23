from django.contrib.auth.models import AbstractUser
from django.db import models
import hashlib
import base58
import pgpy

class User(AbstractUser):
    public_key = models.TextField(blank=True, null=True)
    private_key = models.TextField(blank=True, null=True)
    ora_id = models.CharField(max_length=32, blank=True, null=True)

    def generate_ora_id(self):
        # SHA256 de la clé publique
        sha256_hash = hashlib.sha256(self.public_key.encode()).hexdigest()
        short_hash = sha256_hash[:16]

        # Encode en Base58
        base58_encoded = base58.b58encode(bytes.fromhex(short_hash)).decode()

        # Checksum 2 caractères
        checksum = hashlib.sha256(base58_encoded.encode()).hexdigest()[:2].upper()
        id_with_checksum = f"{base58_encoded}-{checksum}"

        # Format lisible
        clean = id_with_checksum.replace("-", "")
        chunks = [clean[i:i+4] for i in range(0, len(clean)-2, 4)]
        self.ora_id = "ora:" + "-".join(chunks) + "-" + clean[-2:]

    def generate_pgp_keys(self, name=None, email=None):
        key = pgpy.PGPKey.new(pgpy.constants.PubKeyAlgorithm.RSAEncryptOrSign, 2048)
        uid = pgpy.PGPUID.new(name or self.username, email=email or self.email)
        key.add_uid(uid, usage={pgpy.constants.KeyFlags.Sign, pgpy.constants.KeyFlags.EncryptCommunications},
                    hashes=[pgpy.constants.HashAlgorithm.SHA256],
                    ciphers=[pgpy.constants.SymmetricKeyAlgorithm.AES256],
                    compression=[pgpy.constants.CompressionAlgorithm.ZLIB])

        self.private_key = str(key)
        self.public_key = str(key.pubkey)
