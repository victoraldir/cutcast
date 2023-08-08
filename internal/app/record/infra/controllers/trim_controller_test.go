package controllers

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	usecases "github.com/victoraldir/cutcast/internal/app/record/usecases"
	usecases_mock "github.com/victoraldir/cutcast/internal/app/record/usecases/mocks"
	"go.uber.org/mock/gomock"
)

var trimRecordGroupUseCaseMock *usecases_mock.MockTrimRecordGroupUseCase
var listTrimRecordGroupUseCaseMock *usecases_mock.MockListTrimRecordGroupUseCase

func setup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trimRecordGroupUseCaseMock = usecases_mock.NewMockTrimRecordGroupUseCase(ctrl)
	listTrimRecordGroupUseCaseMock = usecases_mock.NewMockListTrimRecordGroupUseCase(ctrl)

}

func TestTrimController_List(t *testing.T) {

	setup(t)

	t.Run("should return 200 when list trim record group", func(t *testing.T) {
		//Arrange
		expectedRecordId := "1"

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{
			gin.Param{
				Key:   "id",
				Value: expectedRecordId,
			},
		}

		trimController := NewTrimController(listTrimRecordGroupUseCaseMock, trimRecordGroupUseCaseMock)

		// Act
		listTrimRecordGroupUseCaseMock.EXPECT().Execute(expectedRecordId).Return([]usecases.TrimRecordGroupResponse{{
			StartTime: "00:00:00",
			EndTime:   "00:00:10",
			RecordId:  expectedRecordId,
		}}, nil)
		trimController.List(c)

		// Assert
		assert.Equal(t, 200, w.Code)
	})

	t.Run("should return 400. Missing recordId", func(t *testing.T) {
		//Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		trimController := NewTrimController(listTrimRecordGroupUseCaseMock, trimRecordGroupUseCaseMock)

		// Act
		trimController.List(c)

		// Assert
		assert.Equal(t, 400, w.Code)
	})

	t.Run("should return 500. Usecase execution error", func(t *testing.T) {
		//Arrange
		expectedRecordId := "1"

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{
			gin.Param{
				Key:   "id",
				Value: expectedRecordId,
			},
		}

		trimController := NewTrimController(listTrimRecordGroupUseCaseMock, trimRecordGroupUseCaseMock)

		// Act
		listTrimRecordGroupUseCaseMock.EXPECT().Execute(expectedRecordId).Return(nil, assert.AnError)
		trimController.List(c)

		// Assert
		assert.Equal(t, 500, w.Code)
	})

}

func TestTrimController_Create(t *testing.T) {

	setup(t)

	t.Run("should return 201 when trim record group", func(t *testing.T) {
		//Arrange
		expectedRecordId := "1"
		expectedStartTime := "00:00:00"
		expectedEndTime := "00:00:10"

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{
			gin.Param{
				Key:   "id",
				Value: expectedRecordId,
			},
		}

		c.Request = httptest.NewRequest("POST", "/trim", bytes.NewBuffer([]byte(`{"start_time":"`+expectedStartTime+`","end_time":"`+expectedEndTime+`"}`)))

		trimController := NewTrimController(listTrimRecordGroupUseCaseMock, trimRecordGroupUseCaseMock)

		// Act
		trimRecordGroupUseCaseMock.EXPECT().Execute(usecases.TrimRecordGroupCommand{
			RecordId:  expectedRecordId,
			StartTime: expectedStartTime,
			EndTime:   expectedEndTime,
		}).Return(&usecases.TrimRecordGroupResponse{
			StartTime: expectedStartTime,
			EndTime:   expectedEndTime,
			RecordId:  expectedRecordId,
		}, nil)
		trimController.Create(c)

		// Assert
		assert.Equal(t, 201, w.Code)
	})

	t.Run("should return 400. Missing recordId", func(t *testing.T) {
		//Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		trimController := NewTrimController(listTrimRecordGroupUseCaseMock, trimRecordGroupUseCaseMock)

		// Act
		trimController.Create(c)

		// Assert
		assert.Equal(t, 400, w.Code)
	})

	t.Run("should return 400. Missing startTime", func(t *testing.T) {
		//Arrange
		expectedRecordId := "1"

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{
			gin.Param{
				Key:   "id",
				Value: expectedRecordId,
			},
		}

		trimController := NewTrimController(listTrimRecordGroupUseCaseMock, trimRecordGroupUseCaseMock)

		// Act
		trimController.Create(c)

		// Assert
		assert.Equal(t, 400, w.Code)
	})

	t.Run("should return 400. Missing endTime", func(t *testing.T) {
		//Arrange
		expectedRecordId := "1"

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{
			gin.Param{
				Key:   "id",
				Value: expectedRecordId,
			},
		}

		trimController := NewTrimController(listTrimRecordGroupUseCaseMock, trimRecordGroupUseCaseMock)

		// Act
		trimController.Create(c)

		// Assert
		assert.Equal(t, 400, w.Code)
	})
}
