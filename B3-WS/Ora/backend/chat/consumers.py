from channels.generic.websocket import AsyncJsonWebsocketConsumer

class WebRTCConsumer(AsyncJsonWebsocketConsumer):
    async def connect(self):
        self.room_name = self.scope["url_route"]["kwargs"]["room"]
        self.group_name = f"webrtc_{self.room_name}"
        await self.channel_layer.group_add(self.group_name, self.channel_name)
        await self.accept()

    async def disconnect(self, close_code):
        await self.channel_layer.group_discard(self.group_name, self.channel_name)

    async def receive_json(self, content):
        await self.channel_layer.group_send(
            self.group_name,
            {
                "type": "webrtc.signal",
                "message": content,
                "sender": self.channel_name,
            },
        )

    async def webrtc_signal(self, event):
        # Ne pas renvoyer à soi-même
        if self.channel_name != event["sender"]:
            await self.send_json(event["message"])
