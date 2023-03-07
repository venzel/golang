package main

type MapCheck map[string]bool;

func show(err error) {
    println(err.Error())
}

func (e *MapCheck) Error() string {
    list := make([]string, 0)

    for key, value := range *e {
        if value {
            list = append(list, key);
        }
    }

    text := ""

    if len(list) > 0 {
        text += "Erros:\n"

        for i := range list {
            text += list[i] + "\n";
        }
    } else {
        text = "Não possui!"
    }

    return text;
}

func main() {
    errors := MapCheck{
        "network": false,
        "database": false,
    }

    errors["network"] = true;
    errors["database"] = true;
    
    show(&errors)
}