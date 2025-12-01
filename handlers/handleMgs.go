package handlers


func HandleMessages(){
	for{
		msg := <- broacast

		for clients := range client{
			if err := clients.WriteJSON(&msg); err != nil{
				delete(client, clients)
				clients.Close()
			}
		}
	}
}
