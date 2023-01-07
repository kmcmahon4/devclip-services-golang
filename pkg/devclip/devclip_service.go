package devclip_service

type DevclipService interface {
	// declare functions here
}

type devclipService struct {
}

var _ DevclipService = (*devclipService)(nil)

func NewDevclipService() DevclipService {
	return &devclipService{}
}

// functions
