package services

import (
	"delivery_svc/config"
	"delivery_svc/dto"
	"delivery_svc/interfaces/mocks"
	"delivery_svc/model"
	"delivery_svc/utils"
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"googlemaps.github.io/maps"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	response = `{
  "destination_addresses": ["-6.362764,106.827049"],
  "origin_addresses": ["-6.238270,106.975571"],
  "rows": [
    {
      "elements": [
        {
          "distance": {
            "text": "12.5 km",
            "value": 12535
          },
          "duration": {
            "text": "18 mins",
            "value": 1083
          },
          "duration_in_traffic": {
            "text": "19 mins",
            "value": 1134
          },
          "status": "OK"
        }
      ]
    }
  ],
  "status": "OK"
}`
	GmapApiKey = "FAKE_GOOGLE_API_KEY_FOR_TESTING"
)

type DeliveryServiceSuite struct {
	suite.Suite
	mockRepo        *mocks.IDeliveryRepository
	deliveryService *DeliveryService
	mapsClient      *maps.Client
}

func TestDeliveryServiceSuiteInit(t *testing.T) {
	suite.Run(t, new(DeliveryServiceSuite))
}

func (s *DeliveryServiceSuite) SetupSuite() {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		fmt.Fprintln(w, response)
	}))
	s.mapsClient, _ = maps.NewClient(maps.WithAPIKey(GmapApiKey), maps.WithBaseURL(server.URL))
	s.mockRepo = &mocks.IDeliveryRepository{}
	s.deliveryService = &DeliveryService{
		DeliveryRepository: s.mockRepo,
		mapsClientInstance: s.mapsClient,
	}
}

func (s *DeliveryServiceSuite) Test_service_ListOrders() {

	//create fake data
	var m model.Orders
	errM := faker.FakeData(&m)
	assert.NoError(s.T(), errM)
	mockList := make([]model.Orders, 0)
	mockList = append(mockList, m)

	page := 1
	limit := 10
	paginator := &utils.Paginator{
		TotalRecord: 10,
		TotalPage:   1,
		Records:     &mockList,
		Offset:      0,
		Limit:       limit,
		Page:        page,
		PrevPage:    1,
		NextPage:    1,
	}

	mockDetail := s.mockRepo.On("FindAllOrders", mock.AnythingOfType("int"), mock.AnythingOfType("int"))
	_ = mockDetail.Return(paginator, nil)

	listOrders, _ := s.deliveryService.ListOfDeliveryOrders(page, limit)
	assert.NotEmpty(s.T(), listOrders)

}

func (s *DeliveryServiceSuite) Test_service_CreateOrder() {
	dtoDeliveryReq := dto.DeliveryRequest{
		Origin:      []string{"-6.362764", "106.827049"},
		Destination: []string{"-6.238270", "106.975571"},
	}

	responseGmapdata := maps.DistanceMatrixResponse{}
	json.Unmarshal([]byte(response), &responseGmapdata)

	s.mockRepo.On("DistanceMatrix", mock.Anything).Return(responseGmapdata, nil)

	mResponse := model.Orders{
		Id:       1,
		Status:   config.UNASSIGNED_STATUS,
		Distance: responseGmapdata.Rows[0].Elements[0].Distance.Meters,
	}
	mockTakeUpOrder := s.mockRepo.On("SaveOrders", mock.Anything)
	_ = mockTakeUpOrder.Return(mResponse, nil)

	createOrder, errResp := s.deliveryService.CreateOrder(dtoDeliveryReq)
	assert.Nil(s.T(), errResp)
	assert.NotEmpty(s.T(), createOrder)
	assert.NotNil(s.T(), createOrder)

}

func (s *DeliveryServiceSuite) Test_service_TakeOrder() {
	dtoTakeOrder := dto.TakeOrder{
		Status: config.TAKEN_STATUS,
	}

	id := 1
	mockTakeUpOrder := s.mockRepo.On("TakeUpOrder", mock.AnythingOfType("int"))
	mockTakeUpOrder.Return(nil)

	respTakeOrder, errResps := s.deliveryService.TakeOrder(dtoTakeOrder, id)

	assert.NotEmpty(s.T(), respTakeOrder)
	assert.NotNil(s.T(), respTakeOrder)
	assert.Equal(s.T(), respTakeOrder.Status, config.SUCCESS_STATUS)
	assert.Nil(s.T(), errResps)

}
