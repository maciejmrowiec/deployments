// Copyright 2017 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package mongo

import (
	"context"

	"github.com/mendersoftware/go-lib-micro/store"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/mendersoftware/deployments/resources/deployments"
)

// Database settings
const (
	// TODO: do we have any naming convention for mongo collections?
	CollectionDeviceDeploymentLogs = "devices.logs"
)

// Database keys
const (
	StorageKeyDeviceDeploymentLogMessages = "messages"
)

// DeviceDeploymentLogsStorage is a data layer for deployment logs based on MongoDB
type DeviceDeploymentLogsStorage struct {
	session *mgo.Session
}

func NewDeviceDeploymentLogsStorage(session *mgo.Session) *DeviceDeploymentLogsStorage {
	return &DeviceDeploymentLogsStorage{
		session: session,
	}
}

func (d *DeviceDeploymentLogsStorage) SaveDeviceDeploymentLog(ctx context.Context,
	log deployments.DeploymentLog) error {

	if err := log.Validate(); err != nil {
		return err
	}

	session := d.session.Copy()
	defer session.Close()

	query := bson.M{
		StorageKeyDeviceDeploymentDeviceId:     log.DeviceID,
		StorageKeyDeviceDeploymentDeploymentID: log.DeploymentID,
	}

	// update log messages
	// if the deployment log is already present than messages will be overwritten
	update := bson.M{
		"$set": bson.M{
			StorageKeyDeviceDeploymentLogMessages: log.Messages,
		},
	}
	if _, err := session.DB(store.DbFromContext(ctx, DatabaseName)).
		C(CollectionDeviceDeploymentLogs).Upsert(query, update); err != nil {
		return err
	}

	return nil
}

func (d *DeviceDeploymentLogsStorage) GetDeviceDeploymentLog(ctx context.Context,
	deviceID, deploymentID string) (*deployments.DeploymentLog, error) {

	session := d.session.Copy()
	defer session.Close()

	query := bson.M{
		StorageKeyDeviceDeploymentDeviceId:     deviceID,
		StorageKeyDeviceDeploymentDeploymentID: deploymentID,
	}

	var depl deployments.DeploymentLog
	if err := session.DB(store.DbFromContext(ctx, DatabaseName)).
		C(CollectionDeviceDeploymentLogs).Find(query).One(&depl); err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &depl, nil
}
