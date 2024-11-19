package access

import (
	"taffeite.com/taffeite-underlying-service/conf"
	//entity "taffeite.com/taffeite-underlying-service/domain"

	"taffeite.com/taffeite-underlying-service/repository"
)

type TaffeiteModule struct {
	InfoDataRepository  repository.InfoDataRepository
	PanelViewRepository repository.PanelViewRepository
	MongoConf           *conf.MongoConf
}

func NewTaffeiteModule() *TaffeiteModule {
	var mongoConf = conf.NewMongoConf()
	var panelViewRepository = repository.PanelViewRepository{Conf: *mongoConf}
	var infoDataRepository = repository.InfoDataRepository{Conf: *mongoConf}
	//var defaultData = entity.GetDefaultPanelViewData()
	//panelViewRepository.InsertPanelViewData(defaultData)
	return &TaffeiteModule{InfoDataRepository: infoDataRepository, PanelViewRepository: panelViewRepository, MongoConf: mongoConf}
}

func (tf *TaffeiteModule) Clear() {
	tf.MongoConf.Clear()
}
