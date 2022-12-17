package sys

import (
	"strconv"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pkg/errors"

	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dto"
)

type CasbinServiceGetter interface {
	CasbinService() CasbinService
}

type CasbinService interface {
	UpdateCasbin(AuthorityID uint, casbinInfos []dto.CasbinInfo) error
	UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error
	GetPolicyPathByAuthorityId(AuthorityID uint) (pathMaps []dto.CasbinInfo)
	Casbin() *casbin.CachedEnforcer
}

type casbinService struct {
	factory dao.ShareDaoFactory
}

func NewCasbinService(factory dao.ShareDaoFactory) CasbinService {
	return &casbinService{factory: factory}
}

func (c *casbinService) UpdateCasbin(AuthorityID uint, casbinInfos []dto.CasbinInfo) error {
	authorityId := strconv.Itoa(int(AuthorityID))
	c.ClearCasbin(0, authorityId)
	var rules [][]string
	for _, v := range casbinInfos {
		rules = append(rules, []string{authorityId, v.Path, v.Method})
	}
	e := c.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	err := e.InvalidateCache()
	if err != nil {
		return err
	}
	return nil
}

func (c *casbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := c.factory.GetDB().Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	e := c.Casbin()
	err = e.InvalidateCache()
	if err != nil {
		return err
	}
	return err
}

func (c *casbinService) GetPolicyPathByAuthorityId(AuthorityID uint) (pathMaps []dto.CasbinInfo) {
	e := c.Casbin()
	authorityId := strconv.Itoa(int(AuthorityID))
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, dto.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

func (c *casbinService) ClearCasbin(v int, p ...string) bool {
	e := c.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

var (
	cachedEnforcer *casbin.CachedEnforcer
	once           sync.Once
)

func (c *casbinService) Casbin() *casbin.CachedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(c.factory.GetDB())
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub)  && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act) 
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			return
		}
		cachedEnforcer, _ = casbin.NewCachedEnforcer(m, a)
		cachedEnforcer.SetExpireTime(60 * 60)
		_ = cachedEnforcer.LoadPolicy()
	})
	return cachedEnforcer
}
