package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samsonannan/prizepicks-assessment/pkg/db"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/cage"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/dinosaur"
	"github.com/samsonannan/prizepicks-assessment/pkg/logger"
	"github.com/samsonannan/prizepicks-assessment/pkg/models"
)

// GetDinosaurs godoc
//
//	@Summary		Get Dinosaurs
//	@Description	Retrieve a listing of dinosaurs. Can filter on attributes i.e. species or group
//	@Tags			Dinosaurs
//	@Accept			json
//	@Produce		json
//	@Param			species		query		string	false	"retrieve dinosaur listing, filter on species"
//	@Param			group		query		string	false	"retrieve dinosaur listing, filter on group i.e. herbivore, carnivore"
//	@Success		200			{object}	models.DinosaursResponse
//	@Failure		400			{object}	models.DinosaurResponse
//	@Failure		404			{object}	models.DinosaurResponse
//	@Failure		500			{object}	models.DinosaurResponse
//	@Router			/dinosaurs/ [get]
func GetDinosaurs(ctx *gin.Context) {
	// Get the query parameters from the request URL.
	query := ctx.Request.URL.Query()
	group, species := query.Get("group"), query.Get("species")

	// Get the initial builder query for the Dinosaur entity.
	builderQuery := db.PostgresClient.Dinosaur.Query().WithCage()

	// Check if the 'group' query parameter is specified.
	if group != "" {
		// If 'group' is specified, add a WHERE clause to filter dinosaurs by group (case-insensitive).
		builderQuery = builderQuery.Where(dinosaur.GroupEQ(dinosaur.Group(strings.ToUpper(group))))
	}

	// Check if the 'species' query parameter is specified.
	if species != "" {
		// If 'species' is specified, add a WHERE clause to filter dinosaurs by species (case-insensitive).
		builderQuery = builderQuery.Where(dinosaur.Species(strings.ToUpper(species)))
	}

	// Execute the final query and retrieve the dinosaurs.
	dinosaurs, err := builderQuery.All(ctx)

	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to retrieve dinosaurs", "err", err.Error())
		// If there was an error while querying the database, respond with a bad request and the error message.
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
		return
	}

	// Respond with the list of dinosaurs in the HTTP response.
	ctx.JSON(http.StatusOK, models.DinosaursSuccessResponse(dinosaurs))
}

// GetDinosaur godoc
//
//	@Summary		Get Dinosaur By ID
//	@Description	Retrieve a dinosaur by ID
//	@Tags			Dinosaurs
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string	true	"retrieve a dinosaur by id"
//	@Success		200			{object}	models.DinosaurResponse
//	@Failure		400			{object}	models.DinosaurResponse
//	@Failure		404			{object}	models.DinosaurResponse
//	@Failure		500			{object}	models.DinosaurResponse
//	@Router			/dinosaurs/{id} [get]
func GetDinosaurById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(errors.New("must provide dinosaur id")))
		return
	}

	// Parse the ID string into a UUID.
	u, err := uuid.Parse(id)
	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to parse id as uuid", "id", id, "err", err.Error())
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(errors.New("could not parse id as uuid")))
		return
	}

	// Use the UUID to fetch the dinosaur from the database.
	dino, err := db.PostgresClient.Dinosaur.Query().Where(dinosaur.ID(u)).WithCage().Only(ctx)
	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to retrieve dinosaur", "err", err.Error())
		// If there was an error while querying the database, respond with a bad request and the error message.
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
		return
	}

	// Respond with the dinosaur in the HTTP response.
	ctx.JSON(http.StatusOK, models.DinosaurSuccessResponse(dino))
}

// GetCageByDinosaurId godoc
//
//	@Summary		Get Cage By Dinosaur ID
//	@Description	Retrieve the cage for dinosaur.
//	@Tags			Dinosaurs
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string	true	"retrieve cage by dinosaur id"
//	@Success		200			{object}	models.CageResponse
//	@Failure		400			{object}	models.CageResponse
//	@Failure		404			{object}	models.CageResponse
//	@Failure		500			{object}	models.CageResponse
//	@Router			/dinosaurs/{id}/cage [get]
func GetCageByDinosaurId(ctx *gin.Context) {
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

	// Get the initial builder query for the Cage entity.
	targetCage, err := db.PostgresClient.Dinosaur.Query().Where(dinosaur.ID(u)).QueryCage().Only(ctx)
	// Cage.Query().Where(cage.ID(u)).QueryDinosaurs()

	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to retrieve cage", "err", err.Error())
		// If there was an error while querying the database, respond with a bad request and the error message.
		ctx.JSON(http.StatusBadRequest, models.CageErrorResponse(err))
		return
	}

	// Respond with the cage details in the HTTP response.
	ctx.JSON(http.StatusOK, models.CageSuccessResponse(targetCage))
}

