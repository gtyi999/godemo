package mypkg

import "strconv"

type Home struct {

    Id  int

}

func (h *Home) AddPeople(Name string) string {
    return Name + strconv.Itoa(h.Id)
}
