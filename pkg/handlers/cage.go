package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/samsonannan/prizepicks-assessment/pkg/db"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/cage"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/dinosaur"
	"github.com/samsonannan/prizepicks-assessment/pkg/logger"
	"github.com/samsonannan/prizepicks-assessment/pkg/models"
)

// GetCages godoc
// @Summary		Get Cages
// @Description	Retrieve a listing of cages. Can filter on power status
// @Tags			Cages
// @Accept			json
// @Produce		json
// @Param			status		query		string	false	"retrieve cage listing, filter on power status"
// @Success		200			{object}	models.CagesResponse
// @Failure		400			{object}	models.CageResponse
// @Failure		404			{object}	models.CageResponse
// @Failure		500			{object}	models.CageResponse
// @Router			/cages/ [get]
func GetCages(ctx *gin.Context) {
	// Get the query parameters from the request URL.
	query := ctx.Request.URL.Query()
	status := query.Get("status")

	err := validation.Validate(status, validation.By(models.CheckStatus))
	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to validate status from query", "status", status, "options", "[ACTIVE, DOWN]", err, err.Error())
		// If there was an error validating status, respond with a bad request and the error message.
		ctx.JSON(http.StatusBadRequest, models.CageErrorResponse(err))
		return
	}

	// Get the initial builder query for the Cage entity.
	builderQuery := db.PostgresClient.Cage.Query()

	// Check if the 'status' query parameter is specified.
	if status != "" {
		// If 'status' is specified, add a WHERE clause to filter cages by status (case-insensitive).
		builderQuery = builderQuery.Where(cage.StatusEQ(cage.Status(strings.ToUpper(status))))
	}

	// Execute the final query and retrieve the cages.
	cages, err := builderQuery.All(ctx)

	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to retrieve cages", "err", err.Error())
		// If there was an error while querying the database, respond with a bad request and the error message.
		ctx.JSON(http.StatusBadRequest, models.CageErrorResponse(err))
		return
	}

	// Respond with the list of cages in the HTTP response.
	ctx.JSON(http.StatusOK, models.CagesSuccessResponse(cages))
}

// GetCage godoc
//
//	@Summary		Get Cage By ID
//	@Description	Retrieve a cage by ID
//	@Tags			Cages
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string	true	"retrieve a cage by id"
//	@Success		200			{object}	models.CageResponse
//	@Failure		400			{object}	models.CageResponse
//	@Failure		404			{object}	models.CageResponse
//	@Failure		500			{object}	models.CageResponse
//	@Router			/cages/{id} [get]
func GetCageById(ctx *gin.Context) {
	// Parse the cage ID from the URL parameter.
	id := ctx.Param("id")
	if id == "" {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to retrieve id from path, id not provided")
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(errors.New("must provide cage id")))
		return
	}

	// Parse the cage ID string into a UUID.
	u, err := uuid.Parse(id)
	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to parse id as UUID", "id", id, "err", err.Error())
		ctx.JSON(http.StatusBadRequest, models.CageErrorResponse(errors.New("could not parse id as UUID")))
		return
	}

	// Fetch the cage based on the provided cage ID.
	targetCage, err := db.PostgresClient.Cage.Query().Where(cage.ID(u)).Only(ctx)
	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to retrieve cages", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, models.CageErrorResponse(err))
		return
	}

	// Respond with the cage details in the HTTP response.
	ctx.JSON(http.StatusOK, models.CageSuccessResponse(targetCage))
}

