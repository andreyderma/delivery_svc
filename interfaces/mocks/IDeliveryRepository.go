// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	model "delivery_svc/model"

	mock "github.com/stretchr/testify/mock"

	utils "delivery_svc/utils"
)

// IDeliveryRepository is an autogenerated mock type for the IDeliveryRepository type
type IDeliveryRepository struct {
	mock.Mock
}

// FindAllOrders provides a mock function with given fields: page, limit
func (_m *IDeliveryRepository) FindAllOrders(page int, limit int) *utils.Paginator {
	ret := _m.Called(page, limit)

	var r0 *utils.Paginator
	if rf, ok := ret.Get(0).(func(int, int) *utils.Paginator); ok {
		r0 = rf(page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.Paginator)
		}
	}

	return r0
}

// SaveOrders provides a mock function with given fields: orders
func (_m *IDeliveryRepository) SaveOrders(orders model.Orders) (model.Orders, error) {
	ret := _m.Called(orders)

	var r0 model.Orders
	if rf, ok := ret.Get(0).(func(model.Orders) model.Orders); ok {
		r0 = rf(orders)
	} else {
		r0 = ret.Get(0).(model.Orders)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Orders) error); ok {
		r1 = rf(orders)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TakeUpOrder provides a mock function with given fields: id
func (_m *IDeliveryRepository) TakeUpOrder(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}