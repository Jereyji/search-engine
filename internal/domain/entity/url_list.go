package entity

type URLList struct {
	ID        int
	Link      string
	Is_parsed bool
}

func (u *URLList) ChangeLink(link string) {
	u.Link = link
}

func (u *URLList) ChangeParseStatus(is_parsed bool) {
	u.Is_parsed = is_parsed
}