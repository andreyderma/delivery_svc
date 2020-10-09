package apis

import (
	"delivery_svc/dto"
	"delivery_svc/services"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type API struct {
	DeliveryService services.DeliveryService
}

func ProvideAPI(
	deliveryService services.DeliveryService) API {
	return API{
		DeliveryService: deliveryService,
	}
}

// Delivery svc godoc
// @tags delivery
// @Summary Create an order
// @Description Create an order
// @Param req body dto.DeliveryRequest true "DeliveryRequest"
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.DeliveryResponse
// @Failure 400 {object} utils.ErrResponse
// @Router /delivery_svc/orders [post]
func (u API) Orders(w http.ResponseWriter, r *http.Request) {
	var dtoDeliveryRequest dto.DeliveryRequest
	json.NewDecoder(r.Body).Decode(&dtoDeliveryRequest)

	resps, err := u.DeliveryService.CreateOrder(dtoDeliveryRequest)
	if err != nil {
		render.Render(w, r, err)
		return
	}
	render.JSON(w, r, resps)
}

// Delivery svc godoc
// @tags delivery
// @Summary Take an order
// @Description Take an order
// @Param req body dto.TakeOrder true "TakeOrder"
// @Param   id     path    int     true        "ID"
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.TakeOrder
// @Failure 400 {object} utils.ErrResponse
// @Router /delivery_svc/orders/{id} [patch]
func (u API) TakeOrder(w http.ResponseWriter, r *http.Request) {
	var dtoTakeOrder dto.TakeOrder
	json.NewDecoder(r.Body).Decode(&dtoTakeOrder)
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	resps, err := u.DeliveryService.TakeOrder(dtoTakeOrder, id)
	if err != nil {
		render.Render(w, r, err)
		return
	}
	render.JSON(w, r, resps)
}

// Delivery svc godoc
// @tags delivery
// @Summary List of orders
// @Description List of orders
// @Param   page     query    string     false        "page"
// @Param   limit     query    string     false        "limit"
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.DeliveryResponse
// @Failure 400 {object} utils.ErrResponse
// @Router /delivery_svc/orders [get]
func (u API) ListOrders(w http.ResponseWriter, r *http.Request) {
	pageNumber, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	resps, err := u.DeliveryService.ListOfDeliveryOrders(pageNumber, limit)
	if err != nil {
		render.Render(w, r, err)
		return
	}
	render.JSON(w, r, resps)
}
