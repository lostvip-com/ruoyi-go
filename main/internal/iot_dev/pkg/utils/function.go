/*******************************************************************************
 * Copyright 2017 Dell Inc.
 * Copyright (c) 2019 Intel Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/
package utils

import "fmt"

func InStringSlice(key string, keys []string) bool {
    ok := false
    for _, item := range keys {
        if key == item {
            return true
        }
    }
    return ok
}

func SliceStringUnique(strings []string) []string {
    temp := map[string]struct{}{}
    result := make([]string, 0, len(strings))
    for _, item := range strings {
        key := fmt.Sprint(item)
        if _, ok := temp[key]; !ok {
            temp[key] = struct{}{}
            result = append(result, item)
        }
    }
    return result
}

