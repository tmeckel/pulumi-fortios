# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetMulticastaddress6Result',
    'AwaitableGetMulticastaddress6Result',
    'get_multicastaddress6',
    'get_multicastaddress6_output',
]

@pulumi.output_type
class GetMulticastaddress6Result:
    """
    A collection of values returned by getMulticastaddress6.
    """
    def __init__(__self__, color=None, comment=None, id=None, ip6=None, name=None, taggings=None, vdomparam=None, visibility=None):
        if color and not isinstance(color, int):
            raise TypeError("Expected argument 'color' to be a int")
        pulumi.set(__self__, "color", color)
        if comment and not isinstance(comment, str):
            raise TypeError("Expected argument 'comment' to be a str")
        pulumi.set(__self__, "comment", comment)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip6 and not isinstance(ip6, str):
            raise TypeError("Expected argument 'ip6' to be a str")
        pulumi.set(__self__, "ip6", ip6)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if taggings and not isinstance(taggings, list):
            raise TypeError("Expected argument 'taggings' to be a list")
        pulumi.set(__self__, "taggings", taggings)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter
    def color(self) -> int:
        """
        Color of icon on the GUI.
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def comment(self) -> str:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip6(self) -> str:
        """
        IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
        """
        return pulumi.get(self, "ip6")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Tag name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def taggings(self) -> Sequence['outputs.GetMulticastaddress6TaggingResult']:
        """
        Config object tagging. The structure of `tagging` block is documented below.
        """
        return pulumi.get(self, "taggings")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def visibility(self) -> str:
        """
        Enable/disable visibility of the IPv6 multicast address on the GUI.
        """
        return pulumi.get(self, "visibility")


class AwaitableGetMulticastaddress6Result(GetMulticastaddress6Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMulticastaddress6Result(
            color=self.color,
            comment=self.comment,
            id=self.id,
            ip6=self.ip6,
            name=self.name,
            taggings=self.taggings,
            vdomparam=self.vdomparam,
            visibility=self.visibility)


def get_multicastaddress6(name: Optional[str] = None,
                          vdomparam: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMulticastaddress6Result:
    """
    Use this data source to get information on an fortios firewall multicastaddress6


    :param str name: Specify the name of the desired firewall multicastaddress6.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:firewall/getMulticastaddress6:getMulticastaddress6', __args__, opts=opts, typ=GetMulticastaddress6Result).value

    return AwaitableGetMulticastaddress6Result(
        color=__ret__.color,
        comment=__ret__.comment,
        id=__ret__.id,
        ip6=__ret__.ip6,
        name=__ret__.name,
        taggings=__ret__.taggings,
        vdomparam=__ret__.vdomparam,
        visibility=__ret__.visibility)


@_utilities.lift_output_func(get_multicastaddress6)
def get_multicastaddress6_output(name: Optional[pulumi.Input[str]] = None,
                                 vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMulticastaddress6Result]:
    """
    Use this data source to get information on an fortios firewall multicastaddress6


    :param str name: Specify the name of the desired firewall multicastaddress6.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
