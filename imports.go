package main

import (
	_ "admin/business/controller"
	_ "admin/business/controller/system"
	_ "admin/business/pogo/entity/system"
	_ "admin/business/service/system/impl"
	_ "admin/common/service/impl"
	_ "admin/component/cache/redis"
	_ "admin/component/db/mysql"
	_ "admin/component/logger"
	_ "admin/config"
	_ "admin/filter/authFilter"
)
