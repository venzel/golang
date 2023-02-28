package send_mail

import "fmt"

/* */

type email struct {
	email string
}

/* */

type google struct {
	email
}

func (g google) send_mail(origin string) {
	fmt.Printf("Google. %s para: %s\n", origin, g.email.email)
}

/* */

type yahoo struct {
	email
}

func (y yahoo) send_mail(origin string) {
	fmt.Printf("Yahoo. %s para: %s\n", origin, y.email.email)
}

/* */

type model interface {
	send_mail(origin string)
}

func send(m model) {
	m.send_mail("fonte@gmail.com")
}

/* */

func Execute(print bool) {
	
	if print {

		e := email{"tiago@gmail.com"}
	
		g := google{e}

		send(g)

		y := yahoo{e}

		send(y)
	}
}