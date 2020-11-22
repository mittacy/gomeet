package service

import (
	"com/mittacy/gomeet/model"
	"com/mittacy/gomeet/repository"
)

type IMeetingService interface {
	AddMeeting(meeting model.Meeting) error
	DeleteMeeting(id int) error
	UpdateMeeting(meeting model.Meeting) error
	GetMeetingByID(id int) (model.Meeting, error)
	GetMeetingCountByBuilding(buildingID int) (int, error)
	GetMeetingsByPage(page, onePageCount, buildingID int) ([]model.Meeting, error)
	GetAllMeetingTypes() []string
	GetAllScaleTypes() []string
	GetAllMeetingByLayer() ([]model.Meeting, error)
	GetAllMeetingByBuilding(buildingID int) ([]model.Meeting, error)
}

func NewMeetingService(meetingRepo repository.IMeetingRepository) IMeetingService {
	return &MeetingService{meetingRepo}
}

type MeetingService struct {
	MeetingRepository repository.IMeetingRepository
}

// AddMeeting 添加会议室
func (ms *MeetingService) AddMeeting(meeting model.Meeting) error {
	return ms.MeetingRepository.InsertMeeting(meeting)
}

func (ms *MeetingService) DeleteMeeting(id int) error {
	return ms.MeetingRepository.DeleteMeeting(id)
}

func (ms *MeetingService) UpdateMeeting(meeting model.Meeting) error {
	//// 确保建筑存在
	//exists := false
	//var err error
	//if exists, err = ms.BuildingRepository.IsBuildingExists(meeting.BuildingID); err != nil {
	//	return err
	//}
	//if !exists {
	//	return errors.New("the building no exists")
	//}
	//// 确保会议室类型符合
	//if !model.IsMeetingType(meeting.MeetingType) {
	//	return errors.New("the meeting type no exists")
	//}
	//// 确保会议室容量符合
	//if !model.IsScaleType(meeting.Scale) {
	//	return errors.New("the meeting scale no exists")
	//}
	return ms.MeetingRepository.UpdateMeeting(meeting)
}

func (ms *MeetingService) GetMeetingByID(id int) (model.Meeting, error) {
	return ms.MeetingRepository.SelectMeetingByID((id))
}

func (ms *MeetingService) GetMeetingCountByBuilding(buildingID int) (int, error) {
	return ms.MeetingRepository.SelectMeetingCountCountByBuilding(buildingID)
}

func (ms *MeetingService) GetMeetingsByPage(page, onePageCount, buildingID int) ([]model.Meeting, error) {
	return ms.MeetingRepository.SelectMeetingsByBuilding(buildingID, page, onePageCount)
}

func (ms *MeetingService) GetAllMeetingTypes() []string {
	return ms.MeetingRepository.SelectAllMeetingTypes()
}

func (ms *MeetingService) GetAllScaleTypes() []string {
	return ms.MeetingRepository.SelectAllScaleTypes()
}

func (ms *MeetingService) GetAllMeetingByLayer() ([]model.Meeting, error) {
	return []model.Meeting{}, nil
}

func (ms *MeetingService) GetAllMeetingByBuilding(buildingID int) ([]model.Meeting, error) {
	return ms.MeetingRepository.SelectMeetingsByBuilding(buildingID)
}