// GetDinosaursByCageId godoc
//
//	@Summary		Get Dinosaurs By Cage ID
//	@Description	Retrieve a listing of dinosaurs in a specific cage. Can filter on attributes i.e. species or group
//	@Tags			Cages
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string	true	"retrieve a list of dinosaurs by cage id"
//	@Param			species		query		string	false	"retrieve dinosaur listing, filter on species"
//	@Param			group		query		string	false	"retrieve dinosaur listing, filter on group i.e. HERBIVORE, CARNIVORE"
//	@Success		200			{object}	models.DinosaursResponse
//	@Failure		400			{object}	models.DinosaurResponse
//	@Failure		404			{object}	models.DinosaurResponse
//	@Failure		500			{object}	models.DinosaurResponse
//	@Router			/cages/{id}/dinosaurs [get]
func GetDinosaursByCageId(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(errors.New("must provide cage id")))
		return
	}

	// Parse the cage ID string into a UUID.
	u, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(errors.New("could not parse id as UUID")))
		return
	}

	// Get the initial builder query for the Cage entity.
	builderQuery := db.PostgresClient.Cage.Query().Where(cage.ID(u)).QueryDinosaurs().WithCage()

	// Get the query parameters from the request URL.
	query := ctx.Request.URL.Query()
	group, species := query.Get("group"), query.Get("species")

	// Check if the 'group' query parameter is specified.
	if group != "" {
		// If 'group' is specified, add a WHERE clause to filter dinosaurs by group (case-insensitive).
		builderQuery = builderQuery.Where(dinosaur.GroupEQ(dinosaur.Group(strings.ToUpper(group))))
	}

	// Check if the 'species' query parameter is specified.
	if species != "" {
		// If 'species' is specified, add a WHERE clause to filter dinosaurs by species (case-insensitive).
		builderQuery = builderQuery.Where(dinosaur.SpeciesEQ(strings.ToUpper(species)))
	}

	// Execute the final query and retrieve the dinosaurs associated with the cage.
	dinosaurs, err := builderQuery.All(ctx)

	if err != nil {
		// If there was an error while querying the database, respond with a bad request and the error message.
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
		return
	}

	// Respond with the list of dinosaurs in the HTTP response.
	ctx.JSON(http.StatusOK, models.DinosaursSuccessResponse(dinosaurs))
}

