package gamer

type Contact struct {
    Emails []string
}

type Gamer struct {
    ID string
    Name string
    Winners int
    Contacts Contact
}

func CreateGamer(name string, winners int, contacts Contact) *Gamer {
    return &Gamer{
        ID: "1",
        Name: name,
        Winners: winners,
        Contacts: contacts,
    }
}