package services

import (
	"context"
	"delivery_svc/config"
	"delivery_svc/dto"
	"delivery_svc/interfaces"
	"delivery_svc/model"
	"delivery_svc/utils"
	"encoding/json"
	"errors"
	"github.com/go-chi/render"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"
	"googlemaps.github.io/maps"
	"log"
	"strings"
	"time"
)

type DeliveryService struct {
	DeliveryRepository interfaces.IDeliveryRepository
	mapsClientInstance *maps.Client
}

func ProvideDeliveryService(
	d interfaces.IDeliveryRepository,
	mapsClientInstance *maps.Client) DeliveryService {
	return DeliveryService{
		DeliveryRepository: d,
		mapsClientInstance: mapsClientInstance,
	}
}

func (p *DeliveryService) TakeOrder(dtoTakeOrder dto.TakeOrder, id int) (resps dto.TakeOrder, errRender render.Renderer) {
	errMessageVal, errValidation := utils.CustomValidator{
		CustomV: validator.New(),
		Trans:   en.New(),
		Dto:     dtoTakeOrder,
	}.TranslateValidator()

	jsm, _ := json.Marshal(dtoTakeOrder)
	log.Println("Incoming request", string(jsm))

	if errValidation != nil {
		log.Println(errMessageVal)
		return resps, utils.ErrInvalidRequest(errValidation, errMessageVal...)
	}

	log.Println(dtoTakeOrder.Status)
	if dtoTakeOrder.Status != config.TAKEN_STATUS {
		return dto.TakeOrder{}, utils.ErrInvalidRequest(errors.New("Invalid status"), utils.INVALID_REQUEST)
	}

	err := p.DeliveryRepository.TakeUpOrder(id)
	if err != nil {
		return resps, utils.ErrInvalidRequest(err, utils.INVALID_REQUEST)
	}

	resps.Status = config.SUCCESS_STATUS

	return resps, nil
}

func (p *DeliveryService) CreateOrder(dtoDeliveryRequest dto.DeliveryRequest) (resps dto.DeliveryResponse, errRender render.Renderer) {
	errMessageVal, errValidation := utils.CustomValidator{
		CustomV: validator.New(),
		Trans:   en.New(),
		Dto:     dtoDeliveryRequest,
	}.TranslateValidator()

	jsm, _ := json.Marshal(dtoDeliveryRequest)
	log.Println("Incoming request", string(jsm))

	if errValidation != nil {
		log.Println(errMessageVal)
		return resps, utils.ErrInvalidRequest(errValidation, errMessageVal...)
	}

	origin := []string{strings.Join(dtoDeliveryRequest.Origin, ",")}
	dest := []string{strings.Join(dtoDeliveryRequest.Destination, ",")}

	if len(dtoDeliveryRequest.Destination) != 2 && len(dtoDeliveryRequest.Origin) != 2 {
		return resps, utils.ErrInvalidRequest(errors.New("Array of destination should contains 2 elements"), utils.INVALID_REQUEST)
	}

	r := &maps.DistanceMatrixRequest{
		Origins:       origin,
		Destinations:  dest,
		DepartureTime: "now",
	}
	route, errMatrix := p.mapsClientInstance.DistanceMatrix(context.Background(), r)
	if errMatrix != nil {
		return resps, utils.ErrInternal(errMatrix, utils.INTERNAL_ERROR)
	}

	distance := route.Rows[0].Elements[0].Distance.Meters

	var m model.Orders
	m.StartLatitude = dtoDeliveryRequest.Origin[0]
	m.StartLongitude = dtoDeliveryRequest.Origin[1]

	m.EndLatitude = dtoDeliveryRequest.Destination[0]
	m.EndLongitude = dtoDeliveryRequest.Destination[1]
	m.CreatedAt = time.Now()
	m.Distance = distance
	m.Status = config.UNASSIGNED_STATUS

	modelResponse, err := p.DeliveryRepository.SaveOrders(m)
	if err != nil {
		return resps, utils.ErrInternal(err, utils.INTERNAL_ERROR)
	}

	resp := dto.DeliveryResponse{
		Id:       modelResponse.Id,
		Distance: distance,
		Status:   modelResponse.Status,
	}

	return resp, nil
}

func (p *DeliveryService) ListOfDeliveryOrders(page int, limit int) (resps []dto.DeliveryResponse, errRender render.Renderer) {
	deliveryOrders := p.DeliveryRepository.FindAllOrders(page, limit)
	var recordData = deliveryOrders.Records.(*[]model.Orders)
	r := make([]dto.DeliveryResponse, 0)
	for _, d := range *recordData {
		r = append(r, dto.DeliveryResponse{
			Id:       d.Id,
			Distance: d.Distance,
			Status:   d.Status,
		})
	}

	return r, nil
}
