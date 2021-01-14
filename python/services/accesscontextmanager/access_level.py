# Copyright 2021 Google LLC. All Rights Reserved.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
from connector import channel
from google3.cloud.graphite.mmv2.services.google.access_context_manager import (
    access_level_pb2,
)
from google3.cloud.graphite.mmv2.services.google.access_context_manager import (
    access_level_pb2_grpc,
)

from typing import List


class AccessLevel(object):
    def __init__(
        self,
        title: str = None,
        create_time: str = None,
        update_time: str = None,
        description: str = None,
        basic: dict = None,
        name: str = None,
        policy: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.title = title
        self.create_time = create_time
        self.update_time = update_time
        self.description = description
        self.basic = basic
        self.name = name
        self.policy = policy
        self.service_account_file = service_account_file

    def apply(self):
        stub = access_level_pb2_grpc.AccesscontextmanagerAccessLevelServiceStub(
            channel.Channel()
        )
        request = access_level_pb2.ApplyAccesscontextmanagerAccessLevelRequest()
        if Primitive.to_proto(self.title):
            request.resource.title = Primitive.to_proto(self.title)

        if Primitive.to_proto(self.create_time):
            request.resource.create_time = Primitive.to_proto(self.create_time)

        if Primitive.to_proto(self.update_time):
            request.resource.update_time = Primitive.to_proto(self.update_time)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if AccessLevelBasic.to_proto(self.basic):
            request.resource.basic.CopyFrom(AccessLevelBasic.to_proto(self.basic))
        else:
            request.resource.ClearField("basic")
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.policy):
            request.resource.policy = Primitive.to_proto(self.policy)

        request.service_account_file = self.service_account_file

        response = stub.ApplyAccesscontextmanagerAccessLevel(request)
        self.title = Primitive.from_proto(response.title)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.description = Primitive.from_proto(response.description)
        self.basic = AccessLevelBasic.from_proto(response.basic)
        self.name = Primitive.from_proto(response.name)
        self.policy = Primitive.from_proto(response.policy)

    @classmethod
    def delete(self, policy, name, service_account_file=""):
        stub = access_level_pb2_grpc.AccesscontextmanagerAccessLevelServiceStub(
            channel.Channel()
        )
        request = access_level_pb2.DeleteAccesscontextmanagerAccessLevelRequest()
        request.service_account_file = service_account_file
        request.Policy = policy

        request.Name = name

        response = stub.DeleteAccesscontextmanagerAccessLevel(request)

    @classmethod
    def list(self, policy, service_account_file=""):
        stub = access_level_pb2_grpc.AccesscontextmanagerAccessLevelServiceStub(
            channel.Channel()
        )
        request = access_level_pb2.ListAccesscontextmanagerAccessLevelRequest()
        request.service_account_file = service_account_file
        request.Policy = policy

        return stub.ListAccesscontextmanagerAccessLevel(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = access_level_pb2.AccesscontextmanagerAccessLevel()
        any_proto.Unpack(res_proto)

        res = AccessLevel()
        res.title = Primitive.from_proto(res_proto.title)
        res.create_time = Primitive.from_proto(res_proto.create_time)
        res.update_time = Primitive.from_proto(res_proto.update_time)
        res.description = Primitive.from_proto(res_proto.description)
        res.basic = AccessLevelBasic.from_proto(res_proto.basic)
        res.name = Primitive.from_proto(res_proto.name)
        res.policy = Primitive.from_proto(res_proto.policy)
        return res


class AccessLevelBasic(object):
    def __init__(self, combining_function: str = None, conditions: list = None):
        self.combining_function = combining_function
        self.conditions = conditions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = access_level_pb2.AccesscontextmanagerAccessLevelBasic()
        if AccessLevelBasicCombiningFunctionEnum.to_proto(resource.combining_function):
            res.combining_function = AccessLevelBasicCombiningFunctionEnum.to_proto(
                resource.combining_function
            )
        if AccessLevelBasicConditionsArray.to_proto(resource.conditions):
            res.conditions.extend(
                AccessLevelBasicConditionsArray.to_proto(resource.conditions)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AccessLevelBasic(
            combining_function=resource.combining_function,
            conditions=resource.conditions,
        )


class AccessLevelBasicArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AccessLevelBasic.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AccessLevelBasic.from_proto(i) for i in resources]


class AccessLevelBasicConditions(object):
    def __init__(
        self,
        regions: list = None,
        ip_subnetworks: list = None,
        required_access_levels: list = None,
        members: list = None,
        negate: bool = None,
        device_policy: dict = None,
    ):
        self.regions = regions
        self.ip_subnetworks = ip_subnetworks
        self.required_access_levels = required_access_levels
        self.members = members
        self.negate = negate
        self.device_policy = device_policy

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = access_level_pb2.AccesscontextmanagerAccessLevelBasicConditions()
        if Primitive.to_proto(resource.regions):
            res.regions.extend(Primitive.to_proto(resource.regions))
        if Primitive.to_proto(resource.ip_subnetworks):
            res.ip_subnetworks.extend(Primitive.to_proto(resource.ip_subnetworks))
        if Primitive.to_proto(resource.required_access_levels):
            res.required_access_levels.extend(
                Primitive.to_proto(resource.required_access_levels)
            )
        if Primitive.to_proto(resource.members):
            res.members.extend(Primitive.to_proto(resource.members))
        if Primitive.to_proto(resource.negate):
            res.negate = Primitive.to_proto(resource.negate)
        if AccessLevelBasicConditionsDevicePolicy.to_proto(resource.device_policy):
            res.device_policy.CopyFrom(
                AccessLevelBasicConditionsDevicePolicy.to_proto(resource.device_policy)
            )
        else:
            res.ClearField("device_policy")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AccessLevelBasicConditions(
            regions=resource.regions,
            ip_subnetworks=resource.ip_subnetworks,
            required_access_levels=resource.required_access_levels,
            members=resource.members,
            negate=resource.negate,
            device_policy=resource.device_policy,
        )


class AccessLevelBasicConditionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AccessLevelBasicConditions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AccessLevelBasicConditions.from_proto(i) for i in resources]


