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

type MockTenantModel struct {
	mock.Mock
}

func (m *MockTenantModel) GetTenants() (*[]models.Tenant, error) {
	args := m.Called()
	if tmp := args.Get(0); tmp != nil {
		return tmp.(*[]models.Tenant), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTenantModel) CreateTenant(tenant *models.Tenant) (*models.Tenant, error) {
	args := m.Called(tenant)
	if tmp := args.Get(0); tmp != nil {
		return tmp.(*models.Tenant), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTenantModel) UpdateTenant(tenant *models.Tenant) (*models.Tenant, error) {
	args := m.Called(tenant)
	if tmp := args.Get(0); tmp != nil {
		return tmp.(*models.Tenant), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTenantModel) DeleteTenant(id int64) (*models.Tenant, error) {
	args := m.Called(id)
	if tmp := args.Get(0); tmp != nil {
		return tmp.(*models.Tenant), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestTenantHandlers(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		url          string
		body         interface{}
		mockSetup    func(*MockTenantModel)
		handler      gin.HandlerFunc
		expectedCode int
	}{
		{
			name:   "GetTenants - Success",
			method: http.MethodGet,
			url:    "/tenants",
			mockSetup: func(m *MockTenantModel) {
				m.On("GetTenants").Return(&[]models.Tenant{{ID: 1}}, nil)
			},
			handler:      GetTenants,
			expectedCode: 200,
		},
		{
			name:   "GetTenants - Failure",
			method: http.MethodGet,
			url:    "/tenants",
			mockSetup: func(m *MockTenantModel) {
				m.On("GetTenants").Return(nil, errors.New("fetch error"))
			},
			handler:      GetTenants,
			expectedCode: 400,
		},
		{
			name:   "CreateTenant - Success",
			method: http.MethodPost,
			url:    "/tenants",
			body: models.Tenant{
				TenantName:        "ABC Corp",
				RegisteredAddress: "123 Street",
				TenantContact:     "9876543210",
				TenantEmail:       "abc@corp.com",
			},
			mockSetup: func(m *MockTenantModel) {
				m.On("CreateTenant", &models.Tenant{
					TenantName:        "ABC Corp",
					RegisteredAddress: "123 Street",
					TenantContact:     "9876543210",
					TenantEmail:       "abc@corp.com",
				}).Return(&models.Tenant{ID: 1}, nil)
			},
			handler:      CreateTenant,
			expectedCode: 201,
		},
		{
			name:         "CreateTenant - Invalid JSON",
			method:       http.MethodPost,
			url:          "/tenants",
			body:         "{invalid-json}",
			mockSetup:    func(m *MockTenantModel) {},
			handler:      CreateTenant,
			expectedCode: 400,
		},
		{
			name:   "CreateTenant - Internal Error",
			method: http.MethodPost,
			url:    "/tenants",
			body: &models.Tenant{
				TenantName:        "ABC Corp",
				RegisteredAddress: "123 Street",
				TenantContact:     "9876543210",
				TenantEmail:       "abc@corp.com",
			},
			mockSetup: func(m *MockTenantModel) {
				m.On("CreateTenant", &models.Tenant{
					TenantName:        "ABC Corp",
					RegisteredAddress: "123 Street",
					TenantContact:     "9876543210",
					TenantEmail:       "abc@corp.com",
				}).Return(nil, errors.New("create error"))
			},
			handler:      CreateTenant,
			expectedCode: 400,
		},
		{
			name:         "UpdateTenant - Invalid ID",
			method:       http.MethodPut,
			url:          "/tenants/invalid",
			mockSetup:    func(m *MockTenantModel) {},
			handler:      UpdateTenant,
			expectedCode: 400,
		},
		{
			name:         "UpdateTenant - Invalid JSON",
			method:       http.MethodPut,
			url:          "/tenants/1",
			body:         "{invalid-json}",
			mockSetup:    func(m *MockTenantModel) {},
			handler:      UpdateTenant,
			expectedCode: 400,
		},
		{
			name:   "UpdateTenant - Success",
			method: http.MethodPut,
			url:    "/tenants/1",
			body: &models.Tenant{
				TenantName:        "ABC Corp",
				RegisteredAddress: "123 Street",
				TenantContact:     "9876543210",
				TenantEmail:       "abc@corp.com",
			},
			mockSetup: func(m *MockTenantModel) {
				m.On("UpdateTenant", &models.Tenant{
					ID:                1,
					TenantName:        "ABC Corp",
					RegisteredAddress: "123 Street",
					TenantContact:     "9876543210",
					TenantEmail:       "abc@corp.com",
				}).Return(&models.Tenant{ID: 1}, nil)
			},
			handler:      UpdateTenant,
			expectedCode: 200,
		},
		{
			name:   "UpdateTenant - Error",
			method: http.MethodPut,
			url:    "/tenants/1",
			body: &models.Tenant{
				TenantName:        "ABC Corp",
				RegisteredAddress: "123 Street",
				TenantContact:     "9876543210",
				TenantEmail:       "abc@corp.com",
			},
			mockSetup: func(m *MockTenantModel) {
				m.On("UpdateTenant", &models.Tenant{
					ID:                1,
					TenantName:        "ABC Corp",
					RegisteredAddress: "123 Street",
					TenantContact:     "9876543210",
					TenantEmail:       "abc@corp.com",
				}).Return(nil, errors.New("update error"))
			},
			handler:      UpdateTenant,
			expectedCode: 400,
		},
		{
			name:         "DeleteTenant - Invalid ID",
			method:       http.MethodDelete,
			url:          "/tenants/invalid",
			mockSetup:    func(m *MockTenantModel) {},
			handler:      DeleteTenant,
			expectedCode: 400,
		},
		{
			name:   "DeleteTenant - Success",
			method: http.MethodDelete,
			url:    "/tenants/1",
			mockSetup: func(m *MockTenantModel) {
				m.On("DeleteTenant", int64(1)).Return(&models.Tenant{ID: 1}, nil)
			},
			handler:      DeleteTenant,
			expectedCode: 200,
		},
		{
			name:   "DeleteTenant - Error",
			method: http.MethodDelete,
			url:    "/tenants/1",
			mockSetup: func(m *MockTenantModel) {
				m.On("DeleteTenant", int64(1)).Return(nil, errors.New("delete error"))
			},
			handler:      DeleteTenant,
			expectedCode: 400,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockModel := new(MockTenantModel)
			if tc.mockSetup != nil {
				tc.mockSetup(mockModel)
			}

			// Swap actual with mock
			origGet := models.GetTenants
			origCreate := models.CreateTenant
			origUpdate := models.UpdateTenant
			origDelete := models.DeleteTenant

			defer func() {
				models.GetTenants = origGet
				models.CreateTenant = origCreate
				models.UpdateTenant = origUpdate
				models.DeleteTenant = origDelete
			}()

			models.GetTenants = mockModel.GetTenants
			models.CreateTenant = mockModel.CreateTenant
			models.UpdateTenant = mockModel.UpdateTenant
			models.DeleteTenant = mockModel.DeleteTenant

			// Prepare request
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			var req *http.Request
			if strBody, ok := tc.body.(string); ok {
				req = httptest.NewRequest(tc.method, tc.url, bytes.NewBufferString(strBody))
			} else if tc.body != nil {
				jsonBody, _ := json.Marshal(tc.body)
				req = httptest.NewRequest(tc.method, tc.url, bytes.NewReader(jsonBody))
			} else {
				req = httptest.NewRequest(tc.method, tc.url, nil)
			}
			req.Header.Set("Content-Type", "application/json")

			c.Request = req
			if strings.Contains(tc.url, "/tenants/") {
				parts := strings.Split(tc.url, "/")
				c.Params = gin.Params{{Key: "id", Value: parts[len(parts)-1]}}
			}

			tc.handler(c)
			assert.Equal(t, tc.expectedCode, w.Code)
			mockModel.AssertExpectations(t)
		})
	}
}
