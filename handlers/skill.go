package handlers

import (
	"net/http"

	"github.com/arjun-saseendran/skill-map/dto"
	"github.com/arjun-saseendran/skill-map/service"
	"github.com/gin-gonic/gin"
)

type SkillHandler struct {
	groupName    string
	skillService service.SkillService
}

func NewSkillHandleFrom(skillService service.SkillService) *SkillHandler {
	return &SkillHandler{
		groupName:    "api/",
		skillService: skillService,
	}

}

func (handler *SkillHandler) RegisterEndpoints(r *gin.Engine) {
	skillGroup := r.Group(handler.groupName)

	skillGroup.GET("/skills", handler.Skills)
	skillGroup.POST("/skill", handler.CreateSkill)
	skillGroup.GET("/skill/:id", handler.Skill)
	skillGroup.PATCH("/skill/:id", handler.UpdateSkill)
	skillGroup.DELETE("/skill/:id", handler.DeleteSkill)

	skillGroup.GET("/skill-groups", handler.SkillGroups)
	skillGroup.GET("/skill/:id", handler.SkillGroup)
	skillGroup.POST("/skill-group", handler.CreateSkillGroup)
	skillGroup.PATCH("/skill-group/:id", handler.UpdateSkillGroup)
	skillGroup.DELETE("/skill-group/:id", handler.DeleteSkillGroup)

}

func (handler *SkillHandler) CreateSkill(ctx *gin.Context) {
	inputData := dto.NewSkillCreateInput()
	err := ctx.BindJSON(&inputData)
	if err != nil {
		dto.BadResponse(ctx, "failed to bind input data.")
		return
	}
	newSkill, err := handler.skillService.Create(inputData)
	if err != nil {
		dto.BadResponse(ctx, "failed to create skill.")
		return
	}
	ctx.JSON(http.StatusOK, newSkill)
}

func (handler *SkillHandler) Skills(ctx *gin.Context) {
	skills, err := handler.skillService.List()
	if err != nil {
		dto.BadResponse(ctx, "failed to get skill data.")
		return
	}
	ctx.JSON(http.StatusOK, skills)
}

func (handler *SkillHandler) Skill(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		dto.BadResponse(ctx, "invalid id.")
		return
	}
	skill, err := handler.skillService.Get(id)
	if err != nil {
		dto.BadResponse(ctx, "failed to get skill data.")
		return
	}
	ctx.JSON(http.StatusOK, skill)

}

func (handler *SkillHandler) DeleteSkill(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		dto.BadResponse(ctx, "invalid id")
		return
	}
	err := handler.skillService.Delete(id)
	if err != nil {
		dto.BadResponse(ctx, "failed to delete skill.")
		return
	}
	dto.SuccessResponse(ctx, "skill deleted.")
}

func (handler *SkillHandler) UpdateSkill(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		dto.BadResponse(ctx, "invalid skill id.")
		return
	}
	inputData := dto.NewSkillUpdateInput()
	err := ctx.BindJSON(&inputData)
	if err != nil {
		dto.BadResponse(ctx, "failed to bind skill data.")
		return
	}
	skill, err := handler.skillService.Update(id, inputData)
	if err != nil {
		dto.BadResponse(ctx, "failed to update skill.")
		return
	}
	ctx.JSON(http.StatusOK, skill)

}

func (handler *SkillHandler) CreateSkillGroup(ctx *gin.Context) {
	inputData := dto.NewSkillGroupCreateInput()
	err := ctx.BindJSON(&inputData)
	if err != nil {
		dto.BadResponse(ctx, "failed to bind skill group data.")
		return
	}
	skillGroup, err := handler.skillService.CreateGroup(inputData)
	if err != nil {
		dto.BadResponse(ctx, "failed to create new skill group")
		return
	}
	ctx.JSON(http.StatusOK, skillGroup)
}

func (handler *SkillHandler) SkillGroups(ctx *gin.Context) {
	skillGroups, err := handler.skillService.ListGroup()
	if err != nil {
		dto.BadResponse(ctx, "failed to get skill group data.")
		return
	}
	ctx.JSON(http.StatusOK, skillGroups)
}

func (handler *SkillHandler) SkillGroup(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		dto.BadResponse(ctx, "invalid skill group id.")
		return
	}
	skillGroup, err := handler.skillService.GetGroup(id)
	if err != nil {
		dto.BadResponse(ctx, "invalid skill group id.")
		return
	}
	ctx.JSON(http.StatusOK, skillGroup)
}

func (handler *SkillHandler) DeleteSkillGroup(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		dto.BadResponse(ctx, "invalid skill group id.")
		return
	}
	err := handler.skillService.DeleteGroup(id)
	if err != nil {
		dto.BadResponse(ctx, "failed to delete skill group id.")
		return
	}
	dto.SuccessResponse(ctx, "skill group deleted.")
}

func (handler *SkillHandler) UpdateSkillGroup(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		dto.BadResponse(ctx, "invalid skill group id.")
		return
	}
	inputData := dto.NewSkillGroupUpdateInput()
	err := ctx.BindJSON(&inputData)
	if err != nil {
		dto.BadResponse(ctx, "failed bind skill group input data.")
		return
	}

	skillGroup, err := handler.skillService.UpdateGroup(id, inputData)
	if err != nil {
		dto.BadResponse(ctx, "failed to update skill group.")
		return
	}
	ctx.JSON(http.StatusOK, skillGroup)
}
