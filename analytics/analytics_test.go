package analytics

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

var (
	key       = os.Getenv("GOOGLE_ANALYTICS_KEY")
	accountID = os.Getenv("GOOGLE_ANALYTICS_ACCOUNT_ID")
	webID     = os.Getenv("GOOGLE_ANALYTICS_WEB_ID")
	profileID = os.Getenv("GOOGLE_ANALYTICS_PROFILE_ID")
)

var _ = Describe("Account List with invalid Key", func() {

	os.Setenv("KEY", "mockKey")

	analytics := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(analytics)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/accountList", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(AccountList)
	handler.ServeHTTP(recorder, request)

	Describe("Account List", func() {
		Context("account list", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Account List with valid param", func() {

	os.Setenv("KEY", key)

	analytics := Arguments{}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(analytics)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/accountList", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(AccountList)
	handler.ServeHTTP(recorder, request)

	Describe("Account List", func() {
		Context("account list", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("WebProperties List with invalid Key", func() {

	os.Setenv("KEY", "mockKey")

	analytics := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(analytics)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/webPropertiesList", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(WebPropertiesList)
	handler.ServeHTTP(recorder, request)

	Describe("Web Properties List", func() {
		Context("web properties list", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("WebProperties List with invalid param", func() {

	os.Setenv("KEY", key)

	analytics := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(analytics)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/webPropertiesList", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(WebPropertiesList)
	handler.ServeHTTP(recorder, request)

	Describe("Web Properties List", func() {
		Context("web properties list", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("WebProperties List with invalid param", func() {

	os.Setenv("KEY", key)

	analytics := Arguments{AccountID: "mockID"}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(analytics)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/webPropertiesList", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(WebPropertiesList)
	handler.ServeHTTP(recorder, request)

	Describe("Web Properties List", func() {
		Context("web properties list", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("WebProperties List with valid param", func() {

	os.Setenv("KEY", key)

	analytics := Arguments{AccountID: accountID}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(analytics)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/webPropertiesList", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(WebPropertiesList)
	handler.ServeHTTP(recorder, request)

	Describe("Web Properties List", func() {
		Context("web properties list", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Profiles List with invalid key", func() {

	os.Setenv("KEY", "mockKey")

	analytics := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(analytics)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/profileList", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ProfilesList)
	handler.ServeHTTP(recorder, request)

	Describe("Profiles List", func() {
		Context("Profiles list", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Profiles List with invalid param", func() {

	os.Setenv("KEY", key)

	analytics := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(analytics)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/profileList", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ProfilesList)
	handler.ServeHTTP(recorder, request)

	Describe("Profiles List", func() {
		Context("Profiles list", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Profiles List with invalid param", func() {

	os.Setenv("KEY", key)

	analytics := Arguments{AccountID: "mockID", WebPropertyID: "mockID"}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(analytics)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/profileList", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ProfilesList)
	handler.ServeHTTP(recorder, request)

	Describe("Profiles List", func() {
		Context("Profiles list", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Profiles List with valid param", func() {

	os.Setenv("KEY", key)

	analytics := Arguments{AccountID: accountID, WebPropertyID: webID}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(analytics)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/profileList", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ProfilesList)
	handler.ServeHTTP(recorder, request)

	Describe("Profiles List", func() {
		Context("Profiles list", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("RealtimeData with invalid key", func() {

	os.Setenv("KEY", "mockKey")

	analytics := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(analytics)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/realtime", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(RealtimeData)
	handler.ServeHTTP(recorder, request)

	Describe("RealtimeData", func() {
		Context("Realtime Data", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("RealtimeData with invalid param", func() {

	os.Setenv("KEY", key)

	analytics := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(analytics)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/realtime", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(RealtimeData)
	handler.ServeHTTP(recorder, request)

	Describe("RealtimeData", func() {
		Context("Realtime Data", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("RealtimeData with invalid param", func() {

	os.Setenv("KEY", key)

	analytics := Arguments{ProfileID: "mockID"}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(analytics)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/realtime", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(RealtimeData)
	handler.ServeHTTP(recorder, request)

	Describe("RealtimeData", func() {
		Context("Realtime Data", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("RealtimeData with valid param", func() {

	os.Setenv("KEY", key)

	analytics := Arguments{ProfileID: profileID}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(analytics)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/realtime", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(RealtimeData)
	handler.ServeHTTP(recorder, request)

	Describe("RealtimeData", func() {
		Context("Realtime Data", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
