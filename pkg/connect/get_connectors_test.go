package connect

import (
	"testing"

	"github.com/cloudhut/connect-client"
	"github.com/stretchr/testify/assert"
)

func Test_traceToErrorContent(t *testing.T) {
	type test struct {
		name         string
		initialValue string
		trace        string
		expected     string
	}

	tests := []test{
		{
			name:         "empty trace",
			initialValue: "connector 0 error.",
			trace:        "",
			expected:     "connector 0 error.",
		},
		{
			name:         "caused by",
			initialValue: "connector 0 error",
			trace: `org.apache.kafka.common.KafkaException: The constructor of org.apache.kafka.clients.admin.ForwardingAdmin threw an exception
			at org.apache.kafka.common.utils.Utils.newParameterizedInstance(Utils.java:469)
			at org.apache.kafka.connect.mirror.MirrorConnectorConfig.forwardingAdmin(MirrorConnectorConfig.java:211)
			at org.apache.kafka.connect.mirror.MirrorHeartbeatConnector.start(MirrorHeartbeatConnector.java:51)
			at org.apache.kafka.connect.runtime.WorkerConnector.doStart(WorkerConnector.java:190)
			at org.apache.kafka.connect.runtime.WorkerConnector.start(WorkerConnector.java:215)
			at org.apache.kafka.connect.runtime.WorkerConnector.doTransitionTo(WorkerConnector.java:360)
			at org.apache.kafka.connect.runtime.WorkerConnector.doTransitionTo(WorkerConnector.java:343)
			at org.apache.kafka.connect.runtime.WorkerConnector.doRun(WorkerConnector.java:143)
			at org.apache.kafka.connect.runtime.WorkerConnector.run(WorkerConnector.java:121)
			at org.apache.kafka.connect.runtime.isolation.Plugins.lambda$withClassLoader$1(Plugins.java:177)
			at java.base/java.util.concurrent.Executors$RunnableAdapter.call(Executors.java:539)
			at java.base/java.util.concurrent.FutureTask.run(FutureTask.java:264)
			at java.base/java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1136)
			at java.base/java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:635)
			at java.base/java.lang.Thread.run(Thread.java:833)
		Caused by: org.apache.kafka.common.config.ConfigException: Missing required configuration "bootstrap.servers" which has no default value.
			at org.apache.kafka.common.config.ConfigDef.parseValue(ConfigDef.java:493)
			at org.apache.kafka.common.config.ConfigDef.parse(ConfigDef.java:483)
			at org.apache.kafka.common.config.AbstractConfig.<init>(AbstractConfig.java:113)
			at org.apache.kafka.common.config.AbstractConfig.<init>(AbstractConfig.java:146)
			at org.apache.kafka.clients.admin.AdminClientConfig.<init>(AdminClientConfig.java:244)
			at org.apache.kafka.clients.admin.Admin.create(Admin.java:144)
			at org.apache.kafka.clients.admin.ForwardingAdmin.<init>(ForwardingAdmin.java:51)
			at java.base/jdk.internal.reflect.NativeConstructorAccessorImpl.newInstance0(Native Method)
			at java.base/jdk.internal.reflect.NativeConstructorAccessorImpl.newInstance(NativeConstructorAccessorImpl.java:77)
			at java.base/jdk.internal.reflect.DelegatingConstructorAccessorImpl.newInstance(DelegatingConstructorAccessorImpl.java:45)
			at java.base/java.lang.reflect.Constructor.newInstanceWithCaller(Constructor.java:499)
			at java.base/java.lang.reflect.Constructor.newInstance(Constructor.java:480)
			at org.apache.kafka.common.utils.Utils.newParameterizedInstance(Utils.java:458)
			... 14 more`,
			expected: `Missing required configuration "bootstrap.servers" which has no default value.`,
		},
		{
			name:         "first line",
			initialValue: "connector 0 error",
			trace: `org.apache.kafka.common.config.ConfigException: Cannot connect to 'c' S3 bucket due to: The specified bucket is not valid.
			at com.redpanda.kafka.connect.s3.config.AwsConfigValidator.validate(AwsConfigValidator.java:57)
			at com.redpanda.kafka.connect.s3.S3SinkConnector.start(S3SinkConnector.java:71)
			at org.apache.kafka.connect.runtime.WorkerConnector.doStart(WorkerConnector.java:190)
			at org.apache.kafka.connect.runtime.WorkerConnector.start(WorkerConnector.java:215)
			at org.apache.kafka.connect.runtime.WorkerConnector.doTransitionTo(WorkerConnector.java:360)
			at org.apache.kafka.connect.runtime.WorkerConnector.doTransitionTo(WorkerConnector.java:343)
			at org.apache.kafka.connect.runtime.WorkerConnector.doRun(WorkerConnector.java:143)
			at org.apache.kafka.connect.runtime.WorkerConnector.run(WorkerConnector.java:121)
			at org.apache.kafka.connect.runtime.isolation.Plugins.lambda$withClassLoader$1(Plugins.java:177)
			at java.base/java.util.concurrent.Executors$RunnableAdapter.call(Executors.java:539)
			at java.base/java.util.concurrent.FutureTask.run(FutureTask.java:264)
			at java.base/java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1136)
			at java.base/java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:635)
			at java.base/java.lang.Thread.run(Thread.java:833)`,
			expected: `Cannot connect to 'c' S3 bucket due to: The specified bucket is not valid.`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := traceToErrorContent(tc.initialValue, tc.trace)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func Test_connectorsResponseToClusterConnectorInfo(t *testing.T) {
	type test struct {
		name       string
		input      *connect.ListConnectorsResponseExpanded
		configHook KafkaConnectToConsoleHook
		expected   *ClusterConnectorInfo
	}

	tests := []test{
		{
			name: "healthy",
			input: &connect.ListConnectorsResponseExpanded{
				Info: connect.ConnectorInfo{
					Name: "http-source-connector-wtue",
					Config: map[string]string{
						"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
						"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
						"http.request.url":                          "https://httpbin.org/uuid",
						"http.timer.catchup.interval.millis":        "30000",
						"http.timer.interval.millis":                "180000",
						"kafka.topic":                               "httpbin-input",
						"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
						"key.converter.schemas.enable":              "false",
						"name":                                      "http-source-connector-wtue",
						"topic.creation.default.partitions":         "1",
						"topic.creation.default.replication.factor": "1",
						"topic.creation.enable":                     "true",
						"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
						"value.converter.schemas.enable":            "false",
					},
					Tasks: []connect.ConnectorTaskID{
						{
							Connector: "http-source-connector-wtue",
							Task:      0,
						},
					},
					Type: "source",
				},
				Status: connect.ConnectorStateInfo{
					Name: "http-source-connector-wtue",
					Connector: struct {
						State    string `json:"state"`
						WorkerID string `json:"worker_id"`
						Trace    string `json:"trace,omitempty"`
					}{
						State:    "RUNNING",
						WorkerID: "172.21.0.5:8083",
					},
					Tasks: []connect.TaskState{
						{
							ID:       0,
							State:    "RUNNING",
							WorkerID: "172.21.0.5:8083",
						},
					},
				},
			},
			expected: &ClusterConnectorInfo{
				Name:  "http-source-connector-wtue",
				Class: "com.github.castorm.kafka.connect.http.HttpSourceConnector",
				Config: map[string]string{
					"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
					"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
					"http.request.url":                          "https://httpbin.org/uuid",
					"http.timer.catchup.interval.millis":        "30000",
					"http.timer.interval.millis":                "180000",
					"kafka.topic":                               "httpbin-input",
					"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
					"key.converter.schemas.enable":              "false",
					"name":                                      "http-source-connector-wtue",
					"topic.creation.default.partitions":         "1",
					"topic.creation.default.replication.factor": "1",
					"topic.creation.enable":                     "true",
					"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
					"value.converter.schemas.enable":            "false",
				},
				WorkerID:     "172.21.0.5:8083",
				Type:         "source",
				Topic:        "httpbin-input",
				State:        connectorStateRunning,
				Status:       ConnectorStatusHealthy,
				TotalTasks:   1,
				RunningTasks: 1,
				Trace:        "",
				Errors:       []ClusterConnectorInfoError{},
				Tasks: []ClusterConnectorTaskInfo{
					{
						TaskID:   0,
						State:    connectorStateRunning,
						WorkerID: "172.21.0.5:8083",
					},
				},
			},
		},
		{
			name: "unhealthy - running and 0 tasks",
			input: &connect.ListConnectorsResponseExpanded{
				Info: connect.ConnectorInfo{
					Name: "http-source-connector-wtue",
					Config: map[string]string{
						"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
						"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
						"http.request.url":                          "https://httpbin.org/uuid",
						"http.timer.catchup.interval.millis":        "30000",
						"http.timer.interval.millis":                "180000",
						"kafka.topic":                               "httpbin-input",
						"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
						"key.converter.schemas.enable":              "false",
						"name":                                      "http-source-connector-wtue",
						"topic.creation.default.partitions":         "1",
						"topic.creation.default.replication.factor": "1",
						"topic.creation.enable":                     "true",
						"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
						"value.converter.schemas.enable":            "false",
					},
					Tasks: []connect.ConnectorTaskID{},
					Type:  "source",
				},
				Status: connect.ConnectorStateInfo{
					Name: "http-source-connector-wtue",
					Connector: struct {
						State    string `json:"state"`
						WorkerID string `json:"worker_id"`
						Trace    string `json:"trace,omitempty"`
					}{
						State:    "RUNNING",
						WorkerID: "172.21.0.5:8083",
					},
					Tasks: []connect.TaskState{},
				},
			},
			expected: &ClusterConnectorInfo{
				Name:  "http-source-connector-wtue",
				Class: "com.github.castorm.kafka.connect.http.HttpSourceConnector",
				Config: map[string]string{
					"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
					"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
					"http.request.url":                          "https://httpbin.org/uuid",
					"http.timer.catchup.interval.millis":        "30000",
					"http.timer.interval.millis":                "180000",
					"kafka.topic":                               "httpbin-input",
					"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
					"key.converter.schemas.enable":              "false",
					"name":                                      "http-source-connector-wtue",
					"topic.creation.default.partitions":         "1",
					"topic.creation.default.replication.factor": "1",
					"topic.creation.enable":                     "true",
					"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
					"value.converter.schemas.enable":            "false",
				},
				WorkerID:     "172.21.0.5:8083",
				Type:         "source",
				Topic:        "httpbin-input",
				State:        connectorStateRunning,
				Status:       ConnectorStatusUnhealthy,
				TotalTasks:   0,
				RunningTasks: 0,
				Trace:        "",
				Errors: []ClusterConnectorInfoError{
					{
						Type:    connectorErrorTypeError,
						Title:   "Connector http-source-connector-wtue is in unhealthy state.",
						Content: "Connector http-source-connector-wtue is in running state but has no tasks.",
					},
				},
				Tasks: []ClusterConnectorTaskInfo{},
			},
		},
		{
			name: "unhealthy - connector in failed state",
			input: &connect.ListConnectorsResponseExpanded{
				Info: connect.ConnectorInfo{
					Name: "http-source-connector-wtue",
					Config: map[string]string{
						"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
						"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
						"http.request.url":                          "https://httpbin.org/uuid",
						"http.timer.catchup.interval.millis":        "30000",
						"http.timer.interval.millis":                "180000",
						"kafka.topic":                               "httpbin-input",
						"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
						"key.converter.schemas.enable":              "false",
						"name":                                      "http-source-connector-wtue",
						"topic.creation.default.partitions":         "1",
						"topic.creation.default.replication.factor": "1",
						"topic.creation.enable":                     "true",
						"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
						"value.converter.schemas.enable":            "false",
					},
					Tasks: []connect.ConnectorTaskID{},
					Type:  "source",
				},
				Status: connect.ConnectorStateInfo{
					Name: "http-source-connector-wtue",
					Connector: struct {
						State    string `json:"state"`
						WorkerID string `json:"worker_id"`
						Trace    string `json:"trace,omitempty"`
					}{
						State:    "FAILED",
						WorkerID: "172.21.0.5:8083",
					},
					Tasks: []connect.TaskState{
						{
							ID:       0,
							State:    "RUNNING",
							WorkerID: "172.21.0.5:8083",
						},
					},
				},
			},
			expected: &ClusterConnectorInfo{
				Name:  "http-source-connector-wtue",
				Class: "com.github.castorm.kafka.connect.http.HttpSourceConnector",
				Config: map[string]string{
					"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
					"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
					"http.request.url":                          "https://httpbin.org/uuid",
					"http.timer.catchup.interval.millis":        "30000",
					"http.timer.interval.millis":                "180000",
					"kafka.topic":                               "httpbin-input",
					"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
					"key.converter.schemas.enable":              "false",
					"name":                                      "http-source-connector-wtue",
					"topic.creation.default.partitions":         "1",
					"topic.creation.default.replication.factor": "1",
					"topic.creation.enable":                     "true",
					"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
					"value.converter.schemas.enable":            "false",
				},
				WorkerID:     "172.21.0.5:8083",
				Type:         "source",
				Topic:        "httpbin-input",
				State:        connectorStateFailed,
				Status:       ConnectorStatusUnhealthy,
				TotalTasks:   1,
				RunningTasks: 1,
				Trace:        "",
				Errors: []ClusterConnectorInfoError{
					{
						Type:    connectorErrorTypeError,
						Title:   "Connector http-source-connector-wtue is in unhealthy state.",
						Content: "Connector http-source-connector-wtue is in failed state.",
					},
				},
				Tasks: []ClusterConnectorTaskInfo{
					{
						TaskID:   0,
						State:    connectorStateRunning,
						WorkerID: "172.21.0.5:8083",
					},
				},
			},
		},
		{
			name: "unhealthy - running and all failed tasks",
			input: &connect.ListConnectorsResponseExpanded{
				Info: connect.ConnectorInfo{
					Name: "http-source-connector-wtue",
					Config: map[string]string{
						"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
						"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
						"http.request.url":                          "https://httpbin.org/uuid",
						"http.timer.catchup.interval.millis":        "30000",
						"http.timer.interval.millis":                "180000",
						"kafka.topic":                               "httpbin-input",
						"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
						"key.converter.schemas.enable":              "false",
						"name":                                      "http-source-connector-wtue",
						"topic.creation.default.partitions":         "1",
						"topic.creation.default.replication.factor": "1",
						"topic.creation.enable":                     "true",
						"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
						"value.converter.schemas.enable":            "false",
					},
					Tasks: []connect.ConnectorTaskID{
						{
							Connector: "http-source-connector-wtue",
							Task:      0,
						},
					},
					Type: "source",
				},
				Status: connect.ConnectorStateInfo{
					Name: "http-source-connector-wtue",
					Connector: struct {
						State    string `json:"state"`
						WorkerID string `json:"worker_id"`
						Trace    string `json:"trace,omitempty"`
					}{
						State:    "RUNNING",
						WorkerID: "172.21.0.5:8083",
					},
					Tasks: []connect.TaskState{
						{
							ID:       0,
							State:    "FAILED",
							WorkerID: "172.21.0.5:8083",
						},
						{
							ID:       1,
							State:    "FAILED",
							WorkerID: "172.21.0.5:8083",
						},
						{
							ID:       2,
							State:    "FAILED",
							WorkerID: "172.21.0.5:8083",
						},
					},
				},
			},
			expected: &ClusterConnectorInfo{
				Name:  "http-source-connector-wtue",
				Class: "com.github.castorm.kafka.connect.http.HttpSourceConnector",
				Config: map[string]string{
					"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
					"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
					"http.request.url":                          "https://httpbin.org/uuid",
					"http.timer.catchup.interval.millis":        "30000",
					"http.timer.interval.millis":                "180000",
					"kafka.topic":                               "httpbin-input",
					"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
					"key.converter.schemas.enable":              "false",
					"name":                                      "http-source-connector-wtue",
					"topic.creation.default.partitions":         "1",
					"topic.creation.default.replication.factor": "1",
					"topic.creation.enable":                     "true",
					"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
					"value.converter.schemas.enable":            "false",
				},
				WorkerID:     "172.21.0.5:8083",
				Type:         "source",
				Topic:        "httpbin-input",
				State:        connectorStateRunning,
				Status:       ConnectorStatusUnhealthy,
				TotalTasks:   3,
				RunningTasks: 0,
				Trace:        "",
				Errors: []ClusterConnectorInfoError{
					{
						Type:    connectorErrorTypeError,
						Title:   "Connector http-source-connector-wtue is in unhealthy state.",
						Content: "Connector http-source-connector-wtue is in running state. All tasks are in failed state.",
					},
					{
						Type:    connectorErrorTypeError,
						Title:   "Connector http-source-connector-wtue Task 0 is in failed state.",
						Content: "Connector http-source-connector-wtue Task 0 is in failed state.",
					},
					{
						Type:    connectorErrorTypeError,
						Title:   "Connector http-source-connector-wtue Task 1 is in failed state.",
						Content: "Connector http-source-connector-wtue Task 1 is in failed state.",
					},
					{
						Type:    connectorErrorTypeError,
						Title:   "Connector http-source-connector-wtue Task 2 is in failed state.",
						Content: "Connector http-source-connector-wtue Task 2 is in failed state.",
					},
				},
				Tasks: []ClusterConnectorTaskInfo{
					{
						TaskID:   0,
						State:    connectorStateFailed,
						WorkerID: "172.21.0.5:8083",
					},
					{
						TaskID:   1,
						State:    connectorStateFailed,
						WorkerID: "172.21.0.5:8083",
					},
					{
						TaskID:   2,
						State:    connectorStateFailed,
						WorkerID: "172.21.0.5:8083",
					},
				},
			},
		},
		{
			name: "paused - paused and all failed tasks",
			input: &connect.ListConnectorsResponseExpanded{
				Info: connect.ConnectorInfo{
					Name: "http-source-connector-wtue",
					Config: map[string]string{
						"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
						"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
						"http.request.url":                          "https://httpbin.org/uuid",
						"http.timer.catchup.interval.millis":        "30000",
						"http.timer.interval.millis":                "180000",
						"kafka.topic":                               "httpbin-input",
						"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
						"key.converter.schemas.enable":              "false",
						"name":                                      "http-source-connector-wtue",
						"topic.creation.default.partitions":         "1",
						"topic.creation.default.replication.factor": "1",
						"topic.creation.enable":                     "true",
						"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
						"value.converter.schemas.enable":            "false",
					},
					Tasks: []connect.ConnectorTaskID{
						{
							Connector: "http-source-connector-wtue",
							Task:      0,
						},
					},
					Type: "source",
				},
				Status: connect.ConnectorStateInfo{
					Name: "http-source-connector-wtue",
					Connector: struct {
						State    string `json:"state"`
						WorkerID string `json:"worker_id"`
						Trace    string `json:"trace,omitempty"`
					}{
						State:    "PAUSED",
						WorkerID: "172.21.0.5:8083",
					},
					Tasks: []connect.TaskState{
						{
							ID:       0,
							State:    "FAILED",
							WorkerID: "172.21.0.5:8083",
						},
						{
							ID:       1,
							State:    "FAILED",
							WorkerID: "172.21.0.5:8083",
						},
						{
							ID:       2,
							State:    "FAILED",
							WorkerID: "172.21.0.5:8083",
						},
					},
				},
			},
			expected: &ClusterConnectorInfo{
				Name:  "http-source-connector-wtue",
				Class: "com.github.castorm.kafka.connect.http.HttpSourceConnector",
				Config: map[string]string{
					"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
					"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
					"http.request.url":                          "https://httpbin.org/uuid",
					"http.timer.catchup.interval.millis":        "30000",
					"http.timer.interval.millis":                "180000",
					"kafka.topic":                               "httpbin-input",
					"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
					"key.converter.schemas.enable":              "false",
					"name":                                      "http-source-connector-wtue",
					"topic.creation.default.partitions":         "1",
					"topic.creation.default.replication.factor": "1",
					"topic.creation.enable":                     "true",
					"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
					"value.converter.schemas.enable":            "false",
				},
				WorkerID:     "172.21.0.5:8083",
				Type:         "source",
				Topic:        "httpbin-input",
				State:        connectorStatePaused,
				Status:       ConnectorStatusPaused,
				TotalTasks:   3,
				RunningTasks: 0,
				Trace:        "",
				Errors: []ClusterConnectorInfoError{
					{
						Type:    connectorErrorTypeError,
						Title:   "Connector http-source-connector-wtue Task 0 is in failed state.",
						Content: "Connector http-source-connector-wtue Task 0 is in failed state.",
					},
					{
						Type:    connectorErrorTypeError,
						Title:   "Connector http-source-connector-wtue Task 1 is in failed state.",
						Content: "Connector http-source-connector-wtue Task 1 is in failed state.",
					},
					{
						Type:    connectorErrorTypeError,
						Title:   "Connector http-source-connector-wtue Task 2 is in failed state.",
						Content: "Connector http-source-connector-wtue Task 2 is in failed state.",
					},
				},
				Tasks: []ClusterConnectorTaskInfo{
					{
						TaskID:   0,
						State:    connectorStateFailed,
						WorkerID: "172.21.0.5:8083",
					},
					{
						TaskID:   1,
						State:    connectorStateFailed,
						WorkerID: "172.21.0.5:8083",
					},
					{
						TaskID:   2,
						State:    connectorStateFailed,
						WorkerID: "172.21.0.5:8083",
					},
				},
			},
		},
		{
			name: "degraded - connector running, has tasks and 1 task is in failed state",
			input: &connect.ListConnectorsResponseExpanded{
				Info: connect.ConnectorInfo{
					Name: "http-source-connector-wtue",
					Config: map[string]string{
						"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
						"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
						"http.request.url":                          "https://httpbin.org/uuid",
						"http.timer.catchup.interval.millis":        "30000",
						"http.timer.interval.millis":                "180000",
						"kafka.topic":                               "httpbin-input",
						"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
						"key.converter.schemas.enable":              "false",
						"name":                                      "http-source-connector-wtue",
						"topic.creation.default.partitions":         "1",
						"topic.creation.default.replication.factor": "1",
						"topic.creation.enable":                     "true",
						"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
						"value.converter.schemas.enable":            "false",
					},
					Tasks: []connect.ConnectorTaskID{
						{
							Connector: "http-source-connector-wtue",
							Task:      0,
						},
					},
					Type: "source",
				},
				Status: connect.ConnectorStateInfo{
					Name: "http-source-connector-wtue",
					Connector: struct {
						State    string `json:"state"`
						WorkerID string `json:"worker_id"`
						Trace    string `json:"trace,omitempty"`
					}{
						State:    "RUNNING",
						WorkerID: "172.21.0.5:8083",
					},
					Tasks: []connect.TaskState{
						{
							ID:       0,
							State:    "RUNNING",
							WorkerID: "172.21.0.5:8083",
						},
						{
							ID:       1,
							State:    "FAILED",
							WorkerID: "172.21.0.5:8083",
						},
						{
							ID:       2,
							State:    "RUNNING",
							WorkerID: "172.21.0.5:8083",
						},
					},
				},
			},
			expected: &ClusterConnectorInfo{
				Name:  "http-source-connector-wtue",
				Class: "com.github.castorm.kafka.connect.http.HttpSourceConnector",
				Config: map[string]string{
					"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
					"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
					"http.request.url":                          "https://httpbin.org/uuid",
					"http.timer.catchup.interval.millis":        "30000",
					"http.timer.interval.millis":                "180000",
					"kafka.topic":                               "httpbin-input",
					"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
					"key.converter.schemas.enable":              "false",
					"name":                                      "http-source-connector-wtue",
					"topic.creation.default.partitions":         "1",
					"topic.creation.default.replication.factor": "1",
					"topic.creation.enable":                     "true",
					"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
					"value.converter.schemas.enable":            "false",
				},
				WorkerID:     "172.21.0.5:8083",
				Type:         "source",
				Topic:        "httpbin-input",
				State:        connectorStateRunning,
				Status:       ConnectorStatusDegraded,
				TotalTasks:   3,
				RunningTasks: 2,
				Trace:        "",
				Errors: []ClusterConnectorInfoError{
					{
						Type:    connectorErrorTypeError,
						Title:   "Connector http-source-connector-wtue is in degraded state.",
						Content: "Connector http-source-connector-wtue is in running state but has 1 / 3 failed tasks.",
					},
					{
						Type:    connectorErrorTypeError,
						Title:   "Connector http-source-connector-wtue Task 1 is in failed state.",
						Content: "Connector http-source-connector-wtue Task 1 is in failed state.",
					},
				},
				Tasks: []ClusterConnectorTaskInfo{
					{
						TaskID:   0,
						State:    connectorStateRunning,
						WorkerID: "172.21.0.5:8083",
					},
					{
						TaskID:   1,
						State:    connectorStateFailed,
						WorkerID: "172.21.0.5:8083",
					},
					{
						TaskID:   2,
						State:    connectorStateRunning,
						WorkerID: "172.21.0.5:8083",
					},
				},
			},
		},
		{
			name: "paused - connector paused, has tasks but 1 task failed",
			input: &connect.ListConnectorsResponseExpanded{
				Info: connect.ConnectorInfo{
					Name: "http-source-connector-wtue",
					Config: map[string]string{
						"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
						"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
						"http.request.url":                          "https://httpbin.org/uuid",
						"http.timer.catchup.interval.millis":        "30000",
						"http.timer.interval.millis":                "180000",
						"kafka.topic":                               "httpbin-input",
						"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
						"key.converter.schemas.enable":              "false",
						"name":                                      "http-source-connector-wtue",
						"topic.creation.default.partitions":         "1",
						"topic.creation.default.replication.factor": "1",
						"topic.creation.enable":                     "true",
						"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
						"value.converter.schemas.enable":            "false",
					},
					Tasks: []connect.ConnectorTaskID{
						{
							Connector: "http-source-connector-wtue",
							Task:      0,
						},
					},
					Type: "source",
				},
				Status: connect.ConnectorStateInfo{
					Name: "http-source-connector-wtue",
					Connector: struct {
						State    string `json:"state"`
						WorkerID string `json:"worker_id"`
						Trace    string `json:"trace,omitempty"`
					}{
						State:    "PAUSED",
						WorkerID: "172.21.0.5:8083",
					},
					Tasks: []connect.TaskState{
						{
							ID:       0,
							State:    "PAUSED",
							WorkerID: "172.21.0.5:8083",
						},
						{
							ID:       1,
							State:    "PAUSED",
							WorkerID: "172.21.0.5:8083",
						},
						{
							ID:       2,
							State:    "FAILED",
							WorkerID: "172.21.0.5:8083",
						},
					},
				},
			},
			expected: &ClusterConnectorInfo{
				Name:  "http-source-connector-wtue",
				Class: "com.github.castorm.kafka.connect.http.HttpSourceConnector",
				Config: map[string]string{
					"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
					"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
					"http.request.url":                          "https://httpbin.org/uuid",
					"http.timer.catchup.interval.millis":        "30000",
					"http.timer.interval.millis":                "180000",
					"kafka.topic":                               "httpbin-input",
					"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
					"key.converter.schemas.enable":              "false",
					"name":                                      "http-source-connector-wtue",
					"topic.creation.default.partitions":         "1",
					"topic.creation.default.replication.factor": "1",
					"topic.creation.enable":                     "true",
					"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
					"value.converter.schemas.enable":            "false",
				},
				WorkerID:     "172.21.0.5:8083",
				Type:         "source",
				Topic:        "httpbin-input",
				State:        connectorStatePaused,
				Status:       ConnectorStatusPaused,
				TotalTasks:   3,
				RunningTasks: 0,
				Trace:        "",
				Errors: []ClusterConnectorInfoError{
					{
						Type:    connectorErrorTypeError,
						Title:   "Connector http-source-connector-wtue Task 2 is in failed state.",
						Content: "Connector http-source-connector-wtue Task 2 is in failed state.",
					},
				},
				Tasks: []ClusterConnectorTaskInfo{
					{
						TaskID:   0,
						State:    connectorStatePaused,
						WorkerID: "172.21.0.5:8083",
					},
					{
						TaskID:   1,
						State:    connectorStatePaused,
						WorkerID: "172.21.0.5:8083",
					},
					{
						TaskID:   2,
						State:    connectorStateFailed,
						WorkerID: "172.21.0.5:8083",
					},
				},
			},
		},
		{
			name: "paused - connector paused, all tasks paused",
			input: &connect.ListConnectorsResponseExpanded{
				Info: connect.ConnectorInfo{
					Name: "http-source-connector-wtue",
					Config: map[string]string{
						"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
						"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
						"http.request.url":                          "https://httpbin.org/uuid",
						"http.timer.catchup.interval.millis":        "30000",
						"http.timer.interval.millis":                "180000",
						"kafka.topic":                               "httpbin-input",
						"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
						"key.converter.schemas.enable":              "false",
						"name":                                      "http-source-connector-wtue",
						"topic.creation.default.partitions":         "1",
						"topic.creation.default.replication.factor": "1",
						"topic.creation.enable":                     "true",
						"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
						"value.converter.schemas.enable":            "false",
					},
					Tasks: []connect.ConnectorTaskID{
						{
							Connector: "http-source-connector-wtue",
							Task:      0,
						},
					},
					Type: "source",
				},
				Status: connect.ConnectorStateInfo{
					Name: "http-source-connector-wtue",
					Connector: struct {
						State    string `json:"state"`
						WorkerID string `json:"worker_id"`
						Trace    string `json:"trace,omitempty"`
					}{
						State:    "PAUSED",
						WorkerID: "172.21.0.5:8083",
					},
					Tasks: []connect.TaskState{
						{
							ID:       0,
							State:    "PAUSED",
							WorkerID: "172.21.0.5:8083",
						},
						{
							ID:       1,
							State:    "PAUSED",
							WorkerID: "172.21.0.5:8083",
						},
						{
							ID:       2,
							State:    "PAUSED",
							WorkerID: "172.21.0.5:8083",
						},
					},
				},
			},
			expected: &ClusterConnectorInfo{
				Name:  "http-source-connector-wtue",
				Class: "com.github.castorm.kafka.connect.http.HttpSourceConnector",
				Config: map[string]string{
					"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
					"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
					"http.request.url":                          "https://httpbin.org/uuid",
					"http.timer.catchup.interval.millis":        "30000",
					"http.timer.interval.millis":                "180000",
					"kafka.topic":                               "httpbin-input",
					"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
					"key.converter.schemas.enable":              "false",
					"name":                                      "http-source-connector-wtue",
					"topic.creation.default.partitions":         "1",
					"topic.creation.default.replication.factor": "1",
					"topic.creation.enable":                     "true",
					"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
					"value.converter.schemas.enable":            "false",
				},
				WorkerID:     "172.21.0.5:8083",
				Type:         "source",
				Topic:        "httpbin-input",
				State:        connectorStatePaused,
				Status:       ConnectorStatusPaused,
				TotalTasks:   3,
				RunningTasks: 0,
				Trace:        "",
				Errors:       []ClusterConnectorInfoError{},
				Tasks: []ClusterConnectorTaskInfo{
					{
						TaskID:   0,
						State:    connectorStatePaused,
						WorkerID: "172.21.0.5:8083",
					},
					{
						TaskID:   1,
						State:    connectorStatePaused,
						WorkerID: "172.21.0.5:8083",
					},
					{
						TaskID:   2,
						State:    connectorStatePaused,
						WorkerID: "172.21.0.5:8083",
					},
				},
			},
		},
		{
			name: "paused - connector paused, no tasks",
			input: &connect.ListConnectorsResponseExpanded{
				Info: connect.ConnectorInfo{
					Name: "http-source-connector-wtue",
					Config: map[string]string{
						"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
						"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
						"http.request.url":                          "https://httpbin.org/uuid",
						"http.timer.catchup.interval.millis":        "30000",
						"http.timer.interval.millis":                "180000",
						"kafka.topic":                               "httpbin-input",
						"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
						"key.converter.schemas.enable":              "false",
						"name":                                      "http-source-connector-wtue",
						"topic.creation.default.partitions":         "1",
						"topic.creation.default.replication.factor": "1",
						"topic.creation.enable":                     "true",
						"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
						"value.converter.schemas.enable":            "false",
					},
					Tasks: []connect.ConnectorTaskID{
						{
							Connector: "http-source-connector-wtue",
							Task:      0,
						},
					},
					Type: "source",
				},
				Status: connect.ConnectorStateInfo{
					Name: "http-source-connector-wtue",
					Connector: struct {
						State    string `json:"state"`
						WorkerID string `json:"worker_id"`
						Trace    string `json:"trace,omitempty"`
					}{
						State:    "PAUSED",
						WorkerID: "172.21.0.5:8083",
					},
					Tasks: []connect.TaskState{},
				},
			},
			expected: &ClusterConnectorInfo{
				Name:  "http-source-connector-wtue",
				Class: "com.github.castorm.kafka.connect.http.HttpSourceConnector",
				Config: map[string]string{
					"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
					"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
					"http.request.url":                          "https://httpbin.org/uuid",
					"http.timer.catchup.interval.millis":        "30000",
					"http.timer.interval.millis":                "180000",
					"kafka.topic":                               "httpbin-input",
					"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
					"key.converter.schemas.enable":              "false",
					"name":                                      "http-source-connector-wtue",
					"topic.creation.default.partitions":         "1",
					"topic.creation.default.replication.factor": "1",
					"topic.creation.enable":                     "true",
					"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
					"value.converter.schemas.enable":            "false",
				},
				WorkerID:     "172.21.0.5:8083",
				Type:         "source",
				Topic:        "httpbin-input",
				State:        connectorStatePaused,
				Status:       ConnectorStatusPaused,
				TotalTasks:   0,
				RunningTasks: 0,
				Trace:        "",
				Errors:       []ClusterConnectorInfoError{},
				Tasks:        []ClusterConnectorTaskInfo{},
			},
		},
		{
			name: "stopped - connector stopped, no tasks",
			input: &connect.ListConnectorsResponseExpanded{
				Info: connect.ConnectorInfo{
					Name: "http-source-connector-wtue",
					Config: map[string]string{
						"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
						"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
						"http.request.url":                          "https://httpbin.org/uuid",
						"http.timer.catchup.interval.millis":        "30000",
						"http.timer.interval.millis":                "180000",
						"kafka.topic":                               "httpbin-input",
						"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
						"key.converter.schemas.enable":              "false",
						"name":                                      "http-source-connector-wtue",
						"topic.creation.default.partitions":         "1",
						"topic.creation.default.replication.factor": "1",
						"topic.creation.enable":                     "true",
						"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
						"value.converter.schemas.enable":            "false",
					},
					Tasks: []connect.ConnectorTaskID{
						{
							Connector: "http-source-connector-wtue",
							Task:      0,
						},
					},
					Type: "source",
				},
				Status: connect.ConnectorStateInfo{
					Name: "http-source-connector-wtue",
					Connector: struct {
						State    string `json:"state"`
						WorkerID string `json:"worker_id"`
						Trace    string `json:"trace,omitempty"`
					}{
						State:    "STOPPED",
						WorkerID: "172.21.0.5:8083",
					},
					Tasks: []connect.TaskState{},
				},
			},
			expected: &ClusterConnectorInfo{
				Name:  "http-source-connector-wtue",
				Class: "com.github.castorm.kafka.connect.http.HttpSourceConnector",
				Config: map[string]string{
					"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
					"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
					"http.request.url":                          "https://httpbin.org/uuid",
					"http.timer.catchup.interval.millis":        "30000",
					"http.timer.interval.millis":                "180000",
					"kafka.topic":                               "httpbin-input",
					"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
					"key.converter.schemas.enable":              "false",
					"name":                                      "http-source-connector-wtue",
					"topic.creation.default.partitions":         "1",
					"topic.creation.default.replication.factor": "1",
					"topic.creation.enable":                     "true",
					"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
					"value.converter.schemas.enable":            "false",
				},
				WorkerID:     "172.21.0.5:8083",
				Type:         "source",
				Topic:        "httpbin-input",
				State:        connectorStateStopped,
				Status:       ConnectorStatusStopped,
				TotalTasks:   0,
				RunningTasks: 0,
				Trace:        "",
				Errors:       []ClusterConnectorInfoError{},
				Tasks:        []ClusterConnectorTaskInfo{},
			},
		},
		{
			name: "restarting - connector restarting, > 0 tasks",
			input: &connect.ListConnectorsResponseExpanded{
				Info: connect.ConnectorInfo{
					Name: "http-source-connector-wtue",
					Config: map[string]string{
						"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
						"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
						"http.request.url":                          "https://httpbin.org/uuid",
						"http.timer.catchup.interval.millis":        "30000",
						"http.timer.interval.millis":                "180000",
						"kafka.topic":                               "httpbin-input",
						"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
						"key.converter.schemas.enable":              "false",
						"name":                                      "http-source-connector-wtue",
						"topic.creation.default.partitions":         "1",
						"topic.creation.default.replication.factor": "1",
						"topic.creation.enable":                     "true",
						"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
						"value.converter.schemas.enable":            "false",
					},
					Tasks: []connect.ConnectorTaskID{
						{
							Connector: "http-source-connector-wtue",
							Task:      0,
						},
					},
					Type: "source",
				},
				Status: connect.ConnectorStateInfo{
					Name: "http-source-connector-wtue",
					Connector: struct {
						State    string `json:"state"`
						WorkerID string `json:"worker_id"`
						Trace    string `json:"trace,omitempty"`
					}{
						State:    "RESTARTING",
						WorkerID: "172.21.0.5:8083",
					},
					Tasks: []connect.TaskState{
						{
							ID:       0,
							State:    "RUNNING",
							WorkerID: "172.21.0.5:8083",
						},
						{
							ID:       1,
							State:    "PAUSED",
							WorkerID: "172.21.0.5:8083",
						},
						{
							ID:       2,
							State:    "RUNNING",
							WorkerID: "172.21.0.5:8083",
						},
					},
				},
			},
			expected: &ClusterConnectorInfo{
				Name:  "http-source-connector-wtue",
				Class: "com.github.castorm.kafka.connect.http.HttpSourceConnector",
				Config: map[string]string{
					"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
					"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
					"http.request.url":                          "https://httpbin.org/uuid",
					"http.timer.catchup.interval.millis":        "30000",
					"http.timer.interval.millis":                "180000",
					"kafka.topic":                               "httpbin-input",
					"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
					"key.converter.schemas.enable":              "false",
					"name":                                      "http-source-connector-wtue",
					"topic.creation.default.partitions":         "1",
					"topic.creation.default.replication.factor": "1",
					"topic.creation.enable":                     "true",
					"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
					"value.converter.schemas.enable":            "false",
				},
				WorkerID:     "172.21.0.5:8083",
				Type:         "source",
				Topic:        "httpbin-input",
				State:        connectorStateRestarting,
				Status:       ConnectorStatusRestarting,
				TotalTasks:   3,
				RunningTasks: 2,
				Trace:        "",
				Errors:       []ClusterConnectorInfoError{},
				Tasks: []ClusterConnectorTaskInfo{
					{
						TaskID:   0,
						State:    connectorStateRunning,
						WorkerID: "172.21.0.5:8083",
					},
					{
						TaskID:   1,
						State:    connectorStatePaused,
						WorkerID: "172.21.0.5:8083",
					},
					{
						TaskID:   2,
						State:    connectorStateRunning,
						WorkerID: "172.21.0.5:8083",
					},
				},
			},
		},
		{
			name: "restarting - connector running, > 0 tasks, 1 task is restarting",
			input: &connect.ListConnectorsResponseExpanded{
				Info: connect.ConnectorInfo{
					Name: "http-source-connector-wtue",
					Config: map[string]string{
						"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
						"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
						"http.request.url":                          "https://httpbin.org/uuid",
						"http.timer.catchup.interval.millis":        "30000",
						"http.timer.interval.millis":                "180000",
						"kafka.topic":                               "httpbin-input",
						"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
						"key.converter.schemas.enable":              "false",
						"name":                                      "http-source-connector-wtue",
						"topic.creation.default.partitions":         "1",
						"topic.creation.default.replication.factor": "1",
						"topic.creation.enable":                     "true",
						"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
						"value.converter.schemas.enable":            "false",
					},
					Tasks: []connect.ConnectorTaskID{
						{
							Connector: "http-source-connector-wtue",
							Task:      0,
						},
					},
					Type: "source",
				},
				Status: connect.ConnectorStateInfo{
					Name: "http-source-connector-wtue",
					Connector: struct {
						State    string `json:"state"`
						WorkerID string `json:"worker_id"`
						Trace    string `json:"trace,omitempty"`
					}{
						State:    "RUNNING",
						WorkerID: "172.21.0.5:8083",
					},
					Tasks: []connect.TaskState{
						{
							ID:       0,
							State:    "RUNNING",
							WorkerID: "172.21.0.5:8083",
						},
						{
							ID:       1,
							State:    "RESTARTING",
							WorkerID: "172.21.0.5:8083",
						},
						{
							ID:       2,
							State:    "RUNNING",
							WorkerID: "172.21.0.5:8083",
						},
					},
				},
			},
			expected: &ClusterConnectorInfo{
				Name:  "http-source-connector-wtue",
				Class: "com.github.castorm.kafka.connect.http.HttpSourceConnector",
				Config: map[string]string{
					"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
					"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
					"http.request.url":                          "https://httpbin.org/uuid",
					"http.timer.catchup.interval.millis":        "30000",
					"http.timer.interval.millis":                "180000",
					"kafka.topic":                               "httpbin-input",
					"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
					"key.converter.schemas.enable":              "false",
					"name":                                      "http-source-connector-wtue",
					"topic.creation.default.partitions":         "1",
					"topic.creation.default.replication.factor": "1",
					"topic.creation.enable":                     "true",
					"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
					"value.converter.schemas.enable":            "false",
				},
				WorkerID:     "172.21.0.5:8083",
				Type:         "source",
				Topic:        "httpbin-input",
				State:        connectorStateRunning,
				Status:       ConnectorStatusRestarting,
				TotalTasks:   3,
				RunningTasks: 2,
				Trace:        "",
				Errors:       []ClusterConnectorInfoError{},
				Tasks: []ClusterConnectorTaskInfo{
					{
						TaskID:   0,
						State:    connectorStateRunning,
						WorkerID: "172.21.0.5:8083",
					},
					{
						TaskID:   1,
						State:    connectorStateRestarting,
						WorkerID: "172.21.0.5:8083",
					},
					{
						TaskID:   2,
						State:    connectorStateRunning,
						WorkerID: "172.21.0.5:8083",
					},
				},
			},
		},
		{
			name: "unassigned - connector is unassigned",
			input: &connect.ListConnectorsResponseExpanded{
				Info: connect.ConnectorInfo{
					Name: "http-source-connector-wtue",
					Config: map[string]string{
						"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
						"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
						"http.request.url":                          "https://httpbin.org/uuid",
						"http.timer.catchup.interval.millis":        "30000",
						"http.timer.interval.millis":                "180000",
						"kafka.topic":                               "httpbin-input",
						"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
						"key.converter.schemas.enable":              "false",
						"name":                                      "http-source-connector-wtue",
						"topic.creation.default.partitions":         "1",
						"topic.creation.default.replication.factor": "1",
						"topic.creation.enable":                     "true",
						"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
						"value.converter.schemas.enable":            "false",
					},
					Tasks: []connect.ConnectorTaskID{
						{
							Connector: "http-source-connector-wtue",
							Task:      0,
						},
					},
					Type: "source",
				},
				Status: connect.ConnectorStateInfo{
					Name: "http-source-connector-wtue",
					Connector: struct {
						State    string `json:"state"`
						WorkerID string `json:"worker_id"`
						Trace    string `json:"trace,omitempty"`
					}{
						State:    "UNASSIGNED",
						WorkerID: "172.21.0.5:8083",
					},
					Tasks: []connect.TaskState{
						{
							ID:       0,
							State:    "UNASSIGNED",
							WorkerID: "172.21.0.5:8083",
						},
					},
				},
			},
			expected: &ClusterConnectorInfo{
				Name:  "http-source-connector-wtue",
				Class: "com.github.castorm.kafka.connect.http.HttpSourceConnector",
				Config: map[string]string{
					"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
					"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
					"http.request.url":                          "https://httpbin.org/uuid",
					"http.timer.catchup.interval.millis":        "30000",
					"http.timer.interval.millis":                "180000",
					"kafka.topic":                               "httpbin-input",
					"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
					"key.converter.schemas.enable":              "false",
					"name":                                      "http-source-connector-wtue",
					"topic.creation.default.partitions":         "1",
					"topic.creation.default.replication.factor": "1",
					"topic.creation.enable":                     "true",
					"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
					"value.converter.schemas.enable":            "false",
				},
				WorkerID:     "172.21.0.5:8083",
				Type:         "source",
				Topic:        "httpbin-input",
				State:        connectorStateUnassigned,
				Status:       ConnectorStatusUnassigned,
				TotalTasks:   1,
				RunningTasks: 0,
				Trace:        "",
				Errors:       []ClusterConnectorInfoError{},
				Tasks: []ClusterConnectorTaskInfo{
					{
						TaskID:   0,
						State:    connectorStateUnassigned,
						WorkerID: "172.21.0.5:8083",
					},
				},
			},
		},
		{
			name: "unassigned - connector is running, some tasks are unassigned",
			input: &connect.ListConnectorsResponseExpanded{
				Info: connect.ConnectorInfo{
					Name: "http-source-connector-wtue",
					Config: map[string]string{
						"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
						"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
						"http.request.url":                          "https://httpbin.org/uuid",
						"http.timer.catchup.interval.millis":        "30000",
						"http.timer.interval.millis":                "180000",
						"kafka.topic":                               "httpbin-input",
						"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
						"key.converter.schemas.enable":              "false",
						"name":                                      "http-source-connector-wtue",
						"topic.creation.default.partitions":         "1",
						"topic.creation.default.replication.factor": "1",
						"topic.creation.enable":                     "true",
						"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
						"value.converter.schemas.enable":            "false",
					},
					Tasks: []connect.ConnectorTaskID{
						{
							Connector: "http-source-connector-wtue",
							Task:      0,
						},
						{
							Connector: "http-source-connector-wtue",
							Task:      1,
						},
						{
							Connector: "http-source-connector-wtue",
							Task:      2,
						},
					},
					Type: "source",
				},
				Status: connect.ConnectorStateInfo{
					Name: "http-source-connector-wtue",
					Connector: struct {
						State    string `json:"state"`
						WorkerID string `json:"worker_id"`
						Trace    string `json:"trace,omitempty"`
					}{
						State:    "RUNNING",
						WorkerID: "172.21.0.5:8083",
					},
					Tasks: []connect.TaskState{
						{
							ID:       0,
							State:    "RUNNING",
							WorkerID: "172.21.0.5:8083",
						},
						{
							ID:       1,
							State:    "UNASSIGNED",
							WorkerID: "172.21.0.5:8083",
						},
						{
							ID:       2,
							State:    "RUNNING",
							WorkerID: "172.21.0.5:8083",
						},
					},
				},
			},
			expected: &ClusterConnectorInfo{
				Name:  "http-source-connector-wtue",
				Class: "com.github.castorm.kafka.connect.http.HttpSourceConnector",
				Config: map[string]string{
					"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
					"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
					"http.request.url":                          "https://httpbin.org/uuid",
					"http.timer.catchup.interval.millis":        "30000",
					"http.timer.interval.millis":                "180000",
					"kafka.topic":                               "httpbin-input",
					"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
					"key.converter.schemas.enable":              "false",
					"name":                                      "http-source-connector-wtue",
					"topic.creation.default.partitions":         "1",
					"topic.creation.default.replication.factor": "1",
					"topic.creation.enable":                     "true",
					"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
					"value.converter.schemas.enable":            "false",
				},
				WorkerID:     "172.21.0.5:8083",
				Type:         "source",
				Topic:        "httpbin-input",
				State:        connectorStateRunning,
				Status:       ConnectorStatusUnassigned,
				TotalTasks:   3,
				RunningTasks: 2,
				Trace:        "",
				Errors:       []ClusterConnectorInfoError{},
				Tasks: []ClusterConnectorTaskInfo{
					{
						TaskID:   0,
						State:    connectorStateRunning,
						WorkerID: "172.21.0.5:8083",
					},
					{
						TaskID:   1,
						State:    connectorStateUnassigned,
						WorkerID: "172.21.0.5:8083",
					},
					{
						TaskID:   2,
						State:    connectorStateRunning,
						WorkerID: "172.21.0.5:8083",
					},
				},
			},
		},
		{
			name: "destroyed - connector is destroyed",
			input: &connect.ListConnectorsResponseExpanded{
				Info: connect.ConnectorInfo{
					Name: "http-source-connector-wtue",
					Config: map[string]string{
						"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
						"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
						"http.request.url":                          "https://httpbin.org/uuid",
						"http.timer.catchup.interval.millis":        "30000",
						"http.timer.interval.millis":                "180000",
						"kafka.topic":                               "httpbin-input",
						"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
						"key.converter.schemas.enable":              "false",
						"name":                                      "http-source-connector-wtue",
						"topic.creation.default.partitions":         "1",
						"topic.creation.default.replication.factor": "1",
						"topic.creation.enable":                     "true",
						"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
						"value.converter.schemas.enable":            "false",
					},
					Tasks: []connect.ConnectorTaskID{
						{
							Connector: "http-source-connector-wtue",
							Task:      0,
						},
					},
					Type: "source",
				},
				Status: connect.ConnectorStateInfo{
					Name: "http-source-connector-wtue",
					Connector: struct {
						State    string `json:"state"`
						WorkerID string `json:"worker_id"`
						Trace    string `json:"trace,omitempty"`
					}{
						State:    "DESTROYED",
						WorkerID: "172.21.0.5:8083",
					},
					Tasks: []connect.TaskState{
						{
							ID:       0,
							State:    "DESTROYED",
							WorkerID: "172.21.0.5:8083",
						},
					},
				},
			},
			expected: &ClusterConnectorInfo{
				Name:  "http-source-connector-wtue",
				Class: "com.github.castorm.kafka.connect.http.HttpSourceConnector",
				Config: map[string]string{
					"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
					"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
					"http.request.url":                          "https://httpbin.org/uuid",
					"http.timer.catchup.interval.millis":        "30000",
					"http.timer.interval.millis":                "180000",
					"kafka.topic":                               "httpbin-input",
					"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
					"key.converter.schemas.enable":              "false",
					"name":                                      "http-source-connector-wtue",
					"topic.creation.default.partitions":         "1",
					"topic.creation.default.replication.factor": "1",
					"topic.creation.enable":                     "true",
					"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
					"value.converter.schemas.enable":            "false",
				},
				WorkerID:     "172.21.0.5:8083",
				Type:         "source",
				Topic:        "httpbin-input",
				State:        connectorStateDestroyed,
				Status:       ConnectorStatusDestroyed,
				TotalTasks:   1,
				RunningTasks: 0,
				Trace:        "",
				Errors:       []ClusterConnectorInfoError{},
				Tasks: []ClusterConnectorTaskInfo{
					{
						TaskID:   0,
						State:    connectorStateDestroyed,
						WorkerID: "172.21.0.5:8083",
					},
				},
			},
		},
		{
			name: "unknown - connector is state we do not track",
			input: &connect.ListConnectorsResponseExpanded{
				Info: connect.ConnectorInfo{
					Name: "http-source-connector-wtue",
					Config: map[string]string{
						"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
						"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
						"http.request.url":                          "https://httpbin.org/uuid",
						"http.timer.catchup.interval.millis":        "30000",
						"http.timer.interval.millis":                "180000",
						"kafka.topic":                               "httpbin-input",
						"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
						"key.converter.schemas.enable":              "false",
						"name":                                      "http-source-connector-wtue",
						"topic.creation.default.partitions":         "1",
						"topic.creation.default.replication.factor": "1",
						"topic.creation.enable":                     "true",
						"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
						"value.converter.schemas.enable":            "false",
					},
					Tasks: []connect.ConnectorTaskID{
						{
							Connector: "http-source-connector-wtue",
							Task:      0,
						},
					},
					Type: "source",
				},
				Status: connect.ConnectorStateInfo{
					Name: "http-source-connector-wtue",
					Connector: struct {
						State    string `json:"state"`
						WorkerID string `json:"worker_id"`
						Trace    string `json:"trace,omitempty"`
					}{
						State:    "PROVISIONING", // some fictitious state
						WorkerID: "172.21.0.5:8083",
					},
					Tasks: []connect.TaskState{
						{
							ID:       0,
							State:    "RUNNING",
							WorkerID: "172.21.0.5:8083",
						},
					},
				},
			},
			expected: &ClusterConnectorInfo{
				Name:  "http-source-connector-wtue",
				Class: "com.github.castorm.kafka.connect.http.HttpSourceConnector",
				Config: map[string]string{
					"connector.class":                           "com.github.castorm.kafka.connect.http.HttpSourceConnector",
					"header.converter":                          "org.apache.kafka.connect.storage.SimpleHeaderConverter",
					"http.request.url":                          "https://httpbin.org/uuid",
					"http.timer.catchup.interval.millis":        "30000",
					"http.timer.interval.millis":                "180000",
					"kafka.topic":                               "httpbin-input",
					"key.converter":                             "org.apache.kafka.connect.json.JsonConverter",
					"key.converter.schemas.enable":              "false",
					"name":                                      "http-source-connector-wtue",
					"topic.creation.default.partitions":         "1",
					"topic.creation.default.replication.factor": "1",
					"topic.creation.enable":                     "true",
					"value.converter":                           "org.apache.kafka.connect.json.JsonConverter",
					"value.converter.schemas.enable":            "false",
				},
				WorkerID:     "172.21.0.5:8083",
				Type:         "source",
				Topic:        "httpbin-input",
				State:        "PROVISIONING",
				Status:       ConnectorStatusUnknown,
				TotalTasks:   1,
				RunningTasks: 1,
				Trace:        "",
				Errors:       []ClusterConnectorInfoError{},
				Tasks: []ClusterConnectorTaskInfo{
					{
						TaskID:   0,
						State:    connectorStateRunning,
						WorkerID: "172.21.0.5:8083",
					},
				},
			},
		},
		{
			name: "healthy - apply hook",
			configHook: func(_ string, configs map[string]string) map[string]string {
				result := make(map[string]string)
				for key, value := range configs {
					switch key {
					case "delete":
						// skip
					case "change":
						result[key] = value + ".changed"
					default:
						result[key] = value
					}
				}
				result["added"] = "added value"

				return result
			},
			input: &connect.ListConnectorsResponseExpanded{
				Info: connect.ConnectorInfo{
					Name: "mirror-source-connector-apfe",
					Config: map[string]string{
						"change":          "value1",
						"delete":          "value2",
						"leave":           "value3",
						"connector.class": "org.apache.kafka.connect.mirror.MirrorSourceConnector",
					},
					Tasks: []connect.ConnectorTaskID{
						{
							Connector: "mirror-source-connector-apfe",
							Task:      0,
						},
					},
					Type: "source",
				},
				Status: connect.ConnectorStateInfo{
					Name: "mirror-source-connector-apfe",
					Connector: struct {
						State    string `json:"state"`
						WorkerID string `json:"worker_id"`
						Trace    string `json:"trace,omitempty"`
					}{
						State:    "RUNNING",
						WorkerID: "172.21.0.5:8083",
					},
					Tasks: []connect.TaskState{
						{
							ID:       0,
							State:    "RUNNING",
							WorkerID: "172.21.0.5:8083",
						},
					},
				},
			},
			expected: &ClusterConnectorInfo{
				Name:  "mirror-source-connector-apfe",
				Class: "org.apache.kafka.connect.mirror.MirrorSourceConnector",
				Config: map[string]string{
					"change":          "value1.changed",
					"leave":           "value3",
					"connector.class": "org.apache.kafka.connect.mirror.MirrorSourceConnector",
					"added":           "added value",
				},
				WorkerID:     "172.21.0.5:8083",
				Type:         "source",
				Topic:        "unknown",
				State:        connectorStateRunning,
				Status:       ConnectorStatusHealthy,
				TotalTasks:   1,
				RunningTasks: 1,
				Trace:        "",
				Errors:       []ClusterConnectorInfoError{},
				Tasks: []ClusterConnectorTaskInfo{
					{
						TaskID:   0,
						State:    connectorStateRunning,
						WorkerID: "172.21.0.5:8083",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := connectorsResponseToClusterConnectorInfo(tc.configHook, tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
