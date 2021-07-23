package service

func (s *Service) GetSetting(key string) interface{} {
	return s.repo.GetSetting(key)
}

// TODO
func (s *Service) GetOrDefault(key string, defaultValue interface{}) interface{} {
	return nil
}

// TODO
func (s *Service) Set(key string, value interface{}) {

}
