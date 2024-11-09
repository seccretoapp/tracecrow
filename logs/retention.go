package logs

import (
	"fmt"
	"time"
)

// RetainLogs mantém apenas logs dentro de um intervalo de tempo específico
func RetainLogs(logs []LogEntry, retentionPeriod time.Duration) ([]LogEntry, error) {
	var retainedLogs []LogEntry
	currentTime := time.Now().Unix()

	for _, log := range logs {
		if currentTime-int64(log.Timestamp) <= int64(retentionPeriod.Seconds()) {
			retainedLogs = append(retainedLogs, log)
		}
	}

	if len(retainedLogs) == 0 {
		return nil, fmt.Errorf("no logs retained within the retention period")
	}

	return retainedLogs, nil
}
