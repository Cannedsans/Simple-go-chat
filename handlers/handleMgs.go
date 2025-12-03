package handlers


func HandleMessages(){
	for{
		msg := <-broadcast

		for client := range clients{
			if err := client.WriteJSON(&msg); err != nil{
				delete(clients, client)
				client.Close()
			}
		}
	}
}
