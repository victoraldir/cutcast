package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victoraldir/cutcast/internal/app/record/usecases"
)

type RecordController struct {
	createRecordGroupUseCase usecases.CreateRecordGroupUseCase
	finishRecordGroupUseCase usecases.FinishRecordGroupUseCase
	listRecordGroupUseCase   usecases.ListRecordGroupUseCase
}

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

func NewRecordController(
	createRecordGroupUseCase usecases.CreateRecordGroupUseCase,
	finishRecordGroupUseCase usecases.FinishRecordGroupUseCase,
	listRecordGroupUseCase usecases.ListRecordGroupUseCase) *RecordController {
	return &RecordController{
		createRecordGroupUseCase: createRecordGroupUseCase,
		finishRecordGroupUseCase: finishRecordGroupUseCase,
		listRecordGroupUseCase:   listRecordGroupUseCase,
	}
}

func (rc *RecordController) Create(ctx *gin.Context) {

	var command usecases.RecordGroupCommand

	if err := ctx.ShouldBindJSON(&command); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := rc.createRecordGroupUseCase.Execute(command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)

}

func (rc *RecordController) Finish(ctx *gin.Context) {

	id := ctx.Param("id")

	err := rc.finishRecordGroupUseCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Record finished"})

}

func (rc *RecordController) List(ctx *gin.Context) {

	records, err := rc.listRecordGroupUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, records)

}

func (tc *TrimController) List(ctx *gin.Context) {

	recordId := ctx.Param("id")

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
