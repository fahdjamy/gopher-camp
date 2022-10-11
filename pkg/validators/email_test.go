package validators

import "testing"

func TestValidateEmail(t *testing.T) {
	isValid, _ := ValidateEmail("testing@gmail.com")
	if !isValid {
		t.Errorf("validation was incorrect: expected: %v got %v", true, isValid)
	}

	invalidEmail := "testing@k.com"
	notValid, _ := ValidateEmail(invalidEmail)

	if notValid {
		t.Errorf("validation was incorrect: expected: %v got %v", false, notValid)
	}
}

func TestValidateEmailDomain(t *testing.T) {
	correctDomain := "gmail.com"
	isValid, _ := validateMailDomain(correctDomain)
	if !isValid {
		t.Errorf("validation was incorrect: expected: %v got %v", true, isValid)
	}

	wrongDomain := "kk.com"
	invalid, _ := validateMailDomain(wrongDomain)
	if invalid {
		t.Errorf("validation was incorrect: expected: %v got %v", false, invalid)
	}
}
