package helpers

import (
	"encoding/json"
	"github.com/verottaa/user-module/common"
	"net/http"
)

type jsonWorker struct {
}

type Encoder interface {
	Encode(w http.ResponseWriter, result interface{}) error
}

type Decoder interface {
	Decode(r *http.Request, target interface{}) error
}

type JsonWorker interface {
	common.Destroyable
	Encoder
	Decoder
}

var destroyJsonWorkerCh = make(chan bool)
var jsonWorkerInstance *jsonWorker

func GetJsonWorker() JsonWorker {
	once.Do(func() {
		jsonWorkerInstance = createJsonWorkerInstance()
		go func() {
			for
			{
				select {
				case <-destroyJsonWorkerCh:
					return
				}
			}
		}()
	})

	return jsonWorkerInstance
}

func createJsonWorkerInstance() *jsonWorker {
	return &jsonWorker{}
}

func (jsonWorker jsonWorker) Destroy() {
	destroyJsonWorkerCh <- true
	close(destroyJsonWorkerCh)
	jsonWorkerInstance = nil
}

func (jsonWorker jsonWorker) Encode(w http.ResponseWriter, result interface{}) error {
	return json.NewEncoder(w).Encode(result)
}

func (jsonWorker jsonWorker) Decode(r *http.Request, target interface{}) error {
	return json.NewDecoder(r.Body).Decode(&target)
}
