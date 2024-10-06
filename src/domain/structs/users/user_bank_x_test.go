package users

import (
	"encoding/json"
	"testing"
	"time"
)

func TestMandatoryFields(t *testing.T) {
	mandatoryFields := [5]string{
		"name",
		"email",
		"password",
		"timestamps",
		"account",
	}
	user := &UserBankX{
		Name:     "Test",
		Email:    "test@email.com",
		Password: "hash",
		Timestamps: Timestamps{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Account: []BankAccount{
			{
				Account: "00 00 000",
				Balance: 235,
				Timestamps: Timestamps{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
		},
	}

	var result map[string]interface{}
	obj, _ := json.Marshal(user)
	json.Unmarshal([]byte(obj), &result)
	for _, mandatoryField := range mandatoryFields {
		if result[mandatoryField] == "" {
			t.Errorf("%s is mandatory field", mandatoryField)
		}
	}
}
