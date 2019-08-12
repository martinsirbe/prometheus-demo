package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/martinsirbe/prometheus-demo/internal/mocks"
	"github.com/martinsirbe/prometheus-demo/internal/service"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type testSuite struct {
	mockController *gomock.Controller
	mockedStorage  *mocks.MockStorage

	ll *service.LogicLayer
}

func TestSuccessfullyCalledInsert(t *testing.T) {
	ts := setupTest(t)

	expectedErr := errors.New("run once insert")

	ts.mockedStorage.EXPECT().Insert(gomock.Any()).Times(1).Return(expectedErr)
	ts.mockedStorage.EXPECT().Delete(gomock.Any()).Return(nil)

	actualErr := ts.ll.Run()

	assert.NotNil(t, actualErr)
	assert.Equal(t, expectedErr, actualErr)
}

func TestSuccessfullyCalledDelete(t *testing.T) {
	ts := setupTest(t)

	ts.mockedStorage.EXPECT().Insert(gomock.Any()).Return(nil)

	expectedErr := errors.New("run once delete")
	ts.mockedStorage.EXPECT().Delete(gomock.Any()).Times(1).Return(expectedErr)

	actualErr := ts.ll.Run()

	assert.NotNil(t, actualErr)
	assert.Equal(t, expectedErr, actualErr)
}

func setupTest(t *testing.T) *testSuite {
	mockController := gomock.NewController(t)
	mockedStorage := mocks.NewMockStorage(mockController)

	ll := service.NewLogicLayer(mockedStorage)

	return &testSuite{
		mockController: mockController,
		mockedStorage:  mockedStorage,
		ll:             ll,
	}
}
