package domain

import (
	"testing"

	"github.com/desafio-senhaForte-GraphQL/graph/model"
)

// TestRegrasValidaSenha - teste de cada regra da função ValidaSenha
func TestRegrasValidaSenha(t *testing.T) {

	t.Run("teste Falha minSize", func(t *testing.T) {
		var senha model.SenhaInput
		var rule model.Rule

		senha.Password = "teste"
		rulesEsperado := []string{"minSize"}
		rule.Rule = "minSize"
		rule.Value = 6

		senha.Rules = append(senha.Rules, &rule)

		verify, rules, _ := ValidaSenha(senha)

		if *verify {
			t.Errorf("esperado: %v, resultado: %v", false, verify)
		}
		if len(rules) == 0 {
			t.Errorf("esperado: %v, resultado: %v", rulesEsperado, len(rules))
		}
	})

	t.Run("teste OK minSize", func(t *testing.T) {
		var senha model.SenhaInput
		var rule model.Rule

		senha.Password = "testes"
		rulesEsperado := []string{}
		rule.Rule = "minSize"
		rule.Value = 6

		senha.Rules = append(senha.Rules, &rule)

		verify, rules, _ := ValidaSenha(senha)

		if !*verify {
			t.Errorf("esperado: %v, resultado: %v", true, verify)
		}
		if len(rules) > 0 {
			t.Errorf("esperado: %v, resultado: %v", rulesEsperado, len(rules))
		}
	})

	t.Run("teste Falha minUppercase", func(t *testing.T) {
		var senha model.SenhaInput
		var rule model.Rule

		senha.Password = "testE"
		rulesEsperado := []string{"minUppercase"}
		rule.Rule = "minUppercase"
		rule.Value = 2

		senha.Rules = append(senha.Rules, &rule)

		verify, rules, _ := ValidaSenha(senha)

		if *verify {
			t.Errorf("esperado: %v, resultado: %v", false, verify)
		}
		if len(rules) == 0 {
			t.Errorf("esperado: %v, resultado: %v", rulesEsperado, len(rules))
		}
	})

	t.Run("teste OK minUppercase", func(t *testing.T) {
		var senha model.SenhaInput
		var rule model.Rule

		senha.Password = "tEstEs"
		rulesEsperado := []string{}
		rule.Rule = "minUppercase"
		rule.Value = 2

		senha.Rules = append(senha.Rules, &rule)

		verify, rules, _ := ValidaSenha(senha)

		if !*verify {
			t.Errorf("esperado: %v, resultado: %v", true, verify)
		}
		if len(rules) > 0 {
			t.Errorf("esperado: %v, resultado: %v", rulesEsperado, len(rules))
		}
	})

	t.Run("teste Falha minLowercase", func(t *testing.T) {
		var senha model.SenhaInput
		var rule model.Rule

		senha.Password = "TEsTE"
		rulesEsperado := []string{"minLowercase"}
		rule.Rule = "minLowercase"
		rule.Value = 2

		senha.Rules = append(senha.Rules, &rule)

		verify, rules, _ := ValidaSenha(senha)

		if *verify {
			t.Errorf("esperado: %v, resultado: %v", false, verify)
		}
		if len(rules) == 0 {
			t.Errorf("esperado: %v, resultado: %v", rulesEsperado, len(rules))
		}
	})

	t.Run("teste OK minLowercase", func(t *testing.T) {
		var senha model.SenhaInput
		var rule model.Rule

		senha.Password = "tEsTE"
		rulesEsperado := []string{}
		rule.Rule = "minLowercase"
		rule.Value = 2

		senha.Rules = append(senha.Rules, &rule)

		verify, rules, _ := ValidaSenha(senha)

		if !*verify {
			t.Errorf("esperado: %v, resultado: %v", true, verify)
		}
		if len(rules) > 0 {
			t.Errorf("esperado: %v, resultado: %v", rulesEsperado, len(rules))
		}
	})

	t.Run("teste Falha minDigit", func(t *testing.T) {
		var senha model.SenhaInput
		var rule model.Rule

		senha.Password = "TEsTE1"
		rulesEsperado := []string{"minDigit"}
		rule.Rule = "minDigit"
		rule.Value = 2

		senha.Rules = append(senha.Rules, &rule)

		verify, rules, _ := ValidaSenha(senha)

		if *verify {
			t.Errorf("esperado: %v, resultado: %v", false, verify)
		}
		if len(rules) == 0 {
			t.Errorf("esperado: %v, resultado: %v", rulesEsperado, len(rules))
		}
	})

	t.Run("teste OK minDigit", func(t *testing.T) {
		var senha model.SenhaInput
		var rule model.Rule

		senha.Password = "tE1sT2E"
		rulesEsperado := []string{}
		rule.Rule = "minDigit"
		rule.Value = 2

		senha.Rules = append(senha.Rules, &rule)

		verify, rules, _ := ValidaSenha(senha)

		if !*verify {
			t.Errorf("esperado: %v, resultado: %v", true, *verify)
		}
		if len(rules) > 0 {
			t.Errorf("esperado: %v, resultado: %v", rulesEsperado, len(rules))
		}
	})

	t.Run("teste Falha minSpecialChars", func(t *testing.T) {
		var senha model.SenhaInput
		var rule model.Rule

		senha.Password = "TEsTE1@"
		rulesEsperado := []string{"minSpecialChars"}
		rule.Rule = "minSpecialChars"
		rule.Value = 2

		senha.Rules = append(senha.Rules, &rule)

		verify, rules, _ := ValidaSenha(senha)

		if *verify {
			t.Errorf("esperado: %v, resultado: %v", false, verify)
		}
		if len(rules) == 0 {
			t.Errorf("esperado: %v, resultado: %v", rulesEsperado, len(rules))
		}
	})

	t.Run("teste OK minSpecialChars", func(t *testing.T) {
		var senha model.SenhaInput
		var rule model.Rule

		senha.Password = "tE1#sTE!"
		rulesEsperado := []string{}
		rule.Rule = "minSpecialChars"
		rule.Value = 2

		senha.Rules = append(senha.Rules, &rule)

		verify, rules, _ := ValidaSenha(senha)

		if !*verify {
			t.Errorf("esperado: %v, resultado: %v", true, verify)
		}
		if len(rules) > 0 {
			t.Errorf("esperado: %v, resultado: %v", rulesEsperado, len(rules))
		}
	})

	t.Run("teste Falha noRepeted Minuscula", func(t *testing.T) {
		var senha model.SenhaInput
		var rule model.Rule

		senha.Password = "TEssTE1@"
		rulesEsperado := []string{"noRepeted"}
		rule.Rule = "noRepeted"

		senha.Rules = append(senha.Rules, &rule)

		verify, rules, _ := ValidaSenha(senha)

		if *verify {
			t.Errorf("esperado: %v, resultado: %v", false, verify)
		}
		if len(rules) == 0 {
			t.Errorf("esperado: %v, resultado: %v", rulesEsperado, len(rules))
		}
	})

	t.Run("teste Falha noRepeted letra Maiuscula", func(t *testing.T) {
		var senha model.SenhaInput
		var rule model.Rule

		senha.Password = "TETTE1@"
		rulesEsperado := []string{"noRepeted"}
		rule.Rule = "noRepeted"

		senha.Rules = append(senha.Rules, &rule)

		verify, rules, _ := ValidaSenha(senha)

		if *verify {
			t.Errorf("esperado: %v, resultado: %v", false, verify)
		}
		if len(rules) == 0 {
			t.Errorf("esperado: %v, resultado: %v", rulesEsperado, len(rules))
		}
	})

	t.Run("teste OK noRepeted", func(t *testing.T) {
		var senha model.SenhaInput
		var rule model.Rule

		senha.Password = "tE1#sTE!"
		rulesEsperado := []string{}
		rule.Rule = "noRepeted"

		senha.Rules = append(senha.Rules, &rule)

		verify, rules, _ := ValidaSenha(senha)

		if !*verify {
			t.Errorf("esperado: %v, resultado: %v", true, verify)
		}
		if len(rules) > 0 {
			t.Errorf("esperado: %v, resultado: %v", rulesEsperado, len(rules))
		}
	})

}

