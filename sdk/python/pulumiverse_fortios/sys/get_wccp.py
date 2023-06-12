# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetWccpResult',
    'AwaitableGetWccpResult',
    'get_wccp',
    'get_wccp_output',
]

@pulumi.output_type
class GetWccpResult:
    """
    A collection of values returned by getWccp.
    """
    def __init__(__self__, assignment_bucket_format=None, assignment_dstaddr_mask=None, assignment_method=None, assignment_srcaddr_mask=None, assignment_weight=None, authentication=None, cache_engine_method=None, cache_id=None, forward_method=None, group_address=None, id=None, password=None, ports=None, ports_defined=None, primary_hash=None, priority=None, protocol=None, return_method=None, router_id=None, router_list=None, server_list=None, server_type=None, service_id=None, service_type=None, vdomparam=None):
        if assignment_bucket_format and not isinstance(assignment_bucket_format, str):
            raise TypeError("Expected argument 'assignment_bucket_format' to be a str")
        pulumi.set(__self__, "assignment_bucket_format", assignment_bucket_format)
        if assignment_dstaddr_mask and not isinstance(assignment_dstaddr_mask, str):
            raise TypeError("Expected argument 'assignment_dstaddr_mask' to be a str")
        pulumi.set(__self__, "assignment_dstaddr_mask", assignment_dstaddr_mask)
        if assignment_method and not isinstance(assignment_method, str):
            raise TypeError("Expected argument 'assignment_method' to be a str")
        pulumi.set(__self__, "assignment_method", assignment_method)
        if assignment_srcaddr_mask and not isinstance(assignment_srcaddr_mask, str):
            raise TypeError("Expected argument 'assignment_srcaddr_mask' to be a str")
        pulumi.set(__self__, "assignment_srcaddr_mask", assignment_srcaddr_mask)
        if assignment_weight and not isinstance(assignment_weight, int):
            raise TypeError("Expected argument 'assignment_weight' to be a int")
        pulumi.set(__self__, "assignment_weight", assignment_weight)
        if authentication and not isinstance(authentication, str):
            raise TypeError("Expected argument 'authentication' to be a str")
        pulumi.set(__self__, "authentication", authentication)
        if cache_engine_method and not isinstance(cache_engine_method, str):
            raise TypeError("Expected argument 'cache_engine_method' to be a str")
        pulumi.set(__self__, "cache_engine_method", cache_engine_method)
        if cache_id and not isinstance(cache_id, str):
            raise TypeError("Expected argument 'cache_id' to be a str")
        pulumi.set(__self__, "cache_id", cache_id)
        if forward_method and not isinstance(forward_method, str):
            raise TypeError("Expected argument 'forward_method' to be a str")
        pulumi.set(__self__, "forward_method", forward_method)
        if group_address and not isinstance(group_address, str):
            raise TypeError("Expected argument 'group_address' to be a str")
        pulumi.set(__self__, "group_address", group_address)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if ports and not isinstance(ports, str):
            raise TypeError("Expected argument 'ports' to be a str")
        pulumi.set(__self__, "ports", ports)
        if ports_defined and not isinstance(ports_defined, str):
            raise TypeError("Expected argument 'ports_defined' to be a str")
        pulumi.set(__self__, "ports_defined", ports_defined)
        if primary_hash and not isinstance(primary_hash, str):
            raise TypeError("Expected argument 'primary_hash' to be a str")
        pulumi.set(__self__, "primary_hash", primary_hash)
        if priority and not isinstance(priority, int):
            raise TypeError("Expected argument 'priority' to be a int")
        pulumi.set(__self__, "priority", priority)
        if protocol and not isinstance(protocol, int):
            raise TypeError("Expected argument 'protocol' to be a int")
        pulumi.set(__self__, "protocol", protocol)
        if return_method and not isinstance(return_method, str):
            raise TypeError("Expected argument 'return_method' to be a str")
        pulumi.set(__self__, "return_method", return_method)
        if router_id and not isinstance(router_id, str):
            raise TypeError("Expected argument 'router_id' to be a str")
        pulumi.set(__self__, "router_id", router_id)
        if router_list and not isinstance(router_list, str):
            raise TypeError("Expected argument 'router_list' to be a str")
        pulumi.set(__self__, "router_list", router_list)
        if server_list and not isinstance(server_list, str):
            raise TypeError("Expected argument 'server_list' to be a str")
        pulumi.set(__self__, "server_list", server_list)
        if server_type and not isinstance(server_type, str):
            raise TypeError("Expected argument 'server_type' to be a str")
        pulumi.set(__self__, "server_type", server_type)
        if service_id and not isinstance(service_id, str):
            raise TypeError("Expected argument 'service_id' to be a str")
        pulumi.set(__self__, "service_id", service_id)
        if service_type and not isinstance(service_type, str):
            raise TypeError("Expected argument 'service_type' to be a str")
        pulumi.set(__self__, "service_type", service_type)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="assignmentBucketFormat")
    def assignment_bucket_format(self) -> str:
        """
        Assignment bucket format for the WCCP cache engine.
        """
        return pulumi.get(self, "assignment_bucket_format")

    @property
    @pulumi.getter(name="assignmentDstaddrMask")
    def assignment_dstaddr_mask(self) -> str:
        """
        Assignment destination address mask.
        """
        return pulumi.get(self, "assignment_dstaddr_mask")

    @property
    @pulumi.getter(name="assignmentMethod")
    def assignment_method(self) -> str:
        """
        Hash key assignment preference.
        """
        return pulumi.get(self, "assignment_method")

    @property
    @pulumi.getter(name="assignmentSrcaddrMask")
    def assignment_srcaddr_mask(self) -> str:
        """
        Assignment source address mask.
        """
        return pulumi.get(self, "assignment_srcaddr_mask")

    @property
    @pulumi.getter(name="assignmentWeight")
    def assignment_weight(self) -> int:
        """
        Assignment of hash weight/ratio for the WCCP cache engine.
        """
        return pulumi.get(self, "assignment_weight")

    @property
    @pulumi.getter
    def authentication(self) -> str:
        """
        Enable/disable MD5 authentication.
        """
        return pulumi.get(self, "authentication")

    @property
    @pulumi.getter(name="cacheEngineMethod")
    def cache_engine_method(self) -> str:
        """
        Method used to forward traffic to the routers or to return to the cache engine.
        """
        return pulumi.get(self, "cache_engine_method")

    @property
    @pulumi.getter(name="cacheId")
    def cache_id(self) -> str:
        """
        IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
        """
        return pulumi.get(self, "cache_id")

    @property
    @pulumi.getter(name="forwardMethod")
    def forward_method(self) -> str:
        """
        Method used to forward traffic to the cache servers.
        """
        return pulumi.get(self, "forward_method")

    @property
    @pulumi.getter(name="groupAddress")
    def group_address(self) -> str:
        """
        IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
        """
        return pulumi.get(self, "group_address")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        Password for MD5 authentication.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def ports(self) -> str:
        """
        Service ports.
        """
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter(name="portsDefined")
    def ports_defined(self) -> str:
        """
        Match method.
        """
        return pulumi.get(self, "ports_defined")

    @property
    @pulumi.getter(name="primaryHash")
    def primary_hash(self) -> str:
        """
        Hash method.
        """
        return pulumi.get(self, "primary_hash")

    @property
    @pulumi.getter
    def priority(self) -> int:
        """
        Service priority.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def protocol(self) -> int:
        """
        Service protocol.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="returnMethod")
    def return_method(self) -> str:
        """
        Method used to decline a redirected packet and return it to the FortiGate.
        """
        return pulumi.get(self, "return_method")

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> str:
        """
        IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
        """
        return pulumi.get(self, "router_id")

    @property
    @pulumi.getter(name="routerList")
    def router_list(self) -> str:
        """
        IP addresses of one or more WCCP routers.
        """
        return pulumi.get(self, "router_list")

    @property
    @pulumi.getter(name="serverList")
    def server_list(self) -> str:
        """
        IP addresses and netmasks for up to four cache servers.
        """
        return pulumi.get(self, "server_list")

    @property
    @pulumi.getter(name="serverType")
    def server_type(self) -> str:
        """
        Cache server type.
        """
        return pulumi.get(self, "server_type")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> str:
        """
        Service ID.
        """
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> str:
        """
        WCCP service type used by the cache server for logical interception and redirection of traffic.
        """
        return pulumi.get(self, "service_type")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetWccpResult(GetWccpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWccpResult(
            assignment_bucket_format=self.assignment_bucket_format,
            assignment_dstaddr_mask=self.assignment_dstaddr_mask,
            assignment_method=self.assignment_method,
            assignment_srcaddr_mask=self.assignment_srcaddr_mask,
            assignment_weight=self.assignment_weight,
            authentication=self.authentication,
            cache_engine_method=self.cache_engine_method,
            cache_id=self.cache_id,
            forward_method=self.forward_method,
            group_address=self.group_address,
            id=self.id,
            password=self.password,
            ports=self.ports,
            ports_defined=self.ports_defined,
            primary_hash=self.primary_hash,
            priority=self.priority,
            protocol=self.protocol,
            return_method=self.return_method,
            router_id=self.router_id,
            router_list=self.router_list,
            server_list=self.server_list,
            server_type=self.server_type,
            service_id=self.service_id,
            service_type=self.service_type,
            vdomparam=self.vdomparam)


def get_wccp(service_id: Optional[str] = None,
             vdomparam: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWccpResult:
    """
    Use this data source to get information on an fortios system wccp


    :param str service_id: Specify the service_id of the desired system wccp.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['serviceId'] = service_id
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:sys/getWccp:getWccp', __args__, opts=opts, typ=GetWccpResult).value

    return AwaitableGetWccpResult(
        assignment_bucket_format=__ret__.assignment_bucket_format,
        assignment_dstaddr_mask=__ret__.assignment_dstaddr_mask,
        assignment_method=__ret__.assignment_method,
        assignment_srcaddr_mask=__ret__.assignment_srcaddr_mask,
        assignment_weight=__ret__.assignment_weight,
        authentication=__ret__.authentication,
        cache_engine_method=__ret__.cache_engine_method,
        cache_id=__ret__.cache_id,
        forward_method=__ret__.forward_method,
        group_address=__ret__.group_address,
        id=__ret__.id,
        password=__ret__.password,
        ports=__ret__.ports,
        ports_defined=__ret__.ports_defined,
        primary_hash=__ret__.primary_hash,
        priority=__ret__.priority,
        protocol=__ret__.protocol,
        return_method=__ret__.return_method,
        router_id=__ret__.router_id,
        router_list=__ret__.router_list,
        server_list=__ret__.server_list,
        server_type=__ret__.server_type,
        service_id=__ret__.service_id,
        service_type=__ret__.service_type,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_wccp)
def get_wccp_output(service_id: Optional[pulumi.Input[str]] = None,
                    vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWccpResult]:
    """
    Use this data source to get information on an fortios system wccp


    :param str service_id: Specify the service_id of the desired system wccp.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...