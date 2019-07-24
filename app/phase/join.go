package phase

func Join(rq valueProvider) Data { return Data{name: rq.Value()} }
