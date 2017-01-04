package main

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dice as a service", func() {
	Describe("GET /roll", func() {
		var (
			writer  *httptest.ResponseRecorder
			request *http.Request
		)

		Context("Default D6 dice", func() {
			It("returns a random roll with the range of 1 and 6 and D6 dice type", func() {
				writer = httptest.NewRecorder()
				request, _ = http.NewRequest("GET", "/roll", nil)
				diceServer(writer, request)
				Expect(writer.Body.Bytes()).To(SatisfyAny(
					(MatchJSON(`{"Roll":1,"Type":"D6"}`)),
					(MatchJSON(`{"Roll":2,"Type":"D6"}`)),
					(MatchJSON(`{"Roll":3,"Type":"D6"}`)),
					(MatchJSON(`{"Roll":4,"Type":"D6"}`)),
					(MatchJSON(`{"Roll":5,"Type":"D6"}`)),
					(MatchJSON(`{"Roll":6,"Type":"D6"}`))))
			})
		})

		Context("With die params set to D10", func() {
			It("returns a random roll within the range of 1 and 10 and D10 dice type", func() {
				writer = httptest.NewRecorder()
				request, _ = http.NewRequest("GET", "/roll?die=D10", nil)

				diceServer(writer, request)
				Expect(writer.Body.Bytes()).To(SatisfyAny(
					(MatchJSON(`{"Roll":1,"Type":"D10"}`)),
					(MatchJSON(`{"Roll":2,"Type":"D10"}`)),
					(MatchJSON(`{"Roll":3,"Type":"D10"}`)),
					(MatchJSON(`{"Roll":4,"Type":"D10"}`)),
					(MatchJSON(`{"Roll":5,"Type":"D10"}`)),
					(MatchJSON(`{"Roll":6,"Type":"D10"}`)),
					(MatchJSON(`{"Roll":7,"Type":"D10"}`)),
					(MatchJSON(`{"Roll":8,"Type":"D10"}`)),
					(MatchJSON(`{"Roll":9,"Type":"D10"}`)),
					(MatchJSON(`{"Roll":10,"Type":"D10"}`)),
				))
			})
		})

		Context("With invalid die param value", func() {
			BeforeEach(func() {
				writer = httptest.NewRecorder()
				request, _ = http.NewRequest("GET", "/roll?die=XYZ", nil)
				diceServer(writer, request)
			})

			It("returns a 400 status code", func() {
				Expect(writer.Code).To(Equal(400))
			})

			It("returns an error message", func() {
				Expect(writer.Body.Bytes()).To(ContainSubstring("Error: Invalid die params"))
			})
		})
	})
})
