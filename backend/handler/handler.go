package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/joshm998/drover/httphandler"
	"github.com/joshm998/drover/model"
	log "github.com/sirupsen/logrus"
)

type ctx struct {
	store Service
	h     func(Service, http.ResponseWriter, *http.Request)
}

func (g *ctx) handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		g.h(g.store, w, r)
	}
}

func Handler(store Service) http.Handler {
	r := chi.NewRouter()
	getPrinterById := ctx{store: store, h: getPrinterById}
	createPrinter := ctx{store: store, h: createPrinter}
	updatePrinter := ctx{store: store, h: updatePrinter}
	deletePrinter := ctx{store: store, h: deletePrinter}

	r.Get(httphandler.WrapHandlerFunc("/printer/{id}", "get printer", getPrinterById.handle()))
	r.Post(httphandler.WrapHandlerFunc("/printer", "create printer", createPrinter.handle()))
	r.Put(httphandler.WrapHandlerFunc("/printer/{id}", "update printer", updatePrinter.handle()))
	r.Delete(httphandler.WrapHandlerFunc("/printer/{id}}", "delete printer", deletePrinter.handle()))

	return r
}

func createPrinter(store Service, w http.ResponseWriter, r *http.Request) {

	data := &Request{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httphandler.ErrInvalidRequest(err, "Invalid Request"))
		return
	}
	recordSchema := data.Printers
	// services, err :=
	store.CreateRecordCoreTeam(recordSchema)
	// if err != nil {
	// 	log.Errorf("Unable To Fetch stats ", httphandler.Error(err).Code, services, err)
	// 	httphandler.ErrInvalidRequest(err, "Unable To Fetch Services ")
	// 	return
	// }
	render.Status(r, http.StatusOK)
	render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, nil))
}

func getPrinterById(store Service, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	services, err := store.GetRecordSetPost(id)
	if err != nil {
		render.Render(w, r, httphandler.ErrNotFound(err, "Printer not Found"))
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, services))
}

func updatePrinter(store Service, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	payload := &Request{}
	if err := render.Bind(r, payload); err != nil {
		render.Render(w, r, httphandler.ErrInvalidRequest(err, "Invalid Request"))
		return
	}
	recordSchema := payload.Printers
	// services, err :=
	store.UpdatePrinter(id, recordSchema)
	// if err != nil {
	// 	log.Errorf("Unable To Fetch stats ", httphandler.Error(err).Code, services, err)
	// 	httphandler.ErrInvalidRequest(err, "Unable To Fetch Services ")
	// 	return
	// }
	render.Status(r, http.StatusOK)
	render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, nil))
}

func deletePrinter(store Service, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	services, err := store.DeletePrinter(id)
	if err != nil {
		log.Errorf("Unable To Fetch stats ", httphandler.Error(err).Code, services, err)
		httphandler.ErrInvalidRequest(err, "Unable To Fetch Services ")
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, services))
}

type Request struct {
	*model.Printers
}

func (a *Request) Bind(r *http.Request) error {
	//TODO: to be expanded
	return nil
}

type Response struct {
	Meta interface{}
	Data interface{}
}
