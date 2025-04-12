package external

type aiService struct{}

type IAiService interface {
	Reply(message string) (string, error)
}

func NewAiService() IAiService {
	return &aiService{}
}

func (ai *aiService) Reply(message string) (string, error) {
	return "Hello, I replied", nil
}
