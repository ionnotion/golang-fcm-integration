package firebase

import "firebase.google.com/go/messaging"

func (f FCM) GetClient(i int) *messaging.Client {
	return f.Clients[i]
}
