package getconf

import (
	"github.com/udonetsm/help/helper"
	"github.com/udonetsm/help/models"
)

var Storeconf string         //for configure postgres conn str(using in db)
var Server models.Srver_Conf //for configure server and secret key for tokens

func init() {
	home := helper.Home() + "/"
	conf := models.Postgres_conf{}
	conf = conf.StoreConf(home + ".confs/conn_config.yaml")
	Storeconf = conf.Dbname + conf.Dbpassword + conf.SslMode + conf.Dbport + conf.Dbuser + conf.Dbhost
	Server = Server.ServerConf(home + ".confs/server_conf_auth.yaml")
}
