package service

// ISiteSettingService mplements SiteSettingServicer
type SiteSettingService struct {
}

func (s *SiteSettingService) Get(key string) interface{} {
	return nil
}
func (s *SiteSettingService) GetOrDefault(key string, defaultValue interface{}) interface{}
func (s *SiteSettingService) Set(key string, value interface{})
