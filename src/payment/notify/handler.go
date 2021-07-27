package notify

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"net/http"
)

type Handler struct {
	App        *kernel.ApplicationPaymentInterface
	Messages   *object.HashMap
	Fail       string
	Attributes *object.HashMap
	Check      bool
	Sign       bool

	Handle func(closure func(payload ...interface{}) interface{}) *http.Response
}

const SUCCESS = "SUCCESS"
const FAIL = "FAIL"

func NewHandler(app *kernel.ApplicationPaymentInterface) *Handler {
	return &Handler{
		App: app,
		Check: true,
		Sign: false,
	}
}

func (handler *Handler) RespondWith(attributes *object.HashMap, sign bool) *Handler {

	handler.Attributes = attributes
	handler.Sign = sign

	return handler
}
func (handler *Handler) ToResponse() (response *http.Response) {

	return nil
}

func (handler *Handler) GetMessage() (message *object.HashMap, err error) {

	if handler.Messages!=nil{
		return handler.Messages, nil
	}



	return message,nil
}

func (handler *Handler) Strict(result interface{}) {

	return
}
