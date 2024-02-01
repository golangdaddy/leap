package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/golangdaddy/leap/sdk/assetlayer"
	"github.com/golangdaddy/leap/sdk/cloudfunc"
)

type AssetlayerWallet struct {
}

// api-assetlayer
func (app *App) EntrypointASSETLAYER(w http.ResponseWriter, r *http.Request) {

	if cloudfunc.HandleCORS(w, r, "*") {
		return
	}

	_, err := app.GetSessionUser(r)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusUnauthorized)
		return
	}

	// get collection metadata
	parentID, err := cloudfunc.QueryParam(r, "parent")
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}
	parent, err := app.GetMetadata(parentID)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusNotFound)
		return
	}

	// get function
	function, err := cloudfunc.QueryParam(r, "function")
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}

	switch r.Method {

	case "GET":

		switch function {

		case "overview":

			data := map[string][]*assetlayer.Collection{}

			slots, err := app.Assetlayer().GetSlots()
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			for _, slot := range slots {
				if data[slot.SlotName] == nil {
					data[slot.SlotName] = []*assetlayer.Collection{}
				}
				data[slot.SlotName], err = app.Assetlayer().GetCollections(slot.SlotID)
				if err != nil {
					cloudfunc.HttpError(w, err, http.StatusInternalServerError)
					return
				}
			}

			if err := cloudfunc.ServeJSON(w, data); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

		// return a list of assets by account
		case "assets":

			println(parent.ID)

			println("checking wallet:", parent.Wallet)
			list, err := app.Assetlayer().GetWalletAssets(parent.Wallet, false, false)
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}
			if err := cloudfunc.ServeJSON(w, list); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			return

		// return a list of wallet assets by account
		case "walletassets":

			println(parent.ID)

			println("checking wallet:", parent.Wallet)
			list, err := app.Assetlayer().GetWalletAssets(parent.Wallet, false, false)
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}
			if err := cloudfunc.ServeJSON(w, list); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			return

		default:
			err := fmt.Errorf("function not found: %s", function)
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return
		}

	default:
		err := errors.New("method not allowed: " + r.Method)
		cloudfunc.HttpError(w, err, http.StatusMethodNotAllowed)
		return
	}
}
