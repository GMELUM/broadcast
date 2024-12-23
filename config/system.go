package config

import "broadcast/utils/env"

var (
	Token    = env.GetEnvString("TOKEN", "")
	RPSLimit = env.GetEnvInt("RPS_LIMIT", 0)
	SplitChunkCount = env.GetEnvInt("SPLIT_CHUNK_COUNT", 1e3)
)
