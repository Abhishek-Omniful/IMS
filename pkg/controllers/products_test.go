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

type MockProductModel struct {
	mock.Mock
}

func (m *MockProductModel) GetProducts() (*[]models.Product, error) {
	args := m.Called()
	if prod := args.Get(0); prod != nil {
		return prod.(*[]models.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockProductModel) CreateProduct(p *models.Product) (*models.Product, error) {
	args := m.Called(p)
	if prod := args.Get(0); prod != nil {
		return prod.(*models.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockProductModel) UpdateProduct(p *models.Product) (*models.Product, error) {
	args := m.Called(p)
	if prod := args.Get(0); prod != nil {
		return prod.(*models.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockProductModel) DeleteProduct(id int64) (*models.Product, error) {
	args := m.Called(id)
	if prod := args.Get(0); prod != nil {
		return prod.(*models.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestProductHandlers(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		url          string
		body         interface{}
		mockSetup    func(*MockProductModel)
		handler      gin.HandlerFunc
		expectedCode int
	}{
		{
			name:   "GetProducts - Success",
			method: http.MethodGet,
			url:    "/products",
			mockSetup: func(m *MockProductModel) {
				m.On("GetProducts").Return(&[]models.Product{{ID: 1, ProductName: "Pen"}}, nil)
			},
			handler:      GetProducts,
			expectedCode: 200,
		},
		{
			name:   "GetProducts - Error",
			method: http.MethodGet,
			url:    "/products",
			mockSetup: func(m *MockProductModel) {
				m.On("GetProducts").Return(nil, errors.New("db error"))
			},
			handler:      GetProducts,
			expectedCode: 400,
		},
		{
			name:   "CreateProduct - Success",
			method: http.MethodPost,
			url:    "/products",
			body:   &models.Product{ProductName: "Pen", SellerId: 1, GeneralDescription: "Blue Ink"},
			mockSetup: func(m *MockProductModel) {
				m.On("CreateProduct", &models.Product{ProductName: "Pen", SellerId: 1, GeneralDescription: "Blue Ink"}).Return(&models.Product{ProductName: "Pen", SellerId: 1, GeneralDescription: "Blue Ink"}, nil)
			},
			handler:      CreateProduct,
			expectedCode: 201,
		},
		{
			name:         "CreateProduct - Invalid JSON",
			method:       http.MethodPost,
			url:          "/products",
			body:         "{invalid-json}",
			mockSetup:    func(m *MockProductModel) {},
			handler:      CreateProduct,
			expectedCode: 400,
		},
		{
			name:   "CreateProduct Internal Error",
			method: http.MethodPost,
			url:    "/products",
			body: &models.Product{
				ProductName:        "Pen",
				SellerId:           1,
				GeneralDescription: "Blue Ink",
			},
			mockSetup: func(m *MockProductModel) {
				m.On("CreateProduct", &models.Product{ProductName: "Pen", SellerId: 1, GeneralDescription: "Blue Ink"}).Return(nil, errors.New("create error"))
			},
			handler:      CreateProduct,
			expectedCode: 400,
		},
		{
			name:   "DeleteProduct - Success",
			method: http.MethodDelete,
			url:    "/products/1",
			mockSetup: func(m *MockProductModel) {
				m.On("DeleteProduct", int64(1)).Return(&models.Product{
					ProductName:        "Pen",
					SellerId:           1,
					GeneralDescription: "Blue Ink",
				}, nil)
			},
			handler:      DeleteProduct,
			expectedCode: 200,
		},
		{
			name:         "DeleteProduct - Invalid ID",
			method:       http.MethodDelete,
			url:          "/products/invalid",
			mockSetup:    func(m *MockProductModel) {},
			handler:      DeleteProduct,
			expectedCode: 400,
		},
		{
			name:         "UpdateProduct - Invalid ID",
			method:       http.MethodPut,
			url:          "/products/invalid",
			mockSetup:    func(m *MockProductModel) {},
			handler:      UpdateProduct,
			expectedCode: 400,
		},
		{
			name:         "UpdateProduct - Invalid JSON",
			method:       http.MethodPut,
			url:          "/products/1",
			body:         "{invalid-json}",
			mockSetup:    func(m *MockProductModel) {},
			handler:      UpdateProduct,
			expectedCode: 400,
		},
		{
			name:   "UpdateProduct - Internal Error",
			method: http.MethodPut,
			url:    "/products/1",
			body: &models.Product{
				ProductName:        "Widget",
				SellerId:           42,
				GeneralDescription: "Useful gadget",
			},
			mockSetup: func(m *MockProductModel) {
				m.On("UpdateProduct", &models.Product{
					ID:                 1,
					ProductName:        "Widget",
					SellerId:           42,
					GeneralDescription: "Useful gadget",
				}).Return(nil, errors.New("update error"))
			},
			handler:      UpdateProduct,
			expectedCode: 400,
		},
		{
			name:   "UpdateProduct - Success",
			method: http.MethodPut,
			url:    "/products/1",
			body: &models.Product{
				ProductName:        "Widget",
				SellerId:           42,
				GeneralDescription: "Useful gadget",
			},
			mockSetup: func(m *MockProductModel) {
				m.On("UpdateProduct", &models.Product{
					ID:                 1,
					ProductName:        "Widget",
					SellerId:           42,
					GeneralDescription: "Useful gadget",
				}).Return(&models.Product{
					ID:                 1,
					ProductName:        "Widget",
					SellerId:           42,
					GeneralDescription: "Useful gadget",
				}, nil)
			},
			handler:      UpdateProduct,
			expectedCode: 200,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockModel := new(MockProductModel)
			if tc.mockSetup != nil {
				tc.mockSetup(mockModel)
			}

			// Backup
			origGet := models.GetProducts
			origCreate := models.CreateProduct
			origUpdate := models.UpdateProduct
			origDelete := models.DeleteProduct

			// Restore
			defer func() {
				models.GetProducts = origGet
				models.CreateProduct = origCreate
				models.UpdateProduct = origUpdate
				models.DeleteProduct = origDelete
			}()

			// Replace
			models.GetProducts = mockModel.GetProducts
			models.CreateProduct = mockModel.CreateProduct
			models.UpdateProduct = mockModel.UpdateProduct
			models.DeleteProduct = mockModel.DeleteProduct

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
			if strings.Contains(tc.url, "/products/") {
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
