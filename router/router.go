package router

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/controller"
	"github.com/noovertime7/kubemanage/controller/kubeController"
	"github.com/noovertime7/kubemanage/docs"
	"github.com/noovertime7/kubemanage/middleware"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

// InitRouter  初始化路由规则
func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Kubemanage API"
	docs.SwaggerInfo.Description = "Kubemanage"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:6180"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	podGroup := router.Group("/api/k8s/pod")
	podGroup.Use(
		middleware.TranslationMiddleware(),
		middleware.JWTAuth(),
		middleware.CasbinHandler(),
	)
	{
		kubeController.PodRegister(podGroup)
	}

	deploymentGroup := router.Group("/api/k8s/deployment")
	deploymentGroup.Use(
		middleware.TranslationMiddleware(),
		middleware.JWTAuth(),
		middleware.CasbinHandler(),
	)
	{
		kubeController.DeploymentRegister(deploymentGroup)
	}

	daemonSetGroup := router.Group("/api/k8s/daemonset")
	daemonSetGroup.Use(
		middleware.TranslationMiddleware(),
		middleware.JWTAuth(),
		middleware.CasbinHandler())
	{
		kubeController.DaemonSetRegister(daemonSetGroup)
	}

	statefulSetGroup := router.Group("/api/k8s/statefulset")
	statefulSetGroup.Use(middleware.TranslationMiddleware(), middleware.JWTAuth(), middleware.CasbinHandler())
	{
		kubeController.StatefulSetRegister(statefulSetGroup)
	}

	nodeGroup := router.Group("/api/k8s/node")
	nodeGroup.Use(middleware.TranslationMiddleware(), middleware.JWTAuth(), middleware.CasbinHandler())
	{
		kubeController.NodeRegister(nodeGroup)
	}

	namespaceGroup := router.Group("/api/k8s/namespace")
	namespaceGroup.Use(middleware.TranslationMiddleware(), middleware.JWTAuth(), middleware.CasbinHandler())
	{
		kubeController.NameSpaceRegister(namespaceGroup)
	}

	persistentVolumeGroup := router.Group("/api/k8s/persistentvolume")
	persistentVolumeGroup.Use(middleware.TranslationMiddleware(), middleware.JWTAuth(), middleware.CasbinHandler())
	{
		kubeController.PersistentVolumeRegister(persistentVolumeGroup)
	}

	serviceGroup := router.Group("/api/k8s/service")
	serviceGroup.Use(middleware.TranslationMiddleware(), middleware.JWTAuth(), middleware.CasbinHandler())
	{
		kubeController.ServiceRegister(serviceGroup)
	}

	ingressGroup := router.Group("/api/k8s/ingress")
	ingressGroup.Use(middleware.TranslationMiddleware(), middleware.JWTAuth(), middleware.CasbinHandler())
	{
		kubeController.IngressRegister(ingressGroup)
	}

	configmapGroup := router.Group("/api/k8s/configmap")
	configmapGroup.Use(middleware.TranslationMiddleware(), middleware.JWTAuth(), middleware.CasbinHandler())
	{
		kubeController.ConfigmapRegister(configmapGroup)
	}

	persistentVolumeClaimGroup := router.Group("/api/k8s/persistentvolumeclaim")
	persistentVolumeClaimGroup.Use(middleware.TranslationMiddleware(), middleware.JWTAuth(), middleware.CasbinHandler())
	{
		kubeController.PersistentVolumeClaimRegister(persistentVolumeClaimGroup)
	}

	secretGroup := router.Group("/api/k8s/secret")
	secretGroup.Use(middleware.TranslationMiddleware(), middleware.JWTAuth(), middleware.CasbinHandler())
	{
		kubeController.SecretRegister(secretGroup)
	}

	workGroup := router.Group("/api/k8s/workflow")
	workGroup.Use(middleware.TranslationMiddleware(), middleware.JWTAuth(), middleware.CasbinHandler())
	{
		kubeController.WorkFlowRegister(workGroup)
	}

	loginGroup := router.Group("/api/user")
	loginGroup.Use(middleware.TranslationMiddleware(), middleware.JWTAuth(), middleware.CasbinHandler())
	{
		controller.UserRegister(loginGroup)
	}

	monitorGroup := router.Group("/api/monitor")
	monitorGroup.Use(middleware.TranslationMiddleware(), middleware.JWTAuth(), middleware.CasbinHandler())
	{
		kubeController.MonitroRegister(monitorGroup)
	}

	menuGroup := router.Group("/api/menu")
	menuGroup.Use(middleware.TranslationMiddleware(), middleware.JWTAuth(), middleware.CasbinHandler())
	{
		controller.MenuRegister(menuGroup)
	}

	casbinGroup := router.Group("/api/casbin")
	casbinGroup.Use(middleware.TranslationMiddleware(), middleware.JWTAuth(), middleware.CasbinHandler())
	{
		controller.CasbinRegister(casbinGroup)
	}

	return router
}
