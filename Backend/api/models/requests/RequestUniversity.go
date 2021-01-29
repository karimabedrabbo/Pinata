package requests

type UniversityGet struct {
	UniversityId int64 `uri:"university_id" json:"university_id" binding:"required,numeric"`
}

type UniversityList struct {
	ListRequest
}

type UniversityPut struct {
	Name     string  `json:"name" binding:"required,max=255"`
	Alias    string  `json:"alias" binding:"omitempty"`
	Domain    string  `json:"domain" binding:"required,hostname_rfc1123,endswith=.edu"`
	Lat      float32 `json:"lat" binding:"required_with=Lat,latitude"`
	Lng      float32 `json:"lng" binding:"required_with=Lng,longitude"`
	Size    int  `json:"size" binding:"omitempty,numeric"`
}

type UniversityPost struct {
	UniversityId int64 `uri:"university_id" json:"university_id" binding:"required,numeric"`
	Name     string  `json:"name" binding:"required,max=255"`
	Alias    string  `json:"alias" binding:"omitempty"`
	Size    int  `json:"size" binding:"omitempty,numeric"`
}

type UniversityDelete struct {
	UniversityId int64 `uri:"university_id" json:"university_id" binding:"required,numeric"`
}


type UniversityAvatarGet struct {
	UniversityId int64 `uri:"university_id" json:"university_id" binding:"required,numeric"`
}

type UniversityAvatarPost struct {
	UniversityId int64 `uri:"university_id" json:"university_id" binding:"required,numeric"`
	*ImageAttachmentCreatable
}
