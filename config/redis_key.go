package config

const(
	RdsKeyVersion = "rdm_version"	// hset no xx url xxx.com
	RdsKeyQaId = "rdm_aQaId"	// inc [RdsKeyQaId]
	RdsKeyQa = "rdm_zQa"		// zadd [RdsKeyQa] [inc RdsKeyQaId] [data.Qa]
	RdsKeyQaReply = "rdm_zQaReply"		// zadd [RdsKeyQa] [inc RdsKeyQaId] [data.QaReply]
)

