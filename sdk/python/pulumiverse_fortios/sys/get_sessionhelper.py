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
    'GetSessionhelperResult',
    'AwaitableGetSessionhelperResult',
    'get_sessionhelper',
    'get_sessionhelper_output',
]

@pulumi.output_type
class GetSessionhelperResult:
    """
    A collection of values returned by getSessionhelper.
    """
    def __init__(__self__, fosid=None, id=None, name=None, port=None, protocol=None, vdomparam=None):
        if fosid and not isinstance(fosid, int):
            raise TypeError("Expected argument 'fosid' to be a int")
        pulumi.set(__self__, "fosid", fosid)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if protocol and not isinstance(protocol, int):
            raise TypeError("Expected argument 'protocol' to be a int")
        pulumi.set(__self__, "protocol", protocol)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def fosid(self) -> int:
        """
        Session helper ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Helper name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        Protocol port.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> int:
        """
        Protocol number.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSessionhelperResult(GetSessionhelperResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSessionhelperResult(
            fosid=self.fosid,
            id=self.id,
            name=self.name,
            port=self.port,
            protocol=self.protocol,
            vdomparam=self.vdomparam)


def get_sessionhelper(fosid: Optional[int] = None,
                      vdomparam: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSessionhelperResult:
    """
    Use this data source to get information on an fortios system sessionhelper


    :param int fosid: Specify the fosid of the desired system sessionhelper.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['fosid'] = fosid
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:sys/getSessionhelper:getSessionhelper', __args__, opts=opts, typ=GetSessionhelperResult).value

    return AwaitableGetSessionhelperResult(
        fosid=__ret__.fosid,
        id=__ret__.id,
        name=__ret__.name,
        port=__ret__.port,
        protocol=__ret__.protocol,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_sessionhelper)
def get_sessionhelper_output(fosid: Optional[pulumi.Input[int]] = None,
                             vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSessionhelperResult]:
    """
    Use this data source to get information on an fortios system sessionhelper


    :param int fosid: Specify the fosid of the desired system sessionhelper.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
