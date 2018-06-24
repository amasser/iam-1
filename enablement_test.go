package iam_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/maurofran/iam"
)

var _ = Describe("Indefinite enablement", func() {
	var e Enablement

	BeforeEach(func() {
		e = IndefiniteEnablement()
	})

	Describe("#IsTimeExpired", func() {
		It("should return false", func() {
			Expect(e.IsTimeExpired()).To(BeFalse())
		})
	})
	Describe("#IsEnabled", func() {
		It("should return true", func() {
			Expect(e.IsEnabled()).To(BeTrue())
		})
	})
})

var _ = Describe("Past enablement", func() {
	var e Enablement

	BeforeEach(func() {
		e = IndefiniteEnablement()
		e.StartDate = time.Now().Add(-48 * time.Hour)
		e.EndDate = e.StartDate.Add(24 * time.Hour)
	})

	Describe("#IsTimeExpired", func() {
		It("should return true", func() {
			Expect(e.IsTimeExpired()).To(BeTrue())
		})
	})
	Describe("#IsEnabled", func() {
		It("should return false", func() {
			Expect(e.IsEnabled()).To(BeFalse())
		})
	})
})

var _ = Describe("Future enablement", func() {
	var e Enablement

	BeforeEach(func() {
		e = IndefiniteEnablement()
		e.StartDate = time.Now().Add(24 * time.Hour)
		e.EndDate = e.StartDate.Add(24 * time.Hour)
	})

	Describe("#IsTimeExpired", func() {
		It("should return true", func() {
			Expect(e.IsTimeExpired()).To(BeTrue())
		})
	})
	Describe("#IsEnabled", func() {
		It("should return false", func() {
			Expect(e.IsEnabled()).To(BeFalse())
		})
	})
})

var _ = Describe("Actual enablement", func() {
	var e Enablement

	BeforeEach(func() {
		e = IndefiniteEnablement()
		e.StartDate = time.Now().Add(-24 * time.Hour)
		e.EndDate = e.StartDate.Add(48 * time.Hour)
	})

	Describe("#IsTimeExpired", func() {
		It("should return true", func() {
			Expect(e.IsTimeExpired()).To(BeFalse())
		})
	})
	Describe("#IsEnabled", func() {
		It("should return false", func() {
			Expect(e.IsEnabled()).To(BeTrue())
		})
	})
})

var _ = Describe("Indefinite disabled enablement", func() {
	var e Enablement

	BeforeEach(func() {
		e = IndefiniteEnablement()
		e.Enabled = false
	})

	Describe("#IsTimeExpired", func() {
		It("should return false", func() {
			Expect(e.IsTimeExpired()).To(BeFalse())
		})
	})
	Describe("#IsEnabled", func() {
		It("should return false", func() {
			Expect(e.IsEnabled()).To(BeFalse())
		})
	})
})
