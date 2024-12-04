package inference_tasks

import (
	"context"
	"crynux_relay/api/v1/response"
	"crynux_relay/blockchain"
	"crynux_relay/config"
	"crynux_relay/models"
	"crynux_relay/utils"
	"errors"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ResultInput struct {
	TaskIDCommitment string                  `path:"task_id_commitment" description:"Task id commitment" validate:"required"`
	Files            []*multipart.FileHeader `form:"files" validate:"required" description:"Result files (PNG images for task of type sd and sd_finetune, JSON files for task of type gpt)"`
	Checkpoint       *multipart.FileHeader   `form:"checkpoint" description:"Result checkpoint file for task of type sd_finetune"`
}

type ResultInputWithSignature struct {
	ResultInput
	Timestamp int64  `form:"timestamp" description:"Signature timestamp" validate:"required"`
	Signature string `form:"signature" description:"Signature" validate:"required"`
}

func UploadResult(c *gin.Context, in *ResultInputWithSignature) (*response.Response, error) {

	match, address, err := ValidateSignature(in.ResultInput, in.Timestamp, in.Signature)

	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	if !match {
		validationErr := response.NewValidationErrorResponse("signature", "Invalid signature")
		return nil, validationErr
	}

	var task models.InferenceTask

	dbCtx, dbCancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer dbCancel()

	if err := config.GetDB().WithContext(dbCtx).Where(&models.InferenceTask{TaskIDCommitment: in.TaskIDCommitment}).First(&task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			validationErr := response.NewValidationErrorResponse("task_id_commitment", "Task not found")
			return nil, validationErr
		} else {
			return nil, response.NewExceptionResponse(err)
		}
	}

	if task.SelectedNode != address {
		return nil, response.NewValidationErrorResponse("Signature", "Signer not allowed")
	}

	taskIDCommitmentBytes, err := utils.HexStrToCommitment(in.TaskIDCommitment)
	if err != nil {
		return nil, response.NewValidationErrorResponse("task_id_commitment", "Invalid task id commitment")
	}

	chainCtx, chainCancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer chainCancel()
	chainTask, err := blockchain.GetTaskByCommitment(chainCtx, *taskIDCommitmentBytes)
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	if task.Status != models.InferenceTaskParamsUploaded ||
		(models.ChainTaskStatus(chainTask.Status) != models.ChainTaskValidated && models.ChainTaskStatus(chainTask.Status) != models.ChainTaskGroupValidated) {
		validationErr := response.NewValidationErrorResponse("task_id_commitment", "Task not validated")
		return nil, validationErr
	}
	// Check whether the images are correct
	var uploadedScoreBytes []byte

	for _, file := range in.Files {

		fileObj, err := file.Open()

		if err != nil {
			return nil, response.NewExceptionResponse(err)
		}

		var hash []byte
		if task.TaskType == models.TaskTypeSD || task.TaskType == models.TaskTypeSDFTLora {
			hash, err = blockchain.GetPHashForImage(fileObj)
		} else {
			hash, err = blockchain.GetHashForGPTResponse(fileObj)
		}

		if err != nil {
			return nil, response.NewExceptionResponse(err)
		}

		uploadedScoreBytes = append(uploadedScoreBytes, hash...)

		err = fileObj.Close()
		if err != nil {
			return nil, response.NewExceptionResponse(err)
		}
	}

	uploadedScore := hexutil.Encode(uploadedScoreBytes)
	chainScore := hexutil.Encode(chainTask.Score)

	log.Debugln("image compare: result from the blockchain: " + chainScore)
	log.Debugln("image compare: result from the uploaded file: " + uploadedScore)

	if chainScore != uploadedScore {
		validationErr := response.NewValidationErrorResponse("files", "Wrong result files uploaded")
		return nil, validationErr
	}

	appConfig := config.GetConfig()

	taskDir := filepath.Join(appConfig.DataDir.InferenceTasks, task.TaskIDCommitment, "results")
	if err = os.MkdirAll(taskDir, 0o711); err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	var fileExt string
	if task.TaskType == models.TaskTypeSD || task.TaskType == models.TaskTypeSDFTLora {
		fileExt = ".png"
	} else {
		fileExt = ".json"
	}

	for i, file := range in.Files {
		filename := filepath.Join(taskDir, strconv.Itoa(i)+fileExt)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			return nil, response.NewExceptionResponse(err)
		}
	}

	// store checkpoint of finetune type task
	if task.TaskType == models.TaskTypeSDFTLora {
		if in.Checkpoint == nil {
			return nil, response.NewValidationErrorResponse("checkpoint", "Checkpoint not uploaded")
		}
		checkpointFilename := filepath.Join(taskDir, "checkpoint.zip")
		if err := c.SaveUploadedFile(in.Checkpoint, checkpointFilename); err != nil {
			return nil, response.NewExceptionResponse(err)
		}
	}

	// Update task status
	task.Status = models.InferenceTaskResultsReady

	dbCtx1, dbCancel1 := context.WithTimeout(c.Request.Context(), time.Second)
	defer dbCancel1()

	if err := config.GetDB().WithContext(dbCtx1).Save(&task).Error; err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	return &response.Response{}, nil
}
