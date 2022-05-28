package movies

import (
	"bytes"
	"encoding/json"
	"log"
)


var (
	Movies []Movie
)


func Main() {

}



func JsonMarshal(data interface{}) []byte {
	var buff bytes.Buffer
	encoder := json.NewEncoder(&buff)
	encoder.SetEscapeHTML(false)
	encoder.Encode(data)
	return buff.Bytes()
}


func HanleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}