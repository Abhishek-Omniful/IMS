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

// MockHubModel struct
type MockHubModel struct {
	mock.Mock
}

func (m *MockHubModel) GetHubs() (*[]models.Hub, error) {
	args := m.Called()
	var res *[]models.Hub
	if tmp := args.Get(0); tmp != nil {
		res = tmp.(*[]models.Hub)
	}
	return res, args.Error(1)
}

func (m *MockHubModel) CreateHub(hub *models.Hub) (*models.Hub, error) {
	args := m.Called(hub)
	var res *models.Hub
	if tmp := args.Get(0); tmp != nil {
		res = tmp.(*models.Hub)
	}
	return res, args.Error(1)
}

func (m *MockHubModel) UpdateHub(hub *models.Hub) (*models.Hub, error) {
	args := m.Called(hub)
	var res *models.Hub
	if tmp := args.Get(0); tmp != nil {
		res = tmp.(*models.Hub)
	}
	return res, args.Error(1)
}

func (m *MockHubModel) DeleteHub(id int64) (*models.Hub, error) {
	args := m.Called(id)
	var res *models.Hub
	if tmp := args.Get(0); tmp != nil {
		res = tmp.(*models.Hub)
	}
	return res, args.Error(1)
}

func TestHubHandlers_TableDriven(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		url          string
		body         interface{}
		mockSetup    func(*MockHubModel)
		handler      gin.HandlerFunc
		expectedCode int
	}{
		{
			name:   "GetHubs Success",
			method: http.MethodGet,
			url:    "/hubs",
			mockSetup: func(m *MockHubModel) {
				m.On("GetHubs").Return(&[]models.Hub{{ID: 1}}, nil)
			},
			handler:      GetHubs,
			expectedCode: 200,
		},
		{
			name:   "GetHubs Failure",
			method: http.MethodGet,
			url:    "/hubs",
			mockSetup: func(m *MockHubModel) {
				m.On("GetHubs").Return((*[]models.Hub)(nil), errors.New("error"))
			},
			handler:      GetHubs,
			expectedCode: 400,
		},
		{
			name:   "CreateHub Success",
			method: http.MethodPost,
			url:    "/hubs",
			body: models.Hub{
				TenantID:       1,
				ManagerName:    "Alice",
				ManagerContact: "9876543210",
				ManagerEmail:   "alice@example.com",
			},
			mockSetup: func(m *MockHubModel) {
				m.On("CreateHub", &models.Hub{
					TenantID:       1,
					ManagerName:    "Alice",
					ManagerContact: "9876543210",
					ManagerEmail:   "alice@example.com",
				}).Return(&models.Hub{TenantID: 1,
					ManagerName:    "Alice",
					ManagerContact: "9876543210",
					ManagerEmail:   "alice@example.com"}, nil)
			},
			handler:      CreateHub,
			expectedCode: 201,
		},
		{
			name:         "CreateHub Invalid JSON",
			method:       http.MethodPost,
			url:          "/hubs",
			body:         "{invalid-json}",
			mockSetup:    func(m *MockHubModel) {},
			handler:      CreateHub,
			expectedCode: 400,
		},
		{
			name:   "CreateHub Internal Error",
			method: http.MethodPost,
			url:    "/hubs",
			body: &models.Hub{
				TenantID:       1,
				ManagerName:    "Alice",
				ManagerContact: "9876543210",
				ManagerEmail:   "alice@example.com",
			},
			mockSetup: func(m *MockHubModel) {
				m.On("CreateHub", &models.Hub{
					TenantID:       1,
					ManagerName:    "Alice",
					ManagerContact: "9876543210",
					ManagerEmail:   "alice@example.com",
				}).Return(nil, errors.New("create error"))
			},
			handler:      CreateHub,
			expectedCode: 400,
		},
		{
			name:         "UpdateHub Invalid ID",
			method:       http.MethodPut,
			url:          "/hubs/invalid",
			mockSetup:    func(m *MockHubModel) {},
			handler:      UpdateHub,
			expectedCode: 400,
		},
		{
			name:         "UpdateHub Invalid JSON",
			method:       http.MethodPut,
			url:          "/hubs/1",
			body:         "{invalid-json}",
			mockSetup:    func(m *MockHubModel) {},
			handler:      UpdateHub,
			expectedCode: 400,
		},
		{
			name:   "UpdateHub Internal Error",
			method: http.MethodPut,
			url:    "/hubs/1",
			body: &models.Hub{
				TenantID:       1,
				ManagerName:    "Alice",
				ManagerContact: "9876543210",
				ManagerEmail:   "alice@example.com",
			},
			mockSetup: func(m *MockHubModel) {
				m.On("UpdateHub", &models.Hub{
					ID:             1,
					TenantID:       1,
					ManagerName:    "Alice",
					ManagerContact: "9876543210",
					ManagerEmail:   "alice@example.com",
				}).Return(nil, errors.New("update error"))
			},
			handler:      UpdateHub,
			expectedCode: 400,
		},
		{
			name:   "UpdateHub Success",
			method: http.MethodPut,
			url:    "/hubs/1",
			body: &models.Hub{
				TenantID:       1,
				ManagerName:    "Alice",
				ManagerContact: "9876543210",
				ManagerEmail:   "alice@example.com",
			},
			mockSetup: func(m *MockHubModel) {
				m.On("UpdateHub", &models.Hub{
					ID:             1,
					TenantID:       1,
					ManagerName:    "Alice",
					ManagerContact: "9876543210",
					ManagerEmail:   "alice@example.com",
				}).Return(&models.Hub{
					TenantID:       1,
					ManagerName:    "Alice",
					ManagerContact: "9876543210",
					ManagerEmail:   "alice@example.com",
				}, nil)
			},
			handler:      UpdateHub,
			expectedCode: 200,
		},
		{
			name:         "DeleteHub Invalid ID",
			method:       http.MethodDelete,
			url:          "/hubs/invalid",
			mockSetup:    func(m *MockHubModel) {},
			handler:      DeleteHub,
			expectedCode: 400,
		},
		{
			name:   "DeleteHub Error",
			method: http.MethodDelete,
			url:    "/hubs/1",
			mockSetup: func(m *MockHubModel) {
				m.On("DeleteHub", int64(1)).Return(nil, errors.New("delete error"))
			},
			handler:      DeleteHub,
			expectedCode: 400,
		},
		{
			name:   "DeleteHub Success",
			method: http.MethodDelete,
			url:    "/hubs/1",
			mockSetup: func(m *MockHubModel) {
				m.On("DeleteHub", int64(1)).Return(&models.Hub{
					TenantID:       1,
					ManagerName:    "Alice",
					ManagerContact: "9876543210",
					ManagerEmail:   "alice@example.com",
				}, nil)
			},
			handler:      DeleteHub,
			expectedCode: 200,
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			mockModel := new(MockHubModel)
			if tc.mockSetup != nil {
				tc.mockSetup(mockModel)
			}

			originalGet := models.GetHubs
			originalCreate := models.CreateHub
			originalUpdate := models.UpdateHub
			originalDelete := models.DeleteHub

			defer func() {
				models.GetHubs = originalGet
				models.CreateHub = originalCreate
				models.UpdateHub = originalUpdate
				models.DeleteHub = originalDelete
			}()

			models.GetHubs = mockModel.GetHubs
			models.CreateHub = mockModel.CreateHub
			models.UpdateHub = mockModel.UpdateHub
			models.DeleteHub = mockModel.DeleteHub

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			var req *http.Request
			if strBody, ok := tc.body.(string); ok {
				req = httptest.NewRequest(tc.method, tc.url, bytes.NewBufferString(strBody))
				req.Header.Set("Content-Type", "application/json")
			} else if tc.body != nil {
				jsonBytes, _ := json.Marshal(tc.body)
				req = httptest.NewRequest(tc.method, tc.url, bytes.NewReader(jsonBytes))
				req.Header.Set("Content-Type", "application/json")
			} else {
				req = httptest.NewRequest(tc.method, tc.url, nil)
			}

			c.Request = req
			if tc.url == "/hubs/1" || tc.url == "/hubs/invalid" {
				parts := strings.Split(tc.url, "/")
				idStr := parts[len(parts)-1]
				c.Params = gin.Params{gin.Param{Key: "id", Value: idStr}}
			}

			tc.handler(c)
			assert.Equal(t, tc.expectedCode, w.Code)
			mockModel.AssertExpectations(t)
		})
	}
}
