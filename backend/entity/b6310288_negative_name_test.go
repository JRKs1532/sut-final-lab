package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegativeNameCustomer(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Negative Name Customer Check blank", func(t *testing.T) {
		c := Customer{
			Name:       "",
			Email:      "jee@gmail.com",
			CustomerID: "M1234567",
		}

		ok, err := govalidator.ValidateStruct(c)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Name: non zero value required"))
	})
}
