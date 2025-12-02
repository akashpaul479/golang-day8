package day8

type Notifier interface {
	send(msg string) error
}

func Notifyuser(n Notifier, msg string) error {
	return n.send(msg)
}

type Fakenotifier struct {
	Received string
}

func (f *Fakenotifier) send(msg string) error {
	f.Received = msg
	return nil
}
