package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Abhishek-Omniful/IMS/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSKUModel struct {
	mock.Mock
}

func (m *MockSKUModel) GetSKUs() (*[]models.SKU, error) {
	args := m.Called()
	if v := args.Get(0); v != nil {
		return v.(*[]models.SKU), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockSKUModel) CreateSKU(sku *models.SKU) (*models.SKU, error) {
	args := m.Called(sku)
	if v := args.Get(0); v != nil {
		return v.(*models.SKU), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockSKUModel) UpdateSKU(sku *models.SKU) (*models.SKU, error) {
	args := m.Called(sku)
	if v := args.Get(0); v != nil {
		return v.(*models.SKU), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockSKUModel) DeleteSKU(id int64) (*models.SKU, error) {
	args := m.Called(id)
	if v := args.Get(0); v != nil {
		return v.(*models.SKU), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestSKUHandlers(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		url          string
		body         interface{}
		mockSetup    func(*MockSKUModel)
		handler      gin.HandlerFunc
		expectedCode int
	}{
		{
			name:   "GetSKUs - Success",
			method: http.MethodGet,
			url:    "/skus",
			mockSetup: func(m *MockSKUModel) {
				m.On("GetSKUs").Return(&[]models.SKU{{ID: 1, Description: "Basic item"}}, nil)
			},
			handler:      GetSKUs,
			expectedCode: 200,
		},
		{
			name:   "GetSKUs - Error",
			method: http.MethodGet,
			url:    "/skus",
			mockSetup: func(m *MockSKUModel) {
				m.On("GetSKUs").Return(nil, errors.New("fetch error"))
			},
			handler:      GetSKUs,
			expectedCode: 400,
		},
		{
			name:   "CreateSKU - Success",
			method: http.MethodPost,
			url:    "/skus",
			body:   &models.SKU{SellerID: 1, ProductID: 2, Description: "Test SKU"},
			mockSetup: func(m *MockSKUModel) {
				m.On("CreateSKU", &models.SKU{SellerID: 1, ProductID: 2, Description: "Test SKU"}).Return(&models.SKU{ID: 1, SellerID: 1, ProductID: 2, Description: "Test SKU"}, nil)
			},
			handler:      CreateSKU,
			expectedCode: 201,
		},
		{
			name:         "CreateSKU - Invalid JSON",
			method:       http.MethodPost,
			url:          "/skus",
			body:         "{invalid-json}",
			mockSetup:    func(m *MockSKUModel) {},
			handler:      CreateSKU,
			expectedCode: 400,
		},
		{
			name:   "CreateSKU - Internal Error",
			method: http.MethodPost,
			url:    "/skus",
			body: &models.SKU{
				SellerID:    1,
				ProductID:   2,
				Description: "Test SKU",
			},
			mockSetup: func(m *MockSKUModel) {
				m.On("CreateSKU", &models.SKU{
					SellerID:    1,
					ProductID:   2,
					Description: "Test SKU",
				}).Return(nil, errors.New("create error"))
			},
			handler:      CreateSKU,
			expectedCode: 400,
		},
		{
			name:   "UpdateSKU - Success",
			method: http.MethodPut,
			url:    "/skus/1",
			body:   &models.SKU{SellerID: 1, ProductID: 2, Description: "Updated"},
			mockSetup: func(m *MockSKUModel) {
				m.On("UpdateSKU", &models.SKU{ID: 1, SellerID: 1, ProductID: 2, Description: "Updated"}).Return(&models.SKU{ID: 1, SellerID: 1, ProductID: 2, Description: "Updated"}, nil)
			},
			handler:      UpdateSKU,
			expectedCode: 200,
		},
		{
			name:         "UpdateSKU - Invalid ID",
			method:       http.MethodPut,
			url:          "/skus/invalid",
			mockSetup:    func(m *MockSKUModel) {},
			handler:      UpdateSKU,
			expectedCode: 400,
		},
		{
			name:         "UpdateSKU - Invalid JSON",
			method:       http.MethodPut,
			url:          "/skus/1",
			body:         "{invalid-json}",
			mockSetup:    func(m *MockSKUModel) {},
			handler:      UpdateSKU,
			expectedCode: 400,
		},
		{
			name:   "UpdateSKU - Internal Error",
			method: http.MethodPut,
			url:    "/skus/1",
			body: &models.SKU{
				ID:          1,
				SellerID:    42,
				ProductID:   2,
				Description: "Useful gadget",
			},
			mockSetup: func(m *MockSKUModel) {
				m.On("UpdateSKU", &models.SKU{
					ID:          1,
					SellerID:    42,
					ProductID:   2,
					Description: "Useful gadget",
				}).Return(nil, errors.New("update error"))
			},
			handler:      UpdateSKU,
			expectedCode: 400,
		},
		{
			name:   "DeleteSKU - Success",
			method: http.MethodDelete,
			url:    "/skus/1",
			mockSetup: func(m *MockSKUModel) {
				m.On("DeleteSKU", int64(1)).Return(&models.SKU{
					ID:          1,
					SellerID:    42,
					ProductID:   2,
					Description: "Useful gadget",
				}, nil)
			},
			handler:      DeleteSKU,
			expectedCode: 200,
		},
		{
			name:         "DeleteSKU - Invalid ID",
			method:       http.MethodDelete,
			url:          "/skus/invalid",
			mockSetup:    func(m *MockSKUModel) {},
			handler:      DeleteSKU,
			expectedCode: 400,
		},
		{
			name:   "DeleteSKU - Error",
			method: http.MethodDelete,
			url:    "/skus/1",
			mockSetup: func(m *MockSKUModel) {
				m.On("DeleteSKU", int64(1)).Return(nil, errors.New("delete error"))
			},
			handler:      DeleteSKU,
			expectedCode: 400,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockModel := new(MockSKUModel)
			if tc.mockSetup != nil {
				tc.mockSetup(mockModel)
			}

			origGet := models.GetSKUs
			origCreate := models.CreateSKU
			origUpdate := models.UpdateSKU
			origDelete := models.DeleteSKU

			defer func() {
				models.GetSKUs = origGet
				models.CreateSKU = origCreate
				models.UpdateSKU = origUpdate
				models.DeleteSKU = origDelete
			}()

			models.GetSKUs = mockModel.GetSKUs
			models.CreateSKU = mockModel.CreateSKU
			models.UpdateSKU = mockModel.UpdateSKU
			models.DeleteSKU = mockModel.DeleteSKU

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			var req *http.Request
			if str, ok := tc.body.(string); ok {
				req = httptest.NewRequest(tc.method, tc.url, bytes.NewBufferString(str))
				req.Header.Set("Content-Type", "application/json")
			} else if tc.body != nil {
				b, _ := json.Marshal(tc.body)
				req = httptest.NewRequest(tc.method, tc.url, bytes.NewReader(b))
				req.Header.Set("Content-Type", "application/json")
			} else {
				req = httptest.NewRequest(tc.method, tc.url, nil)
			}

			c.Request = req
			if strings.Contains(tc.url, "/skus/") {
				parts := strings.Split(tc.url, "/")
				id := parts[len(parts)-1]
				c.Params = gin.Params{{Key: "id", Value: id}}
			}

			tc.handler(c)
			assert.Equal(t, tc.expectedCode, w.Code)
			mockModel.AssertExpectations(t)
		})
	}
}
