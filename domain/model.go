package domain

type Senha struct {
	Password string `json:"password"`
	Rules    []Rule
}

type Rule struct {
	Rule  string `json:"rule"`
	Value int    `json:"value"`
}
