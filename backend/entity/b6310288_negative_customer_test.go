package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegativeCustomerID(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Check Negative CustomerID", func(t *testing.T) {

		c := Customer{
			Name:       "Jee",
			Email:      "jee@gmail.com",
			CustomerID: "M1234564455",
		}

		ok, err := govalidator.ValidateStruct(c)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("CustomerID ไม่ถูกต้อง"))
	})
}