// TestGeralValidaSenha - teste completo da função
func TestGeralValidaSenha(t *testing.T) {
	t.Run("teste Completo Fail", func(t *testing.T) {
		var senha model.SenhaInput

		senha.Password = "tEs21#T!"
		noMatch := []*string{}
		rulesEsperado := []string{"minLowercase", "minDigit", "minSpecialChars"}
		rules := []string{"noRepeted", "minSize", "minUppercase", "minLowercase", "minDigit", "minSpecialChars"}
		for i, v := range rules {
			rule := model.Rule{
				Rule:  v,
				Value: i,
			}
			senha.Rules = append(senha.Rules, &rule)
		}

		verify, noMatch, _ := ValidaSenha(senha)

		if *verify {
			t.Errorf("esperado: %v, resultado: %v", false, *verify)
		}
		if len(noMatch) == 0 {
			t.Errorf("esperado: %v, resultado: %v", rulesEsperado, noMatch)
		}
	})

	t.Run("teste Completo OK", func(t *testing.T) {
		var senha model.SenhaInput

		senha.Password = "tEsr#T!@$%&1487"
		rulesEsperado := []string{""}
		rules := []string{"noRepeted", "minSize", "minUppercase", "minLowercase", "minDigit", "minSpecialChars"}
		for i, v := range rules {
			rule := model.Rule{
				Rule:  v,
				Value: i,
			}
			senha.Rules = append(senha.Rules, &rule)
		}

		verify, noMatch, _ := ValidaSenha(senha)

		if !*verify {
			t.Errorf("esperado: %v, resultado: %v", true, *verify)
		}
		if len(noMatch) > 0 {
			t.Errorf("esperado: %v, resultado: %v", rulesEsperado, noMatch)
		}
	})
}
