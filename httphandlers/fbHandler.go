package httphandlers

import (
	"net/http"
	"fmt"
	"github.com/Instamojo/Instamojo++/models"
	"encoding/json"
)

func fbHook(w http.ResponseWriter, r *http.Request) {
	fbModel := &models.FBModel{}
	err := json.NewDecoder(r.Body).Decode(fbModel)
	if err!= nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusOK)
	}

	fmt.Println(fbModel.Entry[0].Messaging[0].Message.Text)

}
