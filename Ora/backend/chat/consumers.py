from channels.generic.websocket import AsyncJsonWebsocketConsumer

class ChatConsumer(AsyncJsonWebsocketConsumer):
    async def connect(self):
        await self.accept()

    async def receive_json(self, content):
        # Echo simple pour commencer
        await self.send_json({"message": f"Reçu: {content}"})

    async def disconnect(self, close_code):
        pass
