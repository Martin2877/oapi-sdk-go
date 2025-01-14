// Package dispatcher code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package dispatcher

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/v3/service/corehr/v2"
)

// -
//
// - 事件描述文档链接:
func (dispatcher *EventDispatcher) OnP2EmployeeDomainEventV2(handler func(ctx context.Context, event *larkcorehr.P2EmployeeDomainEventV2) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.employee.domain_event_v2"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.employee.domain_event_v2")
	}
	dispatcher.eventType2EventHandler["corehr.employee.domain_event_v2"] = larkcorehr.NewP2EmployeeDomainEventV2Handler(handler)
	return dispatcher
}

// -
//
// - 事件描述文档链接:
func (dispatcher *EventDispatcher) OnP2JobChangeUpdatedV2(handler func(ctx context.Context, event *larkcorehr.P2JobChangeUpdatedV2) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.job_change.updated_v2"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.job_change.updated_v2")
	}
	dispatcher.eventType2EventHandler["corehr.job_change.updated_v2"] = larkcorehr.NewP2JobChangeUpdatedV2Handler(handler)
	return dispatcher
}

// -
//
// - 事件描述文档链接:
func (dispatcher *EventDispatcher) OnP2OffboardingChecklistUpdatedV2(handler func(ctx context.Context, event *larkcorehr.P2OffboardingChecklistUpdatedV2) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.offboarding.checklist_updated_v2"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.offboarding.checklist_updated_v2")
	}
	dispatcher.eventType2EventHandler["corehr.offboarding.checklist_updated_v2"] = larkcorehr.NewP2OffboardingChecklistUpdatedV2Handler(handler)
	return dispatcher
}

// -
//
// - 事件描述文档链接:
func (dispatcher *EventDispatcher) OnP2OffboardingStatusUpdatedV2(handler func(ctx context.Context, event *larkcorehr.P2OffboardingStatusUpdatedV2) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.offboarding.status_updated_v2"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.offboarding.status_updated_v2")
	}
	dispatcher.eventType2EventHandler["corehr.offboarding.status_updated_v2"] = larkcorehr.NewP2OffboardingStatusUpdatedV2Handler(handler)
	return dispatcher
}

// -
//
// - 事件描述文档链接:
func (dispatcher *EventDispatcher) OnP2OffboardingUpdatedV2(handler func(ctx context.Context, event *larkcorehr.P2OffboardingUpdatedV2) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.offboarding.updated_v2"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.offboarding.updated_v2")
	}
	dispatcher.eventType2EventHandler["corehr.offboarding.updated_v2"] = larkcorehr.NewP2OffboardingUpdatedV2Handler(handler)
	return dispatcher
}

// -
//
// - 事件描述文档链接:
func (dispatcher *EventDispatcher) OnP2ProbationUpdatedV2(handler func(ctx context.Context, event *larkcorehr.P2ProbationUpdatedV2) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.probation.updated_v2"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.probation.updated_v2")
	}
	dispatcher.eventType2EventHandler["corehr.probation.updated_v2"] = larkcorehr.NewP2ProbationUpdatedV2Handler(handler)
	return dispatcher
}

// -
//
// - 事件描述文档链接:
func (dispatcher *EventDispatcher) OnP2ProcessUpdatedV2(handler func(ctx context.Context, event *larkcorehr.P2ProcessUpdatedV2) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.process.updated_v2"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.process.updated_v2")
	}
	dispatcher.eventType2EventHandler["corehr.process.updated_v2"] = larkcorehr.NewP2ProcessUpdatedV2Handler(handler)
	return dispatcher
}

// -
//
// - 事件描述文档链接:
func (dispatcher *EventDispatcher) OnP2ProcessApproverUpdatedV2(handler func(ctx context.Context, event *larkcorehr.P2ProcessApproverUpdatedV2) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.process.approver.updated_v2"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.process.approver.updated_v2")
	}
	dispatcher.eventType2EventHandler["corehr.process.approver.updated_v2"] = larkcorehr.NewP2ProcessApproverUpdatedV2Handler(handler)
	return dispatcher
}

// -
//
// - 事件描述文档链接:
func (dispatcher *EventDispatcher) OnP2ProcessCcUpdatedV2(handler func(ctx context.Context, event *larkcorehr.P2ProcessCcUpdatedV2) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.process.cc.updated_v2"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.process.cc.updated_v2")
	}
	dispatcher.eventType2EventHandler["corehr.process.cc.updated_v2"] = larkcorehr.NewP2ProcessCcUpdatedV2Handler(handler)
	return dispatcher
}

// -
//
// - 事件描述文档链接:
func (dispatcher *EventDispatcher) OnP2ProcessNodeUpdatedV2(handler func(ctx context.Context, event *larkcorehr.P2ProcessNodeUpdatedV2) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.process.node.updated_v2"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.process.node.updated_v2")
	}
	dispatcher.eventType2EventHandler["corehr.process.node.updated_v2"] = larkcorehr.NewP2ProcessNodeUpdatedV2Handler(handler)
	return dispatcher
}
