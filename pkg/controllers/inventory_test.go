package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Abhishek-Omniful/IMS/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockInventoryModel mocks the models layer
type MockInventoryModel struct {
	mock.Mock
}

func (m *MockInventoryModel) UpsertInventory(inv *models.Inventory) error {
	args := m.Called(inv)
	return args.Error(0)
}

func (m *MockInventoryModel) CheckInventoryStatus(skuID, hubID int64, qty int) bool {
	args := m.Called(skuID, hubID, qty)
	return args.Bool(0)
}

func (m *MockInventoryModel) ValidateOrderByHubAndSKU(hubID, skuID int64) bool {
	args := m.Called(hubID, skuID)
	return args.Bool(0)
}

func (m *MockInventoryModel) GetInventoryByHub(hubID int64) ([]models.Inventory, error) {
	args := m.Called(hubID)
	return args.Get(0).([]models.Inventory), args.Error(1)
}

func (m *MockInventoryModel) GetInventoryBySKU(skuID int64) ([]models.Inventory, error) {
	args := m.Called(skuID)
	return args.Get(0).([]models.Inventory), args.Error(1)
}

func (m *MockInventoryModel) GetInventoryBySKUAndHub(skuID, hubID int64) (*models.Inventory, error) {
	args := m.Called(skuID, hubID)
	return args.Get(0).(*models.Inventory), args.Error(1)
}

func (m *MockInventoryModel) GetAllInventory() (*[]models.Inventory, error) {
	args := m.Called()
	return args.Get(0).(*[]models.Inventory), args.Error(1)
}

