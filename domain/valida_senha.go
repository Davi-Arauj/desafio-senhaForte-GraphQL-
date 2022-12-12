package domain

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/desafio-senhaForte-GraphQL/graph/model"
)

func ValidaSenha(senha model.SenhaInput) (*bool, []*string, error) {

	rules := []*string{}
	verify := true
	for _, rule := range senha.Rules {
		switch rule.Rule {
		case "minSize":
			valorString := strconv.Itoa(rule.Value)
			match := regexp.MustCompile(`(.){` + valorString + `}`)
			if !match.MatchString(senha.Password) {
				rules = append(rules, &rule.Rule)
			}
		case "minUppercase":
			valorString := strconv.Itoa(rule.Value)
			match := regexp.MustCompile(`(.*[A-Z]){` + valorString + `}`)
			if !match.MatchString(senha.Password) {
				rules = append(rules, &rule.Rule)
			}
		case "minLowercase":
			valorString := strconv.Itoa(rule.Value)
			match := regexp.MustCompile(`(.*[a-z]){` + valorString + `}`)
			if !match.MatchString(senha.Password) {
				rules = append(rules, &rule.Rule)
			}
		case "minDigit":
			valorString := strconv.Itoa(rule.Value)
			match := regexp.MustCompile(`(.*[0-9]){` + valorString + `}`)
			if !match.MatchString(senha.Password) {
				rules = append(rules, &rule.Rule)
			}
		case "minSpecialChars":
			valorString := strconv.Itoa(rule.Value)
			match := regexp.MustCompile(`(.*[!@#$%^&*()-+{}/]){` + valorString + `}`)
			if !match.MatchString(senha.Password) {
				rules = append(rules, &rule.Rule)
			}
		case "noRepeted":
			splitSenha := strings.Split(senha.Password, "")
			for _, letra := range splitSenha {
				match := regexp.MustCompile(`[` + letra + `]{2}`)
				if match.MatchString(senha.Password) {
					rules = append(rules, &rule.Rule)
					break
				}
			}
		}

	}

	if len(rules) > 0 {
		verify = false
		return &verify, rules, nil
	}
	return &verify, rules, nil
}
