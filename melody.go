package main

// Melody is the websocket sender
// var Melody *melody.Melody

// func init(){
// 	r := chi.NewRouter()
// 	r.Get("/ws", upgradeWebsocket) // for Melody websockets
// 	http.ListenAndServe(":5000", r)
// }

// func initMelody() {
// 	Melody = melody.New()

// 	Melody.HandleConnect(func(s *melody.Session) {
// 		fmt.Println("CONNECTED")
// 		Melody.BroadcastFilter([]byte("HI"), func(q *melody.Session) bool {
// 			return q.Request.URL.Path == s.Request.URL.Path
// 		})
// 	})
// }

// func doSomething(){
// 	Melody.BroadcastFilter(buf[:n], func(q *melody.Session) bool {
// 		return true
// 	})
// }