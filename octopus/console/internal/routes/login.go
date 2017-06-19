package user

import (
	"context"
	"fmt"
	"net/http"
)

type loginPayload struct {
}

type responseLoginOK struct {
}

// login user in system
// @Route {
// 		url = /login
//		method = post
//      payload = loginPayload
//		200 = responseLoginOK
//		400 = controller.ErrorResponseSimple
// }
func (c Controller) login(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	payload := c.MustGetPayload(ctx).(*loginPayload)
	fmt.Println(payload)

	resp := responseLoginOK{}
	c.OKResponse(w, resp)
}