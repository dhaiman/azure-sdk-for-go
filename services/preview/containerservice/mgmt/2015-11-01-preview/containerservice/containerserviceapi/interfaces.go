package containerserviceapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "context"
    "github.com/Azure/azure-sdk-for-go/services/preview/containerservice/mgmt/2015-11-01-preview/containerservice"
)

        // ClientAPI contains the set of methods on the Client type.
        type ClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, containerServiceName string, parameters containerservice.ContainerService) (result containerservice.CreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, containerServiceName string) (result containerservice.DeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, containerServiceName string) (result containerservice.ContainerService, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result containerservice.ListResult, err error)
        }

        var _ ClientAPI = (*containerservice.Client)(nil)
