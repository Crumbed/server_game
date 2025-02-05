package client

import (
	"fmt"
	"main/packet"
)



func SCJoinResponseListener(context *packet.PacketContext, p packet.Packet) {
    client := context.Handler.(*Client)
    response := p.(*packet.JoinResponse)
    if response.IsOk() {
        fmt.Println("Successfully connected to game")
        client.Iam = PlayerN(response.PlayerN)
        return 
    }

    fmt.Println("Game is full, connection refused...")
    client.Conn.Close()
}

func SCGameStartListener(context *packet.PacketContext, p packet.Packet) {
    client := context.Handler.(*Client)
    client.Started = true
    client.SendPacket(p)
    fmt.Println("Starting game...")
}

func CCPaddleMoveListener(context *packet.PacketContext, data packet.Packet) {
    client := context.Handler.(*Client)
    update := data.(*packet.PaddleMove)

    client.Players[update.PlayerN].NewPos = update.Pos
}


