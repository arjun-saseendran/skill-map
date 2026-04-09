package service

import (
	"errors"
	"fmt"

	"github.com/arjun-saseendran/skill-map/db"
	"github.com/arjun-saseendran/skill-map/dto"
	"github.com/arjun-saseendran/skill-map/models"
)

type SkillService interface {
	Create(inputData *dto.SkillCreateInput) (*models.Skill, error)
	List() ([]models.Skill, error)
	Get(id string) (*models.Skill, error)
	Update(id string, inputData *dto.SkillUpdateInput) (*models.Skill, error)
	Delete(id string) error
	CreateGroup(inputData *dto.SkillGroupCreateInput) (*models.SkillGroup, error)
	ListGroup() ([]models.SkillGroup, error)
	GetGroup(id string) (*models.SkillGroup, error)
	UpdateGroup(id string, inputData *dto.SkillGroupUpdateInput) (*models.SkillGroup, error)
	DeleteGroup(id string) error
}

type skillService struct {
}

func NewSkillService() SkillService {
	return &skillService{}
}

func (skillServ *skillService) Create(inputData *dto.SkillCreateInput) (*models.Skill, error) {
	newSkill := &models.Skill{Name: "Java"}
	db.DB.Create(newSkill)
	if newSkill.ID == 0 {
		return nil, errors.New("skill creation failed")
	}
	return newSkill, nil
}

func (skillServ *skillService) List() ([]models.Skill, error) {
	var skills []models.Skill
	db.DB.Find(&skills)
	return skills, nil
}

func (skillServ *skillService) Get(id string) (*models.Skill, error) {
	skill := models.NewSkill()
	db.DB.First(skill, id)
	if skill.ID == 0 {
		return nil, errors.New("skill not found")
	}
	return skill, nil
}

func (skillServ *skillService) Update(id string, inputData *dto.SkillUpdateInput) (*models.Skill, error) {
	skill := &models.Skill{}
	db.DB.First(skill, id)
	if skill.ID == 0 {
		return nil, errors.New("skill not found")
	}
	db.DB.Model(skill).Updates(models.Skill{Name: inputData.Name})
	return skill, nil
}

func (skillServ *skillService) Delete(id string) error {
	var skill models.Skill
	db.DB.First(&skill, id)
	if skill.ID == 0 {
		return errors.New("skill not found")
	}
	db.DB.Delete(&skill)
	return nil
}

func (skillServ *skillService) CreateGroup(inputData *dto.SkillGroupCreateInput) (*models.SkillGroup, error) {
	newSkillGroup := &models.SkillGroup{Name: inputData.Name}
	db.DB.Create(newSkillGroup)
	if newSkillGroup.ID == 0 {
		return nil, errors.New("skill group creation failed")
	}
	return newSkillGroup, nil
}

func (skillServ *skillService) ListGroup() ([]models.SkillGroup, error) {
	var skillGroups []models.SkillGroup
	db.DB.Model(&skillGroups).Preload("Skills").Find(&skillGroups)
	return skillGroups, nil
}

func (skillServ *skillService) GetGroup(id string) (*models.SkillGroup, error) {
	skillGroup := models.NewSkillGroup()
	//db.DB.Model(skillGroup).Preload("Skills").Find(skillGroup)
	db.DB.First(skillGroup, id)
	if skillGroup.ID == 0 {
		return nil, errors.New("skill group not found")
	}
	return skillGroup, nil
}

func (skillServ *skillService) UpdateGroup(id string, inputData *dto.SkillGroupUpdateInput) (*models.SkillGroup, error) {
	skillGroup := models.NewSkillGroup()
	db.DB.First(skillGroup, id)
	if skillGroup.ID == 0 {
		return nil, errors.New("skill group not found")
	}
	skillMapping, err := skillServ.GetSkillList(inputData.Skills)
	if err != nil {
		return nil, err
	}
	db.DB.Model(skillGroup).Updates(models.Skill{Name: inputData.Name})
	db.DB.Model(skillGroup).Association("Skills").Replace(skillMapping)
	return skillGroup, nil

}

func (skillServ *skillService) GetSkillList(inputList []int) ([]*models.Skill, error) {
	var skills []*models.Skill
	for _, id := range inputList {
		skill := models.NewSkill()
		db.DB.First(skill, id)
		if skill.ID == 0 {
			return nil, fmt.Errorf("skill with %v not exists", id)
			skills = append(skills, skill)
		}
	}
	return skills, nil
}

func (skillServ *skillService) DeleteGroup(id string) error {
	skillGroup := models.NewSkillGroup()
	db.DB.First(skillGroup, id)
	if skillGroup.ID == 0 {
		return errors.New("skill group not found")
	}
	db.DB.Delete(skillGroup)
	return nil
}
