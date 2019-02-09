package reisapi-ruter






type Service struct {
	Client interface {
		Do(v interface{}) ([]byte, error)
	}
}

func New(c echo.Context) *Service {
	timeout := viper.GetDuration("service.age.timeout") * time.Second
	url := viper.GetString("service.age.url")

	header := https.Header{
		ContentType: viper.GetString("service.age.header.contenttype"),
		// BasicAuthen: https.BasicAuthen{
		// 	Username: viper.GetString("service.age.header.basicauthen.username"),
		// 	Password: viper.GetString("service.age.header.basicauthen.password"),
		// },
	}

	b, err := ioutil.ReadFile(viper.GetString("service.age.cert"))
	if err != nil {
		return &Service{
			Client: https.NewXMLClient(timeout, url, header, c),
		}
	}

	return &Service{
		Client: https.NewXMLClients(timeout, url, header, b, c),
	}
}