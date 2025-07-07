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

type MockSellerModel struct {
	mock.Mock
}

func (m *MockSellerModel) GetSellers() (*[]models.Seller, error) {
	args := m.Called()
	if sellers := args.Get(0); sellers != nil {
		return sellers.(*[]models.Seller), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockSellerModel) CreateSeller(s *models.Seller) (*models.Seller, error) {
	args := m.Called(s)
	if seller := args.Get(0); seller != nil {
		return seller.(*models.Seller), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockSellerModel) UpdateSeller(s *models.Seller) (*models.Seller, error) {
	args := m.Called(s)
	if seller := args.Get(0); seller != nil {
		return seller.(*models.Seller), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockSellerModel) DeleteSeller(id int64) (*models.Seller, error) {
	args := m.Called(id)
	if seller := args.Get(0); seller != nil {
		return seller.(*models.Seller), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestSellerHandlers(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		url          string
		body         interface{}
		mockSetup    func(*MockSellerModel)
		handler      gin.HandlerFunc
		expectedCode int
	}{
		{
			name:   "GetSellers - Success",
			method: http.MethodGet,
			url:    "/sellers",
			mockSetup: func(m *MockSellerModel) {
				m.On("GetSellers").Return(&[]models.Seller{{ID: 1, SellerName: "John"}}, nil)
			},
			handler:      GetSellers,
			expectedCode: 200,
		},
		{
			name:   "GetSellers - Error",
			method: http.MethodGet,
			url:    "/sellers",
			mockSetup: func(m *MockSellerModel) {
				m.On("GetSellers").Return(nil, errors.New("db error"))
			},
			handler:      GetSellers,
			expectedCode: 400,
		},
		{
			name:   "CreateSeller - Success",
			method: http.MethodPost,
			url:    "/sellers",
			body: &models.Seller{
				HubID:         1,
				TenantID:      1,
				SellerName:    "Alice",
				SellerContact: "9876543210",
				SellerEmail:   "alice@example.com",
			},
			mockSetup: func(m *MockSellerModel) {
				m.On("CreateSeller", mock.Anything).Return(&models.Seller{SellerName: "Alice"}, nil)
			},
			handler:      CreateSeller,
			expectedCode: 201,
		},
		{
			name:         "CreateSeller - Invalid JSON",
			method:       http.MethodPost,
			url:          "/sellers",
			body:         "{invalid-json}",
			mockSetup:    func(m *MockSellerModel) {},
			handler:      CreateSeller,
			expectedCode: 400,
		},
		{
			name:   "CreateSeller - Internal Error",
			method: http.MethodPost,
			url:    "/sellers",
			body: &models.Seller{
				HubID:         1,
				TenantID:      1,
				SellerName:    "Alice",
				SellerContact: "9876543210",
				SellerEmail:   "alice@example.com",
			},
			mockSetup: func(m *MockSellerModel) {
				m.On("CreateSeller", mock.Anything).Return(nil, errors.New("create error"))
			},
			handler:      CreateSeller,
			expectedCode: 400,
		},
		{
			name:         "UpdateSeller - Invalid ID",
			method:       http.MethodPut,
			url:          "/sellers/invalid",
			mockSetup:    func(m *MockSellerModel) {},
			handler:      UpdateSeller,
			expectedCode: 400,
		},
		{
			name:         "UpdateSeller - Invalid JSON",
			method:       http.MethodPut,
			url:          "/sellers/1",
			body:         "{invalid-json}",
			mockSetup:    func(m *MockSellerModel) {},
			handler:      UpdateSeller,
			expectedCode: 400,
		},
		{
			name:   "UpdateSeller - Internal Error",
			method: http.MethodPut,
			url:    "/sellers/1",
			body: &models.Seller{
				SellerName:    "Updated",
				SellerContact: "1234567890",
				SellerEmail:   "updated@example.com",
			},
			mockSetup: func(m *MockSellerModel) {
				m.On("UpdateSeller", mock.Anything).Return(nil, errors.New("update error"))
			},
			handler:      UpdateSeller,
			expectedCode: 400,
		},
		{
			name:   "UpdateSeller - Success",
			method: http.MethodPut,
			url:    "/sellers/1",
			body: &models.Seller{
				SellerName:    "Updated",
				SellerContact: "1234567890",
				SellerEmail:   "updated@example.com",
			},
			mockSetup: func(m *MockSellerModel) {
				m.On("UpdateSeller", mock.Anything).Return(&models.Seller{ID: 1, SellerName: "Updated"}, nil)
			},
			handler:      UpdateSeller,
			expectedCode: 200,
		},
		{
			name:         "DeleteSeller - Invalid ID",
			method:       http.MethodDelete,
			url:          "/sellers/invalid",
			mockSetup:    func(m *MockSellerModel) {},
			handler:      DeleteSeller,
			expectedCode: 400,
		},
		{
			name:   "DeleteSeller - Error",
			method: http.MethodDelete,
			url:    "/sellers/1",
			mockSetup: func(m *MockSellerModel) {
				m.On("DeleteSeller", int64(1)).Return(nil, errors.New("delete error"))
			},
			handler:      DeleteSeller,
			expectedCode: 400,
		},
		{
			name:   "DeleteSeller - Success",
			method: http.MethodDelete,
			url:    "/sellers/1",
			mockSetup: func(m *MockSellerModel) {
				m.On("DeleteSeller", int64(1)).Return(&models.Seller{ID: 1, SellerName: "Deleted"}, nil)
			},
			handler:      DeleteSeller,
			expectedCode: 200,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockModel := new(MockSellerModel)
			if tc.mockSetup != nil {
				tc.mockSetup(mockModel)
			}

			// Backup
			origGet := models.GetSellers
			origCreate := models.CreateSeller
			origUpdate := models.UpdateSeller
			origDelete := models.DeleteSeller

			// Restore
			defer func() {
				models.GetSellers = origGet
				models.CreateSeller = origCreate
				models.UpdateSeller = origUpdate
				models.DeleteSeller = origDelete
			}()

			// Replace
			models.GetSellers = mockModel.GetSellers
			models.CreateSeller = mockModel.CreateSeller
			models.UpdateSeller = mockModel.UpdateSeller
			models.DeleteSeller = mockModel.DeleteSeller

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
			if strings.Contains(tc.url, "/sellers/") {
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
