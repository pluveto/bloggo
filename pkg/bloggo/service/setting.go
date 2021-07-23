package service

func (s *Service) GetSetting(key string) (value string, err error) {
	return s.repo.GetSetting(key)
}

func (s *Service) GetOrDefault(key string, defaultValue string) (value string, err error) {
	value, err = s.repo.GetSetting(key)
	if err != nil {
		value = defaultValue
		err = nil
	}
	return
}

// TODO
func (s *Service) Set(key string, value string) {

}
