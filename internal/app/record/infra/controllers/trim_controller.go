package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victoraldir/cutcast/internal/app/record/usecases"
)

type TrimController struct {
	trimRecordGroupUseCase     usecases.TrimRecordGroupUseCase
	listTrimRecordGroupUseCase usecases.ListTrimRecordGroupUseCase
}

func NewTrimController(listTrimRecordGroupUseCase usecases.ListTrimRecordGroupUseCase,
	trimRecordGroupUseCase usecases.TrimRecordGroupUseCase) *TrimController {
	return &TrimController{
		listTrimRecordGroupUseCase: listTrimRecordGroupUseCase,
		trimRecordGroupUseCase:     trimRecordGroupUseCase,
	}
}

func (tc *TrimController) List(ctx *gin.Context) {

	recordId := ctx.Param("id")

	if recordId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record id is required"})
		return
	}

	records, err := tc.listTrimRecordGroupUseCase.Execute(recordId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, records)

}

func (tc *TrimController) Create(ctx *gin.Context) {

	recordId := ctx.Param("id")

	var command usecases.TrimRecordGroupCommand

	if err := ctx.ShouldBindJSON(&command); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	command.RecordId = recordId

	response, err := tc.trimRecordGroupUseCase.Execute(command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)

}
