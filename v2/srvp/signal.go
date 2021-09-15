package srvp

func Signal(signals <-chan struct{}) string {
	<-signals
	return "play"
}
