package resolver_test

import (
	"context"
	"net"
	"time"

	. "github.com/jmalloc/dissolve/src/resolver"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StandardResolver (net.Resolver parity)", func() {
	var (
		subject, builtin Resolver
		ctx              context.Context
		cancel           func()
	)

	BeforeEach(func() {
		subject = &StandardResolver{}
		builtin = &net.Resolver{}

		c, f := context.WithTimeout(context.Background(), 3*time.Second)
		ctx, cancel = c, f // assign in separate statement to silence "go vet" error
	})

	AfterEach(func() {
		cancel()
	})

	Describe("LookupAddr", func() {
		It("returns the same results as the built-in implementation", func() {
			s, err := subject.LookupAddr(ctx, "8.8.8.8")
			Expect(err).ShouldNot(HaveOccurred())

			r, err := builtin.LookupAddr(ctx, "8.8.8.8")
			Expect(err).ShouldNot(HaveOccurred())

			Expect(s).To(ConsistOf(r))
		})
	})

	Describe("LookupCNAME", func() {
		It("returns the same results as the built-in implementation", func() {
			s, err := subject.LookupCNAME(ctx, "mail.icecave.com.au")
			Expect(err).ShouldNot(HaveOccurred())

			r, err := builtin.LookupCNAME(ctx, "mail.icecave.com.au")
			Expect(err).ShouldNot(HaveOccurred())

			Expect(s).To(Equal(r))
		})
	})

	Describe("LookupHost", func() {
		It("returns the same results as the built-in implementation", func() {
			s, err := subject.LookupHost(ctx, "www.icecave.com.au")
			Expect(err).ShouldNot(HaveOccurred())

			r, err := builtin.LookupHost(ctx, "www.icecave.com.au")
			Expect(err).ShouldNot(HaveOccurred())

			Expect(s).To(ConsistOf(r))
		})
	})

	Describe("LookupIPAddr", func() {
		It("returns the same results as the built-in implementation", func() {
			s, err := subject.LookupIPAddr(ctx, "icecave.com.au")
			Expect(err).ShouldNot(HaveOccurred())

			r, err := builtin.LookupIPAddr(ctx, "icecave.com.au")
			Expect(err).ShouldNot(HaveOccurred())

			Expect(s).To(Equal(r))
		})
	})

	Describe("LookupMX", func() {
		It("returns the same results as the built-in implementation", func() {
			s, err := subject.LookupMX(ctx, "icecave.com.au")
			Expect(err).ShouldNot(HaveOccurred())

			r, err := builtin.LookupMX(ctx, "icecave.com.au")
			Expect(err).ShouldNot(HaveOccurred())

			Expect(s).To(HaveLen(len(r)))

			// expect preferences to be the same at each entry
			for idx := 0; idx < len(r); idx++ {
				a, b := s[idx], r[idx]
				Expect(a.Pref).To(Equal(b.Pref))
			}
		})
	})

	Describe("LookupNS", func() {
		It("returns the same results as the built-in implementation", func() {
			s, err := subject.LookupNS(ctx, "icecave.com.au")
			Expect(err).ShouldNot(HaveOccurred())

			r, err := builtin.LookupNS(ctx, "icecave.com.au")
			Expect(err).ShouldNot(HaveOccurred())

			Expect(s).To(ConsistOf(r))
		})
	})

	Describe("LookupPort", func() {
		It("returns the same results as the built-in implementation", func() {
			s, err := subject.LookupPort(ctx, "tcp", "https")
			Expect(err).ShouldNot(HaveOccurred())

			r, err := builtin.LookupPort(ctx, "tcp", "https")
			Expect(err).ShouldNot(HaveOccurred())

			Expect(s).To(Equal(r))
		})
	})

	Describe("LookupSRV", func() {
		It("returns the same results as the built-in implementation", func() {
			sc, s, err := subject.LookupSRV(ctx, "jabber", "tcp", "icecave.com.au")
			Expect(err).ShouldNot(HaveOccurred())

			rc, r, err := builtin.LookupSRV(ctx, "jabber", "tcp", "icecave.com.au")
			Expect(err).ShouldNot(HaveOccurred())

			Expect(sc).To(Equal(rc))
			Expect(s).To(ConsistOf(r))

			for i, rec := range s {
				Expect(rec.Priority).To(Equal(r[i].Priority))
			}
		})
	})

	Describe("LookupTXT", func() {
		It("returns the same results as the built-in implementation", func() {
			s, err := subject.LookupTXT(ctx, "icecave.com.au")
			Expect(err).ShouldNot(HaveOccurred())

			r, err := builtin.LookupTXT(ctx, "icecave.com.au")
			Expect(err).ShouldNot(HaveOccurred())

			Expect(s).To(ConsistOf(r))
		})
	})
})
