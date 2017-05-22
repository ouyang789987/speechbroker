package parseibm_test

import (
	. "github.com/blabbertabber/DiarizerServer/IBMJson/parseibm"

	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
)

var _ = Describe("Parseibm", func() {

	It("should parse an empty JSON properly", func() {
		source := []byte(`{}`)
		expectation := IBMTranscription{}
		result := IBMTranscription{}
		err := json.Unmarshal(source, &result)
		Expect(err).To(BeNil())
		Expect(result).To(Equal(expectation))
	})
	It("should parse a minimal JSON properly", func() {
		source := []byte(`{"result_index": 1, "results": [], "speaker_labels": []}`)
		expectation := IBMTranscription{ResultIndex: 1, Results: []Result{}, SpeakerLabels: []Speaker{}}
		result := IBMTranscription{}
		err := json.Unmarshal(source, &result)
		Expect(err).To(BeNil())
		Expect(result).To(Equal(expectation))
	})
	It("should parse a full JSON properly", func() {
		source, err := ioutil.ReadFile("../../assets/test/ibm.json")
		Expect(err).To(BeNil())
		expectation := IBMTranscription{
			ResultIndex: 0,
			Results: []Result{{
				Alternatives: []Alternative{{
					Confidence: 0.694,
					Timestamps: []Timestamp{{
						Word: "design",
						From: 2.37,
						To:   3.13,
					}},
					Transcript: "design",
				}},
				Final: true,
			}},
			SpeakerLabels: []Speaker{{
				Confidence: 0.488,
				Final:      false,
				From:       2.37,
				To:         3.13,
			}},
		}
		result := IBMTranscription{}
		err = json.Unmarshal(source, &result)
		Expect(err).To(BeNil())
		Expect(result).To(Equal(expectation))
	})
})