package configuration

import "maildefender/validator/internal/utils"

var (
	engineBaseEndpoint string = utils.GetEnvString("ENGINE_BASE_ENDPOINT", "")
)

func EngineBaseEndpoint() string {
	return engineBaseEndpoint
}