func TestInventoryHandlers(t *testing.T) {

	tests := []struct {
		name         string
		method       string
		url          string
		body         interface{}
		queryParams  map[string]string
		mockSetup    func(*MockInventoryModel)
		handler      gin.HandlerFunc
		expectedCode int
	}{
		{
			name:   "UpsertInventory - Success",
			method: http.MethodPost,
			url:    "/inventory",
			body:   models.Inventory{HubID: 1, SkuID: 1, Quantity: 10, UnitPrice: 100},
			mockSetup: func(mockmodel *MockInventoryModel) {
				mockmodel.On("UpsertInventory", &models.Inventory{HubID: 1, SkuID: 1, Quantity: 10, UnitPrice: 100}).Return(nil)
			},
			handler:      UpsertInventory,
			expectedCode: 200,
		},
		{
			name:         "UpsertInventory - Invalid JSON",
			method:       http.MethodPost,
			url:          "/inventory",
			body:         "{invalid-json}",
			mockSetup:    func(_ *MockInventoryModel) {},
			handler:      UpsertInventory,
			expectedCode: 400,
		},
		{
			name:   "UpsertInventory - Internal Error",
			method: http.MethodPost,
			url:    "/inventory",
			body:   models.Inventory{HubID: 1, SkuID: 1, Quantity: 10, UnitPrice: 100},
			mockSetup: func(mockmodel *MockInventoryModel) {
				mockmodel.On("UpsertInventory", &models.Inventory{HubID: 1, SkuID: 1, Quantity: 10, UnitPrice: 100}).Return(errors.New("db error"))
			},
			handler:      UpsertInventory,
			expectedCode: 500,
		},
		{
			name:   "CheckInventoryStatus - Success",
			method: http.MethodGet,
			url:    "/check-inventory",
			queryParams: map[string]string{
				"sku_id":   "1",
				"hub_id":   "2",
				"quantity": "5",
			},
			mockSetup: func(mockmodel *MockInventoryModel) {
				mockmodel.On("CheckInventoryStatus", int64(1), int64(2), 5).Return(true)
			},
			handler:      CheckInventoryStatus,
			expectedCode: 200,
		},
		{
			name:   "CheckInventoryStatus - Not Available",
			method: http.MethodGet,
			url:    "/check-inventory",
			queryParams: map[string]string{
				"sku_id":   "1",
				"hub_id":   "2",
				"quantity": "99",
			},
			mockSetup: func(mockmodel *MockInventoryModel) {
				mockmodel.On("CheckInventoryStatus", int64(1), int64(2), 99).Return(false)
			},
			handler:      CheckInventoryStatus,
			expectedCode: 404,
		},
		{
			name:   "CheckInventoryStatus - Invalid Params",
			method: http.MethodGet,
			url:    "/check-inventory",
			queryParams: map[string]string{
				"sku_id":   "a",
				"hub_id":   "b",
				"quantity": "c",
			},
			mockSetup:    func(mockmodel *MockInventoryModel) {},
			handler:      CheckInventoryStatus,
			expectedCode: 400,
		},
		{
			name:   "ValidateOrderRequest - Success",
			method: http.MethodGet,
			url:    "/validate/1/2",
			mockSetup: func(mockmodel *MockInventoryModel) {
				mockmodel.On("ValidateOrderByHubAndSKU", int64(2), int64(1)).Return(true)
			},
			handler: func(c *gin.Context) {
				c.Params = gin.Params{
					{Key: "sku_id", Value: "1"},
					{Key: "hub_id", Value: "2"},
				}
				ValidateOrderRequest(c)
			},
			expectedCode: 200,
		},
		{
			name:   "ValidateOrderRequest - Fail",
			method: http.MethodGet,
			url:    "/validate/1/2",
			mockSetup: func(mockmodel *MockInventoryModel) {
				mockmodel.On("ValidateOrderByHubAndSKU", int64(2), int64(1)).Return(false)
			},
			handler: func(c *gin.Context) {
				c.Params = gin.Params{
					{Key: "sku_id", Value: "1"},
					{Key: "hub_id", Value: "2"},
				}
				ValidateOrderRequest(c)
			},
			expectedCode: 400,
		},
		{
			name:      "ValidateOrderRequest - Invalid Params",
			method:    http.MethodGet,
			url:       "/validate/invalid/hub",
			mockSetup: func(mockmodel *MockInventoryModel) {},
			handler: func(c *gin.Context) {
				c.Params = gin.Params{
					{Key: "sku_id", Value: "invalid"},
					{Key: "hub_id", Value: "hub"},
				}
				ValidateOrderRequest(c)
			},
			expectedCode: 400,
		},
		{
			name:   "GetInventoryByHub - Success",
			method: http.MethodGet,
			url:    "/inventory/hub/1",
			mockSetup: func(mockmodel *MockInventoryModel) {
				mockmodel.On("GetInventoryByHub", int64(1)).Return([]models.Inventory{
					{SkuID: 1, HubID: 1, Quantity: 5, UnitPrice: 100},
				}, nil)
			},
			handler: func(c *gin.Context) {
				c.Params = gin.Params{{Key: "hub_id", Value: "1"}}
				GetInventoryByHub(c)
			},
			expectedCode: 200,
		},
		{
			name:   "GetInventoryBySKU - Success",
			method: http.MethodGet,
			url:    "/inventory/sku/1",
			mockSetup: func(mockmodel *MockInventoryModel) {
				mockmodel.On("GetInventoryBySKU", int64(1)).Return([]models.Inventory{
					{SkuID: 1, HubID: 1, Quantity: 20, UnitPrice: 150},
				}, nil)
			},
			handler: func(c *gin.Context) {
				c.Params = gin.Params{{Key: "sku_id", Value: "1"}}
				GetInventoryBySKU(c)
			},
			expectedCode: 200,
		},
		{
			name:   "GetInventoryBySKUAndHub - Success",
			method: http.MethodGet,
			url:    "/inventory/sku/1/hub/2",
			mockSetup: func(mockmodel *MockInventoryModel) {
				mockmodel.On("GetInventoryBySKUAndHub", int64(1), int64(2)).Return(&models.Inventory{SkuID: 1, HubID: 2, Quantity: 15, UnitPrice: 120}, nil)
			},
			handler: func(c *gin.Context) {
				c.Params = gin.Params{{Key: "sku_id", Value: "1"}, {Key: "hub_id", Value: "2"}}
				GetInventoryBySKUAndHub(c)
			},
			expectedCode: 200,
		},
		{
			name:   "GetAllInventory - Success",
			method: http.MethodGet,
			url:    "/inventory",
			mockSetup: func(mockmodel *MockInventoryModel) {
				mockmodel.On("GetAllInventory").Return(&[]models.Inventory{
					{SkuID: 1, HubID: 1, Quantity: 50, UnitPrice: 100},
					{SkuID: 2, HubID: 2, Quantity: 60, UnitPrice: 200},
				}, nil)
			},
			handler:      GetAllInventory,
			expectedCode: 200,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockmodel := new(MockInventoryModel)
			if tc.mockSetup != nil {
				tc.mockSetup(mockmodel)
			}
			// Backup original functions
			origUpsert := models.UpsertInventory
			origCheck := models.CheckInventoryStatus
			origValidate := models.ValidateOrderByHubAndSKU
			origGetByHub := models.GetInventoryByHub
			origGetBySKU := models.GetInventoryBySKU
			origGetBySKUHub := models.GetInventoryBySKUAndHub
			origGetAll := models.GetAllInventory

			// Reassign to mock functions
			models.UpsertInventory = mockmodel.UpsertInventory
			models.CheckInventoryStatus = mockmodel.CheckInventoryStatus
			models.ValidateOrderByHubAndSKU = mockmodel.ValidateOrderByHubAndSKU
			models.GetInventoryByHub = mockmodel.GetInventoryByHub
			models.GetInventoryBySKU = mockmodel.GetInventoryBySKU
			models.GetInventoryBySKUAndHub = mockmodel.GetInventoryBySKUAndHub
			models.GetAllInventory = mockmodel.GetAllInventory

			// Restore after tests
			defer func() {
				models.UpsertInventory = origUpsert
				models.CheckInventoryStatus = origCheck
				models.ValidateOrderByHubAndSKU = origValidate
				models.GetInventoryByHub = origGetByHub
				models.GetInventoryBySKU = origGetBySKU
				models.GetInventoryBySKUAndHub = origGetBySKUHub
				models.GetAllInventory = origGetAll
			}()

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			var req *http.Request
			if strBody, ok := tc.body.(string); ok {
				req = httptest.NewRequest(tc.method, tc.url, bytes.NewBufferString(strBody))
				req.Header.Set("Content-Type", "application/json")
			} else if tc.body != nil {
				jsonBody, _ := json.Marshal(tc.body)
				req = httptest.NewRequest(tc.method, tc.url, bytes.NewReader(jsonBody))
				req.Header.Set("Content-Type", "application/json")
			} else {
				req = httptest.NewRequest(tc.method, tc.url, nil)
			}

			q := req.URL.Query()
			for k, v := range tc.queryParams {
				q.Set(k, v)
			}
			req.URL.RawQuery = q.Encode()

			c.Request = req

			tc.handler(c)
			assert.Equal(t, tc.expectedCode, w.Code)
			mockmodel.AssertExpectations(t)
		})
	}
}
