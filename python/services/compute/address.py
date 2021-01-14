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
from google3.cloud.graphite.mmv2.services.google.compute import address_pb2
from google3.cloud.graphite.mmv2.services.google.compute import address_pb2_grpc

from typing import List


class Address(object):
    def __init__(
        self,
        id: int = None,
        name: str = None,
        description: str = None,
        address: str = None,
        prefix_length: int = None,
        status: str = None,
        region: str = None,
        self_link: str = None,
        network_tier: str = None,
        ip_version: str = None,
        address_type: str = None,
        purpose: str = None,
        subnetwork: str = None,
        network: str = None,
        project: str = None,
        creation_timestamp: str = None,
        users: list = None,
        label_fingerprint: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.address = address
        self.prefix_length = prefix_length
        self.region = region
        self.network_tier = network_tier
        self.ip_version = ip_version
        self.address_type = address_type
        self.purpose = purpose
        self.subnetwork = subnetwork
        self.network = network
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = address_pb2_grpc.ComputeAddressServiceStub(channel.Channel())
        request = address_pb2.ApplyComputeAddressRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.address):
            request.resource.address = Primitive.to_proto(self.address)

        if Primitive.to_proto(self.prefix_length):
            request.resource.prefix_length = Primitive.to_proto(self.prefix_length)

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if AddressNetworkTierEnum.to_proto(self.network_tier):
            request.resource.network_tier = AddressNetworkTierEnum.to_proto(
                self.network_tier
            )

        if AddressIPVersionEnum.to_proto(self.ip_version):
            request.resource.ip_version = AddressIPVersionEnum.to_proto(self.ip_version)

        if AddressAddressTypeEnum.to_proto(self.address_type):
            request.resource.address_type = AddressAddressTypeEnum.to_proto(
                self.address_type
            )

        if AddressPurposeEnum.to_proto(self.purpose):
            request.resource.purpose = AddressPurposeEnum.to_proto(self.purpose)

        if Primitive.to_proto(self.subnetwork):
            request.resource.subnetwork = Primitive.to_proto(self.subnetwork)

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeAddress(request)
        self.id = Primitive.from_proto(response.id)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.address = Primitive.from_proto(response.address)
        self.prefix_length = Primitive.from_proto(response.prefix_length)
        self.status = AddressStatusEnum.from_proto(response.status)
        self.region = Primitive.from_proto(response.region)
        self.self_link = Primitive.from_proto(response.self_link)
        self.network_tier = AddressNetworkTierEnum.from_proto(response.network_tier)
        self.ip_version = AddressIPVersionEnum.from_proto(response.ip_version)
        self.address_type = AddressAddressTypeEnum.from_proto(response.address_type)
        self.purpose = AddressPurposeEnum.from_proto(response.purpose)
        self.subnetwork = Primitive.from_proto(response.subnetwork)
        self.network = Primitive.from_proto(response.network)
        self.project = Primitive.from_proto(response.project)
        self.creation_timestamp = Primitive.from_proto(response.creation_timestamp)
        self.users = Primitive.from_proto(response.users)
        self.label_fingerprint = Primitive.from_proto(response.label_fingerprint)
        self.location = Primitive.from_proto(response.location)

    @classmethod
    def delete(self, project, location, name, service_account_file=""):
        stub = address_pb2_grpc.ComputeAddressServiceStub(channel.Channel())
        request = address_pb2.DeleteComputeAddressRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.Name = name

        response = stub.DeleteComputeAddress(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = address_pb2_grpc.ComputeAddressServiceStub(channel.Channel())
        request = address_pb2.ListComputeAddressRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListComputeAddress(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = address_pb2.ComputeAddress()
        any_proto.Unpack(res_proto)

        res = Address()
        res.id = Primitive.from_proto(res_proto.id)
        res.name = Primitive.from_proto(res_proto.name)
        res.description = Primitive.from_proto(res_proto.description)
        res.address = Primitive.from_proto(res_proto.address)
        res.prefix_length = Primitive.from_proto(res_proto.prefix_length)
        res.status = AddressStatusEnum.from_proto(res_proto.status)
        res.region = Primitive.from_proto(res_proto.region)
        res.self_link = Primitive.from_proto(res_proto.self_link)
        res.network_tier = AddressNetworkTierEnum.from_proto(res_proto.network_tier)
        res.ip_version = AddressIPVersionEnum.from_proto(res_proto.ip_version)
        res.address_type = AddressAddressTypeEnum.from_proto(res_proto.address_type)
        res.purpose = AddressPurposeEnum.from_proto(res_proto.purpose)
        res.subnetwork = Primitive.from_proto(res_proto.subnetwork)
        res.network = Primitive.from_proto(res_proto.network)
        res.project = Primitive.from_proto(res_proto.project)
        res.creation_timestamp = Primitive.from_proto(res_proto.creation_timestamp)
        res.users = Primitive.from_proto(res_proto.users)
        res.label_fingerprint = Primitive.from_proto(res_proto.label_fingerprint)
        res.location = Primitive.from_proto(res_proto.location)
        return res


class AddressStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return address_pb2.ComputeAddressStatusEnum.Value(
            "ComputeAddressStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return address_pb2.ComputeAddressStatusEnum.Name(resource)[
            len("ComputeAddressStatusEnum") :
        ]


class AddressNetworkTierEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return address_pb2.ComputeAddressNetworkTierEnum.Value(
            "ComputeAddressNetworkTierEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return address_pb2.ComputeAddressNetworkTierEnum.Name(resource)[
            len("ComputeAddressNetworkTierEnum") :
        ]


class AddressIPVersionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return address_pb2.ComputeAddressIPVersionEnum.Value(
            "ComputeAddressIPVersionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return address_pb2.ComputeAddressIPVersionEnum.Name(resource)[
            len("ComputeAddressIPVersionEnum") :
        ]


class AddressAddressTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return address_pb2.ComputeAddressAddressTypeEnum.Value(
            "ComputeAddressAddressTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return address_pb2.ComputeAddressAddressTypeEnum.Name(resource)[
            len("ComputeAddressAddressTypeEnum") :
        ]


class AddressPurposeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return address_pb2.ComputeAddressPurposeEnum.Value(
            "ComputeAddressPurposeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return address_pb2.ComputeAddressPurposeEnum.Name(resource)[
            len("ComputeAddressPurposeEnum") :
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
