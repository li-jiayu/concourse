package pipelineserver_test

import (
	"errors"
	"net/http/httptest"

	"github.com/concourse/concourse/atc/api/pipelineserver"
	"github.com/concourse/concourse/atc/api/pipelineserver/pipelineserverfakes"
	"github.com/concourse/concourse/atc/db/dbfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go:generate counterfeiter code.cloudfoundry.org/lager.Logger

var _ = Describe("Archive Handler", func() {
	It("logs database errors", func() {
		fakeLogger := new(pipelineserverfakes.FakeLogger)
		server := pipelineserver.NewServer(
			fakeLogger,
			new(dbfakes.FakeTeamFactory),
			new(dbfakes.FakePipelineFactory),
			"",
		)
		dbPipeline := new(dbfakes.FakePipeline)
		expectedError := errors.New("db error")
		dbPipeline.ArchiveReturns(expectedError)

		server.ArchivePipeline(dbPipeline).ServeHTTP(
			httptest.NewRecorder(),
			httptest.NewRequest("PUT", "http://example.com", nil),
		)

		Expect(fakeLogger.ErrorCallCount()).To(Equal(1))
		action, actualError, _ := fakeLogger.ErrorArgsForCall(0)
		Expect(action).To(Equal("archive-pipeline"), "wrong action name")
		Expect(actualError).To(Equal(expectedError))
	})

	// TODO do not log when there is no error
})