class AccessLevelBasicConditionsDevicePolicy(object):
    def __init__(
        self,
        require_screenlock: bool = None,
        require_admin_approval: bool = None,
        require_corp_owned: bool = None,
        allowed_encryption_statuses: list = None,
        allowed_device_management_levels: list = None,
        os_constraints: list = None,
    ):
        self.require_screenlock = require_screenlock
        self.require_admin_approval = require_admin_approval
        self.require_corp_owned = require_corp_owned
        self.allowed_encryption_statuses = allowed_encryption_statuses
        self.allowed_device_management_levels = allowed_device_management_levels
        self.os_constraints = os_constraints

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            access_level_pb2.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicy()
        )
        if Primitive.to_proto(resource.require_screenlock):
            res.require_screenlock = Primitive.to_proto(resource.require_screenlock)
        if Primitive.to_proto(resource.require_admin_approval):
            res.require_admin_approval = Primitive.to_proto(
                resource.require_admin_approval
            )
        if Primitive.to_proto(resource.require_corp_owned):
            res.require_corp_owned = Primitive.to_proto(resource.require_corp_owned)
        if AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnumArray.to_proto(
            resource.allowed_encryption_statuses
        ):
            res.allowed_encryption_statuses.extend(
                AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnumArray.to_proto(
                    resource.allowed_encryption_statuses
                )
            )
        if AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnumArray.to_proto(
            resource.allowed_device_management_levels
        ):
            res.allowed_device_management_levels.extend(
                AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnumArray.to_proto(
                    resource.allowed_device_management_levels
                )
            )
        if AccessLevelBasicConditionsDevicePolicyOsConstraintsArray.to_proto(
            resource.os_constraints
        ):
            res.os_constraints.extend(
                AccessLevelBasicConditionsDevicePolicyOsConstraintsArray.to_proto(
                    resource.os_constraints
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AccessLevelBasicConditionsDevicePolicy(
            require_screenlock=resource.require_screenlock,
            require_admin_approval=resource.require_admin_approval,
            require_corp_owned=resource.require_corp_owned,
            allowed_encryption_statuses=resource.allowed_encryption_statuses,
            allowed_device_management_levels=resource.allowed_device_management_levels,
            os_constraints=resource.os_constraints,
        )


class AccessLevelBasicConditionsDevicePolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AccessLevelBasicConditionsDevicePolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AccessLevelBasicConditionsDevicePolicy.from_proto(i) for i in resources]


class AccessLevelBasicConditionsDevicePolicyOsConstraints(object):
    def __init__(
        self,
        minimum_version: str = None,
        os_type: str = None,
        require_verified_chrome_os: bool = None,
    ):
        self.minimum_version = minimum_version
        self.os_type = os_type
        self.require_verified_chrome_os = require_verified_chrome_os

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            access_level_pb2.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOsConstraints()
        )
        if Primitive.to_proto(resource.minimum_version):
            res.minimum_version = Primitive.to_proto(resource.minimum_version)
        if AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum.to_proto(
            resource.os_type
        ):
            res.os_type = AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum.to_proto(
                resource.os_type
            )
        if Primitive.to_proto(resource.require_verified_chrome_os):
            res.require_verified_chrome_os = Primitive.to_proto(
                resource.require_verified_chrome_os
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AccessLevelBasicConditionsDevicePolicyOsConstraints(
            minimum_version=resource.minimum_version,
            os_type=resource.os_type,
            require_verified_chrome_os=resource.require_verified_chrome_os,
        )


class AccessLevelBasicConditionsDevicePolicyOsConstraintsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AccessLevelBasicConditionsDevicePolicyOsConstraints.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AccessLevelBasicConditionsDevicePolicyOsConstraints.from_proto(i)
            for i in resources
        ]


class AccessLevelBasicCombiningFunctionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return access_level_pb2.AccesscontextmanagerAccessLevelBasicCombiningFunctionEnum.Value(
            "AccesscontextmanagerAccessLevelBasicCombiningFunctionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return access_level_pb2.AccesscontextmanagerAccessLevelBasicCombiningFunctionEnum.Name(
            resource
        )[
            len("AccesscontextmanagerAccessLevelBasicCombiningFunctionEnum") :
        ]


class AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return access_level_pb2.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum.Value(
            "AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return access_level_pb2.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum.Name(
            resource
        )[
            len(
                "AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum"
            ) :
        ]


class AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return access_level_pb2.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum.Value(
            "AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return access_level_pb2.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum.Name(
            resource
        )[
            len(
                "AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum"
            ) :
        ]


class AccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return access_level_pb2.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum.Value(
            "AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return access_level_pb2.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum.Name(
            resource
        )[
            len(
                "AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOsTypeEnum"
            ) :
        ]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
