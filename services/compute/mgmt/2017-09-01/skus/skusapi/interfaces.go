package skusapi

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
    "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2017-09-01/skus"
)

        // ResourceSkusClientAPI contains the set of methods on the ResourceSkusClient type.
        type ResourceSkusClientAPI interface {
            List(ctx context.Context) (result skus.ResourceSkusResultPage, err error)
        }

        var _ ResourceSkusClientAPI = (*skus.ResourceSkusClient)(nil)
