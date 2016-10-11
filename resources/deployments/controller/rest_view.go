// Copyright 2016 Mender Software AS
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

package controller

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/mendersoftware/deployments/resources/deployments"
	"github.com/mendersoftware/deployments/utils/log"
)

type RESTView interface {
	RenderNoUpdateForDevice(w rest.ResponseWriter)
	RenderSuccessPost(w rest.ResponseWriter, r *rest.Request, id string)
	RenderSuccessGet(w rest.ResponseWriter, object interface{})
	RenderEmptySuccessResponse(w rest.ResponseWriter)
	RenderError(w rest.ResponseWriter, err error, status int, l *log.Logger)
	RenderErrorNotFound(w rest.ResponseWriter, l *log.Logger)
	RenderDeploymentLog(w rest.ResponseWriter, dlog deployments.DeploymentLog)
}
