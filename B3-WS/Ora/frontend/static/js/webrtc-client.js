const room = "room1"
const ws = new WebSocket(`ws://${window.location.host}/ws/chat/${room}/`)

const pc = new RTCPeerConnection()
const dataChannel = pc.createDataChannel("chat")

dataChannel.onmessage = (e) => console.log("Message reçu:", e.data)

ws.onmessage = async (event) => {
  const data = JSON.parse(event.data)
  if (data.sdp) {
    await pc.setRemoteDescription(new RTCSessionDescription(data.sdp))
    if (data.sdp.type === "offer") {
      const answer = await pc.createAnswer()
      await pc.setLocalDescription(answer)
      ws.send(JSON.stringify({ "sdp": pc.localDescription }))
    }
  } else if (data.candidate) {
    try {
      await pc.addIceCandidate(new RTCIceCandidate(data.candidate))
    } catch (e) {
      console.error("Erreur ICE", e)
    }
  }
}

pc.onicecandidate = (event) => {
  if (event.candidate) {
    ws.send(JSON.stringify({ "candidate": event.candidate }))
  }
}

// Bouton pour initier un appel
async function startCall() {
  const offer = await pc.createOffer()
  await pc.setLocalDescription(offer)
  ws.send(JSON.stringify({ "sdp": pc.localDescription }))
}
