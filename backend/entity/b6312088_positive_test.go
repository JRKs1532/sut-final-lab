package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPositiveCustomer(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Positive Customer Check correct", func(t *testing.T) {
		c := Customer{
			Name:       "Jeerapat",
			Email:      "jee@gmail.com",
			CustomerID: "M1234567",
		}

		ok, err := govalidator.ValidateStruct(c)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
