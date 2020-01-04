package common

import "os"

// 项目绝对路径
var ProjectBasePath,_ = os.Getwd()

// common目录
var ProjectCommonPath = ProjectBasePath + PROJECT_COMMON_PATH

// config目录
var ProjectConfigPath = ProjectBasePath + PROJECT_CONFIG_PATH

// controller目录
var ProjectControllerPath = ProjectBasePath + PROJECT_CONTROLLER_PATH

// model目录
var ProjectModelPath = ProjectBasePath + PROJECT_MODEL_PATH

// router目录
var ProjectRouterPath = ProjectBasePath + PROJECT_ROUTER_PATH

// router目录
var ProjectLogPath = ProjectBasePath + PROJECT_LOG_PATH

const (
	PROJECT_COMMON_PATH = "/common"

	PROJECT_CONFIG_PATH = "/config"

	PROJECT_CONTROLLER_PATH = "/controller"

	PROJECT_MODEL_PATH = "/model"

	PROJECT_ROUTER_PATH = "/router"

	PROJECT_LOG_PATH = "/log"
)

