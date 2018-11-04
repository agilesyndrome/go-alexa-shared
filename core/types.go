package core

import (
  "context"
  "github.com/arienmalec/alexa-go"
)


type (
  AlexaHandler func (alexa.Request) (alexa.Response, error)
  LambdaHandler func (context.Context, alexa.Request) (alexa.Response, error)
)

func init() {

}
