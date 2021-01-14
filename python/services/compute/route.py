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
from google3.cloud.graphite.mmv2.services.google.compute import route_pb2
from google3.cloud.graphite.mmv2.services.google.compute import route_pb2_grpc

from typing import List


class Route(object):
    def __init__(
        self,
        id: int = None,
        name: str = None,
        description: str = None,
        network: str = None,
        tag: list = None,
        dest_range: str = None,
        priority: int = None,
        next_hop_instance: str = None,
        next_hop_ip: str = None,
        next_hop_network: str = None,
        next_hop_gateway: str = None,
        next_hop_peering: str = None,
        next_hop_ilb: str = None,
        warning: list = None,
        next_hop_vpn_tunnel: str = None,
        self_link: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.network = network
        self.tag = tag
        self.dest_range = dest_range
        self.priority = priority
        self.next_hop_instance = next_hop_instance
        self.next_hop_ip = next_hop_ip
        self.next_hop_network = next_hop_network
        self.next_hop_gateway = next_hop_gateway
        self.next_hop_ilb = next_hop_ilb
        self.next_hop_vpn_tunnel = next_hop_vpn_tunnel
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = route_pb2_grpc.ComputeRouteServiceStub(channel.Channel())
        request = route_pb2.ApplyComputeRouteRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.tag):
            request.resource.tag.extend(Primitive.to_proto(self.tag))
        if Primitive.to_proto(self.dest_range):
            request.resource.dest_range = Primitive.to_proto(self.dest_range)

        if Primitive.to_proto(self.priority):
            request.resource.priority = Primitive.to_proto(self.priority)

        if Primitive.to_proto(self.next_hop_instance):
            request.resource.next_hop_instance = Primitive.to_proto(
                self.next_hop_instance
            )

        if Primitive.to_proto(self.next_hop_ip):
            request.resource.next_hop_ip = Primitive.to_proto(self.next_hop_ip)

        if Primitive.to_proto(self.next_hop_network):
            request.resource.next_hop_network = Primitive.to_proto(
                self.next_hop_network
            )

        if Primitive.to_proto(self.next_hop_gateway):
            request.resource.next_hop_gateway = Primitive.to_proto(
                self.next_hop_gateway
            )

        if Primitive.to_proto(self.next_hop_ilb):
            request.resource.next_hop_ilb = Primitive.to_proto(self.next_hop_ilb)

        if Primitive.to_proto(self.next_hop_vpn_tunnel):
            request.resource.next_hop_vpn_tunnel = Primitive.to_proto(
                self.next_hop_vpn_tunnel
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeRoute(request)
        self.id = Primitive.from_proto(response.id)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.network = Primitive.from_proto(response.network)
        self.tag = Primitive.from_proto(response.tag)
        self.dest_range = Primitive.from_proto(response.dest_range)
        self.priority = Primitive.from_proto(response.priority)
        self.next_hop_instance = Primitive.from_proto(response.next_hop_instance)
        self.next_hop_ip = Primitive.from_proto(response.next_hop_ip)
        self.next_hop_network = Primitive.from_proto(response.next_hop_network)
        self.next_hop_gateway = Primitive.from_proto(response.next_hop_gateway)
        self.next_hop_peering = Primitive.from_proto(response.next_hop_peering)
        self.next_hop_ilb = Primitive.from_proto(response.next_hop_ilb)
        self.warning = RouteWarningArray.from_proto(response.warning)
        self.next_hop_vpn_tunnel = Primitive.from_proto(response.next_hop_vpn_tunnel)
        self.self_link = Primitive.from_proto(response.self_link)
        self.project = Primitive.from_proto(response.project)

    def hcl(self):
        stub = route_pb2_grpc.ComputeRouteServiceStub(channel.Channel())
        request = route_pb2.ComputeRouteAsHclRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.tag):
            request.resource.tag.extend(Primitive.to_proto(self.tag))
        if Primitive.to_proto(self.dest_range):
            request.resource.dest_range = Primitive.to_proto(self.dest_range)

        if Primitive.to_proto(self.priority):
            request.resource.priority = Primitive.to_proto(self.priority)

        if Primitive.to_proto(self.next_hop_instance):
            request.resource.next_hop_instance = Primitive.to_proto(
                self.next_hop_instance
            )

        if Primitive.to_proto(self.next_hop_ip):
            request.resource.next_hop_ip = Primitive.to_proto(self.next_hop_ip)

        if Primitive.to_proto(self.next_hop_network):
            request.resource.next_hop_network = Primitive.to_proto(
                self.next_hop_network
            )

        if Primitive.to_proto(self.next_hop_gateway):
            request.resource.next_hop_gateway = Primitive.to_proto(
                self.next_hop_gateway
            )

        if Primitive.to_proto(self.next_hop_ilb):
            request.resource.next_hop_ilb = Primitive.to_proto(self.next_hop_ilb)

        if Primitive.to_proto(self.next_hop_vpn_tunnel):
            request.resource.next_hop_vpn_tunnel = Primitive.to_proto(
                self.next_hop_vpn_tunnel
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.ComputeRouteAsHcl(request)
        return response.hcl

    @classmethod
    def delete(self, project, name, service_account_file=""):
        stub = route_pb2_grpc.ComputeRouteServiceStub(channel.Channel())
        request = route_pb2.DeleteComputeRouteRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Name = name

        response = stub.DeleteComputeRoute(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = route_pb2_grpc.ComputeRouteServiceStub(channel.Channel())
        request = route_pb2.ListComputeRouteRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListComputeRoute(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = route_pb2.ComputeRoute()
        any_proto.Unpack(res_proto)

        res = Route()
        res.id = Primitive.from_proto(res_proto.id)
        res.name = Primitive.from_proto(res_proto.name)
        res.description = Primitive.from_proto(res_proto.description)
        res.network = Primitive.from_proto(res_proto.network)
        res.tag = Primitive.from_proto(res_proto.tag)
        res.dest_range = Primitive.from_proto(res_proto.dest_range)
        res.priority = Primitive.from_proto(res_proto.priority)
        res.next_hop_instance = Primitive.from_proto(res_proto.next_hop_instance)
        res.next_hop_ip = Primitive.from_proto(res_proto.next_hop_ip)
        res.next_hop_network = Primitive.from_proto(res_proto.next_hop_network)
        res.next_hop_gateway = Primitive.from_proto(res_proto.next_hop_gateway)
        res.next_hop_peering = Primitive.from_proto(res_proto.next_hop_peering)
        res.next_hop_ilb = Primitive.from_proto(res_proto.next_hop_ilb)
        res.warning = RouteWarningArray.from_proto(res_proto.warning)
        res.next_hop_vpn_tunnel = Primitive.from_proto(res_proto.next_hop_vpn_tunnel)
        res.self_link = Primitive.from_proto(res_proto.self_link)
        res.project = Primitive.from_proto(res_proto.project)
        return res


class RouteWarning(object):
    def __init__(self, code: str = None, message: str = None, data: dict = None):
        self.code = code
        self.message = message
        self.data = data

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = route_pb2.ComputeRouteWarning()
        if RouteWarningCodeEnum.to_proto(resource.code):
            res.code = RouteWarningCodeEnum.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if Primitive.to_proto(resource.data):
            res.data = Primitive.to_proto(resource.data)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return RouteWarning(
            code=resource.code, message=resource.message, data=resource.data,
        )


class RouteWarningArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [RouteWarning.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [RouteWarning.from_proto(i) for i in resources]


class RouteWarningCodeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return route_pb2.ComputeRouteWarningCodeEnum.Value(
            "ComputeRouteWarningCodeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return route_pb2.ComputeRouteWarningCodeEnum.Name(resource)[
            len("ComputeRouteWarningCodeEnum") :
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
