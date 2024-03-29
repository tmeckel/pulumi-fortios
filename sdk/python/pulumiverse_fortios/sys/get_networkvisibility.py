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
    'GetNetworkvisibilityResult',
    'AwaitableGetNetworkvisibilityResult',
    'get_networkvisibility',
    'get_networkvisibility_output',
]

@pulumi.output_type
class GetNetworkvisibilityResult:
    """
    A collection of values returned by getNetworkvisibility.
    """
    def __init__(__self__, destination_hostname_visibility=None, destination_location=None, destination_visibility=None, hostname_limit=None, hostname_ttl=None, id=None, source_location=None, vdomparam=None):
        if destination_hostname_visibility and not isinstance(destination_hostname_visibility, str):
            raise TypeError("Expected argument 'destination_hostname_visibility' to be a str")
        pulumi.set(__self__, "destination_hostname_visibility", destination_hostname_visibility)
        if destination_location and not isinstance(destination_location, str):
            raise TypeError("Expected argument 'destination_location' to be a str")
        pulumi.set(__self__, "destination_location", destination_location)
        if destination_visibility and not isinstance(destination_visibility, str):
            raise TypeError("Expected argument 'destination_visibility' to be a str")
        pulumi.set(__self__, "destination_visibility", destination_visibility)
        if hostname_limit and not isinstance(hostname_limit, int):
            raise TypeError("Expected argument 'hostname_limit' to be a int")
        pulumi.set(__self__, "hostname_limit", hostname_limit)
        if hostname_ttl and not isinstance(hostname_ttl, int):
            raise TypeError("Expected argument 'hostname_ttl' to be a int")
        pulumi.set(__self__, "hostname_ttl", hostname_ttl)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if source_location and not isinstance(source_location, str):
            raise TypeError("Expected argument 'source_location' to be a str")
        pulumi.set(__self__, "source_location", source_location)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="destinationHostnameVisibility")
    def destination_hostname_visibility(self) -> str:
        """
        Enable/disable logging of destination hostname visibility.
        """
        return pulumi.get(self, "destination_hostname_visibility")

    @property
    @pulumi.getter(name="destinationLocation")
    def destination_location(self) -> str:
        """
        Enable/disable logging of destination geographical location visibility.
        """
        return pulumi.get(self, "destination_location")

    @property
    @pulumi.getter(name="destinationVisibility")
    def destination_visibility(self) -> str:
        """
        Enable/disable logging of destination visibility.
        """
        return pulumi.get(self, "destination_visibility")

    @property
    @pulumi.getter(name="hostnameLimit")
    def hostname_limit(self) -> int:
        """
        Limit of the number of hostname table entries (0 - 50000).
        """
        return pulumi.get(self, "hostname_limit")

    @property
    @pulumi.getter(name="hostnameTtl")
    def hostname_ttl(self) -> int:
        """
        TTL of hostname table entries (60 - 86400).
        """
        return pulumi.get(self, "hostname_ttl")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="sourceLocation")
    def source_location(self) -> str:
        """
        Enable/disable logging of source geographical location visibility.
        """
        return pulumi.get(self, "source_location")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetNetworkvisibilityResult(GetNetworkvisibilityResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkvisibilityResult(
            destination_hostname_visibility=self.destination_hostname_visibility,
            destination_location=self.destination_location,
            destination_visibility=self.destination_visibility,
            hostname_limit=self.hostname_limit,
            hostname_ttl=self.hostname_ttl,
            id=self.id,
            source_location=self.source_location,
            vdomparam=self.vdomparam)


def get_networkvisibility(vdomparam: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkvisibilityResult:
    """
    Use this data source to get information on fortios system networkvisibility


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:sys/getNetworkvisibility:getNetworkvisibility', __args__, opts=opts, typ=GetNetworkvisibilityResult).value

    return AwaitableGetNetworkvisibilityResult(
        destination_hostname_visibility=__ret__.destination_hostname_visibility,
        destination_location=__ret__.destination_location,
        destination_visibility=__ret__.destination_visibility,
        hostname_limit=__ret__.hostname_limit,
        hostname_ttl=__ret__.hostname_ttl,
        id=__ret__.id,
        source_location=__ret__.source_location,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_networkvisibility)
def get_networkvisibility_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNetworkvisibilityResult]:
    """
    Use this data source to get information on fortios system networkvisibility


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