// EditDinosaur godoc
//
//	@Summary		Update Dinosaur
//	@Description	Updates information stored on a dinosaur. Can move dinosaur to cage.
//	@Tags			Dinosaurs
//	@Accept			json
//	@Produce		json
//	@Param			dinosaur		body	models.DinosaurRequest	true	"update dinosaur information"
//	@Success		200			{object}	models.DinosaurResponse
//	@Failure		400			{object}	models.DinosaurResponse
//	@Failure		404			{object}	models.DinosaurResponse
//	@Failure		500			{object}	models.DinosaurResponse
//	@Router			/dinosaurs/{id} [put]
func EditDinosaur(ctx *gin.Context) {
	var err error
	var targetCage *ent.Cage
	var request models.DinosaurRequest
	// Parse the request JSON data into the 'request' variable.
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
		return
	}

	// Validate the 'request' data to ensure the required fields are provided (including the CageID).
	if err := request.ValidateIfUpdate(); err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to validate request payload", "payload", request, err, err.Error())
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
		return
	}

	// Parse the dinosaur ID from the URL parameter.
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(errors.New("must provide dinosaur id")))
		return
	}

	// Parse the cage ID string into a UUID.
	u, err := uuid.Parse(id)
	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to parse id as uuid", "id", id, "err", err.Error())
		ctx.JSON(http.StatusBadRequest, models.CageErrorResponse(errors.New("could not parse id as uuid")))
		return
	}

	// Fetch the target dinosaur based on the provided dinosaur ID.
	targetDino, err := db.PostgresClient.Dinosaur.Query().Where(dinosaur.ID(u)).Only(ctx)
	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to retrieve target dinosaur", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
		return
	}

	// Fetch the current cage associated with the target dinosaur.
	currentCage, err := targetDino.QueryCage().Only(ctx)
	if err != nil {
		logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to retrieve current cage", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
		return
	}

	// Prepare the modifications to be applied to the target dinosaur.
	mod := db.PostgresClient.Dinosaur.UpdateOneID(u)

	if request.Name != "" {
		mod.SetName(request.Name)
	}

	var grp dinosaur.Group
	if request.Species != "" {
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
	} else {
		grp = targetDino.Group
		if request.Group != "" && !strings.EqualFold(request.Group, grp.String()) {
			logger.SugaredLogger.Ctx(ctx).Errorw("failed to match species with group in request payload", "payload", request)
			ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(errors.New("could not match species with group")))
			return
		}
		request.Species = targetDino.Species
	}

	// Fetch the target cage based on the CageID provided in the request.
	if request.CageID != "" {
		targetCage, err = db.PostgresClient.Cage.Query().Where(cage.ID(uuid.MustParse(request.CageID))).Only(ctx)
		if err != nil {
			logger.SugaredLogger.Ctx(ctx).Errorw("failed to execute query to retrieve target cage", "err", err.Error())
			ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
			return
		}

		if targetCage.ID != currentCage.ID {
			mod.SetCage(targetCage)
			// Check if the target cage is powered down. If it is, the dinosaur cannot be edited.
			if targetCage.Status == cage.StatusDOWN {
				ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(errors.New("cage is powered down")))
				return
			}
			// Check if the target cage is already full to capacity. If so, the dinosaur cannot be moved.
			if targetCage.Size == targetCage.Capacity {
				ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(errors.New("cage is full to capacity")))
				return
			}

			if targetCage.Size == 0 {
				// Save the modifications to the dinosaur.
				dino, err := mod.Save(ctx)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
					return
				}

				// Increase the size of the target cage (adding the dinosaur to it).
				_, err = targetCage.Update().AddSize(1).Save(ctx)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
					return
				}

				// Reduce the size of the current cage (removing the dinosaur from it).
				_, err = currentCage.Update().AddSize(-1).Save(ctx)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
					return
				}
				// Respond with the edited dinosaur in the HTTP response.
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

				// Reduce the size of the current cage (removing the dinosaur from it).
				_, err = currentCage.Update().AddSize(-1).Save(ctx)
				if err != nil {
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

				// Reduce the size of the current cage (removing the dinosaur from it).
				_, err = currentCage.Update().AddSize(-1).Save(ctx)
				if err != nil {
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
			return
		}
	}

	// Save the modifications to the dinosaur.
	dino, err := mod.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.DinosaurErrorResponse(err))
		return
	}
	// Respond with the edited dinosaur in the HTTP response.
	ctx.JSON(http.StatusOK, models.DinosaurSuccessResponse(dino))
}
