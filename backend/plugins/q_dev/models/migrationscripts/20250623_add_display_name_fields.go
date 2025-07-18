/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package migrationscripts

import (
	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
)

var _ plugin.MigrationScript = (*addDisplayNameFields)(nil)

type addDisplayNameFields struct{}

func (*addDisplayNameFields) Up(basicRes context.BasicRes) errors.Error {
	db := basicRes.GetDal()

	// Add Identity Center fields to connections table
	// Ignore error if column already exists (MySQL error 1060)
	_ = db.Exec("ALTER TABLE _tool_q_dev_connections ADD COLUMN identity_store_id VARCHAR(255)")
	_ = db.Exec("ALTER TABLE _tool_q_dev_connections ADD COLUMN identity_store_region VARCHAR(255)")

	// Add display_name column to user_data table
	// Ignore error if column already exists (MySQL error 1060)
	_ = db.Exec("ALTER TABLE _tool_q_dev_user_data ADD COLUMN display_name VARCHAR(255)")

	// Add display_name column to user_metrics table
	// Ignore error if column already exists (MySQL error 1060)
	_ = db.Exec("ALTER TABLE _tool_q_dev_user_metrics ADD COLUMN display_name VARCHAR(255)")

	return nil
}

func (*addDisplayNameFields) Version() uint64 {
	return 20250623000001
}

func (*addDisplayNameFields) Name() string {
	return "add Identity Center fields to connections and display_name fields to user tables"
}
