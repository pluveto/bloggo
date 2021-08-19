package service

func (s *Service) GetSetting(key string) (value string, err error) {
	return s.repo.GetSetting(key)
}

func (s *Service) GetSettingOrDefault(key string, defaultValue string) (value string, err error) {
	value, err = s.repo.GetSetting(key)
	if err != nil {
		value = defaultValue
		err = nil
	}
	return
}

// TODO
func (s *Service) SetSetting(key string, value string) {

}
