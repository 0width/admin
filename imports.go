package main

import (
	_ "admin/business/controller"
	_ "admin/business/controller/system"
	_ "admin/business/filter/roleFilter"
	_ "admin/business/pogo/entity"
	_ "admin/business/service/common/impl"
	_ "admin/business/service/system/impl"
	_ "admin/component/cache/redis"
	_ "admin/component/db/mysql"
	_ "admin/component/jwt"
	_ "admin/config"
)