// CageDinosaur godoc
//
//	@Summary		Cage Dinosaur
//	@Description	Cage a dinosaur. CageID must be provided as destination
//	@Tags			Cages
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string	true	"cage id for dinosaur"
//	@Param			dinosaur		body	models.DinosaurRequest	true	"cage a dinosaur given a cage id"
//	@Success		200			{object}	models.DinosaurResponse
//	@Failure		400			{object}	models.DinosaurResponse
//	@Failure		404			{object}	models.DinosaurResponse
//	@Failure		500			{object}	models.DinosaurResponse
//	@Router			/cages/{id}/dinosaur [post]
func CageDinosaur(ctx *gin.Context) {
	var request models.DinosaurRequest
	// Parse the request JSON data into the 'request' variable.
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to deserialize request payload", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
		return
	}

	// Parse the cage ID from the URL parameter.
	id := ctx.Param("id")
	if id == "" {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to retrieve id from path, id not provided")
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(errors.New("must provide dinosaur id")))
		return
	}

	// Set the 'CageID' in the request to the parsed cage ID from the URL parameter.
	request.CageID = id
	// Validate the 'request' data to ensure the required fields are provided (including the CageID).
	if err := request.ValidateWithCageIdRequired(); err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to validate request payload", "payload", request, err, err.Error())
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
		return
	}

	// Parse the cage ID string into a UUID.
	cageID := uuid.MustParse(request.CageID)

	// Fetch the target cage based on the provided cage ID.
	targetCage, err := db.PostgresClient.Cage.Query().Where(cage.ID(cageID)).Only(ctx)
	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to retrieve target cage", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
		return
	}

	// Check if the target cage is powered down. If it is, the dinosaur cannot be added.
	if targetCage.Status == cage.StatusDOWN {
		logger.SugaredLogger.Ctx(ctx).Errorw("cannot move dinosaur into cage when powered down")
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(errors.New("cage is powered down")))
		return
	}

	// Get the current size of the target cage.
	size := targetCage.Size

	// Check if the target cage is already full to capacity. If so, the dinosaur cannot be added.
	if size == targetCage.Capacity {
		logger.SugaredLogger.Ctx(ctx).Errorw("cage cannot hold more dinosaurs than its capacity", "size", size, "capacity", targetCage.Capacity)
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(errors.New("cage is full to capacity")))
		return
	}

	// Prepare the modifications to be applied to the dinosaur.
	mod := db.PostgresClient.Dinosaur.
		Create().
		SetName(request.Name).
		SetCage(targetCage)

	// Determine the dinosaur group based on the provided group or the species.
	var grp dinosaur.Group
	grp, err = models.GetGroup(request.Species)

	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to identity dinosaur species with any group", "payload", request, "err", err.Error())
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
		return
	}

	if request.Group != "" && !strings.EqualFold(request.Group, grp.String()) {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to match species with group in request payload", "payload", request)
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(errors.New("could not match species with group")))
		return
	}

	// Apply the dinosaur group modification.
	mod.SetGroup(grp)
	mod.SetSpecies(strings.ToUpper(request.Species))

	// Check if the cage is empty (size == 0). If it is, the dinosaur can be added directly.
	if size == 0 {
		dino, err := mod.Save(ctx)
		if err != nil {
			logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to add dinosaur to cage", "err", err.Error())
			ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
			return
		}

		// Update the cage size to add the new dinosaur.
		_, err = targetCage.Update().AddSize(1).Save(ctx)
		if err != nil {
			logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to update size of cage", "err", err.Error())
			ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, models.DinosaurSuccessResponse(dino))
		return
	}

	// Get the first dinosaur in the target cage (neighbor).
	neighbor, err := targetCage.QueryDinosaurs().First(ctx)
	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to retrieve first neighbor dinosaur in cage", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
		return
	}

	// Check if the species of the new dinosaur matches the neighbor's species.
	if strings.EqualFold(request.Species, neighbor.Species) {
		dino, err := mod.Save(ctx)
		if err != nil {
			logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to add dinosaur to cage", "err", err.Error())
			ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
			return
		}

		// Update the cage size to add the new dinosaur.
		_, err = targetCage.Update().AddSize(1).Save(ctx)
		if err != nil {
			logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to update size of cage", "err", err.Error())
			ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, models.DinosaurSuccessResponse(dino))
		return
	}

	// Check if the new dinosaur is a herbivore and the neighbor is also a herbivore.
	if grp == dinosaur.GroupHERBIVORE && neighbor.Group == grp {
		dino, err := mod.Save(ctx)
		if err != nil {
			logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to add dinosaur to cage", "err", err.Error())
			ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
			return
		}

		// Update the cage size to add the new dinosaur.
		_, err = targetCage.Update().AddSize(1).Save(ctx)
		if err != nil {
			logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to update size of cage", "err", err.Error())
			ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, models.DinosaurSuccessResponse(dino))
		return
	}

	// Check if the new dinosaur is a carnivore and the neighbor is also a carnivore.
	if grp == dinosaur.GroupCARNIVORE && neighbor.Group == grp {
		logger.SugaredLogger.Ctx(ctx).Errorw("carnivores can only be in a cage with other dinosaurs of the same species", request.Species, neighbor.Species)
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(errors.New("carnivores can only be in a cage with other dinosaurs of the same species")))
		return
	}

	// One is herbivore and the other carnivore cannot be in same cage
	logger.SugaredLogger.Ctx(ctx).Errorw("herbivores cannot be in the same cage as carnivores")
	ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(errors.New("herbivores cannot be in the same cage as carnivores")))
}

// CreateCage godoc
//
//	@Summary		Create Cage
//	@Description	Create a new cage entry.
//	@Tags			Cages
//	@Accept			json
//	@Produce		json
//	@Param			cage		body	models.CageRequest	true	"create a new cage"
//	@Success		200			{object}	models.CageResponse
//	@Failure		400			{object}	models.CageResponse
//	@Failure		404			{object}	models.CageResponse
//	@Failure		500			{object}	models.CageResponse
//	@Router			/cages [post]
func CreateCage(ctx *gin.Context) {
	var request models.CageRequest

	// Parse the request JSON data into the 'request' variable.
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to deserialize request payload", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, models.CageErrorResponse(err))
		return
	}

	// Validate the 'request' data to ensure all necessary fields are provided and have valid values.
	if err := request.Validate(); err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to validate request payload", "payload", request, err, err.Error())
		ctx.JSON(http.StatusBadRequest, models.CageErrorResponse(err))
		return
	}

	// Create the cage entity in the database based on the request data.
	builderQuery := db.PostgresClient.Cage.Create()

	if *request.Capacity > 0 {
		builderQuery = builderQuery.SetCapacity(*request.Capacity)
	}

	if request.Status != "" {
		builderQuery = builderQuery.SetStatus(cage.Status(strings.ToUpper(request.Status)))
	}

	cage, err := builderQuery.Save(ctx)

	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to create cage", "err", err.Error())
		// If there was an error while creating the cage, respond with a bad request and the error message.
		ctx.JSON(http.StatusBadRequest, models.CageErrorResponse(err))
		return
	}

	// Respond with the details of the newly created cage in the HTTP response.
	ctx.JSON(http.StatusOK, models.CageSuccessResponse(cage))
}

