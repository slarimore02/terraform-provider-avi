<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_serverautoscalepolicy"
sidebar_current: "docs-avi-datasource-serverautoscalepolicy"
description: |-
  Get information of Avi ServerAutoScalePolicy.
---

# avi_serverautoscalepolicy

This data source is used to to get avi_serverautoscalepolicy objects.

## Example Usage

```hcl
data "avi_serverautoscalepolicy" "foo_serverautoscalepolicy" {
    uuid = "serverautoscalepolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ServerAutoScalePolicy by name.
* `uuid` - (Optional) Search ServerAutoScalePolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `delay_for_server_garbage_collection` - Delay in minutes after which a down server will be removed from pool. Value 0 disables this functionality. Field introduced in 20.1.3.
* `description` - User defined description for the object.
* `intelligent_autoscale` - Use avi intelligent autoscale algorithm where autoscale is performed by comparing load on the pool against estimated capacity of all the servers.
* `intelligent_scalein_margin` - Maximum extra capacity as percentage of load used by the intelligent scheme. Scale-in is triggered when available capacity is more than this margin. Allowed values are 1-99.
* `intelligent_scaleout_margin` - Minimum extra capacity as percentage of load used by the intelligent scheme. Scale-out is triggered when available capacity is less than this margin. Allowed values are 1-99.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `max_scalein_adjustment_step` - Maximum number of servers to scale-in simultaneously. The actual number of servers to scale-in is chosen such that target number of servers is always more than or equal to the min_size.
* `max_scaleout_adjustment_step` - Maximum number of servers to scale-out simultaneously. The actual number of servers to scale-out is chosen such that target number of servers is always less than or equal to the max_size.
* `max_size` - Maximum number of servers after scale-out. Allowed values are 0-400.
* `min_size` - No scale-in happens once number of operationally up servers reach min_servers. Allowed values are 0-400.
* `name` - Name of the object.
* `scalein_alertconfig_refs` - Trigger scale-in when alerts due to any of these alert configurations are raised. It is a reference to an object of type alertconfig.
* `scalein_cooldown` - Cooldown period during which no new scale-in is triggered to allow previous scale-in to successfully complete. Unit is sec.
* `scaleout_alertconfig_refs` - Trigger scale-out when alerts due to any of these alert configurations are raised. It is a reference to an object of type alertconfig.
* `scaleout_cooldown` - Cooldown period during which no new scale-out is triggered to allow previous scale-out to successfully complete. Unit is sec.
* `scheduled_scalings` - Scheduled-based scale-in/out policy. During scheduled intervals, metrics based autoscale is not enabled and number of servers will be solely derived from schedulescale policy. Field introduced in 21.1.1. Maximum of 1 items allowed.
* `tenant_ref` - It is a reference to an object of type tenant.
* `use_predicted_load` - Use predicted load rather than current load.
* `uuid` - Unique object identifier of the object.

