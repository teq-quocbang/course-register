package connection

import (
	"github.com/teq-quocbang/course-register/cache"
	"github.com/teq-quocbang/course-register/cache/services/register"
)

func (r redisDB) Register() cache.RegisterService {
	return register.NewRedisRegisterCache(r.redis)
}