// EditCage godoc
//
//	@Summary		Update Cage
//	@Description	Update details of cage entry by id
//	@Tags			Cages
//	@Accept			json
//	@Produce		json
//	@Param			cage		body	models.CageRequest	true	"update cage details"
//	@Success		200			{object}	models.CageResponse
//	@Failure		400			{object}	models.CageResponse
//	@Failure		404			{object}	models.CageResponse
//	@Failure		500			{object}	models.CageResponse
//	@Router			/cages/{id} [put]
func EditCage(ctx *gin.Context) {
	var request models.CageRequest

	// Parse the request JSON data into the 'request' variable.
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to deserialize request payload", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
		return
	}

	// Validate the 'request' data to ensure all necessary fields are provided and have valid values.
	if err := request.ValidateIfUpdate(); err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to validate request payload", "payload", request, err, err.Error())
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
		return
	}

	// Parse the cage ID from the URL parameter.
	id := ctx.Param("id")
	if id == "" {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to retrieve id from path, id not provided")
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(errors.New("must provide dinosaur id")))
		return
	}

	// Parse the cage ID string into a UUID.
	u, err := uuid.Parse(id)
	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to parse id as UUID", "id", id, "err", err.Error())
		ctx.JSON(http.StatusBadRequest, models.CageErrorResponse(errors.New("could not parse id as UUID")))
		return
	}

	// Fetch the target cage based on the provided cage ID.
	targetCage, err := db.PostgresClient.Cage.Query().Where(cage.ID(u)).Only(ctx)
	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to retrieve target cage", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, models.CageErrorResponse(err))
		return
	}

	updateQuery := targetCage.Update()
	// Check if the requested status for the cage is "DOWN" and the cage contains dinosaurs.
	// If so, the cage cannot be powered off, and an error response is returned.
	if request.Status != "" {
		if strings.EqualFold(request.Status, cage.StatusDOWN.String()) && targetCage.Size > 0 {
			logger.SugaredLogger.Ctx(ctx).Errorw("failed to power down target cage, has dinosaurs", "payload", request, "capacity", targetCage.Size)
			ctx.JSON(http.StatusBadRequest, models.CageErrorResponse(errors.New("cage cannot be powered down")))
			return
		}
		updateQuery.SetStatus(cage.Status(strings.ToUpper(request.Status)))
	}

	// Check if the requested capacity is less than the current number of dinosaurs in the cage.
	// If so, the capacity cannot be decreased to a value less than the current number of dinosaurs, and an error response is returned.
	if request.Capacity != nil {
		if *request.Capacity < targetCage.Size {
			logger.SugaredLogger.Ctx(ctx).Errorw("failed to reduce cage capacity below current size", "payload", request, "size", targetCage.Size)
			ctx.JSON(http.StatusBadRequest, models.CageErrorResponse(errors.New("cage size cannot be greater than cage capacity")))
			return
		}
		updateQuery.SetCapacity(*request.Capacity)
	}

	// Update the properties of the target cage with the new values from the request.
	updatedCage, err := updateQuery.Save(ctx)

	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to update cage", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, models.CageErrorResponse(err))
		return
	}

	// Respond with the details of the updated cage in the HTTP response.
	ctx.JSON(http.StatusOK, models.CageSuccessResponse(updatedCage))
}
