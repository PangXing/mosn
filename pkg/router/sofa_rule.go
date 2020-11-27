/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package router

import (
	"context"
	"mosn.io/api"
	"mosn.io/mosn/pkg/log"
	"mosn.io/mosn/pkg/types"
	"mosn.io/mosn/pkg/variable"
)

type SofaRouteRuleImpl struct {
	*RouteRuleImplBase
	matchValue string
}

func (srri *SofaRouteRuleImpl) PathMatchCriterion() api.PathMatchCriterion {
	return srri
}

func (srri *SofaRouteRuleImpl) RouteRule() api.RouteRule {
	return srri
}

func (srri *SofaRouteRuleImpl) Matcher() string {
	return srri.matchValue
}

func (srri *SofaRouteRuleImpl) MatchType() api.PathMatchType {
	return api.SofaHeader
}

func (srri *SofaRouteRuleImpl) FinalizeRequestHeaders(headers api.HeaderMap, requestInfo api.RequestInfo) {
}

func (srri *SofaRouteRuleImpl) Match(ctx context.Context, headers api.HeaderMap) api.Route {
	value, err := variable.GetValueFromVariableAndLegacyHeader(ctx, headers, types.SofaRouteMatchKey, false)
	if value != nil {
		if *value == srri.matchValue || srri.matchValue == ".*" {
			return srri
		}
	}
	if err != nil {
		log.DefaultLogger.Errorf(RouterLogFormat, "sofa rotue rule", "get from ctx error", err)
	}
	log.DefaultLogger.Errorf(RouterLogFormat, "sofa rotue rule", "failed match", headers)
	return nil
}
